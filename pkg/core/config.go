package core

import (
	"os"
	"strconv"
	"strings"
)

type Config struct {
	BaseURL string

	IncognitoImageURL     string
	IncognitoName         string
	IncognitoDescription  string
	IncognitoExternalLink string

	NumberOfTokens int64
	RevealUpTo     int64
	OtherReveals   [][2]int64

	IsERC1155 bool
}

func GetOtherReveals(otherRevealsString string) [][2]int64 {
	if otherRevealsString == "" {
		return [][2]int64{}
	}
	ranges := strings.Split(otherRevealsString, ",")
	otherReveals := [][2]int64{}
	for _, r := range ranges {
		components := strings.Split(r, "-")
		// TODO: Handle those errors
		from, _ := strconv.ParseInt(components[0], 10, 64)
		to, _ := strconv.ParseInt(components[1], 10, 64)
		otherReveals = append(otherReveals, [2]int64{from, to})
	}
	return otherReveals
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
		OtherReveals:          GetOtherReveals(os.Getenv("OTHER_REVEALS")),
		IsERC1155:             erc1155,
	}
}
