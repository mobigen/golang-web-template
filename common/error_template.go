package common

import "errors"

var (
	// ErrNoHaveResult 데이터 베이스 조회 결과 없음
	ErrNoHaveResult = errors.New("no have result")
)
