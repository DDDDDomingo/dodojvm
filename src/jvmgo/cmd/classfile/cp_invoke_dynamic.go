package classfile

/*
CONSTANT_MethodHandle_info {
    u1 tag;
    u1 reference_kind;
    u2 reference_index;
}
1、tag，值为15
2、reference_kind，值范围为[1-9]
3、reference_index，表示对常量池的有效索引；索引处的项的类别由reference_kind决定，
    [1、2、3、4]——CONSTANT_Fieldref_info，表示由一个字段创建的方法句柄
    [5、6、7、8]——CONSTANT_Methodref_info，表示由类的方法或构造函数创建的方法句柄
    [9]        ——CONSTANT_InterfaceMethodref_info，表示由接口方法创建的方法句柄
    [5、6、7、9]——对应的方法不能为实例初始化(<init>)方法或类初始化方法(<clinit>)
    [8]        ——对应的方法必须为实例初始化(<init>)方法
*/
type ConstantMethodHandleInfo struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *ConstantMethodHandleInfo) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

/*
CONSTANT_MethodType_info {
    u1 tag;
    u2 descriptor_index;
}
1、tag，值为16
2、descriptor_index，是对常量池的有效索引且必须是CONSTANT_Utf8_info结构，表示方法的描述符
*/
type ConstantMethodTypeInfo struct {
}
