package server

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/Arch-4ng3l/GoServerHololens/storage"
	"github.com/Arch-4ng3l/GoServerHololens/types"
)

var store = storage.NewPostgres()
var err = store.Init()

func BenchmarkGetObjects(b *testing.B) {
	obj := &types.Object{}
	obj.Name = "Player"
	for n := 0; n < b.N; n++ {
		randomObject(obj)
		objs, err := store.GetObjects(obj)
		if err != nil {
			b.Fatal(err.Error())
		}
		if len(objs) < 1 {
			b.Fatal("No Objects got found")
		}
	}
}

func BenchmarkAddObject(b *testing.B) {
	obj := &types.Object{}
	for n := 0; n < b.N; n++ {
		obj.Name = fmt.Sprintf("cube%d", n)
		randomObject(obj)
		err := store.UpdateObject(obj)
		if err != nil {
			b.Fatal(err.Error())
		}
	}
}

func randomObject(obj *types.Object) {
	obj.X = float32(rand.Int() % 200)
	obj.Y = float32(rand.Int() % 200)
	obj.Z = float32(rand.Int() % 200)
}
