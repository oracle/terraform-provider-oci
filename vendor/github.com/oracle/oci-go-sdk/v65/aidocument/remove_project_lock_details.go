// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveProjectLockDetails Request payload to remove lock to the resource.
type RemoveProjectLockDetails struct {

	// Type of the lock.
	Type RemoveProjectLockDetailsTypeEnum `mandatory:"true" json:"type"`
}

func (m RemoveProjectLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveProjectLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveProjectLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveProjectLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveProjectLockDetailsTypeEnum Enum with underlying type: string
type RemoveProjectLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveProjectLockDetailsTypeEnum
const (
	RemoveProjectLockDetailsTypeFull   RemoveProjectLockDetailsTypeEnum = "FULL"
	RemoveProjectLockDetailsTypeDelete RemoveProjectLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveProjectLockDetailsTypeEnum = map[string]RemoveProjectLockDetailsTypeEnum{
	"FULL":   RemoveProjectLockDetailsTypeFull,
	"DELETE": RemoveProjectLockDetailsTypeDelete,
}

var mappingRemoveProjectLockDetailsTypeEnumLowerCase = map[string]RemoveProjectLockDetailsTypeEnum{
	"full":   RemoveProjectLockDetailsTypeFull,
	"delete": RemoveProjectLockDetailsTypeDelete,
}

// GetRemoveProjectLockDetailsTypeEnumValues Enumerates the set of values for RemoveProjectLockDetailsTypeEnum
func GetRemoveProjectLockDetailsTypeEnumValues() []RemoveProjectLockDetailsTypeEnum {
	values := make([]RemoveProjectLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveProjectLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveProjectLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveProjectLockDetailsTypeEnum
func GetRemoveProjectLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveProjectLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveProjectLockDetailsTypeEnum(val string) (RemoveProjectLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveProjectLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
