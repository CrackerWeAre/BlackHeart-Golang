package utils

// CheckErr func
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
