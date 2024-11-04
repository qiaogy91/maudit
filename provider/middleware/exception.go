package middleware

import "github.com/qiaogy91/ioc/utils"

func ErrTokenAuth(e error) *utils.ApiException {
	return utils.NewApiException(401, "Token 校验失败", e)
}

func ErrRouteSelect(e error) *utils.ApiException {
	return utils.NewApiException(501, "Route 匹配失败", e)
}
