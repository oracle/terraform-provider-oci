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

// SubstringFormatEntry The Substring masking format extracts a portion of the original column
// value and uses it to replace the original value. It internally uses the
// Oracle SUBSTR function. It takes the start position and length as input,
// extracts substring from the original value using SUBSTR, and uses the
// substring to replace the original value. To learn more, check Substring
// in the Data Safe documentation.
type SubstringFormatEntry struct {

	// The starting position in the original string from where the substring
	// should be extracted. It can be either a positive or a negative integer.
	// If It's negative, the counting starts from the end of the string.
	StartPosition *int `mandatory:"true" json:"startPosition"`

	// The number of characters that should be there in the substring. It should
	// be an integer and greater than zero.
	Length *int `mandatory:"true" json:"length"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m SubstringFormatEntry) GetDescription() *string {
	return m.Description
}

func (m SubstringFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubstringFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SubstringFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSubstringFormatEntry SubstringFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSubstringFormatEntry
	}{
		"SUBSTRING",
		(MarshalTypeSubstringFormatEntry)(m),
	}

	return json.Marshal(&s)
}
