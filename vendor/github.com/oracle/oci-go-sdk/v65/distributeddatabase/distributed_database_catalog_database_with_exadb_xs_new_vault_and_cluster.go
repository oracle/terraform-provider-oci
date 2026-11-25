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

// DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster Details of a distributed database catalog on ExaDB-XS with a new cluster and storage vault.
type DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster struct {

	// The name of the catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
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

	// The name of the availability domain that the distributed database catalog will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	// Status of the distributed database catalog.
	Status DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
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
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster
	}{
		"XS_NEW_VAULT_AND_CLUSTER",
		(MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndCluster)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum
const (
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusFailed                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleting              DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleted               DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusUpdating              DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreating              DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreated               DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusConfigured            DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention        DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum = map[string]DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumValues() []DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum(val string) (DistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
