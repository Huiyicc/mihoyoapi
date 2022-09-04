package request

import (
	"fmt"
	"github.com/Huiyicc/mihoyoapi/define"
	"github.com/Huiyicc/mihoyoapi/tools"
	"math"
	"math/rand"
	"time"
)

// GetDsSign 为签到ds
func GetDsSign() string {
	n := "N50pqm7FSy2AkFz2B3TqtuZMJ5TOl3Ep"
	t := time.Now().Unix()
	r := tools.RandomStr(6)
	DS := tools.GetMd5(fmt.Sprintf("salt=%s&t=%d&r=%s", n, t, r))
	return fmt.Sprintf("%d,%s,%s", t, r, DS)
}

// GetDsSign2 为签到ds2
func GetDsSign2() string {
	n := "z8DRIUjNDT7IT5IZXvrUAxyupA1peND9"
	t := time.Now().Unix()
	randomStr := tools.RandomStr(6)
	singsf := fmt.Sprintf("salt=%s&t=%d&r=%s", n, t, randomStr)
	sing := tools.GetMd5(singsf)
	DS := fmt.Sprintf("%d,%s,%s", t, randomStr, sing)
	return DS
}

// GetDs 获取请求ds
func GetDs(server, query, body string) string {
	n := ""
	if server == define.SERVER_SHIJIESHU || server == define.SERVER_TIANKONGDAO {
		n = "xV8v4Qu54lUKrEYFZkJhB8cuOh9Asafs"
	}
	rand.Seed(time.Now().UnixNano())
	t := time.Now().Unix()
	r := int64(math.Floor(float64(rand.Intn(1))*900000 + 100000))
	DS := tools.GetMd5(fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", n, t, r, body, query))
	return fmt.Sprintf("%d,%d,%s", t, r, DS)
}

func GetDs2(query, body string) string {
	n := "t0qEgfub6cvueAPgR5m9aQWWVciEer7v"
	t := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(100001) + 99999
	s := tools.GetMd5(fmt.Sprintf("salt=%s&t=%d&r=%d&b=%s&q=%s", n, t, r, body, query))
	return fmt.Sprintf("%d,%d,%s", t, r, s)
}
