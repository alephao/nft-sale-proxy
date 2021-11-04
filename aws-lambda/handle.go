package nft_proxy_lambda

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	core "github.com/alephao/nft-sale-proxy/core"
	"github.com/aws/aws-lambda-go/events"
)

type Request events.APIGatewayProxyRequest
type Response events.APIGatewayProxyResponse

func HandleRequest(ctx context.Context, request Request) (Response, error) {
	config := core.NewConfigFromEnv()

	tokenId := request.PathParameters["tokenId"]
	if tokenId == "" {
		response := Response{
			StatusCode: 404,
			Body:       "\"message\": \"invalid token id\"",
		}
		return response, nil
	}

	id, err := strconv.Atoi(tokenId)
	if err != nil || id < 0 || id > config.NumberOfTokens {
		response := Response{
			StatusCode: 404,
			Body:       "\"message\": \"id out of bounds\"",
		}
		return response, nil
	}

	metadata, err := core.FetchMetdata(config, id)

	if err != nil {
		response := Response{
			StatusCode: 500,
		}
		return response, fmt.Errorf("failed to get metadata")
	}

	body, err := json.Marshal(metadata)
	if err != nil {
		response := Response{
			StatusCode: 500,
		}
		return response, fmt.Errorf("failed to marshal metadata")
	}

	return Response{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
