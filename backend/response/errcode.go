package response

const (
	ErrCodeSuccess      = 0
	ErrCodeFailed       = 1
	ErrCodeParamInvalid = 2
	ErrCodeNoLogin      = 3
	ErrCodeTokenExpire  = 4

	ErrCodeInitDataFailed   = 10
	ErrCodeFileUploadFailed = 11
	ErrCodeFileRemoveFailed = 12

	ErrCodeUserHasExist         = 10001
	ErrCodeUserNotExist         = 10002
	ErrCodeUserEmailOrPass      = 10003
	ErrCodeVerityCodeSendFailed = 10004
	ErrCodeVerityCodeInvalid    = 10005
	ErrCodeEmailFormatInvalid   = 10008
	ErrCodeUserPassResetFailed  = 10009

	ErrCodeStudentHasExist = 20001

	ErrCodeFileExportFailed = 70001

	ErrCodeOutputHasExist = 80001

	ErrCodeMailConfigInvalid = 90001
	ErrCodeMailSendFailed    = 90002
	ErrCodeMailSendUnEnable  = 90003
	ErrCodeMailConfigUnSet   = 90004
)

var msg = map[int]string{
	ErrCodeSuccess:              "success",
	ErrCodeFailed:               "failed",
	ErrCodeParamInvalid:         "param invalid",
	ErrCodeNoLogin:              "no login",
	ErrCodeTokenExpire:          "token expire",
	ErrCodeInitDataFailed:       "data init failed",
	ErrCodeFileUploadFailed:     "file upload failed",
	ErrCodeUserHasExist:         "user has exist",
	ErrCodeUserNotExist:         "user not exist",
	ErrCodeUserEmailOrPass:      "user email or password error",
	ErrCodeVerityCodeSendFailed: "verify code send failed",
	ErrCodeVerityCodeInvalid:    "verify code invalid",
	ErrCodeEmailFormatInvalid:   "email format invalid",
	ErrCodeUserPassResetFailed:  "user password reset failed",
	ErrCodeStudentHasExist:      "student has exist",
	ErrCodeFileExportFailed:     "file export failed",
	ErrCodeOutputHasExist:       "output has exist",
	ErrCodeMailConfigInvalid:    "mail config invalid",
	ErrCodeMailSendFailed:       "mail send failed",
	ErrCodeMailSendUnEnable:     "mail send server unEnable",
	ErrCodeMailConfigUnSet:      "mail config un set",
}
