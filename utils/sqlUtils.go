package utils

import "strconv"

func MakeCateFilter(cateList []string) string {
	base := " where "

	for idx, cate := range cateList {

		if cate != "" {
			if idx > 0 {
				base += " and "
			}
			base += "pCategory"+strconv.Itoa(idx+1)+"='"+cate+"'"
		}
	}

	return base

}
