package main

import "fmt"

func main() {
	var i int
	var j int
	var k int
	const cant int = 3
	var strings [cant]string
	var string1 string
	var stringelegido string
	var coincidencia bool

	fmt.Println("Ingresar un string")
	fmt.Scan(&string1)

	fmt.Println("Ingresar los", cant, " strings")

	for i = 0; i < cant; i++ {
		fmt.Scan(&strings[i])
	}

	for i = 0; i < len(strings); i++ {
		stringelegido = strings[i]
		for j = 0; j < len(string1); j++ {
			coincidencia = true
			for k = 0; k < len(stringelegido); k++ {
				if stringelegido[k] != string1[j+k] {
					coincidencia = false
					break
				}
			}
			if coincidencia {
				break
			}
		}
		if coincidencia {
			break
		}
	}
	if coincidencia {
		fmt.Println("Al menos un strings coincide con el string elegido")
	} else {
		fmt.Println("No coincide ningun string con el elegido")
	}

}
