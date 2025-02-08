package Soma

func Soma(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	return arr[0] + Soma(arr[1:])
}
