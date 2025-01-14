// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

	// Position of PL/SQL procedure OUT parameter. The value of this property is ignored during update, if "outParamType" is set to NO_OUT_PARAM value.
	OutParamPosition *int `mandatory:"true" json:"outParamPosition"`

	// SQL Type of PL/SQL procedure OUT parameter. During the update, to completely remove the out parameter, use the value NO_OUT_PARAM. In that case, the value of "outParamPosition" will be ignored.
	OutParamType SqlOutParamTypesEnum `mandatory:"true" json:"outParamType"`

	// Name of the Out Parameter
	OutParamName *string `mandatory:"false" json:"outParamName"`
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
