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

// ShardPeerWithBaseDb Details of the shard peer
type ShardPeerWithBaseDb struct {

	// The name of the availability domain that the peer base database system will be located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of shard peer for the Globally distributed base database.
	Status ShardPeerWithBaseDbStatusEnum `mandatory:"true" json:"status"`

	// Identifier of the subnet in which peer shard exists.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The time the shard peer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard peer was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the fault domain that the peer base database system will be located in.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// FLag to enable active Data Guard.
	IsActiveDataGuardEnabled *bool `mandatory:"false" json:"isActiveDataGuardEnabled"`

	// The protection mode for the Data Guard association's primary and standby Base database based shard.
	ProtectionMode BaseDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for Data Guard association for Base database based shard.
	TransportType BaseDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m ShardPeerWithBaseDb) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ShardPeerWithBaseDb) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingShardPeerWithBaseDbStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetShardPeerWithBaseDbStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetBaseDbProtectionModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbTransportTypeEnum(string(m.TransportType)); !ok && m.TransportType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TransportType: %s. Supported values are: %s.", m.TransportType, strings.Join(GetBaseDbTransportTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ShardPeerWithBaseDbStatusEnum Enum with underlying type: string
type ShardPeerWithBaseDbStatusEnum string

// Set of constants representing the allowable values for ShardPeerWithBaseDbStatusEnum
const (
	ShardPeerWithBaseDbStatusFailed                ShardPeerWithBaseDbStatusEnum = "FAILED"
	ShardPeerWithBaseDbStatusDeleting              ShardPeerWithBaseDbStatusEnum = "DELETING"
	ShardPeerWithBaseDbStatusDeleted               ShardPeerWithBaseDbStatusEnum = "DELETED"
	ShardPeerWithBaseDbStatusUpdating              ShardPeerWithBaseDbStatusEnum = "UPDATING"
	ShardPeerWithBaseDbStatusCreating              ShardPeerWithBaseDbStatusEnum = "CREATING"
	ShardPeerWithBaseDbStatusCreated               ShardPeerWithBaseDbStatusEnum = "CREATED"
	ShardPeerWithBaseDbStatusReadyForConfiguration ShardPeerWithBaseDbStatusEnum = "READY_FOR_CONFIGURATION"
	ShardPeerWithBaseDbStatusConfigured            ShardPeerWithBaseDbStatusEnum = "CONFIGURED"
	ShardPeerWithBaseDbStatusNeedsAttention        ShardPeerWithBaseDbStatusEnum = "NEEDS_ATTENTION"
)

var mappingShardPeerWithBaseDbStatusEnum = map[string]ShardPeerWithBaseDbStatusEnum{
	"FAILED":                  ShardPeerWithBaseDbStatusFailed,
	"DELETING":                ShardPeerWithBaseDbStatusDeleting,
	"DELETED":                 ShardPeerWithBaseDbStatusDeleted,
	"UPDATING":                ShardPeerWithBaseDbStatusUpdating,
	"CREATING":                ShardPeerWithBaseDbStatusCreating,
	"CREATED":                 ShardPeerWithBaseDbStatusCreated,
	"READY_FOR_CONFIGURATION": ShardPeerWithBaseDbStatusReadyForConfiguration,
	"CONFIGURED":              ShardPeerWithBaseDbStatusConfigured,
	"NEEDS_ATTENTION":         ShardPeerWithBaseDbStatusNeedsAttention,
}

var mappingShardPeerWithBaseDbStatusEnumLowerCase = map[string]ShardPeerWithBaseDbStatusEnum{
	"failed":                  ShardPeerWithBaseDbStatusFailed,
	"deleting":                ShardPeerWithBaseDbStatusDeleting,
	"deleted":                 ShardPeerWithBaseDbStatusDeleted,
	"updating":                ShardPeerWithBaseDbStatusUpdating,
	"creating":                ShardPeerWithBaseDbStatusCreating,
	"created":                 ShardPeerWithBaseDbStatusCreated,
	"ready_for_configuration": ShardPeerWithBaseDbStatusReadyForConfiguration,
	"configured":              ShardPeerWithBaseDbStatusConfigured,
	"needs_attention":         ShardPeerWithBaseDbStatusNeedsAttention,
}

// GetShardPeerWithBaseDbStatusEnumValues Enumerates the set of values for ShardPeerWithBaseDbStatusEnum
func GetShardPeerWithBaseDbStatusEnumValues() []ShardPeerWithBaseDbStatusEnum {
	values := make([]ShardPeerWithBaseDbStatusEnum, 0)
	for _, v := range mappingShardPeerWithBaseDbStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetShardPeerWithBaseDbStatusEnumStringValues Enumerates the set of values in String for ShardPeerWithBaseDbStatusEnum
func GetShardPeerWithBaseDbStatusEnumStringValues() []string {
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

// GetMappingShardPeerWithBaseDbStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingShardPeerWithBaseDbStatusEnum(val string) (ShardPeerWithBaseDbStatusEnum, bool) {
	enum, ok := mappingShardPeerWithBaseDbStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
