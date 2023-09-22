package wxapi

type menuInter interface {
	Create(body map[string]any) *MenuErr
	GetCurrentSelfMenuInfo(body map[string]string) (mc *MenuRet)
	Delete(body map[string]string)
	AddConditional(body map[string]any)
	DelConditional(body map[string]any)
	TryMatch(body map[string]any)
}
type connectInter interface {
	newConnect(appId, appSecret string) *Result
}
type MenuRetInter interface {
	HandleData(data *MenuRet) error
}
