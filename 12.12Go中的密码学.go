package main

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"io"
	"log"
)

//通过网络传输的数据必须加密，以防止被hacker读取或篡改，并且保证发出的数据和收到的数据校验和一致。
//鉴于Go母公司的业务，我们毫不惊讶地看到Go的标准库为该领域提供了超过30个包：
//1.hash包：实现了adler32\crc32\crc64\fnc校验；
//2.crypto包：实现了其他的hash算法，比如 md4\md5\sha1等。以及完整地实现了aes\blowfish\rc4\rsa\xtea等加密算法。
//下面的示例用sha1和md5计算并输出了一些校验值。

func main() {
	hasher := sha1.New()
	io.WriteString(hasher, "test")
	var b []byte
	fmt.Printf("result: %x\n", hasher.Sum(b))
	fmt.Printf("result: %d\n", hasher.Sum(b))
	//
	hasher.Reset()
	data := []byte("we shall overcome!")
	n, err := hasher.Write(data)
	if n!=len(data) || err!=nil {
		log.Printf("hash write error: %v / %v\n", n, err)
	}
	checkSum := hasher.Sum(b)
	fmt.Printf("result: %x\n", checkSum)
	/**
	result: a94a8fe5ccb19ba61c4c0873d391e987982fbbd3
	result: [169 74 143 229 204 177 155 166 28 76 8 115 211 145 233 135 152 47 187 211]
	result: 7969927e8717df4ed8ba0ae2cf7eed5a8580ff64
	*/
	//通过调用sha1.New()创建了一个新的hash.Hash对象，用来计算SHA1校验值。Hash类型实际上是一个接口，它实现了io.Writer接口。
	//...
	//通过io.WriteString或hasher.Write将给定的[]byte附加到当前的hash.Hash对象中。

	//练习12.9
	hasher2 := md5.New()
	hasher2.Write(data)
	var b2 []byte
	checkSum2 := hasher2.Sum(b2)
	fmt.Printf("result: %x\n", checkSum2)
	/**
	result: f9eff88cd28c9d512c92b4d8a9776d2b
	*/
}
