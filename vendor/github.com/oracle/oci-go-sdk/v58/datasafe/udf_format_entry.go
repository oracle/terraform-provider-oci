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

// UdfFormatEntry The User Defined Function masking format lets you define your own logic to
// mask column data. The return value of the user-defined function is used to
// replace the original values. The user-defined function has a fixed signature
// and is a PL/SQL function that can be invoked in a SELECT statement. To learn
// more, check User Defined Function in the Data Safe documentation.
type UdfFormatEntry struct {

	// The user-defined function in SCHEMA_NAME.PACKAGE_NAME.FUNCTION_NAME format.
	// It can be a standalone or packaged function, so PACKAGE_NAME is optional.
	UserDefinedFunction *string `mandatory:"true" json:"userDefinedFunction"`

	// The description of the format entry.
	Description *string `mandatory:"false" json:"description"`
}

//GetDescription returns Description
func (m UdfFormatEntry) GetDescription() *string {
	return m.Description
}

func (m UdfFormatEntry) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UdfFormatEntry) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UdfFormatEntry) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUdfFormatEntry UdfFormatEntry
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUdfFormatEntry
	}{
		"USER_DEFINED_FUNCTION",
		(MarshalTypeUdfFormatEntry)(m),
	}

	return json.Marshal(&s)
}
