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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FacetedSearchAggregation Aggregation/facets on properties of data object.
type FacetedSearchAggregation struct {

	// Name of data object property
	Type *string `mandatory:"false" json:"type"`

	// Count of number of data objects having property.
	Aggregation map[string]int64 `mandatory:"false" json:"aggregation"`

	// Data type of object property.
	DataType *string `mandatory:"false" json:"dataType"`

	// Type of property that indicates if it was defined by the user or system.
	// CUSTOM_PROPERTY is defined by the user on a data object.
	// DEFAULT_PROPERTY is defined by the system on a data object.
	PropertyType FacetedSearchAggregationPropertyTypeEnum `mandatory:"false" json:"propertyType,omitempty"`
}

func (m FacetedSearchAggregation) String() string {
	return common.PointerString(m)
}

// FacetedSearchAggregationPropertyTypeEnum Enum with underlying type: string
type FacetedSearchAggregationPropertyTypeEnum string

// Set of constants representing the allowable values for FacetedSearchAggregationPropertyTypeEnum
const (
	FacetedSearchAggregationPropertyTypeCustomProperty  FacetedSearchAggregationPropertyTypeEnum = "CUSTOM_PROPERTY"
	FacetedSearchAggregationPropertyTypeDefaultProperty FacetedSearchAggregationPropertyTypeEnum = "DEFAULT_PROPERTY"
)

var mappingFacetedSearchAggregationPropertyType = map[string]FacetedSearchAggregationPropertyTypeEnum{
	"CUSTOM_PROPERTY":  FacetedSearchAggregationPropertyTypeCustomProperty,
	"DEFAULT_PROPERTY": FacetedSearchAggregationPropertyTypeDefaultProperty,
}

// GetFacetedSearchAggregationPropertyTypeEnumValues Enumerates the set of values for FacetedSearchAggregationPropertyTypeEnum
func GetFacetedSearchAggregationPropertyTypeEnumValues() []FacetedSearchAggregationPropertyTypeEnum {
	values := make([]FacetedSearchAggregationPropertyTypeEnum, 0)
	for _, v := range mappingFacetedSearchAggregationPropertyType {
		values = append(values, v)
	}
	return values
}
