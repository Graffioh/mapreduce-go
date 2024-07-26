package worker

import (
	"fmt"
	"log"
	"os"
	"wc-mapreduce-go/kv"
)

func MapWorker(filename string, mapf func(string, string) []kv.KV, res chan []kv.KV) {
	d, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("File content:\n %v\n", string(d))

	content := string(d)

	res <- mapf(filename, content)
}

func ReduceWorker(intermediate []kv.KV, reducef func(string, []string) string, res chan []kv.KV) {
	prev := ""
	var kva_result []kv.KV

	for i, int := range intermediate {
		j := i

		for j < len(intermediate) && intermediate[j].Key == intermediate[i].Key {
			j++
		}

		if int.Key != prev {
			values := []string{}
			for k := i; k < j; k++ {
				values = append(values, intermediate[k].Value)
			}

			kv_reduced := kv.KV{Key: int.Key, Value: reducef(int.Key, values)}
			kva_result = append(kva_result, kv_reduced)
		}

		prev = int.Key
		i = j
	}

	res <- kva_result
}
