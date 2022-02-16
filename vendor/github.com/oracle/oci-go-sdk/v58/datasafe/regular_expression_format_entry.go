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

// RegularExpressionFormatEntry The Regular Expression masking format gives the flexibility to use regular
// expressions to search for sensitive data in a column of Large Object data
// type (LOB), and replace the data with a fixed string, fixed number, null
// value, or SQL expression. It can also be used for columns of VARCHAR2 type
// to mask parts of strings. To learn more, check Regular Expressions in the
// Data Safe documentation.
type RegularExpressionFormatEntry struct {

	// The pattern that should be used to search for data.
	RegularExpression *string `mandatory:"true" json:"regularExpression"`

	// The value that should be used to replace the data matching the regular
	// expression. It can be a fixed string, fixed number, null value, or
	// SQL expression.
	ReplaceWith *string `mandatory:"true" json:"replaceWith"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m RegularExpressionFormatEntry) GetDescription() *string {
	return m.Description
}

func (m RegularExpressionFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RegularExpressionFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m RegularExpressionFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeRegularExpressionFormatEntry RegularExpressionFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeRegularExpressionFormatEntry
	}{
		"REGULAR_EXPRESSION",
		(MarshalTypeRegularExpressionFormatEntry)(m),
	}

	return json.Marshal(&s)
}
