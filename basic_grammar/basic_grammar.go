package main

import (
	"fmt"
	hello "go_learn/hello"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	//data()

	//cnst()

	//packg()

	//dataType()

	//io()

	//ifelse()

	//for_()

	//slice()

	//str()

	map_()
}

func data() {
	// 可以不指定类型，但是不能改变类型
	// var a int
	// := 只用于函数内部，局部变量
	a := 10
	fmt.Println(a)

	// 改变类型，编译报错
	// a = 10.1

	// 此处是 取反
	a ^= a
	fmt.Println(a)

	var b = 100
	fmt.Println(b)

	aaa := 1
	aaa &^= 2
	println(aaa)

	// change
	num1, num2, num3 := 25, 36, 49
	num1, num2 = num2, num1
	num1, num2, num3 = num3, num2, num1

	//它们的参数支持所有的可比较类型，go中的可比较类型有
	//布尔
	//数字
	//字符串
	//指针
	//通道 （仅支持判断是否相等）
	//元素是可比较类型的数组（切片不可比较）
	//字段类型都是可比较类型的结构体（仅支持判断是否相等）
	_ = min(1, 2, -1, 1.2)
	_ = max(100, 22, -1, 1.12)

	fmt.Println(min(1, 1, 2, '1', 'a'))
}

func cnst() {
	// 常量 只能是基本数据类型
	const (
		cnt     = iota
		n1  int = 100
		n2
		n3
		n4 string = "111"
		n5
		pi float32 = 3.14159
		c  bool    = false
	)

	const (
		nn1 = iota //0
		nn2 = 100  //100
		nn3 = iota //2
		nn4        //3
	)
	const nn5 = iota //0

	// 定义数量级（这里的<<表示左移操作，1<<10表示将1的二进制表示向左移10位，也就是由1变成了10000000000，也就是十进制的1024。
	// 同理2<<2表示将2的二进制表示向左移2位，也就是由10变成了1000，也就是十进制的8。）
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	const (
		aa, bb = iota + 1, iota + 2 //1,2
		cc, dd                      //2,3
		ee, ff                      //3,4
	)

	const (
		Num  = iota<<2*3 + 1 // 1
		Num1                 // 13
		Num2                 // 25
		Num3 = iota          // 3
		Num4                 // 4
	)

	//通过上面几个例子可以发现，iota是递增的，第一个常量使用iota值的表达式，
	//根据序号值的变化会自动的赋值给后续的常量，直到用新的iota重置，
	//这个序号其实就是代码的相对行号，是相对于当前分组的起始行号
}

func packg() {
	fmt.Println(hello.Hello("chen"))

	fmt.Println(hello.Bye("chen"))
}

func dataType() {
	//uint8	无符号 8 位整型
	var ut8 uint8 = 'a'
	println(ut8) // 97
	//uint16	无符号 16 位整型
	//uint32	无符号 32 位整型
	//uint64	无符号 64 位整型
	//int8	有符号 8 位整型
	//int16	有符号 16 位整型
	//int32	有符号 32 位整型
	//int64	有符号 64 位整型
	//uint	无符号整型 至少32位
	//int	整型 至少32位
	//uintptr	等价于无符号64位整型，但是专用于存放指针运算，用于存放死的指针地址。

	//float32	IEEE-754 32位浮点数
	//float64	IEEE-754 64位浮点数

	//complex128	64位实数和虚数
	//complex64	32位实数和虚数
	cmplx := 1i + 8
	println(cmplx)

	// // 字符类型
	//byte	等价 uint8 可以表达ANSCII字符
	//rune	等价 int32 可以表达Unicode字符
	//string	字符串即字节序列，可以转换为[]byte类型即字节切片
	str := "123456aaa"
	println(str[1:3])

	//数组	[5]int，长度为5的整型数组
	ints := [5]int{0, 1, 2, 3, 4}
	for item := range ints {
		println(item)
	}
	str2 := [10]string{"1", "a"}
	for item := range str2 {
		println(item)
	}

	//切片	[]float64，64位浮点数切片
	//映射表	map[string]int，键为字符串类型，值为整型的映射表
	//结构体	type Gopher struct{}，Gopher结构体
	//指针	*int，一个整型指针。
	//函数	type f func()，一个没有参数，没有返回值的函数类型
	//接口	type Gopher interface{}，Gopher接口
	//通道	chan int，整型通道

	//nil
	//短变量不能+nil，因为nil不属于任何类型，编译错误
	//name := nil
}

