package tools

import (
	"fmt"

	"github.com/google/uuid"
)

func NewUuid() string {
	defer func() {
		if er := recover(); er != nil {
			fmt.Printf("Recovered after following error: %v\n", er)
		}
	}()

	u := uuid.NewString()

	return u
}
