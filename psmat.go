package psmat

import (
	"sync"

	python3 "go.nhat.io/python/v3"
)

const moduleName = "psMat"

var getModule = sync.OnceValue(func() *python3.Object {
	return python3.MustImportModule(moduleName)
})

// Scale returns a matrix which will scale by x in the horizontal direction and y in the vertical.
func Scale(x, y float64) []float64 {
	matrix := getModule().CallMethodArgs("scale", x, y)
	defer matrix.DecRef()

	return python3.MustUnmarshalAs[[]float64](matrix)
}
