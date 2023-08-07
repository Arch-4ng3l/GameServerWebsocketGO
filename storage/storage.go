package storage

import "github.com/Arch-4ng3l/GoServerHololens/types"

type Storage interface {
	GetObjects(chan *types.Object, *types.Object) error
	UpdateObject(*types.Object) error
	GetObjectsWeb(int) ([]*types.Object, error)
	GetUser(string) *types.User
	GetAmountOfObjects() uint
}
