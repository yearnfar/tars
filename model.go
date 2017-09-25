package main

import (
//"fmt"
)

var cmdModel = &Command{
	Run:       genModel,
	UsageLine: "model",
	Short:     "生成model文件",
	Long:      `在指定目录生成model文件。`,
}

func genModel(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
	}

}
