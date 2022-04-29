package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/speps/go-hashids"
)

func main() {
	sign := "aaaaWTZ4QUJHMWxRT0wzOTpBOTU5OEI1NzlEMkU2MTRCNTM3RTc4QUZDNTkyRUIzQUE2MkI5MjQzOTkwNTY1OTg5MjFDMjEwNkZFQzI4QkE1OjE2NDc4NTI4NzY="
	bs, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(bs), ":")
	if len(parts) != 3 {
		panic("len 3")
	}

	fmt.Println(AppCode2AppID(parts[0]), string(bs))

}

// AppCode2AppID ...
func AppCode2AppID(appCode string) (appID int64) {

	if len(appCode) != 16 {
		appID, _ = strconv.ParseInt(appCode, 10, 64)
		return
	}

	hd := hashids.NewData()
	hd.Salt = "#sss#"

	h, _ := hashids.NewWithData(hd)
	result, _ := h.DecodeInt64WithError(appCode)

	if len(result) > 0 {
		appID = result[0]

	}

	return
}
