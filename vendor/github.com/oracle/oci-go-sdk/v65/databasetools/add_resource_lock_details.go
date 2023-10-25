// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// AddResourceLockDetails The representation of AddResourceLockDetails
type AddResourceLockDetails struct {

	// Type of the lock.
	Type AddResourceLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`

	// When the lock was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m AddResourceLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddResourceLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddResourceLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddResourceLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddResourceLockDetailsTypeEnum Enum with underlying type: string
type AddResourceLockDetailsTypeEnum string

// Set of constants representing the allowable values for AddResourceLockDetailsTypeEnum
const (
	AddResourceLockDetailsTypeFull   AddResourceLockDetailsTypeEnum = "FULL"
	AddResourceLockDetailsTypeDelete AddResourceLockDetailsTypeEnum = "DELETE"
)

var mappingAddResourceLockDetailsTypeEnum = map[string]AddResourceLockDetailsTypeEnum{
	"FULL":   AddResourceLockDetailsTypeFull,
	"DELETE": AddResourceLockDetailsTypeDelete,
}

var mappingAddResourceLockDetailsTypeEnumLowerCase = map[string]AddResourceLockDetailsTypeEnum{
	"full":   AddResourceLockDetailsTypeFull,
	"delete": AddResourceLockDetailsTypeDelete,
}

// GetAddResourceLockDetailsTypeEnumValues Enumerates the set of values for AddResourceLockDetailsTypeEnum
func GetAddResourceLockDetailsTypeEnumValues() []AddResourceLockDetailsTypeEnum {
	values := make([]AddResourceLockDetailsTypeEnum, 0)
	for _, v := range mappingAddResourceLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddResourceLockDetailsTypeEnumStringValues Enumerates the set of values in String for AddResourceLockDetailsTypeEnum
func GetAddResourceLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingAddResourceLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddResourceLockDetailsTypeEnum(val string) (AddResourceLockDetailsTypeEnum, bool) {
	enum, ok := mappingAddResourceLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
