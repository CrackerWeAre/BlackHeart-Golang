package mysql

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/mysql/crud"
)

// GetReviewCommentList func
func GetReviewCommentList(uID string) []model.ReviewComment {
	query := "SELECT rcID, rID, uID, rcContent, rcCreateDate FROM review_comment WHERE uID='" + uID + "'"

	cmList := crud.GetCommentList(query)

	return cmList
}

// GetReviewBoardList func
func GetReviewBoardList(uID string) []model.ReviewBoard {
	query := "SELECT * FROM review_board WHERE uID='" + uID + "'"

	brList := crud.GetBoardList(query)

	return brList
}


