package gpt

import (
	"fmt"
	"testing"
)

func Test_ShouldReadHeader(t *testing.T) {
	testfileOk := "gpt_test_hd"
	expectedUUID := "d0f2f537-fdf7-41ad-acb9-42f0aeb3781d"

	table, err := ReadFrom(testfileOk)

	if err != nil {
		t.Errorf("ReadFrom(%q) threw error while reading testfile", testfileOk)
	}

	actualUUID := fmt.Sprintf("%s", table.Header.DiskGUID.AsUUID())
	if actualUUID != expectedUUID {
		t.Errorf("ReadFrom(%q): Actual %s != expected %s", testfileOk, actualUUID, expectedUUID)
	}
}

func Test_ShouldReadFirstPartition(t *testing.T) {
	testfileOk := "gpt_test_hd"

	var expectedFirstLBA uint64 = 34
	var expectedLastLBA uint64 = 1000

	table, err := ReadFrom(testfileOk)

	if err != nil {
		t.Errorf("ReadFrom(%q) threw error while reading testfile", testfileOk)
	}

	actualFirstLBA := table.Entries[0].FirstLBA
	actualLastLBA := table.Entries[0].LastLBA

	if expectedFirstLBA != actualFirstLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedFirstLBA, actualFirstLBA)
	}

	if expectedLastLBA != actualLastLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedLastLBA, actualLastLBA)
	}
}

func Test_ShouldReadSecondPartition(t *testing.T) {
	testfileOk := "gpt_test_hd"

	var expectedFirstLBA uint64 = 1001
	var expectedLastLBA uint64 = 1500

	table, err := ReadFrom(testfileOk)

	if err != nil {
		t.Errorf("ReadFrom(%q) threw error while reading testfile", testfileOk)
	}

	actualFirstLBA := table.Entries[1].FirstLBA
	actualLastLBA := table.Entries[1].LastLBA

	if expectedFirstLBA != actualFirstLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedFirstLBA, actualFirstLBA)
	}

	if expectedLastLBA != actualLastLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedLastLBA, actualLastLBA)
	}
}

// TODO: check stuff like GUIDs and so on for all partitions

func Test_ShouldReadThirdPartition(t *testing.T) {
	testfileOk := "gpt_test_hd"

	var expectedFirstLBA uint64 = 1501
	var expectedLastLBA uint64 = 2014

	table, err := ReadFrom(testfileOk)

	if err != nil {
		t.Errorf("ReadFrom(%q) threw error while reading testfile", testfileOk)
	}

	actualFirstLBA := table.Entries[127].FirstLBA
	actualLastLBA := table.Entries[127].LastLBA

	if expectedFirstLBA != actualFirstLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedFirstLBA, actualFirstLBA)
	}

	if expectedLastLBA != actualLastLBA {
		t.Errorf("ReadFrom(%q): Actual %d != %d", testfileOk, expectedLastLBA, actualLastLBA)
	}
}

func Test_ShouldBreakOnBrokenEFISignature(t *testing.T) {
	testfile := "gpt_test_hd_brokenEfiSignature"
	_, err := ReadFrom(testfile)

	if err == nil {
		t.Errorf("gpt.ReadHeader(%q) should throw an error", testfile)
	}
}

func Test_ShouldBreakOnNonexistantFile(t *testing.T) {
	nonexistantfile := "/nonexistant/file"
	_, err := ReadFrom(nonexistantfile)

	if err == nil {
		t.Errorf("ReadHeader(%q) should throw an error", nonexistantfile)
	}
}
