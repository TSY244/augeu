package utils

func CheckUuid[T any](events []*T, f func(*T) bool) []*T {
	var res []*T
	for _, event := range events {
		if event == nil {
			continue
		}
		if f(event) {
			res = append(res, event)
		}
	}
	return res
}
