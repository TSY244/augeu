package windowsWmi

import "testing"

//const (
//	QueryUuidKey      = "SELECT UUID FROM Win32_ComputerSystemProduct"
//	QueryOsNameKey    = "SELECT Caption FROM Win32_OperatingSystem"
//	QueryOsVersionKey = "SELECT Version FROM Win32_OperatingSystem"
//	QueryHotFixKey    = "SELECT Description, HotFixID, InstalledBy, InstalledOn FROM Win32_QuickFixEngineering"
//)

func TestQuery(t *testing.T) {
	ret, err := QueryUuid()
	if err != nil {
		t.Error(err)
	}
	t.Log("uuid: ", ret)
	ret, err = QueryOsName()
	if err != nil {
		t.Error(err)
	}
	t.Log("osName: ", ret)
	ret, err = QueryOsVersion()
	if err != nil {
		t.Error(err)
	}
	t.Log("osVersion: ", ret)
	ret2, err := QueryHotFix()
	if err != nil {
		t.Error(err)
	}
	t.Log("hotFix: ", ret2)

}
