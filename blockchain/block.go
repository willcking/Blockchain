package blockchain

import (
	"blockchain/crypto/sha3"
	"blockchain/trie"
	"blockchain/txpool"
	"blockchain/types"
	"blockchain/utils/hash"
	"blockchain/utils/rlp"
)

type Header struct {
	Root       hash.Hash
	ParentHash hash.Hash
	Height     uint64
	Coinbase   types.Address
	Timestamp  uint64
	
	Nonce uint64
}

type Body struct {
	Transactions []types.Transaction
	Receiptions  []types.Receiption
}

type Blockchain struct {
	CurrentHeader Header
	statedb trie.ITrie
	Txpool  txpool.TxPool
}