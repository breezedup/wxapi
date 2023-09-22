package common

import (
	"encoding/json"
	"fmt"
	"os"
)

var MenuConf = &menuType{}

// 菜单配置

type menuType struct {
	MenuCreatePost                string `json:"menu_create"`                    //创建接口
	MenuGetCurrentSelfMenuInfoGet string `json:"menu_get_current_selfmenu_info"` //查询接口
	MenuDeleteGet                 string `json:"menu_delete"`                    //删除接口
	MenuAddConditionalPost        string `json:"menu_addconditional"`            //创建个性化菜单
	MenuDelConditionalPost        string `json:"menu_delconditional"`            //删除个性化菜单
	MenuTryMatchPost              string `json:"menu_trymatch"`                  //测试个性化菜单匹配结果
}

func init() {
	data, err := os.ReadFile("./conf/menu_conf.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, MenuConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(MenuConf)
}
