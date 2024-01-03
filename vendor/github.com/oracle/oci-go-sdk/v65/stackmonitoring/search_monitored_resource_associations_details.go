// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SearchMonitoredResourceAssociationsDetails The information required to search monitored resource associations.
type SearchMonitoredResourceAssociationsDetails struct {

	// Compartment Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Source Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SourceResourceId *string `mandatory:"false" json:"sourceResourceId"`

	// Source Monitored Resource Name.
	SourceResourceName *string `mandatory:"false" json:"sourceResourceName"`

	// Source Monitored Resource Type.
	SourceResourceType *string `mandatory:"false" json:"sourceResourceType"`

	// Destination Monitored Resource Identifier OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	DestinationResourceId *string `mandatory:"false" json:"destinationResourceId"`

	// Source Monitored Resource Name.
	DestinationResourceName *string `mandatory:"false" json:"destinationResourceName"`

	// Source Monitored Resource Type.
	DestinationResourceType *string `mandatory:"false" json:"destinationResourceType"`

	// Association type filter to search associated resources.
	AssociationType *string `mandatory:"false" json:"associationType"`

	// The field to sort by. Only one sort order may be provided.
	// Default order for timeCreated is descending. Default order for assocType is descending.
	SortBy SearchMonitoredResourceAssociationsDetailsSortByEnum `mandatory:"false" json:"sortBy,omitempty"`

	// The sort order to use, either 'ASC' or 'DESC'.
	SortOrder SortOrderEnum `mandatory:"false" json:"sortOrder,omitempty"`
}

func (m SearchMonitoredResourceAssociationsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SearchMonitoredResourceAssociationsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSearchMonitoredResourceAssociationsDetailsSortByEnum(string(m.SortBy)); !ok && m.SortBy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortBy: %s. Supported values are: %s.", m.SortBy, strings.Join(GetSearchMonitoredResourceAssociationsDetailsSortByEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSortOrderEnum(string(m.SortOrder)); !ok && m.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", m.SortOrder, strings.Join(GetSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SearchMonitoredResourceAssociationsDetailsSortByEnum Enum with underlying type: string
type SearchMonitoredResourceAssociationsDetailsSortByEnum string

// Set of constants representing the allowable values for SearchMonitoredResourceAssociationsDetailsSortByEnum
const (
	SearchMonitoredResourceAssociationsDetailsSortByTimeCreated SearchMonitoredResourceAssociationsDetailsSortByEnum = "TIME_CREATED"
	SearchMonitoredResourceAssociationsDetailsSortByAssocType   SearchMonitoredResourceAssociationsDetailsSortByEnum = "ASSOC_TYPE"
)

var mappingSearchMonitoredResourceAssociationsDetailsSortByEnum = map[string]SearchMonitoredResourceAssociationsDetailsSortByEnum{
	"TIME_CREATED": SearchMonitoredResourceAssociationsDetailsSortByTimeCreated,
	"ASSOC_TYPE":   SearchMonitoredResourceAssociationsDetailsSortByAssocType,
}

var mappingSearchMonitoredResourceAssociationsDetailsSortByEnumLowerCase = map[string]SearchMonitoredResourceAssociationsDetailsSortByEnum{
	"time_created": SearchMonitoredResourceAssociationsDetailsSortByTimeCreated,
	"assoc_type":   SearchMonitoredResourceAssociationsDetailsSortByAssocType,
}

// GetSearchMonitoredResourceAssociationsDetailsSortByEnumValues Enumerates the set of values for SearchMonitoredResourceAssociationsDetailsSortByEnum
func GetSearchMonitoredResourceAssociationsDetailsSortByEnumValues() []SearchMonitoredResourceAssociationsDetailsSortByEnum {
	values := make([]SearchMonitoredResourceAssociationsDetailsSortByEnum, 0)
	for _, v := range mappingSearchMonitoredResourceAssociationsDetailsSortByEnum {
		values = append(values, v)
	}
	return values
}

// GetSearchMonitoredResourceAssociationsDetailsSortByEnumStringValues Enumerates the set of values in String for SearchMonitoredResourceAssociationsDetailsSortByEnum
func GetSearchMonitoredResourceAssociationsDetailsSortByEnumStringValues() []string {
	return []string{
		"TIME_CREATED",
		"ASSOC_TYPE",
	}
}

// GetMappingSearchMonitoredResourceAssociationsDetailsSortByEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSearchMonitoredResourceAssociationsDetailsSortByEnum(val string) (SearchMonitoredResourceAssociationsDetailsSortByEnum, bool) {
	enum, ok := mappingSearchMonitoredResourceAssociationsDetailsSortByEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
