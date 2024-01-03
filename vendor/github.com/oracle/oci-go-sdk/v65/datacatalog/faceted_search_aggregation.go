// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Catalog API
//
// Use the Data Catalog APIs to collect, organize, find, access, understand, enrich, and activate technical, business, and operational metadata.
// For more information, see Data Catalog (https://docs.oracle.com/iaas/data-catalog/home.htm).
//

package datacatalog

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FacetedSearchAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingFacetedSearchAggregationPropertyTypeEnum(string(m.PropertyType)); !ok && m.PropertyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PropertyType: %s. Supported values are: %s.", m.PropertyType, strings.Join(GetFacetedSearchAggregationPropertyTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FacetedSearchAggregationPropertyTypeEnum Enum with underlying type: string
type FacetedSearchAggregationPropertyTypeEnum string

// Set of constants representing the allowable values for FacetedSearchAggregationPropertyTypeEnum
const (
	FacetedSearchAggregationPropertyTypeCustomProperty  FacetedSearchAggregationPropertyTypeEnum = "CUSTOM_PROPERTY"
	FacetedSearchAggregationPropertyTypeDefaultProperty FacetedSearchAggregationPropertyTypeEnum = "DEFAULT_PROPERTY"
)

var mappingFacetedSearchAggregationPropertyTypeEnum = map[string]FacetedSearchAggregationPropertyTypeEnum{
	"CUSTOM_PROPERTY":  FacetedSearchAggregationPropertyTypeCustomProperty,
	"DEFAULT_PROPERTY": FacetedSearchAggregationPropertyTypeDefaultProperty,
}

var mappingFacetedSearchAggregationPropertyTypeEnumLowerCase = map[string]FacetedSearchAggregationPropertyTypeEnum{
	"custom_property":  FacetedSearchAggregationPropertyTypeCustomProperty,
	"default_property": FacetedSearchAggregationPropertyTypeDefaultProperty,
}

// GetFacetedSearchAggregationPropertyTypeEnumValues Enumerates the set of values for FacetedSearchAggregationPropertyTypeEnum
func GetFacetedSearchAggregationPropertyTypeEnumValues() []FacetedSearchAggregationPropertyTypeEnum {
	values := make([]FacetedSearchAggregationPropertyTypeEnum, 0)
	for _, v := range mappingFacetedSearchAggregationPropertyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFacetedSearchAggregationPropertyTypeEnumStringValues Enumerates the set of values in String for FacetedSearchAggregationPropertyTypeEnum
func GetFacetedSearchAggregationPropertyTypeEnumStringValues() []string {
	return []string{
		"CUSTOM_PROPERTY",
		"DEFAULT_PROPERTY",
	}
}

// GetMappingFacetedSearchAggregationPropertyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFacetedSearchAggregationPropertyTypeEnum(val string) (FacetedSearchAggregationPropertyTypeEnum, bool) {
	enum, ok := mappingFacetedSearchAggregationPropertyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
