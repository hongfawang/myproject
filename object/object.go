package object

import "fmt"

type Age int

func (a *Age) add(target Age) { //a是指针
	*a += target //*a是指针的值
}

func (a Age) errorAdd(target Age) { //a是值
	a += target //值的改变不会改变指针的值
}

type Book struct {
	name   string
	author string
}

func (this Book) setName1(value string) {
	this.name = value
}

func (this *Book) setName2(value string) {
	this.name = value
}

func Test() {
	var v1 Age = 5
	v1.add(10)
	fmt.Println(v1) //15
	v1.errorAdd(10)
	fmt.Println(v1) //15
	//v2 := &Book{"golang","aaa"} //类的地址
	v2 := new(Book) //new是用来分配内存的内建函数，返回指针，指向新分配的零值 等价于v3 := &Book{}
	v2.name = "golang"
	v2.setName1("golang1")
	fmt.Println(*v2)
	v2.setName2("golang2")
	fmt.Println(*v2)
	hello1()
	hello2()
	hello3()
}

var a string

func f() {
	fmt.Println(a)
}
func hello1() {
	a = "hello!"
	go f()
}

func hello2() {
	go func() { //对 a 进行赋值后并没有任何同步事件，因此它无法保证被其它任何Go程检测到
		a = "hello"
	}()
	fmt.Println(a)
}

var c = make(chan int, 10)

func f1() {
	a = "hello world"
	c <- 0
}
func hello3() {
	go f1()
	<-c
	fmt.Println(a)
}
