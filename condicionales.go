package main

import "fmt"

func main() {

	//	var edad int
	//	fmt.Print("Ingresa tu edad: ")
	//	fmt.Scanf("%d", &edad)
	//
	//	if edad >= 18 {
	//		fmt.Println("El usuario es mayor de edad")
	//	} else {
	//		fmt.Println("El usuario es menor de edad")
	//	}

	// else if
	// 10 - 8/9 - 6/7 - 5

	//var calificacion int
	//
	//fmt.Print("Ingresa una calificacion: ")
	//fmt.Scanf("%d", &calificacion)
	//
	//if calificacion >= 10 {
	//	fmt.Println("Felicidades, tu calificaicon es perfecta")
	//} else if calificacion == 8 || calificacion == 9 {
	//	fmt.Println("Aprobaste la materia.")
	//} else if calificacion == 6 || calificacion == 7 {
	//	fmt.Println("Aprobaste la materia, pero necesitas estudiar mas.")
	//} else if calificacion >= 0 && calificacion <= 5 {
	//	fmt.Println("No aprobaste")
	//} else {
	//	fmt.Println("Calificacion no valida")
	//}

	//if calificacion >= 10 {
	//	fmt.Println("Felicidades, tu calificaicon es perfecta")
	//} else {
	//	if calificacion == 8 || calificacion == 9 {
	//		fmt.Println("Aprobaste la materia.")
	//	} else {
	//		if calificacion == 6 || calificacion == 7 {
	//			fmt.Println("Aprobaste la materia, pero necesitas estudiar mas.")
	//		} else {
	//			if calificacion >= 0 && calificacion <= 5 {
	//				fmt.Println("No aprobaste")
	//			}
	//		}
	//	}
	//}

	if nombre, edad := "Cody", 7; nombre == "Cody" {
		fmt.Println("Hola", nombre, "Te damos la bienvenida!")
	} else {
		fmt.Println("Los valores son:", nombre, edad)
	}

}
