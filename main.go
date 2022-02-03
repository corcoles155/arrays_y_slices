package main

import "fmt"

//Se inicializa con 10 posiciones con el valor 0
var tabla [10]int

//matriz con 5 filas y 7 columnas
var matriz [5][7]

//si colocamos valor entre [] es un array y si no es un slice
var slice := []int{1,2,3}

func main()  {
	fmt.Println(tabla)

	tabla2 := [10]int{1,2,3,4,5,6,7,8,9,10}

	for i:=0; i<len(table); i++{
		fmt.Println(tabla2[i])
	}

	variante2()

	variante3()

	variante4()
}

func variante2()  {
	elementos := [5]int{1,2,3,4,5}
	porcion := elementos[3:] //Desde la tercena posición del array elementos hasta la última posición ([4,5])
	fmt.Println(porcion)
}

func variante3()  {
	elementos := make([int], 5, 20) //inicialmente voy a tener un array de 5 elementos, aunque en memoria tendrá capacidad para 20 elementos
	fmt.Printf("Largo %d, Capacidad %d \n", len(elementos), cap(elementos))
}

func variante4()  {
	nums := make([int], 0, 0)

	for i:=0; i<100; i++ {
		nums = append(nums, i) //Crea un elemento nums con el valor de i
	}

	fmt.Printf("Largo %d, Capacidad %d \n", len(nums), cap(nums)) 
}
