package main

// Reverse input string
func Reverse(s string) string {
	arr := []byte(s)
	mid := len(arr) / 2

	for i, j := 0, len(arr)-1; i < mid; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}
