package OrdenacaoPorSelecao

func BuscaMenorIndice(arr []int) int {

	menor := arr[0]
	indice := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < menor {
			menor = arr[i]
			indice = i
		}
	}
	return indice
}

func OrdenacaoPorSelecao(arr []int) []int {
	var listaOrdenada []int
	var listaDesordenada = arr

	for i := 0; i < len(arr); i++ {
		indice := BuscaMenorIndice(listaDesordenada)
		listaOrdenada = append(listaOrdenada, arr[indice])
		listaDesordenada = append(listaDesordenada[:indice], listaDesordenada[indice+1:]...)
	}

	return listaOrdenada
}
