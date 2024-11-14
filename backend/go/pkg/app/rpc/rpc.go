package rpc

import "context"

type RPC interface {
	Diary() Diary
}

type Diary interface {
	Create(ctx context.Context) error
}
