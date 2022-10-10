package main

import (
	proxy "github.com/alephao/nft-sale-proxy/pkg/aws-lambda"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(proxy.HandleRequest)
}
