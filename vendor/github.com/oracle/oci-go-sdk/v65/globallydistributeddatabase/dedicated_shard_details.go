// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage distributed databases.
//

package globallydistributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DedicatedShardDetails Details of ATP-D based shard.
type DedicatedShardDetails struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The compute amount available to the underlying autonomous database associated with shard.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Name of the shard-group to which the shard belongs.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The time the the shard was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Status of shard or catalog or gsm for the sharded database.
	Status DedicatedShardDetailsStatusEnum `mandatory:"true" json:"status"`

	// Determines the auto-scaling mode.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// Identifier of the primary cloudAutonomousVmCluster for the shard.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	EncryptionKeyDetails *DedicatedShardOrCatalogEncryptionKeyDetails `mandatory:"false" json:"encryptionKeyDetails"`

	// The time the ssl certificate associated with shard expires. An RFC3339 formatted datetime string
	TimeSslCertificateExpires *common.SDKTime `mandatory:"false" json:"timeSslCertificateExpires"`

	// Shard space name.
	ShardSpace *string `mandatory:"false" json:"shardSpace"`

	// Identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// Identifier of the underlying container database.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// Identifier of the underlying container database parent.
	ContainerDatabaseParentId *string `mandatory:"false" json:"containerDatabaseParentId"`

	// Identifier of the peer cloudAutonomousVmCluster for the shard.
	PeerCloudAutonomousVmClusterId *string `mandatory:"false" json:"peerCloudAutonomousVmClusterId"`

	// Additional metadata related to shard's underlying supporting resource.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`
}

func (m DedicatedShardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DedicatedShardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDedicatedShardDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDedicatedShardDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DedicatedShardDetailsStatusEnum Enum with underlying type: string
type DedicatedShardDetailsStatusEnum string

// Set of constants representing the allowable values for DedicatedShardDetailsStatusEnum
const (
	DedicatedShardDetailsStatusFailed                DedicatedShardDetailsStatusEnum = "FAILED"
	DedicatedShardDetailsStatusDeleting              DedicatedShardDetailsStatusEnum = "DELETING"
	DedicatedShardDetailsStatusDeleted               DedicatedShardDetailsStatusEnum = "DELETED"
	DedicatedShardDetailsStatusUpdating              DedicatedShardDetailsStatusEnum = "UPDATING"
	DedicatedShardDetailsStatusCreating              DedicatedShardDetailsStatusEnum = "CREATING"
	DedicatedShardDetailsStatusCreated               DedicatedShardDetailsStatusEnum = "CREATED"
	DedicatedShardDetailsStatusReadyForConfiguration DedicatedShardDetailsStatusEnum = "READY_FOR_CONFIGURATION"
	DedicatedShardDetailsStatusConfigured            DedicatedShardDetailsStatusEnum = "CONFIGURED"
	DedicatedShardDetailsStatusNeedsAttention        DedicatedShardDetailsStatusEnum = "NEEDS_ATTENTION"
)

var mappingDedicatedShardDetailsStatusEnum = map[string]DedicatedShardDetailsStatusEnum{
	"FAILED":                  DedicatedShardDetailsStatusFailed,
	"DELETING":                DedicatedShardDetailsStatusDeleting,
	"DELETED":                 DedicatedShardDetailsStatusDeleted,
	"UPDATING":                DedicatedShardDetailsStatusUpdating,
	"CREATING":                DedicatedShardDetailsStatusCreating,
	"CREATED":                 DedicatedShardDetailsStatusCreated,
	"READY_FOR_CONFIGURATION": DedicatedShardDetailsStatusReadyForConfiguration,
	"CONFIGURED":              DedicatedShardDetailsStatusConfigured,
	"NEEDS_ATTENTION":         DedicatedShardDetailsStatusNeedsAttention,
}

var mappingDedicatedShardDetailsStatusEnumLowerCase = map[string]DedicatedShardDetailsStatusEnum{
	"failed":                  DedicatedShardDetailsStatusFailed,
	"deleting":                DedicatedShardDetailsStatusDeleting,
	"deleted":                 DedicatedShardDetailsStatusDeleted,
	"updating":                DedicatedShardDetailsStatusUpdating,
	"creating":                DedicatedShardDetailsStatusCreating,
	"created":                 DedicatedShardDetailsStatusCreated,
	"ready_for_configuration": DedicatedShardDetailsStatusReadyForConfiguration,
	"configured":              DedicatedShardDetailsStatusConfigured,
	"needs_attention":         DedicatedShardDetailsStatusNeedsAttention,
}

// GetDedicatedShardDetailsStatusEnumValues Enumerates the set of values for DedicatedShardDetailsStatusEnum
func GetDedicatedShardDetailsStatusEnumValues() []DedicatedShardDetailsStatusEnum {
	values := make([]DedicatedShardDetailsStatusEnum, 0)
	for _, v := range mappingDedicatedShardDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDedicatedShardDetailsStatusEnumStringValues Enumerates the set of values in String for DedicatedShardDetailsStatusEnum
func GetDedicatedShardDetailsStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"DELETING",
		"DELETED",
		"UPDATING",
		"CREATING",
		"CREATED",
		"READY_FOR_CONFIGURATION",
		"CONFIGURED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDedicatedShardDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDedicatedShardDetailsStatusEnum(val string) (DedicatedShardDetailsStatusEnum, bool) {
	enum, ok := mappingDedicatedShardDetailsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
