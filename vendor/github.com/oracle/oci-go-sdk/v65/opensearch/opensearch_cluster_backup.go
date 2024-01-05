// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpensearchClusterBackup An OpenSearch cluster backup resource. An cluster is set of instances that provide OpenSearch functionality in OCI Search Service with OpenSearch.
// For more information, see Cluster Backups (https://docs.cloud.oracle.com/iaas/Content/search-opensearch/Concepts/ociopensearchbackups.htm).
type OpensearchClusterBackup struct {

	// The OCID of the cluster backup.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment where the cluster backup is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Specifies whether the cluster backup was created manually, or automatically as a scheduled backup.
	BackupType OpensearchClusterBackupBackupTypeEnum `mandatory:"true" json:"backupType"`

	// The current state of the cluster backup.
	LifecycleState OpensearchClusterBackupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the source OpenSearch cluster for the cluster backup.
	SourceClusterId *string `mandatory:"true" json:"sourceClusterId"`

	// The name of the cluster backup. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The date and time the cluster backup was created. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cluster backup was updated. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Additional information about the current lifecycle state of the cluster backup.
	LifecyleDetails *string `mandatory:"false" json:"lifecyleDetails"`

	// The Object Storage namespace for the cluster backup.
	Namespace *string `mandatory:"false" json:"namespace"`

	// The name of the Object Storage bucket for the cluster backup.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// The prefix within the Object Storage bucket for the cluster backup.
	Prefix *string `mandatory:"false" json:"prefix"`

	// The date and time the cluster backup expires. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeExpired *common.SDKTime `mandatory:"false" json:"timeExpired"`

	// The size in GB of the cluster backup.
	BackupSize *float64 `mandatory:"false" json:"backupSize"`

	// The name of the source OpenSearch cluster for the cluster backup.
	SourceClusterDisplayName *string `mandatory:"false" json:"sourceClusterDisplayName"`

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

func (m OpensearchClusterBackup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchClusterBackup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterBackupBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetOpensearchClusterBackupBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOpensearchClusterBackupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterBackupLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpensearchClusterBackupBackupTypeEnum Enum with underlying type: string
type OpensearchClusterBackupBackupTypeEnum string

// Set of constants representing the allowable values for OpensearchClusterBackupBackupTypeEnum
const (
	OpensearchClusterBackupBackupTypeScheduled OpensearchClusterBackupBackupTypeEnum = "SCHEDULED"
	OpensearchClusterBackupBackupTypeManual    OpensearchClusterBackupBackupTypeEnum = "MANUAL"
)

var mappingOpensearchClusterBackupBackupTypeEnum = map[string]OpensearchClusterBackupBackupTypeEnum{
	"SCHEDULED": OpensearchClusterBackupBackupTypeScheduled,
	"MANUAL":    OpensearchClusterBackupBackupTypeManual,
}

var mappingOpensearchClusterBackupBackupTypeEnumLowerCase = map[string]OpensearchClusterBackupBackupTypeEnum{
	"scheduled": OpensearchClusterBackupBackupTypeScheduled,
	"manual":    OpensearchClusterBackupBackupTypeManual,
}

// GetOpensearchClusterBackupBackupTypeEnumValues Enumerates the set of values for OpensearchClusterBackupBackupTypeEnum
func GetOpensearchClusterBackupBackupTypeEnumValues() []OpensearchClusterBackupBackupTypeEnum {
	values := make([]OpensearchClusterBackupBackupTypeEnum, 0)
	for _, v := range mappingOpensearchClusterBackupBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetOpensearchClusterBackupBackupTypeEnumStringValues Enumerates the set of values in String for OpensearchClusterBackupBackupTypeEnum
func GetOpensearchClusterBackupBackupTypeEnumStringValues() []string {
	return []string{
		"SCHEDULED",
		"MANUAL",
	}
}

// GetMappingOpensearchClusterBackupBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpensearchClusterBackupBackupTypeEnum(val string) (OpensearchClusterBackupBackupTypeEnum, bool) {
	enum, ok := mappingOpensearchClusterBackupBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OpensearchClusterBackupLifecycleStateEnum Enum with underlying type: string
type OpensearchClusterBackupLifecycleStateEnum string

// Set of constants representing the allowable values for OpensearchClusterBackupLifecycleStateEnum
const (
	OpensearchClusterBackupLifecycleStateCreating OpensearchClusterBackupLifecycleStateEnum = "CREATING"
	OpensearchClusterBackupLifecycleStateUpdating OpensearchClusterBackupLifecycleStateEnum = "UPDATING"
	OpensearchClusterBackupLifecycleStateActive   OpensearchClusterBackupLifecycleStateEnum = "ACTIVE"
	OpensearchClusterBackupLifecycleStateDeleting OpensearchClusterBackupLifecycleStateEnum = "DELETING"
	OpensearchClusterBackupLifecycleStateDeleted  OpensearchClusterBackupLifecycleStateEnum = "DELETED"
	OpensearchClusterBackupLifecycleStateFailed   OpensearchClusterBackupLifecycleStateEnum = "FAILED"
)

var mappingOpensearchClusterBackupLifecycleStateEnum = map[string]OpensearchClusterBackupLifecycleStateEnum{
	"CREATING": OpensearchClusterBackupLifecycleStateCreating,
	"UPDATING": OpensearchClusterBackupLifecycleStateUpdating,
	"ACTIVE":   OpensearchClusterBackupLifecycleStateActive,
	"DELETING": OpensearchClusterBackupLifecycleStateDeleting,
	"DELETED":  OpensearchClusterBackupLifecycleStateDeleted,
	"FAILED":   OpensearchClusterBackupLifecycleStateFailed,
}

var mappingOpensearchClusterBackupLifecycleStateEnumLowerCase = map[string]OpensearchClusterBackupLifecycleStateEnum{
	"creating": OpensearchClusterBackupLifecycleStateCreating,
	"updating": OpensearchClusterBackupLifecycleStateUpdating,
	"active":   OpensearchClusterBackupLifecycleStateActive,
	"deleting": OpensearchClusterBackupLifecycleStateDeleting,
	"deleted":  OpensearchClusterBackupLifecycleStateDeleted,
	"failed":   OpensearchClusterBackupLifecycleStateFailed,
}

// GetOpensearchClusterBackupLifecycleStateEnumValues Enumerates the set of values for OpensearchClusterBackupLifecycleStateEnum
func GetOpensearchClusterBackupLifecycleStateEnumValues() []OpensearchClusterBackupLifecycleStateEnum {
	values := make([]OpensearchClusterBackupLifecycleStateEnum, 0)
	for _, v := range mappingOpensearchClusterBackupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpensearchClusterBackupLifecycleStateEnumStringValues Enumerates the set of values in String for OpensearchClusterBackupLifecycleStateEnum
func GetOpensearchClusterBackupLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOpensearchClusterBackupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpensearchClusterBackupLifecycleStateEnum(val string) (OpensearchClusterBackupLifecycleStateEnum, bool) {
	enum, ok := mappingOpensearchClusterBackupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
