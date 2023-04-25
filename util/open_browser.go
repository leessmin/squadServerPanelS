package util

import (
	"fmt"
	"os/exec"
	"runtime"
)

// 打开浏览器
func OpenBrowser(uri string) {

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	args = append(args, uri)

	e := exec.Command(cmd, args...)

	if err := e.Start(); err != nil {
		fmt.Println("打开浏览器失败，err:", err)
	}
}
