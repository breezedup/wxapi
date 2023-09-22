package common

func Cache() cacheInter {
	return &cacheType{}
}

type cacheType struct {
	accessToken string
	expiresIn   int
}

func (this *cacheType) SetAccessToken(accessToken string, expiresIn int) {
	this.accessToken = accessToken
	this.expiresIn = expiresIn
}
func (this *cacheType) GetAccessToken() string {
	return this.accessToken
}
