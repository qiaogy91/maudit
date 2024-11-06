package api

import (
	"github.com/qiaogy91/ioc/utils"
	"net/http"
)

func ErrEventDecode(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10001, "Event请求参数错误", e)
}

func ErrEventQuery(e error) *utils.ApiException {
	return utils.NewApiException(http.StatusBadRequest, 10002, "Event查询过程错误", e)
}
