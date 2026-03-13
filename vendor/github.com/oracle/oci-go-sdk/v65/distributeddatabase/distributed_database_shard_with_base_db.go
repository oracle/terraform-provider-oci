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

// DistributedDatabaseShardWithBaseDb Globally distributed database shard with base database.
type DistributedDatabaseShardWithBaseDb struct {

	// Name of the shard.
	Name *string `mandatory:"true" json:"name"`

	// The time the shard was created. An RFC3339 formatted datetime string
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the shard was last updated. An RFC3339 formatted datetime string
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// The name of the shardGroup for the shard.
	ShardGroup *string `mandatory:"true" json:"shardGroup"`

	// The shard space name for the Globally distributed database. Shard space for existing shard cannot be changed, once shard is created.
	// Shard space name shall be used while creation of new shards.
	ShardSpace *string `mandatory:"true" json:"shardSpace"`

	// The number of CPU cores to enable.
	CpuCoreCount *int `mandatory:"true" json:"cpuCoreCount"`

	// The data storage size to be allocated in GBs.
	DataStorageSizeInGbs *int `mandatory:"true" json:"dataStorageSizeInGbs"`

	// The name of the availability domain that the base database infrastructure resource is located in.
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The identifier of the subnet for the Dbsystem instance.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The shape of the DBSystem instance. The shape determines the amount of CPU, storage, and memory resources allocated to the instance.
	Shape *string `mandatory:"true" json:"shape"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Home.
	DbSystemId *string `mandatory:"false" json:"dbSystemId"`

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

	// Peer details for the shard.
	PeerDetails []ShardPeerWithBaseDb `mandatory:"false" json:"peerDetails"`

	// Fault Domain in which this base database is provisioned.
	FaultDomain *string `mandatory:"false" json:"faultDomain"`

	DataCollectionOptions *BaseDbDistributedDatabaseDataCollectionOptions `mandatory:"false" json:"dataCollectionOptions"`

	Metadata *DistributedDbMetadata `mandatory:"false" json:"metadata"`

	// Status of Base database based shard for the Globally distributed database.
	Status DistributedDatabaseShardWithBaseDbStatusEnum `mandatory:"true" json:"status"`

	// The Oracle Database edition to use for creating a Base database based shard.
	DbEdition BaseDbEditionEnum `mandatory:"true" json:"dbEdition"`

	// The type of redundancy configured for the Base database based shard.
	DiskRedundancy BaseDbDiskRedundancyEnum `mandatory:"false" json:"diskRedundancy,omitempty"`

	// The Oracle license model which applies to the Base database based shard.
	LicenseModel BaseDbLicenseModelEnum `mandatory:"true" json:"licenseModel"`

	// The storage option used for Base database based shard.
	StorageManagement BaseDbStorageManagementEnum `mandatory:"false" json:"storageManagement,omitempty"`

	// Block Volume Performance mode for Base database based shard.
	StorageVolumePerformanceMode BaseDbStorageVolumePerformanceModeEnum `mandatory:"false" json:"storageVolumePerformanceMode,omitempty"`
}

// GetName returns Name
func (m DistributedDatabaseShardWithBaseDb) GetName() *string {
	return m.Name
}

// GetTimeCreated returns TimeCreated
func (m DistributedDatabaseShardWithBaseDb) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m DistributedDatabaseShardWithBaseDb) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

func (m DistributedDatabaseShardWithBaseDb) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DistributedDatabaseShardWithBaseDb) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDistributedDatabaseShardWithBaseDbStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetDistributedDatabaseShardWithBaseDbStatusEnumStringValues(), ",")))
	}

	if _, ok := GetMappingBaseDbEditionEnum(string(m.DbEdition)); !ok && m.DbEdition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DbEdition: %s. Supported values are: %s.", m.DbEdition, strings.Join(GetBaseDbEditionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbDiskRedundancyEnum(string(m.DiskRedundancy)); !ok && m.DiskRedundancy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DiskRedundancy: %s. Supported values are: %s.", m.DiskRedundancy, strings.Join(GetBaseDbDiskRedundancyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbLicenseModelEnum(string(m.LicenseModel)); !ok && m.LicenseModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LicenseModel: %s. Supported values are: %s.", m.LicenseModel, strings.Join(GetBaseDbLicenseModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbStorageManagementEnum(string(m.StorageManagement)); !ok && m.StorageManagement != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageManagement: %s. Supported values are: %s.", m.StorageManagement, strings.Join(GetBaseDbStorageManagementEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBaseDbStorageVolumePerformanceModeEnum(string(m.StorageVolumePerformanceMode)); !ok && m.StorageVolumePerformanceMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageVolumePerformanceMode: %s. Supported values are: %s.", m.StorageVolumePerformanceMode, strings.Join(GetBaseDbStorageVolumePerformanceModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m DistributedDatabaseShardWithBaseDb) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeDistributedDatabaseShardWithBaseDb DistributedDatabaseShardWithBaseDb
	s := struct {
		DiscriminatorParam string `json:"source"`
		MarshalTypeDistributedDatabaseShardWithBaseDb
	}{
		"BASE_DB",
		(MarshalTypeDistributedDatabaseShardWithBaseDb)(m),
	}

	return json.Marshal(&s)
}

// DistributedDatabaseShardWithBaseDbStatusEnum Enum with underlying type: string
type DistributedDatabaseShardWithBaseDbStatusEnum string

// Set of constants representing the allowable values for DistributedDatabaseShardWithBaseDbStatusEnum
const (
	DistributedDatabaseShardWithBaseDbStatusFailed                DistributedDatabaseShardWithBaseDbStatusEnum = "FAILED"
	DistributedDatabaseShardWithBaseDbStatusDeleting              DistributedDatabaseShardWithBaseDbStatusEnum = "DELETING"
	DistributedDatabaseShardWithBaseDbStatusDeleted               DistributedDatabaseShardWithBaseDbStatusEnum = "DELETED"
	DistributedDatabaseShardWithBaseDbStatusUpdating              DistributedDatabaseShardWithBaseDbStatusEnum = "UPDATING"
	DistributedDatabaseShardWithBaseDbStatusCreating              DistributedDatabaseShardWithBaseDbStatusEnum = "CREATING"
	DistributedDatabaseShardWithBaseDbStatusCreated               DistributedDatabaseShardWithBaseDbStatusEnum = "CREATED"
	DistributedDatabaseShardWithBaseDbStatusReadyForConfiguration DistributedDatabaseShardWithBaseDbStatusEnum = "READY_FOR_CONFIGURATION"
	DistributedDatabaseShardWithBaseDbStatusConfigured            DistributedDatabaseShardWithBaseDbStatusEnum = "CONFIGURED"
	DistributedDatabaseShardWithBaseDbStatusNeedsAttention        DistributedDatabaseShardWithBaseDbStatusEnum = "NEEDS_ATTENTION"
)

var mappingDistributedDatabaseShardWithBaseDbStatusEnum = map[string]DistributedDatabaseShardWithBaseDbStatusEnum{
	"FAILED":                  DistributedDatabaseShardWithBaseDbStatusFailed,
	"DELETING":                DistributedDatabaseShardWithBaseDbStatusDeleting,
	"DELETED":                 DistributedDatabaseShardWithBaseDbStatusDeleted,
	"UPDATING":                DistributedDatabaseShardWithBaseDbStatusUpdating,
	"CREATING":                DistributedDatabaseShardWithBaseDbStatusCreating,
	"CREATED":                 DistributedDatabaseShardWithBaseDbStatusCreated,
	"READY_FOR_CONFIGURATION": DistributedDatabaseShardWithBaseDbStatusReadyForConfiguration,
	"CONFIGURED":              DistributedDatabaseShardWithBaseDbStatusConfigured,
	"NEEDS_ATTENTION":         DistributedDatabaseShardWithBaseDbStatusNeedsAttention,
}

var mappingDistributedDatabaseShardWithBaseDbStatusEnumLowerCase = map[string]DistributedDatabaseShardWithBaseDbStatusEnum{
	"failed":                  DistributedDatabaseShardWithBaseDbStatusFailed,
	"deleting":                DistributedDatabaseShardWithBaseDbStatusDeleting,
	"deleted":                 DistributedDatabaseShardWithBaseDbStatusDeleted,
	"updating":                DistributedDatabaseShardWithBaseDbStatusUpdating,
	"creating":                DistributedDatabaseShardWithBaseDbStatusCreating,
	"created":                 DistributedDatabaseShardWithBaseDbStatusCreated,
	"ready_for_configuration": DistributedDatabaseShardWithBaseDbStatusReadyForConfiguration,
	"configured":              DistributedDatabaseShardWithBaseDbStatusConfigured,
	"needs_attention":         DistributedDatabaseShardWithBaseDbStatusNeedsAttention,
}

// GetDistributedDatabaseShardWithBaseDbStatusEnumValues Enumerates the set of values for DistributedDatabaseShardWithBaseDbStatusEnum
func GetDistributedDatabaseShardWithBaseDbStatusEnumValues() []DistributedDatabaseShardWithBaseDbStatusEnum {
	values := make([]DistributedDatabaseShardWithBaseDbStatusEnum, 0)
	for _, v := range mappingDistributedDatabaseShardWithBaseDbStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetDistributedDatabaseShardWithBaseDbStatusEnumStringValues Enumerates the set of values in String for DistributedDatabaseShardWithBaseDbStatusEnum
func GetDistributedDatabaseShardWithBaseDbStatusEnumStringValues() []string {
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

// GetMappingDistributedDatabaseShardWithBaseDbStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDistributedDatabaseShardWithBaseDbStatusEnum(val string) (DistributedDatabaseShardWithBaseDbStatusEnum, bool) {
	enum, ok := mappingDistributedDatabaseShardWithBaseDbStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
