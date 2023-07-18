package main

import "fmt"

func main() {
	// switch
	// 10 - 8/9 - 6/7 - 5

	var calificacion int

	fmt.Print("Ingresa una calificaion: ")
	fmt.Scanf("%d", &calificacion)

	switch calificacion {
	case 10:
		fmt.Println("Felicidades, tu calificaicon es perfecta")
	case 8, 9:
		fmt.Println("Aprobaste la materia.")
	case 7, 6:
		fmt.Println("Aprobaste pero tienes que estudiar")
	case 1, 2, 3, 4, 5:
		fmt.Println("No aprobaste")
	default:
		fmt.Println("Calificacion no valida")
	}

	//switch {
	//case calificacion >= 10:
	//	fmt.Println("Felicidades, tu calificaicon es perfecta")
	//case calificacion == 8 || calificacion == 9:
	//	fmt.Println("Aprobaste la materia.")
	//case calificacion == 6 || calificacion == 7:
	//	fmt.Println("Aprobaste pero tienes que estudiar")
	//case calificacion >= 0 && calificacion <= 5:
	//	fmt.Println("No aprobaste")
	//default:
	//	fmt.Println("Calificacion no valida")
	//}
}
