package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type KV struct {
	key   string
	value string
}

func sortKVA(kva []KV) {
	sort.Slice(kva, func(i, j int) bool {
		return kva[i].key < kva[j].key
	})
}

// map(k1, v1) -> list(k2, v2)
func Map(filename string, content string) []KV {
	words := strings.Fields(content)
	var kva []KV

	for _, w := range words {
		kva = append(kva, KV{key: w, value: "1"})
	}

	sortKVA(kva)

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

func main() {
	filename := "./input-file.txt"

	d, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content:\n %v\n", string(d))

	content := string(d)

	intermediate := Map(filename, content)
	fmt.Printf("Intermediate: %v", intermediate)

	prev := ""
	var kva_result []KV

	for i, int := range intermediate {
		j := i

		for j < len(intermediate) && intermediate[j].key == intermediate[i].key {
			j++
		}

		if int.key != prev {
			values := []string{}
			for k := i; k < j; k++ {
				values = append(values, intermediate[k].value)
			}

			kv_reduced := KV{key: int.key, value: Reduce(int.key, values)}
			kva_result = append(kva_result, kv_reduced)
		}

		prev = int.key
		i = j
	}

	fmt.Printf("\n\nFinal result: %v", kva_result)
}
