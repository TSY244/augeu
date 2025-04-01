package msg

// JsonMsgResp 提供给websocket 返回信息格式
//
// Code 模拟的是http的状态码
// Msg 额外的信息
// Type 表示 Data 的类型，通过该值可以给前端返回信息，使用不同的类型去展示这个 Data
type JsonMsgResp struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

type JsonMsgReq struct {
	UUID    string `json:"uuid"`
	CmdType string `json:"cmd_type"`
	Cmd     string `json:"cmd"`
}

type JsonMsg struct {
	Type     string `json:"type"`
	ClientId string `json:"clientId"`
	Message  string `json:"message"`
}

type ClintConnectMsg struct {
	Type     string `json:"type"`
	ClientId string `json:"clientId"`
}

type WelcomeMsg struct {
	Type     string `json:"type"`
	ClientId string `json:"clientId"`
	Message  string `json:"message"`
}

type HelloMsg struct {
	UUID       string     `json:"uuid"`
	IP         *[]string  `json:"ip"`
	SystemInfo SystemInfo `json:"system_info"`
}
type SystemInfo struct {
	OSName    string  `json:"os_name"`
	OSVersion string  `json:"os_version"`
	OSArch    string  `json:"os_arch"`
	Patchs    []Patch `json:"patchs"`
}

type Patch struct {
	Description string
	HotFixID    string
	InstalledBy string
	InstalledOn string
}
