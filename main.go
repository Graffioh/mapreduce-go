package main

import (
	"fmt"
	"wc-mapreduce-go/kv"
	"wc-mapreduce-go/mr"
	"wc-mapreduce-go/worker"
)

func main() {
	var intermediate []kv.KV
	kva_res_chan := make(chan []kv.KV)

	filenames := []string{"./input-file-1.txt", "./input-file-2.txt"}

	for _, f := range filenames {
		go worker.MapWorker(f, mr.Map, kva_res_chan)
	}

	for range filenames {
		kva_res := <-kva_res_chan
		intermediate = append(intermediate, kva_res...)
	}

	kv.SortKVA(intermediate)

	fmt.Printf("Intermediate: %v", intermediate)

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

			kv_reduced := kv.KV{Key: int.Key, Value: mr.Reduce(int.Key, values)}
			kva_result = append(kva_result, kv_reduced)
		}

		prev = int.Key
		i = j
	}

	fmt.Printf("\n\nFinal result: %v", kva_result)
}
