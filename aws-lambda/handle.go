package nft_proxy_lambda

import (
	"context"

	core "github.com/alephao/nft-sale-proxy/core"
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	genericResponse, err := core.HandleRequest(func() string {
		return request.PathParameters["tokenId"]
	})

	return Response{
		StatusCode: genericResponse.StatusCode,
		Body:       genericResponse.Body,
		Headers:    genericResponse.Headers,
	}, err
}
