package milvus_meta

import clientv3 "go.etcd.io/etcd/client/v3"

// EtcdKV implements TxnKV interface, it supports to process multiple kvs in a transaction.
type EtcdKV struct {
	client   *clientv3.Client
	rootPath string
}

// NewEtcdKV creates a new etcd kv.
func NewEtcdKV(client *clientv3.Client, rootPath string) *EtcdKV {
	kv := &EtcdKV{
		client:   client,
		rootPath: rootPath,
	}
	return kv
}
