package arrays

type Array[T comparable] []T

func (v Array[T]) Map(f func(T) T) Array[T] {
	var mappedValues Array[T]
	for _, val := range v {
		mappedValues = append(mappedValues, f(val))
	}
	return mappedValues
}

func (v Array[T]) Filter(f func(T) bool) Array[T] {
	var filteredValues Array[T]
	for _, val := range v {
		if f(val) {
			filteredValues = append(filteredValues, val)
		}
	}

	return filteredValues
}

func (v Array[T]) Includes(val T) bool {
	for _, value := range v {
		if value == val {
			return true
		}
	}

	return false
}
