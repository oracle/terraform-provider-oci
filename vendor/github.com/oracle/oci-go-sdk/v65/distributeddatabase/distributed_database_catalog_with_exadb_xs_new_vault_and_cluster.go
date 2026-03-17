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

// DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster Globally distributed database catalog based on ExaDbXs.
type DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster struct {

	// The name of catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the shardGroup for the catalog.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbHomeId *string `mandatory:"false" json:"dbHomeId"`

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

	// The name of the availability domain that the distributed database catalog will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	// Peer details for the catalog.
	PeerDetails []CatalogPeerWithExadbXsNewVaultAndCluster `mandatory:"false" json:"peerDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Status of EXADB_XS based catalog.
	Status DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`
}

// GetName returns Name
func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseCatalogWithExadbXsNewVaultAndCluster DistributedDatabaseCatalogWithExadbXsNewVaultAndCluster
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseCatalogWithExadbXsNewVaultAndCluster
	}{
		"NEW_VAULT_AND_CLUSTER",
		(MarshalTypeDistributedDatabaseCatalogWithExadbXsNewVaultAndCluster)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum
const (
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusFailed                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleting              DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleted               DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusUpdating              DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreating              DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreated               DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusReadyForConfiguration DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusConfigured            DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusNeedsAttention        DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum = map[string]DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumValues() []DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum(val string) (DistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
