package bootstrap

import (
	"go-quickstart/constant/sys"
	"go-quickstart/infrastructure/pool"
)

func NewPoolFactory() *PoolsFactory {
	p := make(map[int]*pool.Pool, sys.GoRoutinePoolTypesAmount)
	return &PoolsFactory{Pools: p}
}
