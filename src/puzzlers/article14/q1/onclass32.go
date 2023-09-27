package main

import "fmt"

type Pet interface {
	Name() string
	Category() string
}

type Dog struct {
	name string // 名字。
}

func (dog *Dog) SetName(name string) {
	dog.name = name
}

func (dog Dog) Name() string {
	return dog.name
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	// 示例1
	dog := Dog{"little pig"}
	fmt.Printf("The dog's name: %q\n", dog.Name())
	var pet Pet = dog
	dog.SetName("monster")
	fmt.Printf("The dog's name: %q\n", dog.Name())
	fmt.Printf("This pet is %s, name is %q\n", pet.Category(), pet.Name())
	fmt.Println()

	// 示例2
	dog1 := Dog{"little pig"}
	dog2 := dog1
	fmt.Printf("The first dog' is %q\n", dog1.name)
	fmt.Printf("The secend dog is %q\n", dog2.name)
	dog1.name = "monster"
	fmt.Printf("The first dog is %q\n", dog1.name)
	fmt.Printf("The secend dog is %q\n", dog2.name)
	fmt.Println()

	// 示例3
	dog = Dog{"little pig"}
	pet = &dog
	_, ok := interface{}(dog).(Pet)
	fmt.Printf("The Dog is implements interface Pet: %v\n", ok)
	_, ok = interface{}(&dog).(Pet)
	fmt.Printf("The Dog is implements interface Pet: %v\n", ok)
	dog.SetName("monster")
	fmt.Printf("The dog's name: %q\n", dog.Name())
	fmt.Printf("This pet is %s, name is %q\n", pet.Category(), pet.Name())
	fmt.Println()

}
