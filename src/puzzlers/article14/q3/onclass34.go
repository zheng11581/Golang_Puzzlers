package main

import "fmt"

type Animal interface {
	ScientificName() string
	Category() string
}

type Named interface {
	Name() string
}

type Pet interface {
	Animal
	Named
}

type PetTag struct {
	name  string
	owner string
}

func (pt PetTag) Name() string {
	return pt.name
}

func (pt PetTag) Owner() string {
	return pt.owner
}

type Dog struct {
	PetTag
	scientificName string
}

func (dog Dog) ScientificName() string {
	return dog.scientificName
}

func (dog Dog) Category() string {
	return "dog"
}

func main() {
	petTag := PetTag{name: "little pig"}
	_, ok := interface{}(petTag).(Named)
	fmt.Printf("The PetTag implements interface Named: %v\n", ok)

	dog := Dog{PetTag: petTag, scientificName: "America shorthair"}
	_, ok = interface{}(dog).(Named)
	fmt.Printf("The Dog implements interface Named: %v\n", ok)
	_, ok = interface{}(dog).(Animal)
	fmt.Printf("The Dog implements interface Animal: %v\n", ok)
	_, ok = interface{}(dog).(Pet)
	fmt.Printf("The Dog implements interface Pet: %v\n", ok)
}
