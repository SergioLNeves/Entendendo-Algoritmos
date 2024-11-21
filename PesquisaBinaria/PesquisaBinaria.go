package PesquisaBinaria

import "fmt"

func PesquisaBinaria(lista []int, item int) (int, error) {
	baixo := 0
	alto := len(lista) - 1

	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]

		if chute == item {
			return meio, nil
		}
		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}

	return item, fmt.Errorf("Item não encontrado")
}
