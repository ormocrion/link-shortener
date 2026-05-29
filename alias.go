package main

import (
	"math/rand"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func checkLink(input string) string {
	if (strings.Contains(input, "http://") || strings.Contains(input, "https://")) {
		return input
	}
		return "http://" + input
} 

func generateRandomString(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func redirectLinkFormer(alias string) string {
	aliasLink := "localhost/url/" + alias;
	
	return aliasLink;
}
