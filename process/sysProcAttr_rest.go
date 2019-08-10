// +build !windows

package process

import "syscall"

var sysProcAttr = &syscall.SysProcAttr{}
