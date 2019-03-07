package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("********MI PROGRAMA EN GO**********")
	fmt.Println("Hola " + os.Args[1] + " Bienvenido a go")
	var edad, _ = strconv.Atoi(os.Args[2])

	if edad >= 18 && edad != 25 && edad != 99 {
		fmt.Println("Eres mayor de edad o tienes 17")
	} else if edad == 25 {
		fmt.Println("tienes 25 años")
	} else if edad == 99 {
		fmt.Println("Estas a punto de morir")
	} else {
		fmt.Println("Eres menor de edad")
	}
}
