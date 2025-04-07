package windowsLog

const (
	LoginEvenType EventNameType = "LoginEvenType"
	RdpEventType  EventNameType = "RdpEventType"
)

var (
	EventToFilePath = map[EventNameType][]string{
		LoginEvenType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
		RdpEventType: {
			"C:\\Windows\\System32\\winevt\\Logs\\Security.evtx",
		},
	}
)

// runBase info key

const (
	MachineUUIDKey = "MachineUUID"
	EventIdKey     = "EventID"
	EventTimeKey   = "EventTime"
)

// login event info key
const (
	LoginTypeKey       = "LoginType"
	SourceIpKey        = "SourceIp"
	UsernameKey        = "Username"
	SubjectUsernameKey = "SubjectUsername"
	SubjectDomainKey   = "SubjectDomain"
	ProcessNameKey     = "ProcessName"
)

// login event value path
// 用于解析GoEvtxMap中的值，需要提供path 去解析map 中的value
const (
	usernamePath        = "/Event/EventData/TargetUserName"
	ipAddressPath       = "/Event/EventData/IpAddress"
	logonTypePath       = "/Event/EventData/LogonType"
	subjectUserNamePath = "/Event/EventData/SubjectUserName"
	subjectDomainPath   = "/Event/EventData/SubjectDomainName"
	processNamePath     = "/Event/EventData/ProcessName"
)

// rdp event info key
const (
	AccountNameKey   = "AccountName"
	AccountDomainKey = "AccountDomain"
	ClientNameKey    = "ClientName"
	ClientAddressKey = "ClientAddress"
)

// rdp event value path
const (
	AccountNamePath   = "/Event/EventData/AccountName"
	AccountDomainPath = "/Event/EventData/AccountDomain"
	ClientNamePath    = "/Event/EventData/ClientName"
	ClientAddressPath = "/Event/EventData/ClientAddress"
)
