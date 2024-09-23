// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetProperty Property associated with a Fleet.
type FleetProperty struct {

	// The unique id of the resource.
	Id *string `mandatory:"true" json:"id"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Text selection of the category
	SelectionType SelectionEnum `mandatory:"true" json:"selectionType"`

	// Format of the value
	ValueType ValueTypeEnum `mandatory:"true" json:"valueType"`

	// The current state of the FleetProperty.
	LifecycleState FleetPropertyLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Value of the Property
	Value *string `mandatory:"false" json:"value"`

	// Property Id Ocid.
	PropertyId *string `mandatory:"false" json:"propertyId"`

	// Values of the category (must be a single value if selection = 'single choice')
	AllowedValues []string `mandatory:"false" json:"allowedValues"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FleetProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSelectionEnum(string(m.SelectionType)); !ok && m.SelectionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SelectionType: %s. Supported values are: %s.", m.SelectionType, strings.Join(GetSelectionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetValueTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetPropertyLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetPropertyLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetPropertyLifecycleStateEnum Enum with underlying type: string
type FleetPropertyLifecycleStateEnum string

// Set of constants representing the allowable values for FleetPropertyLifecycleStateEnum
const (
	FleetPropertyLifecycleStateActive  FleetPropertyLifecycleStateEnum = "ACTIVE"
	FleetPropertyLifecycleStateDeleted FleetPropertyLifecycleStateEnum = "DELETED"
	FleetPropertyLifecycleStateFailed  FleetPropertyLifecycleStateEnum = "FAILED"
)

var mappingFleetPropertyLifecycleStateEnum = map[string]FleetPropertyLifecycleStateEnum{
	"ACTIVE":  FleetPropertyLifecycleStateActive,
	"DELETED": FleetPropertyLifecycleStateDeleted,
	"FAILED":  FleetPropertyLifecycleStateFailed,
}

var mappingFleetPropertyLifecycleStateEnumLowerCase = map[string]FleetPropertyLifecycleStateEnum{
	"active":  FleetPropertyLifecycleStateActive,
	"deleted": FleetPropertyLifecycleStateDeleted,
	"failed":  FleetPropertyLifecycleStateFailed,
}

// GetFleetPropertyLifecycleStateEnumValues Enumerates the set of values for FleetPropertyLifecycleStateEnum
func GetFleetPropertyLifecycleStateEnumValues() []FleetPropertyLifecycleStateEnum {
	values := make([]FleetPropertyLifecycleStateEnum, 0)
	for _, v := range mappingFleetPropertyLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetPropertyLifecycleStateEnumStringValues Enumerates the set of values in String for FleetPropertyLifecycleStateEnum
func GetFleetPropertyLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFleetPropertyLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetPropertyLifecycleStateEnum(val string) (FleetPropertyLifecycleStateEnum, bool) {
	enum, ok := mappingFleetPropertyLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
