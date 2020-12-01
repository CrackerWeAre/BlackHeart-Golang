package utils

import (
	"github.com/gin-gonic/gin"
)

// JSONReturnMsg func
func JSONReturnMsg(status bool, message string) gin.H {

	return gin.H{
		"status":  status,
		"message": message,
	}
}

// JSONReturnRes func
func JSONReturnRes(status bool, result interface{}, lastPage int) gin.H {

	return gin.H{
		"status":   status,
		"result":   result,
		"lastPage": lastPage,
	}
}
