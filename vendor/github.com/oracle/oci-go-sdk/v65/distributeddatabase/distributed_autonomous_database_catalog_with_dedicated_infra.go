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

// DistributedAutonomousDatabaseCatalogWithDedicatedInfra Details of a distributed autonomous database catalog on dedicated infrastructure.
type DistributedAutonomousDatabaseCatalogWithDedicatedInfra struct {

	// The name of the catalog.
	Name *string `mandatory:"true" json:"name"`

	// The time the catalog was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the catalog was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The compute count for the catalog database. It has to be in multiples of 2.
	ComputeCount *float32 `mandatory:"true" json:"computeCount"`

	// The data disk group size to be allocated in GBs for the catalog database.
	DataStorageSizeInGbs *float64 `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Autonomous VM Cluster.
	CloudAutonomousVmClusterId *string `mandatory:"true" json:"cloudAutonomousVmClusterId"`

	Metadata *DistributedAutonomousDbMetadata `mandatory:"false" json:"metadata"`

	// Indicates if vertical auto scaling is enabled for the Autonomous AI Database CPU core count.
	// The default value is `FALSE`.
	IsAutoScalingEnabled *bool `mandatory:"false" json:"isAutoScalingEnabled"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure vault (https://docs.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `kmsKeyId` are required for Customer Managed Keys.
	VaultId *string `mandatory:"false" json:"vaultId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions.
	KmsKeyVersionId *string `mandatory:"false" json:"kmsKeyVersionId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the key store used to create the catalog.
	OkvKeyStoreId *string `mandatory:"false" json:"okvKeyStoreId"`

	// The OKV endpoint name.
	OkvEndPointGroupName *string `mandatory:"false" json:"okvEndPointGroupName"`

	// The lag time preference based on data loss tolerance in seconds.
	FastStartFailOverLagLimitInSeconds *int `mandatory:"false" json:"fastStartFailOverLagLimitInSeconds"`

	// The scheduling detail for the quarterly maintenance window of the standby Autonomous Container Database.
	// This value represents the number of days before scheduled maintenance of the primary database.
	StandbyMaintenanceBufferInDays *int `mandatory:"false" json:"standbyMaintenanceBufferInDays"`

	// the identifier of the underlying supporting resource.
	SupportingResourceId *string `mandatory:"false" json:"supportingResourceId"`

	// the identifier of the container database for underlying supporting resource.
	ContainerDatabaseId *string `mandatory:"false" json:"containerDatabaseId"`

	// Status of the distributed autonomous database catalog.
	Status DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum `mandatory:"true" json:"status"`

	// The protection mode for the catalog peer.
	ProtectionMode DistributedAutonomousDbProtectionModeEnum `mandatory:"false" json:"protectionMode,omitempty"`
}

// GetName returns Name
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetMetadata returns Metadata
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) GetMetadata() *DistributedAutonomousDbMetadata {
	return m.Metadata
}

func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDistributedAutonomousDbProtectionModeEnum(string(m.ProtectionMode)); !ok && m.ProtectionMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProtectionMode: %s. Supported values are: %s.", m.ProtectionMode, strings.Join(GetDistributedAutonomousDbProtectionModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedAutonomousDatabaseCatalogWithDedicatedInfra) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedAutonomousDatabaseCatalogWithDedicatedInfra DistributedAutonomousDatabaseCatalogWithDedicatedInfra
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedAutonomousDatabaseCatalogWithDedicatedInfra
	}{
		"ADBD_EXISTING_CLUSTER",
		(MarshalTypeDistributedAutonomousDatabaseCatalogWithDedicatedInfra)(m),
	}

	return json.Marshal(&s)
}

// DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum Enum with underlying type: string
type DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum string

// Set of constants representing the allowable values for DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum
const (
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusFailed                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "FAILED"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleting              DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "DELETING"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleted               DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "DELETED"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusUpdating              DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "UPDATING"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreating              DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "CREATING"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreated               DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "CREATED"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusReadyForConfiguration DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusConfigured            DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "CONFIGURED"
	DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusNeedsAttention        DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum = map[string]DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum{
	"FAILED":                  DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusFailed,
	"DELETING":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleting,
	"DELETED":                 DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleted,
	"UPDATING":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusUpdating,
	"CREATING":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreating,
	"CREATED":                 DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusReadyForConfiguration,
	"CONFIGURED":              DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusConfigured,
	"NEEDS_ATTENTION":         DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusNeedsAttention,
}

var mappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumLowerCase = map[string]DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum{
	"failed":                  DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusFailed,
	"deleting":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleting,
	"deleted":                 DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusDeleted,
	"updating":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusUpdating,
	"creating":                DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreating,
	"created":                 DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusCreated,
	"ready_for_configuration": DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusReadyForConfiguration,
	"configured":              DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusConfigured,
	"needs_attention":         DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusNeedsAttention,
}

// GetDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumValues Enumerates the set of values for DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum
func GetDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumValues() []DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum {
	values := make([]DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum, 0)
	for _, v := range mappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumStringValues Enumerates the set of values in String for DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum
func GetDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumStringValues() []string {
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

// GetMappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum(val string) (DistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnum, bool) {
	enum, ok := mappingDistributedAutonomousDatabaseCatalogWithDedicatedInfraStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
