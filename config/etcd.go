package config

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func LoadConfigFromEtcd(cfg *Config) error {
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.EtcdConfig.Endpoints,
		DialTimeout: 5 * time.Second,
		//TLS: tls.Config{}
		Username: cfg.EtcdConfig.Username,
		Password: cfg.EtcdConfig.Password,
	})
	if err != nil {
		return err
	}
	defer func(cli *clientv3.Client) {
		err := cli.Close()
		if err != nil {
			log.Printf("error closing etcd client")
		}
	}(etcdClient)

	// Fetch the configuration from etcd
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := etcdClient.Get(ctx, "config/application")
	if err != nil {
		return err
	}

	if len(resp.Kvs) == 0 {
		return fmt.Errorf("configuration not found in etcd")
	}

	// Load the configuration into Viper
	reader := bytes.NewReader(resp.Kvs[0].Value)
	if err := viper.ReadConfig(reader); err != nil {
		return err
	}

	if err := viper.Unmarshal(cfg); err != nil {
		return err
	}

	return nil
}
