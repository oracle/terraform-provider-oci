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

// StackMonitoringConfig The configuration of Stack Monitoring for the external database.
type StackMonitoringConfig struct {

	// The status of Stack Monitoring.
	StackMonitoringStatus StackMonitoringConfigStackMonitoringStatusEnum `mandatory:"true" json:"stackMonitoringStatus"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the
	// CreateExternalDatabaseConnectorDetails.
	StackMonitoringConnectorId *string `mandatory:"false" json:"stackMonitoringConnectorId"`
}

func (m StackMonitoringConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackMonitoringConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStackMonitoringConfigStackMonitoringStatusEnum(string(m.StackMonitoringStatus)); !ok && m.StackMonitoringStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StackMonitoringStatus: %s. Supported values are: %s.", m.StackMonitoringStatus, strings.Join(GetStackMonitoringConfigStackMonitoringStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StackMonitoringConfigStackMonitoringStatusEnum Enum with underlying type: string
type StackMonitoringConfigStackMonitoringStatusEnum string

// Set of constants representing the allowable values for StackMonitoringConfigStackMonitoringStatusEnum
const (
	StackMonitoringConfigStackMonitoringStatusEnabling        StackMonitoringConfigStackMonitoringStatusEnum = "ENABLING"
	StackMonitoringConfigStackMonitoringStatusEnabled         StackMonitoringConfigStackMonitoringStatusEnum = "ENABLED"
	StackMonitoringConfigStackMonitoringStatusDisabling       StackMonitoringConfigStackMonitoringStatusEnum = "DISABLING"
	StackMonitoringConfigStackMonitoringStatusNotEnabled      StackMonitoringConfigStackMonitoringStatusEnum = "NOT_ENABLED"
	StackMonitoringConfigStackMonitoringStatusFailedEnabling  StackMonitoringConfigStackMonitoringStatusEnum = "FAILED_ENABLING"
	StackMonitoringConfigStackMonitoringStatusFailedDisabling StackMonitoringConfigStackMonitoringStatusEnum = "FAILED_DISABLING"
)

var mappingStackMonitoringConfigStackMonitoringStatusEnum = map[string]StackMonitoringConfigStackMonitoringStatusEnum{
	"ENABLING":         StackMonitoringConfigStackMonitoringStatusEnabling,
	"ENABLED":          StackMonitoringConfigStackMonitoringStatusEnabled,
	"DISABLING":        StackMonitoringConfigStackMonitoringStatusDisabling,
	"NOT_ENABLED":      StackMonitoringConfigStackMonitoringStatusNotEnabled,
	"FAILED_ENABLING":  StackMonitoringConfigStackMonitoringStatusFailedEnabling,
	"FAILED_DISABLING": StackMonitoringConfigStackMonitoringStatusFailedDisabling,
}

var mappingStackMonitoringConfigStackMonitoringStatusEnumLowerCase = map[string]StackMonitoringConfigStackMonitoringStatusEnum{
	"enabling":         StackMonitoringConfigStackMonitoringStatusEnabling,
	"enabled":          StackMonitoringConfigStackMonitoringStatusEnabled,
	"disabling":        StackMonitoringConfigStackMonitoringStatusDisabling,
	"not_enabled":      StackMonitoringConfigStackMonitoringStatusNotEnabled,
	"failed_enabling":  StackMonitoringConfigStackMonitoringStatusFailedEnabling,
	"failed_disabling": StackMonitoringConfigStackMonitoringStatusFailedDisabling,
}

// GetStackMonitoringConfigStackMonitoringStatusEnumValues Enumerates the set of values for StackMonitoringConfigStackMonitoringStatusEnum
func GetStackMonitoringConfigStackMonitoringStatusEnumValues() []StackMonitoringConfigStackMonitoringStatusEnum {
	values := make([]StackMonitoringConfigStackMonitoringStatusEnum, 0)
	for _, v := range mappingStackMonitoringConfigStackMonitoringStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetStackMonitoringConfigStackMonitoringStatusEnumStringValues Enumerates the set of values in String for StackMonitoringConfigStackMonitoringStatusEnum
func GetStackMonitoringConfigStackMonitoringStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"NOT_ENABLED",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
	}
}

// GetMappingStackMonitoringConfigStackMonitoringStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackMonitoringConfigStackMonitoringStatusEnum(val string) (StackMonitoringConfigStackMonitoringStatusEnum, bool) {
	enum, ok := mappingStackMonitoringConfigStackMonitoringStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
