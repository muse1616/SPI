package main

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	for i := 0; i < len(arr); {
		if arr[i] == " " || arr[i] == "" {
			arr = append(arr[:i], arr[i+1:]...)
		} else {
			i++
		}
	}
	// 翻转数组 然后组合即可
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
	return strings.Join(arr, " ")
}
