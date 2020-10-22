package util

func MapToInt(input map[string]int32) map[string]int {
	output := make(map[string]int)

	for k, v := range input {
		output[k] = int(v)
	}

	return output
}

func MapToInt32(input map[string]int) map[string]int32 {
	output := make(map[string]int32)

	for k, v := range input {
		output[k] = int32(v)
	}

	return output
}
