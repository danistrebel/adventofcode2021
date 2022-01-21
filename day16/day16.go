package day16

import (
	"adventofcode/utils"
	"fmt"
	"log"
	"strconv"
)

type packet struct {
	version    int
	literal    int
	id         int
	subPackets []packet
}

func expandHex(hex string) string {
	binary := ""
	for _, c := range string(hex) {
		num, _ := strconv.ParseInt(string([]rune{c}), 16, 64)
		binary += fmt.Sprintf("%04b", num)
	}
	return binary
}

func parseDecimal(b string) int {
	i, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		log.Fatalln("Error", err)
	}
	return int(i)
}

func parseLiteralNumber(b string, l *int) int {
	acc := ""
	for {
		acc += b[*l+1 : *l+5]
		if b[*l:*l+1] == "0" {
			break
		}
		*l += 5
	}
	*l += 5
	return parseDecimal(acc)
}

func parseVersion(binary string, l *int) int {
	v := parseDecimal(binary[*l : *l+3])
	*l += 3
	return v
}

func parseId(binary string, l *int) int {
	id := parseDecimal(binary[*l : *l+3])
	*l += 3
	return id
}

func parseLengthTypeId(binary string, l *int) int {
	typeId := parseDecimal(binary[*l : *l+1])
	*l++
	return typeId
}

func parseSubPacketsLocation(binary string, typeId int, l *int) int {
	if typeId == 0 {
		bitlength := parseDecimal(binary[*l : *l+15])
		*l += 15
		return bitlength
	} else if typeId == 1 {
		packetCount := parseDecimal(binary[*l : *l+11])
		*l += 11
		return packetCount
	} else {
		log.Fatal("unknown length ID", typeId)
		return 0
	}
}

func parsePacket(binary string, l *int) packet {

	version := parseVersion(binary, l)
	id := parseId(binary, l)

	if id == 4 {
		value := parseLiteralNumber(binary, l)
		return packet{literal: value, version: version, id: id}
	} else {
		lengthTypeId := parseLengthTypeId(binary, l)
		subPackets := []packet{}
		if lengthTypeId == 0 {
			subpacketsBit := parseSubPacketsLocation(binary, lengthTypeId, l)
			endPackets := *l + subpacketsBit
			for *l < endPackets {
				subPacket := parsePacket(binary, l)
				subPackets = append(subPackets, subPacket)
			}
			return packet{version: version, subPackets: subPackets, id: id}
		} else {
			subpacketsCount := parseSubPacketsLocation(binary, lengthTypeId, l)

			for c := 0; c < subpacketsCount; c++ {
				subPacket := parsePacket(binary, l)
				subPackets = append(subPackets, subPacket)
			}
			return packet{version: version, subPackets: subPackets, id: id}
		}
	}
}

func sumVersions(hex string) int {
	binary := expandHex(hex)
	location := 0
	rootPacket := parsePacket(binary, &location)
	return rootPacket.sumVersions()
}

func (p packet) sumVersions() int {
	v := p.version
	for _, sp := range p.subPackets {
		v += sp.sumVersions()
	}
	return v
}

func rootPacketValue(hex string) int {
	binary := expandHex(hex)
	location := 0
	rootPacket := parsePacket(binary, &location)
	return rootPacket.packetValue()
}

func (p packet) packetValue() int {
	switch p.id {
	case 0:
		acc := 0
		for _, sp := range p.subPackets {
			acc += sp.packetValue()
		}
		return acc
	case 1:
		acc := 1
		for _, sp := range p.subPackets {
			acc *= sp.packetValue()
		}
		return acc
	case 2:
		acc := p.subPackets[0].packetValue()
		for i := 1; i < len(p.subPackets); i++ {
			v := p.subPackets[i].packetValue()
			if v < acc {
				acc = v
			}
		}
		return acc
	case 3:
		acc := p.subPackets[0].packetValue()
		for i := 1; i < len(p.subPackets); i++ {
			v := p.subPackets[i].packetValue()
			if v > acc {
				acc = v
			}
		}
		return acc
	case 4:
		return p.literal
	case 5:
		if p.subPackets[0].packetValue() > p.subPackets[1].packetValue() {
			return 1
		} else {
			return 0
		}
	case 6:
		if p.subPackets[0].packetValue() < p.subPackets[1].packetValue() {
			return 1
		} else {
			return 0
		}
	case 7:
		if p.subPackets[0].packetValue() == p.subPackets[1].packetValue() {
			return 1
		} else {
			return 0
		}
	}

	log.Fatal("unnkown id", p.id)
	return 0
}

func PrintSolution() {
	lines := utils.ParseLines("inputs/day16.txt")
	versionSum := sumVersions(lines[0])
	fmt.Println("Sum Versions (Part 1)", versionSum)
	rootPacketValue := rootPacketValue(lines[0])
	fmt.Println("Root Packet Value (Part 2)", rootPacketValue)
}
