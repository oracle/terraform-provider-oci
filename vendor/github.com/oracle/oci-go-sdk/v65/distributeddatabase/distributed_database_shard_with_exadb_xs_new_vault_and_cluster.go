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

// DistributedDatabaseShardWithExadbXsNewVaultAndCluster Globally distributed database shard based on ExaDbXs.
type DistributedDatabaseShardWithExadbXsNewVaultAndCluster struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the shardGroup for the shard.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

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

	// The name of the availability domain that the distributed database shard will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	// Peer details for the shard.
	PeerDetails []ShardPeerWithExadbXsNewVaultAndCluster `mandatory:"false" json:"peerDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Status of EXADB_XS based shard.
	Status DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`
}

// GetName returns Name
func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedDatabaseShardWithExadbXsNewVaultAndCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseShardWithExadbXsNewVaultAndCluster DistributedDatabaseShardWithExadbXsNewVaultAndCluster
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseShardWithExadbXsNewVaultAndCluster
	}{
		"NEW_VAULT_AND_CLUSTER",
		(MarshalTypeDistributedDatabaseShardWithExadbXsNewVaultAndCluster)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum
const (
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusFailed                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleting              DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleted               DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusUpdating              DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreating              DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreated               DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusReadyForConfiguration DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusConfigured            DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusNeedsAttention        DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum = map[string]DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumValues() []DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum(val string) (DistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
