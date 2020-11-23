package utils

import "fmt"

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
