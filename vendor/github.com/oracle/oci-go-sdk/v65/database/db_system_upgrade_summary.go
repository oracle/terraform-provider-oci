// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// DbSystemUpgradeSummary It describes the various upgrade properties such as component, osversion, giversion for any VM DB system.
type DbSystemUpgradeSummary struct {

	// The component on which upgrade is applicable. OS (Operating System upgrade), GI (Grid Infrastructure upgrade) or OS_GI (both Operating System and Grid Infrastructure upgrade)
	Component DbSystemUpgradeSummaryComponentEnum `mandatory:"true" json:"component"`

	// The version of the OS for this upgrade eg. Oracle Linux Server release 7.9
	OsVersion *string `mandatory:"false" json:"osVersion"`

	// The version of the grid infrastructure for this upgrade. This is only applicable for ASM based DbSystems
	GiVersion *string `mandatory:"false" json:"giVersion"`
}

func (m DbSystemUpgradeSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DbSystemUpgradeSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDbSystemUpgradeSummaryComponentEnum(string(m.Component)); !ok && m.Component != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Component: %s. Supported values are: %s.", m.Component, strings.Join(GetDbSystemUpgradeSummaryComponentEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DbSystemUpgradeSummaryComponentEnum Enum with underlying type: string
type DbSystemUpgradeSummaryComponentEnum string

// Set of constants representing the allowable values for DbSystemUpgradeSummaryComponentEnum
const (
	DbSystemUpgradeSummaryComponentOs   DbSystemUpgradeSummaryComponentEnum = "OS"
	DbSystemUpgradeSummaryComponentOsGi DbSystemUpgradeSummaryComponentEnum = "OS_GI"
	DbSystemUpgradeSummaryComponentGi   DbSystemUpgradeSummaryComponentEnum = "GI"
)

var mappingDbSystemUpgradeSummaryComponentEnum = map[string]DbSystemUpgradeSummaryComponentEnum{
	"OS":    DbSystemUpgradeSummaryComponentOs,
	"OS_GI": DbSystemUpgradeSummaryComponentOsGi,
	"GI":    DbSystemUpgradeSummaryComponentGi,
}

var mappingDbSystemUpgradeSummaryComponentEnumLowerCase = map[string]DbSystemUpgradeSummaryComponentEnum{
	"os":    DbSystemUpgradeSummaryComponentOs,
	"os_gi": DbSystemUpgradeSummaryComponentOsGi,
	"gi":    DbSystemUpgradeSummaryComponentGi,
}

// GetDbSystemUpgradeSummaryComponentEnumValues Enumerates the set of values for DbSystemUpgradeSummaryComponentEnum
func GetDbSystemUpgradeSummaryComponentEnumValues() []DbSystemUpgradeSummaryComponentEnum {
	values := make([]DbSystemUpgradeSummaryComponentEnum, 0)
	for _, v := range mappingDbSystemUpgradeSummaryComponentEnum {
		values = append(values, v)
	}
	return values
}

// GetDbSystemUpgradeSummaryComponentEnumStringValues Enumerates the set of values in String for DbSystemUpgradeSummaryComponentEnum
func GetDbSystemUpgradeSummaryComponentEnumStringValues() []string {
	return []string{
		"OS",
		"OS_GI",
		"GI",
	}
}

// GetMappingDbSystemUpgradeSummaryComponentEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDbSystemUpgradeSummaryComponentEnum(val string) (DbSystemUpgradeSummaryComponentEnum, bool) {
	enum, ok := mappingDbSystemUpgradeSummaryComponentEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
