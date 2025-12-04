package main

import "fmt"
const  (
	
	Lunes = iota +1
	Martes
	Miércoles
	Jueves
	Viernes
	Sábado
	Domingo
)

func main() {
	fmt.Println(Lunes)
	fmt.Println(Martes)
	fmt.Println(Miércoles)
	fmt.Println(Jueves)
	fmt.Println(Viernes)
	fmt.Println(Sábado)
	fmt.Println(Domingo)
}