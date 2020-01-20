package classfile

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
这里还包含一些属性的校验，详见Java虚拟机规范4.7.1
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)
	for i := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)
	attrLen := reader.readUint32()
	attrInfo := newAttributeInfo(attrName, attrLen, cp)
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrlen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Code":
		return &CodeAttribute{cp: cp}
	//case "StackMapTable": return
	case "Exceptions":
		return &ExceptionAttribute{}
	//case "InnerClasses": return
	//case "EnclosingMethod":
	case "Synthetic":
		return &SyntheticAttribute{}
	//case "Signature"
	//case "SourceFile"
	//case "SourceDebugExtension"
	//case "LineNumberTable"
	//case "LocalVariableTable"
	//case "LocalVariableTypeTable"
	case "Deprecated":
		return &DeprecatedAttribute{}
	//case "RuntimeVisibleAnnotations"
	//case "RuntimeInvisibleAnnotations"
	//case "RuntimeVisibleParameterAnnotations"
	//case "RuntimeInvisibleParameterAnnotations"
	//case "AnnotationDefault"
	//case "BootstrapMethods"
	default:
		return &UnparsedAttribute{attrName, attrlen, nil}
	}
}
