package kv

import "sort"

type KV struct {
	Key   string
	Value string
}

func SortKVA(kva []KV) {
	sort.Slice(kva, func(i, j int) bool {
		return kva[i].Key < kva[j].Key
	})
}
