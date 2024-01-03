// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Node Details about a node.
type Node struct {

	// The OCID of the underlying Oracle Cloud Infrastructure Compute instance.
	InstanceId *string `mandatory:"true" json:"instanceId"`

	// The name of the node.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The state of the node.
	LifecycleState NodeLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Cluster node type.
	NodeType NodeNodeTypeEnum `mandatory:"true" json:"nodeType"`

	// Shape of the node.
	Shape *string `mandatory:"true" json:"shape"`

	// The OCID of the subnet in which the node is to be created.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// IP address of the node.
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The fingerprint of the SSH key used for node access.
	SshFingerprint *string `mandatory:"true" json:"sshFingerprint"`

	// The name of the availability domain in which the node is running.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The name of the fault domain in which the node is running.
	FaultDomain *string `mandatory:"true" json:"faultDomain"`

	// The time the node was created, shown as an RFC 3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The list of block volumes attached to a given node.
	AttachedBlockVolumes []VolumeAttachmentDetail `mandatory:"false" json:"attachedBlockVolumes"`

	// The fully-qualified hostname (FQDN) of the node.
	Hostname *string `mandatory:"false" json:"hostname"`

	// The OCID of the image from which the node was created.
	ImageId *string `mandatory:"false" json:"imageId"`

	// The time the cluster was updated, shown as an RFC 3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The total number of OCPUs available to the node.
	Ocpus *int `mandatory:"false" json:"ocpus"`

	// The total amount of memory available to the node, in gigabytes.
	MemoryInGBs *int `mandatory:"false" json:"memoryInGBs"`

	// The number of NVMe drives to be used for storage. A single drive has 6.8 TB available.
	Nvmes *int `mandatory:"false" json:"nvmes"`

	// The aggregate size of all local disks, in gigabytes. If the instance does not have any local disks, this field is null.
	LocalDisksTotalSizeInGBs *float64 `mandatory:"false" json:"localDisksTotalSizeInGBs"`

	// The date and time the instance is expected to be stopped / started, in the format defined by RFC3339.
	TimeMaintenanceRebootDue *common.SDKTime `mandatory:"false" json:"timeMaintenanceRebootDue"`
}

func (m Node) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Node) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNodeLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNodeLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNodeNodeTypeEnum(string(m.NodeType)); !ok && m.NodeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NodeType: %s. Supported values are: %s.", m.NodeType, strings.Join(GetNodeNodeTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
	NodeLifecycleStateStopped  NodeLifecycleStateEnum = "STOPPED"
	NodeLifecycleStateStopping NodeLifecycleStateEnum = "STOPPING"
	NodeLifecycleStateStarting NodeLifecycleStateEnum = "STARTING"
)

var mappingNodeLifecycleStateEnum = map[string]NodeLifecycleStateEnum{
	"CREATING": NodeLifecycleStateCreating,
	"ACTIVE":   NodeLifecycleStateActive,
	"INACTIVE": NodeLifecycleStateInactive,
	"UPDATING": NodeLifecycleStateUpdating,
	"DELETING": NodeLifecycleStateDeleting,
	"DELETED":  NodeLifecycleStateDeleted,
	"FAILED":   NodeLifecycleStateFailed,
	"STOPPED":  NodeLifecycleStateStopped,
	"STOPPING": NodeLifecycleStateStopping,
	"STARTING": NodeLifecycleStateStarting,
}

var mappingNodeLifecycleStateEnumLowerCase = map[string]NodeLifecycleStateEnum{
	"creating": NodeLifecycleStateCreating,
	"active":   NodeLifecycleStateActive,
	"inactive": NodeLifecycleStateInactive,
	"updating": NodeLifecycleStateUpdating,
	"deleting": NodeLifecycleStateDeleting,
	"deleted":  NodeLifecycleStateDeleted,
	"failed":   NodeLifecycleStateFailed,
	"stopped":  NodeLifecycleStateStopped,
	"stopping": NodeLifecycleStateStopping,
	"starting": NodeLifecycleStateStarting,
}

// GetNodeLifecycleStateEnumValues Enumerates the set of values for NodeLifecycleStateEnum
func GetNodeLifecycleStateEnumValues() []NodeLifecycleStateEnum {
	values := make([]NodeLifecycleStateEnum, 0)
	for _, v := range mappingNodeLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeLifecycleStateEnumStringValues Enumerates the set of values in String for NodeLifecycleStateEnum
func GetNodeLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"INACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
		"STOPPED",
		"STOPPING",
		"STARTING",
	}
}

// GetMappingNodeLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeLifecycleStateEnum(val string) (NodeLifecycleStateEnum, bool) {
	enum, ok := mappingNodeLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NodeNodeTypeEnum Enum with underlying type: string
type NodeNodeTypeEnum string

// Set of constants representing the allowable values for NodeNodeTypeEnum
const (
	NodeNodeTypeMaster            NodeNodeTypeEnum = "MASTER"
	NodeNodeTypeEdge              NodeNodeTypeEnum = "EDGE"
	NodeNodeTypeUtility           NodeNodeTypeEnum = "UTILITY"
	NodeNodeTypeWorker            NodeNodeTypeEnum = "WORKER"
	NodeNodeTypeComputeOnlyWorker NodeNodeTypeEnum = "COMPUTE_ONLY_WORKER"
	NodeNodeTypeKafkaBroker       NodeNodeTypeEnum = "KAFKA_BROKER"
	NodeNodeTypeBursting          NodeNodeTypeEnum = "BURSTING"
	NodeNodeTypeCloudSql          NodeNodeTypeEnum = "CLOUD_SQL"
)

var mappingNodeNodeTypeEnum = map[string]NodeNodeTypeEnum{
	"MASTER":              NodeNodeTypeMaster,
	"EDGE":                NodeNodeTypeEdge,
	"UTILITY":             NodeNodeTypeUtility,
	"WORKER":              NodeNodeTypeWorker,
	"COMPUTE_ONLY_WORKER": NodeNodeTypeComputeOnlyWorker,
	"KAFKA_BROKER":        NodeNodeTypeKafkaBroker,
	"BURSTING":            NodeNodeTypeBursting,
	"CLOUD_SQL":           NodeNodeTypeCloudSql,
}

var mappingNodeNodeTypeEnumLowerCase = map[string]NodeNodeTypeEnum{
	"master":              NodeNodeTypeMaster,
	"edge":                NodeNodeTypeEdge,
	"utility":             NodeNodeTypeUtility,
	"worker":              NodeNodeTypeWorker,
	"compute_only_worker": NodeNodeTypeComputeOnlyWorker,
	"kafka_broker":        NodeNodeTypeKafkaBroker,
	"bursting":            NodeNodeTypeBursting,
	"cloud_sql":           NodeNodeTypeCloudSql,
}

// GetNodeNodeTypeEnumValues Enumerates the set of values for NodeNodeTypeEnum
func GetNodeNodeTypeEnumValues() []NodeNodeTypeEnum {
	values := make([]NodeNodeTypeEnum, 0)
	for _, v := range mappingNodeNodeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetNodeNodeTypeEnumStringValues Enumerates the set of values in String for NodeNodeTypeEnum
func GetNodeNodeTypeEnumStringValues() []string {
	return []string{
		"MASTER",
		"EDGE",
		"UTILITY",
		"WORKER",
		"COMPUTE_ONLY_WORKER",
		"KAFKA_BROKER",
		"BURSTING",
		"CLOUD_SQL",
	}
}

// GetMappingNodeNodeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNodeNodeTypeEnum(val string) (NodeNodeTypeEnum, bool) {
	enum, ok := mappingNodeNodeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
