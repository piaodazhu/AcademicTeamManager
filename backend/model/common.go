package model

type Page struct {
	PageNum  int `form:"pageNum"  json:"pageNum"`
	PageSize int `form:"pageSize" json:"pageSize"`
}

type MailParam struct {
	Smtp       string
	Port       int
	AuthCode   string
	Sender     string
	Subject    string
	Content    string
	Attachment string
	Receiver   string
}

type FileParam struct {
	Name string `json:"name" binding:"required,gt=0"`
}

type FileInfo struct {
	Url  string `json:"url"`
	Name string `json:"name"`
}
