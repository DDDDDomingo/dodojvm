package classfile

/*
Deprecated_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
1、attribute_name_index 项的值必须是对常量池的一个有效索引。常量池在该索引处的成员必须是 CONSTANT_Utf8_info（§4.4.7）结构，表示字符串“Deprecated”。
2、attribute_length的值固定为0。
*/
type DeprecatedAttribute struct {
	MarkerAttribute
}

/*
Synthetic_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
}
1、attribute_name_index 项的值必须是对常量池的一个有效索引，常量池在该索引处的成员必须是 CONSTANT_Utf8_info（§4.4.7）结构，表示字符串“Synthetic”。
2、attribute_length的值固定为0。
*/
type SyntheticAttribute struct {
	MarkerAttribute
}

type MarkerAttribute struct {
}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	//attribute_length为0，所以read nothing
}
