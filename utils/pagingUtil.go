package utils

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
)

type result struct {
	result   []model.User
	lastPage int
}

// PgSplit func
func PgSplit(inputData []model.User, page, maxResult int) interface{} {
	if page == 0 {
		return JSONReturnMsg(false, "페이지는 1부터 시작됩니다")
	}

	startNum, lastNum := maxResult*(page-1), maxResult*page
	lastPage := (len(inputData) / maxResult) + 1


	if startNum >= len(inputData) {
		return JSONReturnMsg(false, "해당 페이지에는 결과가 없습니다")
	}

	if lastNum > len(inputData) {
		return JSONReturnRes(true, inputData[startNum:], lastPage)
	}

	return JSONReturnRes(true, inputData[startNum:lastNum], lastPage)
}
