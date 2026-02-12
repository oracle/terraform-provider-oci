// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DependentPropertyValue Dependent values of a property.
type DependentPropertyValue struct {

	// Dependent property value ID.
	Id *string `mandatory:"true" json:"id"`

	// compartment OCID
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	KeyValueMap *DependentPropertyKeyValueMap `mandatory:"true" json:"keyValueMap"`

	// The current state of the Dependent property value.
	LifecycleState DependentPropertyValueLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`
}

func (m DependentPropertyValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DependentPropertyValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDependentPropertyValueLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDependentPropertyValueLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DependentPropertyValueLifecycleStateEnum Enum with underlying type: string
type DependentPropertyValueLifecycleStateEnum string

// Set of constants representing the allowable values for DependentPropertyValueLifecycleStateEnum
const (
	DependentPropertyValueLifecycleStateActive   DependentPropertyValueLifecycleStateEnum = "ACTIVE"
	DependentPropertyValueLifecycleStateDeleted  DependentPropertyValueLifecycleStateEnum = "DELETED"
	DependentPropertyValueLifecycleStateFailed   DependentPropertyValueLifecycleStateEnum = "FAILED"
	DependentPropertyValueLifecycleStateUpdating DependentPropertyValueLifecycleStateEnum = "UPDATING"
)

var mappingDependentPropertyValueLifecycleStateEnum = map[string]DependentPropertyValueLifecycleStateEnum{
	"ACTIVE":   DependentPropertyValueLifecycleStateActive,
	"DELETED":  DependentPropertyValueLifecycleStateDeleted,
	"FAILED":   DependentPropertyValueLifecycleStateFailed,
	"UPDATING": DependentPropertyValueLifecycleStateUpdating,
}

var mappingDependentPropertyValueLifecycleStateEnumLowerCase = map[string]DependentPropertyValueLifecycleStateEnum{
	"active":   DependentPropertyValueLifecycleStateActive,
	"deleted":  DependentPropertyValueLifecycleStateDeleted,
	"failed":   DependentPropertyValueLifecycleStateFailed,
	"updating": DependentPropertyValueLifecycleStateUpdating,
}

// GetDependentPropertyValueLifecycleStateEnumValues Enumerates the set of values for DependentPropertyValueLifecycleStateEnum
func GetDependentPropertyValueLifecycleStateEnumValues() []DependentPropertyValueLifecycleStateEnum {
	values := make([]DependentPropertyValueLifecycleStateEnum, 0)
	for _, v := range mappingDependentPropertyValueLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDependentPropertyValueLifecycleStateEnumStringValues Enumerates the set of values in String for DependentPropertyValueLifecycleStateEnum
func GetDependentPropertyValueLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
		"UPDATING",
	}
}

// GetMappingDependentPropertyValueLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDependentPropertyValueLifecycleStateEnum(val string) (DependentPropertyValueLifecycleStateEnum, bool) {
	enum, ok := mappingDependentPropertyValueLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
