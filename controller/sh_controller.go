package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"http-shell/models"
	"http-shell/service"
	"log"
	"strings"
)

// ExecuteShell 执行shell脚本
func ExecuteShell(c *gin.Context) {
	var reqParam models.ReqParam
	err := c.ShouldBindJSON(&reqParam)
	if err != nil {
		_, _ = fmt.Fprint(c.Writer, err)
		log.Printf("解析executeShell入参出错: %v\n", err)
		return
	}
	fmt.Printf("reqParam:%v\n", reqParam)

	shellPath := reqParam.ShellPath
	if strings.Contains(shellPath, " ") {
		_, _ = fmt.Fprint(c.Writer, "脚本路径不能包含空格")
		return
	}

	if !strings.HasSuffix(shellPath, ".sh") {
		_, _ = fmt.Fprint(c.Writer, "脚本必须以.sh结尾")
		return
	}

	var params = []string{shellPath}
	outputString, err := service.ExecCommand("sh", params)
	if err == nil {
		_, _ = fmt.Fprint(c.Writer, outputString)
	} else {
		_, _ = fmt.Fprint(c.Writer, err)
	}
}
