// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to monitor and manage resources such as
// Oracle Databases, MySQL Databases, and External Database Systems.
// For more information, see Database Management (https://docs.oracle.com/iaas/database-management/home.htm).
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociatedCloudComponent The details of the associated component.
type AssociatedCloudComponent struct {

	// The identifier of the associated component.
	ComponentId *string `mandatory:"true" json:"componentId"`

	// The association type.
	AssociationType AssociatedCloudComponentAssociationTypeEnum `mandatory:"true" json:"associationType"`

	// The type of associated component.
	ComponentType CloudDbSystemComponentTypeEnum `mandatory:"false" json:"componentType,omitempty"`
}

func (m AssociatedCloudComponent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AssociatedCloudComponent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAssociatedCloudComponentAssociationTypeEnum(string(m.AssociationType)); !ok && m.AssociationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssociationType: %s. Supported values are: %s.", m.AssociationType, strings.Join(GetAssociatedCloudComponentAssociationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCloudDbSystemComponentTypeEnum(string(m.ComponentType)); !ok && m.ComponentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComponentType: %s. Supported values are: %s.", m.ComponentType, strings.Join(GetCloudDbSystemComponentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AssociatedCloudComponentAssociationTypeEnum Enum with underlying type: string
type AssociatedCloudComponentAssociationTypeEnum string

// Set of constants representing the allowable values for AssociatedCloudComponentAssociationTypeEnum
const (
	AssociatedCloudComponentAssociationTypeContains AssociatedCloudComponentAssociationTypeEnum = "CONTAINS"
	AssociatedCloudComponentAssociationTypeUses     AssociatedCloudComponentAssociationTypeEnum = "USES"
)

var mappingAssociatedCloudComponentAssociationTypeEnum = map[string]AssociatedCloudComponentAssociationTypeEnum{
	"CONTAINS": AssociatedCloudComponentAssociationTypeContains,
	"USES":     AssociatedCloudComponentAssociationTypeUses,
}

var mappingAssociatedCloudComponentAssociationTypeEnumLowerCase = map[string]AssociatedCloudComponentAssociationTypeEnum{
	"contains": AssociatedCloudComponentAssociationTypeContains,
	"uses":     AssociatedCloudComponentAssociationTypeUses,
}

// GetAssociatedCloudComponentAssociationTypeEnumValues Enumerates the set of values for AssociatedCloudComponentAssociationTypeEnum
func GetAssociatedCloudComponentAssociationTypeEnumValues() []AssociatedCloudComponentAssociationTypeEnum {
	values := make([]AssociatedCloudComponentAssociationTypeEnum, 0)
	for _, v := range mappingAssociatedCloudComponentAssociationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAssociatedCloudComponentAssociationTypeEnumStringValues Enumerates the set of values in String for AssociatedCloudComponentAssociationTypeEnum
func GetAssociatedCloudComponentAssociationTypeEnumStringValues() []string {
	return []string{
		"CONTAINS",
		"USES",
	}
}

// GetMappingAssociatedCloudComponentAssociationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAssociatedCloudComponentAssociationTypeEnum(val string) (AssociatedCloudComponentAssociationTypeEnum, bool) {
	enum, ok := mappingAssociatedCloudComponentAssociationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
