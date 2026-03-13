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

// CatalogPeerWithExadbXs Details of the catalog peer
type CatalogPeerWithExadbXs struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VmCluster.
	VmClusterId *string `mandatory:"true" json:"vmClusterId"`

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of EXADB_XS based catalog peer.
	Status CatalogPeerWithExadbXsStatusEnum `mandatory:"true" json:"status"`

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

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m CatalogPeerWithExadbXs) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogPeerWithExadbXs) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogPeerWithExadbXsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCatalogPeerWithExadbXsStatusEnumStringValues(), ",")))
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

// CatalogPeerWithExadbXsStatusEnum Enum with underlying type: string
type CatalogPeerWithExadbXsStatusEnum string

// Set of constants representing the allowable values for CatalogPeerWithExadbXsStatusEnum
const (
	CatalogPeerWithExadbXsStatusFailed                CatalogPeerWithExadbXsStatusEnum = "FAILED"
	CatalogPeerWithExadbXsStatusDeleting              CatalogPeerWithExadbXsStatusEnum = "DELETING"
	CatalogPeerWithExadbXsStatusDeleted               CatalogPeerWithExadbXsStatusEnum = "DELETED"
	CatalogPeerWithExadbXsStatusUpdating              CatalogPeerWithExadbXsStatusEnum = "UPDATING"
	CatalogPeerWithExadbXsStatusCreating              CatalogPeerWithExadbXsStatusEnum = "CREATING"
	CatalogPeerWithExadbXsStatusCreated               CatalogPeerWithExadbXsStatusEnum = "CREATED"
	CatalogPeerWithExadbXsStatusReadyForConfiguration CatalogPeerWithExadbXsStatusEnum = "READY_FOR_CONFIGURATION"
	CatalogPeerWithExadbXsStatusConfigured            CatalogPeerWithExadbXsStatusEnum = "CONFIGURED"
	CatalogPeerWithExadbXsStatusNeedsAttention        CatalogPeerWithExadbXsStatusEnum = "NEEDS_ATTENTION"
)

var mappingCatalogPeerWithExadbXsStatusEnum = map[string]CatalogPeerWithExadbXsStatusEnum{
	"FAILED":                  CatalogPeerWithExadbXsStatusFailed,
	"DELETING":                CatalogPeerWithExadbXsStatusDeleting,
	"DELETED":                 CatalogPeerWithExadbXsStatusDeleted,
	"UPDATING":                CatalogPeerWithExadbXsStatusUpdating,
	"CREATING":                CatalogPeerWithExadbXsStatusCreating,
	"CREATED":                 CatalogPeerWithExadbXsStatusCreated,
	"READY_FOR_CONFIGURATION": CatalogPeerWithExadbXsStatusReadyForConfiguration,
	"CONFIGURED":              CatalogPeerWithExadbXsStatusConfigured,
	"NEEDS_ATTENTION":         CatalogPeerWithExadbXsStatusNeedsAttention,
}

var mappingCatalogPeerWithExadbXsStatusEnumLowerCase = map[string]CatalogPeerWithExadbXsStatusEnum{
	"failed":                  CatalogPeerWithExadbXsStatusFailed,
	"deleting":                CatalogPeerWithExadbXsStatusDeleting,
	"deleted":                 CatalogPeerWithExadbXsStatusDeleted,
	"updating":                CatalogPeerWithExadbXsStatusUpdating,
	"creating":                CatalogPeerWithExadbXsStatusCreating,
	"created":                 CatalogPeerWithExadbXsStatusCreated,
	"ready_for_configuration": CatalogPeerWithExadbXsStatusReadyForConfiguration,
	"configured":              CatalogPeerWithExadbXsStatusConfigured,
	"needs_attention":         CatalogPeerWithExadbXsStatusNeedsAttention,
}

// GetCatalogPeerWithExadbXsStatusEnumValues Enumerates the set of values for CatalogPeerWithExadbXsStatusEnum
func GetCatalogPeerWithExadbXsStatusEnumValues() []CatalogPeerWithExadbXsStatusEnum {
	values := make([]CatalogPeerWithExadbXsStatusEnum, 0)
	for _, v := range mappingCatalogPeerWithExadbXsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogPeerWithExadbXsStatusEnumStringValues Enumerates the set of values in String for CatalogPeerWithExadbXsStatusEnum
func GetCatalogPeerWithExadbXsStatusEnumStringValues() []string {
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

// GetMappingCatalogPeerWithExadbXsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogPeerWithExadbXsStatusEnum(val string) (CatalogPeerWithExadbXsStatusEnum, bool) {
	enum, ok := mappingCatalogPeerWithExadbXsStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
