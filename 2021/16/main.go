package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

var hex2bin = map[string][]string{
	"0": {"0", "0", "0", "0"},
	"1": {"0", "0", "0", "1"},
	"2": {"0", "0", "1", "0"},
	"3": {"0", "0", "1", "1"},
	"4": {"0", "1", "0", "0"},
	"5": {"0", "1", "0", "1"},
	"6": {"0", "1", "1", "0"},
	"7": {"0", "1", "1", "1"},
	"8": {"1", "0", "0", "0"},
	"9": {"1", "0", "0", "1"},
	"A": {"1", "0", "1", "0"},
	"B": {"1", "0", "1", "1"},
	"C": {"1", "1", "0", "0"},
	"D": {"1", "1", "0", "1"},
	"E": {"1", "1", "1", "0"},
	"F": {"1", "1", "1", "1"},
}

type packet struct {
	version int
	typeId  int

	literal int
	length  int

	packets []*packet
}

func (p *packet) toString(level int) string {
	str := ""
	intend := ""
	for i := 0; i < level; i++ {
		intend += "  "
	}
	str += intend
	if p.typeId == 4 {
		str += fmt.Sprintf("<packet version=%d type=%d literal=%d>\n", p.version, p.typeId, p.literal)
	} else {
		str += fmt.Sprintf("<packet version %d type=%d packets=%d>\n", p.version, p.typeId, len(p.packets))
	}
	for _, sp := range p.packets {
		str += sp.toString(level + 1)
	}
	return str
}

func main() {
	input := readInput("./input.txt")

	bits := hexToBits(strings.Split(input[0], ""))
	p := parsePacket(bits)
	fmt.Println(p.toString(0))

	solveOne(p)
	solveTwo(p)
}

func solveOne(p *packet) {
	sum := sumUpVersions(p)
	fmt.Printf("Part 1: Sum of version numbers %d\n", sum)
}

func solveTwo(p *packet) {
	result := interpretPacket(p)
	fmt.Printf("Part 2: Hexadecimal-encoded BITS trasmitted %d\n", result)
}

func sumUpVersions(p *packet) int {
	sum := p.version
	for _, sp := range p.packets {
		sum += sumUpVersions(sp)
	}
	return sum
}

func parsePacket(bits []string) *packet {
	p := packet{
		version: bitsToInt(bits[0:3]),
		typeId:  bitsToInt(bits[3:6]),
	}

	switch p.typeId {
	case 4:
		// literal value
		offset := 6
		literal := []string{}
		for {
			literal = append(literal, bits[offset+1:offset+5]...)
			if bits[offset] == "0" {
				break
			}
			offset += 5
		}
		p.literal = bitsToInt(literal)
		p.length = offset + 5
	default:
		// operator
		lengthTypeId := bits[6]
		if lengthTypeId == "0" {
			totalLength := bitsToInt(bits[7:22])
			offset := 22
			for offset < totalLength+22 {
				sp := parsePacket(bits[offset:])
				p.packets = append(p.packets, sp)
				offset += sp.length
			}
			p.length = offset
		} else {
			subpackets := bitsToInt(bits[7:18])
			offset := 18
			for i := 0; i < subpackets; i++ {
				sp := parsePacket(bits[offset:])
				p.packets = append(p.packets, sp)
				offset += sp.length
			}
			p.length = offset
		}
	}

	return &p
}

func interpretPacket(p *packet) int {
	var result int
	switch p.typeId {
	case 0:
		// sum
		result = 0
		for _, sp := range p.packets {
			result += interpretPacket(sp)
		}
	case 1:
		// product
		result = 1
		for _, sp := range p.packets {
			result *= interpretPacket(sp)
		}
	case 2:
		// min
		result = math.MaxInt
		for _, sp := range p.packets {
			min := interpretPacket(sp)
			if min < result {
				result = min
			}
		}
	case 3:
		// max
		result = 0
		for _, sp := range p.packets {
			max := interpretPacket(sp)
			if max > result {
				result = max
			}
		}
	case 4:
		// literal
		result = p.literal
	case 5:
		// greater than
		sp1 := interpretPacket(p.packets[0])
		sp2 := interpretPacket(p.packets[1])
		if sp1 > sp2 {
			result = 1
		} else {
			result = 0
		}
	case 6:
		// less than
		sp1 := interpretPacket(p.packets[0])
		sp2 := interpretPacket(p.packets[1])
		if sp1 < sp2 {
			result = 1
		} else {
			result = 0
		}
	case 7:
		// equal
		sp1 := interpretPacket(p.packets[0])
		sp2 := interpretPacket(p.packets[1])
		if sp1 == sp2 {
			result = 1
		} else {
			result = 0
		}
	}

	return result
}

func hexToBits(hex []string) []string {
	bin := []string{}
	for _, h := range hex {
		b := hex2bin[string(h)]
		bin = append(bin, b...)
	}
	return bin
}

func bitsToInt(bits []string) int {
	i, _ := strconv.ParseInt(strings.Join(bits, ""), 2, 64)
	return int(i)
}

func readInput(fileName string) []string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	input := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return input
}
