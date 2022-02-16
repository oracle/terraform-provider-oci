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

// DeleteRowsFormatEntry The Delete Rows masking format deletes the rows that meet a user-specified
// condition. It is useful in conditional masking when you want to delete a
// subset of values in a column and mask the remaining values using some other
// masking formats. You should be careful while using this masking format. If
// no condition is specified, all rows in a table are deleted. If a column is
// being masked using Delete Rows, there must not be a foreign key constraint
// or dependent column referring to the table. To learn more, check Delete Rows
// in the Data Safe documentation.
type DeleteRowsFormatEntry struct {

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m DeleteRowsFormatEntry) GetDescription() *string {
	return m.Description
}

func (m DeleteRowsFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeleteRowsFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeleteRowsFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeleteRowsFormatEntry DeleteRowsFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDeleteRowsFormatEntry
	}{
		"DELETE_ROWS",
		(MarshalTypeDeleteRowsFormatEntry)(m),
	}

	return json.Marshal(&s)
}
