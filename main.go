package main

import (
	"errors"
	"log"
)

func main() {
	myMap := make(map[string]int)

	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3
	myMap["four"] = 4

	log.Println(myMap)

	value, err := divide(200.0, 20.0)

	if err != nil {
		panic(err)
	}

	log.Println(value)
}

func divide(num1 float64, num2 float64) (float64, error) {
	value := 0.0

	if num2 == 0.0 {
		return value, errors.New("Divisor cannot be 0.")
	}

	value = num1 / num2

	return value, nil
}
