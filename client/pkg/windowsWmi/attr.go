package windowsWmi

type QueryString struct {
	Value string
}

type QueryStrings struct {
	Values []string
}

type QueryFunc func(query string) ([]string, error)

type QueryKey string

type Win32_QuickFixEngineering struct {
	Description string
	HotFixID    string
	InstalledBy string
	InstalledOn string
}

// 确保结构体字段与WMI属性映射
type Win32_ComputerSystemProduct struct {
	UUID string `wmi:"UUID"`
}

type Win32_OperatingSystem struct {
	Caption string `wmi:"Caption"`
}

type Win32_OperatingSystemVersion struct {
	Version string `wmi:"Version"`
}
