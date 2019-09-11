package main

import (
	"github.com/hzjiangjian/go-stack-test/src/algorithm"
	"github.com/hzjiangjian/go-stack-test/src/tgrpc"
	"time"
)

func main(){
	Run()
	return
	algorithm.PrintPrime()
	return
	go func(){
		for{
			time.Sleep(time.Second)
			tgrpc.RunClient()
		}
	}()

	tgrpc.RunServer()

}
