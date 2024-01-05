// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// License Manager API
//
// Use the License Manager API to manage product licenses and license records. For more information, see License Manager Overview (https://docs.cloud.oracle.com/iaas/Content/LicenseManager/Concepts/licensemanageroverview.htm).
//

package licensemanager

import (
	"strings"
)

// ProductCategoryEnum Enum with underlying type: string
type ProductCategoryEnum string

// Set of constants representing the allowable values for ProductCategoryEnum
const (
	ProductCategoryBase   ProductCategoryEnum = "BASE"
	ProductCategoryOption ProductCategoryEnum = "OPTION"
)

var mappingProductCategoryEnum = map[string]ProductCategoryEnum{
	"BASE":   ProductCategoryBase,
	"OPTION": ProductCategoryOption,
}

var mappingProductCategoryEnumLowerCase = map[string]ProductCategoryEnum{
	"base":   ProductCategoryBase,
	"option": ProductCategoryOption,
}

// GetProductCategoryEnumValues Enumerates the set of values for ProductCategoryEnum
func GetProductCategoryEnumValues() []ProductCategoryEnum {
	values := make([]ProductCategoryEnum, 0)
	for _, v := range mappingProductCategoryEnum {
		values = append(values, v)
	}
	return values
}

// GetProductCategoryEnumStringValues Enumerates the set of values in String for ProductCategoryEnum
func GetProductCategoryEnumStringValues() []string {
	return []string{
		"BASE",
		"OPTION",
	}
}

// GetMappingProductCategoryEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingProductCategoryEnum(val string) (ProductCategoryEnum, bool) {
	enum, ok := mappingProductCategoryEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
