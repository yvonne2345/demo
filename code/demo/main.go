package main

import (
	"code/test"
	"fmt"
	"strconv"
)

func main() { //基本类型、数组、切片、map、指针
	fmt.Println("2022/9/1")
	fmt.Println(test.Name) //引用别的包的变量要求是大写开头,小写开头是private不能被引用

	var num int = 9223372036854775807 //int表示int64
	var str string = "abc"
	fmt.Println(num)
	fmt.Printf("%T %T", num, str) //fmt.Printf类似于反射,输出的是对应的类型
	fmt.Println()

	str1 := "123"                 //数据类型转换api
	int1, _ := strconv.Atoi(str1) //string->int
	fmt.Println(int1)
	fmt.Printf("%T", int1)
	fmt.Println()
	int2, _ := strconv.ParseInt(str1+"1", 10, 64) //string->int64
	fmt.Println(int2)
	fmt.Printf("%T", int2)
	fmt.Println()
	str2 := strconv.Itoa(int1 + 2) //int->string
	fmt.Println(str2)
	fmt.Printf("%T", str2)
	fmt.Println()
	str3 := strconv.FormatInt(int2+3, 10) //int64->string
	fmt.Println(str3)
	fmt.Printf("%T", str3)
	fmt.Println()
	float1, _ := strconv.ParseFloat(str3+"5", 32) //float->string
	fmt.Println(float1)
	fmt.Printf("%T", float1)
	fmt.Println()

	a := 1 //流程控制语句
	switch a {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
		fallthrough //可以直接输出下一个条件
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("不满足")
	}
	c := 0
	for {
		c++
		fmt.Println(c)
		if c > 10 {
			break //break结束当前循环，continue跳出本次循环，进入下一个循环
		}
	}
	for b := 0; b < 10; b++ {
		fmt.Println(b)
	}
A:
	for d := 0; d < 10; d++ {
		fmt.Println(d)
		if d == 5 {
			break A
			goto B //跳转语句goto
		}
	}
B:
	fmt.Println("B")

	e := [3]int{0, 1, 2} //数组
	f := [...]int{1, 2, 3, 4, 5, 6, 7}
	var g = new([10]int)
	g[2] = 2
	fmt.Println(e, f, g)
	h := [...]string{"a", "b", "c", "d"}
	for i := 0; i < len(h); i++ { //数组循环  fori
		fmt.Println(h[i])
	}
	for i, v := range h { //forr
		fmt.Println(i, v)
	}

	a1 := [3]int{0, 1, 2} //切片slice,是数组的一部分
	sl := a1[:]
	sl1 := a1[2:]
	sl = append(sl, 5)
	copy(sl[1:2], sl1)
	fmt.Println(sl, sl1)
	fmt.Println(len(sl), cap(sl))
	var aa []int
	aaa := make([]int, 4) //有默认值
	fmt.Println(aa, aaa)

	var m map[string]string //map声明,key、value可以接任何类型
	m = map[string]string{}
	m["a"] = "a"
	m["b"] = "b"
	m1 := map[string]string{}
	m1["c"] = "c"
	m1["d"] = "d"
	m2 := make(map[string]string)
	m2["e"] = "e"
	m2["f"] = "f"
	fmt.Println(m, m1, m2)
	m3 := map[int]interface{}{} //可以接任意类型
	m3[0] = "c"
	m3[1] = 1
	m3[2] = true
	m3[3] = []int{1, 2, 3}
	fmt.Println(m3)
	delete(m3, 2) //删除delete(对象, key)
	fmt.Println(len(m3))
	for i, i2 := range m3 { //循环
		fmt.Println(i, i2)
	}

	var array1 = [5]string{"1", "2", "3", "4", "5"} //数组
	var array1P *[5]string                          //数组指针—指向数组的指针
	array1P = &array1                               //&指向array1地址
	fmt.Println(array1, array1P)
	string1 := "str1"
	string2 := "str2"
	string3 := "str3"
	string4 := "str4"
	string5 := "str5"
	var array2 = [5]*string{&string1, &string2, &string3, &string4, &string5} //指针数组—指向指针的数组,地址的数组
	*array2[2] = "string2"
	for i, v := range array2 {
		fmt.Println(i, *v) //*依次打印地址里的值
	}
	fmt.Println(array2) //打印地址
}
