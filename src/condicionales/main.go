package main

import (
	"fmt"
)

func main() {
	// valor:= rand.Intn(100)

	// fmt.Println("El valor es: ", valor)

	// if valor % 2 == 0 {
	// 	fmt.Println("El valor es par")
	// }else {
	// 	fmt.Println("El valor es impar", valor)
	// }

	// fmt.Println("Regresa pronto!")

	// if  valor:= rand.Intn(100); valor % 2 == 0 {
	// 	fmt.Println("El valor es par")
	// }else {
	// 	fmt.Println("El valor es impar", valor)
	// }

	// fmt.Println("El arch es: ")
	// arch := runtime.GOARCH

	// switch arch {
	// case "amd64":
	// 	fmt.Println("El arch es amd64")
	// case "arm64":
	// 	fmt.Println("El arch es arm64")
	// default:
	// 	fmt.Println("El arch es desconocido")
	// }

	// switch os := runtime.GOOS; os {
	// case "darwin":
	// 	fmt.Println("El os es darwin")
	// case "linux":
	// 	fmt.Println("El os es linux")
	// default:
	// 	fmt.Println("El os es desconocido")
	// }

	// letra := "A"
	// switch letra {
	// case "A":
	// 	fmt.Println("La letra es A")
	// 	fallthrough
	// case "a":
	// 	fmt.Println("La letra es a")
	// default:
	// 	fmt.Println("La letra es desconocida")

	// }

	letras := 'A'
	switch letras {
		case 'A','E','I','O','U':
			fmt.Println("La letra es una vocal")
		default:
			fmt.Println("La letra es una consonante")
	}
}
