package geometry

import "testing"

func TestGeom_CalculateDistance(t *testing.T) {
	tests := []struct {
		name         string
		geom         Geom
		wantDistance float64
	}{
		{
			name:         "#1",
			geom:         Geom{X1: -1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: -1,
		},
		{
			name:         "#2",
			geom:         Geom{X1: 1, Y1: 1, X2: 4, Y2: 5},
			wantDistance: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := tt.geom.CalculateDistance(); err != nil {
				t.Log(err)
			}
			if gotDistance, _ := tt.geom.CalculateDistance(); gotDistance != tt.wantDistance {
				t.Errorf("Geom.CalculateDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}
