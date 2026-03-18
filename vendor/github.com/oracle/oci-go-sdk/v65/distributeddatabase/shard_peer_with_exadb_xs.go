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

// ShardPeerWithExadbXs Details of the shard peer
type ShardPeerWithExadbXs struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of EXADB_XS based shard peer.
	Status ShardPeerWithExadbXsStatusEnum `mandatory:"true" json:"status"`

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

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m ShardPeerWithExadbXs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShardPeerWithExadbXs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardPeerWithExadbXsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetShardPeerWithExadbXsStatusEnumStringValues(), ",")))
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

// ShardPeerWithExadbXsStatusEnum Enum with underlying type: string
type ShardPeerWithExadbXsStatusEnum string

// Set of constants representing the allowable values for ShardPeerWithExadbXsStatusEnum
const (
	ShardPeerWithExadbXsStatusFailed                ShardPeerWithExadbXsStatusEnum = "FAILED"
	ShardPeerWithExadbXsStatusDeleting              ShardPeerWithExadbXsStatusEnum = "DELETING"
	ShardPeerWithExadbXsStatusDeleted               ShardPeerWithExadbXsStatusEnum = "DELETED"
	ShardPeerWithExadbXsStatusUpdating              ShardPeerWithExadbXsStatusEnum = "UPDATING"
	ShardPeerWithExadbXsStatusCreating              ShardPeerWithExadbXsStatusEnum = "CREATING"
	ShardPeerWithExadbXsStatusCreated               ShardPeerWithExadbXsStatusEnum = "CREATED"
	ShardPeerWithExadbXsStatusReadyForConfiguration ShardPeerWithExadbXsStatusEnum = "READY_FOR_CONFIGURATION"
	ShardPeerWithExadbXsStatusConfigured            ShardPeerWithExadbXsStatusEnum = "CONFIGURED"
	ShardPeerWithExadbXsStatusNeedsAttention        ShardPeerWithExadbXsStatusEnum = "NEEDS_ATTENTION"
)

var mappingShardPeerWithExadbXsStatusEnum = map[string]ShardPeerWithExadbXsStatusEnum{
	"FAILED":                  ShardPeerWithExadbXsStatusFailed,
	"DELETING":                ShardPeerWithExadbXsStatusDeleting,
	"DELETED":                 ShardPeerWithExadbXsStatusDeleted,
	"UPDATING":                ShardPeerWithExadbXsStatusUpdating,
	"CREATING":                ShardPeerWithExadbXsStatusCreating,
	"CREATED":                 ShardPeerWithExadbXsStatusCreated,
	"READY_FOR_CONFIGURATION": ShardPeerWithExadbXsStatusReadyForConfiguration,
	"CONFIGURED":              ShardPeerWithExadbXsStatusConfigured,
	"NEEDS_ATTENTION":         ShardPeerWithExadbXsStatusNeedsAttention,
}

var mappingShardPeerWithExadbXsStatusEnumLowerCase = map[string]ShardPeerWithExadbXsStatusEnum{
	"failed":                  ShardPeerWithExadbXsStatusFailed,
	"deleting":                ShardPeerWithExadbXsStatusDeleting,
	"deleted":                 ShardPeerWithExadbXsStatusDeleted,
	"updating":                ShardPeerWithExadbXsStatusUpdating,
	"creating":                ShardPeerWithExadbXsStatusCreating,
	"created":                 ShardPeerWithExadbXsStatusCreated,
	"ready_for_configuration": ShardPeerWithExadbXsStatusReadyForConfiguration,
	"configured":              ShardPeerWithExadbXsStatusConfigured,
	"needs_attention":         ShardPeerWithExadbXsStatusNeedsAttention,
}

// GetShardPeerWithExadbXsStatusEnumValues Enumerates the set of values for ShardPeerWithExadbXsStatusEnum
func GetShardPeerWithExadbXsStatusEnumValues() []ShardPeerWithExadbXsStatusEnum {
	values := make([]ShardPeerWithExadbXsStatusEnum, 0)
	for _, v := range mappingShardPeerWithExadbXsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetShardPeerWithExadbXsStatusEnumStringValues Enumerates the set of values in String for ShardPeerWithExadbXsStatusEnum
func GetShardPeerWithExadbXsStatusEnumStringValues() []string {
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

// GetMappingShardPeerWithExadbXsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardPeerWithExadbXsStatusEnum(val string) (ShardPeerWithExadbXsStatusEnum, bool) {
	enum, ok := mappingShardPeerWithExadbXsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
