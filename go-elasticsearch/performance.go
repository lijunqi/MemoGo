package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	//"github.com/elastic/go-elasticsearch/v8/esutil"
)

type LogDoc struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func RunPerformance() {
	const nCli = 10
	const nReq = 100

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://10.24.11.116:9200",
		},
	}

	var channels [nCli]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}

	startTime := time.Now()
	for i := 0; i < nCli; i++ {
		go func(idx int) {
			log.Printf("[Client: %d]: Start\n", idx)
			ctx := context.Background()
			es, err := elasticsearch.NewTypedClient(cfg)
			if err != nil {
				log.Printf("xxx Connect failed: %v\n", err)
			} else {
				log.Printf("[Client %d]: Connected.\n", idx)
				indexName := fmt.Sprintf("myclient-%d", idx)
				cliName := fmt.Sprintf("Cli[%d]", idx)
				es.Indices.Create(indexName).Do(ctx)
				for j := 0; j < nReq; j++ {
					log.Printf("[Client %d]: %d\n", idx, j)
					doc := LogDoc{Name: cliName, Content: fmt.Sprintf("Number %d", j)}
					res, err := es.Index(indexName).Request(doc).Do(ctx)
					if err != nil {
						log.Printf("xxx %s create doc: %v\n", cliName, err)
					} else {
						log.Printf("[Client %d]: %v\n", idx, res)
					}
					//time.Sleep(200 * time.Millisecond)
				}
			}
			channels[idx] <- idx
			log.Printf("[Client: %d]: Done\n", idx)
		}(i)
	}

	for i := range channels {
		<-channels[i]
	}

	elapstedTime := time.Since(startTime)

	log.Printf("Elapsted Time: %.2f sec\n", float64(elapstedTime)/float64(time.Second))
	log.Println("Done")
}
