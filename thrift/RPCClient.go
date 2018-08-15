package main

import (
	"context"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/darionyaphet/go.learning/thrift/gen-go/rpc"
	"os"
	"time"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	t, _ := thrift.NewTSocket("127.0.0.1:8989")
	transport, _ := transportFactory.GetTransport(t)

	client := rpc.NewRpcServiceClientFactory(transport, protocolFactory)
	if err := t.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:19090", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	paramMap := make(map[string]string)
	paramMap["first name"] = "darion"
	paramMap["last name"] = "yaphet"
	result, _ := client.Run(context.Background(), int64(time.Now().Unix()), "hello", paramMap)
	fmt.Println("result", result)
}
