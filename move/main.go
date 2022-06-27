package main

import (
	"fmt"
	"strconv"
)

func main() {
	var renderID int64 = 11
	oldBasicID := renderID & 0xFF
	oldRenderID := renderID >> 1

	fmt.Println(oldBasicID, oldRenderID)
	fmt.Println(strconv.FormatInt(renderID, 2))
}
