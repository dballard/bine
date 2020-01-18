// +build !windows

package process

import "syscall"

var sysProcAttrHide = &syscall.SysProcAttr{}
