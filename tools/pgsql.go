package tools

import (
	"errors"
	"fmt"
	"github.com/lib/pq"
)

func IsDuplicateKeyError(err error) bool {
	var pgErr *pq.Error
	ok := errors.As(err, &pgErr)
	if !ok {
		fmt.Println("not a pq error")
	}
	if ok && pgErr.Code == "23505" {
		// 处理重复键错误
		return true
	}
	return false
}
