package main

import (
	"github.com/ajg/form"
	"strings"
	"testing"
)

func TestMeasurementProtocolDecode(T *testing.T) {
	var mp *MeasurementProtocol
	params := strings.NewReader("v=1&tid=testtid&cid=testcid")

	d := form.NewDecoder(params)
	if err := d.Decode(&mp); err != nil {
		T.Fatal(err.Error())
	}

	T.Logf("Decoded: %#v", mp)

	if err := mp.Validate(); err != nil {
		T.Fatal(err.Error())
	}
}
