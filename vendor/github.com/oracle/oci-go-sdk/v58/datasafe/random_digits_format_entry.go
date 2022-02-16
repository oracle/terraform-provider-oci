// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// RandomDigitsFormatEntry The Random Digits masking format generates random digits of length within a range.
// The length range is defined by the startLength and endLength attributes. The start
// length must be less than or equal to the end length. When masking columns with
// uniqueness constraint, ensure that the length range is sufficient enough to generate
// unique values. This masking format pads to the appropriate length in a string, but
// does not pad when used for a number column. It's a complementary type of Random Number,
// which is not padded.
type RandomDigitsFormatEntry struct {

	// The minimum number of digits the generated values should have. It can be
	// any integer greater than zero, but it must be less than or equal to the
	// end length.
	StartLength *int `mandatory:"true" json:"startLength"`

	// The maximum number of digits the generated values should have. It can
	// be any integer greater than zero, but it must be greater than or equal
	// to the start length.
	EndLength *int `mandatory:"true" json:"endLength"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m RandomDigitsFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomDigitsFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomDigitsFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomDigitsFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomDigitsFormatEntry RandomDigitsFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomDigitsFormatEntry
	}{
		"RANDOM_DIGITS",
		(MarshalTypeRandomDigitsFormatEntry)(m),
	}

	return json.Marshal(&s)
}
