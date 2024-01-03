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

// RandomDateFormatEntry The Random Date masking format generates random and unique dates within a range.
// The date range is defined by the startDate and endDate attributes. The start date
// must be less than or equal to the end date. When masking columns with uniqueness
// constraint, ensure that the date range is sufficient enough to generate unique
// values. To learn more, check Random Date in the Data Safe documentation.
type RandomDateFormatEntry struct {

	// The lower bound of the range within which random dates should be generated.
	// The start date must be less than or equal to the end date.
	StartDate *common.SDKTime `mandatory:"true" json:"startDate"`

	// The upper bound of the range within which random dates should be generated.
	// The end date must be greater than or equal to the start date.
	EndDate *common.SDKTime `mandatory:"true" json:"endDate"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

// GetDescription returns Description
func (m RandomDateFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomDateFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomDateFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomDateFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomDateFormatEntry RandomDateFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomDateFormatEntry
	}{
		"RANDOM_DATE",
		(MarshalTypeRandomDateFormatEntry)(m),
	}

	return json.Marshal(&s)
}
