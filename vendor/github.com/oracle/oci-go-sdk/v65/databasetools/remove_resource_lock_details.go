// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveResourceLockDetails The representation of RemoveResourceLockDetails
type RemoveResourceLockDetails struct {

	// Type of the lock.
	Type RemoveResourceLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`

	// When the lock was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m RemoveResourceLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveResourceLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveResourceLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveResourceLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveResourceLockDetailsTypeEnum Enum with underlying type: string
type RemoveResourceLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveResourceLockDetailsTypeEnum
const (
	RemoveResourceLockDetailsTypeFull   RemoveResourceLockDetailsTypeEnum = "FULL"
	RemoveResourceLockDetailsTypeDelete RemoveResourceLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveResourceLockDetailsTypeEnum = map[string]RemoveResourceLockDetailsTypeEnum{
	"FULL":   RemoveResourceLockDetailsTypeFull,
	"DELETE": RemoveResourceLockDetailsTypeDelete,
}

var mappingRemoveResourceLockDetailsTypeEnumLowerCase = map[string]RemoveResourceLockDetailsTypeEnum{
	"full":   RemoveResourceLockDetailsTypeFull,
	"delete": RemoveResourceLockDetailsTypeDelete,
}

// GetRemoveResourceLockDetailsTypeEnumValues Enumerates the set of values for RemoveResourceLockDetailsTypeEnum
func GetRemoveResourceLockDetailsTypeEnumValues() []RemoveResourceLockDetailsTypeEnum {
	values := make([]RemoveResourceLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveResourceLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveResourceLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveResourceLockDetailsTypeEnum
func GetRemoveResourceLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveResourceLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveResourceLockDetailsTypeEnum(val string) (RemoveResourceLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveResourceLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
