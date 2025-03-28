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

// AddLockDetails Request payload to add lock to the resource. The FULL lock type allows no modifications (delete, create, update).
//
//	The DELETE lock type allows all modifications, but delete is not allowed.
type AddLockDetails struct {

	// Lock type.
	Type AddLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The resource ID that is locking this resource. Indicates that deleting this resource removes the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the lock creator. The message typically gives an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`
}

func (m AddLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddLockDetailsTypeEnum Enum with underlying type: string
type AddLockDetailsTypeEnum string

// Set of constants representing the allowable values for AddLockDetailsTypeEnum
const (
	AddLockDetailsTypeFull   AddLockDetailsTypeEnum = "FULL"
	AddLockDetailsTypeDelete AddLockDetailsTypeEnum = "DELETE"
)

var mappingAddLockDetailsTypeEnum = map[string]AddLockDetailsTypeEnum{
	"FULL":   AddLockDetailsTypeFull,
	"DELETE": AddLockDetailsTypeDelete,
}

var mappingAddLockDetailsTypeEnumLowerCase = map[string]AddLockDetailsTypeEnum{
	"full":   AddLockDetailsTypeFull,
	"delete": AddLockDetailsTypeDelete,
}

// GetAddLockDetailsTypeEnumValues Enumerates the set of values for AddLockDetailsTypeEnum
func GetAddLockDetailsTypeEnumValues() []AddLockDetailsTypeEnum {
	values := make([]AddLockDetailsTypeEnum, 0)
	for _, v := range mappingAddLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddLockDetailsTypeEnumStringValues Enumerates the set of values in String for AddLockDetailsTypeEnum
func GetAddLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingAddLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddLockDetailsTypeEnum(val string) (AddLockDetailsTypeEnum, bool) {
	enum, ok := mappingAddLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
