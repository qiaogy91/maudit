package event

import "github.com/qiaogy91/ioc/utils"

func ErrEventValidate(s error) *utils.ApiException { return utils.NewApiException(10000, s.Error()) }
func ErrEventCreate(s error) *utils.ApiException   { return utils.NewApiException(10001, s.Error()) }
func ErrQueryEvent(s error) *utils.ApiException    { return utils.NewApiException(10002, s.Error()) }
func ErrApiQueryReq(s error) *utils.ApiException   { return utils.NewApiException(10003, s.Error()) }
