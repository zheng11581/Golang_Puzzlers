package main

import "fmt"

func main() {
	{
		// 别名类型：和潜在类型是先同类型，可赋值、可比较、可类型转换
		type MyString = string
		str := "BCD"
		myStr1 := MyString(str)
		myStr2 := MyString("A" + str)
		fmt.Printf("%T(%q) == %T(%q): %v\n", str, str, myStr1, myStr1, str == myStr1)
		fmt.Printf("%T(%q) > %T(%q): %v\n", str, str, myStr2, myStr2, str > myStr2)
		fmt.Printf("Type %T is the same as type %T.\n", myStr1, str)

		strs := []string{"E", "F", "G"}
		myStrs := []MyString(strs)
		fmt.Printf("A value of type []Mystring is: %T(%q)\n", myStrs, myStrs)
		fmt.Printf("Type %T is the same as type %T.\n", myStrs, strs)

	}

	{
		// 类型再定义
		type MyString string
		str := "BCD"
		mystr1 := MyString(str)
		mystr2 := MyString("A" + str)
		_ = mystr2
		// fmt.Printf("%T(%q) == %T(%q): %v\n", str, str, mystr1, mystr1, str == mystr1)
		// fmt.Printf("%T(%q) > %T(%q): %v\n", str, str, mystr2, mystr2, str > mystr2)
		fmt.Printf("Type %T is different from %T\n", mystr1, str)

	}

}
