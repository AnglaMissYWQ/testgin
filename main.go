package testgin

import (
	"fmt"
	"reflect"
)

var _VERSION_ = "unknown"

type testInterface interface {
	Message()
	Success()
}

type test1 struct {
	Name string
	Age  int
}

func (test1) Message() {
	fmt.Println("test1")
}

func (test1) Success() {
	fmt.Println("test2")
}

type test2 struct {
	Name string
	Age  int
}

func (test2) Message() {

	fmt.Println("test11111111111111")
}

func (test2) Success() {
	fmt.Println("test222222222222222")
}

func onegoes(i testInterface) {
	i.Message()
}

type myint int

func (myint) Abccc() {

}
func main() {
	type interfaceA interface{}
	type interfaceB interface {
		Echo()
	}

	var abc myint = 123

	myReflect := reflect.TypeOf(abc)

	fmt.Println(myReflect.Name(), myReflect.Kind())
	// name 是当前的类型  包含包名
	// kind 是底层数组类型

	//是否实现了某些接口
	var (
		interfaceAa *interfaceA
	// interfaceBb interfaceB  , myReflect.Implements(interfaceBb)
	)

	testTmp := reflect.TypeOf((*interfaceB)(nil)).Elem()

	fmt.Println(myReflect.Implements(testTmp))

	//判断当前变量是否能赋值给类型为u的类型变量
	fmt.Println(myReflect.AssignableTo(testTmp))

	//判断当前类型是否支持比较
	fmt.Println(myReflect.Comparable())

	//判断当前类型是否能转换成U类型
	fmt.Println(123)
	fmt.Println(myReflect.ConvertibleTo(reflect.TypeOf(interfaceAa).Elem()))
	fmt.Println(myReflect.ConvertibleTo(testTmp))

	// 返回一个类型的方法个数
	fmt.Println(myReflect.NumMethod())

	//通过索引访问方法
	fmt.Println(myReflect.Method(0))

	// myReflect.Comparable(reflect.TypeOf((*interfaceB)(nil)).Elem())
	// go1 := test1{"123", 12}
	// go2 := test2{"123", 12}

	// onegoes(go1)
	// onegoes(go2)

	// r := gin.Default()

	// r.Run()

}
