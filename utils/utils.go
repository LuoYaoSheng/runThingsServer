package utils

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"strconv"
	"time"
)

// Hex2Dec 二进制转十进制
func Hex2Dec(val string) int {
	n, err := strconv.ParseUint(val, 16, 32)
	if err != nil {
		log.Println(err)
	}
	return int(n)
}

// IntToBytes 整形转换成字节
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

func StructToString(v interface{}) string {
	b, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
		return err.Error()
	}
	return string(b)
}

// ConvertToBin 将十进制数字转化为二进制字符串
func ConvertToBin(num int) string {
	str := ""
	if num == 0 {
		return "0"
	}
	// num /= 2 每次循环的时候 都将num除以2  再把结果赋值给 num
	for ; num > 0; num /= 2 {
		lsb := num % 2
		// strconv.Itoa() 将数字强制性转化为字符串
		str = strconv.Itoa(lsb) + str
	}
	return str
}

// CheckNum 校验和
func CheckNum(data []byte) int {
	num := 0
	for i := 0; i < len(data); i++ {
		num += int(data[i])
	}
	//fmt.Printf("---------------------num: %d %x \n", num, num)
	//s := fmt.Sprintf("%02x", num)
	//b, _ := hex.DecodeString(s[len(s)-2:])
	//fmt.Println("---------s: ", s, s[len(s)-2:], b, num%256)
	return num % 256
}

const (
	UTF8    = string("UTF-8")
	GB18030 = string("GB18030")
)

func ConvertByte2String(byte []byte, charset string) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

// DateBytes 获取 年月日、时分秒
func DateBytes() []byte {
	year, month, day := time.Now().Date()
	hour, min, sec := time.Now().Clock()
	return []byte{byte(sec), byte(min), byte(hour), byte(day), byte(month), byte(year)}
}
