package listgen

func (l *List[T]) Add(element T) {
    *l = append(*l, element)
}