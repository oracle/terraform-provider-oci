// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard and Security Zones API
//
// Use the Cloud Guard and Security Zones API to automate processes that you would otherwise perform through the Cloud Guard Console or the Security Zones Console. For more information on these services, see the Cloud Guard (https://docs.oracle.com/iaas/cloud-guard/home.htm) and Security Zones (https://docs.oracle.com/iaas/security-zone/home.htm) documentation.
// **Note:** For Cloud Guard, you can perform Create, Update, and Delete operations only from the reporting region of your Cloud Guard tenancy. You can perform Read operations from any region.
//

package cloudguard

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RemoveGovernanceTargetLockDetails Request payload to remove lock to the resource.
type RemoveGovernanceTargetLockDetails struct {

	// Type of the lock.
	Type RemoveGovernanceTargetLockDetailsTypeEnum `mandatory:"true" json:"type"`

	// The compartment ID of the lock.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`
}

func (m RemoveGovernanceTargetLockDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RemoveGovernanceTargetLockDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRemoveGovernanceTargetLockDetailsTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetRemoveGovernanceTargetLockDetailsTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RemoveGovernanceTargetLockDetailsTypeEnum Enum with underlying type: string
type RemoveGovernanceTargetLockDetailsTypeEnum string

// Set of constants representing the allowable values for RemoveGovernanceTargetLockDetailsTypeEnum
const (
	RemoveGovernanceTargetLockDetailsTypeFull   RemoveGovernanceTargetLockDetailsTypeEnum = "FULL"
	RemoveGovernanceTargetLockDetailsTypeDelete RemoveGovernanceTargetLockDetailsTypeEnum = "DELETE"
)

var mappingRemoveGovernanceTargetLockDetailsTypeEnum = map[string]RemoveGovernanceTargetLockDetailsTypeEnum{
	"FULL":   RemoveGovernanceTargetLockDetailsTypeFull,
	"DELETE": RemoveGovernanceTargetLockDetailsTypeDelete,
}

var mappingRemoveGovernanceTargetLockDetailsTypeEnumLowerCase = map[string]RemoveGovernanceTargetLockDetailsTypeEnum{
	"full":   RemoveGovernanceTargetLockDetailsTypeFull,
	"delete": RemoveGovernanceTargetLockDetailsTypeDelete,
}

// GetRemoveGovernanceTargetLockDetailsTypeEnumValues Enumerates the set of values for RemoveGovernanceTargetLockDetailsTypeEnum
func GetRemoveGovernanceTargetLockDetailsTypeEnumValues() []RemoveGovernanceTargetLockDetailsTypeEnum {
	values := make([]RemoveGovernanceTargetLockDetailsTypeEnum, 0)
	for _, v := range mappingRemoveGovernanceTargetLockDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRemoveGovernanceTargetLockDetailsTypeEnumStringValues Enumerates the set of values in String for RemoveGovernanceTargetLockDetailsTypeEnum
func GetRemoveGovernanceTargetLockDetailsTypeEnumStringValues() []string {
	return []string{
		"FULL",
		"DELETE",
	}
}

// GetMappingRemoveGovernanceTargetLockDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRemoveGovernanceTargetLockDetailsTypeEnum(val string) (RemoveGovernanceTargetLockDetailsTypeEnum, bool) {
	enum, ok := mappingRemoveGovernanceTargetLockDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
