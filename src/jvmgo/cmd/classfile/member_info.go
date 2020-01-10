package classfile

/*
字段
field_info {
	u2 access_flags;
	u2 name_index;
	u2 descriptor_index;
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}
1、access_flags, 用于定义字段被访问权限和基础属性的掩码标志
标记名			值			说明
ACC_PUBLIC		0x0001		public
ACC_PRIVATE		0x0002		private
ACC_PROTECTED	0x0004		protected
ACC_STATIC		0x0008		static
ACC_FINAL		0x0010		final
ACC_VOLATILE	0x0040		volatile
ACC_TRANSIENT	0x0080		transient
ACC_SYNTHETIC	0x1000		表示字段由编译器自动产生
ACC_ENUM		0x4000		表示字段为枚举类型
2、name_index项的值是对常量池的一个有效索引。索引项为CONSTANT_Utf8_info，表示一个有效的字段的非全限定名
3、descriptor_index项的值是对常量池的一个有效索引。索引项为CONSTANT_Utf8_info，表示一个有效的字段的描述符
4、attribute_count的项的值表示当前字段的附加属性的数量
5、attributes[]表的每一个成员的值必须是attribute，如：ConstantValue, Synthetic, Signature, Deprecated, RuntimeVisibleAnnotations,
RuntimeInvisibleAnnotations
*/
/*
方法
method_info {
	u2 access_flags;
	u2 name_index;
	u2 descriptor_index;
	u2 attributes_count;
	attribute_info attributes[attributes_count];
}
1、access_flags，用于定义当前方法的访问权限和基本属性的掩码标志
标记名				值			说明
ACC_PUBLIC			0x0001		public
ACC_PRIVATE			0x0002		private
ACC_PROTECTED		0x0004		protected
ACC_STATIC			0x0008		static
ACC_FINAL			0x0010		final，无法被Override
ACC_SYNCHRONIZED	0x0020		synchronized，方法由管程同步
ACC_BRIDGE			0x0040		bridge，方法由编译器产生
ACC_VARARGS			0x0080		表示方法带有变长参数
ACC_NATIVE			0x0100		native，方法引用非java语言的本地方法
ACC_ABSTRACT		0x0400		abstract，方法没有具体实现
ACC_STRICT			0x0800		strictfp，方法使用FP-strict浮点格式
ACC_SYNTHETIC		0x1000		方法在源文件中不出现，由编译器产生
2、name_index项的值是对常量池的一个有效索引。索引项为CONSTANT_Utf8_info，表示初始化方法的名字(<init>or<clinit>)一个方法的有效的非全限定名
3、descriptor_index项的值是对常量池的一个有效索引，索引项为CONSTANT_Utf8_info，表示一个有效的方法的描述符
4、attributes_count项的值表示这个方法的附加属性的数量
5、attributes[]表的每一个成员的值必须是attribute，如：Code，Exceptions，Synthetic，Signature，Deprecated，untimeVisibleAnnotations，
RuntimeInvisibleAnnotations，RuntimeVisibleParameterAnnotations，RuntimeInvisibleParameterAnnotations和AnnotationDefault
*/

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members
}

func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}

func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
