package main

import (
	"fmt"
	"reflect"
	"strings"
)

// User User
type User struct {
	ID   int
	Name string
	Age  int
}

// SetValue 修改结构体值
func SetValue(o interface{}, s interface{}, name string, value interface{}) {
	v := reflect.ValueOf(o)
	t := reflect.TypeOf(s)
	v = v.Elem()
	for i := 0; i < v.NumField(); i++ {
		n := strings.ToLower(t.Field(i).Name)
		name = strings.ToLower(name)
		if n == name {
			f := v.FieldByName(t.Field(i).Name)
			switch f.Kind() {
			case reflect.String:
				{
					f.SetString(value.(string))
				}
			case reflect.Int:
				{
					f.SetInt(int64(value.(int)))
				}
			case reflect.Int64:
				{
					f.SetInt(value.(int64))
					break
				}

			case reflect.Bool:
				{
					f.SetBool(value.(bool))
				}
			}
		}

	}
}

// func setValue(o interface{}, v interface {}) {
// 	v := reflect.ValueOf(o)
// 	v = v.Elem()
// 	f := v.FieldByName(v)
// 	//
// }

// Poni 结构体遍历
func Poni(o interface{}) {
	t := reflect.TypeOf(o)
	v := reflect.ValueOf(o)
	v = v.Elem()
	fmt.Println("v", v)
	fmt.Println("类型：", t)
	// t.Name()
	// v.Type().Elem().Field()
	// fmt.Println("字符串类型：", t.Name())

	for i := 0; i < v.NumField(); i++ {
		switch v.Field(i).Kind() {
		case reflect.String:
			{
				println("t", t.Field(i).Name)
				f := v.FieldByName(t.Field(i).Name)
				f.SetString("hello")
				// v.FieldByName(t.Field(i).Name).SetString("hi")
				// v.FieldByIndex(i).SetString("hi")
				// v.Field(i).SetString("hi")
				fmt.Printf("field, %s %T string \n", v.Field(i), v)
				break
			}
		case reflect.Int:
			{
				println("t int", t.Field(i).Name)
				// fmt.Printf("field,Int %s, struct %T", v.Field(i), v)
			}
		}
	}

	// fmt.Println(v)

	// 可以获取所有属性
	// 获取结构体字段个数：t.NumField()

	// for i := 0; i < t.NumField(); i++ {
	// 		f := t.Field(i)
	// 		v = v.Elem()
	// 		// v.SetString
	// 		fmt.Printf("%s : %v", f.Name, f.Type)
	// 		// 获取字段的值信息
	// 		// Interface()：获取字段对应的值
	// 		val := v.Field(i).Interface()
	// 		// fmt.Println(v.Type().FieldByName(f.Name))
	// 		// v.FieldByName(f.Name)

	// 		// // v.Field(i).SetString("hello")
	// 		// switch f.Type.Elem().Kind() {
	// 		// case reflect.String:
	// 		// 	v.Field(i).SetString("hello")
	// 		// case reflect.Int:
	// 		// 	v.Field(i).SetInt(10)
	// 		// }
	// 		fmt.Println("val :", val)
	// }
	fmt.Println("=================方法====================")
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
		fmt.Println(m.Type)
	}

}
func ParseValue(o interface{}) {
	SetValue(&o, o, "name", "cx")
}
func main() {
	// u := User{1, "5lmh.com", 20}
	//var u User
	//u := User{}
	//var   =
	var u = User{}
	// SetValue(&u)
	//SetValue(&u, u, "name", "cjk")
	//SetValue(&u, u, "ID", 12)
	//ParseValue(u)
	// SetValue(&u, "N)
	fmt.Println(u)
	// fmt.Println(u)

	// fmt.Println("=-===============")

	// Poni(u)

	// fmt.Println("==============")

}
