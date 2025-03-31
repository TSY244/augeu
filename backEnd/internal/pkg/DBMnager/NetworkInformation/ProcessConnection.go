package NetworkInformation

import (
	"gorm.io/gorm"
	"net"
	"time"
)

type ProcessConnection struct {
	ProcessName      string         `gorm:"column:process_name;not null;size:255"` // 进程名
	PID              int            `gorm:"column:pid;not null"`                   // 进程ID
	Protocol         *string        `gorm:"column:protocol;size:10"`               // 协议 (TCP/UDP)
	LocalAddress     net.IP         `gorm:"column:local_address;type:inet"`        // 本地监听地址
	LocalPort        uint16         `gorm:"column:local_port"`                     // 本地监听端口 (0-65535)
	RemoteAddress    net.IP         `gorm:"column:remote_address;type:inet"`       // 远程地址
	RemotePort       uint16         `gorm:"column:remote_port"`                    // 远程端口 (0-65535)
	ConnectionStatus string         `gorm:"column:connection_status;size:20"`      // 连接状态
	CreateAt         time.Time      `gorm:"autoCreateTime"`
	DeleteAt         gorm.DeletedAt `gorm:"index"`
	ID               uint           `gorm:"primaryKey;autoIncrement"`
	UUID             string         `gorm:"type:varchar(255);column:uuid"`
}

func (ProcessConnection) TableName() string {
	return "process_connections"
}

func (ProcessConnection) BeforeCreate(tx *gorm.DB) (err error) {
	return tx.Exec(`
        ALTER TABLE process_connections ADD CONSTRAINT ck_local_port_range 
        CHECK (local_port BETWEEN 0 AND 65535);
    `).Error
}

func (ProcessConnection) AfterCreate(tx *gorm.DB) (err error) {
	return tx.Exec(`
        ALTER TABLE process_connections ADD CONSTRAINT ck_remote_port_range 
        CHECK (remote_port BETWEEN 0 AND 65535);
    `).Error
}

func (ProcessConnection) AfterUpdate(tx *gorm.DB) (err error) {
	return tx.Exec(`
        ALTER TABLE process_connections ADD CONSTRAINT ck_connection_status_enum 
        CHECK (connection_status IN (
            'LISTEN', 'ESTABLISHED', 'TIME_WAIT', 'CLOSE_WAIT', 
            'SYN_SENT', 'SYN_RECEIVED', 'CLOSED'
        ));
    `).Error
}
