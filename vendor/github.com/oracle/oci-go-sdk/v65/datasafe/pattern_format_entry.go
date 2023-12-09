// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// PatternFormatEntry The Pattern masking format randomly selects values according to pattern
// to replace the original values.
// Rules
// Max Generated Data Length 30 characters
// Use '%c' for a random lowercase letter
// Use '%C' for a random uppercase letter
// Use '%u[]' for a random character out of all characters enclosed in []
// Use '%%' for a '%'
// Use '%d' for a random digit
// Use '%nd','%nc', '%nC', or '%nu[]' n random letters or digits or characters enclosed in [], n can be 0-9 only
// Any other character will be included as it is
// Examples
// %3d-%5C will generate 416-JQPCS
// %3d-%5c will generate 392-dehco
// %u[$^#] will generate $
// %%%3d will generate %704
type PatternFormatEntry struct {

	// The pattern that should be used to mask data.
	Pattern *string `mandatory:"true" json:"pattern"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m PatternFormatEntry) GetDescription() *string {
	return m.Description
}

func (m PatternFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PatternFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m PatternFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypePatternFormatEntry PatternFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypePatternFormatEntry
	}{
		"PATTERN",
		(MarshalTypePatternFormatEntry)(m),
	}

	return json.Marshal(&s)
}
