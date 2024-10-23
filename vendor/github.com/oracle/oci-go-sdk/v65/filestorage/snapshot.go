// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.cloud.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Snapshot A point-in-time snapshot of a specified file system.
type Snapshot struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system from which the snapshot
	// was created.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the snapshot.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the snapshot.
	LifecycleState SnapshotLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the snapshot. This value is immutable.
	// Avoid entering confidential information.
	// Example: `Sunday`
	Name *string `mandatory:"true" json:"name"`

	// The date and time the snapshot was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Specifies generation type of the snapshot.
	SnapshotType SnapshotSnapshotTypeEnum `mandatory:"false" json:"snapshotType,omitempty"`

	// The date and time the snapshot was taken, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// This value might be the same or different from `timeCreated` depending
	// on the following factors:
	// - If the snapshot is created in the original file system directory.
	// - If the snapshot is cloned from a file system.
	// - If the snapshot is replicated from a file system.
	// Example: `2020-08-25T21:10:29.600Z`
	SnapshotTime *common.SDKTime `mandatory:"false" json:"snapshotTime"`

	// An OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) identifying the parent from which this snapshot was cloned.
	// If this snapshot was not cloned, then the `provenanceId` is the same as the snapshot `id` value.
	// If this snapshot was cloned, then the `provenanceId` value is the parent's `provenanceId`.
	// See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	ProvenanceId *string `mandatory:"false" json:"provenanceId"`

	// Specifies whether the snapshot has been cloned.
	// See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	IsCloneSource *bool `mandatory:"false" json:"isCloneSource"`

	// Additional information about the current `lifecycleState`.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The time when this snapshot will be deleted.
	ExpirationTime *common.SDKTime `mandatory:"false" json:"expirationTime"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system snapshot policy that created this snapshot.
	FilesystemSnapshotPolicyId *string `mandatory:"false" json:"filesystemSnapshotPolicyId"`
}

func (m Snapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Snapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSnapshotLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSnapshotLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSnapshotSnapshotTypeEnum(string(m.SnapshotType)); !ok && m.SnapshotType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SnapshotType: %s. Supported values are: %s.", m.SnapshotType, strings.Join(GetSnapshotSnapshotTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SnapshotLifecycleStateEnum Enum with underlying type: string
type SnapshotLifecycleStateEnum string

// Set of constants representing the allowable values for SnapshotLifecycleStateEnum
const (
	SnapshotLifecycleStateCreating SnapshotLifecycleStateEnum = "CREATING"
	SnapshotLifecycleStateActive   SnapshotLifecycleStateEnum = "ACTIVE"
	SnapshotLifecycleStateDeleting SnapshotLifecycleStateEnum = "DELETING"
	SnapshotLifecycleStateDeleted  SnapshotLifecycleStateEnum = "DELETED"
)

var mappingSnapshotLifecycleStateEnum = map[string]SnapshotLifecycleStateEnum{
	"CREATING": SnapshotLifecycleStateCreating,
	"ACTIVE":   SnapshotLifecycleStateActive,
	"DELETING": SnapshotLifecycleStateDeleting,
	"DELETED":  SnapshotLifecycleStateDeleted,
}

var mappingSnapshotLifecycleStateEnumLowerCase = map[string]SnapshotLifecycleStateEnum{
	"creating": SnapshotLifecycleStateCreating,
	"active":   SnapshotLifecycleStateActive,
	"deleting": SnapshotLifecycleStateDeleting,
	"deleted":  SnapshotLifecycleStateDeleted,
}

// GetSnapshotLifecycleStateEnumValues Enumerates the set of values for SnapshotLifecycleStateEnum
func GetSnapshotLifecycleStateEnumValues() []SnapshotLifecycleStateEnum {
	values := make([]SnapshotLifecycleStateEnum, 0)
	for _, v := range mappingSnapshotLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotLifecycleStateEnumStringValues Enumerates the set of values in String for SnapshotLifecycleStateEnum
func GetSnapshotLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSnapshotLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotLifecycleStateEnum(val string) (SnapshotLifecycleStateEnum, bool) {
	enum, ok := mappingSnapshotLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnapshotSnapshotTypeEnum Enum with underlying type: string
type SnapshotSnapshotTypeEnum string

// Set of constants representing the allowable values for SnapshotSnapshotTypeEnum
const (
	SnapshotSnapshotTypeUser        SnapshotSnapshotTypeEnum = "USER"
	SnapshotSnapshotTypePolicyBased SnapshotSnapshotTypeEnum = "POLICY_BASED"
	SnapshotSnapshotTypeReplication SnapshotSnapshotTypeEnum = "REPLICATION"
)

var mappingSnapshotSnapshotTypeEnum = map[string]SnapshotSnapshotTypeEnum{
	"USER":         SnapshotSnapshotTypeUser,
	"POLICY_BASED": SnapshotSnapshotTypePolicyBased,
	"REPLICATION":  SnapshotSnapshotTypeReplication,
}

var mappingSnapshotSnapshotTypeEnumLowerCase = map[string]SnapshotSnapshotTypeEnum{
	"user":         SnapshotSnapshotTypeUser,
	"policy_based": SnapshotSnapshotTypePolicyBased,
	"replication":  SnapshotSnapshotTypeReplication,
}

// GetSnapshotSnapshotTypeEnumValues Enumerates the set of values for SnapshotSnapshotTypeEnum
func GetSnapshotSnapshotTypeEnumValues() []SnapshotSnapshotTypeEnum {
	values := make([]SnapshotSnapshotTypeEnum, 0)
	for _, v := range mappingSnapshotSnapshotTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotSnapshotTypeEnumStringValues Enumerates the set of values in String for SnapshotSnapshotTypeEnum
func GetSnapshotSnapshotTypeEnumStringValues() []string {
	return []string{
		"USER",
		"POLICY_BASED",
		"REPLICATION",
	}
}

// GetMappingSnapshotSnapshotTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotSnapshotTypeEnum(val string) (SnapshotSnapshotTypeEnum, bool) {
	enum, ok := mappingSnapshotSnapshotTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
