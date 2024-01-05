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

// FixedNumberFormatEntry The Fixed Number masking format uses a constant number for masking. To learn more,
// check Fixed Number in the Data Safe documentation.
type FixedNumberFormatEntry struct {

	// The constant number to be used for masking.
	FixedNumber *float32 `mandatory:"true" json:"fixedNumber"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m FixedNumberFormatEntry) GetDescription() *string {
	return m.Description
}

func (m FixedNumberFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FixedNumberFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FixedNumberFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFixedNumberFormatEntry FixedNumberFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeFixedNumberFormatEntry
	}{
		"FIXED_NUMBER",
		(MarshalTypeFixedNumberFormatEntry)(m),
	}

	return json.Marshal(&s)
}
