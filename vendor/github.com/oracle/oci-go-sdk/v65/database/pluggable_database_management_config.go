// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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

// PluggableDatabaseManagementConfig The configuration of the Pluggable Database Management service.
type PluggableDatabaseManagementConfig struct {

	// The status of the Pluggable Database Management service.
	ManagementStatus PluggableDatabaseManagementConfigManagementStatusEnum `mandatory:"true" json:"managementStatus"`
}

func (m PluggableDatabaseManagementConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabaseManagementConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluggableDatabaseManagementConfigManagementStatusEnum(string(m.ManagementStatus)); !ok && m.ManagementStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagementStatus: %s. Supported values are: %s.", m.ManagementStatus, strings.Join(GetPluggableDatabaseManagementConfigManagementStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseManagementConfigManagementStatusEnum Enum with underlying type: string
type PluggableDatabaseManagementConfigManagementStatusEnum string

// Set of constants representing the allowable values for PluggableDatabaseManagementConfigManagementStatusEnum
const (
	PluggableDatabaseManagementConfigManagementStatusEnabling        PluggableDatabaseManagementConfigManagementStatusEnum = "ENABLING"
	PluggableDatabaseManagementConfigManagementStatusEnabled         PluggableDatabaseManagementConfigManagementStatusEnum = "ENABLED"
	PluggableDatabaseManagementConfigManagementStatusDisabling       PluggableDatabaseManagementConfigManagementStatusEnum = "DISABLING"
	PluggableDatabaseManagementConfigManagementStatusDisabled        PluggableDatabaseManagementConfigManagementStatusEnum = "DISABLED"
	PluggableDatabaseManagementConfigManagementStatusUpdating        PluggableDatabaseManagementConfigManagementStatusEnum = "UPDATING"
	PluggableDatabaseManagementConfigManagementStatusFailedEnabling  PluggableDatabaseManagementConfigManagementStatusEnum = "FAILED_ENABLING"
	PluggableDatabaseManagementConfigManagementStatusFailedDisabling PluggableDatabaseManagementConfigManagementStatusEnum = "FAILED_DISABLING"
	PluggableDatabaseManagementConfigManagementStatusFailedUpdating  PluggableDatabaseManagementConfigManagementStatusEnum = "FAILED_UPDATING"
)

var mappingPluggableDatabaseManagementConfigManagementStatusEnum = map[string]PluggableDatabaseManagementConfigManagementStatusEnum{
	"ENABLING":         PluggableDatabaseManagementConfigManagementStatusEnabling,
	"ENABLED":          PluggableDatabaseManagementConfigManagementStatusEnabled,
	"DISABLING":        PluggableDatabaseManagementConfigManagementStatusDisabling,
	"DISABLED":         PluggableDatabaseManagementConfigManagementStatusDisabled,
	"UPDATING":         PluggableDatabaseManagementConfigManagementStatusUpdating,
	"FAILED_ENABLING":  PluggableDatabaseManagementConfigManagementStatusFailedEnabling,
	"FAILED_DISABLING": PluggableDatabaseManagementConfigManagementStatusFailedDisabling,
	"FAILED_UPDATING":  PluggableDatabaseManagementConfigManagementStatusFailedUpdating,
}

var mappingPluggableDatabaseManagementConfigManagementStatusEnumLowerCase = map[string]PluggableDatabaseManagementConfigManagementStatusEnum{
	"enabling":         PluggableDatabaseManagementConfigManagementStatusEnabling,
	"enabled":          PluggableDatabaseManagementConfigManagementStatusEnabled,
	"disabling":        PluggableDatabaseManagementConfigManagementStatusDisabling,
	"disabled":         PluggableDatabaseManagementConfigManagementStatusDisabled,
	"updating":         PluggableDatabaseManagementConfigManagementStatusUpdating,
	"failed_enabling":  PluggableDatabaseManagementConfigManagementStatusFailedEnabling,
	"failed_disabling": PluggableDatabaseManagementConfigManagementStatusFailedDisabling,
	"failed_updating":  PluggableDatabaseManagementConfigManagementStatusFailedUpdating,
}

// GetPluggableDatabaseManagementConfigManagementStatusEnumValues Enumerates the set of values for PluggableDatabaseManagementConfigManagementStatusEnum
func GetPluggableDatabaseManagementConfigManagementStatusEnumValues() []PluggableDatabaseManagementConfigManagementStatusEnum {
	values := make([]PluggableDatabaseManagementConfigManagementStatusEnum, 0)
	for _, v := range mappingPluggableDatabaseManagementConfigManagementStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseManagementConfigManagementStatusEnumStringValues Enumerates the set of values in String for PluggableDatabaseManagementConfigManagementStatusEnum
func GetPluggableDatabaseManagementConfigManagementStatusEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
		"UPDATING",
		"FAILED_ENABLING",
		"FAILED_DISABLING",
		"FAILED_UPDATING",
	}
}

// GetMappingPluggableDatabaseManagementConfigManagementStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseManagementConfigManagementStatusEnum(val string) (PluggableDatabaseManagementConfigManagementStatusEnum, bool) {
	enum, ok := mappingPluggableDatabaseManagementConfigManagementStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
