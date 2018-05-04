package test

import (
	"fmt"
)

type Phone interface { //这是接口
	call() string
}

type NokiaPhone struct { //这是实现接口的结构体

}

func (nokiaPhone *NokiaPhone) call() string { //这是实现接口方法，就是实现接口Phone
	fmt.Println("i am a nokiaPhone, call()")
	return ""
}

type Iphone struct {
}

func (iphone *Iphone) call() string { 
	fmt.Println("i am a iphone, call()")
	return ""
}

func Test() {
	fmt.Println("well done")
	fmt.Println(GetName())
	var phone Phone
	phone = new(NokiaPhone) //phone = new(NokiaPhone) 这是实现接口 type *NokiaPhone implement Phone 
	phone.call()
	phone = new(Iphone) //cannot use new(Iphone) (type *Iphone) as type Phone in assignment:*Iphone does not implement Phone (missing call method)
	phone.call()
}

func GetName() (string, string, string) {
	return "name1", "name2", "name3"
}

