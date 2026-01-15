// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"strings"
)

// CatalogListingVersionCriteriaEnumEnum Enum with underlying type: string
type CatalogListingVersionCriteriaEnumEnum string

// Set of constants representing the allowable values for CatalogListingVersionCriteriaEnumEnum
const (
	CatalogListingVersionCriteriaEnumListAllVersions         CatalogListingVersionCriteriaEnumEnum = "LIST_ALL_VERSIONS"
	CatalogListingVersionCriteriaEnumListEarliestVersionOnly CatalogListingVersionCriteriaEnumEnum = "LIST_EARLIEST_VERSION_ONLY"
	CatalogListingVersionCriteriaEnumListLatestVersionOnly   CatalogListingVersionCriteriaEnumEnum = "LIST_LATEST_VERSION_ONLY"
)

var mappingCatalogListingVersionCriteriaEnumEnum = map[string]CatalogListingVersionCriteriaEnumEnum{
	"LIST_ALL_VERSIONS":          CatalogListingVersionCriteriaEnumListAllVersions,
	"LIST_EARLIEST_VERSION_ONLY": CatalogListingVersionCriteriaEnumListEarliestVersionOnly,
	"LIST_LATEST_VERSION_ONLY":   CatalogListingVersionCriteriaEnumListLatestVersionOnly,
}

var mappingCatalogListingVersionCriteriaEnumEnumLowerCase = map[string]CatalogListingVersionCriteriaEnumEnum{
	"list_all_versions":          CatalogListingVersionCriteriaEnumListAllVersions,
	"list_earliest_version_only": CatalogListingVersionCriteriaEnumListEarliestVersionOnly,
	"list_latest_version_only":   CatalogListingVersionCriteriaEnumListLatestVersionOnly,
}

// GetCatalogListingVersionCriteriaEnumEnumValues Enumerates the set of values for CatalogListingVersionCriteriaEnumEnum
func GetCatalogListingVersionCriteriaEnumEnumValues() []CatalogListingVersionCriteriaEnumEnum {
	values := make([]CatalogListingVersionCriteriaEnumEnum, 0)
	for _, v := range mappingCatalogListingVersionCriteriaEnumEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogListingVersionCriteriaEnumEnumStringValues Enumerates the set of values in String for CatalogListingVersionCriteriaEnumEnum
func GetCatalogListingVersionCriteriaEnumEnumStringValues() []string {
	return []string{
		"LIST_ALL_VERSIONS",
		"LIST_EARLIEST_VERSION_ONLY",
		"LIST_LATEST_VERSION_ONLY",
	}
}

// GetMappingCatalogListingVersionCriteriaEnumEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogListingVersionCriteriaEnumEnum(val string) (CatalogListingVersionCriteriaEnumEnum, bool) {
	enum, ok := mappingCatalogListingVersionCriteriaEnumEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
