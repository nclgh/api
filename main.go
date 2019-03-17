package main

import (
	"fmt"
	"github.com/nclgh/lakawei_scaffold/kite/passport"
	"github.com/nclgh/lakawei_scaffold/kite/kite_common"
)

func main() {
	passportCli := kite_common.GetServerClient("passport")

	//req := &passport.CreateSessionRequest{
	//	UserId: 123,
	//}
	//rsp, _ := passportCli.Tell("CreateSession", req)
	//realRsp := &passport.CreateSessionResponse{}

	req := &passport.GetSessionRequest{
		SessionId: "hhh",
	}
	rsp, _ := passportCli.Tell("GetSession", req)
	realRsp := &passport.GetSessionResponse{}

	err := rsp.Unmarshal(realRsp)
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		fmt.Printf("%v\n", *realRsp)
	}
}
