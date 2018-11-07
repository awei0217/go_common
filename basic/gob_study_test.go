package basic

import "testing"

func TestEncoderAndDecoder(t *testing.T) {
	EncoderAndDecoder(P{
		X: 10,
		Y: "spw",
	}, Q{})
}
