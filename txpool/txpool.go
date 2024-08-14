package txpool

import(
	"blockchain/types"
	"blockchain/utils/hash"
)



type TxPool struct {
	NewTx(tx *types.Transaction)
	Pop() *types.Transaction
	SetStatRoot(root hash.Hash)
	NotifyTxEvent(txs []*types.Transaction)
}