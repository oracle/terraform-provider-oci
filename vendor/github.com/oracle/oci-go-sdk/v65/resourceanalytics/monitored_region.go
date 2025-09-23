// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MonitoredRegion A MonitoredRegion is a region to collect data for the associated ResourceAnalyticsInstance.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type MonitoredRegion struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the MonitoredRegion.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this MonitoredRegion.
	ResourceAnalyticsInstanceId *string `mandatory:"true" json:"resourceAnalyticsInstanceId"`

	// The Region Identifier (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm) of this MonitoredRegion.
	RegionId *string `mandatory:"true" json:"regionId"`

	// The date and time the MonitoredRegion was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the MonitoredRegion.
	LifecycleState MonitoredRegionLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MonitoredRegion was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the MonitoredRegion in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m MonitoredRegion) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitoredRegion) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMonitoredRegionLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetMonitoredRegionLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MonitoredRegionLifecycleStateEnum Enum with underlying type: string
type MonitoredRegionLifecycleStateEnum string

// Set of constants representing the allowable values for MonitoredRegionLifecycleStateEnum
const (
	MonitoredRegionLifecycleStateCreating MonitoredRegionLifecycleStateEnum = "CREATING"
	MonitoredRegionLifecycleStateUpdating MonitoredRegionLifecycleStateEnum = "UPDATING"
	MonitoredRegionLifecycleStateActive   MonitoredRegionLifecycleStateEnum = "ACTIVE"
	MonitoredRegionLifecycleStateDeleting MonitoredRegionLifecycleStateEnum = "DELETING"
	MonitoredRegionLifecycleStateDeleted  MonitoredRegionLifecycleStateEnum = "DELETED"
	MonitoredRegionLifecycleStateFailed   MonitoredRegionLifecycleStateEnum = "FAILED"
)

var mappingMonitoredRegionLifecycleStateEnum = map[string]MonitoredRegionLifecycleStateEnum{
	"CREATING": MonitoredRegionLifecycleStateCreating,
	"UPDATING": MonitoredRegionLifecycleStateUpdating,
	"ACTIVE":   MonitoredRegionLifecycleStateActive,
	"DELETING": MonitoredRegionLifecycleStateDeleting,
	"DELETED":  MonitoredRegionLifecycleStateDeleted,
	"FAILED":   MonitoredRegionLifecycleStateFailed,
}

var mappingMonitoredRegionLifecycleStateEnumLowerCase = map[string]MonitoredRegionLifecycleStateEnum{
	"creating": MonitoredRegionLifecycleStateCreating,
	"updating": MonitoredRegionLifecycleStateUpdating,
	"active":   MonitoredRegionLifecycleStateActive,
	"deleting": MonitoredRegionLifecycleStateDeleting,
	"deleted":  MonitoredRegionLifecycleStateDeleted,
	"failed":   MonitoredRegionLifecycleStateFailed,
}

// GetMonitoredRegionLifecycleStateEnumValues Enumerates the set of values for MonitoredRegionLifecycleStateEnum
func GetMonitoredRegionLifecycleStateEnumValues() []MonitoredRegionLifecycleStateEnum {
	values := make([]MonitoredRegionLifecycleStateEnum, 0)
	for _, v := range mappingMonitoredRegionLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetMonitoredRegionLifecycleStateEnumStringValues Enumerates the set of values in String for MonitoredRegionLifecycleStateEnum
func GetMonitoredRegionLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingMonitoredRegionLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingMonitoredRegionLifecycleStateEnum(val string) (MonitoredRegionLifecycleStateEnum, bool) {
	enum, ok := mappingMonitoredRegionLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
