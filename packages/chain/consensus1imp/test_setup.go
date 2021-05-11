package consensus1imp

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/iotaledger/wasp/packages/solo"

	"github.com/iotaledger/wasp/packages/testutil"

	"go.dedis.ch/kyber/v3"

	"github.com/iotaledger/wasp/packages/testutil/testpeers"

	"github.com/iotaledger/goshimmer/packages/ledgerstate"
	"github.com/iotaledger/goshimmer/packages/ledgerstate/utxodb"
	"github.com/iotaledger/goshimmer/packages/ledgerstate/utxoutil"
	"github.com/iotaledger/hive.go/crypto/ed25519"
	"github.com/iotaledger/hive.go/logger"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/chain/committeeimpl"
	"github.com/iotaledger/wasp/packages/chain/mempool"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/dbprovider"
	"github.com/iotaledger/wasp/packages/peering"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/testutil/testchain"
	"github.com/iotaledger/wasp/packages/testutil/testlogger"
	"github.com/stretchr/testify/require"
	"go.dedis.ch/kyber/v3/pairing"
	"go.uber.org/zap/zapcore"
)

type mockedEnv struct {
	Suite             *pairing.SuiteBn256
	T                 *testing.T
	N                 uint16
	Quorum            uint16
	Neighbors         []string
	Log               *logger.Logger
	Ledger            *utxodb.UtxoDB
	OriginatorKeyPair *ed25519.KeyPair
	OriginatorAddress ledgerstate.Address
	StateAddress      ledgerstate.Address
	PubKeys           []kyber.Point
	PrivKeys          []kyber.Scalar
	NetworkProviders  []peering.NetworkProvider
	DKSRegistries     []coretypes.DKShareRegistryProvider
	DB                *dbprovider.DBProvider
	SolidState        state.VirtualState
	StateReader       state.StateReader
	StateOutput       *ledgerstate.AliasOutput
	NodeConn          *testchain.MockedNodeConn
	MockedACS         *testchain.MockedAsynchronousCommonSubset
	ChainID           coretypes.ChainID
	mutex             sync.Mutex
	Nodes             []*mockedNode
	push              bool
}

type mockedNode struct {
	OwnIndex  uint16
	Env       *mockedEnv
	ChainCore *testchain.MockedChainCore
	Mempool   chain.Mempool
	Consensus *consensusImpl
	Log       *logger.Logger
}

func NewMockedEnv(t *testing.T, n, quorum uint16, debug bool) (*mockedEnv, *ledgerstate.Transaction) {
	level := zapcore.InfoLevel
	if debug {
		level = zapcore.DebugLevel
	}
	log := testlogger.WithLevel(testlogger.NewLogger(t, "04:05.000"), level, false)
	var err error

	log.Infof("creating test environment with N = %d, T = %d", n, quorum)

	neighbors := make([]string, n)
	for i := range neighbors {
		neighbors[i] = fmt.Sprintf("localhost:%d", 4000+i)
	}

	ret := &mockedEnv{
		Suite:     pairing.NewSuiteBn256(),
		T:         t,
		N:         n,
		Quorum:    quorum,
		Neighbors: neighbors,
		Log:       log,
		Ledger:    utxodb.New(),
		NodeConn:  testchain.NewMockedNodeConnection(),
		Nodes:     make([]*mockedNode, n),
	}
	ret.MockedACS = testchain.NewMockedACS(quorum, func(values [][]byte) {
		for _, n := range ret.Nodes {
			go n.ChainCore.ReceiveMessage(&chain.AsynchronousCommonSubsetMsg{
				ProposedBatchesBin: values,
			})
		}
	})

	_, ret.PubKeys, ret.PrivKeys = testpeers.SetupKeys(n, ret.Suite)
	ret.StateAddress, ret.DKSRegistries = testpeers.SetupDkg(t, quorum, neighbors, ret.PubKeys, ret.PrivKeys, ret.Suite, log.Named("dkg"))
	ret.NetworkProviders = testpeers.SetupNet(neighbors, ret.PubKeys, ret.PrivKeys, testutil.NewPeeringNetReliable(), log)

	ret.OriginatorKeyPair, ret.OriginatorAddress = ret.Ledger.NewKeyPairByIndex(0)
	_, err = ret.Ledger.RequestFunds(ret.OriginatorAddress)
	require.NoError(t, err)

	outputs := ret.Ledger.GetAddressOutputs(ret.OriginatorAddress)
	require.True(t, len(outputs) == 1)

	bals := map[ledgerstate.Color]uint64{ledgerstate.ColorIOTA: 100}
	txBuilder := utxoutil.NewBuilder(outputs...)
	err = txBuilder.AddNewAliasMint(bals, ret.StateAddress, state.OriginStateHash().Bytes())
	require.NoError(t, err)
	err = txBuilder.AddRemainderOutputIfNeeded(ret.OriginatorAddress, nil)
	require.NoError(t, err)
	originTx, err := txBuilder.BuildWithED25519(ret.OriginatorKeyPair)
	require.NoError(t, err)
	err = ret.Ledger.AddTransaction(originTx)
	require.NoError(t, err)

	ret.StateOutput, err = utxoutil.GetSingleChainedAliasOutput(originTx)
	require.NoError(t, err)

	ret.ChainID = *coretypes.NewChainID(ret.StateOutput.GetAliasAddress())

	ret.DB = dbprovider.NewInMemoryDBProvider(log)
	ret.SolidState, err = state.CreateOriginState(ret.DB, &ret.ChainID)
	require.NoError(t, err)
	ret.StateReader, err = state.NewStateReader(ret.DB, &ret.ChainID)
	require.NoError(t, err)

	ret.NodeConn.OnPostTransaction(func(tx *ledgerstate.Transaction) {
		_, exists := ret.Ledger.GetTransaction(tx.ID())
		if exists {
			ret.Log.Debugf("posted repeating originTx: %s", tx.ID().Base58())
			return
		}
		if err := ret.Ledger.AddTransaction(tx); err != nil {
			ret.Log.Error(err)
			return
		}

		ret.Log.Infof("mockedEnv: posted transaction to ledger: %s", tx.ID().Base58())
	})
	pullBacklogOutputClosure := func(addr *ledgerstate.AliasAddress) {}
	ret.NodeConn.OnPullBacklog(pullBacklogOutputClosure)

	for i := range ret.Nodes {
		ret.Nodes[i] = ret.newNode(uint16(i))
	}

	return ret, originTx
}

