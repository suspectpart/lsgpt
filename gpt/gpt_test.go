package gpt_test

import (
	"fmt"
	"testing"

	"github.com/suspectpart/lsgpt/gpt"
)

func Test_ShouldReadHeader(t *testing.T) {
	testfileOk := "gpt_test_hd"
	expectedUUID := "17b49718-e953-4465-8eba-0b6ca15d0ebd"

	table, err := gpt.ReadFrom(testfileOk)

	if err != nil {
		t.Errorf("gpt.ReadHeader(%q) threw error while reading testfile", testfileOk)
	}

	actualUUID := fmt.Sprintf("%s", table.Header.DiskGUID.AsUUID())
	if actualUUID != expectedUUID {
		t.Errorf("gpt.ReadHeader(%q): Actual %s != expected %s", testfileOk, actualUUID, expectedUUID)
	}
}

func Test_ShouldBreakOnBrokenEFISignature(t *testing.T) {
	testfile := "gpt_test_hd_brokenEfiSignature"
	_, err := gpt.ReadFrom(testfile)

	if err == nil {
		t.Errorf("gpt.ReadHeader(%q) should throw an error", testfile)
	}
}

func Test_ShouldBreakOnNonexistantFile(t *testing.T) {
	nonexistantfile := "/nonexistant/file"
	_, err := gpt.ReadFrom(nonexistantfile)

	if err == nil {
		t.Errorf("gpt.ReadHeader(%q) should throw an error", nonexistantfile)
	}
}
