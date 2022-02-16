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

// RandomListFormatEntry The Random List masking format randomly selects values from a list of values
// to replace the original values. To learn more, check Random List in the
// Data Safe documentation.
type RandomListFormatEntry struct {

	// A comma-separated list of values to be used to replace column values.
	// The list can be of strings, numbers, or dates. The data type of each
	// value in the list must be compatible with the data type of the column.
	// The number of entries in the list cannot be more than 999.
	RandomList []string `mandatory:"true" json:"randomList"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m RandomListFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomListFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomListFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomListFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomListFormatEntry RandomListFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomListFormatEntry
	}{
		"RANDOM_LIST",
		(MarshalTypeRandomListFormatEntry)(m),
	}

	return json.Marshal(&s)
}
