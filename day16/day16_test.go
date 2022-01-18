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

	version := parseVersion(binary)
	expectedVersion := 6
	if version != expectedVersion {
		t.Error("Unexpected Version. Got", version, "expected", expectedVersion)
	}

	id := parseId(binary)
	expectedId := 4
	if id != expectedId {
		t.Error("Unexpected ID. Got", id, "expected", expectedId)
	}

	literalNumber := parseLiteralNumber(binary[6:])
	expectedLiteralNumber := 2021
	if literalNumber != expectedLiteralNumber {
		t.Error("Unexpected literal number. Got", literalNumber, "expected", expectedLiteralNumber)
	}

}

func TestOperator(t *testing.T) {
	binary := expandHex("38006F45291200")

	expectedBinary := "00111000000000000110111101000101001010010001001000000000"
	if binary != expectedBinary {
		t.Error("unexpected binary conversion. Got", binary, "expected", expectedBinary)
	}

	version := parseVersion(binary)
	expectedVersion := 1
	if version != expectedVersion {
		t.Error("Unexpected Version. Got", version, "expected", expectedVersion)
	}

	id := parseId(binary)
	expectedId := 6
	if id != expectedId {
		t.Error("Unexpected ID. Got", id, "expected", expectedId)
	}

	subpacketsIdLength := parseIdLengthSubpackets(binary)
	expectedSubpackets := 15
	if subpacketsIdLength != expectedSubpackets {
		t.Error("Unexpected ID Sub Packets Length. Got", subpacketsIdLength, "expected", expectedSubpackets)
	}

	subpacketsLength := parseLengthSubpackets(binary, subpacketsIdLength)
	expectedLength := 27
	if subpacketsLength != expectedLength {
		t.Error("Unexpected ID Sub Packets Length. Got", subpacketsLength, "expected", expectedLength)
	}
}
