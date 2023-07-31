package util

import (
	"math"

	"github.com/Arch-4ng3l/GoServerHololens/types"
)

func CalcDistance(obj1, obj2 *types.Object) float32 {
	d := math.Sqrt(math.Pow(float64(obj2.X-obj1.X), 2) + math.Pow(float64(obj2.Z-obj1.Z), 2))
	return float32(d)
}
