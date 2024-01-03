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

// RemoveModelLockDetails Request payload to remove lock to the resource.
type RemoveModelLockDetails struct {

	// Type of the lock.
	Type RemoveModelLockDetailsTypeEnum `mandatory:"true" json:"type"`
}

func (m RemoveModelLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveModelLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveModelLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveModelLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveModelLockDetailsTypeEnum Enum with underlying type: string
type RemoveModelLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveModelLockDetailsTypeEnum
const (
	RemoveModelLockDetailsTypeFull   RemoveModelLockDetailsTypeEnum = "FULL"
	RemoveModelLockDetailsTypeDelete RemoveModelLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveModelLockDetailsTypeEnum = map[string]RemoveModelLockDetailsTypeEnum{
	"FULL":   RemoveModelLockDetailsTypeFull,
	"DELETE": RemoveModelLockDetailsTypeDelete,
}

var mappingRemoveModelLockDetailsTypeEnumLowerCase = map[string]RemoveModelLockDetailsTypeEnum{
	"full":   RemoveModelLockDetailsTypeFull,
	"delete": RemoveModelLockDetailsTypeDelete,
}

// GetRemoveModelLockDetailsTypeEnumValues Enumerates the set of values for RemoveModelLockDetailsTypeEnum
func GetRemoveModelLockDetailsTypeEnumValues() []RemoveModelLockDetailsTypeEnum {
	values := make([]RemoveModelLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveModelLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveModelLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveModelLockDetailsTypeEnum
func GetRemoveModelLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveModelLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveModelLockDetailsTypeEnum(val string) (RemoveModelLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveModelLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
