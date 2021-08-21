package tegola

import "github.com/go-spatial/geom"

const (
	WebMercator = 3857
	WGS84       = 4326
	NonEarth    = 6707
)

var (
	WebMercatorBounds = &geom.Extent{-20026376.39, -20048966.10, 20026376.39, 20048966.10}
	WGS84Bounds       = &geom.Extent{-180.0, -85.0511, 180.0, 85.0511}
	NonEarthBounds    = &geom.Extent{-180.0, -90.0, 180.0, 90.0}
)
