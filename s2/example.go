package main

import (
	s2 "github.com/golang/geo/s2"
)

func main() {
	s2.CapFromCenterAngle("")
	rc := &s2.RegionCoverer{MaxLevel: 30, MaxCells: 5}
	r := s2.Region(CapFromCenterArea(center, area))
	covering := rc.Covering(r)

}
