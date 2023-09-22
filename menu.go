package wxapi

import (
	"fmt"
	"github.com/breezedup/wxapi/common"
)

func NewMenu() menuInter {
	return &menuType{}
}

type menuType struct {
}

func (this *menuType) Create(body map[string]any) *MenuErr {
	data := common.NewNetWork().SendPost(common.MenuConf.MenuCreatePost, body)
	return &MenuErr{data.Code, data.Msg}
}
func (this *menuType) GetCurrentSelfMenuInfo(body map[string]string) (mc *MenuRet) {
	data := common.NewNetWork().SendGet(common.MenuConf.MenuGetCurrentSelfMenuInfoGet, body)
	fmt.Println(data)
	mc.handleData(data.Msg)
	return
}
func (this *menuType) Delete(body map[string]string) {
	data := common.NewNetWork().SendGet(common.MenuConf.MenuDeleteGet, body)
	fmt.Println(data)
}
func (this *menuType) AddConditional(body map[string]any) {
	data := common.NewNetWork().SendPost(common.MenuConf.MenuAddConditionalPost, body)
	fmt.Println(data)
}
func (this *menuType) DelConditional(body map[string]any) {
	data := common.NewNetWork().SendPost(common.MenuConf.MenuDelConditionalPost, body)
	fmt.Println(data)
}
func (this *menuType) TryMatch(body map[string]any) {
	data := common.NewNetWork().SendPost(common.MenuConf.MenuTryMatchPost, body)
	fmt.Println(data)
}
