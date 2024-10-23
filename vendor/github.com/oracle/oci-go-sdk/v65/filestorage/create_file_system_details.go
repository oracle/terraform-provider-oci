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

// CreateFileSystemDetails Details for creating the file system.
type CreateFileSystemDetails struct {

	// The availability domain to create the file system in.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My file system`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system.
	KmsKeyId *string `mandatory:"false" json:"kmsKeyId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system.
	// See Cloning a File System (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
	SourceSnapshotId *string `mandatory:"false" json:"sourceSnapshotId"`

	// Specifies whether the clone file system is attached to its parent file system.
	// If the value is set to 'DETACH', then the file system will be created, which is deep copied from the snapshot
	// specified by sourceSnapshotId, else will remain attached to its parent.
	CloneAttachStatus CreateFileSystemDetailsCloneAttachStatusEnum `mandatory:"false" json:"cloneAttachStatus,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the associated file system snapshot policy, which
	// controls the frequency of snapshot creation and retention period of the taken snapshots.
	// May be unset as a blank value.
	FilesystemSnapshotPolicyId *string `mandatory:"false" json:"filesystemSnapshotPolicyId"`
}

func (m CreateFileSystemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateFileSystemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateFileSystemDetailsCloneAttachStatusEnum(string(m.CloneAttachStatus)); !ok && m.CloneAttachStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CloneAttachStatus: %s. Supported values are: %s.", m.CloneAttachStatus, strings.Join(GetCreateFileSystemDetailsCloneAttachStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateFileSystemDetailsCloneAttachStatusEnum Enum with underlying type: string
type CreateFileSystemDetailsCloneAttachStatusEnum string

// Set of constants representing the allowable values for CreateFileSystemDetailsCloneAttachStatusEnum
const (
	CreateFileSystemDetailsCloneAttachStatusDetach CreateFileSystemDetailsCloneAttachStatusEnum = "DETACH"
	CreateFileSystemDetailsCloneAttachStatusAttach CreateFileSystemDetailsCloneAttachStatusEnum = "ATTACH"
)

var mappingCreateFileSystemDetailsCloneAttachStatusEnum = map[string]CreateFileSystemDetailsCloneAttachStatusEnum{
	"DETACH": CreateFileSystemDetailsCloneAttachStatusDetach,
	"ATTACH": CreateFileSystemDetailsCloneAttachStatusAttach,
}

var mappingCreateFileSystemDetailsCloneAttachStatusEnumLowerCase = map[string]CreateFileSystemDetailsCloneAttachStatusEnum{
	"detach": CreateFileSystemDetailsCloneAttachStatusDetach,
	"attach": CreateFileSystemDetailsCloneAttachStatusAttach,
}

// GetCreateFileSystemDetailsCloneAttachStatusEnumValues Enumerates the set of values for CreateFileSystemDetailsCloneAttachStatusEnum
func GetCreateFileSystemDetailsCloneAttachStatusEnumValues() []CreateFileSystemDetailsCloneAttachStatusEnum {
	values := make([]CreateFileSystemDetailsCloneAttachStatusEnum, 0)
	for _, v := range mappingCreateFileSystemDetailsCloneAttachStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateFileSystemDetailsCloneAttachStatusEnumStringValues Enumerates the set of values in String for CreateFileSystemDetailsCloneAttachStatusEnum
func GetCreateFileSystemDetailsCloneAttachStatusEnumStringValues() []string {
	return []string{
		"DETACH",
		"ATTACH",
	}
}

// GetMappingCreateFileSystemDetailsCloneAttachStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateFileSystemDetailsCloneAttachStatusEnum(val string) (CreateFileSystemDetailsCloneAttachStatusEnum, bool) {
	enum, ok := mappingCreateFileSystemDetailsCloneAttachStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
