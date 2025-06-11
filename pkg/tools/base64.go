package tools

import (
	"encoding/base64"
	"fmt"
)

func Base64Encode(text string) string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	res := base64.StdEncoding.EncodeToString([]byte(text))

	return res
}

func Base64Decode(text string) string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	res, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic(err)
	}

	return string(res)
}
