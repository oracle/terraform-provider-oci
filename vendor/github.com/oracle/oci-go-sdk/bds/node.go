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

// Node Specific info about a node
type Node struct {

	// The OCID of the underlying compute instance
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The name of the node
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the node
	LifecycleState NodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// BDS instance node type
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// Shape of the node
	Shape *string `mandatory:"true" json:"shape"`

	// The OCID of the subnet in which the node should be created
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// IP address of the node
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The fingerprint of the SSH key used for node access
	SshFingerprint *string `mandatory:"true" json:"sshFingerprint"`

	// The name of the availability domain the node is running in
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The name of the fault domain the node is running in
	FaultDomain *string `mandatory:"true" json:"faultDomain"`

	// The time the node was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The list of block volumes attached to a given node.
	AttachedBlockVolumes []VolumeAttachmentDetail `mandatory:"false" json:"attachedBlockVolumes"`

	// The fully-qualified hostname (FQDN) of the node
	Hostname *string `mandatory:"false" json:"hostname"`

	// The OCID of the image from which the node was created
	ImageId *string `mandatory:"false" json:"imageId"`

	// The time the BDS instance was updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m Node) String() string {
	return common.PointerString(m)
}

// NodeLifecycleStateEnum Enum with underlying type: string
type NodeLifecycleStateEnum string

// Set of constants representing the allowable values for NodeLifecycleStateEnum
const (
	NodeLifecycleStateCreating NodeLifecycleStateEnum = "CREATING"
	NodeLifecycleStateActive   NodeLifecycleStateEnum = "ACTIVE"
	NodeLifecycleStateInactive NodeLifecycleStateEnum = "INACTIVE"
	NodeLifecycleStateUpdating NodeLifecycleStateEnum = "UPDATING"
	NodeLifecycleStateDeleting NodeLifecycleStateEnum = "DELETING"
	NodeLifecycleStateDeleted  NodeLifecycleStateEnum = "DELETED"
	NodeLifecycleStateFailed   NodeLifecycleStateEnum = "FAILED"
	NodeLifecycleStateStopping NodeLifecycleStateEnum = "STOPPING"
	NodeLifecycleStateStarting NodeLifecycleStateEnum = "STARTING"
)

var mappingNodeLifecycleState = map[string]NodeLifecycleStateEnum{
	"CREATING": NodeLifecycleStateCreating,
	"ACTIVE":   NodeLifecycleStateActive,
	"INACTIVE": NodeLifecycleStateInactive,
	"UPDATING": NodeLifecycleStateUpdating,
	"DELETING": NodeLifecycleStateDeleting,
	"DELETED":  NodeLifecycleStateDeleted,
	"FAILED":   NodeLifecycleStateFailed,
	"STOPPING": NodeLifecycleStateStopping,
	"STARTING": NodeLifecycleStateStarting,
}

// GetNodeLifecycleStateEnumValues Enumerates the set of values for NodeLifecycleStateEnum
func GetNodeLifecycleStateEnumValues() []NodeLifecycleStateEnum {
	values := make([]NodeLifecycleStateEnum, 0)
	for _, v := range mappingNodeLifecycleState {
		values = append(values, v)
	}
	return values
}

// NodeNodeTypeEnum Enum with underlying type: string
type NodeNodeTypeEnum string

// Set of constants representing the allowable values for NodeNodeTypeEnum
const (
	NodeNodeTypeMaster   NodeNodeTypeEnum = "MASTER"
	NodeNodeTypeEdge     NodeNodeTypeEnum = "EDGE"
	NodeNodeTypeUtility  NodeNodeTypeEnum = "UTILITY"
	NodeNodeTypeWorker   NodeNodeTypeEnum = "WORKER"
	NodeNodeTypeBursting NodeNodeTypeEnum = "BURSTING"
	NodeNodeTypeCloudSql NodeNodeTypeEnum = "CLOUD_SQL"
)

var mappingNodeNodeType = map[string]NodeNodeTypeEnum{
	"MASTER":    NodeNodeTypeMaster,
	"EDGE":      NodeNodeTypeEdge,
	"UTILITY":   NodeNodeTypeUtility,
	"WORKER":    NodeNodeTypeWorker,
	"BURSTING":  NodeNodeTypeBursting,
	"CLOUD_SQL": NodeNodeTypeCloudSql,
}

// GetNodeNodeTypeEnumValues Enumerates the set of values for NodeNodeTypeEnum
func GetNodeNodeTypeEnumValues() []NodeNodeTypeEnum {
	values := make([]NodeNodeTypeEnum, 0)
	for _, v := range mappingNodeNodeType {
		values = append(values, v)
	}
	return values
}
