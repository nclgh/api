package clients

import (
	"fmt"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/conf"
)

var (
	clients = make(map[string]*client.RpcClient)
)

func GetRpcClient(service string) *client.RpcClient {
	c := conf.GetConfig()
	serviceName := c.DefaultString(fmt.Sprintf("%sServiceName", service), "")
	if serviceName == "" {
		panic(fmt.Sprintf("%s serviceName not found", service))
	}

	if cli, ok := clients[serviceName]; ok {
		return cli
	}
	cli, err := client.InitClient(serviceName)
	if err != nil {
		panic(err)
	}
	clients[serviceName] = cli
	return cli
}
