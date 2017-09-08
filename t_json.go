package main

import (
	"fmt"
	// "github.com/allbuleyu/spider/spider"
	"encoding/json"
)

type Response struct {
	Code    int                    `json:"code"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

var str = []byte(`
{
    "code": 200,
    "message": "xxx",
    "data": {
        "a": 1,
        "b": "c"
    }
}`)

func main() {
	resp := Response{}
	err := json.Unmarshal(str, &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)

	fmt.Println(resp.Code, resp.Message, resp.Data)
}
