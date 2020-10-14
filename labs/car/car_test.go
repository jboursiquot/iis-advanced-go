package car_test

import (
	"testing"

	car "github.com/jboursiquot/iis-advanced-go/labs/car"
)

func TestIncrease(t *testing.T) {
	tests := map[string]struct {
		curr  int
		accel int
		want  int
	}{
		"50-60": {
			curr:  50,
			accel: 10,
			want:  60,
		},
		"70-100": {
			curr:  70,
			accel: 30,
			want:  100,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := car.Increase(tc.accel, tc.curr); got != tc.want {
				t.Errorf("wanted %d but got %d", tc.want, got)
			}
		})
	}
}
