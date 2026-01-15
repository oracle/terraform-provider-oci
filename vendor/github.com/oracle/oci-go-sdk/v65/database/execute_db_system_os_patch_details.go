// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteDbSystemOsPatchDetails Request details for submitting an operating system (OS) patch action on a DB system.
// Use PRECHECK to validate prerequisites and surface the expected changes without applying them. Use APPLY to install the selected updates. Some updates may require a reboot to take effect.
type ExecuteDbSystemOsPatchDetails struct {

	// Operation system (OS) patch action to perform
	// * PRECHECK: No changes applied; runs validation/dry run.
	// * APPLY: Installs updates; may require a reboot (see OS patch history entry details to determine isRebootRequired).
	Action ExecuteDbSystemOsPatchDetailsActionEnum `mandatory:"true" json:"action"`

	// The OCID of the DB node to target for this patch action. If omitted, the action applies to all nodes in the DB system.
	DbNodeId *string `mandatory:"false" json:"dbNodeId"`
}

func (m ExecuteDbSystemOsPatchDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteDbSystemOsPatchDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecuteDbSystemOsPatchDetailsActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetExecuteDbSystemOsPatchDetailsActionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExecuteDbSystemOsPatchDetailsActionEnum Enum with underlying type: string
type ExecuteDbSystemOsPatchDetailsActionEnum string

// Set of constants representing the allowable values for ExecuteDbSystemOsPatchDetailsActionEnum
const (
	ExecuteDbSystemOsPatchDetailsActionPrecheck ExecuteDbSystemOsPatchDetailsActionEnum = "PRECHECK"
	ExecuteDbSystemOsPatchDetailsActionApply    ExecuteDbSystemOsPatchDetailsActionEnum = "APPLY"
)

var mappingExecuteDbSystemOsPatchDetailsActionEnum = map[string]ExecuteDbSystemOsPatchDetailsActionEnum{
	"PRECHECK": ExecuteDbSystemOsPatchDetailsActionPrecheck,
	"APPLY":    ExecuteDbSystemOsPatchDetailsActionApply,
}

var mappingExecuteDbSystemOsPatchDetailsActionEnumLowerCase = map[string]ExecuteDbSystemOsPatchDetailsActionEnum{
	"precheck": ExecuteDbSystemOsPatchDetailsActionPrecheck,
	"apply":    ExecuteDbSystemOsPatchDetailsActionApply,
}

// GetExecuteDbSystemOsPatchDetailsActionEnumValues Enumerates the set of values for ExecuteDbSystemOsPatchDetailsActionEnum
func GetExecuteDbSystemOsPatchDetailsActionEnumValues() []ExecuteDbSystemOsPatchDetailsActionEnum {
	values := make([]ExecuteDbSystemOsPatchDetailsActionEnum, 0)
	for _, v := range mappingExecuteDbSystemOsPatchDetailsActionEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteDbSystemOsPatchDetailsActionEnumStringValues Enumerates the set of values in String for ExecuteDbSystemOsPatchDetailsActionEnum
func GetExecuteDbSystemOsPatchDetailsActionEnumStringValues() []string {
	return []string{
		"PRECHECK",
		"APPLY",
	}
}

// GetMappingExecuteDbSystemOsPatchDetailsActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteDbSystemOsPatchDetailsActionEnum(val string) (ExecuteDbSystemOsPatchDetailsActionEnum, bool) {
	enum, ok := mappingExecuteDbSystemOsPatchDetailsActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
