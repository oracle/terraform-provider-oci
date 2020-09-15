// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
//

package datacatalog

// TermWorkflowStatusEnum Enum with underlying type: string
type TermWorkflowStatusEnum string

// Set of constants representing the allowable values for TermWorkflowStatusEnum
const (
	TermWorkflowStatusNew         TermWorkflowStatusEnum = "NEW"
	TermWorkflowStatusApproved    TermWorkflowStatusEnum = "APPROVED"
	TermWorkflowStatusUnderReview TermWorkflowStatusEnum = "UNDER_REVIEW"
	TermWorkflowStatusEscalated   TermWorkflowStatusEnum = "ESCALATED"
)

var mappingTermWorkflowStatus = map[string]TermWorkflowStatusEnum{
	"NEW":          TermWorkflowStatusNew,
	"APPROVED":     TermWorkflowStatusApproved,
	"UNDER_REVIEW": TermWorkflowStatusUnderReview,
	"ESCALATED":    TermWorkflowStatusEscalated,
}

// GetTermWorkflowStatusEnumValues Enumerates the set of values for TermWorkflowStatusEnum
func GetTermWorkflowStatusEnumValues() []TermWorkflowStatusEnum {
	values := make([]TermWorkflowStatusEnum, 0)
	for _, v := range mappingTermWorkflowStatus {
		values = append(values, v)
	}
	return values
}
