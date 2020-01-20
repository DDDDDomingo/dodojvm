package classfile

/*
它被调试器用于确定方法在执行过程中局部变量的信息。在 Code 属性的属性表中，LocalVariableTable 属性可以按照任意顺序出现。Code 属性中
的每个局部变量最多只能有一个 LocalVariableTable 属性。
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
1、attribute_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示字符串“LocalVariableTable”。
2、attribute_length 项的值给出当前属性的长度，不包括开始的 6 个字节。
3、local_variable_table_length 项的值给出了 local_variable_table[]数组的成员的数量。
4、local_variable_table[]数组的每一个成员表示一个局部变量的值在 code[]数组中的偏移量范围。它同时也是用于从当前帧的局部变量表找出所需的局部变量的索引。
    4.1、start_pc, length：所有给定的局部变量的索引都在范围[start_pc, start_pc+length)，
    4.2、name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示一个局部变量的有效的非全限定名。
    4.3、descriptor_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示源程序中局部变量类型的字段描述符。
    4.4、index 为此局部变量在当前栈帧的局部变量表中的索引。long 或 double 型，则占用 index 和 index+1 两个索引。
*/
type LocalVariableTable struct {
	localVariableTable []*LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc    uint16
	length     uint16
	nameIndex  uint16
	descriptor uint16
	index      uint16
}

func (self *LocalVariableTable) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for i := range self.localVariableTable {
		self.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:    reader.readUint16(),
			length:     reader.readUint16(),
			nameIndex:  reader.readUint16(),
			descriptor: reader.readUint16(),
			index:      reader.readUint16(),
		}
	}
}
