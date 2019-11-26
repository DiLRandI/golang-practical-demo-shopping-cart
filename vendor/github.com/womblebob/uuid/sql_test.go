// Copyright 2015 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"bytes"
	"strings"
	"testing"
)

func TestStandardScannerScan(t *testing.T) {
	var stringTest string = "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	var byteTest []byte = Parse(stringTest)
	var badTypeTest int = 6
	var invalidTest string = "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	sut := &StandardScanner{}

	// sunny day tests

	var uuid UUID
	err := sut.Scan(stringTest, &uuid)
	if err != nil {
		t.Fatal(err)
	}

	err = sut.Scan([]byte(stringTest), &uuid)
	if err != nil {
		t.Fatal(err)
	}

	err = sut.Scan(byteTest, &uuid)
	if err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err = sut.Scan(badTypeTest, &uuid)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}
	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = sut.Scan(invalidTest, &uuid)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid UUID returned an incorrect error message")
	}

	err = sut.Scan(byteTest[:len(byteTest)-2], &uuid)
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid byte UUID returned an incorrect error message")
	}

	// empty tests

	uuid = nil
	var emptySlice []byte
	err = sut.Scan(emptySlice, &uuid)
	if err != nil {
		t.Fatal(err)
	}

	if uuid != nil {
		t.Error("UUID was not nil after scanning empty byte slice")
	}

	uuid = nil
	var emptyString string
	err = sut.Scan(emptyString, &uuid)
	if err != nil {
		t.Fatal(err)
	}

	if uuid != nil {
		t.Error("UUID was not nil after scanning empty string")
	}
}

func TestStandardScannerValue(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	uuid := Parse(stringTest)

	sut := &StandardScanner{}
	val, _ := sut.Value(uuid)
	if val != stringTest {
		t.Error("Value() did not return expected string")
	}
}

func TestSqlServerScanner(t *testing.T) {
	dbUUID := []byte{0x67, 0x45, 0x23, 0x01,
		0xAB, 0x89,
		0xEF, 0xCD,
		0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF,
	}

	uuid := UUID{0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF, 0x01, 0x23, 0x45, 0x67, 0x89, 0xAB, 0xCD, 0xEF}

	t.Run("Scan", func(t *testing.T) {
		t.Run("[]byte", func(t *testing.T) {
			sut := &SqlServerScanner{}
			got := UUID{}
			if err := sut.Scan(dbUUID[:], &got); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, uuid) {
				t.Errorf("bytes not swapped correctly: got %q; want %q", got, uuid)
			}
		})

		t.Run("string", func(t *testing.T) {
			sut := &SqlServerScanner{}
			var got UUID
			if err := sut.Scan(uuid.String(), &got); err != nil {
				t.Fatal(err)
			}
			if !bytes.Equal(got, uuid) {
				t.Errorf("string not scanned correctly: got %q; want %q", got, uuid)
			}
		})
	})

	t.Run("Value", func(t *testing.T) {
		sut := &SqlServerScanner{}
		got := uuid
		v, err := sut.Value(got)
		if err != nil {
			t.Fatal(err)
		}

		b, ok := v.([]byte)
		if !ok {
			t.Fatalf("(%T) is not []byte", v)
		}

		if !bytes.Equal(b, dbUUID[:]) {
			t.Errorf("got %q; want %q", b, dbUUID)
		}
	})
}

func TestScan(t *testing.T) {
	var stringTest string = "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	var byteTest []byte = Parse(stringTest)
	var badTypeTest int = 6
	var invalidTest string = "f47ac10b-58cc-0372-8567-0e02b2c3d4"

	// sunny day tests

	var uuid UUID
	err := (&uuid).Scan(stringTest)
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan([]byte(stringTest))
	if err != nil {
		t.Fatal(err)
	}

	err = (&uuid).Scan(byteTest)
	if err != nil {
		t.Fatal(err)
	}

	// bad type tests

	err = (&uuid).Scan(badTypeTest)
	if err == nil {
		t.Error("int correctly parsed and shouldn't have")
	}
	if !strings.Contains(err.Error(), "unable to scan type") {
		t.Error("attempting to parse an int returned an incorrect error message")
	}

	// invalid/incomplete uuids

	err = (&uuid).Scan(invalidTest)
	if err == nil {
		t.Error("invalid uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid UUID returned an incorrect error message")
	}

	err = (&uuid).Scan(byteTest[:len(byteTest)-2])
	if err == nil {
		t.Error("invalid byte uuid was parsed without error")
	}
	if !strings.Contains(err.Error(), "invalid UUID") {
		t.Error("attempting to parse an invalid byte UUID returned an incorrect error message")
	}

	// empty tests

	uuid = nil
	var emptySlice []byte
	err = (&uuid).Scan(emptySlice)
	if err != nil {
		t.Fatal(err)
	}

	if uuid != nil {
		t.Error("UUID was not nil after scanning empty byte slice")
	}

	uuid = nil
	var emptyString string
	err = (&uuid).Scan(emptyString)
	if err != nil {
		t.Fatal(err)
	}

	if uuid != nil {
		t.Error("UUID was not nil after scanning empty string")
	}
}

func TestValue(t *testing.T) {
	stringTest := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	uuid := Parse(stringTest)
	val, _ := uuid.Value()
	if val != stringTest {
		t.Error("Value() did not return expected string")
	}
}
