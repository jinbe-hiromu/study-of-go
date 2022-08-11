package main

import (
	"fmt"
)

// interface: 異なる型に共通の性質を付与する
type Stringfy interface {
	ToString() string // ToStringメソッドを持つもの＝Stringfy
}

type Car struct {
	Number int
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type Person struct {
	Name string
	Age  int
}

func (c *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", c.Name, c.Age)
}

// interface Stringer
func (p *Person) String() string{
	return fmt.Sprintf("<<%v,%v>>",p.Name,p.Age)
}

func main() {
	stringfy := []Stringfy{
		&Person{Name: "John", Age: 100},
		&Car{Number: 123, Model: "TeslaModel"},
	}

	for _, c := range stringfy {
		fmt.Println(c.ToString())
	}

	// interface Stringer
	p := &Person{Name: "John", Age:100}
	fmt.Println(p)
}
