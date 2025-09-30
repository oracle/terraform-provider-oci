// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Catalog API
//
// Use the Service Catalog API to manage solutions in Oracle Cloud Infrastructure Service Catalog.
// For more information, see Overview of Service Catalog (https://docs.oracle.com/iaas/Content/service-catalog/overview_of_service_catalog.htm).
//

package servicecatalog

import (
	"strings"
)

// ServiceCatalogModeEnumEnum Enum with underlying type: string
type ServiceCatalogModeEnumEnum string

// Set of constants representing the allowable values for ServiceCatalogModeEnumEnum
const (
	ServiceCatalogModeEnumEnabled  ServiceCatalogModeEnumEnum = "ENABLED"
	ServiceCatalogModeEnumDisabled ServiceCatalogModeEnumEnum = "DISABLED"
)

var mappingServiceCatalogModeEnumEnum = map[string]ServiceCatalogModeEnumEnum{
	"ENABLED":  ServiceCatalogModeEnumEnabled,
	"DISABLED": ServiceCatalogModeEnumDisabled,
}

var mappingServiceCatalogModeEnumEnumLowerCase = map[string]ServiceCatalogModeEnumEnum{
	"enabled":  ServiceCatalogModeEnumEnabled,
	"disabled": ServiceCatalogModeEnumDisabled,
}

// GetServiceCatalogModeEnumEnumValues Enumerates the set of values for ServiceCatalogModeEnumEnum
func GetServiceCatalogModeEnumEnumValues() []ServiceCatalogModeEnumEnum {
	values := make([]ServiceCatalogModeEnumEnum, 0)
	for _, v := range mappingServiceCatalogModeEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceCatalogModeEnumEnumStringValues Enumerates the set of values in String for ServiceCatalogModeEnumEnum
func GetServiceCatalogModeEnumEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
	}
}

// GetMappingServiceCatalogModeEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceCatalogModeEnumEnum(val string) (ServiceCatalogModeEnumEnum, bool) {
	enum, ok := mappingServiceCatalogModeEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
