package main

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

//type EtcdObj struct {
//	hostIP string
//	port   string
//	cfg clientv3.Config
//}

func main() {
	config := clientv3.Config{
		//Endpoints:   []string{"localhost:2379"},
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	}

	cli, err := clientv3.New(config)
	if err != nil {
		// handle error!
		fmt.Printf("xxx Error new client: %v\n", err)
	}
	defer cli.Close()

	kv := clientv3.NewKV(cli)
	resp, err := kv.Get(context.Background(), "\x00", clientv3.WithFromKey())
	if err != nil {
		fmt.Printf("xxx Error get / : %v\n", err)
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}
}
