package windowsLog

var (
	LoginEvent = map[int64]string{
		4624: "成功登陆",
		4625: "登陆失败",
	}
	RdpEvent = map[int64]string{
		4778: "成功登陆",
		4777: "登陆失败",
	}
)
