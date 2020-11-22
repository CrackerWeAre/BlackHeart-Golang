package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
)

// PgSplit func
func PgSplit(inputData []model.UserList, page, maxResult int) interface{} {
	if page == 0 {
		return gin.H{"msg": "페이지는 1부터 시작됩니다."}
	}

	startNum, lastNum := maxResult*(page-1), maxResult*page

	if startNum >= len(inputData) {
		return gin.H{"msg": "해당 페이지에는 결과가 없습니다."}
	}

	if lastNum > len(inputData) {
		return inputData[startNum:]
	}
	return inputData[startNum:lastNum]
}
