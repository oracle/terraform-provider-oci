// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CompatibleFormatsForDataTypes A list of basic masking formats compatible with and grouped by the supported data types. The data types are grouped into
// the following categories -
//   Character - Includes CHAR, NCHAR, VARCHAR2, and NVARCHAR2
//   Numeric - Includes NUMBER, FLOAT, RAW, BINARY_FLOAT, and BINARY_DOUBLE
//   Date - Includes DATE and TIMESTAMP
//   LOB - Includes BLOB, CLOB, and NCLOB
//   All - Includes all the supported data types
type CompatibleFormatsForDataTypes struct {

	// An array of lists of basic masking formats compatible with the supported data types.
	FormatsForDataType []FormatsForDataType `mandatory:"true" json:"formatsForDataType"`
}

func (m CompatibleFormatsForDataTypes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CompatibleFormatsForDataTypes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
