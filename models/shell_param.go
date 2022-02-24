package models

type ReqParam struct {
	ShellPath string `json:"shellPath" binding:"required"`
}
