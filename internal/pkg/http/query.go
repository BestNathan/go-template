package http

import (
	"go-template/internal/pkg/json"
	"go-template/internal/pkg/time"
)

type Query struct {
	Page      int               `form:"page" json:"page"`
	Size      int               `form:"size" json:"size"`
	Search    string            `form:"search" json:"search"`
	StartTime time.Timestamp    `form:"startTime" json:"startTime"`
	EndTime   time.Timestamp    `form:"endTime" json:"endTime"`
	Test      json.String2Array `form:"test" json:"test"`
}

func (q *Query) Limit() int {
	if q.Size < 1 || q.Size > 20 {
		return 20
	}

	return q.Size
}

func (q *Query) Offset() int {
	if q.Page < 1 {
		return 0
	}

	return q.Size * (q.Page - 1)
}