func io() {
	// 标准
	//var (
	//	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	//	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	//	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")
	//)

	os.Stdout.WriteString("xxx\n")
	//os.Stdout.Write([]byte("xxxxx"))

	println("Hello 世界!")

	// fmt.Println会用到反射，因此输出的内容通常更容易使人阅读，不过性能很差强人意。
	fmt.Printf("%q\n", "hello world")

	fmt.Printf("%f\n", 1e2)
	fmt.Printf("%e\n", 1e2)
	fmt.Printf("%E\n", 1e2)
	fmt.Printf("%g\n", 1e2)

	type person struct {
		name    string
		age     int
		address string
	}
	//var p1 person = person{}
	//fmt.Println(p1)
	fmt.Printf("%v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%#v\n", person{"lihua", 22, "beijing"})
	fmt.Printf("%t\n", true)
	fmt.Printf("%T\n", person{})
	fmt.Printf("%c%c\n", 20050, 20051)
	fmt.Printf("%U\n", '码')
	fmt.Printf("%p\n", &person{}) // 地址

	str := "abcdefg"
	fmt.Printf("%x\n", str)
	fmt.Printf("% x\n", str) // 61 62 63 64 65 66 67

	fmt.Printf("%09d", 1) // 000000001

	fmt.Printf("%s", 1) // %!s(int=1)

	//// 扫描从os.Stdin读入的文本，根据空格分隔，换行也被当作空格
	//func Scan(a ...any) (n int, err error)
	//
	//// 与Scan类似，但是遇到换行停止扫描
	//func Scanln(a ...any) (n int, err error)
	//
	//// 根据格式化的字符串扫描
	//func Scanf(format string, a ...any) (n int, err error)

}

func ifelse() {
	var a int
	fmt.Scanln(&a)

	if a == 1 {
		println(1)
	} else if a == 2 {
		_, err := os.Stderr.WriteString("2")
		if err != nil {
			return
		}
	} else if a == 3 {
		fmt.Println("3")
	} else {
		fmt.Println("other")
	}

}

func for_() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j <= i {
				//print(i, "*", j, "=", i*j, ", ")
				fmt.Printf("%d*%d =%2d  ", j, i, i*j)
			}
		}
		//println()
		fmt.Println()
	}

	sequence := "hello world"
	for index, value := range sequence {
		fmt.Println(index, string(value))
	}
}

