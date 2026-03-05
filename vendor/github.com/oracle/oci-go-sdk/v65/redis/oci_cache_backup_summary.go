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

// OciCacheBackupSummary Summary of the OCI Cache Backup.
type OciCacheBackupSummary struct {

	// Unique identifier that is immutable on creation
	Id *string `mandatory:"true" json:"id"`

	// Backup identifier, can be renamed
	DisplayName *string `mandatory:"true" json:"displayName"`

	// Compartment identifier
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The source OCI Cache Cluster OCID.
	SourceClusterId *string `mandatory:"true" json:"sourceClusterId"`

	// The date and time the backup was created. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the backup.
	LifecycleState OciCacheBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Specifies whether the cluster is sharded or non-sharded.
	ClusterMode RedisClusterClusterModeEnum `mandatory:"true" json:"clusterMode"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	// The date and time the backup was updated. An RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339) formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Specifies whether the backup was created from a replica or primary node
	BackupSource OciCacheBackupBackupSourceEnum `mandatory:"false" json:"backupSource,omitempty"`

	// Backup retention period in days.
	RetentionPeriodInDays *int `mandatory:"false" json:"retentionPeriodInDays"`

	// Backup size in GB.
	BackupSizeInGBs *float32 `mandatory:"false" json:"backupSizeInGBs"`

	// Backup Type.
	BackupType OciCacheBackupSummaryBackupTypeEnum `mandatory:"false" json:"backupType,omitempty"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OciCacheBackupSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciCacheBackupSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOciCacheBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOciCacheBackupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRedisClusterClusterModeEnum(string(m.ClusterMode)); !ok && m.ClusterMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ClusterMode: %s. Supported values are: %s.", m.ClusterMode, strings.Join(GetRedisClusterClusterModeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingOciCacheBackupBackupSourceEnum(string(m.BackupSource)); !ok && m.BackupSource != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupSource: %s. Supported values are: %s.", m.BackupSource, strings.Join(GetOciCacheBackupBackupSourceEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOciCacheBackupSummaryBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetOciCacheBackupSummaryBackupTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OciCacheBackupSummaryBackupTypeEnum Enum with underlying type: string
type OciCacheBackupSummaryBackupTypeEnum string

// Set of constants representing the allowable values for OciCacheBackupSummaryBackupTypeEnum
const (
	OciCacheBackupSummaryBackupTypeManual    OciCacheBackupSummaryBackupTypeEnum = "MANUAL"
	OciCacheBackupSummaryBackupTypeAutomated OciCacheBackupSummaryBackupTypeEnum = "AUTOMATED"
)

var mappingOciCacheBackupSummaryBackupTypeEnum = map[string]OciCacheBackupSummaryBackupTypeEnum{
	"MANUAL":    OciCacheBackupSummaryBackupTypeManual,
	"AUTOMATED": OciCacheBackupSummaryBackupTypeAutomated,
}

var mappingOciCacheBackupSummaryBackupTypeEnumLowerCase = map[string]OciCacheBackupSummaryBackupTypeEnum{
	"manual":    OciCacheBackupSummaryBackupTypeManual,
	"automated": OciCacheBackupSummaryBackupTypeAutomated,
}

// GetOciCacheBackupSummaryBackupTypeEnumValues Enumerates the set of values for OciCacheBackupSummaryBackupTypeEnum
func GetOciCacheBackupSummaryBackupTypeEnumValues() []OciCacheBackupSummaryBackupTypeEnum {
	values := make([]OciCacheBackupSummaryBackupTypeEnum, 0)
	for _, v := range mappingOciCacheBackupSummaryBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOciCacheBackupSummaryBackupTypeEnumStringValues Enumerates the set of values in String for OciCacheBackupSummaryBackupTypeEnum
func GetOciCacheBackupSummaryBackupTypeEnumStringValues() []string {
	return []string{
		"MANUAL",
		"AUTOMATED",
	}
}

// GetMappingOciCacheBackupSummaryBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOciCacheBackupSummaryBackupTypeEnum(val string) (OciCacheBackupSummaryBackupTypeEnum, bool) {
	enum, ok := mappingOciCacheBackupSummaryBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
