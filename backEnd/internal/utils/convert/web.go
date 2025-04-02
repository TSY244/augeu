package convert

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/public/pkg/swaggerCore/models"
)

//
//type GetClientIDRequest struct {
//
//	// IP 地址列表
//	// Required: true
//	IP []string `json:"ip"`
//
//	// 密钥
//	// Required: true
//	Secret *string `json:"secret"`
//
//	// 系统信息
//	// Required: true
//	SystemInfo *SystemInfo `json:"system_info"`
//
//	// 唯一标识符
//	// Required: true
//	UUID *string `json:"uuid"`
//}

//
//type Account struct {
//	CreatedAt time.Time      `gorm:"autoCreateTime"`
//	DeletedAt gorm.DeletedAt `gorm:"index"`
//	// 使用中间表 account_ips 建立与 IPAddress 的多对多关系
//	IPAddresses []IPAddress `gorm:"many2many:account_ips;"`
//	ClientID    string      `gorm:"type:varchar(255);not null"` // 动态的，用于绑定任务
//	// 通过 UUID 外键关联 System
//	System System `gorm:"foreignKey:UUID"`
//	UUID   string `gorm:"type:varchar(36);not null;primaryKey"` // UUID 通常为 36 位
//}

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
