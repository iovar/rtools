package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JSONMinify(text string) string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	var res = new(bytes.Buffer)

	err := json.Compact(res, []byte(text))

	if err != nil {
		panic(err)
	}

	return res.String()
}

func JSONBeautify(text string) string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	dt := make(map[string]any)
	json.Unmarshal([]byte(text), &dt)

	res, err := json.MarshalIndent(dt, "", "    ")

	if err != nil {
		panic(err)
	}

	return string(res)
}
