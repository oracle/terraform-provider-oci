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

// AddModelLockDetails Request payload to add lock to the resource.
type AddModelLockDetails struct {

	// Type of the lock.
	Type AddModelLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`
}

func (m AddModelLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddModelLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddModelLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddModelLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddModelLockDetailsTypeEnum Enum with underlying type: string
type AddModelLockDetailsTypeEnum string

// Set of constants representing the allowable values for AddModelLockDetailsTypeEnum
const (
	AddModelLockDetailsTypeFull   AddModelLockDetailsTypeEnum = "FULL"
	AddModelLockDetailsTypeDelete AddModelLockDetailsTypeEnum = "DELETE"
)

var mappingAddModelLockDetailsTypeEnum = map[string]AddModelLockDetailsTypeEnum{
	"FULL":   AddModelLockDetailsTypeFull,
	"DELETE": AddModelLockDetailsTypeDelete,
}

var mappingAddModelLockDetailsTypeEnumLowerCase = map[string]AddModelLockDetailsTypeEnum{
	"full":   AddModelLockDetailsTypeFull,
	"delete": AddModelLockDetailsTypeDelete,
}

// GetAddModelLockDetailsTypeEnumValues Enumerates the set of values for AddModelLockDetailsTypeEnum
func GetAddModelLockDetailsTypeEnumValues() []AddModelLockDetailsTypeEnum {
	values := make([]AddModelLockDetailsTypeEnum, 0)
	for _, v := range mappingAddModelLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddModelLockDetailsTypeEnumStringValues Enumerates the set of values in String for AddModelLockDetailsTypeEnum
func GetAddModelLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingAddModelLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddModelLockDetailsTypeEnum(val string) (AddModelLockDetailsTypeEnum, bool) {
	enum, ok := mappingAddModelLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
