package utils

func Unique[T comparable](slice []T) []T {
	uniqueSlice := make([]T, 0, len(slice))
	set := make(map[T]struct{}, len(slice))

	for _, item := range slice {
		if _, ok := set[item]; !ok {
			uniqueSlice = append(uniqueSlice, item)
			set[item] = struct{}{}
		}
	}

	return uniqueSlice
}
