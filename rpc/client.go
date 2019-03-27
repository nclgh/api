package rpc

import (
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/clients"
)

var (
	passportCli *client.RpcClient
	memberCli   *client.RpcClient
	deviceCli   *client.RpcClient
)

func Init() {
	passportCli = clients.GetRpcClient("passport")
	memberCli = clients.GetRpcClient("member")
	deviceCli = clients.GetRpcClient("device")
}

func GetPassportClient() *client.RpcClient {
	return passportCli
}

func GetMemberClient() *client.RpcClient {
	return memberCli
}

func GetDeviceClient() *client.RpcClient {
	return deviceCli
}
