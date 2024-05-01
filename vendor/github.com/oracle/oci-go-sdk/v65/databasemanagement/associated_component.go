// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.cloud.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociatedComponent The details of the associated component.
type AssociatedComponent struct {

	// The identifier of the associated component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The association type.
	AssociationType AssociatedComponentAssociationTypeEnum `mandatory:"true" json:"associationType"`

	// The type of associated component.
	ComponentType ExternalDbSystemComponentTypeEnum `mandatory:"false" json:"componentType,omitempty"`
}

func (m AssociatedComponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedComponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssociatedComponentAssociationTypeEnum(string(m.AssociationType)); !ok && m.AssociationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationType: %s. Supported values are: %s.", m.AssociationType, strings.Join(GetAssociatedComponentAssociationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingExternalDbSystemComponentTypeEnum(string(m.ComponentType)); !ok && m.ComponentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComponentType: %s. Supported values are: %s.", m.ComponentType, strings.Join(GetExternalDbSystemComponentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociatedComponentAssociationTypeEnum Enum with underlying type: string
type AssociatedComponentAssociationTypeEnum string

// Set of constants representing the allowable values for AssociatedComponentAssociationTypeEnum
const (
	AssociatedComponentAssociationTypeContains AssociatedComponentAssociationTypeEnum = "CONTAINS"
	AssociatedComponentAssociationTypeUses     AssociatedComponentAssociationTypeEnum = "USES"
)

var mappingAssociatedComponentAssociationTypeEnum = map[string]AssociatedComponentAssociationTypeEnum{
	"CONTAINS": AssociatedComponentAssociationTypeContains,
	"USES":     AssociatedComponentAssociationTypeUses,
}

var mappingAssociatedComponentAssociationTypeEnumLowerCase = map[string]AssociatedComponentAssociationTypeEnum{
	"contains": AssociatedComponentAssociationTypeContains,
	"uses":     AssociatedComponentAssociationTypeUses,
}

// GetAssociatedComponentAssociationTypeEnumValues Enumerates the set of values for AssociatedComponentAssociationTypeEnum
func GetAssociatedComponentAssociationTypeEnumValues() []AssociatedComponentAssociationTypeEnum {
	values := make([]AssociatedComponentAssociationTypeEnum, 0)
	for _, v := range mappingAssociatedComponentAssociationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedComponentAssociationTypeEnumStringValues Enumerates the set of values in String for AssociatedComponentAssociationTypeEnum
func GetAssociatedComponentAssociationTypeEnumStringValues() []string {
	return []string{
		"CONTAINS",
		"USES",
	}
}

// GetMappingAssociatedComponentAssociationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedComponentAssociationTypeEnum(val string) (AssociatedComponentAssociationTypeEnum, bool) {
	enum, ok := mappingAssociatedComponentAssociationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
