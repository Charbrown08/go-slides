package main

import "fmt"

func main() {
	// for {
	// 	fmt.Println(" Quieres continuar? (s/n): ")
	// 	var c rune
	// 	fmt.Scanf("%c\n", &c)
	// 	if c == 'n' || c == 'N' {
	// 		break
	// 	} else {
	// 		fmt.Println("Bienvenido al programa...")
	// 		continue
	// 	}
	// }
	// fmt.Println("Hasta luego!")

	// var c  rune
	// for c != 's' && c != 'S' {
	// 	fmt.Println(" Quieres continuar? (s/n): ")
	// 	fmt.Scanf("%c\n", &c)
	// }

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// }

	// contexto y ocultacion de variables

	// a := 3
	// if a == 3 {
	// 	i := 1
	// 	fmt.Println(" Aqui podemos Imprimir la variable i: ", i)
	// }
	// fmt.Println("Aqui solo podemor imprimir la variable a: ", a)

	a := 0 // Variable 'a' en el scope de main
	b := 0 // Variable 'b' en el scope de main

	if true {
		a := 1 // ⚠️ SHADOWING: Crea una NUEVA variable 'a' local al bloque if
		// Esta 'a' oculta la variable 'a' externa
		b = 1 // ✅ ASIGNACIÓN: Modifica la variable 'b' del scope externo
		a++   // Incrementa la 'a' LOCAL (la del if)
		b++   // Incrementa la 'b' EXTERNA (la de main)

	}

	fmt.Println("a: ", a) // Imprime 0 (la variable externa no cambió)
	fmt.Println("b: ", b) // Imprime 2 (se asignó 1 y luego se incrementó)
}
