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

	NumberOfTokens int
	RevealUpTo     int
}

func NewConfigFromEnv() *Config {
	numberOfTokens, _ := strconv.Atoi(os.Getenv("NUMBER_OF_TOKENS"))
	revealUpTo, _ := strconv.Atoi(os.Getenv("REVEAL_UP_TO"))

	return &Config{
		BaseURL:               os.Getenv("BASE_URL"),
		IncognitoImageURL:     os.Getenv("INCOGNITO_IMAGE_URL"),
		IncognitoName:         os.Getenv("INCOGNITO_NAME"),
		IncognitoDescription:  os.Getenv("INCOGNITO_DESCRIPTION"),
		IncognitoExternalLink: os.Getenv("INCOGNITO_EXTERNAL_LINK"),
		NumberOfTokens:        numberOfTokens,
		RevealUpTo:            revealUpTo,
	}
}
