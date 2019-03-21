package main

import (
	"fmt"
	"github.com/nclgh/lakawei_rpc/client"
	"github.com/nclgh/lakawei_scaffold/rpc/passport"
	"time"
	"sync"
	"net/rpc"
)

func test() {
	t0()
	//t1()
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
	passportCli := client.RpcClient{}
	passportCli.Init("ServicePassport")
	for {
		var wg sync.WaitGroup
		st := time.Now()
		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				req := passport.CreateSessionRequest{
					UserId: 666,
				}
				var res passport.CreateSessionResponse
				tres, err := passportCli.Call(&client.RpcRequestCtx{}, "CreateSession", req, &res)
				if err != nil {
					fmt.Println(err)
					return
				}
				_,ok := tres.(*passport.CreateSessionResponse)
				if !ok {
					return
				}
				//realRes := tres.(*passport.CreateSessionResponse)
				//fmt.Println(realRes.SessionId)
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