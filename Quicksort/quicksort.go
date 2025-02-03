package Quicksort

func filtrarMenores(arr []int, pivo int) []int {
	menores := make([]int, 0)
	for _, i := range arr[1:] {
		if i <= pivo {
			menores = append(menores, i)
		}
	}
	return menores
}

func filtrarMaiores(arr []int, pivo int) []int {
	maiores := make([]int, 0)
	for _, i := range arr[1:] {
		if i > pivo {
			maiores = append(maiores, i)
		}
	}
	return maiores

}

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
