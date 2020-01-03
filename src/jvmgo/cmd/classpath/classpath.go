package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

//如果没有-classpath/-cp选项，则使用当前目录作为用户类路径。
//依次从启动类路径、拓展类路径和用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	//启动类路径
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	//拓展类路径
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

func (self *Classpath) parseBootAndExtClasspath(jreOption string){
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func (self *Classpath) parseUserClasspath(cpOption string){
	if cpOption =="" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

//优先使用用户输入的-Xjre选项作为jre目录.如果没有输入该选项，则在当前目录下寻找jre目录.
//如果找不到，尝试使用JAVA_HOME
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre"){
		return "./jre"
	}
	if jh:= os.Getenv("JAVA_HOME"); jh!=""{
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (self *Classpath) String() string{
	return self.userClasspath.String()
}