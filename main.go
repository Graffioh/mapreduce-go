package main

import (
	"fmt"
	"log"
	"os"
	"wc-mapreduce-go/kv"
	"wc-mapreduce-go/mr"
	"wc-mapreduce-go/worker"
)

func main() {
	var intermediate []kv.KV

	map_res_chan := make(chan []kv.KV)
	reduce_res_chan := make(chan []kv.KV)

	filenames := []string{"./input-file-1.txt", "./input-file-2.txt"}

	// Create map worker to generate intermediate
	for _, f := range filenames {
		go worker.MapWorker(f, mr.Map, map_res_chan)
	}

	// Merge map worker intermediate
	for range filenames {
		kva_res := <-map_res_chan
		intermediate = append(intermediate, kva_res...)
	}

	kv.SortKVA(intermediate)

	fmt.Printf("Intermediate: %v", intermediate)

	// Reduce the intermediate
	go worker.ReduceWorker(intermediate, mr.Reduce, reduce_res_chan)

	kva_result := <-reduce_res_chan

	fmt.Printf("\n\nFinal result: %v", kva_result)

	d1 := []byte(fmt.Sprintf("%v\n", kva_result))
	err := os.WriteFile("./output-file-1.txt", d1, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
