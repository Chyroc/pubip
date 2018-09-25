package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var ipReg = regexp.MustCompile(`您的本地上网IP是：<h2>(.*?)</h2>`)

func main() {
	run()
}

func run() {
	resp, err := http.Get("http://www.net.cn/static/customercare/yourip.asp")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	bs, err = gbkToUtf8(bs)
	if err != nil {
		panic(err)
	}

	matched := ipReg.FindSubmatch(bs)
	if len(matched) < 2 {
		panic("run error")
	}

	fmt.Println(string(matched[1]))
}

func gbkToUtf8(s []byte) ([]byte, error) {
	return simplifiedchinese.GBK.NewDecoder().Bytes(s)
}
