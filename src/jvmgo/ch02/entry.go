package classpath

import (
	"os"
	"strings"
)

//路径分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	//负责寻找和加载class文件
	readClass(className string) ([]byte, Entry, error)
	//返回变量的字符串显示
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip")|| strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
