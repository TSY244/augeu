package server

type DBFunction func()

var (
	FunctionMap = map[int]DBFunction{}
)
