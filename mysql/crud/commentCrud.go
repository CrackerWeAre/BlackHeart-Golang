package crud

import (
	"github.com/ssoyyoung.p/BlackHeart-Golang/model"
	"github.com/ssoyyoung.p/BlackHeart-Golang/utils"
)

// GetCommentList func
func GetCommentList(query string) []model.Comment {

	DB := ConnectDB()
	defer DB.Close()

	var cm model.Comment
	var allCms []model.Comment

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
