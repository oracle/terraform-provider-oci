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

// SqlExpressionFormatEntry The SQL Expression masking format uses a SQL expression to generate values
// that are used to replace the original data values. SQL expressions with
// dbms_lob and other user-defined functions can be used to mask columns of
// Large Object data type (LOB). To learn more, check SQL Expression in the
// Data Safe documentation.
type SqlExpressionFormatEntry struct {

	// The SQL expression to be used to generate the masked values. It can
	// consist of one or more values, operators, and SQL functions that
	// evaluate to a value. It can also contain substitution columns from
	// the same table. Specify the substitution columns within percent (%)
	// symbols.
	SqlExpression *string `mandatory:"true" json:"sqlExpression"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m SqlExpressionFormatEntry) GetDescription() *string {
	return m.Description
}

func (m SqlExpressionFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlExpressionFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SqlExpressionFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSqlExpressionFormatEntry SqlExpressionFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeSqlExpressionFormatEntry
	}{
		"SQL_EXPRESSION",
		(MarshalTypeSqlExpressionFormatEntry)(m),
	}

	return json.Marshal(&s)
}
