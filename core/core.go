package core

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type Attribute struct {
	Value     string `json:"value"`
	TraitType string `json:"trait_type"`
}

type Metadata struct {
	Name         string      `json:"name,omitempty"`
	Description  string      `json:"description,omitempty"`
	Image        string      `json:"image,omitempty"`
  ImageUrl     string      `json:"image_url,omitempty"`
	AnimationUrl string      `json:"animation_url,omitempty"`
	ExternalLink string      `json:"external_link,omitempty"`
	Attributes   []Attribute `json:"attributes,omitempty"`
}

func fetchMetadata(url string) (*Metadata, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	meta := Metadata{}
	err = json.Unmarshal(body, &meta)

	if err != nil {
		return nil, err
	}

	return &meta, nil
}

func FetchMetadataERC721(config *Config, id int64) (*Metadata, error) {
	url := fmt.Sprintf("%s%d", config.BaseURL, id)
	return fetchMetadata(url)
}

func FetchMetadataERC1155(config *Config, id int64) (*Metadata, error) {
	url := fmt.Sprintf("%s%064x", config.BaseURL, id)
	return fetchMetadata(url)
}

func IncognitoMetadata(config *Config, id int64) *Metadata {
	return &Metadata{
		Name:         strings.ReplaceAll(config.IncognitoName, "{id}", strconv.FormatInt(id, 10)),
		Description:  config.IncognitoDescription,
		Image:        config.IncognitoImageURL,
		ExternalLink: config.IncognitoExternalLink,
		Attributes:   []Attribute{},
	}
}

// Returns the real metadata or incognito metadata depending on the configuration
func FetchMetdata(config *Config, id int64) (*Metadata, error) {
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
