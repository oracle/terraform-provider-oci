// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.cloud.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CccUpgradeSchedule Defines a schedule for preferred upgrade times.
type CccUpgradeSchedule struct {

	// Upgrade schedule OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// This cannot be changed once created.
	Id *string `mandatory:"true" json:"id"`

	// Compute Cloud@Customer upgrade schedule display name.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the
	// Compute Cloud@Customer upgrade schedule.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The time the upgrade schedule was created, using an RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Lifecycle state of the resource.
	LifecycleState CccUpgradeScheduleLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An optional description of the Compute Cloud@Customer upgrade schedule.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// The time the upgrade schedule was updated, using an RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail.
	// For example, the message can be used to provide actionable information for a resource in
	// a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// List of preferred times for Compute Cloud@Customer infrastructures associated with this
	// schedule to be upgraded.
	Events []CccScheduleEvent `mandatory:"false" json:"events"`

	// List of Compute Cloud@Customer infrastructure
	// OCIDs (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that are using this upgrade
	// schedule.
	InfrastructureIds []string `mandatory:"false" json:"infrastructureIds"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m CccUpgradeSchedule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CccUpgradeSchedule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCccUpgradeScheduleLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetCccUpgradeScheduleLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CccUpgradeScheduleLifecycleStateEnum Enum with underlying type: string
type CccUpgradeScheduleLifecycleStateEnum string

// Set of constants representing the allowable values for CccUpgradeScheduleLifecycleStateEnum
const (
	CccUpgradeScheduleLifecycleStateActive         CccUpgradeScheduleLifecycleStateEnum = "ACTIVE"
	CccUpgradeScheduleLifecycleStateNeedsAttention CccUpgradeScheduleLifecycleStateEnum = "NEEDS_ATTENTION"
	CccUpgradeScheduleLifecycleStateDeleted        CccUpgradeScheduleLifecycleStateEnum = "DELETED"
	CccUpgradeScheduleLifecycleStateFailed         CccUpgradeScheduleLifecycleStateEnum = "FAILED"
)

var mappingCccUpgradeScheduleLifecycleStateEnum = map[string]CccUpgradeScheduleLifecycleStateEnum{
	"ACTIVE":          CccUpgradeScheduleLifecycleStateActive,
	"NEEDS_ATTENTION": CccUpgradeScheduleLifecycleStateNeedsAttention,
	"DELETED":         CccUpgradeScheduleLifecycleStateDeleted,
	"FAILED":          CccUpgradeScheduleLifecycleStateFailed,
}

var mappingCccUpgradeScheduleLifecycleStateEnumLowerCase = map[string]CccUpgradeScheduleLifecycleStateEnum{
	"active":          CccUpgradeScheduleLifecycleStateActive,
	"needs_attention": CccUpgradeScheduleLifecycleStateNeedsAttention,
	"deleted":         CccUpgradeScheduleLifecycleStateDeleted,
	"failed":          CccUpgradeScheduleLifecycleStateFailed,
}

// GetCccUpgradeScheduleLifecycleStateEnumValues Enumerates the set of values for CccUpgradeScheduleLifecycleStateEnum
func GetCccUpgradeScheduleLifecycleStateEnumValues() []CccUpgradeScheduleLifecycleStateEnum {
	values := make([]CccUpgradeScheduleLifecycleStateEnum, 0)
	for _, v := range mappingCccUpgradeScheduleLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetCccUpgradeScheduleLifecycleStateEnumStringValues Enumerates the set of values in String for CccUpgradeScheduleLifecycleStateEnum
func GetCccUpgradeScheduleLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETED",
		"FAILED",
	}
}

// GetMappingCccUpgradeScheduleLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCccUpgradeScheduleLifecycleStateEnum(val string) (CccUpgradeScheduleLifecycleStateEnum, bool) {
	enum, ok := mappingCccUpgradeScheduleLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
