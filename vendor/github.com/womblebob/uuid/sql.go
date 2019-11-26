// Copyright 2015 Google Inc.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uuid

import (
	"database/sql/driver"
	//"encoding/hex"
	"errors"
	"fmt"
)

var DefaultScanner Scanner = &StandardScanner{}

type Scanner interface {
	Scan(interface{}, *UUID) error
	Value(UUID) (driver.Value, error)
}

// StandardScanner implements uuid.Scanner using the original logic from pborman/uuid
type StandardScanner struct{}

func (s *StandardScanner) Scan(src interface{}, uuid *UUID) error {
	if uuid == nil {
		return errors.New("Nil uuid parsed to standard scanner")
	}

	switch src.(type) {
	case string:
		// if an empty UUID comes from a table, we return a null UUID
		if src.(string) == "" {
			return nil
		}

		// see uuid.Parse for required string format
		parsed := Parse(src.(string))

		if parsed == nil {
			return errors.New("Scan: invalid UUID format")
		}

		*uuid = parsed
	case []byte:
		b := src.([]byte)

		// if an empty UUID comes from a table, we return a null UUID
		if len(b) == 0 {
			return nil
		}

		// assumes a simple slice of bytes if 16 bytes
		// otherwise attempts to parse
		if len(b) == 16 {
			*uuid = UUID(b)
		} else {
			u := Parse(string(b))

			if u == nil {
				return errors.New("Scan: invalid UUID format")
			}

			*uuid = u
		}

	default:
		return fmt.Errorf("Scan: unable to scan type %T into UUID", src)
	}

	return nil
}

func (s *StandardScanner) Value(uuid UUID) (driver.Value, error) {
	return uuid.String(), nil
}

// SqlServerScanner implements uuid.Scanner using the logic taken from
// https://github.com/denisenkom/go-mssqldb/blob/7a8f90684760b9ebd0001efcfc0102fcfb47c5d9/uniqueidentifier.go
type SqlServerScanner struct{}

func (s *SqlServerScanner) Scan(val interface{}, uuid *UUID) error {
	if uuid == nil {
		return errors.New("Nil uuid parsed to sql scanner")
	}

	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}

	switch vt := val.(type) {
	case []byte:
		if len(vt) != 16 {
			return errors.New("mssql: invalid UniqueIdentifier length")
		}

		var raw [16]byte

		copy(raw[:], vt)

		reverse(raw[0:4])
		reverse(raw[4:6])
		reverse(raw[6:8])
		return uuid.UnmarshalBinary(raw[:])

	case string:
		return uuid.UnmarshalText([]byte(vt))
	default:
		return fmt.Errorf("mssql: cannot convert %T to UUID, %u", val, vt)
	}
}

func (s *SqlServerScanner) Value(uuid UUID) (driver.Value, error) {
	reverse := func(b []byte) {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
	}

	var raw [16]byte
	copy(raw[:], uuid[:])

	reverse(raw[0:4])
	reverse(raw[4:6])
	reverse(raw[6:8])

	return raw[:], nil
}

// Scan implements sql.Scanner so UUIDs can be read from databases transparently
// Currently, database types that map to string and []byte are supported. Please
// consult database-specific driver documentation for matching types.
func (uuid *UUID) Scan(src interface{}) error {
	return DefaultScanner.Scan(src, uuid)
}

// Value implements sql.Valuer so that UUIDs can be written to databases
// transparently. Currently, UUIDs map to strings. Please consult
// database-specific driver documentation for matching types.
func (uuid UUID) Value() (driver.Value, error) {
	return DefaultScanner.Value(uuid)
}
