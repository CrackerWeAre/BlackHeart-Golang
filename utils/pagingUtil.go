package utils

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
)

// PgSplit func
func PgSplit(inputData []model.User, page, maxResult int) interface{} {
	if page == 0 {
		return JSONReturnMsg(false, "페이지는 1부터 시작됩니다")
	}

	startNum, lastNum := maxResult*(page-1), maxResult*page

	if startNum >= len(inputData) {
		return JSONReturnMsg(false, "해당 페이지에는 결과가 없습니다")
	}

	if lastNum > len(inputData) {
		return JSONReturnRes(true, inputData[startNum:])
	}
	return JSONReturnRes(true, inputData[startNum:])
}
