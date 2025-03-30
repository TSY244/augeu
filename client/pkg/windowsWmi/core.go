package windowsWmi

import (
	"errors"
	"fmt"
	"github.com/StackExchange/wmi"
)

// QueryUuid 查询系统UUID
func QueryUuid() (string, error) {
	var dst []Win32_ComputerSystemProduct
	if err := wmi.Query(QueryUuidKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].UUID, nil
}

// QueryOsName 查询系统名称
func QueryOsName() (string, error) {
	var dst []Win32_OperatingSystem
	if err := wmi.Query(QueryOsNameKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].Caption, nil
}

// QueryOsVersion 查询系统版本
func QueryOsVersion() (string, error) {
	var dst []Win32_OperatingSystemVersion
	if err := wmi.Query(QueryOsVersionKey, &dst); err != nil {
		return "", fmt.Errorf("WMI query failed: %w", err)
	}
	if len(dst) == 0 {
		return "", errors.New("empty result set")
	}
	return dst[0].Version, nil
}

// 统一使用基础查询函数
func baseSlice[T any](resultSet *[]T, queryKey string) error {
	if resultSet == nil {
		return errors.New("nil pointer passed for result storage")
	}
	if err := wmi.Query(queryKey, resultSet); err != nil {
		return fmt.Errorf("WMI query failed: %w", err)
	}
	if len(*resultSet) == 0 {
		return errors.New("empty result set from WMI query")
	}
	return nil
}

// QueryHotFix 查询系统补丁
func QueryHotFix() ([]Win32_QuickFixEngineering, error) {
	var dst []Win32_QuickFixEngineering
	if err := baseSlice(&dst, QueryHotFixKey); err != nil {
		return nil, err
	}
	return dst, nil
}
