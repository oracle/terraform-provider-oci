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

// DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster Details of a distributed database catalog on ExaDB-XS with a new cluster.
type DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster struct {

	// The name of the catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the exascale db storage vault.
	ExascaleDbStorageVaultId *string `mandatory:"true" json:"exascaleDbStorageVaultId"`

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

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	// Status of the distributed database catalog.
	Status DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum `mandatory:"true" json:"status"`

	// The protection mode used for the Data Guard association.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The transport type used for the Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) GetMetadata() *DistributedDbMetadata {
	return m.Metadata
}

func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumStringValues(), ",")))
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
func (m DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewCluster DistributedDatabaseCatalogDatabaseWithExadbXsNewCluster
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewCluster
	}{
		"XS_NEW_CLUSTER",
		(MarshalTypeDistributedDatabaseCatalogDatabaseWithExadbXsNewCluster)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum Enum with underlying type: string
type DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum
const (
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusFailed                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "FAILED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleting              DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "DELETING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleted               DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "DELETED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusUpdating              DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "UPDATING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreating              DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "CREATING"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreated               DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "CREATED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusReadyForConfiguration DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusConfigured            DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "CONFIGURED"
	DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusNeedsAttention        DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum = map[string]DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum{
	"FAILED":                  DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusFailed,
	"DELETING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleting,
	"DELETED":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleted,
	"UPDATING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusUpdating,
	"CREATING":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreating,
	"CREATED":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusNeedsAttention,
}

var mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumLowerCase = map[string]DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum{
	"failed":                  DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusFailed,
	"deleting":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleting,
	"deleted":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusDeleted,
	"updating":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusUpdating,
	"creating":                DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreating,
	"created":                 DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusCreated,
	"ready_for_configuration": DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusReadyForConfiguration,
	"configured":              DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusConfigured,
	"needs_attention":         DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusNeedsAttention,
}

// GetDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumValues Enumerates the set of values for DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumValues() []DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum {
	values := make([]DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum
func GetDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum(val string) (DistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseCatalogDatabaseWithExadbXsNewClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
