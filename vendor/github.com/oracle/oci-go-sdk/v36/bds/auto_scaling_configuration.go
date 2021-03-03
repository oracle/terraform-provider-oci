// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// API for the Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service
// build on Hadoop, Spark and Data Science distribution, which can be fully integrated with existing enterprise
// data in Oracle Database and Oracle Applications..
//

package bds

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// AutoScalingConfiguration The information about auto scale configuration.
type AutoScalingConfiguration struct {

	// The unique identifier for autoscaling configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A node type that is managed by an autoscaling configuration. The only supported type is WORKER.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// The state of the autoscaling configuration
	LifecycleState AutoScalingConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the BDS instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the autoscale configuration was updated.
	// An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Policy *AutoScalePolicy `mandatory:"true" json:"policy"`
}

func (m AutoScalingConfiguration) String() string {
	return common.PointerString(m)
}

// AutoScalingConfigurationLifecycleStateEnum Enum with underlying type: string
type AutoScalingConfigurationLifecycleStateEnum string

// Set of constants representing the allowable values for AutoScalingConfigurationLifecycleStateEnum
const (
	AutoScalingConfigurationLifecycleStateCreating AutoScalingConfigurationLifecycleStateEnum = "CREATING"
	AutoScalingConfigurationLifecycleStateActive   AutoScalingConfigurationLifecycleStateEnum = "ACTIVE"
	AutoScalingConfigurationLifecycleStateUpdating AutoScalingConfigurationLifecycleStateEnum = "UPDATING"
	AutoScalingConfigurationLifecycleStateDeleting AutoScalingConfigurationLifecycleStateEnum = "DELETING"
	AutoScalingConfigurationLifecycleStateDeleted  AutoScalingConfigurationLifecycleStateEnum = "DELETED"
	AutoScalingConfigurationLifecycleStateFailed   AutoScalingConfigurationLifecycleStateEnum = "FAILED"
)

var mappingAutoScalingConfigurationLifecycleState = map[string]AutoScalingConfigurationLifecycleStateEnum{
	"CREATING": AutoScalingConfigurationLifecycleStateCreating,
	"ACTIVE":   AutoScalingConfigurationLifecycleStateActive,
	"UPDATING": AutoScalingConfigurationLifecycleStateUpdating,
	"DELETING": AutoScalingConfigurationLifecycleStateDeleting,
	"DELETED":  AutoScalingConfigurationLifecycleStateDeleted,
	"FAILED":   AutoScalingConfigurationLifecycleStateFailed,
}

// GetAutoScalingConfigurationLifecycleStateEnumValues Enumerates the set of values for AutoScalingConfigurationLifecycleStateEnum
func GetAutoScalingConfigurationLifecycleStateEnumValues() []AutoScalingConfigurationLifecycleStateEnum {
	values := make([]AutoScalingConfigurationLifecycleStateEnum, 0)
	for _, v := range mappingAutoScalingConfigurationLifecycleState {
		values = append(values, v)
	}
	return values
}
