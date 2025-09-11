// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Email Delivery API
//
// Use the Email Delivery API to do the necessary set up to send high-volume and application-generated emails through the OCI Email Delivery service.
// For more information, see Overview of the Email Delivery Service (https://docs.oracle.com/iaas/Content/Email/Concepts/overview.htm).
//  **Note:** Write actions (POST, UPDATE, DELETE) may take several minutes to propagate and be reflected by the API.
//  If a subsequent read request fails to reflect your changes, wait a few minutes and try again.
//

package email

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddLockDetails Request payload to add lock to the resource.
type AddLockDetails struct {

	// Lock type.
	Type AddLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The lock compartment ID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

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
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
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
