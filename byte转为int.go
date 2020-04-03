// byte转为int len>0
func BytesToInt2(arr []byte) int {
	if len(arr) > 1 {
		//if len(arr) < 4 {
			var md int = 0
			for i :=0;i<len(arr);i++ {
				fmt.Printf("%b\n", arr[i])
				md = md | int(uint8(arr[i]))
				fmt.Printf("%b\n", int(uint8(arr[i])))
				fmt.Printf("|运算：%b\n", md)
				if i != len(arr)-1 {
					md = md<<8
					fmt.Printf("<<运算：%b\n", md)
				}
				
			}
			fmt.Printf("\n")
			return md
		/*}else {
			return BytesToInt(arr)
		}*/
	}else {
		return int(uint8(arr[0]))
	}
}

//整形转换成字节  len>0
func IntToBytes(n int) []byte {
    x := int32(n)

    bytesBuffer := bytes.NewBuffer([]byte{})
    binary.Write(bytesBuffer, binary.BigEndian, x)
    return bytesBuffer.Bytes()
}

//字节转换成整形  len>4
func BytesToInt(b []byte) int {
    bytesBuffer := bytes.NewBuffer(b)

    var x int32
    binary.Read(bytesBuffer, binary.BigEndian, &x)

    return int(x)
}
