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

// FileSystem An NFS file system. To allow access to a file system, add it
// to an export set and associate the export set with a mount
// target. The same file system can be in multiple export sets and
// associated with multiple mount targets.
// To use any of the API operations, you must be authorized in an
// IAM policy. If you're not authorized, talk to an
// administrator. If you're an administrator who needs to write
// policies to give users access, see Getting Started with
// Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
type FileSystem struct {

	// The number of bytes consumed by the file system, including
	// any snapshots. This number reflects the metered size of the file
	// system and is updated asynchronously with respect to
	// updates to the file system.
	// For more information, see File System Usage and Metering (https://docs.cloud.oracle.com/Content/File/Concepts/FSutilization.htm).
	MeteredBytes *int64 `mandatory:"true" json:"meteredBytes"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My file system`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the file system.
	LifecycleState FileSystemLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the file system was created, expressed in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the KMS key which is the master encryption key for the file system.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	SourceDetails *SourceDetails `mandatory:"false" json:"sourceDetails"`

	// Specifies whether the file system has been cloned.
	// See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	IsCloneParent *bool `mandatory:"false" json:"isCloneParent"`

	// Specifies whether the data has finished copying from the source to the clone.
	// Hydration can take up to several hours to complete depending on the size of the source.
	// The source and clone remain available during hydration, but there may be some performance impact.
	// See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm#hydration).
	IsHydrated *bool `mandatory:"false" json:"isHydrated"`

	// Specifies the total number of children of a file system.
	CloneCount *int `mandatory:"false" json:"cloneCount"`

	// Specifies whether the file system is attached to its parent file system.
	CloneAttachStatus FileSystemCloneAttachStatusEnum `mandatory:"false" json:"cloneAttachStatus,omitempty"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Specifies whether the file system can be used as a target file system for replication. The system sets this value to `true` if the file system is unexported, hasn't yet been specified as a target file system in any replication resource, and has no user snapshots. After the file system has been specified as a target in a replication, or if the file system contains user snapshots, the system sets this value to `false`.
	// For more information, see Using Replication (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/using-replication.htm).
	IsTargetable *bool `mandatory:"false" json:"isTargetable"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication target associated with the file system.
	// Empty if the file system is not being used as target in a replication.
	ReplicationTargetId *string `mandatory:"false" json:"replicationTargetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated file system snapshot policy, which
	// controls the frequency of snapshot creation and retention period of the taken snapshots.
	FilesystemSnapshotPolicyId *string `mandatory:"false" json:"filesystemSnapshotPolicyId"`
}

func (m FileSystem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFileSystemLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFileSystemLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFileSystemCloneAttachStatusEnum(string(m.CloneAttachStatus)); !ok && m.CloneAttachStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneAttachStatus: %s. Supported values are: %s.", m.CloneAttachStatus, strings.Join(GetFileSystemCloneAttachStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSystemLifecycleStateEnum Enum with underlying type: string
type FileSystemLifecycleStateEnum string

// Set of constants representing the allowable values for FileSystemLifecycleStateEnum
const (
	FileSystemLifecycleStateCreating FileSystemLifecycleStateEnum = "CREATING"
	FileSystemLifecycleStateActive   FileSystemLifecycleStateEnum = "ACTIVE"
	FileSystemLifecycleStateUpdating FileSystemLifecycleStateEnum = "UPDATING"
	FileSystemLifecycleStateDeleting FileSystemLifecycleStateEnum = "DELETING"
	FileSystemLifecycleStateDeleted  FileSystemLifecycleStateEnum = "DELETED"
	FileSystemLifecycleStateFailed   FileSystemLifecycleStateEnum = "FAILED"
)

var mappingFileSystemLifecycleStateEnum = map[string]FileSystemLifecycleStateEnum{
	"CREATING": FileSystemLifecycleStateCreating,
	"ACTIVE":   FileSystemLifecycleStateActive,
	"UPDATING": FileSystemLifecycleStateUpdating,
	"DELETING": FileSystemLifecycleStateDeleting,
	"DELETED":  FileSystemLifecycleStateDeleted,
	"FAILED":   FileSystemLifecycleStateFailed,
}

var mappingFileSystemLifecycleStateEnumLowerCase = map[string]FileSystemLifecycleStateEnum{
	"creating": FileSystemLifecycleStateCreating,
	"active":   FileSystemLifecycleStateActive,
	"updating": FileSystemLifecycleStateUpdating,
	"deleting": FileSystemLifecycleStateDeleting,
	"deleted":  FileSystemLifecycleStateDeleted,
	"failed":   FileSystemLifecycleStateFailed,
}

// GetFileSystemLifecycleStateEnumValues Enumerates the set of values for FileSystemLifecycleStateEnum
func GetFileSystemLifecycleStateEnumValues() []FileSystemLifecycleStateEnum {
	values := make([]FileSystemLifecycleStateEnum, 0)
	for _, v := range mappingFileSystemLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemLifecycleStateEnumStringValues Enumerates the set of values in String for FileSystemLifecycleStateEnum
func GetFileSystemLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFileSystemLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemLifecycleStateEnum(val string) (FileSystemLifecycleStateEnum, bool) {
	enum, ok := mappingFileSystemLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FileSystemCloneAttachStatusEnum Enum with underlying type: string
type FileSystemCloneAttachStatusEnum string

// Set of constants representing the allowable values for FileSystemCloneAttachStatusEnum
const (
	FileSystemCloneAttachStatusAttached  FileSystemCloneAttachStatusEnum = "ATTACHED"
	FileSystemCloneAttachStatusDetaching FileSystemCloneAttachStatusEnum = "DETACHING"
	FileSystemCloneAttachStatusDetached  FileSystemCloneAttachStatusEnum = "DETACHED"
)

var mappingFileSystemCloneAttachStatusEnum = map[string]FileSystemCloneAttachStatusEnum{
	"ATTACHED":  FileSystemCloneAttachStatusAttached,
	"DETACHING": FileSystemCloneAttachStatusDetaching,
	"DETACHED":  FileSystemCloneAttachStatusDetached,
}

var mappingFileSystemCloneAttachStatusEnumLowerCase = map[string]FileSystemCloneAttachStatusEnum{
	"attached":  FileSystemCloneAttachStatusAttached,
	"detaching": FileSystemCloneAttachStatusDetaching,
	"detached":  FileSystemCloneAttachStatusDetached,
}

// GetFileSystemCloneAttachStatusEnumValues Enumerates the set of values for FileSystemCloneAttachStatusEnum
func GetFileSystemCloneAttachStatusEnumValues() []FileSystemCloneAttachStatusEnum {
	values := make([]FileSystemCloneAttachStatusEnum, 0)
	for _, v := range mappingFileSystemCloneAttachStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemCloneAttachStatusEnumStringValues Enumerates the set of values in String for FileSystemCloneAttachStatusEnum
func GetFileSystemCloneAttachStatusEnumStringValues() []string {
	return []string{
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

// GetMappingFileSystemCloneAttachStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemCloneAttachStatusEnum(val string) (FileSystemCloneAttachStatusEnum, bool) {
	enum, ok := mappingFileSystemCloneAttachStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
