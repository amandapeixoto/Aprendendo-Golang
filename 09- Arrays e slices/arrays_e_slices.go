package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e slices")

	var array1 [5]int
	array1[0] = 10
	array1[1] = 8
	fmt.Println(array1)

	array2 := [5]string{"posição 1 -", "posição 2 -", "posição 3 -", "posição 4 -", "posição 5 -"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{10, 15, 20, 25, 30, 35, 40} //possui o tamanho dinâmico
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(array3)) //TypeOf tipos das variáveis
	fmt.Println(reflect.TypeOf(slice))

	slice = append(slice, 550) //acrescentando um novo valor ao slice
	fmt.Println(slice)

	slice2 := array2[1:3] //pegar do indíce 1(que é inclusivo) e o 2 (pq o indíce 3 é exclusivo, então entra o que vem  antes)sem alterar o array
	fmt.Println(slice2)

	array2[1] = "posição 2 alterada -"
	fmt.Println(array2)
	fmt.Println(slice2)

	// ARRAYS INTERNOS (A FUNÇÃO MAKE CRIOU UM ARRAY DE 15 POSIÇÕES E RETORNOU UM SLICE(FATIA DO ARRAY) DE 10 POSIÇÕES)
	fmt.Println("____________") // O slice não tem limite, sempre que você ultrapassar a capacidade ele duplica.
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //lenght (Tamanho do slice)
	fmt.Println(cap(slice3)) //capacidade

	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //lenght (Tamanho do slice)
	fmt.Println(cap(slice3)) //capacidade

	slice4 := make([]float32, 5) //caso não informe a capacidade, entende-se que o tamanho e capacidade são os mesmos
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 9)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))
}
