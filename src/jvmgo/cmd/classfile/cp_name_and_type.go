package classfile
/*
CONSTANT_NameAndType_info {
	u1 tag;
	u2 name_index;
	u2 descriptor_index;
}
1、name_index的值必须是对常量池的有效索引，常量池在该索引处的项必须是CONSTANT_Utf8_info结构，要么表示特殊的方法名<init>，
要么表示一个有效的字段或方法的非限定名
2、descriptor_index的值必须是对常量池的有效索引，常量池在该索引处的项必须是CONSTANT_Utf8_info结构，这个结构表示一个有效
的字段描述符或方法描述符

字段描述符
字符            类型	        含义
B               byte        有符号字节整数
C               char        Unicode字符，UTF-16编码
D               double      双精度浮点数
F               float       单精度浮点数
I               int         整型数
J               long        长整型
S               short       有符号短整数
Z               boolean     布尔值 true/false
L CLassname;    reference   一个名为<Classname>的实例
[               reference   一个一维数组
*/

type ConstantNameAndTypeInfo struct {
	nameIndex		uint16
	descriptorIndex	uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader){
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}