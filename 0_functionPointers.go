package main

import (
	"fmt"
	"strconv"

	"github.com/sachinmahanin/commonpassword/config"
	webserver "github.com/zhongjie-cai/web-server"
)

// func pointers for injection / testing: main.go
var (
	fmtErrorf               = fmt.Errorf
	configSetupApplication  = config.SetupApplication
	strconvAtoi             = strconv.Atoi
	webserverNewApplication = webserver.NewApplication
	DownloadFileFunc        = DownloadFile
)
