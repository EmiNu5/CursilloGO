package main

import "fmt"

func main38() {
	var day int

	fmt.Println("Por favor ingrese el numero del dia")
	fmt.Scanln(&day)

	fmt.Println(dayOfWeek(day))
}
func dayOfWeek(day int) string {
	var givenDay string
	switch day {
	case 1:
		givenDay = "Domingo"
	case 2:
		givenDay = "Lunes"
	case 3:
		givenDay = "Martes"
	case 4:
		givenDay = "Miércoles"
	case 5:
		givenDay = "Jueves"
	case 6:
		givenDay = "Viernes"
	case 7:
		givenDay = "Sábado"
	default:
		givenDay = "Error"
	}
	return givenDay
}
