// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SortField Field outlining queryString sort command fields and their corresponding sort order.
type SortField struct {

	// Field display name - will be alias if field is renamed by queryStrng.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Field denoting if this is a declaration of the field in the queryString.
	IsDeclared *bool `mandatory:"false" json:"isDeclared"`

	// Same as displayName unless field renamed in which case this will hold the original display names for the field
	// across all renames.
	OriginalDisplayNames []string `mandatory:"false" json:"originalDisplayNames"`

	// Internal identifier for the field.
	InternalName *string `mandatory:"false" json:"internalName"`

	// Identifies if this field can be used as a grouping field in any grouping command.
	IsGroupable *bool `mandatory:"false" json:"isGroupable"`

	// Identifies if this field format is a duration.
	IsDuration *bool `mandatory:"false" json:"isDuration"`

	// Alias of field if renamed by queryStrng.
	Alias *string `mandatory:"false" json:"alias"`

	// Query used to derive this field if specified.
	FilterQueryString *string `mandatory:"false" json:"filterQueryString"`

	// Field denoting field unit type.
	UnitType *string `mandatory:"false" json:"unitType"`

	// Sort order for the field specified in the queryString.
	Direction SortFieldDirectionEnum `mandatory:"false" json:"direction,omitempty"`

	// Field denoting field data type.
	ValueType ValueTypeEnum `mandatory:"false" json:"valueType,omitempty"`
}

//GetDisplayName returns DisplayName
func (m SortField) GetDisplayName() *string {
	return m.DisplayName
}

//GetIsDeclared returns IsDeclared
func (m SortField) GetIsDeclared() *bool {
	return m.IsDeclared
}

//GetOriginalDisplayNames returns OriginalDisplayNames
func (m SortField) GetOriginalDisplayNames() []string {
	return m.OriginalDisplayNames
}

//GetInternalName returns InternalName
func (m SortField) GetInternalName() *string {
	return m.InternalName
}

//GetValueType returns ValueType
func (m SortField) GetValueType() ValueTypeEnum {
	return m.ValueType
}

//GetIsGroupable returns IsGroupable
func (m SortField) GetIsGroupable() *bool {
	return m.IsGroupable
}

//GetIsDuration returns IsDuration
func (m SortField) GetIsDuration() *bool {
	return m.IsDuration
}

//GetAlias returns Alias
func (m SortField) GetAlias() *string {
	return m.Alias
}

//GetFilterQueryString returns FilterQueryString
func (m SortField) GetFilterQueryString() *string {
	return m.FilterQueryString
}

//GetUnitType returns UnitType
func (m SortField) GetUnitType() *string {
	return m.UnitType
}

func (m SortField) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SortField) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSortFieldDirectionEnum(string(m.Direction)); !ok && m.Direction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Direction: %s. Supported values are: %s.", m.Direction, strings.Join(GetSortFieldDirectionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m SortField) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeSortField SortField
	s := struct {
		DiscriminatorParam string `json:"name"`
		MarshalTypeSortField
	}{
		"SORT",
		(MarshalTypeSortField)(m),
	}

	return json.Marshal(&s)
}

// SortFieldDirectionEnum Enum with underlying type: string
type SortFieldDirectionEnum string

// Set of constants representing the allowable values for SortFieldDirectionEnum
const (
	SortFieldDirectionAscending  SortFieldDirectionEnum = "ASCENDING"
	SortFieldDirectionDescending SortFieldDirectionEnum = "DESCENDING"
)

var mappingSortFieldDirectionEnum = map[string]SortFieldDirectionEnum{
	"ASCENDING":  SortFieldDirectionAscending,
	"DESCENDING": SortFieldDirectionDescending,
}

// GetSortFieldDirectionEnumValues Enumerates the set of values for SortFieldDirectionEnum
func GetSortFieldDirectionEnumValues() []SortFieldDirectionEnum {
	values := make([]SortFieldDirectionEnum, 0)
	for _, v := range mappingSortFieldDirectionEnum {
		values = append(values, v)
	}
	return values
}

// GetSortFieldDirectionEnumStringValues Enumerates the set of values in String for SortFieldDirectionEnum
func GetSortFieldDirectionEnumStringValues() []string {
	return []string{
		"ASCENDING",
		"DESCENDING",
	}
}

// GetMappingSortFieldDirectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSortFieldDirectionEnum(val string) (SortFieldDirectionEnum, bool) {
	mappingSortFieldDirectionEnumIgnoreCase := make(map[string]SortFieldDirectionEnum)
	for k, v := range mappingSortFieldDirectionEnum {
		mappingSortFieldDirectionEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingSortFieldDirectionEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
