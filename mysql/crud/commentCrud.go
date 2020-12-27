package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetCommentList func
func GetCommentList(query string) []model.ReviewComment {

	DB := ConnectDB()
	defer DB.Close()

	var cm model.ReviewComment
	var allCms []model.ReviewComment

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&cm.RCID, &cm.RID, &cm.UID, &cm.RcContent, &cm.RcCreateDate)
		utils.CheckErr(err)

		allCms = append(allCms, cm)
	}

	return allCms
}

// GetBoardList func
func GetBoardList(query string) []model.ReviewBoard {

	DB := ConnectDB()
	defer DB.Close()

	var br model.ReviewBoard
	var allBr []model.ReviewBoard

	rows, err := DB.Query(query)
	utils.CheckErr(err)

	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&br.RID, &br.PID, &br.UID, &br.Rrate, &br.Rtitle, &br.Rcontent, &br.Rfile, &br.RcreateDate, &br.Rhit, &br.Rlike)
		utils.CheckErr(err)

		allBr = append(allBr, br)
	}

	return allBr
}
