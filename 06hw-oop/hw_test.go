package geometry

import (
	"fmt"
	"testing"
)

func TestGeom_CalculateDistance(t *testing.T) {
	wantDistance := 5.0
	wantErr := fmt.Errorf("координаты не могут быть меньше нуля")
	tests := []struct {
		name         string
		geom         Geom
		wantDistance *float64
	}{
		{
			name:         "#1",
			geom:         Geom{X1: -1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: nil,
		},
		{
			name:         "#2",
			geom:         Geom{X1: 1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: &wantDistance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance, err := tt.geom.CalculateDistance(); err != nil {
				if gotDistance != tt.wantDistance {
					t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
				}

				if wantErr.Error() != err.Error() {
					t.Errorf("err = %v, want %v", err, wantErr)
				}
			}

			if gotDistance, err := tt.geom.CalculateDistance(); err == nil {
				// получилось переписать тест, только через if
				// так как была паника, когда было обращение к *gotDistance == nil
				if *gotDistance != *tt.wantDistance {
					t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
				}
			}
		})
	}
}
