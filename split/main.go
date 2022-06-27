package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	bytes, _ := ioutil.ReadFile("demo.txt")
	result := split(string(bytes))
	aa, _ := json.Marshal(result)
	fmt.Println(string(aa))
}

func split(raw string) []string {
	retData := []string{}
	parts := strings.Split(raw, ",")
	for _, item := range parts {
		tmpParts := strings.Split(strings.TrimSpace(item), "\n")
		retData = append(retData, tmpParts...)
	}
	return retData
}
