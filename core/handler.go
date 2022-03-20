package core

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type GenericResponse struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func HandleRequest(
	config *Config,
	httpClient *http.Client,
	getTokenId func() string,
) (GenericResponse, error) {
	tokenId := getTokenId()
	if tokenId == "" {
		response := GenericResponse{
			StatusCode: 404,
			Body:       "\"message\": \"invalid token id\"",
		}
		return response, nil
	}

	var id int64
	var err error
	if config.IsERC1155 {
		id, err = strconv.ParseInt(tokenId, 16, 32)
	} else {
		id, err = strconv.ParseInt(tokenId, 10, 0)
	}
	if err != nil || id < 0 || id > config.NumberOfTokens {
		response := GenericResponse{
			StatusCode: 404,
			Body:       "\"message\": \"id out of bounds\"",
		}
		return response, nil
	}

	metadata, err := FetchMetdata(httpClient, config, id)

	if err != nil {
		response := GenericResponse{
			StatusCode: 500,
		}
		return response, fmt.Errorf("failed to get metadata")
	}

	body, err := json.Marshal(metadata)
	if err != nil {
		response := GenericResponse{
			StatusCode: 500,
		}
		return response, fmt.Errorf("failed to marshal metadata")
	}

	return GenericResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}, nil
}
