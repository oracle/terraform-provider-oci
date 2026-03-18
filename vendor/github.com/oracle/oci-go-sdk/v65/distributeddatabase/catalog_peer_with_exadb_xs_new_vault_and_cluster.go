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

// CatalogPeerWithExadbXsNewVaultAndCluster Details of the catalog peer
type CatalogPeerWithExadbXsNewVaultAndCluster struct {

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of EXADB_XS based catalog peer.
	Status CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum `mandatory:"true" json:"status"`

	// The time the catalog peer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog peer was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// The protectionMode for the catalog peer.
	ProtectionMode DistributedDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The redo transport type to use for this Data Guard association.
	TransportType DistributedDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	// The name of the availability domain that the distributed database shard will be located in.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	DbStorageVaultDetails *DistributedDbStorageVault `mandatory:"false" json:"dbStorageVaultDetails"`

	VmClusterDetails *DistributedDbVmCluster `mandatory:"false" json:"vmClusterDetails"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m CatalogPeerWithExadbXsNewVaultAndCluster) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogPeerWithExadbXsNewVaultAndCluster) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues(), ",")))
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

// CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum Enum with underlying type: string
type CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum string

// Set of constants representing the allowable values for CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum
const (
	CatalogPeerWithExadbXsNewVaultAndClusterStatusFailed                CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "FAILED"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleting              CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "DELETING"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleted               CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "DELETED"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusUpdating              CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "UPDATING"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusCreating              CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "CREATING"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusCreated               CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "CREATED"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "READY_FOR_CONFIGURATION"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusConfigured            CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "CONFIGURED"
	CatalogPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention        CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = "NEEDS_ATTENTION"
)

var mappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnum = map[string]CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum{
	"FAILED":                  CatalogPeerWithExadbXsNewVaultAndClusterStatusFailed,
	"DELETING":                CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleting,
	"DELETED":                 CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleted,
	"UPDATING":                CatalogPeerWithExadbXsNewVaultAndClusterStatusUpdating,
	"CREATING":                CatalogPeerWithExadbXsNewVaultAndClusterStatusCreating,
	"CREATED":                 CatalogPeerWithExadbXsNewVaultAndClusterStatusCreated,
	"READY_FOR_CONFIGURATION": CatalogPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"CONFIGURED":              CatalogPeerWithExadbXsNewVaultAndClusterStatusConfigured,
	"NEEDS_ATTENTION":         CatalogPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

var mappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumLowerCase = map[string]CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum{
	"failed":                  CatalogPeerWithExadbXsNewVaultAndClusterStatusFailed,
	"deleting":                CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleting,
	"deleted":                 CatalogPeerWithExadbXsNewVaultAndClusterStatusDeleted,
	"updating":                CatalogPeerWithExadbXsNewVaultAndClusterStatusUpdating,
	"creating":                CatalogPeerWithExadbXsNewVaultAndClusterStatusCreating,
	"created":                 CatalogPeerWithExadbXsNewVaultAndClusterStatusCreated,
	"ready_for_configuration": CatalogPeerWithExadbXsNewVaultAndClusterStatusReadyForConfiguration,
	"configured":              CatalogPeerWithExadbXsNewVaultAndClusterStatusConfigured,
	"needs_attention":         CatalogPeerWithExadbXsNewVaultAndClusterStatusNeedsAttention,
}

// GetCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumValues Enumerates the set of values for CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum
func GetCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumValues() []CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum {
	values := make([]CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum, 0)
	for _, v := range mappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues Enumerates the set of values in String for CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum
func GetCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumStringValues() []string {
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

// GetMappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnum(val string) (CatalogPeerWithExadbXsNewVaultAndClusterStatusEnum, bool) {
	enum, ok := mappingCatalogPeerWithExadbXsNewVaultAndClusterStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
