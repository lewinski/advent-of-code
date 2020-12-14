package main

import (
	"fmt"
	"strings"

	"github.com/lewinski/advent-of-code/util"
)

func main() {
	input := util.Lines("input.txt")

	fmt.Println("part1:", sumMemory(bootloaderVersion1(input)))
	fmt.Println("part2:", sumMemory(bootloaderVersion2(input)))
}

type memory map[int]int

func bootloaderVersion1(commands []string) memory {
	memory := memory{}

	var maskBits, maskValue int
	for _, command := range commands {
		if strings.HasPrefix(command, "mask = ") {
			maskBits, maskValue = decodeMask(command)
		} else {
			addr, val := decodeMem(command)
			memory[addr] = maskValue | int(val) & ^maskBits
		}
	}

	return memory
}

func bootloaderVersion2(commands []string) memory {
	memory := memory{}

	var maskBits, maskValue int
	for _, command := range commands {
		if strings.HasPrefix(command, "mask = ") {
			maskBits, maskValue = decodeMask(command)
		} else {
			baseAddr, val := decodeMem(command)
			for _, addr := range decodeAddresses(baseAddr, maskBits, maskValue) {
				memory[addr] = val
			}
		}
	}
	return memory
}

func decodeAddresses(addr int, maskBits, maskValue int) []int {
	// apply mask 1s
	addr = addr | maskValue

	// apply mask floating values
	addrs := []int{addr}
	for i := 0; i < 36; i++ {
		if maskBits&(1<<i) == 0 {
			newAddrs := []int{}
			for _, x := range addrs {
				newAddrs = append(newAddrs,
					x & ^(1<<i),
					x & ^(1<<i) | (1<<i),
				)
			}
			addrs = newAddrs
		}
	}

	return addrs
}

func decodeMask(line string) (bits int, value int) {
	fields := strings.Fields(line)

	for i, r := range fields[2] {
		if r == '1' {
			value |= 1 << (35 - i)
		}
		if r != 'X' {
			bits |= 1 << (35 - i)
		}
	}

	return
}

func decodeMem(line string) (addr int, val int) {
	fields := strings.Fields(line)
	fmt.Sscanf(fields[0], "mem[%d]", &addr)
	val = util.MustAtoi(fields[2])
	return
}

func sumMemory(mem map[int]int) int {
	var total int
	for _, val := range mem {
		total += val
	}
	return total
}
