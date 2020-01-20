package classfile

/*
ConstantValue_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 constantvalue_index;
}
1、attribute_name_index 项的值，必须是一个对常量池的有效索引。常量池在该索引处的项必须是 CONSTANT_Utf8_info（§4.4.7）结构，表示字符串“ConstantValue”。
2、ConstantValue_attribute 结构的 attribute_length 项的值固定为 2。
3、constantvalue_index 项的值，必须是一个对常量池的有效索引。常量池在该索引处的项给出该属性表示的常量值。常量池的项的类型表示的字段类型如表 4.7 所示。
字段类型                        项类型
long                            CONSTANT_Long
float                           CONSTANT_Float
double                          CONSTANT_Double
int.short,char,byte,boolean     CONSTANT_Integer
String                          CONSTANT_String
*/

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
