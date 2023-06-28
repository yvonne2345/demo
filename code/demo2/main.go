package main

import "fmt"

func main() { //函数、接口(结构体)
	//func 函数名(参数列表)(返回值列表){————函数,函数不需要对象就可以调用
	//	函数体
	//}
	//func (对象)函数名(参数列表)(返回值列表){————方法,方法需要通过对象来调用
	//	函数体
	//}
	res1, res2 := f1(1, "a") //普通函数
	fmt.Println(res1, res2)
	bb := func(data1 string) { //匿名函数
		fmt.Println(data1)
	}
	bb("匿名")
	arr := []string{"1", "2", "3", "4"}

	mo(789, arr...) //不定项参数传参
	func() {
		fmt.Println("自执行") //自执行函数,定时任务
	}()

	mo1()(4) //函数的返回可以是函数

	//接口,规范
	m := Method1{Worker: "zhangsan", Bug: "3个"}
	var test method //定义一个接口test
	test = m
	test.write() //接口没有结构体的属性,但结构体实现了接口,通过接口,可以调用到原始实例的方法
	test.use()
	c := Method1{
		Worker: "c",
		Bug:    "3",
	}
	MyFun(c)
	L.write()

	nextNumber := getSequence() /* nextNumber 为一个函数，函数 i 为 0 */
	fmt.Println(nextNumber())   /* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber) //nextNumber()是值,nextNumber是匿名函数
	fmt.Println(getSequence()())
	fmt.Println(getSequence())
	nextNumber1 := getSequence() /* 创建新的函数 nextNumber1，并查看结果 */
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
	/**
	getSequence() 返回的是一个函数，因此nextNumber :=的结果其实是一个匿名函数（而不是getSequence()这个函数）
	此时，当前的i 值带入函数，得到的结果被打印出来。
	因此fmt.Println(nextNumber()) ，第一次i=0，结果是1。第二次i=1，结果是2……一直调用的是getSequence()里的匿名函数，而非getSequence()
	如果修改成fmt.Println(getSequence()())则不管打印多少次都是1.
	*/

}

func f1(data1 int, data2 string) (resp1 int, resp2 string) { //方法名小写开头私有,大写可以被调用
	resp1 = data1
	resp2 = data2
	return
}

func mo(data1 int, data2 ...string) { //不定项参数传参
	fmt.Println(data1, data2)
	fmt.Printf("%T", data2)
	fmt.Println()
}

func mo1() func(int2 int) { //函数的返回可以是函数
	return func(int2 int) {
		fmt.Println(int2)
	}
}

// 1、空接口可以接任何参数interface{} 2、
type method interface { //接口,声明接口用type,实现了该接口就属于该接口
	write()
	use()
}
type Method1 struct { //结构体
	Worker string
	Bug    string
}
type Method2 struct { //结构体
	Worker string
}

func (m1 Method1) write() { //直接实现  方法-方法前面那个小括号里面的参数称之为接受者，接受者决定了这个方法属于哪个结构体
	fmt.Println(m1.Worker, "写")
}
func (m1 Method1) use() {
	fmt.Println(m1.Bug, "写")
}
func (m2 Method2) write() {
	fmt.Println(m2.Worker, "用")
}
func (m2 Method2) use() {
	fmt.Println(m2.Worker, "用")
}

var L method

func MyFun(a method) {
	L = a
}

func getSequence() func() int { //闭包函数，返回是函数---类似Java中的lambda表达式
	i := 0
	return func() int {
		i += 1
		return i
	}
}
