package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	println("hello milvus tool")

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
		// Endpogiints: []string{"localhost:2379", "localhost:22379", "localhost:32379"}
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Printf("Initialize etcd client failed. err: %v\n", err)
	}
	defer cli.Close()

	resp, err := cli.Get(context.TODO(), "/y-dev/", clientv3.WithPrefix())
	if err != nil {
		fmt.Printf("Fail to get from etcd: %v\n", err)
	}
	for _, kv := range resp.Kvs {
		fmt.Printf("%s -> %s", string(kv.Key), string(kv.Value))
	}

}
