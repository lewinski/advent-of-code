package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	scans := []scan{}
	for _, rec := range util.Records("input.txt") {
		scans = append(scans, parseScan(rec))
	}
	alignScanners(scans)

	fmt.Println("part1:", part1(scans))
	fmt.Println("part2:", part2(scans))
}

type scan struct {
	beacons  util.IntGrid3
	position util.Point3 // relative to scan 0
	solved   bool
}

func newScan() scan {
	return scan{
		beacons:  util.IntGrid3{},
		position: util.Origin3(),
		solved:   false,
	}
}

func parseScan(s string) scan {
	rv := newScan()

	for i, line := range strings.Split(s, "\n") {
		if i == 0 {
			continue
		}
		coords := strings.Split(line, ",")
		rv.beacons.SetCoords(
			util.MustAtoi(coords[0]),
			util.MustAtoi(coords[1]),
			util.MustAtoi(coords[2]),
			1)
	}

	return rv
}

func alignScanners(scans []scan) {
	// first one is solved
	scans[0].solved = true

	for {
		// check if we've solved every one
		solved := 0
		for i := 0; i < len(scans); i++ {
			if scans[i].solved {
				solved++
			}
		}
		if solved == len(scans) {
			break
		}

		for i := 0; i < len(scans); i++ {
			if !scans[i].solved {
				continue
			}
			for j := 0; j < len(scans); j++ {
				if i == j {
					continue
				}
				if scans[j].solved {
					continue
				}

				// for each pair i, j where i is solved and j is unsolved, try to match them up
				matched, rotation, offset := tryMatch(scans[j], scans[i])
				if matched {
					// if matched up, then realign j into the shared reference point
					newBeacons := util.IntGrid3{}
					scans[j].beacons.Each(func(p util.Point3, x int) {
						newBeacons.Set(rotation(p).Offset(offset), x)
					})
					scans[j].beacons = newBeacons

					// record the position and mark this one solved
					scans[j].position = offset
					scans[j].solved = true
				}
			}
		}
	}
}

func part1(scans []scan) int {
	allBeacons := util.IntGrid3{}
	for i := 0; i < len(scans); i++ {
		scans[i].beacons.Each(func(p util.Point3, x int) {
			allBeacons.Set(p, allBeacons.Get(p)+1)
		})
	}

	return len(allBeacons)
}

func part2(scans []scan) int {
	best := 0

	for i := 0; i < len(scans); i++ {
		for j := 0; j < len(scans); j++ {
			dist := util.IAbs(scans[i].position[0]-scans[j].position[0]) +
				util.IAbs(scans[i].position[1]-scans[j].position[1]) +
				util.IAbs(scans[i].position[2]-scans[j].position[2])
			if dist > best {
				best = dist
			}
		}
	}
	return best
}
func tryMatch(unknown, known scan) (bool, rotationFunc, util.Point3) {
	rotations := allRotationFuncs()

	matched := false
	matchedRotation := rotations[0]
	matchedOffset := util.Origin3()

	// try all 24 rotations
	for _, rot := range rotations {
		// reorient the unknown points
		b := util.IntGrid3{}
		unknown.beacons.Each(func(p util.Point3, x int) {
			b.Set(rot(p), x)
		})

		// for each point in the unknown set
		b.Each(func(unkref util.Point3, x int) {
			// already done
			if matched {
				return
			}

			// try a point in the known set
			known.beacons.Each(func(knref util.Point3, x int) {
				// already done
				if matched {
					return
				}

				// and calculate an offset between them
				offset := util.Point3{knref[0] - unkref[0], knref[1] - unkref[1], knref[2] - unkref[2]}

				// and then see if this offset causes other overlaps
				count := 0
				b.Each(func(p util.Point3, x int) {
					if known.beacons.Contains(p.Offset(offset)) {
						count++
					}
				})

				// if we counted more than 12, then we aligned this region and want to return
				if count >= 12 {
					matched = true
					matchedRotation = rot
					matchedOffset = offset
				}
			})
		})

		if matched {
			break
		}
	}

	return matched, matchedRotation, matchedOffset
}

/*

24 non-mirrored orientations of cube

finally a use for my xyz calibration cube collection!

x, y, z
-y, x, z
-x, -y, z
y, -x, z

x, -y, -z
y, x, -z
-x, y, z
-y, -x, -z

x, -z, y
z, x, y
-x, z, y
-z, -x, y

x, z, -y
-z, x, -y
-x, -z, -y
z, -x, -y

y, z, x
-z, y, x
-y, -z, x
z, -y, x

y, -z, -x
z, y, -x
-y, z, -x
-z, -y, -x

*/

type rotationFunc func(p util.Point3) util.Point3

func allRotationFuncs() []rotationFunc {
	return []rotationFunc{
		func(p util.Point3) util.Point3 { return util.Point3{p[0], p[1], p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[1], p[0], p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[0], -p[1], p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[1], -p[0], p[2]} },

		func(p util.Point3) util.Point3 { return util.Point3{p[0], -p[1], -p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[1], p[0], -p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[0], p[1], -p[2]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[1], -p[0], -p[2]} },

		func(p util.Point3) util.Point3 { return util.Point3{p[0], -p[2], p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[2], p[0], p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[0], p[2], p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[2], -p[0], p[1]} },

		func(p util.Point3) util.Point3 { return util.Point3{p[0], p[2], -p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[2], p[0], -p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[0], -p[2], -p[1]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[2], -p[0], -p[1]} },

		func(p util.Point3) util.Point3 { return util.Point3{p[1], p[2], p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[2], p[1], p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[1], -p[2], p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[2], -p[1], p[0]} },

		func(p util.Point3) util.Point3 { return util.Point3{p[1], -p[2], -p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{p[2], p[1], -p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[1], p[2], -p[0]} },
		func(p util.Point3) util.Point3 { return util.Point3{-p[2], -p[1], -p[0]} },
	}
}
