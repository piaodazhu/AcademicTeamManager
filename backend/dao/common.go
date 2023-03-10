package dao

import "context"

const (
	USER    = "user"
	STUDENT = "student"
	PROJECT = "project"
	OUTPUT  = "output"
	MAILCONFIG  = "mail_config"
)

const (
	NumberNull = 0
	StringNull = ""
)

const (
	Dev  = "dev"
	Prod = "prod"
)

var ctx = context.Background()
