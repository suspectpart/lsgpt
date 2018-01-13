package gpt

import (
	"fmt"
	"testing"

	"github.com/suspectpart/lsgpt/gpt"
)

func TestReadHeader(t *testing.T) {
	testfileOk := "gpt_test_hdok"
	expectedUUID := "17b49718-e953-4465-8eba-0b6ca15d0ebd"

	header, err := gpt.ReadHeader(testfileOk)

	if err != nil {
		t.Errorf("gpt.ReadHeader(%q) threw error while reading testfile", testfileOk)
	}

	actualUUID := fmt.Sprintf("%s", header.DiskGUID.AsUUID())
	if actualUUID != expectedUUID {
		t.Errorf("gpt.ReadHeader(%q): Actual %s != expected %s", testfileOk, actualUUID, expectedUUID)
	}
}
