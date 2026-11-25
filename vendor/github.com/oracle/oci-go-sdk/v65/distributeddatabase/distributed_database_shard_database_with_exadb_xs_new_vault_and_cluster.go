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

// DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster Details of a distributed database shard on ExaDB-XS with a new cluster and storage vault.
type DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
	// This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// The name of the availability domain that the distributed database shard will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	// Status of distributed database shard.
	Status DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
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

// MarshalJSON marshals to json representation
func (m DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster
	}{
		"XS_NEW_VAULT_AND_CLUSTER",
		(MarshalTypeDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndCluster)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum
const (
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusFailed                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleting              DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleted               DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusUpdating              DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreating              DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreated               DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusConfigured            DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention        DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum = map[string]DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumValues() []DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum(val string) (DistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardDatabaseWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
