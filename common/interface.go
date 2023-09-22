package common

type netInterface interface {
	SendPost(url string, body map[string]any) *netError
	SendGet(baseURL string, data map[string]string) *netError
}
type cacheInter interface {
	SetAccessToken(accessToken string, expiresIn int)
	GetAccessToken() string
}
