package main


type itype int
const len = 5

const (
	i1 itype = iota
	i2
	i3
	i4
	i5
)

const (
	val1 itype = iota
	val2
)

func main() {
	//注：itype替换为基本类型
	var start = 2
	var end = 4
	//创建:
	arr1 := new([len]itype)
	slice1 := make([]itype, len)

	//初始化
	arr2 := [...]itype{i1,i2,i3,i4,i5}
	arrKeyValue := [len]itype{i1: val1, i2: val2}//初始化部分i2初始化为val2...
	var slice1 []itype = arr1[start:end]
	//1.如何截断数组或切片的最后一个元素
	line = line[:len(line)-1]
	//2.如何使用for或者for-range遍历一个数组（或切片）
	for i:=0;i<len(arr);i++{
		//...arr[i]
	}
	for ix, value := range arr {
		//...
	}
	//3.如何在一个二维数组或切片arr2Dim中查找一个指定值V
	found := false
	Found: for row := range arr2Dim {
		for column := range arr2Dim[row] {
			if arr2Dim[row][column] == V {
				found = true
				break Found
			}
		}
	}
}