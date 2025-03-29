package augue_mq

import (
	"context"
	"fmt"
	"testing"
)

func TestMqCore(t *testing.T) {
	// 创建背景上下文
	ctx := context.Background()

	// 1. 初始化mq
	Init(1)
	if !IsExit() {
		t.Fatal("mq is not exist")
	}

	// 2. 添加一个topic
	err := NewCell(ctx, "test", 1)
	if err != nil {
		t.Fatalf("failed to create topic 'test': %v", err)
	}

	// 3. 添加消息到topic
	err = Set(ctx, "test", "hello")
	if err != nil {
		t.Fatalf("failed to set message to 'test': %v", err)
	}

	// 4. 从topic中获取消息
	msg, err := Get(ctx, "test")
	if err != nil {
		t.Fatalf("failed to get message from 'test': %v", err)
	}
	if msg.(string) != "hello" {
		t.Errorf("expected 'hello', got: %v", msg)
	}
	fmt.Println(msg)

	// 5. 超出消息队列容量
	err = Set(ctx, "test", "world")
	if err != nil {
		t.Fatalf("failed to set message to 'test': %v", err)
	}

	// 尝试超出容量，应该失败
	err = Set(ctx, "test", "august")
	if err == nil {
		t.Errorf("expected 'MqCell is full' error, got: %v", err)
	}

	// 6. 设置mq的大小
	err = SetMqCoreSize(ctx, 2)
	if err != nil {
		t.Fatalf("failed to set mq size: %v", err)
	}

	// 7. 创建一个新的topic
	err = NewCell(ctx, "test2", 1)
	if err != nil {
		t.Fatalf("failed to create topic 'test2': %v", err)
	}

	// 8. 销毁第一个topic
	err = DeleteCell(ctx, "test")
	if err != nil {
		t.Fatalf("failed to delete topic 'test': %v", err)
	}

	// 9. 退出mq
	Exit()

	// 10. 检查退出后的行为
	err = Set(ctx, "test2", "hello")
	if err == nil {
		t.Fatal("expected error when setting message after Exit, but got nil")
	}
}
