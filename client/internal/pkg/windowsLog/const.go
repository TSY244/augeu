package windowsLog

const (
	LoginFilePath = "LoginFilePath"
)

var (
	//FilePath = []string{
	//	"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
	//}
	EventToFilePath = map[string][]string{
		LoginFilePath: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		"RdpFilePath": {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
	}
)
