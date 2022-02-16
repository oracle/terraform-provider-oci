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

// NullValueFormatEntry The Null Value masking format replaces column data with NULL. The column
// being masked must be allowed to contain null values. To learn more,
// check Null Value in the Data Safe documentation.
type NullValueFormatEntry struct {

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m NullValueFormatEntry) GetDescription() *string {
	return m.Description
}

func (m NullValueFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NullValueFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NullValueFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNullValueFormatEntry NullValueFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNullValueFormatEntry
	}{
		"NULL_VALUE",
		(MarshalTypeNullValueFormatEntry)(m),
	}

	return json.Marshal(&s)
}
