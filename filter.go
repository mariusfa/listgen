package listgen


// Filter filters the list elements based on a predicate function.
// It returns a new List containing only the elements that satisfy the predicate.
func (l List[T]) Filter(predicate func(T) bool) List[T] {
	var result List[T]
	for _, v := range l {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}
