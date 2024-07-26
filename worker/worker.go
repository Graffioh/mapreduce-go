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
