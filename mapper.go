package geo

// The Mapper Interface
// This interface specifies the responsibilities of a Mapper.
// A Mapper should be able to find all know Point structs within a given radius of a certain Point.
type Mapper interface {
	PointsWithinRadius(p *Point, radius int) bool
}
