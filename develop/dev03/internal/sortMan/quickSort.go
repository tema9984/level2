package sortMan

//Алгоритм сортировки
func SortStrings(com forSort) []string {
	return quickSort(com)
}

func SortStringsRevers(com forSort) []string {
	rev := quickSort(com)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	return rev
}
func partition(com forSort, low, high int) ([]string, int) {
	arr := com.getSorted()
	//pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if com.compLower(j, high) {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		} else if com.equalsLower(j, high) {
			if com.comp(j, high) {
				arr[i], arr[j] = arr[j], arr[i]
				i++
			}
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}
func sort(com forSort, low, high int) []string {
	arr := com.getSorted()
	if low < high {
		var p int
		arr, p = partition(com, low, high)
		arr = sort(com, low, p-1)
		arr = sort(com, p+1, high)
	}
	return arr
}

func quickSort(com forSort) []string {
	return sort(com, 0, com.getLen()-1)
}
