package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

//创建结构体实例
func newDirEntry(path string) *DirEntry{
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

//把目录和class文件名拼成一个完整的路径，然后调用ioutil包提供的ReadFile()函数读取class文件内容，最后返回
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string{
	return self.absDir
}