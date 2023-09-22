package common

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func NewNetWork() netInterface {
	return &netWork{}
}

type netWork struct {
}
type netError struct {
	Code int
	Msg  string
}

func newNetError(code int, msg string) *netError {
	return &netError{code, msg}
}
func (this *netWork) SendPost(url string, body map[string]any) *netError {
	url = url + "?access_token=" + Cache().GetAccessToken()
	var bodyStr string
	if body != nil && len(body) > 0 {
		data, err := json.Marshal(body)
		if err != nil {
			fmt.Println("POST Marshal:", err)
		}
		bodyStr = string(data)
	}
	// 发起POST请求
	postResp, err := http.Post(url, "application/json",
		strings.NewReader(bodyStr))
	if err != nil {
		fmt.Println("POST请求出错:", err)
		return newNetError(-1, err.Error())
	}
	defer postResp.Body.Close()

	// 读取POST请求响应的内容
	postBody, err := io.ReadAll(postResp.Body)
	if err != nil {
		fmt.Println("读取POST响应出错:", err)
		return newNetError(-1, err.Error())
	}
	fmt.Println("POST响应内容:", string(postBody))
	return newNetError(0, string(postBody))
}
func (this *netWork) SendGet(baseURL string, data map[string]string) *netError {
	if baseURL != TokenConf.TokenGet {
		baseURL = baseURL + "?access_token=" + Cache().GetAccessToken()
	}
	params := url.Values{}
	for s, i := range data {
		params.Add(s, i)
	}
	urlObj, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("URL parsing error:", err)
		return newNetError(-1, err.Error())
	}
	urlObj.RawQuery = params.Encode()

	// 发起GET请求
	getResp, err := http.Get(urlObj.String())
	if err != nil {
		fmt.Println("GET请求出错:", err)
		return newNetError(-1, err.Error())
	}
	defer getResp.Body.Close()

	// 读取GET请求响应的内容
	getBody, err := io.ReadAll(getResp.Body)
	if err != nil {
		fmt.Println("读取GET响应出错:", err)
		return newNetError(-1, err.Error())
	}
	fmt.Println("GET响应内容:", string(getBody))
	return newNetError(0, string(getBody))
}
