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

// Replication Replications are the primary resource that governs the policy of cross-region replication between source
// and target file systems. Replications are associated with a secondary resource called a ReplicationTarget
// located in another availability domain in the same or different region.
// The replication retrieves the delta of data between two snapshots of a source file system
// and sends it to the associated `ReplicationTarget`, which applies it to the target
// file system. For more information, see File System Replication (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/FSreplication.htm).
type Replication struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	Id *string `mandatory:"true" json:"id"`

	// The current lifecycle state of the replication.
	LifecycleState ReplicationLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My replication`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the replication was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2021-01-04T20:01:29.100Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the source file system.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the target file system.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the ReplicationTarget.
	ReplicationTargetId *string `mandatory:"true" json:"replicationTargetId"`

	// The availability domain that contains the replication. May be unset as a blank or `NULL` value.
	// Example: `Uocm:PHX-AD-2`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Duration in minutes between replication snapshots.
	ReplicationInterval *int64 `mandatory:"false" json:"replicationInterval"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last snapshot that has been replicated completely.
	// Empty if the copy of the initial snapshot is not complete.
	LastSnapshotId *string `mandatory:"false" json:"lastSnapshotId"`

	// The snapshotTime of the most recent recoverable replication snapshot
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2021-04-04T20:01:29.100Z`
	RecoveryPointTime *common.SDKTime `mandatory:"false" json:"recoveryPointTime"`

	// The current state of the snapshot during replication operations.
	DeltaStatus ReplicationDeltaStatusEnum `mandatory:"false" json:"deltaStatus,omitempty"`

	// Additional information about the current 'lifecycleState'.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Percentage progress of the current replication cycle.
	DeltaProgress *int64 `mandatory:"false" json:"deltaProgress"`

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

func (m Replication) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Replication) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicationLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReplicationDeltaStatusEnum(string(m.DeltaStatus)); !ok && m.DeltaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaStatus: %s. Supported values are: %s.", m.DeltaStatus, strings.Join(GetReplicationDeltaStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationLifecycleStateEnum Enum with underlying type: string
type ReplicationLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicationLifecycleStateEnum
const (
	ReplicationLifecycleStateCreating ReplicationLifecycleStateEnum = "CREATING"
	ReplicationLifecycleStateActive   ReplicationLifecycleStateEnum = "ACTIVE"
	ReplicationLifecycleStateDeleting ReplicationLifecycleStateEnum = "DELETING"
	ReplicationLifecycleStateDeleted  ReplicationLifecycleStateEnum = "DELETED"
	ReplicationLifecycleStateFailed   ReplicationLifecycleStateEnum = "FAILED"
)

var mappingReplicationLifecycleStateEnum = map[string]ReplicationLifecycleStateEnum{
	"CREATING": ReplicationLifecycleStateCreating,
	"ACTIVE":   ReplicationLifecycleStateActive,
	"DELETING": ReplicationLifecycleStateDeleting,
	"DELETED":  ReplicationLifecycleStateDeleted,
	"FAILED":   ReplicationLifecycleStateFailed,
}

var mappingReplicationLifecycleStateEnumLowerCase = map[string]ReplicationLifecycleStateEnum{
	"creating": ReplicationLifecycleStateCreating,
	"active":   ReplicationLifecycleStateActive,
	"deleting": ReplicationLifecycleStateDeleting,
	"deleted":  ReplicationLifecycleStateDeleted,
	"failed":   ReplicationLifecycleStateFailed,
}

// GetReplicationLifecycleStateEnumValues Enumerates the set of values for ReplicationLifecycleStateEnum
func GetReplicationLifecycleStateEnumValues() []ReplicationLifecycleStateEnum {
	values := make([]ReplicationLifecycleStateEnum, 0)
	for _, v := range mappingReplicationLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicationLifecycleStateEnum
func GetReplicationLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingReplicationLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationLifecycleStateEnum(val string) (ReplicationLifecycleStateEnum, bool) {
	enum, ok := mappingReplicationLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReplicationDeltaStatusEnum Enum with underlying type: string
type ReplicationDeltaStatusEnum string

// Set of constants representing the allowable values for ReplicationDeltaStatusEnum
const (
	ReplicationDeltaStatusIdle         ReplicationDeltaStatusEnum = "IDLE"
	ReplicationDeltaStatusCapturing    ReplicationDeltaStatusEnum = "CAPTURING"
	ReplicationDeltaStatusApplying     ReplicationDeltaStatusEnum = "APPLYING"
	ReplicationDeltaStatusServiceError ReplicationDeltaStatusEnum = "SERVICE_ERROR"
	ReplicationDeltaStatusUserError    ReplicationDeltaStatusEnum = "USER_ERROR"
	ReplicationDeltaStatusFailed       ReplicationDeltaStatusEnum = "FAILED"
	ReplicationDeltaStatusTransferring ReplicationDeltaStatusEnum = "TRANSFERRING"
)

var mappingReplicationDeltaStatusEnum = map[string]ReplicationDeltaStatusEnum{
	"IDLE":          ReplicationDeltaStatusIdle,
	"CAPTURING":     ReplicationDeltaStatusCapturing,
	"APPLYING":      ReplicationDeltaStatusApplying,
	"SERVICE_ERROR": ReplicationDeltaStatusServiceError,
	"USER_ERROR":    ReplicationDeltaStatusUserError,
	"FAILED":        ReplicationDeltaStatusFailed,
	"TRANSFERRING":  ReplicationDeltaStatusTransferring,
}

var mappingReplicationDeltaStatusEnumLowerCase = map[string]ReplicationDeltaStatusEnum{
	"idle":          ReplicationDeltaStatusIdle,
	"capturing":     ReplicationDeltaStatusCapturing,
	"applying":      ReplicationDeltaStatusApplying,
	"service_error": ReplicationDeltaStatusServiceError,
	"user_error":    ReplicationDeltaStatusUserError,
	"failed":        ReplicationDeltaStatusFailed,
	"transferring":  ReplicationDeltaStatusTransferring,
}

// GetReplicationDeltaStatusEnumValues Enumerates the set of values for ReplicationDeltaStatusEnum
func GetReplicationDeltaStatusEnumValues() []ReplicationDeltaStatusEnum {
	values := make([]ReplicationDeltaStatusEnum, 0)
	for _, v := range mappingReplicationDeltaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationDeltaStatusEnumStringValues Enumerates the set of values in String for ReplicationDeltaStatusEnum
func GetReplicationDeltaStatusEnumStringValues() []string {
	return []string{
		"IDLE",
		"CAPTURING",
		"APPLYING",
		"SERVICE_ERROR",
		"USER_ERROR",
		"FAILED",
		"TRANSFERRING",
	}
}

// GetMappingReplicationDeltaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationDeltaStatusEnum(val string) (ReplicationDeltaStatusEnum, bool) {
	enum, ok := mappingReplicationDeltaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
