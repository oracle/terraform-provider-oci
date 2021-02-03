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
	"github.com/oracle/oci-go-sdk/v35/common"
)

// AutoScalingConfigurationSummary The information about auto scale configuration.
type AutoScalingConfigurationSummary struct {

	// The OCID of the autoscaling configuration.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the autoscaling configuration
	LifecycleState AutoScalingConfigurationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A node type that is managed by an autoscaling configuration. The only supported type is WORKER.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// The time the BDS instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the autoscale configuration was updated.
	// An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Policy *AutoScalePolicy `mandatory:"true" json:"policy"`
}

func (m AutoScalingConfigurationSummary) String() string {
	return common.PointerString(m)
}
