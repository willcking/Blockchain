package types

type Address [20]byte

func PubKeyToAddress(pub []byte) Address {
	var addr Address
	hash := sha3.Keccak256(pub)
	address := hash[12:]
	return addr
}