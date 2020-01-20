package classfile

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
1、attribute_name_index 项的值必须是一个对常量池的有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示字符串“SourceFile”。
2、attribute_length的值必须为2；
3、sourcefile_index 项的值必须是一个对常量池的有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示一个字符串。
*/

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
