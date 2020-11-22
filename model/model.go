package model

// DBinfo strutct
type DBinfo struct {
	Username string
	Password string
	Hostname string
	Port     string
	Database string
}

// UserList struct > user_list
type UserList struct {
	UID       int
	Uemail    string
	Upw       string
	Uname     string
	Ugender   string
	Uaddr     string
	Uphone    int
	UjoinDate string
	Ubirth    string
	Uagree    int
	Ulevel    int
	UjoinPath string
}
