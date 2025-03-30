package windowsWmi

const (
	QueryUuidKey      = "SELECT UUID FROM Win32_ComputerSystemProduct"
	QueryOsNameKey    = "SELECT Caption FROM Win32_OperatingSystem"
	QueryOsVersionKey = "SELECT Version FROM Win32_OperatingSystem"
	QueryHotFixKey    = "SELECT Description, HotFixID, InstalledBy, InstalledOn FROM Win32_QuickFixEngineering"
)
