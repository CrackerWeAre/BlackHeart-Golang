package model

// DBinfo strutct
type DBinfo struct {
	Username string
	Password string
	Hostname string
	Port     string
	Database string
}

// User struct > user_list
type User struct {
	UID       int    `json:"uID"`
	Uemail    string `json:"uEmail"`
	Upw       string `json:"uPW"`
	Uname     string `json:"uName"`
	Ugender   string `json:"uGender"`
	Uaddr     string `json:"uAddr"`
	Upost     int    `json:"uPost"`
	Uphone    string `json:"uPhone"`
	UjoinDate string `json:"uJoinDate"`
	Ubirth    string `json:"uBirth"`
	Uagree    int    `json:"uAgree"`
	Ulevel    int    `json:"uLevel"`
	UjoinPath string `json:"uJoinPath"`
}

// Result struct
type Result struct {
	Ulist    []User
	LastPage int
}
