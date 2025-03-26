package ProcessInfo

type ProcessImport struct {
	ImportID        uint   `gorm:"primaryKey;autoIncrement;column:import_id;type:serial"`
	PID             int    `gorm:"column:pid;type:integer;not null"`
	ImportPath      string `gorm:"column:import_path;type:text;not null"`
	ImportMD5       string `gorm:"column:import_md5;type:char(32)"`
	ImportSignature string `gorm:"column:import_signature;type:text"`
}

func (ProcessImport) TableName() string {
	return "process_imports"
}
