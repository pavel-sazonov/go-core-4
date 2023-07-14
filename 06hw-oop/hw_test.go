package geometry

import (
	"fmt"
	"testing"
)

func TestGeom_CalculateDistance(t *testing.T) {
	wantDistance := 5.0
	wantErr := fmt.Errorf("координаты не могут быть меньше нуля")
	tests := []struct {
		name           string
		x1, y1, x2, y2 float64
		wantDistance   *float64
	}{
		{
			name: "#1",
			x1:   -1, y1: 1, x2: 4, y2: 5,
			wantDistance: nil,
		},
		{
			name: "#2",
			x1:   1, y1: 1, x2: 4, y2: 5,
			wantDistance: &wantDistance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance, err := CalculateDistance(tt.x1, tt.y1, tt.x2, tt.y2); err != nil {
				if gotDistance != tt.wantDistance {
					t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
				}

				if wantErr.Error() != err.Error() {
					t.Errorf("err = %v, want %v", err, wantErr)
				}
			}

			if gotDistance, err := CalculateDistance(tt.x1, tt.y1, tt.x2, tt.y2); err == nil {
				// получилось переписать тест, только через if
				// так как была паника, когда было обращение к *gotDistance == nil
				if *gotDistance != *tt.wantDistance {
					t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
				}
			}
		})
	}
}
