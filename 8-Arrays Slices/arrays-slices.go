package main

import "fmt"

func main() {

	//Arrays
	var array1 [5]string
	array1[0] = "Posicao 1"
	fmt.Println(array1)

	array2 := [5]string{"Camila", "Ana", "Joao", "Andre", "Paulo"}
	fmt.Println(array2)

	array3 := [...]int8{2, 4, 6, 9, 18, 1, 1}
	fmt.Println(array3)

	//Slice
	slice := []int{103, 102, 20, 32, 399, 69}
	fmt.Println(slice)

	slice = append(slice, 1000)
	fmt.Println(slice)

	slice2 := array2[0:4]
	fmt.Println(slice2)

	//Arrays Internos
	slice3 := make([]float32, 10, 15)

	slice3 = append(slice3, 5)
	slice3 = append(slice3, 6)
	fmt.Println(slice3)

	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}
