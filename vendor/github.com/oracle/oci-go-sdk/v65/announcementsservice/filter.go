// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Announcements Service API
//
// Manage Oracle Cloud Infrastructure console announcements.
//

package announcementsservice

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Filter Criteria that the Announcements service uses to match announcements so it can provide only desired announcements to subscribers.
type Filter struct {

	// The type of filter. You cannot combine the RESOURCE_ID filter with any other type of filter within a given filter group. For filter types that support multiple values, specify the values individually.
	Type FilterTypeEnum `mandatory:"true" json:"type"`

	// The value of the filter.
	Value *string `mandatory:"true" json:"value"`
}

func (m Filter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Filter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFilterTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetFilterTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FilterTypeEnum Enum with underlying type: string
type FilterTypeEnum string

// Set of constants representing the allowable values for FilterTypeEnum
const (
	FilterTypeCompartmentId    FilterTypeEnum = "COMPARTMENT_ID"
	FilterTypePlatformType     FilterTypeEnum = "PLATFORM_TYPE"
	FilterTypeRegion           FilterTypeEnum = "REGION"
	FilterTypeService          FilterTypeEnum = "SERVICE"
	FilterTypeResourceId       FilterTypeEnum = "RESOURCE_ID"
	FilterTypeAnnouncementType FilterTypeEnum = "ANNOUNCEMENT_TYPE"
)

var mappingFilterTypeEnum = map[string]FilterTypeEnum{
	"COMPARTMENT_ID":    FilterTypeCompartmentId,
	"PLATFORM_TYPE":     FilterTypePlatformType,
	"REGION":            FilterTypeRegion,
	"SERVICE":           FilterTypeService,
	"RESOURCE_ID":       FilterTypeResourceId,
	"ANNOUNCEMENT_TYPE": FilterTypeAnnouncementType,
}

var mappingFilterTypeEnumLowerCase = map[string]FilterTypeEnum{
	"compartment_id":    FilterTypeCompartmentId,
	"platform_type":     FilterTypePlatformType,
	"region":            FilterTypeRegion,
	"service":           FilterTypeService,
	"resource_id":       FilterTypeResourceId,
	"announcement_type": FilterTypeAnnouncementType,
}

// GetFilterTypeEnumValues Enumerates the set of values for FilterTypeEnum
func GetFilterTypeEnumValues() []FilterTypeEnum {
	values := make([]FilterTypeEnum, 0)
	for _, v := range mappingFilterTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetFilterTypeEnumStringValues Enumerates the set of values in String for FilterTypeEnum
func GetFilterTypeEnumStringValues() []string {
	return []string{
		"COMPARTMENT_ID",
		"PLATFORM_TYPE",
		"REGION",
		"SERVICE",
		"RESOURCE_ID",
		"ANNOUNCEMENT_TYPE",
	}
}

// GetMappingFilterTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFilterTypeEnum(val string) (FilterTypeEnum, bool) {
	enum, ok := mappingFilterTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
