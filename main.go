package main 

import (
	"blockchain/kvstore"
	"blockchain/trie"
	"fmt"
)

func main() {
	db := kvstore.NewLevelDB("./leveldb")
	state := trie.NewState(db, trie.EmptyHash)
	state.Store([]byte("apple"), []byte("apple"))
	state.Store([]byte("apply"), []byte("apply"))
	state.Store([]byte("application"), []byte("application"))
	state.Store([]byte("banana"), []byte("banana"))
	value, err := state.Load([]byte("apple"))
	fmt.Println(string(value), err)
}
