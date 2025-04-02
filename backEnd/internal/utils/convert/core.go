package convert

type listCopyFunc[T, E any] func(T) E

func ArrayCopy[T, E any](lists []T, copyFunc listCopyFunc[T, E]) []E {
	if len(lists) == 0 {
		return nil
	}
	var res []E
	for _, item := range lists {
		res = append(res, copyFunc(item))
	}
	return res
}
