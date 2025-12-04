package main

import "fmt"

func main() {
	// Calcula el valor de esta expresión
	// Recuerda: Paréntesis primero, luego *, /, %, y finalmente +, -
	
	a := (8 + 3) * (3 % 5) - 2 * (4 + 1)
	
	fmt.Printf("(8 + 3) * (3 %% 5) - 2 * (4 + 1) = %d\n", a)
}
