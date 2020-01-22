// +build windows

package process

import (
	"fmt"
	"syscall"
)

var sysProcAttrHide = &syscall.SysProcAttr{HideWindow: true}

func printOs() { fmt.Println("OS windows")}