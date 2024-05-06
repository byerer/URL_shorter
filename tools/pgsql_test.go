package tools

import (
	"github.com/lib/pq"
	"testing"
)

func TestIsDuplicateKeyError(t *testing.T) {
	// 模拟一个pq错误
	err := &pq.Error{Code: "23505"}

	// 测试函数
	if !IsDuplicateKeyError(err) {
		t.Errorf("Expected true, got false")
	}
}
