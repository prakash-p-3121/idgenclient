package idgenclient

import (
	"github.com/prakash-p-3121/errorlib"
	"github.com/prakash-p-3121/idgenmodel"
)

type IDGenClient interface {
	NextID(tableName string) (*idgenmodel.IDGenResp, errorlib.AppError)
}
