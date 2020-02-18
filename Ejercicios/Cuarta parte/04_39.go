package main

import "fmt"

func main39() {

	var day int
	fmt.Println("Por favor ingrese el numero del dia")
	fmt.Scanln(&day)
	var result bool
	var daysName [7]string = [7]string{"Domingo", "Lunes", "Martes", "Miércoles", "Jueves", "Viernes", "Sábado"}
	var i int

	for i = 0; i < len(daysName); i++ {
		if (dayOfWeek(day)) == daysName[i] {
			result = true
			break
		}
	}

	if result {
		fmt.Printf("El caso %d -> %s", day, inGreen("Ok"))
	} else {
		fmt.Printf("El caso %d -> %s", day, inRed("ERROR"))
	}

}
