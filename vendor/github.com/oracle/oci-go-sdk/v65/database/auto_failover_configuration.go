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

// AutoFailoverConfiguration The properties for defining auto failover configuration.
type AutoFailoverConfiguration struct {

	// The state of managed auto failover.
	ManagedAutoFailover AutoFailoverConfigurationManagedAutoFailoverEnum `mandatory:"false" json:"managedAutoFailover,omitempty"`

	// Specifies the `DB_UNIQUE_NAME` of the data guard group member databases.
	FailoverTargets []string `mandatory:"false" json:"failoverTargets"`
}

func (m AutoFailoverConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoFailoverConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAutoFailoverConfigurationManagedAutoFailoverEnum(string(m.ManagedAutoFailover)); !ok && m.ManagedAutoFailover != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ManagedAutoFailover: %s. Supported values are: %s.", m.ManagedAutoFailover, strings.Join(GetAutoFailoverConfigurationManagedAutoFailoverEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AutoFailoverConfigurationManagedAutoFailoverEnum Enum with underlying type: string
type AutoFailoverConfigurationManagedAutoFailoverEnum string

// Set of constants representing the allowable values for AutoFailoverConfigurationManagedAutoFailoverEnum
const (
	AutoFailoverConfigurationManagedAutoFailoverEnable  AutoFailoverConfigurationManagedAutoFailoverEnum = "ENABLE"
	AutoFailoverConfigurationManagedAutoFailoverDisable AutoFailoverConfigurationManagedAutoFailoverEnum = "DISABLE"
)

var mappingAutoFailoverConfigurationManagedAutoFailoverEnum = map[string]AutoFailoverConfigurationManagedAutoFailoverEnum{
	"ENABLE":  AutoFailoverConfigurationManagedAutoFailoverEnable,
	"DISABLE": AutoFailoverConfigurationManagedAutoFailoverDisable,
}

var mappingAutoFailoverConfigurationManagedAutoFailoverEnumLowerCase = map[string]AutoFailoverConfigurationManagedAutoFailoverEnum{
	"enable":  AutoFailoverConfigurationManagedAutoFailoverEnable,
	"disable": AutoFailoverConfigurationManagedAutoFailoverDisable,
}

// GetAutoFailoverConfigurationManagedAutoFailoverEnumValues Enumerates the set of values for AutoFailoverConfigurationManagedAutoFailoverEnum
func GetAutoFailoverConfigurationManagedAutoFailoverEnumValues() []AutoFailoverConfigurationManagedAutoFailoverEnum {
	values := make([]AutoFailoverConfigurationManagedAutoFailoverEnum, 0)
	for _, v := range mappingAutoFailoverConfigurationManagedAutoFailoverEnum {
		values = append(values, v)
	}
	return values
}

// GetAutoFailoverConfigurationManagedAutoFailoverEnumStringValues Enumerates the set of values in String for AutoFailoverConfigurationManagedAutoFailoverEnum
func GetAutoFailoverConfigurationManagedAutoFailoverEnumStringValues() []string {
	return []string{
		"ENABLE",
		"DISABLE",
	}
}

// GetMappingAutoFailoverConfigurationManagedAutoFailoverEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAutoFailoverConfigurationManagedAutoFailoverEnum(val string) (AutoFailoverConfigurationManagedAutoFailoverEnum, bool) {
	enum, ok := mappingAutoFailoverConfigurationManagedAutoFailoverEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
