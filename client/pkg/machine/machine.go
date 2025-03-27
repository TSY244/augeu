package machine

import (
	_const "augeu/client/internal/utils/const"
	"golang.org/x/sys/windows/registry"
)

var (
	WindowsId = "" // 提高效率，只获取一次
)

// GetWindowsGuid 通过注册表获取windows 机器的guid
//
// 注意：
//   - 通过注册表获取的guid 只有当用户重新安装了系统，或者是修改了windows 的安装位置才会发生变化
func GetWindowsGuid() (string, error) {
	if WindowsId != "" {
		return WindowsId, nil
	}
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, _const.GuidKeyPath, registry.QUERY_VALUE)
	if err != nil {
		return "", err
	}
	defer key.Close()
	value, _, err := key.GetStringValue(_const.GuidKeyName)
	if err != nil {
		return "", err
	}
	WindowsId = value
	return value, nil
}
