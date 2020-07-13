package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func httpGet() {

	resp, err := http.Get("http://www.baidu.com")
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(resp.Body)

	fmt.Println(string(body))

	resp.Body.Close()
        foo := 0
        var str string = "test"
        foo += 1
}

func main() {
	httpGet()
}
