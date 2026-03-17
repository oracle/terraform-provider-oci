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

// DistributedDatabaseShardWithExadbXs Globally distributed database shard based on exadbxs.
type DistributedDatabaseShardWithExadbXs struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the shardGroup for the shard.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created.
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

	// Peer details for the shard.
	PeerDetails []ShardPeerWithExadbXs `mandatory:"false" json:"peerDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Status of EXADB_XS based shard.
	Status DistributedDatabaseShardWithExadbXsStatusEnum `mandatory:"true" json:"status"`
}

// GetName returns Name
func (m DistributedDatabaseShardWithExadbXs) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseShardWithExadbXs) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseShardWithExadbXs) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedDatabaseShardWithExadbXs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseShardWithExadbXs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseShardWithExadbXsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseShardWithExadbXsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedDatabaseShardWithExadbXs) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseShardWithExadbXs DistributedDatabaseShardWithExadbXs
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseShardWithExadbXs
	}{
		"EXADB_XS",
		(MarshalTypeDistributedDatabaseShardWithExadbXs)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseShardWithExadbXsStatusEnum Enum with underlying type: string
type DistributedDatabaseShardWithExadbXsStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardWithExadbXsStatusEnum
const (
	DistributedDatabaseShardWithExadbXsStatusFailed                DistributedDatabaseShardWithExadbXsStatusEnum = "FAILED"
	DistributedDatabaseShardWithExadbXsStatusDeleting              DistributedDatabaseShardWithExadbXsStatusEnum = "DELETING"
	DistributedDatabaseShardWithExadbXsStatusDeleted               DistributedDatabaseShardWithExadbXsStatusEnum = "DELETED"
	DistributedDatabaseShardWithExadbXsStatusUpdating              DistributedDatabaseShardWithExadbXsStatusEnum = "UPDATING"
	DistributedDatabaseShardWithExadbXsStatusCreating              DistributedDatabaseShardWithExadbXsStatusEnum = "CREATING"
	DistributedDatabaseShardWithExadbXsStatusCreated               DistributedDatabaseShardWithExadbXsStatusEnum = "CREATED"
	DistributedDatabaseShardWithExadbXsStatusReadyForConfiguration DistributedDatabaseShardWithExadbXsStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseShardWithExadbXsStatusConfigured            DistributedDatabaseShardWithExadbXsStatusEnum = "CONFIGURED"
	DistributedDatabaseShardWithExadbXsStatusNeedsAttention        DistributedDatabaseShardWithExadbXsStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseShardWithExadbXsStatusEnum = map[string]DistributedDatabaseShardWithExadbXsStatusEnum{
	"FAILED":                  DistributedDatabaseShardWithExadbXsStatusFailed,
	"DELETING":                DistributedDatabaseShardWithExadbXsStatusDeleting,
	"DELETED":                 DistributedDatabaseShardWithExadbXsStatusDeleted,
	"UPDATING":                DistributedDatabaseShardWithExadbXsStatusUpdating,
	"CREATING":                DistributedDatabaseShardWithExadbXsStatusCreating,
	"CREATED":                 DistributedDatabaseShardWithExadbXsStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseShardWithExadbXsStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseShardWithExadbXsStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseShardWithExadbXsStatusNeedsAttention,
}

var mappingDistributedDatabaseShardWithExadbXsStatusEnumLowerCase = map[string]DistributedDatabaseShardWithExadbXsStatusEnum{
	"failed":                  DistributedDatabaseShardWithExadbXsStatusFailed,
	"deleting":                DistributedDatabaseShardWithExadbXsStatusDeleting,
	"deleted":                 DistributedDatabaseShardWithExadbXsStatusDeleted,
	"updating":                DistributedDatabaseShardWithExadbXsStatusUpdating,
	"creating":                DistributedDatabaseShardWithExadbXsStatusCreating,
	"created":                 DistributedDatabaseShardWithExadbXsStatusCreated,
	"ready_for_configuration": DistributedDatabaseShardWithExadbXsStatusReadyForConfiguration,
	"configured":              DistributedDatabaseShardWithExadbXsStatusConfigured,
	"needs_attention":         DistributedDatabaseShardWithExadbXsStatusNeedsAttention,
}

// GetDistributedDatabaseShardWithExadbXsStatusEnumValues Enumerates the set of values for DistributedDatabaseShardWithExadbXsStatusEnum
func GetDistributedDatabaseShardWithExadbXsStatusEnumValues() []DistributedDatabaseShardWithExadbXsStatusEnum {
	values := make([]DistributedDatabaseShardWithExadbXsStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseShardWithExadbXsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardWithExadbXsStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardWithExadbXsStatusEnum
func GetDistributedDatabaseShardWithExadbXsStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseShardWithExadbXsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardWithExadbXsStatusEnum(val string) (DistributedDatabaseShardWithExadbXsStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardWithExadbXsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
