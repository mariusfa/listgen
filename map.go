package listgen

// Map transforms each element in the list using the provided function.
func Map[T any, U any](l List[T], mapper func(T) U) List[U] {
    var result List[U]
    for _, v := range l {
        result = append(result, mapper(v))
    }
    return result
}

