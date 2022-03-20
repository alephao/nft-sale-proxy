package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func fetchMetadata(url string) (map[string]interface{}, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	meta := map[string]interface{}{}
	err = json.Unmarshal(body, &meta)

	if err != nil {
		return nil, err
	}

	return meta, nil
}

func FetchMetadataERC721(config *Config, id int64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%d", config.BaseURL, id)
	return fetchMetadata(url)
}

func FetchMetadataERC1155(config *Config, id int64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%064x", config.BaseURL, id)
	return fetchMetadata(url)
}

func IncognitoMetadata(config *Config, id int64) map[string]interface{} {
	meta := map[string]interface{}{}

	meta["name"] = strings.ReplaceAll(config.IncognitoName, "{id}", strconv.FormatInt(id, 10))

	if config.IncognitoDescription != "" {
		meta["description"] = config.IncognitoDescription
	}

	if config.IncognitoImageURL != "" {
		meta["image"] = config.IncognitoImageURL
	}

	if config.IncognitoExternalLink != "" {
		meta["external_link"] = config.IncognitoExternalLink
	}

	return meta
}

// Returns the real metadata or incognito metadata depending on the configuration
func FetchMetdata(config *Config, id int64) (map[string]interface{}, error) {
	if id <= config.RevealUpTo {
		if config.IsERC1155 {
			return FetchMetadataERC1155(config, id)
		} else {
			return FetchMetadataERC721(config, id)
		}
	} else {
		return IncognitoMetadata(config, id), nil
	}
}
