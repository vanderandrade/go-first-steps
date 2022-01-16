package attributes

import (
	"fmt"
	"errors"
)

const xpto int = 10

const (
	const1 string = "const1"
	const2 string = "const2"
	const3 string = "const3"
)


func Show(show bool) {
	var name string
	name = "Name"
	last_name := "Last Name"

	a := 10
	b := "Hello"
	c := 10.44
	d := false
	e := 'h'
	f := `Hello 
		World
	`

	if show {
		print(name +" "+ last_name + "\n")

		fmt.Printf("%v [%T]\n", a, a)
		fmt.Printf("%v [%T]\n", b, b)
		fmt.Printf("%v [%T]\n", c, c)
		fmt.Printf("%v [%T]\n", d, d)
		fmt.Printf("%v [%T]\n", e, e)
		fmt.Printf("%v [%T]\n", f, f)
	}

	_, _, err := abc(&a)
	if err != nil {
		panic("panico")
	}
}

func Ponteiros() {
	x := 10
	y := &x

	fmt.Println(y) // Address
	fmt.Println(*y) // Value at the address

	var z *int = &x
	fmt.Println(z)
	fmt.Println(*z)

}

func Arrays() {
	var x [10]int
	x[0] = 1
	x[1] = 2

	y := [5]int{1,2,3} // [1, 2, 3, 0, 0]
	fmt.Println(y)
}

func ArraysDinamicos() {
	slice := make([]int, 3)
	slice[0] = 1
	slice[1] = 1
	slice = append(slice, 2,3,4,5,6,7,8,9,10)
	fmt.Println(slice)

	sliceString := []string {
		"Hello",
		"World",
	}
	fmt.Println(sliceString)

	Dicionarios()
}

func Dicionarios() {
	m := make(map[string]int)
	m["first"] = 1
	m["second"] = 2

	fmt.Println(m)
}


func abc(a *int) (int, string, error) {
	if *a < xpto {
		return 0, "", errors.New("Valor menor que xpto!")
	}
	return *a, "a", nil
}
