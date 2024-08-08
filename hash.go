package main

import (
	"hash/fnv"
)

func hash(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}

func getIndex(key string, length int) int {
	return int(hash(key) % uint64(length))
}
