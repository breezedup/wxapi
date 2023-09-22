package wxapi

import (
	"encoding/json"
	"fmt"
)

type MenuErr struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type MenuRet struct {
	Data1 MenuCustom
	Data2 MenuWebPub
}

// 自定义菜单配置(通过api接口创建的)

type MenuCustom struct {
	IsMenuOpen   int `json:"is_menu_open"`
	SelfmenuInfo struct {
		Button []struct {
			Type      string `json:"type,omitempty"`
			Name      string `json:"name"`
			Key       string `json:"key,omitempty"`
			SubButton struct {
				List []struct {
					Type string `json:"type"`
					Name string `json:"name"`
					Url  string `json:"url,omitempty"`
					Key  string `json:"key,omitempty"`
				} `json:"list"`
			} `json:"sub_button,omitempty"`
		} `json:"button"`
	} `json:"selfmenu_info"`
}

// 网站功能发布菜单

type MenuWebPub struct {
	IsMenuOpen   int `json:"is_menu_open"`
	SelfmenuInfo struct {
		Button []struct {
			Name      string `json:"name"`
			SubButton struct {
				List []struct {
					Type     string `json:"type"`
					Name     string `json:"name"`
					Url      string `json:"url,omitempty"`
					Value    string `json:"value,omitempty"`
					NewsInfo struct {
						List []struct {
							Title      string `json:"title"`
							Author     string `json:"author"`
							Digest     string `json:"digest"`
							ShowCover  int    `json:"show_cover"`
							CoverUrl   string `json:"cover_url"`
							ContentUrl string `json:"content_url"`
							SourceUrl  string `json:"source_url"`
						} `json:"list"`
					} `json:"news_info,omitempty"`
				} `json:"list"`
			} `json:"sub_button,omitempty"`
			Type  string `json:"type,omitempty"`
			Value string `json:"value,omitempty"`
		} `json:"button"`
	} `json:"selfmenu_info"`
}

func (h *MenuRet) handleData(data string) error {
	err := json.Unmarshal([]byte(data), h)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	// 处理数据的逻辑
	fmt.Println("Data1 Field1:", h.Data1)
	fmt.Println("Data2 Field2:", h.Data2)
	return nil
}
