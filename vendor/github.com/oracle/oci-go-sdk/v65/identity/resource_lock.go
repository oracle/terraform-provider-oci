// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, policies, and identity domains.
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceLock Resource locks are used to prevent certain APIs from being called for the resource.
// A full lock prevents both updating the resource and deleting the resource. A delete
// lock prevents deleting the resource.
type ResourceLock struct {

	// Type of the lock.
	Type ResourceLockTypeEnum `mandatory:"true" json:"type"`

	// The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`

	// When the lock was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Indicates if the lock is active or not. For example, if there are mutliple FULL locks, the first-created FULL lock will be effective.
	IsActive *bool `mandatory:"false" json:"isActive"`
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
