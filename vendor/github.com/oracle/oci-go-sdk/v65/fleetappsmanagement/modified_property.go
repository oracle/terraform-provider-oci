// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ModifiedProperty Modified Property represent the modification performed on the property
type ModifiedProperty struct {

	// Property name
	PropertyName *string `mandatory:"true" json:"propertyName"`

	// Modification performed on the property
	PropertyModification ModifiedPropertyPropertyModificationEnum `mandatory:"true" json:"propertyModification"`

	// Modification description. For example: changed From A - To B.
	ModificationDescription *string `mandatory:"false" json:"modificationDescription"`
}

func (m ModifiedProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ModifiedProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingModifiedPropertyPropertyModificationEnum(string(m.PropertyModification)); !ok && m.PropertyModification != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PropertyModification: %s. Supported values are: %s.", m.PropertyModification, strings.Join(GetModifiedPropertyPropertyModificationEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ModifiedPropertyPropertyModificationEnum Enum with underlying type: string
type ModifiedPropertyPropertyModificationEnum string

// Set of constants representing the allowable values for ModifiedPropertyPropertyModificationEnum
const (
	ModifiedPropertyPropertyModificationAdded    ModifiedPropertyPropertyModificationEnum = "ADDED"
	ModifiedPropertyPropertyModificationRemoved  ModifiedPropertyPropertyModificationEnum = "REMOVED"
	ModifiedPropertyPropertyModificationModified ModifiedPropertyPropertyModificationEnum = "MODIFIED"
)

var mappingModifiedPropertyPropertyModificationEnum = map[string]ModifiedPropertyPropertyModificationEnum{
	"ADDED":    ModifiedPropertyPropertyModificationAdded,
	"REMOVED":  ModifiedPropertyPropertyModificationRemoved,
	"MODIFIED": ModifiedPropertyPropertyModificationModified,
}

var mappingModifiedPropertyPropertyModificationEnumLowerCase = map[string]ModifiedPropertyPropertyModificationEnum{
	"added":    ModifiedPropertyPropertyModificationAdded,
	"removed":  ModifiedPropertyPropertyModificationRemoved,
	"modified": ModifiedPropertyPropertyModificationModified,
}

// GetModifiedPropertyPropertyModificationEnumValues Enumerates the set of values for ModifiedPropertyPropertyModificationEnum
func GetModifiedPropertyPropertyModificationEnumValues() []ModifiedPropertyPropertyModificationEnum {
	values := make([]ModifiedPropertyPropertyModificationEnum, 0)
	for _, v := range mappingModifiedPropertyPropertyModificationEnum {
		values = append(values, v)
	}
	return values
}

// GetModifiedPropertyPropertyModificationEnumStringValues Enumerates the set of values in String for ModifiedPropertyPropertyModificationEnum
func GetModifiedPropertyPropertyModificationEnumStringValues() []string {
	return []string{
		"ADDED",
		"REMOVED",
		"MODIFIED",
	}
}

// GetMappingModifiedPropertyPropertyModificationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingModifiedPropertyPropertyModificationEnum(val string) (ModifiedPropertyPropertyModificationEnum, bool) {
	enum, ok := mappingModifiedPropertyPropertyModificationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
