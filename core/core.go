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
	ExternalLink string      `json:"external_link,omitempty"`
	Attributes   []Attribute `json:"attributes,omitempty"`
}

func FetchMetadataERC721(config *Config, id int) (*Metadata, error) {
	url := fmt.Sprintf("%s%d", config.BaseURL, id)
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

func IncognitoMetadata(config *Config, id int) *Metadata {
	return &Metadata{
		Name:         strings.ReplaceAll(config.IncognitoName, "{id}", strconv.Itoa(id)),
		Description:  config.IncognitoDescription,
		Image:        config.IncognitoImageURL,
		ExternalLink: config.IncognitoExternalLink,
		Attributes:   []Attribute{},
	}
}

// Returns the real metadata or incognito metadata depending on the configuration
func FetchMetdata(config *Config, id int) (*Metadata, error) {
	if id <= config.RevealUpTo {
		return FetchMetadataERC721(config, id)
	} else {
		return IncognitoMetadata(config, id), nil
	}
}
