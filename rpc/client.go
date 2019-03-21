package rpc

import (
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/clients"
)

var (
	passportCli *client.RpcClient
)

func Init() {
	passportCli = clients.GetRpcClient("passport")
}


func GetPassportClient() *client.RpcClient {
	return passportCli
}