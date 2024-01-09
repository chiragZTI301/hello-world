package main

import "fmt"

func mapBasics() {
	var map_1 map[int]int

	if map_1 == nil {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }

	map_2 := map[int]string{
		90: "Dog",
		91: "Cat",
		92: "Cow",
		93: "Bird",
		94: "Rabbit",
	}

	fmt.Println(map_2[94])

	var map_3 = make(map[float64]string)

	map_3[1.3] = "Dog"
	map_3[1.5] = "Camel"

	fmt.Println(map_3)

	for i, v := range map_3 {
		fmt.Println(i, v)
	}

	delete(map_3, 1.5)
	name, ok := map_3[1.5]

	fmt.Println(name, ok)
}