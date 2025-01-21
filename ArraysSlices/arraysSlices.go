package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices em GO")

	var array1[5] int
	array1[0] = 1
	array1[1] = 2
	array1[2] = 3
	array1[3] = 4
	array1[4] = 5
	fmt.Println(array1)

	array2 := [5]int{6, 7, 8, 9, 10}
	fmt.Println(array2)

	array3 := [...]int{11, 12, 13, 14, 15, 16}
	fmt.Println(array3)

	fmt.Println("-----------------------------------------------------------------------------------")
	slice := []int{1, 2, 3}
	fmt.Println(slice)
	slice = append(slice, 4)
	fmt.Println(slice)

	slice2 := array3[3:6]
	fmt.Println(slice2)
	array3[3] = 22
	fmt.Println(slice2)

	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println(reflect.TypeOf(array1))
	fmt.Println(reflect.TypeOf(slice))

	//Arrays internos
	slice3 := make([]float32, 10, 11)
	fmt.Println("-----------------------------------------------------------------------------------")
	fmt.Println(slice3)
	fmt.Println(len(slice3))//tamanho
	fmt.Println(cap(slice3))//capacidade

	fmt.Println("-----------------------------------------------------------------------------------")
	slice3 = append(slice3, 1)
	slice3 = append(slice3, 2)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))
}