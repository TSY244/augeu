package systeminfo

import (
	"augeu/agent/internal/pkg/msg"
	"augeu/agent/internal/utils/convert"
	"augeu/agent/pkg/windowsWmi"
	"os"
)

//type nil struct {
//	Description string
//	HotFixID    string
//	InstalledBy string
//	InstalledOn string
//}

func GetSystemInfo() (*msg.SystemInfo, error) {
	osName, err := windowsWmi.QueryOsName()
	if err != nil {
		return nil, err
	}
	osVersion, err := windowsWmi.QueryOsVersion()
	if err != nil {
		return nil, err
	}
	hotFix, err := windowsWmi.QueryHotFix()
	if err != nil {
		return nil, err
	}
	msgHotFixs := convert.ArrayCopy(hotFix, convert.WmiPatchToMsgPatch)
	return &msg.SystemInfo{
		OSName:    osName,
		OSVersion: osVersion,
		OSArch:    os.Getenv("PROCESSOR_ARCHITECTURE"),
		Patchs:    msgHotFixs,
	}, nil
}

func GetUuid() (string, error) {
	return windowsWmi.QueryUuid()
}
