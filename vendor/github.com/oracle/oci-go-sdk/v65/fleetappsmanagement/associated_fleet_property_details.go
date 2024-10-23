// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociatedFleetPropertyDetails The information about associated FleetProperty.
type AssociatedFleetPropertyDetails struct {

	// Tenancy OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Type of the FleetProperty.
	FleetPropertyType AssociatedFleetPropertyDetailsFleetPropertyTypeEnum `mandatory:"true" json:"fleetPropertyType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Value of the Property.
	Value *string `mandatory:"false" json:"value"`

	// Property is required or not.
	IsRequired *bool `mandatory:"false" json:"isRequired"`
}

func (m AssociatedFleetPropertyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedFleetPropertyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnum(string(m.FleetPropertyType)); !ok && m.FleetPropertyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FleetPropertyType: %s. Supported values are: %s.", m.FleetPropertyType, strings.Join(GetAssociatedFleetPropertyDetailsFleetPropertyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociatedFleetPropertyDetailsFleetPropertyTypeEnum Enum with underlying type: string
type AssociatedFleetPropertyDetailsFleetPropertyTypeEnum string

// Set of constants representing the allowable values for AssociatedFleetPropertyDetailsFleetPropertyTypeEnum
const (
	AssociatedFleetPropertyDetailsFleetPropertyTypeString AssociatedFleetPropertyDetailsFleetPropertyTypeEnum = "STRING"
	AssociatedFleetPropertyDetailsFleetPropertyTypeNumber AssociatedFleetPropertyDetailsFleetPropertyTypeEnum = "NUMBER"
)

var mappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnum = map[string]AssociatedFleetPropertyDetailsFleetPropertyTypeEnum{
	"STRING": AssociatedFleetPropertyDetailsFleetPropertyTypeString,
	"NUMBER": AssociatedFleetPropertyDetailsFleetPropertyTypeNumber,
}

var mappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnumLowerCase = map[string]AssociatedFleetPropertyDetailsFleetPropertyTypeEnum{
	"string": AssociatedFleetPropertyDetailsFleetPropertyTypeString,
	"number": AssociatedFleetPropertyDetailsFleetPropertyTypeNumber,
}

// GetAssociatedFleetPropertyDetailsFleetPropertyTypeEnumValues Enumerates the set of values for AssociatedFleetPropertyDetailsFleetPropertyTypeEnum
func GetAssociatedFleetPropertyDetailsFleetPropertyTypeEnumValues() []AssociatedFleetPropertyDetailsFleetPropertyTypeEnum {
	values := make([]AssociatedFleetPropertyDetailsFleetPropertyTypeEnum, 0)
	for _, v := range mappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedFleetPropertyDetailsFleetPropertyTypeEnumStringValues Enumerates the set of values in String for AssociatedFleetPropertyDetailsFleetPropertyTypeEnum
func GetAssociatedFleetPropertyDetailsFleetPropertyTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
	}
}

// GetMappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnum(val string) (AssociatedFleetPropertyDetailsFleetPropertyTypeEnum, bool) {
	enum, ok := mappingAssociatedFleetPropertyDetailsFleetPropertyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
