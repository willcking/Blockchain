package statdb

import (
	"blockchain/types"
	"hash"
)

type StatDB interface {
	SetStatRoot(root hash.Hash)
	Load(addr types.Address) *types.Account
	Store(addr types.Address, account types.Account)
}

func SetStatRoot(root hash.Hash) {
	// todo
}

func Load(addr types.Address) *types.Account {
	//todo
}

func Store(addr types.Address, account types.Account) {
	//todo
}
