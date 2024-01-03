// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlOutParamDetails Position and SQL Type of PL/SQL OUT parameter
type SqlOutParamDetails struct {

	// Position of PL/SQL procedure OUT parameter
	OutParamPosition *int `mandatory:"true" json:"outParamPosition"`

	// SQL Type of PL/SQL procedure OUT parameter
	OutParamType SqlOutParamTypesEnum `mandatory:"true" json:"outParamType"`
}

func (m SqlOutParamDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlOutParamDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSqlOutParamTypesEnum(string(m.OutParamType)); !ok && m.OutParamType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OutParamType: %s. Supported values are: %s.", m.OutParamType, strings.Join(GetSqlOutParamTypesEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
