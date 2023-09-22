package common

import (
	"encoding/json"
	"fmt"
	"os"
)

var TokenConf = &tokenType{}

// 菜单配置

type tokenType struct {
	TokenGet string `json:"token"` //创建接口
}

func init() {
	data, err := os.ReadFile("./conf/token_conf.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, TokenConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(TokenConf)
}
