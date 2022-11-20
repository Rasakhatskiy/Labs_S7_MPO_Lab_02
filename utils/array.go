package utils

func RemoveIndex[T any](s []T, index int) []T {
	return append(s[:index], s[index+1:]...)
}

func Contains[T comparable](array []T, elem T) bool {
	for _, el := range array {
		if elem == el {
			return true
		}
	}
	return false
}

type ItemNotFoundError struct {
	Err error
}

func (err ItemNotFoundError) Error() string {
	return "item not found"
}

func Find[T comparable](array []T, elem T) (int, error) {
	for i, el := range array {
		if elem == el {
			return i, nil
		}
	}
	return -1, ItemNotFoundError{}
}
