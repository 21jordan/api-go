package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("lector")
	var nuevoTexto = []byte(os.Args[1])
	// ioutil.WriteFile("fichero.txt", nuevoTexto, 0777)

	fmt.Println(nuevoTexto)
	var fichero, errorDeFichero = ioutil.ReadFile("fichero.txt")
	showError(errorDeFichero)
	fmt.Println(string(fichero))
}

func showError(e error) {
	if e != nil {
		panic(e)
	}
}
