package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	lines := util.Lines("input.txt")

	bin := hex2bin(lines[0])
	packet, _ := parsePacket(bin, 0)

	fmt.Println("part1:", versionSum(packet))
	fmt.Println("part2:", evaluate(packet))
}

type packet struct {
	version int
	typeId  int
	literal int
	subs    []packet
}

func hex2bin(s string) string {
	var sb strings.Builder

	for i := range s {
		v, _ := strconv.ParseUint(s[i:i+1], 16, 8)
		sb.WriteString(fmt.Sprintf("%04b", v))
	}

	return sb.String()
}

func bin2int(s string) int {
	v, _ := strconv.ParseInt(fmt.Sprintf("0b%s", s), 0, 0)
	return int(v)
}

func parsePacket(s string, ptr int) (packet, int) {
	p := packet{}

	p.version = bin2int(s[ptr : ptr+3])
	ptr += 3
	p.typeId = bin2int(s[ptr : ptr+3])
	ptr += 3

	if p.typeId == 4 {
		var lit strings.Builder
		for {
			done := s[ptr] == '0'
			lit.WriteString(s[ptr+1 : ptr+5])
			ptr += 5
			if done {
				break
			}
		}
		p.literal = bin2int(lit.String())
	} else {
		if s[ptr] == '0' {
			ptr += 1
			subLength := bin2int(s[ptr : ptr+15])
			ptr += 15
			done := ptr + subLength
			for {
				var subPacket packet
				subPacket, ptr = parsePacket(s, ptr)
				p.subs = append(p.subs, subPacket)
				if ptr == done {
					break
				}
			}
		} else {
			ptr += 1
			subPackets := bin2int(s[ptr : ptr+11])
			ptr += 11
			for i := 0; i < subPackets; i++ {
				var subPacket packet
				subPacket, ptr = parsePacket(s, ptr)
				p.subs = append(p.subs, subPacket)
			}
		}
	}

	return p, ptr
}

func versionSum(p packet) int {
	sum := p.version
	for _, s := range p.subs {
		sum += versionSum(s)
	}
	return sum
}

func evaluate(p packet) int {
	switch p.typeId {
	case 0:
		val := evaluate(p.subs[0])
		for i := 1; i < len(p.subs); i++ {
			val += evaluate(p.subs[i])
		}
		return val
	case 1:
		val := evaluate(p.subs[0])
		for i := 1; i < len(p.subs); i++ {
			val *= evaluate(p.subs[i])
		}
		return val
	case 2:
		val := evaluate(p.subs[0])
		for i := 1; i < len(p.subs); i++ {
			x := evaluate(p.subs[i])
			if x < val {
				val = x
			}
		}
		return val
	case 3:
		val := evaluate(p.subs[0])
		for i := 1; i < len(p.subs); i++ {
			x := evaluate(p.subs[i])
			if x > val {
				val = x
			}
		}
		return val
	case 4:
		return p.literal
	case 5:
		a := evaluate(p.subs[0])
		b := evaluate(p.subs[1])
		if a > b {
			return 1
		} else {
			return 0
		}
	case 6:
		a := evaluate(p.subs[0])
		b := evaluate(p.subs[1])
		if a < b {
			return 1
		} else {
			return 0
		}
	case 7:
		a := evaluate(p.subs[0])
		b := evaluate(p.subs[1])
		if a == b {
			return 1
		} else {
			return 0
		}
	default:
		panic("help")
	}
}
