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

// ServiceCatalogStatusEnumEnum Enum with underlying type: string
type ServiceCatalogStatusEnumEnum string

// Set of constants representing the allowable values for ServiceCatalogStatusEnumEnum
const (
	ServiceCatalogStatusEnumActive   ServiceCatalogStatusEnumEnum = "ACTIVE"
	ServiceCatalogStatusEnumInactive ServiceCatalogStatusEnumEnum = "INACTIVE"
)

var mappingServiceCatalogStatusEnumEnum = map[string]ServiceCatalogStatusEnumEnum{
	"ACTIVE":   ServiceCatalogStatusEnumActive,
	"INACTIVE": ServiceCatalogStatusEnumInactive,
}

var mappingServiceCatalogStatusEnumEnumLowerCase = map[string]ServiceCatalogStatusEnumEnum{
	"active":   ServiceCatalogStatusEnumActive,
	"inactive": ServiceCatalogStatusEnumInactive,
}

// GetServiceCatalogStatusEnumEnumValues Enumerates the set of values for ServiceCatalogStatusEnumEnum
func GetServiceCatalogStatusEnumEnumValues() []ServiceCatalogStatusEnumEnum {
	values := make([]ServiceCatalogStatusEnumEnum, 0)
	for _, v := range mappingServiceCatalogStatusEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceCatalogStatusEnumEnumStringValues Enumerates the set of values in String for ServiceCatalogStatusEnumEnum
func GetServiceCatalogStatusEnumEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingServiceCatalogStatusEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceCatalogStatusEnumEnum(val string) (ServiceCatalogStatusEnumEnum, bool) {
	enum, ok := mappingServiceCatalogStatusEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
