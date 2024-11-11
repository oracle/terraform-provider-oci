// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddResourceLockDetails Used to add a resource lock.
// Resource locks are used to prevent certain APIs from being called for the resource.
// A full lock prevents both updating the resource and deleting the resource. A delete
// lock prevents deleting the resource.
type AddResourceLockDetails struct {

	// Type of the lock.
	Type AddResourceLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`
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
