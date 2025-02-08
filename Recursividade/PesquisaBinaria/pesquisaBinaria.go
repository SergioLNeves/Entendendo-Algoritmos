package PesquisaBinaria

import "fmt"

func PesquisaBinaria(baixo, alto, valor int) (int, error) {
	meio := (baixo + alto) / 2
	chute := meio

	if valor > alto || valor < baixo {
		return valor, fmt.Errorf("valor out of range: %d", valor)
	}

	if chute < valor {
		baixo = meio + 1
		return PesquisaBinaria(baixo, alto, valor)
	} else if chute > valor {
		alto = meio - 1
		return PesquisaBinaria(baixo, alto, valor)
	} else if chute == valor {
		return chute, nil
	} else {
		return 0, nil
	}
}
