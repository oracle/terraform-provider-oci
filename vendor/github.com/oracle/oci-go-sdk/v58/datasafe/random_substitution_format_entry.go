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

// RandomSubstitutionFormatEntry The Random Substitution masking format uses the specified substitution column
// as the source of masked values. The values in the substitution column are randomly
// ordered before mapping them to the original column values. As a masking operation
// renames tables temporarily, the substitution column must be in a table that has
// no masking column. Also, you may want to ensure that the substitution column has
// sufficient values to uniquely mask the target column.
// Unlike Deterministic Substitution, Random Substitution doesn't do deterministic
// masking, and thus, doesn't require a seed value. To learn more, check Random
// Substitution in the Data Safe documentation.
type RandomSubstitutionFormatEntry struct {

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
func (m RandomSubstitutionFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RandomSubstitutionFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RandomSubstitutionFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RandomSubstitutionFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRandomSubstitutionFormatEntry RandomSubstitutionFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRandomSubstitutionFormatEntry
	}{
		"RANDOM_SUBSTITUTION",
		(MarshalTypeRandomSubstitutionFormatEntry)(m),
	}

	return json.Marshal(&s)
}
