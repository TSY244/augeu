package systeminfo

import (
	"augeu/agent/internal/utils/convert"
	"augeu/agent/pkg/windowsWmi"
	"augeu/public/pkg/swaggerCore/models"
	convert2 "augeu/public/util/convert"
	"runtime"
)

func GetSystemInfo() (*models.SystemInfo, error) {
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
	return &models.SystemInfo{
		OsName:    convert2.StrPtr(osName),
		OsVersion: convert2.StrPtr(osVersion),
		OsArch:    convert2.StrPtr(runtime.GOARCH),
		Patchs:    msgHotFixs,
	}, nil
}

func GetUuid() (string, error) {
	return windowsWmi.QueryUuid()
}
