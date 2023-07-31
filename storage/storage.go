package storage

import "github.com/Arch-4ng3l/GoServerHololens/types"

type Storage interface {
	GetObjects(*types.Object) ([]*types.Object, error)
	UpdateObject(*types.Object) error
}
