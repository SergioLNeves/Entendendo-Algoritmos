package Recursividade

func PesquisaBinaria(baixo, alto, valor int) int {
	meio := (baixo + alto) / 2
	chute := meio

	if chute < valor {
		baixo = meio + 1
		return PesquisaBinaria(baixo, alto, valor)
	} else if chute > valor {
		alto = meio - 1
		return PesquisaBinaria(baixo, alto, valor)
	} else if chute == valor {
		return chute
	} else {
		return 0
	}
}
