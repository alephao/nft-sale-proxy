package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func fetchMetadata(httpClient *http.Client, url string) (map[string]interface{}, error) {
	resp, err := httpClient.Get(url)

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

func FetchMetadataERC721(config *Config, httpClient *http.Client, id int64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%d", config.BaseURL, id)
	return fetchMetadata(httpClient, url)
}

func FetchMetadataERC1155(config *Config, httpClient *http.Client, id int64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%s%064x", config.BaseURL, id)
	return fetchMetadata(httpClient, url)
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

func idInBoundaries(config *Config, id int64) bool {
	for _, boundaries := range config.OtherReveals {
		if boundaries[0] <= id && boundaries[1] >= id {
			return true
		}
	}
	return false
}

// Returns the real metadata or incognito metadata depending on the configuration
func FetchMetdata(httpClient *http.Client, config *Config, id int64) (map[string]interface{}, error) {

	if id <= config.RevealUpTo || idInBoundaries(config, id) {
		if config.IsERC1155 {
			return FetchMetadataERC1155(config, httpClient, id)
		} else {
			return FetchMetadataERC721(config, httpClient, id)
		}
	} else {
		return IncognitoMetadata(config, id), nil
	}
}
