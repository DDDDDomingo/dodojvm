package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath	Entry
	extClasspath	Entry
	userClasspath	Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {

}

func (self *Classpath) String() string{

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

//优先使用用户输入的-Xjre选项作为jre目录.如果没有输入该选项，则在当前目录下寻找jre目录.
//如果找不到，尝试使用JAVA_HOME
func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {

	}
}

func exists(path string) bool {
	if _, err := os.Stat(path); err!=nil {
		if os.IsNotExist(err) {
			return false
		}
	}
}