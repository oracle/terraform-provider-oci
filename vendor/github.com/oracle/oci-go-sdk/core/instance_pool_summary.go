// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// InstancePoolSummary Condensed InstancePool data when listing instance pools.
type InstancePoolSummary struct {

	// The OCID of the instance pool
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the instance pool
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the instance configuration associated to the intance pool.
	InstanceConfigurationId *string `mandatory:"true" json:"instanceConfigurationId"`

	// The current state of the instance pool.
	LifecycleState InstancePoolSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The availability domains for the instance pool.
	AvailabilityDomains []string `mandatory:"true" json:"availabilityDomains"`

	// The number of instances that should be in the instance pool.
	Size *int `mandatory:"true" json:"size"`

	// The date and time the instance pool was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The user-friendly name.  Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m InstancePoolSummary) String() string {
	return common.PointerString(m)
}

// InstancePoolSummaryLifecycleStateEnum Enum with underlying type: string
type InstancePoolSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for InstancePoolSummaryLifecycleStateEnum
const (
	InstancePoolSummaryLifecycleStateProvisioning InstancePoolSummaryLifecycleStateEnum = "PROVISIONING"
	InstancePoolSummaryLifecycleStateScaling      InstancePoolSummaryLifecycleStateEnum = "SCALING"
	InstancePoolSummaryLifecycleStateStarting     InstancePoolSummaryLifecycleStateEnum = "STARTING"
	InstancePoolSummaryLifecycleStateStopping     InstancePoolSummaryLifecycleStateEnum = "STOPPING"
	InstancePoolSummaryLifecycleStateTerminating  InstancePoolSummaryLifecycleStateEnum = "TERMINATING"
	InstancePoolSummaryLifecycleStateStopped      InstancePoolSummaryLifecycleStateEnum = "STOPPED"
	InstancePoolSummaryLifecycleStateTerminated   InstancePoolSummaryLifecycleStateEnum = "TERMINATED"
	InstancePoolSummaryLifecycleStateRunning      InstancePoolSummaryLifecycleStateEnum = "RUNNING"
)

var mappingInstancePoolSummaryLifecycleState = map[string]InstancePoolSummaryLifecycleStateEnum{
	"PROVISIONING": InstancePoolSummaryLifecycleStateProvisioning,
	"SCALING":      InstancePoolSummaryLifecycleStateScaling,
	"STARTING":     InstancePoolSummaryLifecycleStateStarting,
	"STOPPING":     InstancePoolSummaryLifecycleStateStopping,
	"TERMINATING":  InstancePoolSummaryLifecycleStateTerminating,
	"STOPPED":      InstancePoolSummaryLifecycleStateStopped,
	"TERMINATED":   InstancePoolSummaryLifecycleStateTerminated,
	"RUNNING":      InstancePoolSummaryLifecycleStateRunning,
}

// GetInstancePoolSummaryLifecycleStateEnumValues Enumerates the set of values for InstancePoolSummaryLifecycleStateEnum
func GetInstancePoolSummaryLifecycleStateEnumValues() []InstancePoolSummaryLifecycleStateEnum {
	values := make([]InstancePoolSummaryLifecycleStateEnum, 0)
	for _, v := range mappingInstancePoolSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
