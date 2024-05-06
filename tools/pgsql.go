package tools

import (
	"errors"
	"github.com/lib/pq"
)

func IsDuplicateKeyError(err error) bool {
	var pgErr *pq.Error
	ok := errors.As(err, &pgErr)
	if ok && pgErr.Code == "23505" {
		// 处理重复键错误
		return true
	}
	return false
}
