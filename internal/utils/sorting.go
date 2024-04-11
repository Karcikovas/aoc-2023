package utils

func BinarySearch(successValue int, options []int) bool {
	low := 0
	high := len(options) - 1

	for low <= high {
		middleValueIndex := low + (high-low)/2
		middleValue := options[middleValueIndex]

		if successValue > middleValue {
			low = middleValueIndex + 1
		} else if successValue < middleValue {
			high = middleValueIndex - 1
		} else {
			return true
		}
	}

	if low == len(options) || low != successValue {
		return false
	}

	return true
}
