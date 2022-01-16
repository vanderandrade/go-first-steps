package types

import (
	"fmt"
	"encoding/json"
)

type CarName string
type CarYear string

type vehicle interface {
	start() string
}

type Car struct {
	CarName string `json:"car"` //Adicionado uma tag para, em funções json, usá-la no lugar do nome da variável
	CarYear int `json:"year"`
}

func (car Car) start() string { // Não funciona com o ponteiro do car, fiquei bolado
	return "Start!"
}

func (c *Car) Drive() {
	fmt.Println(c.CarName, "andou!")
}

func (c *Car) ToJson(log bool) string {
	result, _ := json.Marshal(c)
	if log {
		fmt.Println(string(result))
	}
	return string(result)
}

func Tipos() {
	var x CarName
	x = "Fusion"

	car1 := Car {
		CarName: string(x),
		CarYear: 2020,
	}

	car1.Drive()
	car1_json := car1.ToJson(false)
	fmt.Println(car1_json)

	car2_json := []byte (`{"car": "Civic", "year": 1994}`)
	var car2 Car
	json.Unmarshal(car2_json, &car2)
	showVehicle(car2)
}

func showVehicle(h vehicle) {
	fmt.Println(h)
}