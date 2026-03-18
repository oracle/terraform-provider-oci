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

// CatalogPeerWithDedicatedInfra Details of the catalog peer with dedicated infrastructure.
type CatalogPeerWithDedicatedInfra struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloudAutonomousVmCluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	// The name of the shardGroup for the peer.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// Status of catalog with dedicated infrastructure for the Globally distributed autonomous database.
	Status CatalogPeerWithDedicatedInfraStatusEnum `mandatory:"true" json:"status"`

	// The time the catalog peer was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog peer was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// The protectionMode for the shard peer.
	ProtectionMode DistributedAutonomousDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`

	// The lag time for my preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// This field is deprecated. Support for this field will be removed after one year of deprecation cycle.
	IsAutomaticFailoverEnabled *bool `mandatory:"false" json:"isAutomaticFailoverEnabled"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before schedlued maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	Metadata *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`
}

func (m CatalogPeerWithDedicatedInfra) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CatalogPeerWithDedicatedInfra) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCatalogPeerWithDedicatedInfraStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCatalogPeerWithDedicatedInfraStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedAutonomousDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedAutonomousDbProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CatalogPeerWithDedicatedInfraStatusEnum Enum with underlying type: string
type CatalogPeerWithDedicatedInfraStatusEnum string

// Set of constants representing the allowable values for CatalogPeerWithDedicatedInfraStatusEnum
const (
	CatalogPeerWithDedicatedInfraStatusFailed                CatalogPeerWithDedicatedInfraStatusEnum = "FAILED"
	CatalogPeerWithDedicatedInfraStatusDeleting              CatalogPeerWithDedicatedInfraStatusEnum = "DELETING"
	CatalogPeerWithDedicatedInfraStatusDeleted               CatalogPeerWithDedicatedInfraStatusEnum = "DELETED"
	CatalogPeerWithDedicatedInfraStatusUpdating              CatalogPeerWithDedicatedInfraStatusEnum = "UPDATING"
	CatalogPeerWithDedicatedInfraStatusCreating              CatalogPeerWithDedicatedInfraStatusEnum = "CREATING"
	CatalogPeerWithDedicatedInfraStatusCreated               CatalogPeerWithDedicatedInfraStatusEnum = "CREATED"
	CatalogPeerWithDedicatedInfraStatusReadyForConfiguration CatalogPeerWithDedicatedInfraStatusEnum = "READY_FOR_CONFIGURATION"
	CatalogPeerWithDedicatedInfraStatusConfigured            CatalogPeerWithDedicatedInfraStatusEnum = "CONFIGURED"
	CatalogPeerWithDedicatedInfraStatusNeedsAttention        CatalogPeerWithDedicatedInfraStatusEnum = "NEEDS_ATTENTION"
)

var mappingCatalogPeerWithDedicatedInfraStatusEnum = map[string]CatalogPeerWithDedicatedInfraStatusEnum{
	"FAILED":                  CatalogPeerWithDedicatedInfraStatusFailed,
	"DELETING":                CatalogPeerWithDedicatedInfraStatusDeleting,
	"DELETED":                 CatalogPeerWithDedicatedInfraStatusDeleted,
	"UPDATING":                CatalogPeerWithDedicatedInfraStatusUpdating,
	"CREATING":                CatalogPeerWithDedicatedInfraStatusCreating,
	"CREATED":                 CatalogPeerWithDedicatedInfraStatusCreated,
	"READY_FOR_CONFIGURATION": CatalogPeerWithDedicatedInfraStatusReadyForConfiguration,
	"CONFIGURED":              CatalogPeerWithDedicatedInfraStatusConfigured,
	"NEEDS_ATTENTION":         CatalogPeerWithDedicatedInfraStatusNeedsAttention,
}

var mappingCatalogPeerWithDedicatedInfraStatusEnumLowerCase = map[string]CatalogPeerWithDedicatedInfraStatusEnum{
	"failed":                  CatalogPeerWithDedicatedInfraStatusFailed,
	"deleting":                CatalogPeerWithDedicatedInfraStatusDeleting,
	"deleted":                 CatalogPeerWithDedicatedInfraStatusDeleted,
	"updating":                CatalogPeerWithDedicatedInfraStatusUpdating,
	"creating":                CatalogPeerWithDedicatedInfraStatusCreating,
	"created":                 CatalogPeerWithDedicatedInfraStatusCreated,
	"ready_for_configuration": CatalogPeerWithDedicatedInfraStatusReadyForConfiguration,
	"configured":              CatalogPeerWithDedicatedInfraStatusConfigured,
	"needs_attention":         CatalogPeerWithDedicatedInfraStatusNeedsAttention,
}

// GetCatalogPeerWithDedicatedInfraStatusEnumValues Enumerates the set of values for CatalogPeerWithDedicatedInfraStatusEnum
func GetCatalogPeerWithDedicatedInfraStatusEnumValues() []CatalogPeerWithDedicatedInfraStatusEnum {
	values := make([]CatalogPeerWithDedicatedInfraStatusEnum, 0)
	for _, v := range mappingCatalogPeerWithDedicatedInfraStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCatalogPeerWithDedicatedInfraStatusEnumStringValues Enumerates the set of values in String for CatalogPeerWithDedicatedInfraStatusEnum
func GetCatalogPeerWithDedicatedInfraStatusEnumStringValues() []string {
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

// GetMappingCatalogPeerWithDedicatedInfraStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCatalogPeerWithDedicatedInfraStatusEnum(val string) (CatalogPeerWithDedicatedInfraStatusEnum, bool) {
	enum, ok := mappingCatalogPeerWithDedicatedInfraStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
