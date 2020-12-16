package dashboard

import (
	"fmt"
	"net/http"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/chain"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/hashing"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/registry"
	"github.com/iotaledger/wasp/packages/state"
	"github.com/iotaledger/wasp/packages/vm/builtinvm/accountsc"
	"github.com/iotaledger/wasp/packages/vm/builtinvm/blob"
	"github.com/iotaledger/wasp/plugins/chains"
	"github.com/labstack/echo"
)

const chainRoute = "/chains/:chainid"
const chainTplName = "chain"

func addChainEndpoints(e *echo.Echo) {
	e.GET(chainRoute, func(c echo.Context) error {
		chainid, err := coretypes.NewChainIDFromBase58(c.Param("chainid"))
		if err != nil {
			return err
		}

		result := &ChainTemplateParams{
			BaseTemplateParams: BaseParams(c, chainListRoute),
			ChainID:            chainid,
		}

		result.ChainRecord, err = registry.GetChainRecord(&chainid)
		if err != nil {
			return err
		}

		if result.ChainRecord != nil && result.ChainRecord.Active {
			_, result.Block, _, err = state.LoadSolidState(&chainid)
			if err != nil {
				return err
			}

			chain := chains.GetChain(chainid)

			result.Committee.Size = chain.Size()
			result.Committee.Quorum = chain.Quorum()
			result.Committee.NumPeers = chain.NumPeers()
			result.Committee.HasQuorum = chain.HasQuorum()
			result.Committee.PeerStatus = chain.PeerStatus()
			result.RootInfo, err = fetchRootInfo(chain)
			if err != nil {
				return err
			}

			result.Accounts, err = fetchAccounts(chain)
			if err != nil {
				return err
			}

			result.TotalAssets, err = fetchTotalAssets(chain)
			if err != nil {
				return err
			}

			result.Blobs, err = fetchBlobs(chain)
			if err != nil {
				return err
			}
		}

		return c.Render(http.StatusOK, chainTplName, result)
	})
}

func fetchAccounts(chain chain.Chain) ([]coretypes.AgentID, error) {
	accounts, err := callView(chain, accountsc.Interface.Hname(), accountsc.FuncAccounts, nil)
	if err != nil {
		return nil, fmt.Errorf("accountsc view call failed: %v", err)
	}

	ret := make([]coretypes.AgentID, 0)
	for k := range accounts {
		agentid, _, err := codec.DecodeAgentID([]byte(k))
		if err != nil {
			return nil, err
		}
		ret = append(ret, agentid)
	}
	return ret, nil
}

func fetchTotalAssets(chain chain.Chain) (map[balance.Color]int64, error) {
	bal, err := callView(chain, accountsc.Interface.Hname(), accountsc.FuncTotalAssets, nil)
	if err != nil {
		return nil, err
	}
	return accountsc.DecodeBalances(bal)
}

func fetchBlobs(chain chain.Chain) (map[hashing.HashValue]uint32, error) {
	ret, err := callView(chain, blob.Interface.Hname(), blob.FuncListBlobs, nil)
	if err != nil {
		return nil, err
	}
	return blob.DecodeDirectory(ret)
}

type ChainTemplateParams struct {
	BaseTemplateParams

	ChainID coretypes.ChainID

	ChainRecord *registry.ChainRecord
	Block       state.Block
	RootInfo    RootInfo
	Accounts    []coretypes.AgentID
	TotalAssets map[balance.Color]int64
	Blobs       map[hashing.HashValue]uint32
	Committee   struct {
		Size       uint16
		Quorum     uint16
		NumPeers   uint16
		HasQuorum  bool
		PeerStatus []*chain.PeerStatus
	}
}

