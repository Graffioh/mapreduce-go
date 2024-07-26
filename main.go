package main

import (
	"fmt"
	"log"
	"os"
	"wc-mapreduce-go/kv"
	"wc-mapreduce-go/mr"
)

func main() {
	var intermediate []kv.KV

	filenames := []string{"./input-file-1.txt", "./input-file-2.txt"}

	for _, f := range filenames {
		d, err := os.ReadFile(f)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("File content:\n %v\n", string(d))

		content := string(d)

		intermediate = append(intermediate, mr.Map(f, content)...)
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
