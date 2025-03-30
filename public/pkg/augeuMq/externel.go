package augueMq

import (
	"context"
	"errors"
)

var (
	defaultMq *MqCore
)

// Init
// 初始化 MqCore
func Init(maxSize uint) {
	if maxSize <= 0 {
		maxSize = 1
	}
	if maxSize >= 1000 {
		maxSize = 1000
	}
	defaultMq = &MqCore{
		exit:   make(chan bool),
		size:   maxSize,
		topics: make(map[string]*MqCell),
	}
}

// IsExit
// 检查 MqCore 是否存在
func IsExit() bool {
	select {
	case <-defaultMq.exit:
		return false
	default:
	}
	return true
}

// CheckCtxAndExit 检查上下文是否取消或者 MqCore 是否关闭
func CheckCtxAndExit(ctx context.Context) error {
	select {
	case <-defaultMq.exit:
		return errors.New("defaultMq is closed")
	case <-ctx.Done():
		return errors.New("operation cancelled due to timeout")
	default:
	}
	return nil
}

// SetMqCoreSize
// 设置 MqCore 的大小
func SetMqCoreSize(ctx context.Context, size uint) error {
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()
	return defaultMq.SetCoreSize(ctx, size)
}

// Exit 安全退出 MqCore
func Exit() {
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()

	close(defaultMq.exit)
}

// Set 向指定的 topic 添加消息
func Set(ctx context.Context, name string, msg interface{}) error {
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()
	return defaultMq.Set(ctx, name, msg)
}

// NewCell 创建一个新的 topic cell
func NewCell(ctx context.Context, name string, size int) error {
	// 直接操作 defaultMq，不通过内部的 base 方法
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()

	return defaultMq.NewCell(ctx, name, size)
}

// Get 从指定的 topic 获取消息
func Get(ctx context.Context, name string) (interface{}, error) {
	// 直接操作 defaultMq，不通过内部的 base 方法
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()

	return defaultMq.Get(ctx, name)
}

// SetCellSize 设置 MqCore 的大小
func SetCellSize(ctx context.Context, name string, size int) error {
	// 直接操作 defaultMq，不通过内部的 base 方法
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()
	return defaultMq.SetCellSize(ctx, name, size)

}

// IsFull 检查 MqCore 是否已满
func IsFull() bool {
	return len(defaultMq.topics) >= int(defaultMq.size)
}

func DeleteCell(ctx context.Context, name string) error {
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()
	return defaultMq.DeleteCell(ctx, name)
}

// IsExist 检查 MqCore 是否存在
func IsExist(name string) bool {
	defaultMq.lock.RLock()
	defer defaultMq.lock.RUnlock()
	_, ok := defaultMq.topics[name]
	return ok
}

func GetDefaultMq() *MqCore {
	return defaultMq
}
