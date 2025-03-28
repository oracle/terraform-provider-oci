// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage API
//
// Use the File Storage service API to manage file systems, mount targets, and snapshots.
// For more information, see Overview of File Storage (https://docs.oracle.com/iaas/Content/File/Concepts/filestorageoverview.htm).
//

package filestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FileSystemSummary Summary information for a file system.
type FileSystemSummary struct {

	// The number of bytes consumed by the file system, including
	// any snapshots. This number reflects the metered size of the file
	// system and is updated asynchronously with respect to
	// updates to the file system.
	MeteredBytes *int64 `mandatory:"true" json:"meteredBytes"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My file system`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the file system.
	LifecycleState FileSystemSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the file system was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the file system is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource.
	// System tags are applied to resources by internal OCI services.
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// Specifies the total number of replications for which this file system is a source.
	ReplicationSourceCount *int `mandatory:"false" json:"replicationSourceCount"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	SourceDetails *SourceDetails `mandatory:"false" json:"sourceDetails"`

	// Specifies whether the file system has been cloned.
	// See Cloning a File System (https://docs.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	IsCloneParent *bool `mandatory:"false" json:"isCloneParent"`

	// Specifies whether the data has finished copying from the source to the clone.
	// Hydration can take up to several hours to complete depending on the size of the source.
	// The source and clone remain available during hydration, but there may be some performance impact.
	// See Cloning a File System (https://docs.oracle.com/iaas/Content/File/Tasks/cloningFS.htm#hydration).
	IsHydrated *bool `mandatory:"false" json:"isHydrated"`

	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Specifies whether the file system is attached to its parent file system.
	CloneAttachStatus FileSystemSummaryCloneAttachStatusEnum `mandatory:"false" json:"cloneAttachStatus,omitempty"`

	// Displays the state of enforcement of quota rules on the file system.
	QuotaEnforcementState FileSystemSummaryQuotaEnforcementStateEnum `mandatory:"false" json:"quotaEnforcementState,omitempty"`
}

func (m FileSystemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FileSystemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFileSystemSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFileSystemSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingFileSystemSummaryCloneAttachStatusEnum(string(m.CloneAttachStatus)); !ok && m.CloneAttachStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneAttachStatus: %s. Supported values are: %s.", m.CloneAttachStatus, strings.Join(GetFileSystemSummaryCloneAttachStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFileSystemSummaryQuotaEnforcementStateEnum(string(m.QuotaEnforcementState)); !ok && m.QuotaEnforcementState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for QuotaEnforcementState: %s. Supported values are: %s.", m.QuotaEnforcementState, strings.Join(GetFileSystemSummaryQuotaEnforcementStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FileSystemSummaryLifecycleStateEnum Enum with underlying type: string
type FileSystemSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for FileSystemSummaryLifecycleStateEnum
const (
	FileSystemSummaryLifecycleStateCreating FileSystemSummaryLifecycleStateEnum = "CREATING"
	FileSystemSummaryLifecycleStateActive   FileSystemSummaryLifecycleStateEnum = "ACTIVE"
	FileSystemSummaryLifecycleStateUpdating FileSystemSummaryLifecycleStateEnum = "UPDATING"
	FileSystemSummaryLifecycleStateDeleting FileSystemSummaryLifecycleStateEnum = "DELETING"
	FileSystemSummaryLifecycleStateDeleted  FileSystemSummaryLifecycleStateEnum = "DELETED"
	FileSystemSummaryLifecycleStateFailed   FileSystemSummaryLifecycleStateEnum = "FAILED"
)

var mappingFileSystemSummaryLifecycleStateEnum = map[string]FileSystemSummaryLifecycleStateEnum{
	"CREATING": FileSystemSummaryLifecycleStateCreating,
	"ACTIVE":   FileSystemSummaryLifecycleStateActive,
	"UPDATING": FileSystemSummaryLifecycleStateUpdating,
	"DELETING": FileSystemSummaryLifecycleStateDeleting,
	"DELETED":  FileSystemSummaryLifecycleStateDeleted,
	"FAILED":   FileSystemSummaryLifecycleStateFailed,
}

var mappingFileSystemSummaryLifecycleStateEnumLowerCase = map[string]FileSystemSummaryLifecycleStateEnum{
	"creating": FileSystemSummaryLifecycleStateCreating,
	"active":   FileSystemSummaryLifecycleStateActive,
	"updating": FileSystemSummaryLifecycleStateUpdating,
	"deleting": FileSystemSummaryLifecycleStateDeleting,
	"deleted":  FileSystemSummaryLifecycleStateDeleted,
	"failed":   FileSystemSummaryLifecycleStateFailed,
}

// GetFileSystemSummaryLifecycleStateEnumValues Enumerates the set of values for FileSystemSummaryLifecycleStateEnum
func GetFileSystemSummaryLifecycleStateEnumValues() []FileSystemSummaryLifecycleStateEnum {
	values := make([]FileSystemSummaryLifecycleStateEnum, 0)
	for _, v := range mappingFileSystemSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for FileSystemSummaryLifecycleStateEnum
func GetFileSystemSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingFileSystemSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemSummaryLifecycleStateEnum(val string) (FileSystemSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingFileSystemSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FileSystemSummaryCloneAttachStatusEnum Enum with underlying type: string
type FileSystemSummaryCloneAttachStatusEnum string

// Set of constants representing the allowable values for FileSystemSummaryCloneAttachStatusEnum
const (
	FileSystemSummaryCloneAttachStatusAttached  FileSystemSummaryCloneAttachStatusEnum = "ATTACHED"
	FileSystemSummaryCloneAttachStatusDetaching FileSystemSummaryCloneAttachStatusEnum = "DETACHING"
	FileSystemSummaryCloneAttachStatusDetached  FileSystemSummaryCloneAttachStatusEnum = "DETACHED"
)

var mappingFileSystemSummaryCloneAttachStatusEnum = map[string]FileSystemSummaryCloneAttachStatusEnum{
	"ATTACHED":  FileSystemSummaryCloneAttachStatusAttached,
	"DETACHING": FileSystemSummaryCloneAttachStatusDetaching,
	"DETACHED":  FileSystemSummaryCloneAttachStatusDetached,
}

var mappingFileSystemSummaryCloneAttachStatusEnumLowerCase = map[string]FileSystemSummaryCloneAttachStatusEnum{
	"attached":  FileSystemSummaryCloneAttachStatusAttached,
	"detaching": FileSystemSummaryCloneAttachStatusDetaching,
	"detached":  FileSystemSummaryCloneAttachStatusDetached,
}

// GetFileSystemSummaryCloneAttachStatusEnumValues Enumerates the set of values for FileSystemSummaryCloneAttachStatusEnum
func GetFileSystemSummaryCloneAttachStatusEnumValues() []FileSystemSummaryCloneAttachStatusEnum {
	values := make([]FileSystemSummaryCloneAttachStatusEnum, 0)
	for _, v := range mappingFileSystemSummaryCloneAttachStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemSummaryCloneAttachStatusEnumStringValues Enumerates the set of values in String for FileSystemSummaryCloneAttachStatusEnum
func GetFileSystemSummaryCloneAttachStatusEnumStringValues() []string {
	return []string{
		"ATTACHED",
		"DETACHING",
		"DETACHED",
	}
}

// GetMappingFileSystemSummaryCloneAttachStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemSummaryCloneAttachStatusEnum(val string) (FileSystemSummaryCloneAttachStatusEnum, bool) {
	enum, ok := mappingFileSystemSummaryCloneAttachStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FileSystemSummaryQuotaEnforcementStateEnum Enum with underlying type: string
type FileSystemSummaryQuotaEnforcementStateEnum string

// Set of constants representing the allowable values for FileSystemSummaryQuotaEnforcementStateEnum
const (
	FileSystemSummaryQuotaEnforcementStateEnabling  FileSystemSummaryQuotaEnforcementStateEnum = "ENABLING"
	FileSystemSummaryQuotaEnforcementStateEnabled   FileSystemSummaryQuotaEnforcementStateEnum = "ENABLED"
	FileSystemSummaryQuotaEnforcementStateDisabling FileSystemSummaryQuotaEnforcementStateEnum = "DISABLING"
	FileSystemSummaryQuotaEnforcementStateDisabled  FileSystemSummaryQuotaEnforcementStateEnum = "DISABLED"
	FileSystemSummaryQuotaEnforcementStateSyncing   FileSystemSummaryQuotaEnforcementStateEnum = "SYNCING"
	FileSystemSummaryQuotaEnforcementStateFailed    FileSystemSummaryQuotaEnforcementStateEnum = "FAILED"
)

var mappingFileSystemSummaryQuotaEnforcementStateEnum = map[string]FileSystemSummaryQuotaEnforcementStateEnum{
	"ENABLING":  FileSystemSummaryQuotaEnforcementStateEnabling,
	"ENABLED":   FileSystemSummaryQuotaEnforcementStateEnabled,
	"DISABLING": FileSystemSummaryQuotaEnforcementStateDisabling,
	"DISABLED":  FileSystemSummaryQuotaEnforcementStateDisabled,
	"SYNCING":   FileSystemSummaryQuotaEnforcementStateSyncing,
	"FAILED":    FileSystemSummaryQuotaEnforcementStateFailed,
}

var mappingFileSystemSummaryQuotaEnforcementStateEnumLowerCase = map[string]FileSystemSummaryQuotaEnforcementStateEnum{
	"enabling":  FileSystemSummaryQuotaEnforcementStateEnabling,
	"enabled":   FileSystemSummaryQuotaEnforcementStateEnabled,
	"disabling": FileSystemSummaryQuotaEnforcementStateDisabling,
	"disabled":  FileSystemSummaryQuotaEnforcementStateDisabled,
	"syncing":   FileSystemSummaryQuotaEnforcementStateSyncing,
	"failed":    FileSystemSummaryQuotaEnforcementStateFailed,
}

// GetFileSystemSummaryQuotaEnforcementStateEnumValues Enumerates the set of values for FileSystemSummaryQuotaEnforcementStateEnum
func GetFileSystemSummaryQuotaEnforcementStateEnumValues() []FileSystemSummaryQuotaEnforcementStateEnum {
	values := make([]FileSystemSummaryQuotaEnforcementStateEnum, 0)
	for _, v := range mappingFileSystemSummaryQuotaEnforcementStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFileSystemSummaryQuotaEnforcementStateEnumStringValues Enumerates the set of values in String for FileSystemSummaryQuotaEnforcementStateEnum
func GetFileSystemSummaryQuotaEnforcementStateEnumStringValues() []string {
	return []string{
		"ENABLING",
		"ENABLED",
		"DISABLING",
		"DISABLED",
		"SYNCING",
		"FAILED",
	}
}

// GetMappingFileSystemSummaryQuotaEnforcementStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFileSystemSummaryQuotaEnforcementStateEnum(val string) (FileSystemSummaryQuotaEnforcementStateEnum, bool) {
	enum, ok := mappingFileSystemSummaryQuotaEnforcementStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
