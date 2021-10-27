package models

var MexicoBBox = BoundingBox{
	19.296134,
	-99.296741,
	19.661237,
	-98.916339,
}

type Location struct {
	Latitude  float64
	Longitude float64
}

type BoundingBox struct {
	Left   float64
	Bottom float64
	Right  float64
	Top    float64
}

func (bb BoundingBox) InBoundingBox(l Location) bool {
	return bb.Top <= l.Longitude &&
		l.Longitude <= bb.Bottom &&
		bb.Left <= l.Latitude &&
		l.Latitude <= bb.Right
}
