package mempool

import (
	cosmostx "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/types"
)

// isClobOrderTransaction returns true if the provided `tx` is a
// Cosmos transaction containing a `MsgPlaceOrder` or `MsgCancelOrder` message.
func IsClobOrderTransaction(
	tx types.Tx,
	mempoolLogger log.Logger,
) bool {
	cosmosTx := &cosmostx.Tx{}
	err := cosmosTx.Unmarshal(tx)
	if err != nil {
		mempoolLogger.Error("isClobOrderTransaction error. Invalid Cosmos Transaction.")
		return false
	}

	if len(cosmosTx.Body.Messages) == 1 &&
		(cosmosTx.Body.Messages[0].TypeUrl == "/dydxprotocol.clob.MsgPlaceOrder" ||
			cosmosTx.Body.Messages[0].TypeUrl == "/dydxprotocol.clob.MsgCancelOrder") {
		return true
	}

	return false
}
