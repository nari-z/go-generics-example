package sample

type HashSet[T comparable] map[T]struct{}

func New[T comparable]() HashSet[T] {
	v := make(HashSet[T], 0)
	return v
}

func (s HashSet[T]) Insert(x T) {
	s[x] = struct{}{}
}
