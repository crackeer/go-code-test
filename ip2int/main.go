package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/speps/go-hashids"
)

func main() {
	fmt.Println(StringIpToInt("172.16.1.1"))
	fmt.Println(StringIpToInt("172.16.2.9"))
	fmt.Println(IpIntToString(2886729986))

	traceID := GenTraceID("172.16.1.14")
	fmt.Println(traceID)
	list := DecodeTraceID(traceID)
	fmt.Println(list, IpIntToString(int(list[0])))
}

func DecodeTraceID(hostIP string) []int64 {
	hd := hashids.NewData()
	hd.Salt = "#trace-id-hashids-salt@realsee#"
	hd.MinLength = 32
	h, _ := hashids.NewWithData(hd)
	result, _ := h.DecodeInt64WithError(hostIP)
	return result

}

func GenTraceID(hostIP string) string {
	hd := hashids.NewData()
	hd.Salt = "#trace-id-hashids-salt@realsee#"
	hd.MinLength = 32

	ipIntValue := StringIpToInt(hostIP)

	unixNano := time.Now().UnixNano()

	ts := unixNano / 1e9
	nonce := unixNano % 1e9 / 1e3

	h, _ := hashids.NewWithData(hd)
	fmt.Println(unixNano, ts, nonce)
	result, _ := h.Encode([]int{ipIntValue, int(ts), int(nonce)})
	return result

}

func StringIpToInt(ipstring string) int {
	ipSegs := strings.Split(ipstring, ".")
	var ipInt int = 0
	var pos uint = 24
	for _, ipSeg := range ipSegs {
		tempInt, _ := strconv.Atoi(ipSeg)
		tempInt = tempInt << pos
		ipInt = ipInt | tempInt
		pos -= 8
	}
	return ipInt
}

// int转ip串
func IpIntToString(ipInt int) string {
	ipSegs := make([]string, 4)
	var len int = len(ipSegs)
	buffer := bytes.NewBufferString("")
	for i := 0; i < len; i++ {
		tempInt := ipInt & 0xFF
		ipSegs[len-i-1] = strconv.Itoa(tempInt)
		ipInt = ipInt >> 8
	}
	for i := 0; i < len; i++ {
		buffer.WriteString(ipSegs[i])
		if i < len-1 {
			buffer.WriteString(".")
		}
	}
	return buffer.String()
}
