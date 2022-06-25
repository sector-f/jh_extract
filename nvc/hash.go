package nvc

import (
	"fmt"
	"hash/fnv"
)

// Hash is a 64-bit FNV-1a hash
type Hash uint64

func String2Hash(s string) Hash {
	hash := fnv.New64a()
	hash.Write([]byte(s))
	return Hash(hash.Sum64())
}

func (h Hash) String() string {
	return fmt.Sprintf("%016x", uint64(h))
}
