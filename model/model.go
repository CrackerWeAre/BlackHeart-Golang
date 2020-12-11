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

//LoginUser struct
type LoginUser struct {
	Token    string `json:"token"`
	TokenEXP int64  `json:"tokenExp"`
	Uemail   string `json:"uEmail"`
}

// Order sturct
type Order struct {
	OID          int     `json:"oID"`
	UID          string  `json:"uID"`
	OOrderNum    int     `json:"oOrderNum"`
	OOrderDate   string  `json:"oOrderDate"`
	PID          int     `json:"pID"`
	PName        string  `json:"pName"`
	POption      string  `json:"pOption"`
	PPrice       int     `json:"pPrice"`
	PDistount    float32 `json:"pDiscount"`
	OQuantity    int     `json:"oQuantity"`
	ODestination string  `json:"oDestination"`
	OPhoneNum    string  `json:"oPhoneNum"`
	OPayment     int     `json:"oPayment"`
	OInvoice     int     `json:"oInvoice"`
	ODelivery    int     `json:"oDelivery"`
	OStatus      string  `json:"oStatus"`
	PListImage   string  `json:"pListImage"`
}

// Comment struct
type Comment struct {
	RCID         int    `json:"rcID"`
	RID          int    `json:"rID"`
	UID          int    `json:"uID"`
	RCPW         string `json:"rcPW"`
	RcContent    string `json:"rcContent"`
	RcCreateDate string `json:"rcCreateDate"`
}
