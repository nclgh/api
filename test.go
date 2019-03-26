package main

import (
	"fmt"
	"time"
	"sync"
	"net/rpc"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
	"os"
	"github.com/nclgh/lakawei_scaffold/rpc/common"
)

func test() {
	//t0()
	//t1()
	t2()
	os.Exit(0)
}

func t0() {
	fmt.Println("tag1")
	cli, err := rpc.Dial("tcp", "39.108.71.247:9701")
	if err != nil {
		panic(err)
	}
	fmt.Println("connect")

	req := passport.CreateSessionRequest{
		UserId: 666,
	}
	var res passport.CreateSessionResponse

	err = cli.Call("ServicePassport.CreateSession", req, &res)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.SessionId)

	var getSessionRes passport.GetSessionResponse
	err = cli.Call("ServicePassport.GetSession", passport.GetSessionRequest{
		SessionId: res.SessionId,
	}, &getSessionRes)
	if err != nil {
		panic(err)
	}
	fmt.Println(getSessionRes.UserId)
}

func t1() {
	passportCli, err := client.InitClient("ServicePassport")
	if err != nil {
		panic(err)
	}
	for {
		var wg sync.WaitGroup
		st := time.Now()
		for i := 0; i < 1; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req := passport.CreateSessionRequest{
					UserId: 666,
				}
				var res passport.CreateSessionResponse
				err := passportCli.Call(&client.RpcRequestCtx{}, "CreateSession", req, &res)
				if err != nil {
					fmt.Println(err)
					return
				}
				//_, ok := tres.(*passport.CreateSessionResponse)
				//if !ok {
				//	return
				//}
				//realRes := tres.(*passport.CreateSessionResponse)
				fmt.Println(res.SessionId)
			}()
		}
		wg.Wait()
		et := time.Now()
		sub := et.Sub(st).Nanoseconds()
		//fmt.Println(sub)
		if sub > 0 {

		}
	}
}

func t2() {
	passportCli, err := client.InitClient("ServicePassport")
	if err != nil {
		panic(err)
	}
	req := passport.CreateSessionRequest{
		UserId: 666,
	}
	var res passport.CreateSessionResponse
	err = passportCli.Call(&client.RpcRequestCtx{}, "CreateSession", req, &res)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res.SessionId)
	var resDeleteSession passport.DeleteSessionResponse
	err = passportCli.Call(&client.RpcRequestCtx{},"DeleteSession",passport.DeleteSessionRequest{UserId:666},&resDeleteSession)
	if err != nil {
		fmt.Println(err)
		return
	}
	if resDeleteSession.Code != common.CodeSuccess || !resDeleteSession.IsSuccess {
		fmt.Println(resDeleteSession.Msg)
	}
	fmt.Println("delete session success")
}
