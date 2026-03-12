// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage Private Service Access (PSA) endpoints. PSA endpoints are used to create private access between resources in a VCN or on-premises and services in Oracle services network. For important details about how PSA endpoints work, see Access to Oracle Services: Private Service Access Endpoints (https://docs.oracle.com/iaas/Content/Network/Concepts/private-service-access.htm).
//

package psa

import (
	"strings"
)

// PsaEndpointServiceCatalogEnum Enum with underlying type: string
type PsaEndpointServiceCatalogEnum string

// Set of constants representing the allowable values for PsaEndpointServiceCatalogEnum
const (
	PsaEndpointServiceCatalogDevelopment PsaEndpointServiceCatalogEnum = "DEVELOPMENT"
	PsaEndpointServiceCatalogInternal    PsaEndpointServiceCatalogEnum = "INTERNAL"
	PsaEndpointServiceCatalogPublic      PsaEndpointServiceCatalogEnum = "PUBLIC"
)

var mappingPsaEndpointServiceCatalogEnum = map[string]PsaEndpointServiceCatalogEnum{
	"DEVELOPMENT": PsaEndpointServiceCatalogDevelopment,
	"INTERNAL":    PsaEndpointServiceCatalogInternal,
	"PUBLIC":      PsaEndpointServiceCatalogPublic,
}

var mappingPsaEndpointServiceCatalogEnumLowerCase = map[string]PsaEndpointServiceCatalogEnum{
	"development": PsaEndpointServiceCatalogDevelopment,
	"internal":    PsaEndpointServiceCatalogInternal,
	"public":      PsaEndpointServiceCatalogPublic,
}

// GetPsaEndpointServiceCatalogEnumValues Enumerates the set of values for PsaEndpointServiceCatalogEnum
func GetPsaEndpointServiceCatalogEnumValues() []PsaEndpointServiceCatalogEnum {
	values := make([]PsaEndpointServiceCatalogEnum, 0)
	for _, v := range mappingPsaEndpointServiceCatalogEnum {
		values = append(values, v)
	}
	return values
}

// GetPsaEndpointServiceCatalogEnumStringValues Enumerates the set of values in String for PsaEndpointServiceCatalogEnum
func GetPsaEndpointServiceCatalogEnumStringValues() []string {
	return []string{
		"DEVELOPMENT",
		"INTERNAL",
		"PUBLIC",
	}
}

// GetMappingPsaEndpointServiceCatalogEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPsaEndpointServiceCatalogEnum(val string) (PsaEndpointServiceCatalogEnum, bool) {
	enum, ok := mappingPsaEndpointServiceCatalogEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