func slice() {
	// 如果事先就知道了要存放数据的长度，且后续使用中不会有扩容的需求，就可以考虑使用数组，
	// Go中的数组是值类型，而非引用，并不是指向头部元素的指针
	// 正确示例
	var a [5]int
	a = [5]int{1, 2, 3}
	fmt.Println("数组: ", a)

	nums := [5]int{1, 2, 3}
	for index, item := range nums {
		fmt.Print(index, item, "; ")
	}

	fmt.Println("len: ", len(nums))
	fmt.Println("cap: ", cap(nums))

	// 错误示例
	//l := 1
	//var b [l]int

	_ = new([]int) // 指针
	//slice1 = []int{1, 2, 3} // 值
	//var slice1 []int            // 值
	slice1 := make([]int, 5, 5) // 推荐使用make来创建一个空切片，只是对于切片而言，make函数接收三个参数：类型，长度，容量
	fmt.Println("切片: ", slice1)

	nums1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//nums1 = append(nums1, []int{-1, 0}...)
	nums1 = append([]int{-1, 0}, nums1...)
	fmt.Println("head insert: ", nums1) // [-1 0 1 2 3 4 5 6 7 8 9 10]

	i := 3
	nums1 = append(nums1[:i+1], append([]int{999, 999}, nums1[i+1:]...)...) // 切片 从中间插入
	fmt.Println("middle insert: ", nums1)                                   // i=3，[1 2 3 4 999 999 5 6 7 8 9 10]

	fmt.Println("截取: [start, end)")
	fmt.Println(nums1[4:])
	fmt.Println(nums1[:4])
	fmt.Println(nums1[1:4])
	fmt.Println(append(nums[:2], nums[5:]...)) // nums[:i], nums[i+n:]...

	fmt.Println("delete all: ", nums1[:0]) // 删除

	dest := make([]int, len(nums1))
	println("copy return changed_len", copy(dest, nums1))

	// 二维 切片
	slice2 := make([][]int, 5)
	fmt.Println(slice2[2])
	fmt.Println(append(slice2[2], []int{1, 2, 3}...))
	for i := 0; i < len(slice2); i++ {
		slice2[i] = make([]int, 5) // 需要遍历赋值，否则为长度为0的切片
	}

	fmt.Println("只有切片才能使用拓展表达式, 该特性于Go1.2版本添加，主要是为了解决切片共享底层数组的读写问题")
	fmt.Println("扩展表达式中的max用于限制新生成切片的容量，新切片的容量为max - low, 容量不足时，分配新的底层数组")
	low, high, max_ := 2, 7, 10
	slice3 := slice2[low:high:max_]
	fmt.Println(slice3)

	// 将切片内的值置为 0值
	clear(slice3)

	// 清空
	_ = slice3[:0:0]
}

func str() {
	str1 := `这是一个原生字符串，换行
	tab缩进，\t制表符但是无效,换行
	"这是一个普通字符串"
	
	结束`

	fmt.Println(str1)

	str2 := "this is a string"
	fmt.Println(string(str2[0]))   // 字节
	fmt.Println(string(str2[0:2])) // 字节

	//str2[0] = 'a' // 无法通过编译
	str2 = "new str" //可以覆盖

	//str2 := "this is a string"
	//str3 := string(str2[0:2])
	//str4 := []rune("12")
	//fmt.Println(append([]rune(str3), str4...)) // 字节

	str3 := "this is a string"
	// 显式类型转换为字节切片
	bytes1 := []byte(str3)
	fmt.Println(bytes1)
	// 显式类型转换为字符串
	fmt.Println(string(bytes1))

	var dst string
	desBytes := make([]byte, len(str3))
	copy(desBytes, str3) // [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103]
	dst = string(desBytes)
	//dst = strings.Clone(str3)
	fmt.Println(str3, dst)

	desBytes = append(desBytes, "123"...) //拼接

	desBytes = append(desBytes[1:2], 'a')

	// 性能相对高
	builder := strings.Builder{}
	builder.WriteString("this is a string ")
	builder.WriteString("that is a int")
	fmt.Println(builder.String())

	str4 := "hello 世界！!!" // 中文字符占3个字节，
	for i, w := 0, 0; i < len(str4); i += w {
		r, width := utf8.DecodeRuneInString(str4[i:])
		fmt.Println(string(r))
		w = width
	}

	//for len(str4) > 0 {
	//	r, size := utf8.DecodeRuneInString(str4)
	//	fmt.Printf("%c %v\n", r, size)
	//	str4 = str4[size:]
	//}

	for i, r := range str4 {
		fmt.Printf("%d, %d,%x,%s\n", i, r, r, string(r))
	}
}

func map_() {
	mp := make(map[string]int)
	mp["one"] = 1
	mp["two"] = 2

	fmt.Println(mp)

	for k, v := range mp {
		fmt.Println(k, v)
	}

	delete(mp, "two")
	fmt.Println(mp)

	value, ok := mp["one"]
	fmt.Println("value:", value, ", ok:", ok)

	value, ok = mp["two"]
	fmt.Println("value:", value, ", ok:", ok)

	// 安全的map操作
	m1 := map[string]int{"one": 1, "two": 2}
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[k] = v * 2
	}
}
