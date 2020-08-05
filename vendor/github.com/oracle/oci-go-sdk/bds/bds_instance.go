// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// BdsInstance Description of the BDS instance
type BdsInstance struct {

	// The OCID of the BDS resource
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Name of the BDS instance
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the BDS instance
	LifecycleState BdsInstanceLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Boolean flag specifying whether or not the cluster is HA
	IsHighAvailability *bool `mandatory:"true" json:"isHighAvailability"`

	// Boolean flag specifying whether or not the cluster should be setup as secure.
	IsSecure *bool `mandatory:"true" json:"isSecure"`

	// Boolean flag specifying whether we configure Cloud SQL or not
	IsCloudSqlConfigured *bool `mandatory:"true" json:"isCloudSqlConfigured"`

	// The list of nodes in the BDS instance
	Nodes []Node `mandatory:"true" json:"nodes"`

	// Number of nodes that forming the cluster
	NumberOfNodes *int `mandatory:"true" json:"numberOfNodes"`

	// Version of the Hadoop distribution
	ClusterVersion BdsInstanceClusterVersionEnum `mandatory:"false" json:"clusterVersion,omitempty"`

	// Additional configuration of customer's network.
	NetworkConfig *NetworkConfig `mandatory:"false" json:"networkConfig"`

	// Specific info about a Hadoop cluster
	ClusterDetails *ClusterDetails `mandatory:"false" json:"clusterDetails"`

	// The information about added Cloud SQL capability
	CloudSqlDetails *CloudSqlDetails `mandatory:"false" json:"cloudSqlDetails"`

	// The user who created the BDS instance.
	CreatedBy *string `mandatory:"false" json:"createdBy"`

	// The time the BDS instance was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time the BDS instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BdsInstance) String() string {
	return common.PointerString(m)
}

// BdsInstanceLifecycleStateEnum Enum with underlying type: string
type BdsInstanceLifecycleStateEnum string

// Set of constants representing the allowable values for BdsInstanceLifecycleStateEnum
const (
	BdsInstanceLifecycleStateCreating      BdsInstanceLifecycleStateEnum = "CREATING"
	BdsInstanceLifecycleStateActive        BdsInstanceLifecycleStateEnum = "ACTIVE"
	BdsInstanceLifecycleStateUpdating      BdsInstanceLifecycleStateEnum = "UPDATING"
	BdsInstanceLifecycleStateUpdatingInfra BdsInstanceLifecycleStateEnum = "UPDATING_INFRA"
	BdsInstanceLifecycleStateSuspending    BdsInstanceLifecycleStateEnum = "SUSPENDING"
	BdsInstanceLifecycleStateSuspended     BdsInstanceLifecycleStateEnum = "SUSPENDED"
	BdsInstanceLifecycleStateResuming      BdsInstanceLifecycleStateEnum = "RESUMING"
	BdsInstanceLifecycleStateDeleting      BdsInstanceLifecycleStateEnum = "DELETING"
	BdsInstanceLifecycleStateDeleted       BdsInstanceLifecycleStateEnum = "DELETED"
	BdsInstanceLifecycleStateFailed        BdsInstanceLifecycleStateEnum = "FAILED"
)

var mappingBdsInstanceLifecycleState = map[string]BdsInstanceLifecycleStateEnum{
	"CREATING":       BdsInstanceLifecycleStateCreating,
	"ACTIVE":         BdsInstanceLifecycleStateActive,
	"UPDATING":       BdsInstanceLifecycleStateUpdating,
	"UPDATING_INFRA": BdsInstanceLifecycleStateUpdatingInfra,
	"SUSPENDING":     BdsInstanceLifecycleStateSuspending,
	"SUSPENDED":      BdsInstanceLifecycleStateSuspended,
	"RESUMING":       BdsInstanceLifecycleStateResuming,
	"DELETING":       BdsInstanceLifecycleStateDeleting,
	"DELETED":        BdsInstanceLifecycleStateDeleted,
	"FAILED":         BdsInstanceLifecycleStateFailed,
}

// GetBdsInstanceLifecycleStateEnumValues Enumerates the set of values for BdsInstanceLifecycleStateEnum
func GetBdsInstanceLifecycleStateEnumValues() []BdsInstanceLifecycleStateEnum {
	values := make([]BdsInstanceLifecycleStateEnum, 0)
	for _, v := range mappingBdsInstanceLifecycleState {
		values = append(values, v)
	}
	return values
}

// BdsInstanceClusterVersionEnum Enum with underlying type: string
type BdsInstanceClusterVersionEnum string

// Set of constants representing the allowable values for BdsInstanceClusterVersionEnum
const (
	BdsInstanceClusterVersionCdh5 BdsInstanceClusterVersionEnum = "CDH5"
	BdsInstanceClusterVersionCdh6 BdsInstanceClusterVersionEnum = "CDH6"
)

var mappingBdsInstanceClusterVersion = map[string]BdsInstanceClusterVersionEnum{
	"CDH5": BdsInstanceClusterVersionCdh5,
	"CDH6": BdsInstanceClusterVersionCdh6,
}

// GetBdsInstanceClusterVersionEnumValues Enumerates the set of values for BdsInstanceClusterVersionEnum
func GetBdsInstanceClusterVersionEnumValues() []BdsInstanceClusterVersionEnum {
	values := make([]BdsInstanceClusterVersionEnum, 0)
	for _, v := range mappingBdsInstanceClusterVersion {
		values = append(values, v)
	}
	return values
}
