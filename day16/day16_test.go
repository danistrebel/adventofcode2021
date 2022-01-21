package day16

import (
	"testing"
)

func TestLiteralInput(t *testing.T) {

	binary := expandHex("D2FE28")
	expectedBinary := "110100101111111000101000"
	if binary != expectedBinary {
		t.Error("unexpected binary conversion. Got", binary, "expected", expectedBinary)
	}

	l := 0

	version := parseVersion(binary, &l)
	expectedVersion := 6
	if version != expectedVersion {
		t.Error("Unexpected Version. Got", version, "expected", expectedVersion)
	}

	id := parseId(binary, &l)
	expectedId := 4
	if id != expectedId {
		t.Error("Unexpected ID. Got", id, "expected", expectedId)
	}

	literalNumber := parseLiteralNumber(binary, &l)
	expectedLiteralNumber := 2021
	if literalNumber != expectedLiteralNumber {
		t.Error("Unexpected literal number. Got", literalNumber, "expected", expectedLiteralNumber)
	}
}

func TestOperatorLength(t *testing.T) {
	binary := expandHex("38006F45291200")

	expectedBinary := "00111000000000000110111101000101001010010001001000000000"
	if binary != expectedBinary {
		t.Error("unexpected binary conversion. Got", binary, "expected", expectedBinary)
	}
	l := 0
	version := parseVersion(binary, &l)
	expectedVersion := 1
	if version != expectedVersion {
		t.Error("Unexpected Version. Got", version, "expected", expectedVersion)
	}

	id := parseId(binary, &l)
	expectedId := 6
	if id != expectedId {
		t.Error("Unexpected ID. Got", id, "expected", expectedId)
	}

	typeId := parseLengthTypeId(binary, &l)
	expectedTypeId := 0
	if typeId != expectedTypeId {
		t.Error("Unexpected ID. Got", typeId, "expected", expectedTypeId)
	}
}

func TestOperatorCount(t *testing.T) {
	binary := expandHex("EE00D40C823060")

	expectedBinary := "11101110000000001101010000001100100000100011000001100000"
	if binary != expectedBinary {
		t.Error("unexpected binary conversion. Got", binary, "expected", expectedBinary)
	}

	l := 0

	version := parseVersion(binary, &l)
	expectedVersion := 7
	if version != expectedVersion {
		t.Error("Unexpected Version. Got", version, "expected", expectedVersion)
	}

	id := parseId(binary, &l)
	expectedId := 3
	if id != expectedId {
		t.Error("Unexpected ID. Got", id, "expected", expectedId)
	}

	typeId := parseLengthTypeId(binary, &l)
	expectedTypeId := 1
	if typeId != expectedTypeId {
		t.Error("Unexpected ID. Got", typeId, "expected", expectedTypeId)
	}

	subpacketsCount := parseSubPacketsLocation(binary, typeId, &l)
	expectedSubPacketsCount := 3
	if subpacketsCount != expectedSubPacketsCount {
		t.Error("Unexpected Subpackets Count. Got", subpacketsCount, "expected", expectedSubPacketsCount)
	}
}

func TestSumVersions(t *testing.T) {
	testInputs := []struct {
		hex      string
		expected int
	}{
		{"8A004A801A8002F478", 16},
		{"620080001611562C8802118E34", 12},
		{"C0015000016115A2E0802F182340", 23},
		{"A0016C880162017C3686B18A3D4780", 31},
	}

	for _, testInput := range testInputs {
		sum := sumVersions(testInput.hex)
		if sum != testInput.expected {
			t.Error("Unexpected Version Sum. Got", sum, "expected", testInput.expected)
		}
	}
}

func TestPacketValue(t *testing.T) {
	testInputs := []struct {
		hex      string
		expected int
	}{
		{"C200B40A82", 3},
		{"04005AC33890", 54},
		{"880086C3E88112", 7},
		{"CE00C43D881120", 9},
		{"D8005AC2A8F0", 1},
		{"F600BC2D8F", 0},
		{"9C005AC2F8F0", 0},
		{"9C0141080250320F1802104A08", 1},
	}

	for _, testInput := range testInputs {
		sum := rootPacketValue(testInput.hex)
		if sum != testInput.expected {
			t.Error("Unexpected Version Sum. Got", sum, "expected", testInput.expected)
		}
	}
}
