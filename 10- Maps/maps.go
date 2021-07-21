package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Amanda",
		"sobrenome": "Peixoto",
	}
	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Romão",
			"último":   "Silva",
		},
		"curso": {
			"nome":   "Sistemas",
			"campus": "Campus 1",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "nome") // Para deletar chaves
	fmt.Println(usuario2)

	usuario2["signo"] = map[string]string{
		"nome": "Peixes",
	}
	fmt.Println(usuario2)

}
