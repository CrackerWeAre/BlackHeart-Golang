package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetCommentList func
func GetCommentList(uID string) []model.Comment {
	query := "SELECT rcID, rID, uID, rcContent, rcCreateDate FROM review_comment WHERE uID='" + uID + "'"

	cmList := crud.GetCommentList(query)

	return cmList
}
