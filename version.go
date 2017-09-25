package main

import (
	"fmt"
)

var cmdVersion = &Command{
	Run:       runVersion,
	UsageLine: "version",
	Short:     "输出Tars版本",
	Long:      `输出当前Tars版本。`,
}

func runVersion(cmd *Command, args []string) {
	if len(args) != 0 {
		cmd.Usage()
	}

	fmt.Printf("tars version %s\n", version)
}
