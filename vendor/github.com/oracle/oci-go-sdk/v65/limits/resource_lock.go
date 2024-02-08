// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ResourceLock Resource locks prevent certain APIs from being called for the resource.
// A full lock prevents both updating and deleting the resource. A lock delete
// prevents deleting the resource.
type ResourceLock struct {

	// Lock type.
	Type ResourceLockTypeEnum `mandatory:"true" json:"type"`

	// The resource ID that is locking this resource. Indicates that deleting this resource removes the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the lock creator. The message typically gives an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`

	// Indicates when the lock was created, in the format defined by RFC 3339.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ResourceLock) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceLock) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResourceLockTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResourceLockTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ResourceLockTypeEnum Enum with underlying type: string
type ResourceLockTypeEnum string

// Set of constants representing the allowable values for ResourceLockTypeEnum
const (
	ResourceLockTypeFull   ResourceLockTypeEnum = "FULL"
	ResourceLockTypeDelete ResourceLockTypeEnum = "DELETE"
)

var mappingResourceLockTypeEnum = map[string]ResourceLockTypeEnum{
	"FULL":   ResourceLockTypeFull,
	"DELETE": ResourceLockTypeDelete,
}

var mappingResourceLockTypeEnumLowerCase = map[string]ResourceLockTypeEnum{
	"full":   ResourceLockTypeFull,
	"delete": ResourceLockTypeDelete,
}

// GetResourceLockTypeEnumValues Enumerates the set of values for ResourceLockTypeEnum
func GetResourceLockTypeEnumValues() []ResourceLockTypeEnum {
	values := make([]ResourceLockTypeEnum, 0)
	for _, v := range mappingResourceLockTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetResourceLockTypeEnumStringValues Enumerates the set of values in String for ResourceLockTypeEnum
func GetResourceLockTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingResourceLockTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingResourceLockTypeEnum(val string) (ResourceLockTypeEnum, bool) {
	enum, ok := mappingResourceLockTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
