// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"strings"
)

// TermWorkflowStatusEnum Enum with underlying type: string
type TermWorkflowStatusEnum string

// Set of constants representing the allowable values for TermWorkflowStatusEnum
const (
	TermWorkflowStatusNew         TermWorkflowStatusEnum = "NEW"
	TermWorkflowStatusApproved    TermWorkflowStatusEnum = "APPROVED"
	TermWorkflowStatusUnderReview TermWorkflowStatusEnum = "UNDER_REVIEW"
	TermWorkflowStatusEscalated   TermWorkflowStatusEnum = "ESCALATED"
)

var mappingTermWorkflowStatusEnum = map[string]TermWorkflowStatusEnum{
	"NEW":          TermWorkflowStatusNew,
	"APPROVED":     TermWorkflowStatusApproved,
	"UNDER_REVIEW": TermWorkflowStatusUnderReview,
	"ESCALATED":    TermWorkflowStatusEscalated,
}

// GetTermWorkflowStatusEnumValues Enumerates the set of values for TermWorkflowStatusEnum
func GetTermWorkflowStatusEnumValues() []TermWorkflowStatusEnum {
	values := make([]TermWorkflowStatusEnum, 0)
	for _, v := range mappingTermWorkflowStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTermWorkflowStatusEnumStringValues Enumerates the set of values in String for TermWorkflowStatusEnum
func GetTermWorkflowStatusEnumStringValues() []string {
	return []string{
		"NEW",
		"APPROVED",
		"UNDER_REVIEW",
		"ESCALATED",
	}
}

// GetMappingTermWorkflowStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTermWorkflowStatusEnum(val string) (TermWorkflowStatusEnum, bool) {
	mappingTermWorkflowStatusEnumIgnoreCase := make(map[string]TermWorkflowStatusEnum)
	for k, v := range mappingTermWorkflowStatusEnum {
		mappingTermWorkflowStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingTermWorkflowStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
