package classfile

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
1、attribute_name_index 项的值必须是对常量池的有效索引，常量池在该索引处的项必须是 CONSTANT_Utf8_info（§4.4.7）结构，表示字符串“Code”。
2、attribute_length 项的值表示当前属性的长度，不包括开始的 6 个字节（attribute_name_index和attribute_length）。
3、max_stack 项的值给出了当前方法的操作数栈在运行执行的任何时间点的最大深度(long、double类型2个单位，其他类型一个单位)
4、max_locals的值给出了分配在当前方法引用的局部变量表中的局部变量个数，包括调用此方法时用于传递参数的局部变量。long 和 double 型的局部变量的
最大索引是 max_locals-2，其它类型的局部变量的最大索引是 max_locals-1.
5、code_length 项给出了当前方法的 code[]数组的字节数，code_length 的值必须大于 0，即 code[]数组不能为空。
6、code[]数组给出了实现当前方法的 Java 虚拟机字节码。
7、exception_table_length项的值给出了 exception_table[]数组的成员个数量
8、exception_table[]数组的每个成员表示 code[]数组中的一个异常处理器（Exception Handler）。exception_table[]数组中，异常处理器顺序=异常处理顺序；
    8.1、start_pc 和 end_pc两项的值表明了异常处理器在 code[]数组中的有效范围。即设 x 为异常句柄的有效范围内的值，x 满足：start_pc ≤ x < end_pc。
        TIPS:end_pc值本身不处于异常处理器，如果一个code的长度刚好为65535个字节，并且以一个1字节长度的指令结束，那么这条指令将不能被异常处理器所处理。不过编译器可以限制code[]数组的最大长度
    8,2、handler_pc 项表示一个异常处理器的起点，它的值必须同时是一个对当前 code[]数组中某一指令的操作码的有效索引。
    8.3、catch_type=0,表示对常量池的有效索引，且必须是CONSTANT_Class_info结构，表示当前异常处理器指定需要捕捉的异常类型，当且仅当当前抛出的异常是指定的类或其子类的实例时，才会被调用
        catch_type!=0,表示这个异常处理器将会在所有异常抛出时都被调用。可用于实现finally语句
9、attributes_count 项的值给出了 Code 属性中 attributes 表的成员个数。
10、attributes[]
*/

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
