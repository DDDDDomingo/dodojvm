package classfile

import (
	"fmt"
	"unicode/utf16"
)

/*
CONSTANT_Utf8_info {
	u1 tag;
	u2 length;
	u1 bytes[length];
}
1、tag项对应CONSTANT_Utf8
2、length项指名了bytes[]数组的长度，以length属性确定长度
3、bytes[]表示字符串值的byte数组，bytes[]数组中每个成员的byte值都不会是0，也不在0xf0至0xff范围内
Tips: Java的字符串常量可以包含ASCII中的所有非空字符和所有Unicode编码的字符，一个字符一个Byte
\u0000 表示字符null
\u0001 至 \u007F内的字符占1byte，x[0xxx xxxx]
\u0080 至 \u07FF的字符占用2byte，由x[110x xxxx]和y[10xx xxxx]表示，字符值为 ((x & 0x1f) << 6) +(y & 0x3f)
\u0800 至 \uFFFF的字符占用3byte，由x[1110 xxxx]、y[10xx xxxx]、z[10xx xxxx]表示，字符值为 ((x & 0xf) << 12) + ((y & 0x3f) << 6) + (z & 0x3f)
超过\uFFFF的字符称为补充字符，占用6byte，由u[1110 1101]，v[1010 xxxx]，w[10xx xxxx]，x[1110 1101]，y[1011 xxxx]，z[10xx xxxx]，字符值为 0x10000+((v&0x0f)<<16)+((w&0x3f)<<10)+(y&0x0f)<<6)+(z&0x3f)
*/

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)

}

func decodeMUTF8(bytes []byte) string {
	utflen := len(bytes)
	chararr := make([]uint16, utflen)

	var c, char2, char3 uint16
	count := 0
	chararr_count := 0

	for count < utflen {
		c = uint16(bytes[count])
		if c > 127 {
			break
		}
		count++
		chararr[chararr_count] = c
		chararr_count++
	}

	for count < utflen {
		c = uint16(bytes[count])
		switch c >> 4 {
		case 0, 1, 2, 3, 4, 5, 6, 7:
			// 0xxx xxxx [Unicode符号范围 0000 0000-0000 007F]
			count++
			chararr[chararr_count] = c
			chararr_count++
		case 12, 13:
			// 110x xxxx   10xx xxxx [Unicode符号范围 0000 0080-0000 07FF]
			count += 2
			if count > utflen {
				panic("malformed input: oartial character at end")
			}
			char2 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input around byte %v", count))
			}
			chararr[chararr_count] = c&0x1F<<6 | char2&0x3F
			chararr_count++
		case 14:
			// 1110 xxxx  10xx xxxx  10xx xxxx [Unicode符号范围 0000 0800-0000 FFFF]
			count += 3
			if count > utflen {
				panic("malformed input: partial character at end")
			}
			char2 = uint16(bytes[count-2])
			char3 = uint16(bytes[count-1])
			if char2&0xC0 != 0x80 || char3&0xC0 != 0x80 {
				panic(fmt.Errorf("malformed input arount byte %v", count-1))
			}
			chararr[chararr_count] = c&0x0F<<12 | char2&0x3F<<6 | char3&0x3F<<0
			chararr_count++
		default:
			//10xx xxxx  1111 xxxx
			panic(fmt.Errorf("malformed input around byte %v", count))
		}
	}
	chararr = chararr[0:chararr_count]
	runes := utf16.Decode(chararr)
	return string(runes)
}
