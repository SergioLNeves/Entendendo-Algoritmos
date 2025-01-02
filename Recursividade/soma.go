package Recursividade

func Soma(arr []int) int {
	var soma int

	if len(arr) == 0 {
		return 0
	}

	for _, valor := range arr[1:] {
		soma += valor
	}

	return arr[0] + soma
}

func ProcuraBase(arr []int) int {

}
