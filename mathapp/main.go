package main

import (
	"time"
	"fmt"
	"myproject/mymath"
	"myproject/test"
	"net/http"
	// "log"
	"myproject/object"

	"github.com/labstack/echo"
)

var a, b, c int //全局变量是允许声明但不使用的
// Book 注释以被声明的东西开头
type Book struct { //结构体
	title  string
	author string
	id     int
}

func init() { //init函数
	fmt.Println("init func")
}
func main() {
	variablefunc()
	object.Test()
	//启动web服务 begin
	// http.HandleFunc("/", HandleIndex)
	// err := http.ListenAndServe(":8080", nil)
	// if err != nil {
	// 	log.Fatal("Error:", err)
	// }
	//end
	//通过go web 框架 echo 启动web服务 begin
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/users/:id", getUser)
	e.POST("/users", saveUser)
	// e.PUT("/users/:id", updateUser)
	//e.DELETE("/users/:id", deleteUser)
	
	e.Logger.Fatal(e.Start(":8080")) //web启动
	
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	id := c.QueryParam("name")
	fmt.Println(id)
	return c.String(http.StatusOK, "saveSuccess")
}

func HandleIndex(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("PATH:", r.URL.Path)
	fmt.Println("SCHEME:", r.URL.Scheme)
	fmt.Println("METHOD:", r.Method)
	fmt.Fprintf(w, "<h1>Hello world</h1>")
}

func myTest() int {
	return 0
}

func myfunction(num1 int, num2 int) (int, int) {
	return num1, num2
}

func variablefunc() {
	var str = "abc"
	_ = str //空白标识符 抛弃值
	//fmt.Println(str)
	test.Test()
	fmt.Printf("Sqrt(2) = %v\n", mymath.Test(2))
	var v3 = myTest()
	fmt.Println(v3)
	fmt.Println(myfunction(101, 102))
	fmt.Println("***********************")
	var (
		v31 float32
		v32 [3]bool
		v33 []bool
		v34 *bool //指针变量
		v35 struct {
			name string
			age  int
		}
		v36, v37 map[string]bool
		v38      [2][2]int
	)
	v31 = 2.3
	v32 = [3]bool{false, true, false} //数组
	v32[0] = true
	v33 = v32[1:3]            //从下标startIndex到endIndex-1 下的元素创建为一个新的切片
	v331 := make([]string, 3) //创建元素个数为3的切片
	v332 := append(v33, true) //切片添加元素
	v333 := make([]bool, len(v33), (cap(v33))*2)
	copy(v333, v33) //slice切片跟数组类型是不同的，换成v32 报错second argument to copy should be slice or string; have [3]bool
	v34 = &v32[0]   //通过 &i 来获取变量 i 的内存地址
	v35.name = "Sam"
	v35.age = 25
	v36 = make(map[string]bool)
	v36["k1"] = false
	v36["k2"] = true
	v36["k3"] = true
	delete(v36, "k3") //map的delete()函数
	v361, ok := v36["k1"]
	v37 = make(map[string]bool, 10)
	v37 = map[string]bool{"h1": false, "h2": true}
	v38 = [2][2]int{{1, 2}, {3, 4}}
	v39 := map[string]string{"key1": "a", "key2": "b"}
	v40 := v38[0] //:=使用变量的首选形式 声明赋值 但是不可再次对该变量使用:=
	fmt.Println("v31 is", v31)
	fmt.Println("v32 is", v32)
	fmt.Println("v33 is", v33, "space", cap(v33), "lenth", len(v33)) //分配的空间大小cap
	fmt.Println("v331 is", v331)
	fmt.Println("v332 is", v332)
	fmt.Println("v333 is", v333, "space", cap(v333), "lenth", len(v333))
	fmt.Println("v34 is", v34)
	fmt.Println("*v34 is", *v34)
	fmt.Println("v35 is", v35)
	fmt.Println("v36 is", v36, "lenth is", len(v36))
	fmt.Println("v361 is", v361)
	fmt.Println("v361's ok is", ok)
	fmt.Println("v37 is", v37)
	fmt.Println("v38 is", v38)
	fmt.Println("v39 is", v39)
	fmt.Println("v40 is", v40)
	fmt.Println("***********************")
	var vname1, vname2, vname3 int //多变量声明
	vname1, vname2, vname3 = 1, 2, 3
	fmt.Println(vname1, vname2, vname3)
	//常量
	const v41 int = 41 //将变量中的var替换成const就是在定义常量
	//位运算
	v51 := 1
	v52 := 2
	v53 := v51 << 2  //4
	v54 := v51 ^ v52 //3
	v55 := v51 & v53 //0
	v56 := v51 | v53 //5
	v57 := ^v51      //-2

	fmt.Println("v51 is", v51)
	fmt.Println("v52 is", v52)
	fmt.Println("v53 is", v53)
	fmt.Println("v54 is", v54)
	fmt.Println("v55 is", v55)
	fmt.Println("v56 is", v56)
	fmt.Println("v57 is", v57)
	fmt.Println("***********************")
	//遍历
	for i := 0; i < len(v32); i++ {
		fmt.Println(v32[i])
	}
	for i, item := range v32 {
		if item {
			fmt.Println("index:", i)
		}
	}
	for k, v := range v39 {
		fmt.Printf("%s -> %s\n", k, v)
	}
	//
	var book1 Book
	book1.title = "GOLANG"
	book1.author = "xxx"
	book1.id = 1

	mymath.TypeChange() //应用包中方法名称首字母大写，则其他包可见，否则不可见
	// a := time.Now()
	fmt.Println(time.Second)
}
