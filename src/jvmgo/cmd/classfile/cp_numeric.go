package classfile

import "math"

//CONSTANT_Intrger_info 和 CONSTANT_Float_info 结构表示 4 字节（int 和 float）的数值常量
/*
CONSTANT_Integer_info {
	u1 tag;
	u4 bytes;
}
存储顺序按照Big-Endian
*/

type ConstantIntegerInfo struct {
	value int32
}

func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.value = int32(bytes)
}

/*
CONSTANT_Float_info {
	u1 tag;
	u4 bytes;
}
存储顺序按照Big-Endian，bytes按照IEEE754单精度浮点格式，详见Java虚拟机规范
*/

type ConstantFloatInfo struct {
	value float32
}

func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.value = math.Float32frombits(bytes)
}

//CONSTANT_Long_info 和 CONSTANT_Double_info 结构表示 8 字节（long 和 double） 的数值常量
/*
CONSTANT_Long_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
构造形式为((long) high_bytes << 32) + low_bytes，按照Big-Endian顺序存储，详见Java虚拟机规范
*/
type ConstantLongInfo struct {
	value int64
}

func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.value = int64(bytes)
}

/*
CONSTANT_Double_info {
	u1 tag;
	u4 high_bytes;
	u4 low_bytes;
}
high_bytes 和 low_bytes 共同按照 IEEE 754双精度浮点格式，按照Big-Endian顺序存储，详见Java虚拟机规范
*/
type ConstantDoubleInfo struct {
	value float64
}

func (self *ConstantDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.value = math.Float64frombits(bytes)
}
