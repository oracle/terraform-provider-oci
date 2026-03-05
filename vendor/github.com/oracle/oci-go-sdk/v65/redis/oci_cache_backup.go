// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Cache API
//
// Use the OCI Cache API to create and manage clusters. A cluster is a memory-based storage solution. For more information, see OCI Cache (https://docs.oracle.com/iaas/Content/ocicache/home.htm).
//

package redis

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OciCacheBackup OCI Cache cluster backup information
type OciCacheBackup struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Backup display name
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Backup compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current state of the backup.
	LifecycleState OciCacheBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Backup size in GB.
	BackupSizeInGBs *float32 `mandatory:"true" json:"backupSizeInGBs"`

	// The source OCI Cache Cluster OCID.
	SourceClusterId *string `mandatory:"true" json:"sourceClusterId"`

	// Specifies whether the cluster is sharded or non-sharded.
	ClusterMode RedisClusterClusterModeEnum `mandatory:"true" json:"clusterMode"`

	// The OCI Cache engine version that the cluster is running.
	SoftwareVersion RedisClusterSoftwareVersionEnum `mandatory:"true" json:"softwareVersion"`

	// Backup description
	Description *string `mandatory:"false" json:"description"`

	// Specifies whether the backup was created from a replica or primary node
	BackupSource OciCacheBackupBackupSourceEnum `mandatory:"false" json:"backupSource,omitempty"`

	// Backup retention period in days.
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// The amount of memory allocated to the cluster, in gigabytes.
	ClusterMemoryInGBs *float32 `mandatory:"false" json:"clusterMemoryInGBs"`

	// The number of shards in a sharded cluster. Only applicable when clusterMode is SHARDED.
	ShardCount *int `mandatory:"false" json:"shardCount"`

	// The date and time the backup was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the backup was updated. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Backup Type.
	BackupType OciCacheBackupBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OciCacheBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisClusterClusterModeEnum(string(m.ClusterMode)); !ok && m.ClusterMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterMode: %s. Supported values are: %s.", m.ClusterMode, strings.Join(GetRedisClusterClusterModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisClusterSoftwareVersionEnum(string(m.SoftwareVersion)); !ok && m.SoftwareVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SoftwareVersion: %s. Supported values are: %s.", m.SoftwareVersion, strings.Join(GetRedisClusterSoftwareVersionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOciCacheBackupBackupSourceEnum(string(m.BackupSource)); !ok && m.BackupSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupSource: %s. Supported values are: %s.", m.BackupSource, strings.Join(GetOciCacheBackupBackupSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetOciCacheBackupBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciCacheBackupBackupSourceEnum Enum with underlying type: string
type OciCacheBackupBackupSourceEnum string

// Set of constants representing the allowable values for OciCacheBackupBackupSourceEnum
const (
	OciCacheBackupBackupSourceReplica OciCacheBackupBackupSourceEnum = "REPLICA"
	OciCacheBackupBackupSourcePrimary OciCacheBackupBackupSourceEnum = "PRIMARY"
)

var mappingOciCacheBackupBackupSourceEnum = map[string]OciCacheBackupBackupSourceEnum{
	"REPLICA": OciCacheBackupBackupSourceReplica,
	"PRIMARY": OciCacheBackupBackupSourcePrimary,
}

var mappingOciCacheBackupBackupSourceEnumLowerCase = map[string]OciCacheBackupBackupSourceEnum{
	"replica": OciCacheBackupBackupSourceReplica,
	"primary": OciCacheBackupBackupSourcePrimary,
}

// GetOciCacheBackupBackupSourceEnumValues Enumerates the set of values for OciCacheBackupBackupSourceEnum
func GetOciCacheBackupBackupSourceEnumValues() []OciCacheBackupBackupSourceEnum {
	values := make([]OciCacheBackupBackupSourceEnum, 0)
	for _, v := range mappingOciCacheBackupBackupSourceEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheBackupBackupSourceEnumStringValues Enumerates the set of values in String for OciCacheBackupBackupSourceEnum
func GetOciCacheBackupBackupSourceEnumStringValues() []string {
	return []string{
		"REPLICA",
		"PRIMARY",
	}
}

// GetMappingOciCacheBackupBackupSourceEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheBackupBackupSourceEnum(val string) (OciCacheBackupBackupSourceEnum, bool) {
	enum, ok := mappingOciCacheBackupBackupSourceEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OciCacheBackupLifecycleStateEnum Enum with underlying type: string
type OciCacheBackupLifecycleStateEnum string

// Set of constants representing the allowable values for OciCacheBackupLifecycleStateEnum
const (
	OciCacheBackupLifecycleStateCreating OciCacheBackupLifecycleStateEnum = "CREATING"
	OciCacheBackupLifecycleStateUpdating OciCacheBackupLifecycleStateEnum = "UPDATING"
	OciCacheBackupLifecycleStateActive   OciCacheBackupLifecycleStateEnum = "ACTIVE"
	OciCacheBackupLifecycleStateDeleting OciCacheBackupLifecycleStateEnum = "DELETING"
	OciCacheBackupLifecycleStateDeleted  OciCacheBackupLifecycleStateEnum = "DELETED"
	OciCacheBackupLifecycleStateFailed   OciCacheBackupLifecycleStateEnum = "FAILED"
)

var mappingOciCacheBackupLifecycleStateEnum = map[string]OciCacheBackupLifecycleStateEnum{
	"CREATING": OciCacheBackupLifecycleStateCreating,
	"UPDATING": OciCacheBackupLifecycleStateUpdating,
	"ACTIVE":   OciCacheBackupLifecycleStateActive,
	"DELETING": OciCacheBackupLifecycleStateDeleting,
	"DELETED":  OciCacheBackupLifecycleStateDeleted,
	"FAILED":   OciCacheBackupLifecycleStateFailed,
}

var mappingOciCacheBackupLifecycleStateEnumLowerCase = map[string]OciCacheBackupLifecycleStateEnum{
	"creating": OciCacheBackupLifecycleStateCreating,
	"updating": OciCacheBackupLifecycleStateUpdating,
	"active":   OciCacheBackupLifecycleStateActive,
	"deleting": OciCacheBackupLifecycleStateDeleting,
	"deleted":  OciCacheBackupLifecycleStateDeleted,
	"failed":   OciCacheBackupLifecycleStateFailed,
}

// GetOciCacheBackupLifecycleStateEnumValues Enumerates the set of values for OciCacheBackupLifecycleStateEnum
func GetOciCacheBackupLifecycleStateEnumValues() []OciCacheBackupLifecycleStateEnum {
	values := make([]OciCacheBackupLifecycleStateEnum, 0)
	for _, v := range mappingOciCacheBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheBackupLifecycleStateEnumStringValues Enumerates the set of values in String for OciCacheBackupLifecycleStateEnum
func GetOciCacheBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOciCacheBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheBackupLifecycleStateEnum(val string) (OciCacheBackupLifecycleStateEnum, bool) {
	enum, ok := mappingOciCacheBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OciCacheBackupBackupTypeEnum Enum with underlying type: string
type OciCacheBackupBackupTypeEnum string

// Set of constants representing the allowable values for OciCacheBackupBackupTypeEnum
const (
	OciCacheBackupBackupTypeManual    OciCacheBackupBackupTypeEnum = "MANUAL"
	OciCacheBackupBackupTypeAutomated OciCacheBackupBackupTypeEnum = "AUTOMATED"
)

var mappingOciCacheBackupBackupTypeEnum = map[string]OciCacheBackupBackupTypeEnum{
	"MANUAL":    OciCacheBackupBackupTypeManual,
	"AUTOMATED": OciCacheBackupBackupTypeAutomated,
}

var mappingOciCacheBackupBackupTypeEnumLowerCase = map[string]OciCacheBackupBackupTypeEnum{
	"manual":    OciCacheBackupBackupTypeManual,
	"automated": OciCacheBackupBackupTypeAutomated,
}

// GetOciCacheBackupBackupTypeEnumValues Enumerates the set of values for OciCacheBackupBackupTypeEnum
func GetOciCacheBackupBackupTypeEnumValues() []OciCacheBackupBackupTypeEnum {
	values := make([]OciCacheBackupBackupTypeEnum, 0)
	for _, v := range mappingOciCacheBackupBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheBackupBackupTypeEnumStringValues Enumerates the set of values in String for OciCacheBackupBackupTypeEnum
func GetOciCacheBackupBackupTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATED",
	}
}

// GetMappingOciCacheBackupBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheBackupBackupTypeEnum(val string) (OciCacheBackupBackupTypeEnum, bool) {
	enum, ok := mappingOciCacheBackupBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
