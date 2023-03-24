package geo

import (
	"math"
)

const earthRadius = 6371

type mathType uint8

const (
	Cosines mathType = iota
	Haversine
)

type distance interface {
	Distance(lat1, lat2, lng1, lng2 float64) float64
}

type cosines struct{}
type haversine struct{}

var distanceCalculator distance

func (d *cosines) Distance(lat1, lng1, lat2, lng2 float64) float64 {
	if lat1 == lat2 && lng1 == lng2 {
		return 0
	}
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad
	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))
	num := dist * earthRadius * 1000
	return math.Trunc(num*100) / 100
}

func (d *haversine) Distance(lat1, lng1, lat2, lng2 float64) float64 {
	// 将经纬度转换为弧度
	rad := math.Pi / 180.0
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	// 计算差值
	dLat := lat2 - lat1
	dLon := lng2 - lng1

	// 使用 haversine 公式计算距离
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := earthRadius * c * 1000
	return math.Trunc(distance*100) / 100
}

func Calculator(calType mathType) distance {
	switch calType {
	case Cosines:
		distanceCalculator = &cosines{}
	case Haversine:
		distanceCalculator = &haversine{}
	default:

	}
	return distanceCalculator
}