func (env *mockedEnv) newNode(i uint16) *mockedNode {
	log := env.Log.Named(fmt.Sprintf("%d", i))
	chainCore := testchain.NewMockedChainCore(env.ChainID, log)
	mpool := mempool.New(env.StateReader, coretypes.NewInMemoryBlobCache(), log)
	mockCommitteeRegistry := testchain.NewMockedCommitteeRegistry(env.Neighbors)
	cfg, err := peering.NewStaticPeerNetworkConfigProvider(env.Neighbors[i], 4000+int(i), env.Neighbors...)
	require.NoError(env.T, err)

	committee, err := committeeimpl.NewCommittee(env.StateAddress, env.NetworkProviders[i], cfg, env.DKSRegistries[i], mockCommitteeRegistry, log)
	require.NoError(env.T, err)

	ret := &mockedNode{
		OwnIndex:  i,
		Env:       env,
		ChainCore: chainCore,
		Mempool:   mpool,
		Consensus: New(chainCore, mpool, committee, testchain.NewMockedNodeConnection(), log),
		Log:       log,
	}
	ret.Consensus.mockedACS = env.MockedACS
	chainCore.OnReceiveMessage(func(msg interface{}) {
		switch msg := msg.(type) {
		case *chain.AsynchronousCommonSubsetMsg:
			ret.Consensus.EventAsynchronousCommonSubsetMsg(msg)
		}
	})

	return ret
}

func (env *mockedEnv) StartTimers() {
	for _, n := range env.Nodes {
		n.StartTimer()
	}
}

func (n *mockedNode) StartTimer() {
	n.Log.Infof("started timer..")
	go func() {
		counter := 0
		for {
			n.Consensus.EventTimerMsg(chain.TimerTick(counter))
			counter++
			time.Sleep(50 * time.Millisecond)
		}
	}()
}

func (n *mockedNode) WaitTimerTick(until int) {
	for n.Consensus.getTimerTick() < until {
		time.Sleep(10 * time.Millisecond)
	}
}

func (env *mockedEnv) WaitTimerTick(until int) {
	var wg sync.WaitGroup
	wg.Add(int(env.N))
	for _, n := range env.Nodes {
		go func() {
			n.WaitTimerTick(until)
			wg.Done()
		}()
	}
	wg.Wait()
	env.Log.Infof("target timer tick #%d has been reached", until)
}

func (n *mockedNode) WaitStateIndex(until uint32) {
	for n.Consensus.getStateIndex() < until {
		time.Sleep(10 * time.Millisecond)
	}
}

func (env *mockedEnv) WaitStateIndex(until uint32) {
	var wg sync.WaitGroup
	wg.Add(int(env.N))
	for _, n := range env.Nodes {
		go func() {
			n.WaitStateIndex(until)
			wg.Done()
		}()
	}
	wg.Wait()
	env.Log.Infof("target state index #%d has been reached", until)
}

func (env *mockedEnv) eventStateTransition() {
	nowis := time.Now()
	for _, node := range env.Nodes {
		go node.Consensus.EventStateTransitionMsg(&chain.StateTransitionMsg{
			State:          env.SolidState.Clone(),
			StateOutput:    env.StateOutput,
			StateTimestamp: nowis,
		})
	}
}

func (env *mockedEnv) postDummyRequest() {
	req := solo.NewCallParams("dummy", "dummy").
		NewRequestOffLedger(env.OriginatorKeyPair)
	for _, n := range env.Nodes {
		go n.Mempool.ReceiveRequest(req)
	}
}
