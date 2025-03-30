package convert

import (
	"augeu/client/internal/pkg/msg"
	"augeu/client/pkg/windowsWmi"
)

type listCopyFunc[T, E any] func(T) E

func ArrayCopy[T, E any](lists []T, copyFunc listCopyFunc[T, E]) []E {
	if len(lists) == 0 {
		return nil
	}
	var res []E
	for _, item := range lists {
		res = append(res, copyFunc(item))
	}
	return res
}

// WmiPatchToMsgPatch 将 wmi 的 patch 转变成 msg patch
func WmiPatchToMsgPatch(wmiPatch windowsWmi.Win32_QuickFixEngineering) msg.Patch {
	return msg.Patch{
		Description: wmiPatch.Description,
		HotFixID:    wmiPatch.HotFixID,
		InstalledBy: wmiPatch.InstalledBy,
		InstalledOn: wmiPatch.InstalledOn,
	}
}
