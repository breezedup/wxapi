package wxapi

import (
	"encoding/json"
	"wxapi/common"
)

func NewConnect(appId, appSecret string) *Result {
	var ci = connectInter(&connect{})
	return ci.newConnect(appId, appSecret)
}

type connect struct {
}

type Result struct {
	Code        int
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Error       string
}

func (this *connect) newConnect(appId, appSecret string) (ret *Result) {
	netinfo := common.NewNetWork().SendGet(common.TokenConf.TokenGet, map[string]string{
		"grant_type": "client_credential",
		"appid":      appId,
		"secret":     appSecret})
	if netinfo.Code == 0 {
		err := json.Unmarshal([]byte(netinfo.Msg), ret)
		if err != nil {
			ret.Code = -1
			ret.Error = err.Error()
		}
		common.Cache().SetAccessToken(ret.AccessToken, ret.ExpiresIn)
		return
	}
	return nil
}
