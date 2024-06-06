package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
)

type LogDoc struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func PerfSingleInsert() {
	var channels [nCli]chan int
	for i := range channels {
		channels[i] = make(chan int)
	}

	for i := 0; i < nCli; i++ {
		go func(idx int) {
			log.Printf("[Client: %d]: Start\n", idx)
			ctx := context.Background()
			es, err := NewTypedCli()
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

	log.Println("Done")
}

func PerfBulkInsert() {
	es, err := NewCli()
	if err != nil {
		log.Printf("xxx Connect failed: %v\n", err)
	} else {
		var buf bytes.Buffer
		for j := 0; j < nReq; j++ {
			doc := LogDoc{ID: j, Name: "cli[0]", Content: fmt.Sprintf("Number %d", j)}
			//meta := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%d" } }%s`, doc.ID, "\n"))
			data, _ := json.Marshal(doc)
			data = append(data, "\n"...)
			//buf.Grow(len(meta) + len(data))
			buf.Grow(len(data))
			//buf.Write(meta)
			buf.Write(data)
		}

		indexName := "myclient-0"
		_, err := es.Bulk(bytes.NewReader(buf.Bytes()), es.Bulk.WithIndex(indexName))
		if err != nil {
			log.Fatalf("Failure indexing %s", err)
		} else {
			log.Printf("[Success]Bulk insert.\n")
		}
	}
}
