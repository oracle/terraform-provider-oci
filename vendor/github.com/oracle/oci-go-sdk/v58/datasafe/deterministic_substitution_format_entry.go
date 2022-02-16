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

// DeterministicSubstitutionFormatEntry The Deterministic Substitution masking format uses the specified substitution column
// as the source of masked values. It performs hash-based substitution to replace the
// original data in a column with values from the substitution column. As a masking
// operation renames tables temporarily, the substitution column must be in a table
// that has no masking column. Also, you may want to ensure that the substitution column
// has sufficient values to uniquely mask the target column.
// Deterministic Substitution requires a seed value while submitting a masking work
// request. Passing the same seed value when masking multiple times or masking different
// databases ensures that the data is masked deterministically. To learn more, check
// Deterministic Substitution in the Data Safe documentation.
type DeterministicSubstitutionFormatEntry struct {

	// The name of the schema that contains the substitution column.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// The name of the table that contains the substitution column.
	TableName *string `mandatory:"true" json:"tableName"`

	// The name of the substitution column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m DeterministicSubstitutionFormatEntry) GetDescription() *string {
	return m.Description
}

func (m DeterministicSubstitutionFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DeterministicSubstitutionFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DeterministicSubstitutionFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDeterministicSubstitutionFormatEntry DeterministicSubstitutionFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeDeterministicSubstitutionFormatEntry
	}{
		"DETERMINISTIC_SUBSTITUTION",
		(MarshalTypeDeterministicSubstitutionFormatEntry)(m),
	}

	return json.Marshal(&s)
}
