package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"http-shell/models"
	"http-shell/service"
	"log"
	"strings"
)

// ExecuteShell 直接执行shell脚本
func ExecuteShell(context *gin.Context) {
	shellPath := checkAndGetShellPath(context)
	if len(shellPath) == 0 {
		return
	}

	outputString, err := service.ExecCommand(shellPath, []string{})
	if err == nil {
		_, _ = fmt.Fprint(context.Writer, outputString)
	} else {
		_, _ = fmt.Fprint(context.Writer, err)
	}
}

// ExecuteShShell 使用sh执行shell脚本
func ExecuteShShell(context *gin.Context) {
	shellPath := checkAndGetShellPath(context)
	if len(shellPath) == 0 {
		return
	}

	var params = []string{shellPath}
	outputString, err := service.ExecCommand("sh", params)
	if err == nil {
		_, _ = fmt.Fprint(context.Writer, outputString)
	} else {
		_, _ = fmt.Fprint(context.Writer, err)
	}
}

func checkAndGetShellPath(context *gin.Context) string {
	var reqParam models.ReqParam
	err := context.ShouldBindJSON(&reqParam)
	if err != nil {
		_, _ = fmt.Fprint(context.Writer, err)
		log.Printf("解析executeShell入参出错: %v\n", err)
		return ""
	}
	fmt.Printf("reqParam:%v\n", reqParam)

	shellPath := reqParam.ShellPath
	if strings.Contains(shellPath, " ") {
		_, _ = fmt.Fprint(context.Writer, "脚本路径不能包含空格")
		return ""
	}

	if !strings.HasSuffix(shellPath, ".sh") {
		_, _ = fmt.Fprint(context.Writer, "脚本必须以.sh结尾")
		return ""
	}

	return shellPath
}
