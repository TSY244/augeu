package augueMq

import (
	"errors"
	"sync"
)

type MqCell struct {
	cell    chan interface{}
	isExit  bool
	maxSize int        // 最大限制
	lock    sync.Mutex // 对 MqCell 的操作加锁
}

// NewMqCell 创建一个新的 MqCell
func NewMqCell(size int) *MqCell {
	if size <= 0 {
		size = 1
	}
	if size >= 1000 {
		size = 1000
	}

	return &MqCell{
		cell:    make(chan interface{}, size),
		isExit:  false,
		maxSize: size,
	}
}

// Set 添加消息到 MqCell
func (m *MqCell) Set(msg interface{}) error {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果已经退出，不再接受消息
	if m.isExit {
		return errors.New("MqCell is exited")
	}

	// 如果队列已满，返回错误
	if len(m.cell) >= m.maxSize {
		return errors.New("MqCell is full")
	}

	m.cell <- msg
	return nil
}

// Get 从 MqCell 中获取消息
func (m *MqCell) Get() (interface{}, error) {
	m.lock.Lock()
	defer m.lock.Unlock()

	// 如果已经退出，不再接受消息
	if m.isExit {
		return nil, errors.New("MqCell is exited")
	}

	if len(m.cell) == 0 {
		return nil, errors.New("MqCell is empty")
	}

	return <-m.cell, nil
}

// Exit 退出并关闭 MqCell
func (m *MqCell) Exit() {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.isExit {
		return
	}

	close(m.cell)
	m.isExit = true
}

// IsExit 检查 MqCell 是否已退出
func (m *MqCell) IsExit() bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	return m.isExit
}

// Clear 清空 MqCell 并重置状态
func (m *MqCell) Clear() {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.cell = make(chan interface{}, m.maxSize)
	m.isExit = false
}

// GetSize 获取当前 MqCell 中存储的消息数量
func (m *MqCell) GetSize() int {
	m.lock.Lock()
	defer m.lock.Unlock()

	return len(m.cell)
}

func (m *MqCell) SetMaxSize(size int) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if size <= 0 {
		size = 1
	}
	if size >= 1000 {
		size = 1000
	}

	m.maxSize = size
}

func (m *MqCell) IsHaveMsg(name string) bool {
	m.lock.Lock()
	defer m.lock.Unlock()

	if len(m.cell) == 0 {
		return false
	}
	return true
}

func (m *MqCell) GetCell() chan interface{} {
	return m.cell
}
