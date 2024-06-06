package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
)

const SERVER = "http://10.24.11.116:9200"
const nCli = 1
const nReq = 1000

func NewCli() (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{SERVER},
	}
	return elasticsearch.NewClient(cfg)
}

func NewTypedCli() (*elasticsearch.TypedClient, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{SERVER},
	}
	return elasticsearch.NewTypedClient(cfg)
}

func Search(index string) {
	es, err := NewCli()
	if err != nil {
		log.Printf("xxx Connect failed: %v\n", err)
		return
	}

	query := `{ "query": { "match_all": {} } }`
	res, err := es.Search(
		es.Search.WithIndex(index),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Printf("xxx Search error: %v\n", err)
		return
	}

	log.Println("Search result:")
	log.Printf("Status code: %d\n", res.StatusCode)
	log.Printf("Header: %v\n", res.Header)
	//log.Printf("Body: %v\n", res.Body)

	buf := new(strings.Builder)
	n, err := io.Copy(buf, res.Body)
	if err != nil {
		log.Printf("xxx Read body: %v\n", err)
	} else {
		log.Printf("Read %d bytes\n", n)
	}
	log.Printf("Body: %s\n", buf.String())

	/*
		body, err := io.ReadAll(res.Body)
		if err != nil {
			log.Printf("xxx ReadAll error: %v\n", err)
			return
		}
		var jsonMap map[string]interface{}
		json.Unmarshal(body, &jsonMap)
		log.Printf("Body obj: %v\n", jsonMap)
	*/

	/*
		out := new(bytes.Buffer)
		b1 := bytes.NewBuffer([]byte{})
		b2 := bytes.NewBuffer([]byte{})
		tr := io.TeeReader(res.Body, b1)
		if _, err := io.Copy(b2, tr); err != nil {
			out.WriteString(fmt.Sprintf("<error reading response body: %v>", err))
			return
		}
		out.ReadFrom(b2) // errcheck exclude (*bytes.Buffer).ReadFrom
		log.Printf("Body: %s\n", out)
	*/
}

func DeleteIndex(idxList []string) {
	es, err := NewTypedCli()
	if err != nil {
		log.Printf("xxx Connect failed: %v\n", err)
	} else {
		for _, idxName := range idxList {
			_, err := es.Indices.Delete(idxName).Do(context.TODO())
			if err != nil {
				log.Printf("[Failed]Delete index[%s]: %v\n", idxName, err)
			} else {
				log.Printf("[Success]Delete index[%s].\n", idxName)
			}
		}
	}
}

func main() {
	log.Println("====== Start ======")
	startTime := time.Now()

	// ~ ES Search
	// indexName := "test-index"
	// Search(indexName)

	// ~ Single Insert
	//PerfSingleInsert()

	// ~ Bulk Insert
	//PerfBulkInsert()

	// ~ Delete Index
	del := 1
	if del == 1 {
		indexList := []string{}
		for i := 0; i < nCli; i++ {
			indexList = append(indexList, fmt.Sprintf("myclient-%d", i))
		}
		DeleteIndex(indexList)
	}

	elapstedTime := time.Since(startTime)
	log.Printf("[Done]Elapsted Time: %.3f sec\n", float64(elapstedTime)/float64(time.Second))
}
