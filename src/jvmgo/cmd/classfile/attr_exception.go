package classfile

/*
Exceptions_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 number_of_exceptions;
    u2 exception_index_table[number_of_exceptions];
}
1、attribute_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info结构，表示字符串"Exceptions"。
2、attribute_length 项的值给出了当前属性的长度，不包括开始的 6 个字节。
3、number_of_exceptions 项的值给出了 exception_index_table[]数组中成员的数量。
4、exception_index_table[]数组的每个成员的值都必须是对常量池的有效索引。常量池在这些索引处的成员必须都是 CONSTANT_Class_info结构，表示这个方法声明要抛出的异常的类的类型。
*/

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ExceptionAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
