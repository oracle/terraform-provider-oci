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

// TruncateTableFormatEntry The Truncate Table masking format drops all the rows in a table. If one of the
// columns in a table is masked using Truncate Table, the entire table is truncated,
// so no other masking format can be used for any of the other columns in that table.
// If a table is being truncated, it cannot be referred to by a foreign key constraint
// or a dependent column. To learn more, check Truncate Table in the Data Safe documentation.
type TruncateTableFormatEntry struct {

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m TruncateTableFormatEntry) GetDescription() *string {
	return m.Description
}

func (m TruncateTableFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TruncateTableFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TruncateTableFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTruncateTableFormatEntry TruncateTableFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTruncateTableFormatEntry
	}{
		"TRUNCATE_TABLE",
		(MarshalTypeTruncateTableFormatEntry)(m),
	}

	return json.Marshal(&s)
}