const tplChain = `
{{define "title"}}Chain details{{end}}

{{define "body"}}
	{{ $chainid := .ChainID }}
	<h2>Chain <tt>{{printf "%.8s" $chainid}}…</tt></h2>

	{{if .ChainRecord}}
		<div>
			<p>ChainID: <code>{{.ChainRecord.ChainID}}</code></p>
			<p>Chain Address: {{template "address" .ChainRecord.ChainID.Address}}</p>
			<p>Chain Color: <code>{{.ChainRecord.Color}}</code></p>
			<p>Active: <code>{{.ChainRecord.Active}}</code></p>
			{{if .ChainRecord.Active}}
				<p>Owner ID: {{template "agentid" (args $chainid .RootInfo.OwnerID)}}</p>
				<p>Description: <code>{{quoted 50 .RootInfo.Description}}</code></p>
			{{end}}
		</div>
		{{if .ChainRecord.Active}}
			<hr/>
			<div>
				<h3>Contracts</h3>
				<table>
					<thead>
						<tr>
							<th>Name</th>
							<th>Description</th>
							<th>Program Hash</th>
						</tr>
					</thead>
					<tbody>
					{{range $_, $c := .RootInfo.Contracts}}
						<tr>
							<td><code>{{quoted 30 $c.Name}}</code></td>
							<td><code>{{quoted 50 $c.Description}}</code></td>
							<td><code>{{$c.ProgramHash.Short}}</code></td>
						</tr>
					{{end}}
					</tbody>
				</table>
			</div>

			<hr/>
			<div>
				<h3>On-chain accounts</h3>
				<table>
					<thead>
						<tr>
							<th>AgentID</th>
						</tr>
					</thead>
					<tbody>
					{{range $_, $agentid := .Accounts}}
						<tr>
							<td>{{template "agentid" (args $chainid $agentid)}}</td>
						</tr>
					{{end}}
					</tbody>
				</table>
				<h4>Total assets</h4>
				{{ template "balances" .TotalAssets }}
			</div>

			<hr/>
			<div>
				<h3>Blobs</h3>
				<table>
					<thead>
						<tr>
							<th>Hash</th>
							<th>Size (bytes)</th>
							<th></th>
						</tr>
					</thead>
					<tbody>
					{{range $hash, $size := .Blobs}}
						<tr>
							<td><code>{{ hashref $hash }}</code></td>
							<td>{{ $size }}</td>
							<td><a href="/chains/{{$chainid}}/blob/{{hashref $hash}}">Details</a></td>
						</tr>
					{{end}}
					</tbody>
				</table>
			</div>

			<hr/>
			<div>
				<h3>Block</h3>
				<p>State index: <code>{{.Block.StateIndex}}</code></p>
				<p>State Transaction ID: <code>{{.Block.StateTransactionID}}</code></p>
				<p>Timestamp: <code>{{formatTimestamp .Block.Timestamp}}</code></p>
				<p>Essence Hash: <code>{{.Block.EssenceHash}}</code></p>
				<div>
					<table>
						<caption>Requests</caption>
						<thead>
							<tr>
								<th>RequestID</th>
							</tr>
						</thead>
						<tbody>
						{{range $_, $reqId := .Block.RequestIDs}}
							<tr>
								<td><code>{{$reqId}}</code></td>
							</tr>
						{{end}}
						</tbody>
					</table>
				</div>
			</div>

			<hr/>
			<div>
				<h3>Committee</h3>
				<p>Size:           <code>{{.Committee.Size}}</code></p>
				<p>Quorum:         <code>{{.Committee.Quorum}}</code></p>
				<p>NumPeers:       <code>{{.Committee.NumPeers}}</code></p>
				<p>HasQuorum:      <code>{{.Committee.HasQuorum}}</code></p>
				<table>
				<caption>Peer status</caption>
				<thead>
					<tr>
						<th>Index</th>
						<th>ID</th>
						<th>Status</th>
					</tr>
				</thead>
				<tbody>
				{{range $_, $s := .Committee.PeerStatus}}
					<tr>
						<td>{{$s.Index}}</td>
						<td><code>{{$s.PeeringID}}</code></td>
						<td>{{if $s.Connected}}up{{else}}down{{end}}</td>
					</tr>
				{{end}}
				</tbody>
				</table>
			</div>
		{{end}}
	{{else}}
		<p>No chain record for ID <code>{{$chainid}}</code></p>
	{{end}}
{{end}}
`
