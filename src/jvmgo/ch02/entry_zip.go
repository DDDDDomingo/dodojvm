package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

//构造结构体实例
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	//打开zip文件
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	//defer保证文件可以关闭
	defer r.Close()
	//遍历zip中的文件，查找class文件
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			//保证文件可以关闭
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
		return nil, nil, errors.New("class not found: " + className)
	}
}

func (self *ZipEntry) String() string {
	return self.absPath
}
