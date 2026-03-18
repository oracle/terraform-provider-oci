// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Globally Distributed Database
//
// Use the Globally Distributed Database service APIs to create and manage the Globally distributed databases.
//

package distributeddatabase

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ShardPeerWithExadbXsNewVaultAndCluster Details of the shard peer
type ShardPeerWithExadbXsNewVaultAndCluster struct {

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of EXADB_XS based shard peer.
	Status ShardPeerWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`

	// The time the shard peer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard peer was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// The protectionMode for the shard peer.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// The name of the availability domain that the distributed database shard will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m ShardPeerWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShardPeerWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardPeerWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetShardPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedDbProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingDistributedDbTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetDistributedDbTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShardPeerWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type ShardPeerWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for ShardPeerWithExadbXsNewVaultAndClusterStatusEnum
const (
	ShardPeerWithExadbXsNewVaultAndClusterStatusFailed                ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	ShardPeerWithExadbXsNewVaultAndClusterStatusDeleting              ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	ShardPeerWithExadbXsNewVaultAndClusterStatusDeleted               ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	ShardPeerWithExadbXsNewVaultAndClusterStatusUpdating              ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	ShardPeerWithExadbXsNewVaultAndClusterStatusCreating              ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	ShardPeerWithExadbXsNewVaultAndClusterStatusCreated               ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	ShardPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	ShardPeerWithExadbXsNewVaultAndClusterStatusConfigured            ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	ShardPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention        ShardPeerWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingShardPeerWithExadbXsNewVaultAndClusterStatusEnum = map[string]ShardPeerWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  ShardPeerWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                ShardPeerWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 ShardPeerWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                ShardPeerWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                ShardPeerWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 ShardPeerWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": ShardPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              ShardPeerWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         ShardPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingShardPeerWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]ShardPeerWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  ShardPeerWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                ShardPeerWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 ShardPeerWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                ShardPeerWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                ShardPeerWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 ShardPeerWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": ShardPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              ShardPeerWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         ShardPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetShardPeerWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for ShardPeerWithExadbXsNewVaultAndClusterStatusEnum
func GetShardPeerWithExadbXsNewVaultAndClusterStatusEnumValues() []ShardPeerWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]ShardPeerWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingShardPeerWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetShardPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for ShardPeerWithExadbXsNewVaultAndClusterStatusEnum
func GetShardPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingShardPeerWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardPeerWithExadbXsNewVaultAndClusterStatusEnum(val string) (ShardPeerWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingShardPeerWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
