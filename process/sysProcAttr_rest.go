// +build !windows

package process

import "syscall"

var sysProcAttrHide = &syscall.SysProcAttr{}

func printOs() { fmt.Println("OS linux")}