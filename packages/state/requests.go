package state

import (
	"github.com/iotaledger/wasp/packages/database"
	"github.com/iotaledger/wasp/packages/sctransaction"
)

func MarkRequestsProcessed(reqids []*sctransaction.RequestId) error {
	dbase, err := database.GetProcessedRequestsDB()
	if err != nil {
		return err
	}
	for _, rid := range reqids {
		err := dbase.Set(database.Entry{
			Key: database.DbKeyProcessedRequest(rid),
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func IsRequestProcessed(reqid *sctransaction.RequestId) (bool, error) {
	dbase, err := database.GetProcessedRequestsDB()
	if err != nil {
		return false, err
	}
	contains, err := dbase.Contains(database.DbKeyProcessedRequest(reqid))
	if err != nil {
		return false, err
	}
	return contains, nil
}