package msg

// JsonMsg 提供给websocket 返回信息格式
//
// Code 模拟的是http的状态码
// Msg 额外的信息
// Type 表示 Data 的类型，通过该值可以给前端返回信息，使用不同的类型去展示这个 Data
type JsonMsg struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}
