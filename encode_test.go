package main

import (
	"github.com/ajg/form"
	"testing"
)

func TestMeasurementProtocolEncode(T *testing.T) {
	mp := &MeasurementProtocol{
		V:   1,
		Tid: "TestTid",
		Cid: "TestCid",
	}

	if err := mp.Validate(); err != nil {
		T.Fatal(err.Error())
	}

	if str, err := form.EncodeToValues(mp); err != nil {
		T.Fatal("Error:", err.Error())
	} else {
		T.Log("Decoded MP is:", str)
	}

}
