// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"strings"
)

// ClassicSubscriptionEnvironmentNameEnum Enum with underlying type: string
type ClassicSubscriptionEnvironmentNameEnum string

// Set of constants representing the allowable values for ClassicSubscriptionEnvironmentNameEnum
const (
	ClassicSubscriptionEnvironmentNameProd     ClassicSubscriptionEnvironmentNameEnum = "PROD"
	ClassicSubscriptionEnvironmentNamePreprod  ClassicSubscriptionEnvironmentNameEnum = "PREPROD"
	ClassicSubscriptionEnvironmentNamePreprod1 ClassicSubscriptionEnvironmentNameEnum = "PREPROD1"
	ClassicSubscriptionEnvironmentNamePintlab  ClassicSubscriptionEnvironmentNameEnum = "PINTLAB"
	ClassicSubscriptionEnvironmentNameMiglab   ClassicSubscriptionEnvironmentNameEnum = "MIGLAB"
	ClassicSubscriptionEnvironmentNamePool2    ClassicSubscriptionEnvironmentNameEnum = "POOL2"
	ClassicSubscriptionEnvironmentNamePintlab2 ClassicSubscriptionEnvironmentNameEnum = "PINTLAB2"
	ClassicSubscriptionEnvironmentNameMylab0   ClassicSubscriptionEnvironmentNameEnum = "MYLAB0"
	ClassicSubscriptionEnvironmentNameMylab1   ClassicSubscriptionEnvironmentNameEnum = "MYLAB1"
	ClassicSubscriptionEnvironmentNameMylab2   ClassicSubscriptionEnvironmentNameEnum = "MYLAB2"
	ClassicSubscriptionEnvironmentNameMylab3   ClassicSubscriptionEnvironmentNameEnum = "MYLAB3"
	ClassicSubscriptionEnvironmentNameMylab4   ClassicSubscriptionEnvironmentNameEnum = "MYLAB4"
	ClassicSubscriptionEnvironmentNameMylab5   ClassicSubscriptionEnvironmentNameEnum = "MYLAB5"
)

var mappingClassicSubscriptionEnvironmentNameEnum = map[string]ClassicSubscriptionEnvironmentNameEnum{
	"PROD":     ClassicSubscriptionEnvironmentNameProd,
	"PREPROD":  ClassicSubscriptionEnvironmentNamePreprod,
	"PREPROD1": ClassicSubscriptionEnvironmentNamePreprod1,
	"PINTLAB":  ClassicSubscriptionEnvironmentNamePintlab,
	"MIGLAB":   ClassicSubscriptionEnvironmentNameMiglab,
	"POOL2":    ClassicSubscriptionEnvironmentNamePool2,
	"PINTLAB2": ClassicSubscriptionEnvironmentNamePintlab2,
	"MYLAB0":   ClassicSubscriptionEnvironmentNameMylab0,
	"MYLAB1":   ClassicSubscriptionEnvironmentNameMylab1,
	"MYLAB2":   ClassicSubscriptionEnvironmentNameMylab2,
	"MYLAB3":   ClassicSubscriptionEnvironmentNameMylab3,
	"MYLAB4":   ClassicSubscriptionEnvironmentNameMylab4,
	"MYLAB5":   ClassicSubscriptionEnvironmentNameMylab5,
}

var mappingClassicSubscriptionEnvironmentNameEnumLowerCase = map[string]ClassicSubscriptionEnvironmentNameEnum{
	"prod":     ClassicSubscriptionEnvironmentNameProd,
	"preprod":  ClassicSubscriptionEnvironmentNamePreprod,
	"preprod1": ClassicSubscriptionEnvironmentNamePreprod1,
	"pintlab":  ClassicSubscriptionEnvironmentNamePintlab,
	"miglab":   ClassicSubscriptionEnvironmentNameMiglab,
	"pool2":    ClassicSubscriptionEnvironmentNamePool2,
	"pintlab2": ClassicSubscriptionEnvironmentNamePintlab2,
	"mylab0":   ClassicSubscriptionEnvironmentNameMylab0,
	"mylab1":   ClassicSubscriptionEnvironmentNameMylab1,
	"mylab2":   ClassicSubscriptionEnvironmentNameMylab2,
	"mylab3":   ClassicSubscriptionEnvironmentNameMylab3,
	"mylab4":   ClassicSubscriptionEnvironmentNameMylab4,
	"mylab5":   ClassicSubscriptionEnvironmentNameMylab5,
}

// GetClassicSubscriptionEnvironmentNameEnumValues Enumerates the set of values for ClassicSubscriptionEnvironmentNameEnum
func GetClassicSubscriptionEnvironmentNameEnumValues() []ClassicSubscriptionEnvironmentNameEnum {
	values := make([]ClassicSubscriptionEnvironmentNameEnum, 0)
	for _, v := range mappingClassicSubscriptionEnvironmentNameEnum {
		values = append(values, v)
	}
	return values
}

// GetClassicSubscriptionEnvironmentNameEnumStringValues Enumerates the set of values in String for ClassicSubscriptionEnvironmentNameEnum
func GetClassicSubscriptionEnvironmentNameEnumStringValues() []string {
	return []string{
		"PROD",
		"PREPROD",
		"PREPROD1",
		"PINTLAB",
		"MIGLAB",
		"POOL2",
		"PINTLAB2",
		"MYLAB0",
		"MYLAB1",
		"MYLAB2",
		"MYLAB3",
		"MYLAB4",
		"MYLAB5",
	}
}

// GetMappingClassicSubscriptionEnvironmentNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingClassicSubscriptionEnvironmentNameEnum(val string) (ClassicSubscriptionEnvironmentNameEnum, bool) {
	enum, ok := mappingClassicSubscriptionEnvironmentNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
