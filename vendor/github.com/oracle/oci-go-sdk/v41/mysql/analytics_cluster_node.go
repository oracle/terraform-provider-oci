// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// MySQL Database Service API
//
// The API for the MySQL Database Service
//

package mysql

import (
	"github.com/oracle/oci-go-sdk/v41/common"
)

// AnalyticsClusterNode DEPRECATED -- please use HeatWave API instead.
// An Analytics Cluster Node is a compute host that is part of an Analytics Cluster.
type AnalyticsClusterNode struct {

	// The ID of the node within MySQL Analytics Cluster.
	NodeId *string `mandatory:"true" json:"nodeId"`

	// The current state of the MySQL Analytics Cluster node.
	LifecycleState AnalyticsClusterNodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the MySQL Analytics Cluster node was created, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the MySQL Analytics Cluster node was updated, as described by RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m AnalyticsClusterNode) String() string {
	return common.PointerString(m)
}

// AnalyticsClusterNodeLifecycleStateEnum Enum with underlying type: string
type AnalyticsClusterNodeLifecycleStateEnum string

// Set of constants representing the allowable values for AnalyticsClusterNodeLifecycleStateEnum
const (
	AnalyticsClusterNodeLifecycleStateCreating AnalyticsClusterNodeLifecycleStateEnum = "CREATING"
	AnalyticsClusterNodeLifecycleStateActive   AnalyticsClusterNodeLifecycleStateEnum = "ACTIVE"
	AnalyticsClusterNodeLifecycleStateInactive AnalyticsClusterNodeLifecycleStateEnum = "INACTIVE"
	AnalyticsClusterNodeLifecycleStateUpdating AnalyticsClusterNodeLifecycleStateEnum = "UPDATING"
	AnalyticsClusterNodeLifecycleStateDeleting AnalyticsClusterNodeLifecycleStateEnum = "DELETING"
	AnalyticsClusterNodeLifecycleStateDeleted  AnalyticsClusterNodeLifecycleStateEnum = "DELETED"
	AnalyticsClusterNodeLifecycleStateFailed   AnalyticsClusterNodeLifecycleStateEnum = "FAILED"
)

var mappingAnalyticsClusterNodeLifecycleState = map[string]AnalyticsClusterNodeLifecycleStateEnum{
	"CREATING": AnalyticsClusterNodeLifecycleStateCreating,
	"ACTIVE":   AnalyticsClusterNodeLifecycleStateActive,
	"INACTIVE": AnalyticsClusterNodeLifecycleStateInactive,
	"UPDATING": AnalyticsClusterNodeLifecycleStateUpdating,
	"DELETING": AnalyticsClusterNodeLifecycleStateDeleting,
	"DELETED":  AnalyticsClusterNodeLifecycleStateDeleted,
	"FAILED":   AnalyticsClusterNodeLifecycleStateFailed,
}

// GetAnalyticsClusterNodeLifecycleStateEnumValues Enumerates the set of values for AnalyticsClusterNodeLifecycleStateEnum
func GetAnalyticsClusterNodeLifecycleStateEnumValues() []AnalyticsClusterNodeLifecycleStateEnum {
	values := make([]AnalyticsClusterNodeLifecycleStateEnum, 0)
	for _, v := range mappingAnalyticsClusterNodeLifecycleState {
		values = append(values, v)
	}
	return values
}
