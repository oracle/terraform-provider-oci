// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Limits APIs
//
// APIs that interact with the resource limits of a specific resource type.
//

package limits

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveLockDetails Request payload to remove the resource lock.
type RemoveLockDetails struct {

	// Lock type.
	Type RemoveLockDetailsTypeEnum `mandatory:"true" json:"type"`
}

func (m RemoveLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveLockDetailsTypeEnum Enum with underlying type: string
type RemoveLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveLockDetailsTypeEnum
const (
	RemoveLockDetailsTypeFull   RemoveLockDetailsTypeEnum = "FULL"
	RemoveLockDetailsTypeDelete RemoveLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveLockDetailsTypeEnum = map[string]RemoveLockDetailsTypeEnum{
	"FULL":   RemoveLockDetailsTypeFull,
	"DELETE": RemoveLockDetailsTypeDelete,
}

var mappingRemoveLockDetailsTypeEnumLowerCase = map[string]RemoveLockDetailsTypeEnum{
	"full":   RemoveLockDetailsTypeFull,
	"delete": RemoveLockDetailsTypeDelete,
}

// GetRemoveLockDetailsTypeEnumValues Enumerates the set of values for RemoveLockDetailsTypeEnum
func GetRemoveLockDetailsTypeEnumValues() []RemoveLockDetailsTypeEnum {
	values := make([]RemoveLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveLockDetailsTypeEnum
func GetRemoveLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveLockDetailsTypeEnum(val string) (RemoveLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
