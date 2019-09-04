// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ClusterNetworkSummary Condensed Cluster Network data when listing cluster networks.
type ClusterNetworkSummary struct {

	// The OCID of the cluster network.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment containing the cluster netowrk.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the cluster network.
	LifecycleState ClusterNetworkSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the resource was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the resource was updated, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The user-friendly name.  Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The instance pools in the cluster network.
	InstancePools []InstancePoolSummary `mandatory:"false" json:"instancePools"`
}

func (m ClusterNetworkSummary) String() string {
	return common.PointerString(m)
}

// ClusterNetworkSummaryLifecycleStateEnum Enum with underlying type: string
type ClusterNetworkSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ClusterNetworkSummaryLifecycleStateEnum
const (
	ClusterNetworkSummaryLifecycleStateProvisioning ClusterNetworkSummaryLifecycleStateEnum = "PROVISIONING"
	ClusterNetworkSummaryLifecycleStateScaling      ClusterNetworkSummaryLifecycleStateEnum = "SCALING"
	ClusterNetworkSummaryLifecycleStateStarting     ClusterNetworkSummaryLifecycleStateEnum = "STARTING"
	ClusterNetworkSummaryLifecycleStateStopping     ClusterNetworkSummaryLifecycleStateEnum = "STOPPING"
	ClusterNetworkSummaryLifecycleStateTerminating  ClusterNetworkSummaryLifecycleStateEnum = "TERMINATING"
	ClusterNetworkSummaryLifecycleStateStopped      ClusterNetworkSummaryLifecycleStateEnum = "STOPPED"
	ClusterNetworkSummaryLifecycleStateTerminated   ClusterNetworkSummaryLifecycleStateEnum = "TERMINATED"
	ClusterNetworkSummaryLifecycleStateRunning      ClusterNetworkSummaryLifecycleStateEnum = "RUNNING"
)

var mappingClusterNetworkSummaryLifecycleState = map[string]ClusterNetworkSummaryLifecycleStateEnum{
	"PROVISIONING": ClusterNetworkSummaryLifecycleStateProvisioning,
	"SCALING":      ClusterNetworkSummaryLifecycleStateScaling,
	"STARTING":     ClusterNetworkSummaryLifecycleStateStarting,
	"STOPPING":     ClusterNetworkSummaryLifecycleStateStopping,
	"TERMINATING":  ClusterNetworkSummaryLifecycleStateTerminating,
	"STOPPED":      ClusterNetworkSummaryLifecycleStateStopped,
	"TERMINATED":   ClusterNetworkSummaryLifecycleStateTerminated,
	"RUNNING":      ClusterNetworkSummaryLifecycleStateRunning,
}

// GetClusterNetworkSummaryLifecycleStateEnumValues Enumerates the set of values for ClusterNetworkSummaryLifecycleStateEnum
func GetClusterNetworkSummaryLifecycleStateEnumValues() []ClusterNetworkSummaryLifecycleStateEnum {
	values := make([]ClusterNetworkSummaryLifecycleStateEnum, 0)
	for _, v := range mappingClusterNetworkSummaryLifecycleState {
		values = append(values, v)
	}
	return values
}
