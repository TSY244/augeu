package convert

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/backEnd/internal/pkg/DBMnager/Log"
	"augeu/public/pkg/swaggerCore/models"
	"time"
)

func GetClientIDRequest2Db(getClientIDRequest models.GetClientIDRequest) HostInfo.Account {
	return HostInfo.Account{
		ClientID:    "",
		IPAddresses: GetClientIDRequestIps2Db(getClientIDRequest),
		System: HostInfo.System{
			OSArch:    *getClientIDRequest.SystemInfo.OsArch,
			OSName:    *getClientIDRequest.SystemInfo.OsName,
			OSVersion: *getClientIDRequest.SystemInfo.OsVersion,
			Patches:   GetClientIDRequestPatches2Db(getClientIDRequest),
		},
		UUID: *getClientIDRequest.UUID,
	}
}

func GetClientIDRequestIps2Db(getClientIDRequest models.GetClientIDRequest) []HostInfo.IPAddress {
	return ArrayCopy(getClientIDRequest.IP, func(originStr string) HostInfo.IPAddress {
		return HostInfo.IPAddress{
			Value: originStr,
		}
	})
}

func GetClientIDRequestPatches2Db(getClientIDRequest models.GetClientIDRequest) []HostInfo.Patch {
	return ArrayCopy(getClientIDRequest.SystemInfo.Patchs, func(originPatch *models.Patch) HostInfo.Patch {
		return HostInfo.Patch{
			Description: *originPatch.Description,
			HotFixID:    *originPatch.HotFixID,
			InstalledBy: *originPatch.InstalledBy,
			InstalledOn: *originPatch.InstalledOn,
		}
	})
}

//
//type LoginEvent struct {
//	EventId         int64          `json:"event_id"`
//	EventTime       time.Time      `json:"event_time"` // 这个事件创建的时间
//	LoginType       string         `json:"login_type"`
//	SourceIp        string         `json:"source_ip"`
//	Username        string         `json:"username"`
//	SubjectUsername string         `json:"subject_username"` // 主体用户名
//	SubjectDomain   string         `json:"subject_domain"`   // 主体域
//	ProcessName     string         `json:"process_name"`     // 进程名称
//	CreateAt        time.Time      `gorm:"autoCreateTime"`   // 这条sql 记录创建的时间
//	DeleteAt        gorm.DeletedAt `gorm:"index"`
//	ID              uint           `gorm:"primaryKey;autoIncrement"`      // 主键
//	UUID            string         `gorm:"type:varchar(255);column:uuid"` // windows 主机的uuid
//}
//type LoginEvent struct {
//
//	// 事件ID，4624表示成功登录
//	// Required: true
//	EventID *int64 `json:"EventID"`
//
//	// 事件时间，格式：2006-01-02 15:04:05（注意原数据日期时间连写问题）
//	// Required: true
//	// Format: date-time
//	EventTime *strfmt.DateTime `json:"EventTime"`
//
//	// 登录类型（Unknown表示无法识别的类型）
//	// Required: true
//	LoginType *string `json:"LoginType"`
//
//	// 设备唯一标识符（UUID格式）
//	// Required: true
//	MachineUUID *string `json:"MachineUUID"`
//
//	// 触发登录的进程名（-表示无）
//	// Required: true
//	ProcessName *string `json:"ProcessName"`
//
//	// 源IP地址（-表示无）
//	// Required: true
//	SourceIP *string `json:"SourceIP"`
//
//	// 登录用户所属域（-表示无）
//	// Required: true
//	SubjectDomain *string `json:"SubjectDomain"`
//
//	// 登录会话用户（-表示无）
//	// Required: true
//	SubjectUser *string `json:"SubjectUser"`
//
//	// 登录用户名（SYSTEM表示系统账户）
//	// Required: true
//	Username *string `json:"Username"`
//}

func LoginEvent2Db(loginEvent *models.LoginEvent) Log.LoginEvent {
	return Log.LoginEvent{
		EventId:         *loginEvent.EventID,
		EventTime:       time.Time(*loginEvent.EventTime),
		LoginType:       *loginEvent.LoginType,
		SourceIp:        *loginEvent.SourceIP,
		Username:        *loginEvent.Username,
		SubjectUsername: *loginEvent.SubjectUser,
		SubjectDomain:   *loginEvent.SubjectDomain,
		ProcessName:     *loginEvent.ProcessName,
	}
}
