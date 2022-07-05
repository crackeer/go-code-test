package main

import (
	"fmt"

	"github.com/kr/pretty"
)

func main() {
	type myType struct {
		a, b int
	}
	var x = []myType{{1, 2}, {3, 4}, {5, 6}}
	fmt.Printf("%# v", pretty.Formatter(x))
	pretty.Log("sggg", map[string]interface{}{
		"abc": 88,
	})
	fmt.Println(pretty.Diff(map[string]interface{}{
		"aaa": "ssssss",
	}, map[string]interface{}{
		"aaa": "sss",
	}))
}
