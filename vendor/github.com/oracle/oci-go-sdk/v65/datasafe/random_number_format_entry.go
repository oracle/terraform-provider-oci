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

// RandomNumberFormatEntry The Random Number masking format generates random and unique integers within
// a range. The range is defined by the startValue and endValue attributes. The
// start value must be less than or equal to the end value. When masking columns
// with uniqueness constraint, ensure that the range is sufficient enough to
// generate unique values. To learn more, check Random Number in the Data Safe
// documentation.
type RandomNumberFormatEntry struct {

	// The lower bound of the range within which random numbers should be
	// generated. It must be less than or equal to the end value. It
	// supports input of long type.
	StartValue *int64 `mandatory:"true" json:"startValue"`

	// The upper bound of the range within which random numbers should be
	// generated. It must be greater than or equal to the start value.
	// It supports input of long type.
	EndValue *int64 `mandatory:"true" json:"endValue"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m RandomNumberFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomNumberFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomNumberFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomNumberFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomNumberFormatEntry RandomNumberFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomNumberFormatEntry
	}{
		"RANDOM_NUMBER",
		(MarshalTypeRandomNumberFormatEntry)(m),
	}

	return json.Marshal(&s)
}
