package Quicksort

// QuickSort - Ordena um array utilizando o algoritmo QuickSort.
// Arrays vazios ou com um elemento são casos base e já estão ordenados.
// Escolhe-se um pivô normalmente o primeiro elemento de um array, particiona-se o array em elementos menores e maiores que o pivô,
// aplica-se o QuickSort recursivamente nos subarrays e combina-se os resultados.
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	} else {
		pivo := arr[0]
		menores := filtrarMenores(arr, pivo)
		maiores := filtrarMaiores(arr, pivo)

		resultado := append(QuickSort(menores), pivo)
		resultado = append(resultado, QuickSort(maiores)...)
		return resultado

	}
}

// FiltrarMenores - Função auxiliar para filtrar elementos menores ou iguais ao pivô.
// Iteramos pelos elementos do array, exceto o pivô, e adicionamos os menores ou iguais ao pivô em um novo array.
func filtrarMenores(arr []int, pivo int) []int {
	menores := make([]int, 0)
	for _, i := range arr[1:] {
		if i <= pivo {
			menores = append(menores, i)
		}
	}
	return menores
}

// FiltrarMaiores - Função auxiliar para filtrar elementos maiores que o pivô.
// Iteramos pelos elementos do array, exceto o pivô, e adicionamos os maiores que o pivô em um novo array.
func filtrarMaiores(arr []int, pivo int) []int {
	maiores := make([]int, 0)
	for _, i := range arr[1:] {
		if i > pivo {
			maiores = append(maiores, i)
		}
	}
	return maiores

}
