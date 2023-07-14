package geometry // поменял название

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

// По условиям задачи, координаты не могут быть меньше 0.
func (geom Geom) CalculateDistance() (distance float64, err error) {
	// добавил возврат ошибки и проверил в тесте
	if geom.X1 < 0 || geom.X2 < 0 || geom.Y1 < 0 || geom.Y2 < 0 {
		return -1, fmt.Errorf("координаты не могут быть меньше нуля")
	}
	// убрал else

	// возврат расстояния между точками
	distance = math.Sqrt(math.Pow(geom.X2-geom.X1, 2) + math.Pow(geom.Y2-geom.Y1, 2))
	return distance, nil

}
