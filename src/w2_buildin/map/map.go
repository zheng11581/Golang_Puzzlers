package main

import "fmt"

func main() {
	fmt.Println("Creating map")
	m := map[string]string{
		"name": "zhenghc",
		"sex":  "male",
	}
	var m1 map[string]string         // Zero value of map is nil
	m2 := make(map[string]string, 0) // empty map
	fmt.Println("m1=", m1)
	fmt.Println("m2=", m2)

	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Printf("m[%q]==%q\n", k, v)
	}

	fmt.Println("Getting value")
	name, ok := m["name"]
	fmt.Printf("name=%s, %v\n", name, ok)

	if sex, ok := m["sax"]; ok { // 判断key是否存在
		fmt.Printf("sex=%s", sex)
	} else {
		fmt.Println("key dose not exist")
	}

	fmt.Println("Deleting value")
	sex, ok := m["sex"]
	fmt.Printf("sex=%q, %v\n", sex, ok)

	delete(m, "sex")

	sex, ok = m["sex"]
	fmt.Printf("sex=%q, %v\n", sex, ok)

	fmt.Println("Length of non-repeat sub-string")

	fmt.Println("length:=", lenOfNonRepeatSubStr("abcdabcssddde")) // maxLen = 5
	fmt.Println("length:=", lenOfNonRepeatSubStrGlobal("我爱北京天安门")) // maxLen = 7

}
