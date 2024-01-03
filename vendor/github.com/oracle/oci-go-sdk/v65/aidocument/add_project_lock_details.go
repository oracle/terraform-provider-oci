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

// AddProjectLockDetails Request payload to add lock to the resource.
type AddProjectLockDetails struct {

	// Type of the lock.
	Type AddProjectLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`
}

func (m AddProjectLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddProjectLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddProjectLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddProjectLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddProjectLockDetailsTypeEnum Enum with underlying type: string
type AddProjectLockDetailsTypeEnum string

// Set of constants representing the allowable values for AddProjectLockDetailsTypeEnum
const (
	AddProjectLockDetailsTypeFull   AddProjectLockDetailsTypeEnum = "FULL"
	AddProjectLockDetailsTypeDelete AddProjectLockDetailsTypeEnum = "DELETE"
)

var mappingAddProjectLockDetailsTypeEnum = map[string]AddProjectLockDetailsTypeEnum{
	"FULL":   AddProjectLockDetailsTypeFull,
	"DELETE": AddProjectLockDetailsTypeDelete,
}

var mappingAddProjectLockDetailsTypeEnumLowerCase = map[string]AddProjectLockDetailsTypeEnum{
	"full":   AddProjectLockDetailsTypeFull,
	"delete": AddProjectLockDetailsTypeDelete,
}

// GetAddProjectLockDetailsTypeEnumValues Enumerates the set of values for AddProjectLockDetailsTypeEnum
func GetAddProjectLockDetailsTypeEnumValues() []AddProjectLockDetailsTypeEnum {
	values := make([]AddProjectLockDetailsTypeEnum, 0)
	for _, v := range mappingAddProjectLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddProjectLockDetailsTypeEnumStringValues Enumerates the set of values in String for AddProjectLockDetailsTypeEnum
func GetAddProjectLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingAddProjectLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddProjectLockDetailsTypeEnum(val string) (AddProjectLockDetailsTypeEnum, bool) {
	enum, ok := mappingAddProjectLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
