package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/darionyaphet/go.learning/thrift/gen-go/rpc"
	"os"
)

const address = "127.0.0.1:8989"

type RpcServiceImpl struct {
}

func (*RpcServiceImpl) Run(cxt context.Context, time int64, message string, pairs map[string]string) (v string, err error) {
	value := fmt.Sprintf("%d %s with %v", time, message, pairs)
	fmt.Println(value)
	return value, nil
}

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transport, err := thrift.NewTServerSocket(address)

	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := &RpcServiceImpl{}
	processor := rpc.NewRpcServiceProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	server.Serve()
}
