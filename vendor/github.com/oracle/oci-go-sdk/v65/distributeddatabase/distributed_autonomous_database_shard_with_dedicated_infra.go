// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DistributedAutonomousDatabaseShardWithDedicatedInfra Globally distributed autonomous database shard with dedicated autonomous infrastructure.
type DistributedAutonomousDatabaseShardWithDedicatedInfra struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The compute count for the shard database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the shard database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// Determines the auto-scaling mode for the shard database.
	IsAutoScalingEnabled *bool `mandatory:"true" json:"isAutoScalingEnabled"`

	// The name of the shardGroup for the shard.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	// This field is deprecated. For shard peer details please refer peerDetails attribute.
	PeerCloudAutonomousVmClusterIds []string `mandatory:"false" json:"peerCloudAutonomousVmClusterIds"`

	// The shard space name for the Globally distributed autonomous database. Shard space for existing shard cannot be changed, once shard is created.
	// Shard space name shall be used while creation of new shards.
	ShardSpace *string `mandatory:"false" json:"shardSpace"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// Peer details for the shard with dedicated infrastructure.
	PeerDetails []ShardPeerWithDedicatedInfra `mandatory:"false" json:"peerDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the shard.
	OkvKeyStoreId *string `mandatory:"false" json:"okvKeyStoreId"`

	// The OKV endpoint name.
	OkvEndPointGroup *string `mandatory:"false" json:"okvEndPointGroup"`

	Metadata *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`

	// Status of shard with dedicated infrastructure for the Globally distributed autonomous database.
	Status DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum `mandatory:"true" json:"status"`
}

// GetName returns Name
func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedAutonomousDatabaseShardWithDedicatedInfra) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedAutonomousDatabaseShardWithDedicatedInfra DistributedAutonomousDatabaseShardWithDedicatedInfra
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedAutonomousDatabaseShardWithDedicatedInfra
	}{
		"ADB_D",
		(MarshalTypeDistributedAutonomousDatabaseShardWithDedicatedInfra)(m),
	}

	return json.Marshal(&s)
}

// DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum Enum with underlying type: string
type DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum
const (
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusFailed                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "FAILED"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleting              DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "DELETING"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleted               DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "DELETED"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusUpdating              DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "UPDATING"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreating              DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "CREATING"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreated               DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "CREATED"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusReadyForConfiguration DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusConfigured            DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "CONFIGURED"
	DistributedAutonomousDatabaseShardWithDedicatedInfraStatusNeedsAttention        DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum = map[string]DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum{
	"FAILED":                  DistributedAutonomousDatabaseShardWithDedicatedInfraStatusFailed,
	"DELETING":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleting,
	"DELETED":                 DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleted,
	"UPDATING":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusUpdating,
	"CREATING":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreating,
	"CREATED":                 DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedAutonomousDatabaseShardWithDedicatedInfraStatusReadyForConfiguration,
	"CONFIGURED":              DistributedAutonomousDatabaseShardWithDedicatedInfraStatusConfigured,
	"NEEDS_ATTENTION":         DistributedAutonomousDatabaseShardWithDedicatedInfraStatusNeedsAttention,
}

var mappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumLowerCase = map[string]DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum{
	"failed":                  DistributedAutonomousDatabaseShardWithDedicatedInfraStatusFailed,
	"deleting":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleting,
	"deleted":                 DistributedAutonomousDatabaseShardWithDedicatedInfraStatusDeleted,
	"updating":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusUpdating,
	"creating":                DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreating,
	"created":                 DistributedAutonomousDatabaseShardWithDedicatedInfraStatusCreated,
	"ready_for_configuration": DistributedAutonomousDatabaseShardWithDedicatedInfraStatusReadyForConfiguration,
	"configured":              DistributedAutonomousDatabaseShardWithDedicatedInfraStatusConfigured,
	"needs_attention":         DistributedAutonomousDatabaseShardWithDedicatedInfraStatusNeedsAttention,
}

// GetDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumValues Enumerates the set of values for DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum
func GetDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumValues() []DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum {
	values := make([]DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum
func GetDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumStringValues() []string {
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

// GetMappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum(val string) (DistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseShardWithDedicatedInfraStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
