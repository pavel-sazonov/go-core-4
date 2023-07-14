package geometry // поменял название

import (
	"fmt"
	"math"
)

// Geom представляет две точки.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// CalculateDistance рассчитывает расстояние между двумя точками, координаты не могут быть меньше 0.
func (geom Geom) CalculateDistance() (*float64, error) {
	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
		// добавил возврат ошибки
		// добавил возврат nil вместо -1
		return nil, fmt.Errorf("координаты не могут быть меньше нуля")
	}
	// убрал else

	res := math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	return &res, nil

}
