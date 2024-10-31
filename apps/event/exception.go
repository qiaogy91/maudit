package event

import "github.com/qiaogy91/ioc/utils"

func ErrEventValidate(e error) *utils.ApiException {
	return utils.NewApiException(10011, "Event 参数校验错误", e)
}
func ErrEventCreate(e error) *utils.ApiException {
	return utils.NewApiException(10011, "Event 创建错误", e)
}
func ErrQueryEvent(e error) *utils.ApiException {
	return utils.NewApiException(10011, "Event 查询错误", e)
}
