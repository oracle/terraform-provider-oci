// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// FleetTarget A confirmed target within a fleet.
type FleetTarget struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Tenancy OCID
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Current version of target.
	Version *string `mandatory:"false" json:"version"`

	// Product to which the target belongs to.
	Product *string `mandatory:"false" json:"product"`

	Resource *TargetResource `mandatory:"false" json:"resource"`

	// The last known compliance state of the target.
	ComplianceState ComplianceStateEnum `mandatory:"false" json:"complianceState,omitempty"`

	// The time when the last successful discovery was made.
	TimeOfLastSuccessfulDiscovery *common.SDKTime `mandatory:"false" json:"timeOfLastSuccessfulDiscovery"`

	// The time when last discovery was attempted.
	TimeOfLastDiscoveryAttempt *common.SDKTime `mandatory:"false" json:"timeOfLastDiscoveryAttempt"`

	// A boolean flag that represents whether the last discovery attempt was successful.
	IsLastDiscoveryAttemptSuccessful *bool `mandatory:"false" json:"isLastDiscoveryAttemptSuccessful"`

	// The current state of the FleetTarget.
	LifecycleState FleetTargetLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m FleetTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingComplianceStateEnum(string(m.ComplianceState)); !ok && m.ComplianceState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ComplianceState: %s. Supported values are: %s.", m.ComplianceState, strings.Join(GetComplianceStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFleetTargetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFleetTargetLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FleetTargetLifecycleStateEnum Enum with underlying type: string
type FleetTargetLifecycleStateEnum string

// Set of constants representing the allowable values for FleetTargetLifecycleStateEnum
const (
	FleetTargetLifecycleStateActive  FleetTargetLifecycleStateEnum = "ACTIVE"
	FleetTargetLifecycleStateDeleted FleetTargetLifecycleStateEnum = "DELETED"
	FleetTargetLifecycleStateFailed  FleetTargetLifecycleStateEnum = "FAILED"
)

var mappingFleetTargetLifecycleStateEnum = map[string]FleetTargetLifecycleStateEnum{
	"ACTIVE":  FleetTargetLifecycleStateActive,
	"DELETED": FleetTargetLifecycleStateDeleted,
	"FAILED":  FleetTargetLifecycleStateFailed,
}

var mappingFleetTargetLifecycleStateEnumLowerCase = map[string]FleetTargetLifecycleStateEnum{
	"active":  FleetTargetLifecycleStateActive,
	"deleted": FleetTargetLifecycleStateDeleted,
	"failed":  FleetTargetLifecycleStateFailed,
}

// GetFleetTargetLifecycleStateEnumValues Enumerates the set of values for FleetTargetLifecycleStateEnum
func GetFleetTargetLifecycleStateEnumValues() []FleetTargetLifecycleStateEnum {
	values := make([]FleetTargetLifecycleStateEnum, 0)
	for _, v := range mappingFleetTargetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFleetTargetLifecycleStateEnumStringValues Enumerates the set of values in String for FleetTargetLifecycleStateEnum
func GetFleetTargetLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFleetTargetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFleetTargetLifecycleStateEnum(val string) (FleetTargetLifecycleStateEnum, bool) {
	enum, ok := mappingFleetTargetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
