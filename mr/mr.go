package mr

import (
	"strconv"
	"strings"
	"wc-mapreduce-seq-go/kv"
)

// map(k1, v1) -> list(k2, v2)
func Map(filename string, content string) []kv.KV {
	words := strings.Fields(content)
	var kva []kv.KV

	for _, w := range words {
		kva = append(kva, kv.KV{Key: w, Value: "1"})
	}

	kv.SortKVA(kva)

	return kva
}

// reduce(k2, list(v2)) -> list(v2)
func Reduce(key string, values []string) string {
	result := 0

	for range values {
		result++
	}

	return strconv.Itoa(result)
}
