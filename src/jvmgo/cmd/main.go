package main

import (
	"fmt"
	"jvmgo/cmd/classpath"
	"strings"
)

func main(){
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class =="" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	//输出命令行参数
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("Classpath:%s class:%s args:%v\n",
		cp, cmd.class, cmd.args)
	//读取主类数据
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("Could not find or load main class *s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}