// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddPolicyLockDetails Request payload to add lock to the resource.
type AddPolicyLockDetails struct {

	// Type of the lock.
	Type AddPolicyLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The compartment ID of the lock.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock.
	RelatedResourceId *string `mandatory:"false" json:"relatedResourceId"`

	// A message added by the creator of the lock. This is typically used to give an
	// indication of why the resource is locked.
	Message *string `mandatory:"false" json:"message"`
}

func (m AddPolicyLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddPolicyLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAddPolicyLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAddPolicyLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AddPolicyLockDetailsTypeEnum Enum with underlying type: string
type AddPolicyLockDetailsTypeEnum string

// Set of constants representing the allowable values for AddPolicyLockDetailsTypeEnum
const (
	AddPolicyLockDetailsTypeFull   AddPolicyLockDetailsTypeEnum = "FULL"
	AddPolicyLockDetailsTypeDelete AddPolicyLockDetailsTypeEnum = "DELETE"
)

var mappingAddPolicyLockDetailsTypeEnum = map[string]AddPolicyLockDetailsTypeEnum{
	"FULL":   AddPolicyLockDetailsTypeFull,
	"DELETE": AddPolicyLockDetailsTypeDelete,
}

var mappingAddPolicyLockDetailsTypeEnumLowerCase = map[string]AddPolicyLockDetailsTypeEnum{
	"full":   AddPolicyLockDetailsTypeFull,
	"delete": AddPolicyLockDetailsTypeDelete,
}

// GetAddPolicyLockDetailsTypeEnumValues Enumerates the set of values for AddPolicyLockDetailsTypeEnum
func GetAddPolicyLockDetailsTypeEnumValues() []AddPolicyLockDetailsTypeEnum {
	values := make([]AddPolicyLockDetailsTypeEnum, 0)
	for _, v := range mappingAddPolicyLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAddPolicyLockDetailsTypeEnumStringValues Enumerates the set of values in String for AddPolicyLockDetailsTypeEnum
func GetAddPolicyLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingAddPolicyLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAddPolicyLockDetailsTypeEnum(val string) (AddPolicyLockDetailsTypeEnum, bool) {
	enum, ok := mappingAddPolicyLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
