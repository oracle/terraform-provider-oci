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

// RemoveResourceLockDetails Used to remove a resource lock.
// Resource locks are used to prevent certain APIs from being called for the resource.
// A full lock prevents both updating the resource and deleting the resource. A delete
// lock prevents deleting the resource.
type RemoveResourceLockDetails struct {

	// Type of the lock.
	Type RemoveResourceLockDetailsTypeEnum `mandatory:"true" json:"type"`
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
