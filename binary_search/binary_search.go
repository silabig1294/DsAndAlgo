package binarysearch



func BinarySearch(bigstack []int,needle int) bool {

	low := 0
	high := len(bigstack) - 1

	for low <= high{
		median := (low + high) / 2

		if bigstack[median] < needle {
			low = median + 1
		}else{
			high = median - 1
		}
	}

	if low == len(bigstack) || bigstack[low] != needle {
		return false
	}

	return true
}