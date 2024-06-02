package utils

func RemoveDuplicatesList(intSlice []string) []string {
	result := make([]string, 0, len(intSlice))
	for i := range intSlice {
		if i == 0 || intSlice[i] != intSlice[i-1] {
			result = append(result, intSlice[i])
		}
	}
	return result
}
