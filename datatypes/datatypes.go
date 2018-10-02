package main

import "fmt"

func main() {

	// Uint
	var u8 uint8 = 1
	var u16 uint16 = 1
	var u32 uint32 = 1
	var u64 uint64 = 1
	var u uint = 1

	fmt.Printf("u8=%d, u16=%d, u32=%d, u64=%d, u=%d\n", u8, u16, u32, u64, u)

	// int
	var i8 int8 = 1
	var i16 int16 = 1
	var i32 int32 = 1
	var i64 int64 = 1
	var i int = 1

	fmt.Printf("i8=%d, i16=%d, i32=%d, i64=%d, i=%d\n", i8, i16, i32, i64, i)

	// float
	f32 := float32(1)
	f64 := float64(1)

	fmt.Printf("f32=%f, f64=%f\n", f32, f64)

	// String
	var firstName string = "vikram"
	lastName := "jakhar"

	fmt.Printf("%s %s\n", firstName, lastName)

	// boolean
	var b bool = true

	fmt.Println(b)

	// struct

	type User struct {
		firstName string
		lastName string
		age int
	}

	user := User{firstName:"vikram", lastName:"jakhar", age:26}

	fmt.Println(user)
	fmt.Println(user.firstName)

	// own type
	type MyType string

	var me MyType = MyType("Aaaalee")

	fmt.Println(me)

	// map
	var m1 map[string]string
	m1 = make(map[string]string)
	m1["firstName"] = "vikram"
	m1["lastName"] = "jakhar"

	fmt.Println(m1["firstName"], m1["lastName"])

	m2 := map[string]string{"firstName":"vikram"}
	fmt.Println(m2)

	// interface type
	var ifType interface{} = "blahblah"
	fmt.Println(ifType)
}
