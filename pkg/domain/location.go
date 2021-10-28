package domain

var MexicoBBox = BoundingBox{
	-99.296741,
	19.296134,
	-98.916339,
	19.661237,
}

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type BoundingBox struct {
	MinLongitude float64
	MinLatitude  float64
	MaxLongitude float64
	MaxLatitude  float64
}

func (bb BoundingBox) InBoundingBox(l Location) bool {
	return bb.MinLatitude <= l.Latitude && l.Latitude <= bb.MaxLatitude &&
		bb.MinLongitude <= l.Longitude && l.Longitude <= bb.MaxLongitude
}

func (l Location) InBoundingBox(box BoundingBox) bool {
	return box.InBoundingBox(l)
}
