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

// CatalogPeerWithBaseDb Details of the catalog peer
type CatalogPeerWithBaseDb struct {

	// The name of the availability domain that the base database system will be located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of Base database based catalog peer.
	Status CatalogPeerWithBaseDbStatusEnum `mandatory:"true" json:"status"`

	// Identifier of the subnet in which peer catalog exists.
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

	// The redo transport type to use for Data Guard association for Base database based catalog.
	TransportType BaseDbTransportTypeEnum `mandatory:"false" json:"transportType,omitempty"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`
}

func (m CatalogPeerWithBaseDb) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogPeerWithBaseDb) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogPeerWithBaseDbStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCatalogPeerWithBaseDbStatusEnumStringValues(), ",")))
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

// CatalogPeerWithBaseDbStatusEnum Enum with underlying type: string
type CatalogPeerWithBaseDbStatusEnum string

// Set of constants representing the allowable values for CatalogPeerWithBaseDbStatusEnum
const (
	CatalogPeerWithBaseDbStatusFailed                CatalogPeerWithBaseDbStatusEnum = "FAILED"
	CatalogPeerWithBaseDbStatusDeleting              CatalogPeerWithBaseDbStatusEnum = "DELETING"
	CatalogPeerWithBaseDbStatusDeleted               CatalogPeerWithBaseDbStatusEnum = "DELETED"
	CatalogPeerWithBaseDbStatusUpdating              CatalogPeerWithBaseDbStatusEnum = "UPDATING"
	CatalogPeerWithBaseDbStatusCreating              CatalogPeerWithBaseDbStatusEnum = "CREATING"
	CatalogPeerWithBaseDbStatusCreated               CatalogPeerWithBaseDbStatusEnum = "CREATED"
	CatalogPeerWithBaseDbStatusReadyForConfiguration CatalogPeerWithBaseDbStatusEnum = "READY_FOR_CONFIGURATION"
	CatalogPeerWithBaseDbStatusConfigured            CatalogPeerWithBaseDbStatusEnum = "CONFIGURED"
	CatalogPeerWithBaseDbStatusNeedsAttention        CatalogPeerWithBaseDbStatusEnum = "NEEDS_ATTENTION"
)

var mappingCatalogPeerWithBaseDbStatusEnum = map[string]CatalogPeerWithBaseDbStatusEnum{
	"FAILED":                  CatalogPeerWithBaseDbStatusFailed,
	"DELETING":                CatalogPeerWithBaseDbStatusDeleting,
	"DELETED":                 CatalogPeerWithBaseDbStatusDeleted,
	"UPDATING":                CatalogPeerWithBaseDbStatusUpdating,
	"CREATING":                CatalogPeerWithBaseDbStatusCreating,
	"CREATED":                 CatalogPeerWithBaseDbStatusCreated,
	"READY_FOR_CONFIGURATION": CatalogPeerWithBaseDbStatusReadyForConfiguration,
	"CONFIGURED":              CatalogPeerWithBaseDbStatusConfigured,
	"NEEDS_ATTENTION":         CatalogPeerWithBaseDbStatusNeedsAttention,
}

var mappingCatalogPeerWithBaseDbStatusEnumLowerCase = map[string]CatalogPeerWithBaseDbStatusEnum{
	"failed":                  CatalogPeerWithBaseDbStatusFailed,
	"deleting":                CatalogPeerWithBaseDbStatusDeleting,
	"deleted":                 CatalogPeerWithBaseDbStatusDeleted,
	"updating":                CatalogPeerWithBaseDbStatusUpdating,
	"creating":                CatalogPeerWithBaseDbStatusCreating,
	"created":                 CatalogPeerWithBaseDbStatusCreated,
	"ready_for_configuration": CatalogPeerWithBaseDbStatusReadyForConfiguration,
	"configured":              CatalogPeerWithBaseDbStatusConfigured,
	"needs_attention":         CatalogPeerWithBaseDbStatusNeedsAttention,
}

// GetCatalogPeerWithBaseDbStatusEnumValues Enumerates the set of values for CatalogPeerWithBaseDbStatusEnum
func GetCatalogPeerWithBaseDbStatusEnumValues() []CatalogPeerWithBaseDbStatusEnum {
	values := make([]CatalogPeerWithBaseDbStatusEnum, 0)
	for _, v := range mappingCatalogPeerWithBaseDbStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogPeerWithBaseDbStatusEnumStringValues Enumerates the set of values in String for CatalogPeerWithBaseDbStatusEnum
func GetCatalogPeerWithBaseDbStatusEnumStringValues() []string {
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

// GetMappingCatalogPeerWithBaseDbStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogPeerWithBaseDbStatusEnum(val string) (CatalogPeerWithBaseDbStatusEnum, bool) {
	enum, ok := mappingCatalogPeerWithBaseDbStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
