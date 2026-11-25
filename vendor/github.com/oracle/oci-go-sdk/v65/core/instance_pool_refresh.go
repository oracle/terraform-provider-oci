// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
// The Core Services API provides both IPv4-only and dual-stack (IPv4 and IPv6) endpoints. By default, OCI-published SDKs and the CLI use IPv4-only endpoints.
// Use a dual-stack endpoint to connect over IPv6 while retaining IPv4 compatibility.
// Dual-stack endpoints use the following format: https://{service}.{region}.ds.oci.{secondLevelDomain}.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InstancePoolRefresh An instance pool refresh is an asynchronous operation that replaces eligible instances in an
// instance pool with instances launched from the pool's target instance configuration.
type InstancePoolRefresh struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool refresh.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the instance pool.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance pool.
	InstancePoolId *string `mandatory:"true" json:"instancePoolId"`

	// The current state of the instance pool refresh.
	LifecycleState InstancePoolRefreshLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance configuration captured
	// from the instance pool when the refresh started.
	TargetInstanceConfigurationId *string `mandatory:"true" json:"targetInstanceConfigurationId"`

	// The date and time the instance pool refresh was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{ "orcl-cloud": { "free-tier-retained": "true" } }`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The maximum healthy percentage requested for launch-before-terminate capacity.
	MaxHealthyPercent *int `mandatory:"false" json:"maxHealthyPercent"`

	// The requested warmup duration in seconds.
	WarmupSeconds *int `mandatory:"false" json:"warmupSeconds"`

	// Whether the refresh can select instances that already match the target instance configuration.
	IsFullRefresh *bool `mandatory:"false" json:"isFullRefresh"`

	// The maximum number of instances refreshed in a batch.
	RefreshBatchSize *int `mandatory:"false" json:"refreshBatchSize"`

	// The total number of instances selected for this refresh.
	TotalMemberCount *int `mandatory:"false" json:"totalMemberCount"`

	// The number of instances remaining to be refreshed.
	RemainingMemberCount *int `mandatory:"false" json:"remainingMemberCount"`

	// The percentage of the refresh that is complete.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// The number of instances skipped by the refresh.
	SkippedMemberCount *int `mandatory:"false" json:"skippedMemberCount"`

	// The date and time the instance pool refresh reached a terminal state, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`
}

func (m InstancePoolRefresh) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstancePoolRefresh) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstancePoolRefreshLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInstancePoolRefreshLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstancePoolRefreshLifecycleStateEnum Enum with underlying type: string
type InstancePoolRefreshLifecycleStateEnum string

// Set of constants representing the allowable values for InstancePoolRefreshLifecycleStateEnum
const (
	InstancePoolRefreshLifecycleStateAccepted   InstancePoolRefreshLifecycleStateEnum = "ACCEPTED"
	InstancePoolRefreshLifecycleStateInProgress InstancePoolRefreshLifecycleStateEnum = "IN_PROGRESS"
	InstancePoolRefreshLifecycleStateSucceeded  InstancePoolRefreshLifecycleStateEnum = "SUCCEEDED"
	InstancePoolRefreshLifecycleStateFailed     InstancePoolRefreshLifecycleStateEnum = "FAILED"
	InstancePoolRefreshLifecycleStateCanceling  InstancePoolRefreshLifecycleStateEnum = "CANCELING"
	InstancePoolRefreshLifecycleStateCanceled   InstancePoolRefreshLifecycleStateEnum = "CANCELED"
)

var mappingInstancePoolRefreshLifecycleStateEnum = map[string]InstancePoolRefreshLifecycleStateEnum{
	"ACCEPTED":    InstancePoolRefreshLifecycleStateAccepted,
	"IN_PROGRESS": InstancePoolRefreshLifecycleStateInProgress,
	"SUCCEEDED":   InstancePoolRefreshLifecycleStateSucceeded,
	"FAILED":      InstancePoolRefreshLifecycleStateFailed,
	"CANCELING":   InstancePoolRefreshLifecycleStateCanceling,
	"CANCELED":    InstancePoolRefreshLifecycleStateCanceled,
}

var mappingInstancePoolRefreshLifecycleStateEnumLowerCase = map[string]InstancePoolRefreshLifecycleStateEnum{
	"accepted":    InstancePoolRefreshLifecycleStateAccepted,
	"in_progress": InstancePoolRefreshLifecycleStateInProgress,
	"succeeded":   InstancePoolRefreshLifecycleStateSucceeded,
	"failed":      InstancePoolRefreshLifecycleStateFailed,
	"canceling":   InstancePoolRefreshLifecycleStateCanceling,
	"canceled":    InstancePoolRefreshLifecycleStateCanceled,
}

// GetInstancePoolRefreshLifecycleStateEnumValues Enumerates the set of values for InstancePoolRefreshLifecycleStateEnum
func GetInstancePoolRefreshLifecycleStateEnumValues() []InstancePoolRefreshLifecycleStateEnum {
	values := make([]InstancePoolRefreshLifecycleStateEnum, 0)
	for _, v := range mappingInstancePoolRefreshLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInstancePoolRefreshLifecycleStateEnumStringValues Enumerates the set of values in String for InstancePoolRefreshLifecycleStateEnum
func GetInstancePoolRefreshLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingInstancePoolRefreshLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstancePoolRefreshLifecycleStateEnum(val string) (InstancePoolRefreshLifecycleStateEnum, bool) {
	enum, ok := mappingInstancePoolRefreshLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
