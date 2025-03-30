package augueMq

import (
	"context"
	"errors"
	"sync"
)

const (
	CheckTrue  = true
	CheckFalse = false
)

type MqCore struct {
	exit   chan bool
	size   uint // 订阅者的限制
	topics map[string]*MqCell
	lock   sync.RWMutex // 读写锁
}

// SetCoreSize
// 设置 MqCore 的大小
//
// 注意：
//   - 如果 size 小于当前的 size，返回错误
//   - 这个Size 是指 MqCore 的大小，而不是某个 topic 的大小
func (mq *MqCore) SetCoreSize(ctx context.Context, size uint) error {
	return mq.base(ctx, CheckFalse, func() error {
		if len(mq.topics) >= int(size) {
			return errors.New("defaultMq is full")
		}

		mq.size = size
		return nil
	})
}

// Exit
// 安全退出 MqCore
func (mq *MqCore) Exit(ctx context.Context) error {
	return mq.base(ctx, CheckFalse, func() error {
		close(mq.exit)
		return nil
	})
}

// Set
// 向指定的 topic 添加消息
func (mq *MqCore) Set(ctx context.Context, name string, msg interface{}) error {
	return mq.base(ctx, CheckFalse, func() error {
		if cell, ok := mq.topics[name]; ok {
			return cell.Set(msg)
		}
		return errors.New("topic does not exist")
	})
}

// NewCell
// 创建一个新的 topic cell
func (mq *MqCore) NewCell(ctx context.Context, name string, size int) error {
	return mq.base(ctx, CheckTrue, func() error {
		if _, ok := mq.topics[name]; ok {
			return TopicAlreadyExists
		}
		mq.topics[name] = NewMqCell(size)
		return nil
	})
}

// IsFull
// 检查 MqCore 是否已满
func (mq *MqCore) IsFull() bool {
	return uint(len(mq.topics)) >= mq.size
}

// Get
// 从指定的 topic 获取消息
func (mq *MqCore) Get(ctx context.Context, name string) (interface{}, error) {
	var msg interface{}
	var err error
	err = mq.base(ctx, CheckFalse, func() error {
		if cell, ok := mq.topics[name]; ok {
			msg, err = cell.Get()
			return err
		} else {
			return errors.New("topic does not exist")
		}
	})
	return msg, err
}

func (mq *MqCore) checkCtxAndExit(ctx context.Context) error {
	select {
	case <-mq.exit:
		return errors.New("defaultMq is closed")
	case <-ctx.Done():
		return errors.New("operation cancelled due to timeout")
	default:
	}
	return nil
}

func (mq *MqCore) base(ctx context.Context, check bool, f func() error) error {
	if err := mq.checkCtxAndExit(ctx); err != nil {
		return err
	}
	if check {
		if mq.IsFull() {
			return errors.New("defaultMq is full")
		}
	}
	return f()
}

func (mq *MqCore) SetCellSize(ctx context.Context, name string, size int) error {
	return mq.base(ctx, CheckFalse, func() error {
		if cell, ok := mq.topics[name]; ok {
			cell.SetMaxSize(size)
			return nil
		}
		return errors.New("topic does not exist")
	})
}

func (mq *MqCore) DeleteCell(ctx context.Context, name string) error {
	return mq.base(ctx, CheckFalse, func() error {
		if _, ok := mq.topics[name]; ok {
			delete(mq.topics, name)
			return nil
		}
		return errors.New("topic does not exist")
	})
}

func (mq *MqCore) GetTopic(ctx context.Context, name string) (*MqCell, error) {
	var topics *MqCell
	err := mq.base(ctx, CheckFalse, func() error {
		if cell, ok := mq.topics[name]; ok {
			topics = cell
			return nil
		}
		return nil
	})
	return topics, err
}
