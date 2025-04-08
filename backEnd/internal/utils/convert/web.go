package convert

import (
	"augeu/backEnd/internal/pkg/DBMnager/HostInfo"
	"augeu/backEnd/internal/pkg/DBMnager/Log"
	"augeu/public/pkg/logger"
	"augeu/public/pkg/swaggerCore/models"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"time"
)

func GetClientIDRequest2Db(getClientIDRequest models.GetClientIDRequest) HostInfo.Account {
	return HostInfo.Account{
		ClientID:    "",
		IPAddresses: GetClientIDRequestIps2Db(getClientIDRequest),
		System: HostInfo.System{
			OSArch:    *getClientIDRequest.ClientInfo.SystemInfo.OsArch,
			OSName:    *getClientIDRequest.ClientInfo.SystemInfo.OsName,
			OSVersion: *getClientIDRequest.ClientInfo.SystemInfo.OsVersion,
			Patches:   GetClientIDRequestPatches2Db(getClientIDRequest),
		},
		UUID: *getClientIDRequest.ClientInfo.UUID,
	}
}

func GetClientIDRequestIps2Db(getClientIDRequest models.GetClientIDRequest) []HostInfo.IPAddress {
	return ArrayCopy(getClientIDRequest.ClientInfo.IP, func(originStr string) HostInfo.IPAddress {
		return HostInfo.IPAddress{
			Value: originStr,
		}
	})
}

func GetClientIDRequestPatches2Db(getClientIDRequest models.GetClientIDRequest) []HostInfo.Patch {
	return ArrayCopy(getClientIDRequest.ClientInfo.SystemInfo.Patchs, func(originPatch *models.Patch) HostInfo.Patch {
		return HostInfo.Patch{
			Description: *originPatch.Description,
			HotFixID:    *originPatch.HotFixID,
			InstalledBy: *originPatch.InstalledBy,
			InstalledOn: *originPatch.InstalledOn,
		}
	})
}

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

func DbHostinfo2moduleHostinfo(hostInfo HostInfo.Account) *models.ClientInfo {
	return &models.ClientInfo{
		SystemInfo: DbSystem2moduleSystemInfo(hostInfo.System),
		IP:         DbIpaddress2StrSlice(hostInfo.IPAddresses),
		UUID:       &hostInfo.UUID,
	}
}

func DbSystem2moduleSystemInfo(system HostInfo.System) *models.SystemInfo {
	return &models.SystemInfo{
		OsArch:    &system.OSArch,
		OsName:    &system.OSName,
		OsVersion: &system.OSVersion,
		Patchs:    ArrayCopy(system.Patches, DbPatch2modulePatch),
	}
}

func DbPatch2modulePatch(patch HostInfo.Patch) *models.Patch {
	return &models.Patch{
		Description: &patch.Description,
		HotFixID:    &patch.HotFixID,
		InstalledBy: &patch.InstalledBy,
		InstalledOn: &patch.InstalledOn,
	}
}

func DbIpaddress2StrSlice(ipAddresses []HostInfo.IPAddress) []string {
	var ipAddressesStrSlice []string
	for _, ipAddress := range ipAddresses {
		ipAddressesStrSlice = append(ipAddressesStrSlice, ipAddress.Value)
	}
	return ipAddressesStrSlice
}

func DbLoginEvent2modelLogEvent(loginEvent Log.LoginEvent) *models.LoginEvent {
	// time.time -> strfmt.DateTime
	evnetTime := strfmt.DateTime(loginEvent.EventTime)
	return &models.LoginEvent{
		EventID:       swag.Int64(loginEvent.EventId),
		EventTime:     &evnetTime,
		LoginType:     swag.String(loginEvent.LoginType),
		MachineUUID:   swag.String(loginEvent.UUID),
		ProcessName:   swag.String(loginEvent.ProcessName),
		SourceIP:      swag.String(loginEvent.SourceIp),
		SubjectDomain: swag.String(loginEvent.SubjectDomain),
	}
}

func ModelRdpEvent2DbRdpEvent(rdpEvent *models.RDPEventUnit) Log.EventRDPLogon {
	if rdpEvent == nil {
		logger.Errorf("rdpEvent is nil")
		return Log.EventRDPLogon{}
	}
	if rdpEvent.Base == nil {
		logger.Errorf("rdpEvent.Base is nil")
		return Log.EventRDPLogon{}
	}
	if rdpEvent.AccountName == nil || rdpEvent.AccountDomain == nil ||
		rdpEvent.ClientAddress == nil || rdpEvent.ClientName == nil ||
		rdpEvent.Base.EventID == nil || rdpEvent.Base.UUID == nil {
		logger.Errorf("rdpEvent is invalid")
		return Log.EventRDPLogon{}
	}
	return Log.EventRDPLogon{
		AccountDomain: *rdpEvent.AccountDomain,
		AccountName:   *rdpEvent.AccountName,
		ClientAddress: *rdpEvent.ClientAddress,
		ClientName:    *rdpEvent.ClientName,
		EventID:       *rdpEvent.Base.EventID,
		UUID:          *rdpEvent.Base.UUID,
	}
}
