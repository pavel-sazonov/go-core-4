package geometry // поменял название

import (
	"fmt"
	"math"
)

// Geom представляет две точки.
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// CalculateDistance рассчитывает расстояние между двумя точками
// По условиям задачи, координаты не могут быть меньше 0.
func (geom Geom) CalculateDistance() (distance float64, err error) {
	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
		// добавил возврат ошибки и проверил в тесте
		return -1, fmt.Errorf("координаты не могут быть меньше нуля")
	}
	// убрал else

	distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	return distance, nil

}
