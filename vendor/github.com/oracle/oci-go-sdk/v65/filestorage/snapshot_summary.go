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

// SnapshotSummary Summary information for a snapshot.
type SnapshotSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the file system from which the snapshot was created.
	FileSystemId *string `mandatory:"true" json:"fileSystemId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the snapshot.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the snapshot.
	LifecycleState SnapshotSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Name of the snapshot. This value is immutable.
	// Avoid entering confidential information.
	// Example: `Sunday`
	Name *string `mandatory:"true" json:"name"`

	// The date and time the snapshot was created, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Specifies the generation type of the snapshot.
	SnapshotType SnapshotSummarySnapshotTypeEnum `mandatory:"false" json:"snapshotType,omitempty"`

	// The date and time the snapshot was taken, expressed
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// This value might be the same or different from `timeCreated` depending
	// on the following factors:
	// - If the snapshot is created in the original file system directory.
	// - If the snapshot is cloned from a file system.
	// - If the snapshot is replicated from a file system.
	// Example: `2020-08-25T21:10:29.600Z`
	SnapshotTime *common.SDKTime `mandatory:"false" json:"snapshotTime"`

	// The time when this snapshot will be deleted.
	ExpirationTime *common.SDKTime `mandatory:"false" json:"expirationTime"`

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

	// Free-form tags for this resource. Each tag is a simple key-value pair
	//  with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m SnapshotSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SnapshotSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSnapshotSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSnapshotSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSnapshotSummarySnapshotTypeEnum(string(m.SnapshotType)); !ok && m.SnapshotType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SnapshotType: %s. Supported values are: %s.", m.SnapshotType, strings.Join(GetSnapshotSummarySnapshotTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SnapshotSummaryLifecycleStateEnum Enum with underlying type: string
type SnapshotSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for SnapshotSummaryLifecycleStateEnum
const (
	SnapshotSummaryLifecycleStateCreating SnapshotSummaryLifecycleStateEnum = "CREATING"
	SnapshotSummaryLifecycleStateActive   SnapshotSummaryLifecycleStateEnum = "ACTIVE"
	SnapshotSummaryLifecycleStateDeleting SnapshotSummaryLifecycleStateEnum = "DELETING"
	SnapshotSummaryLifecycleStateDeleted  SnapshotSummaryLifecycleStateEnum = "DELETED"
)

var mappingSnapshotSummaryLifecycleStateEnum = map[string]SnapshotSummaryLifecycleStateEnum{
	"CREATING": SnapshotSummaryLifecycleStateCreating,
	"ACTIVE":   SnapshotSummaryLifecycleStateActive,
	"DELETING": SnapshotSummaryLifecycleStateDeleting,
	"DELETED":  SnapshotSummaryLifecycleStateDeleted,
}

var mappingSnapshotSummaryLifecycleStateEnumLowerCase = map[string]SnapshotSummaryLifecycleStateEnum{
	"creating": SnapshotSummaryLifecycleStateCreating,
	"active":   SnapshotSummaryLifecycleStateActive,
	"deleting": SnapshotSummaryLifecycleStateDeleting,
	"deleted":  SnapshotSummaryLifecycleStateDeleted,
}

// GetSnapshotSummaryLifecycleStateEnumValues Enumerates the set of values for SnapshotSummaryLifecycleStateEnum
func GetSnapshotSummaryLifecycleStateEnumValues() []SnapshotSummaryLifecycleStateEnum {
	values := make([]SnapshotSummaryLifecycleStateEnum, 0)
	for _, v := range mappingSnapshotSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for SnapshotSummaryLifecycleStateEnum
func GetSnapshotSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingSnapshotSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotSummaryLifecycleStateEnum(val string) (SnapshotSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingSnapshotSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SnapshotSummarySnapshotTypeEnum Enum with underlying type: string
type SnapshotSummarySnapshotTypeEnum string

// Set of constants representing the allowable values for SnapshotSummarySnapshotTypeEnum
const (
	SnapshotSummarySnapshotTypeUser        SnapshotSummarySnapshotTypeEnum = "USER"
	SnapshotSummarySnapshotTypePolicyBased SnapshotSummarySnapshotTypeEnum = "POLICY_BASED"
	SnapshotSummarySnapshotTypeReplication SnapshotSummarySnapshotTypeEnum = "REPLICATION"
)

var mappingSnapshotSummarySnapshotTypeEnum = map[string]SnapshotSummarySnapshotTypeEnum{
	"USER":         SnapshotSummarySnapshotTypeUser,
	"POLICY_BASED": SnapshotSummarySnapshotTypePolicyBased,
	"REPLICATION":  SnapshotSummarySnapshotTypeReplication,
}

var mappingSnapshotSummarySnapshotTypeEnumLowerCase = map[string]SnapshotSummarySnapshotTypeEnum{
	"user":         SnapshotSummarySnapshotTypeUser,
	"policy_based": SnapshotSummarySnapshotTypePolicyBased,
	"replication":  SnapshotSummarySnapshotTypeReplication,
}

// GetSnapshotSummarySnapshotTypeEnumValues Enumerates the set of values for SnapshotSummarySnapshotTypeEnum
func GetSnapshotSummarySnapshotTypeEnumValues() []SnapshotSummarySnapshotTypeEnum {
	values := make([]SnapshotSummarySnapshotTypeEnum, 0)
	for _, v := range mappingSnapshotSummarySnapshotTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetSnapshotSummarySnapshotTypeEnumStringValues Enumerates the set of values in String for SnapshotSummarySnapshotTypeEnum
func GetSnapshotSummarySnapshotTypeEnumStringValues() []string {
	return []string{
		"USER",
		"POLICY_BASED",
		"REPLICATION",
	}
}

// GetMappingSnapshotSummarySnapshotTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSnapshotSummarySnapshotTypeEnum(val string) (SnapshotSummarySnapshotTypeEnum, bool) {
	enum, ok := mappingSnapshotSummarySnapshotTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
