package main

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	for i := 0; i < 1000; i++ {
		id, err := gonanoid.New(16)
		fmt.Println(id, err)
	}
}
