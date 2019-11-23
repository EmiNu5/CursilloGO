package main

import "fmt"

func main(){
	var i int
	var string1 string
	var palindromo bool

	fmt.Println("Ingrese un string")
	fmt.Scan(&string1)

	for i=0;i<len(string1);i++{
		palindromo = true
		if string1[i] != string1[len(string1)-1-i]{
			palindromo = false
			break
		}
	}
	if palindromo {
		fmt.Println("Es un palindromo")
	} else{
		fmt.Println("No es un palindromo")
	}
}