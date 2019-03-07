package main

import "fmt"

type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	// var gorra_negra = Gorra{"Nike", "Negro", 25.50, true}
	// var numero1 float32 = 10
	// var numero2 float32 = 6

	// fmt.Print("gorras: ")
	// fmt.Println(gorra_negra.marca)

	// fmt.Print("la suma es: ")
	// fmt.Println(operacion(numero1, numero2, "+"))

	// fmt.Print("la resta es: ")
	// fmt.Println(operacion(numero1, numero2, "-"))

	// fmt.Print("la multiplicacion es: ")
	// fmt.Println(operacion(numero1, numero2, "*"))

	// fmt.Print("la division es: ")
	// fmt.Println(operacion(numero1, numero2, "/"))

	// holaMundo()
	// fmt.Println(gorras(45, "$"))
	// fmt.Println(devolverTexto())
	// pantalon("rolo", "largo", "mezclilla", "de vestir")
	// var peliculas [3][2]string
	// peliculas[0][0] = "la verdad duele"
	// peliculas[0][1] = "mientras duermes"
	// peliculas[1][0] = "El señor de los anillos"
	// peliculas[1][1] = "El señor de los anillos 2"
	// peliculas[2][0] = "Gran Torino"
	// peliculas[2][1] = "A todo gas"

	peliculas := []string{
		"La verdad duele",
		"Ciudadano ejemplar",
		"Batman",
	}
	peliculas = append(peliculas, "sin limites")

	fmt.Println(len(peliculas))

}
func pantalon(caracteristicas ...string) {
	for _, caracteristica := range caracteristicas {
		fmt.Println(caracteristica)
	}
}

func gorras(pedido float32, moneda string) (string, string, float32) {
	precio := func() float32 {
		return pedido * 7
	}
	return "el precio de gorras pedidas es: ", moneda, precio()
}
func devolverTexto() string {
	return "Texto por defecto"
}

func holaMundo() {
	fmt.Println("Hola mundo!!")
}

func operacion(n1 float32, n2 float32, op string) float32 {
	var resultado float32
	if op == "+" {
		resultado = n1 + n2
	}
	if op == "-" {
		resultado = n1 - n2
	}
	if op == "*" {
		resultado = n1 * n2
	}
	if op == "/" {
		resultado = n1 / n2
	}
	return resultado
}
