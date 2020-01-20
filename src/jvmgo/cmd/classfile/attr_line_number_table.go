package classfile

/*
它被调试器用于确定源文件中行号表示的内容在 Java 虚拟机的 code[]数组中对应的部分。LineNumberTable可以按照任意顺序出现，该属性不需要与
源文件的行一一对应，可以按照任意顺序出现。
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
1、attribute_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示字符串“LineNumberTable”
2、attribute_length 给出了当前属性的长度，不包括开始的 6 个字节。
3、line_number_table_length 项的值给出了 line_number_table[]数组的成员个数。
4、line_number_table[]数组的每个成员都表明源文件中行号的变化在 code[]数组中都会有对应的标记点
    4.1、start_pc 项的值必须是 code[]数组的一个索引，code[]数组在该索引处的字符表示源文件中新的行的起点。start_pc 项的值必须小于当前 LineNumberTable
         属性所在的 Code 属性的 code_length 项的值。
    4.2、line_number 项的值必须与源文件的行数相匹配。
*/
type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
