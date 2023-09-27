package main

import "fmt"

type Pet interface {
	SetName(name string)
	Name() string
	Category() string
	ScientificName() string
}

type Cat struct {
	scientificName string
	category       string
	name           string
}

func New(scientificName, category, name string) Cat {
	return Cat{
		scientificName: scientificName,
		category:       category,
		name:           name,
	}
}

func (cat Cat) Name() string {
	return cat.name
}

func (cat Cat) Category() string {
	return cat.category
}

func (cat Cat) ScientificName() string {
	return cat.scientificName
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}

func (cat Cat) SetNameOfCopy(name string) {
	cat.name = name
}

func (cat Cat) String() string {
	return fmt.Sprintf("%q (category: %s, scientificName: %s)", cat.name, cat.category, cat.scientificName)
}

func main() {
	cat := New("America shorthair", "cat", "little pig")
	cat.SetName("littile monster")
	fmt.Printf("The cat: %s\n", cat)

	cat.SetNameOfCopy("little pig")
	fmt.Printf("The cat: %s\n", cat)

	_, ok1 := interface{}(cat).(Pet)
	_, ok2 := interface{}(&cat).(Pet)

	fmt.Printf("Cat implements interface Pet: %v\n", ok1)
	fmt.Printf("*Cat implements interface Pet: %v\n", ok2)
}
