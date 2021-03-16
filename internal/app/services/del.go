package services

import "context"

//删除的接口
type IDel interface {
	Execution(ctx context.Context, uid string) error
}
