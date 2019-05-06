package cmd

import (
	"fmt"
	"os"
)

//确认退出，防止双击exe执行时，程序自动退出cmd
func EnterExit(status int) {
	fmt.Println("按 Enter键 退出")
	fmt.Scanln()
	os.Exit(status)
}
