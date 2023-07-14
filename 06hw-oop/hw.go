package geometry // поменял название

import (
	"fmt"
	"math"
)

// CalculateDistance рассчитывает расстояние между двумя точками, координаты не могут быть меньше 0.
func CalculateDistance(x1, y1, x2, y2 float64) (*float64, error) {
	if x1 < 0 || y1 < 0 || x2 < 0 || y2 < 0 {
		return nil, fmt.Errorf("координаты не могут быть меньше нуля")
	}

	res := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
	return &res, nil

}
