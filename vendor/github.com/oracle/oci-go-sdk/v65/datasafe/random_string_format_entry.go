// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RandomStringFormatEntry The Random String masking format generates random and unique strings of length
// within a range. The length range is defined by the startLength and endLength
// attributes. The start length must be less than or equal to the end length. When
// masking columns with uniqueness constraint, ensure that the length range is
// sufficient enough to generate unique values. To learn more, check Random String
// in the Data Safe documentation.
type RandomStringFormatEntry struct {

	// The minimum number of characters the generated strings should have. It can
	// be any integer greater than zero, but it must be less than or equal to the
	// end length.
	StartLength *int `mandatory:"true" json:"startLength"`

	// The maximum number of characters the generated strings should have. It can
	// be any integer greater than zero, but it must be greater than or equal to
	// the start length.
	EndLength *int `mandatory:"true" json:"endLength"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m RandomStringFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomStringFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomStringFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomStringFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomStringFormatEntry RandomStringFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomStringFormatEntry
	}{
		"RANDOM_STRING",
		(MarshalTypeRandomStringFormatEntry)(m),
	}

	return json.Marshal(&s)
}
