package geo_test

import (
	"testing"

	"github.com/supernarsi/gotool/geo"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		name  string
		input [4]float64
		want  float64
	}{
		{name: "test 101, 102, 101.5, 102.5", input: [4]float64{101, 102, 101.5, 102.5}, want: 56645.09},
		{name: "test 83, 102, 89.40123, 42.11074", input: [4]float64{83, 102, 78.61059437770405, 15.995051388871412}, want: 1436805.4},
		{name: "test 33, 20.1234, 19.40123, 42.8765", input: [4]float64{33, 20.1234, 19.40123, 42.8765}, want: 2717422.87},
		{name: "test 0, 0.001, 1.001, 0.10101", input: [4]float64{0, 0.001, 1.001, 0.10101}, want: 111860.21},
		{name: "test 4,3,2,1", input: [4]float64{4, 3, 2, 1}, want: 314283.25},
		{name: "test 10,10,10,10", input: [4]float64{10, 10, 10, 10}, want: 0},
		{name: "test 10,10,10,10.00001", input: [4]float64{10, 10, 10, 10.00001}, want: 1.09},
		{name: "test Barcelona", input: [4]float64{41.40364, 2.17440, 41.36385, 2.15246}, want: 4788.13},
		{name: "test BJ to SH", input: [4]float64{39.9075, 116.39723, 31.23037, 121.4737}, want: 1068033.42},
		{name: "test small distance", input: [4]float64{39.875211897994895, 116.40458367375038, 39.87448996132426, 116.40498141917334}, want: 87.15},
		{name: "test min distance", input: [4]float64{89.999999, 0, 90.000001, 0}, want: 0.21},
		{name: "test near distance", input: [4]float64{39.908254, 116.397242, 39.908246, 116.398055}, want: 69.35},
		{name: "test_13", input: [4]float64{29.53691983396626, 106.50190658067017, 29.536914999999997, 106.50191000000001}, want: 0.62},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := geo.Calculator(geo.Cosines).Distance(tt.input[0], tt.input[1], tt.input[2], tt.input[3]); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDistanceByGPT(t *testing.T) {
	tests := []struct {
		name  string
		input [4]float64
		want  float64
	}{
		{name: "test 101, 102, 101.5, 102.5", input: [4]float64{101, 102, 101.5, 102.5}, want: 56645.09},
		{name: "test 83, 102, 89.40123, 42.11074", input: [4]float64{83, 102, 78.61059437770405, 15.995051388871412}, want: 1436805.4},
		{name: "test 33, 20.1234, 19.40123, 42.8765", input: [4]float64{33, 20.1234, 19.40123, 42.8765}, want: 2717422.87},
		{name: "test 0, 0.001, 1.001, 0.10101", input: [4]float64{0, 0.001, 1.001, 0.10101}, want: 111860.21},
		{name: "test 4,3,2,1", input: [4]float64{4, 3, 2, 1}, want: 314283.25},
		{name: "test 10,10,10,10", input: [4]float64{10, 10, 10, 10}, want: 0},
		{name: "test 10,10,10,10.00001", input: [4]float64{10, 10, 10, 10.00001}, want: 1.09},
		{name: "test Barcelona", input: [4]float64{41.40364, 2.17440, 41.36385, 2.15246}, want: 4788.13},
		{name: "test BJ to SH", input: [4]float64{39.9075, 116.39723, 31.23037, 121.4737}, want: 1068033.42},
		{name: "test small distance", input: [4]float64{39.875211897994895, 116.40458367375038, 39.87448996132426, 116.40498141917334}, want: 87.15},
		{name: "test min distance", input: [4]float64{89.999999, 0, 90.000001, 0}, want: 0.22},
		{name: "test near distance", input: [4]float64{39.908254, 116.397242, 39.908246, 116.398055}, want: 69.35},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Cleanup(func() {})
			if got := geo.Calculator(geo.Haversine).Distance(tt.input[0], tt.input[1], tt.input[2], tt.input[3]); got != tt.want {
				t.Errorf("got result %v, want %v", got, tt.want)
			}
		})
	}
}
