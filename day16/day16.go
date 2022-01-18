package day16

import (
	"fmt"
	"log"
	"strconv"
)

func expandHex(hex string) string {
	i, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		log.Fatalln("Error", err)
	}

	out := fmt.Sprintf("%b", i)
	for len(out)%4 != 0 {
		out = "0" + out
	}
	return out
}

func parseDecimal(b string) int {
	i, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		log.Fatalln("Error", err)
	}
	return int(i)
}

func parseLiteralNumber(b string) int {
	i := 0
	acc := ""
	for {
		acc += b[i+1 : i+5]
		if b[i:i+1] == "0" {
			break
		}
		i += 5
	}
	return parseDecimal(acc)
}

func parseVersion(binary string) int {
	return parseDecimal(binary[0:3])
}

func parseId(binary string) int {
	return parseDecimal(binary[3:6])
}

func parseIdLengthSubpackets(binary string) int {
	lengthId := parseDecimal(binary[6:7])
	if lengthId == 0 {
		return 15
	} else if lengthId == 1 {
		return 11
	} else {
		log.Fatal("unknown length ID", lengthId)
		return 0
	}
}

func parseLengthSubpackets(binary string, length int) int {
	return parseDecimal(binary[7 : 7+length])
}
