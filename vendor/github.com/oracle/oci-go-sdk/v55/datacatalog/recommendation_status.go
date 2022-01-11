// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

// RecommendationStatusEnum Enum with underlying type: string
type RecommendationStatusEnum string

// Set of constants representing the allowable values for RecommendationStatusEnum
const (
	RecommendationStatusAccepted RecommendationStatusEnum = "ACCEPTED"
	RecommendationStatusRejected RecommendationStatusEnum = "REJECTED"
	RecommendationStatusInferred RecommendationStatusEnum = "INFERRED"
)

var mappingRecommendationStatus = map[string]RecommendationStatusEnum{
	"ACCEPTED": RecommendationStatusAccepted,
	"REJECTED": RecommendationStatusRejected,
	"INFERRED": RecommendationStatusInferred,
}

// GetRecommendationStatusEnumValues Enumerates the set of values for RecommendationStatusEnum
func GetRecommendationStatusEnumValues() []RecommendationStatusEnum {
	values := make([]RecommendationStatusEnum, 0)
	for _, v := range mappingRecommendationStatus {
		values = append(values, v)
	}
	return values
}
