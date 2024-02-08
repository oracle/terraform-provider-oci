// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// VmClusterUpdateDetails Details specifying which maintenance update to apply to the VM Cluster and which action is to be performed by the maintenance update. Applies to Exadata Cloud@Customer instances only.
type VmClusterUpdateDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the maintenance update.
	UpdateId *string `mandatory:"false" json:"updateId"`

	// The update action to perform.
	UpdateAction VmClusterUpdateDetailsUpdateActionEnum `mandatory:"false" json:"updateAction,omitempty"`
}

func (m VmClusterUpdateDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VmClusterUpdateDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingVmClusterUpdateDetailsUpdateActionEnum(string(m.UpdateAction)); !ok && m.UpdateAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for UpdateAction: %s. Supported values are: %s.", m.UpdateAction, strings.Join(GetVmClusterUpdateDetailsUpdateActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// VmClusterUpdateDetailsUpdateActionEnum Enum with underlying type: string
type VmClusterUpdateDetailsUpdateActionEnum string

// Set of constants representing the allowable values for VmClusterUpdateDetailsUpdateActionEnum
const (
	VmClusterUpdateDetailsUpdateActionRollingApply VmClusterUpdateDetailsUpdateActionEnum = "ROLLING_APPLY"
	VmClusterUpdateDetailsUpdateActionPrecheck     VmClusterUpdateDetailsUpdateActionEnum = "PRECHECK"
	VmClusterUpdateDetailsUpdateActionRollback     VmClusterUpdateDetailsUpdateActionEnum = "ROLLBACK"
)

var mappingVmClusterUpdateDetailsUpdateActionEnum = map[string]VmClusterUpdateDetailsUpdateActionEnum{
	"ROLLING_APPLY": VmClusterUpdateDetailsUpdateActionRollingApply,
	"PRECHECK":      VmClusterUpdateDetailsUpdateActionPrecheck,
	"ROLLBACK":      VmClusterUpdateDetailsUpdateActionRollback,
}

var mappingVmClusterUpdateDetailsUpdateActionEnumLowerCase = map[string]VmClusterUpdateDetailsUpdateActionEnum{
	"rolling_apply": VmClusterUpdateDetailsUpdateActionRollingApply,
	"precheck":      VmClusterUpdateDetailsUpdateActionPrecheck,
	"rollback":      VmClusterUpdateDetailsUpdateActionRollback,
}

// GetVmClusterUpdateDetailsUpdateActionEnumValues Enumerates the set of values for VmClusterUpdateDetailsUpdateActionEnum
func GetVmClusterUpdateDetailsUpdateActionEnumValues() []VmClusterUpdateDetailsUpdateActionEnum {
	values := make([]VmClusterUpdateDetailsUpdateActionEnum, 0)
	for _, v := range mappingVmClusterUpdateDetailsUpdateActionEnum {
		values = append(values, v)
	}
	return values
}

// GetVmClusterUpdateDetailsUpdateActionEnumStringValues Enumerates the set of values in String for VmClusterUpdateDetailsUpdateActionEnum
func GetVmClusterUpdateDetailsUpdateActionEnumStringValues() []string {
	return []string{
		"ROLLING_APPLY",
		"PRECHECK",
		"ROLLBACK",
	}
}

// GetMappingVmClusterUpdateDetailsUpdateActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingVmClusterUpdateDetailsUpdateActionEnum(val string) (VmClusterUpdateDetailsUpdateActionEnum, bool) {
	enum, ok := mappingVmClusterUpdateDetailsUpdateActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
