package OrdenacaoPorSelecao

// OrdenacaoPorSelecao - Ordena um array utilizando o algoritmo de ordenação por seleção.
// Em cada iteração, encontra o menor elemento do array desordenado e o move para o array ordenado.
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

// BuscaMenorIndice - Encontra o índice do menor elemento em um array.
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
