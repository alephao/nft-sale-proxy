package core

import (
	"os"
	"strconv"
)

type Config struct {
	BaseURL string

	IncognitoImageURL     string
	IncognitoName         string
	IncognitoDescription  string
	IncognitoExternalLink string

	NumberOfTokens int64
	RevealUpTo     int64

	IsERC1155 bool
}

func NewConfigFromEnv() *Config {
	numberOfTokens, _ := strconv.ParseInt(os.Getenv("NUMBER_OF_TOKENS"), 10, 0)
	revealUpTo, _ := strconv.ParseInt(os.Getenv("REVEAL_UP_TO"), 10, 0)
	erc1155 := os.Getenv("ERC1155") == "true"

	return &Config{
		BaseURL:               os.Getenv("BASE_URL"),
		IncognitoImageURL:     os.Getenv("INCOGNITO_IMAGE_URL"),
		IncognitoName:         os.Getenv("INCOGNITO_NAME"),
		IncognitoDescription:  os.Getenv("INCOGNITO_DESCRIPTION"),
		IncognitoExternalLink: os.Getenv("INCOGNITO_EXTERNAL_LINK"),
		NumberOfTokens:        numberOfTokens,
		RevealUpTo:            revealUpTo,
		IsERC1155:             erc1155,
	}
}
