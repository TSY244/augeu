package windowsLog

const (
	LoginEvenType    eventNameType = "LoginEvenType"
	RdpFileEventType eventNameType = "RdpFileEventType"
)

var (
	EventToFilePath = map[eventNameType][]string{
		LoginEvenType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		RdpFileEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
	}
)

// base info key

const (
	EventIdKey = "EventID"
)
