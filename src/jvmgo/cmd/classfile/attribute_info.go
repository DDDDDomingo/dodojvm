package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
    readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo{
    attributesCount := reader.readUint16()
    att
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo{

}

func newAttributeInfo(attrName string, attrlen uint32, cp ConstantPool) AttributeInfo{

}