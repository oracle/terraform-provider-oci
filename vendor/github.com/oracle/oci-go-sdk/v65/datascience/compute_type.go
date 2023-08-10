// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"strings"
)

// ComputeTypeEnum Enum with underlying type: string
type ComputeTypeEnum string

// Set of constants representing the allowable values for ComputeTypeEnum
const (
	ComputeTypeCustomerOke ComputeTypeEnum = "CUSTOMER_OKE"
)

var mappingComputeTypeEnum = map[string]ComputeTypeEnum{
	"CUSTOMER_OKE": ComputeTypeCustomerOke,
}

var mappingComputeTypeEnumLowerCase = map[string]ComputeTypeEnum{
	"customer_oke": ComputeTypeCustomerOke,
}

// GetComputeTypeEnumValues Enumerates the set of values for ComputeTypeEnum
func GetComputeTypeEnumValues() []ComputeTypeEnum {
	values := make([]ComputeTypeEnum, 0)
	for _, v := range mappingComputeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetComputeTypeEnumStringValues Enumerates the set of values in String for ComputeTypeEnum
func GetComputeTypeEnumStringValues() []string {
	return []string{
		"CUSTOMER_OKE",
	}
}

// GetMappingComputeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingComputeTypeEnum(val string) (ComputeTypeEnum, bool) {
	enum, ok := mappingComputeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
