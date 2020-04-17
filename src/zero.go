package src

func CheckZero(n int, array []int) string {
	sum := check(array)
	if sum%2 != 0 {
		return "NO"
	}
	handledArray := handleArray(array)
	sum = check(handledArray)
	if sum > 0 {
		return "NO"
	}
	return "YES"
}

func check(array []int) int {
	sum := 0
	for _, value := range array {
		sum += value
	}
	return sum
}

func handleArray(array []int) []int {
	max := array[0]
	indexMax1 := 0
	indexMax2 := 1
	for index, value := range array {
		if value >= max && len(array) > 2 {
			max = value
			indexMax2 = indexMax1
			indexMax1 = index
		}
	}
	if array[indexMax1] == 0 && array[indexMax2] > 0 || array[indexMax2] == 0 && array[indexMax1] > 0 {
		return array
	}
	if array[indexMax1] == 0 && array[indexMax2] == 0 {
		return array
	}
	for array[indexMax1] > 0 && array[indexMax2] > 0 {
		array[indexMax1] -= 1
		array[indexMax2] -= 1
	}
	if array[indexMax1] == 0 || array[indexMax2] == 0 {
		handleArray(array)
	}
	return array
}
