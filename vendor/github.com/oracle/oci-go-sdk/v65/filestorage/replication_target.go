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

// ReplicationTarget Replication targets are associated with a primary resource called a Replication
// located in another availability domain in the same or different region.
// The replication retrieves the delta of data between two snapshots of a source file system
// and sends it to the associated `ReplicationTarget`,  which applies it to the target
// file system.
// All operations (except `DELETE`) must be done using the associated replication resource.
// Deleting a `ReplicationTarget` allows the target file system to be exported.
// Deleting a `ReplicationTarget` does not delete the associated `Replication` resource, but places it in a `FAILED` state.
// For more information, see File System Replication (https://docs.cloud.oracle.com/iaas/Content/File/Tasks/FSreplication.htm).
type ReplicationTarget struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this replication.
	LifecycleState ReplicationTargetLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. This name is same as the replication display name for the associated resource.
	// Example: `My Replication`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the replication target was created in target region.
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2021-01-04T20:01:29.100Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of source filesystem.
	SourceId *string `mandatory:"true" json:"sourceId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of target filesystem.
	TargetId *string `mandatory:"true" json:"targetId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of replication.
	ReplicationId *string `mandatory:"true" json:"replicationId"`

	// The availability domain the replication resource is in. May be unset
	// as a blank or NULL value.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the last snapshot snapshot which was completely applied to the target file system.
	// Empty while the initial snapshot is being applied.
	LastSnapshotId *string `mandatory:"false" json:"lastSnapshotId"`

	// The snapshotTime of the most recent recoverable replication snapshot
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2021-04-04T20:01:29.100Z`
	RecoveryPointTime *common.SDKTime `mandatory:"false" json:"recoveryPointTime"`

	// The current state of the snapshot during replication operations.
	DeltaStatus ReplicationTargetDeltaStatusEnum `mandatory:"false" json:"deltaStatus,omitempty"`

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

	// Additional information about the current `lifecycleState`.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m ReplicationTarget) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationTarget) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationTargetLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicationTargetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingReplicationTargetDeltaStatusEnum(string(m.DeltaStatus)); !ok && m.DeltaStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DeltaStatus: %s. Supported values are: %s.", m.DeltaStatus, strings.Join(GetReplicationTargetDeltaStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationTargetLifecycleStateEnum Enum with underlying type: string
type ReplicationTargetLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicationTargetLifecycleStateEnum
const (
	ReplicationTargetLifecycleStateCreating ReplicationTargetLifecycleStateEnum = "CREATING"
	ReplicationTargetLifecycleStateActive   ReplicationTargetLifecycleStateEnum = "ACTIVE"
	ReplicationTargetLifecycleStateDeleting ReplicationTargetLifecycleStateEnum = "DELETING"
	ReplicationTargetLifecycleStateDeleted  ReplicationTargetLifecycleStateEnum = "DELETED"
	ReplicationTargetLifecycleStateFailed   ReplicationTargetLifecycleStateEnum = "FAILED"
)

var mappingReplicationTargetLifecycleStateEnum = map[string]ReplicationTargetLifecycleStateEnum{
	"CREATING": ReplicationTargetLifecycleStateCreating,
	"ACTIVE":   ReplicationTargetLifecycleStateActive,
	"DELETING": ReplicationTargetLifecycleStateDeleting,
	"DELETED":  ReplicationTargetLifecycleStateDeleted,
	"FAILED":   ReplicationTargetLifecycleStateFailed,
}

var mappingReplicationTargetLifecycleStateEnumLowerCase = map[string]ReplicationTargetLifecycleStateEnum{
	"creating": ReplicationTargetLifecycleStateCreating,
	"active":   ReplicationTargetLifecycleStateActive,
	"deleting": ReplicationTargetLifecycleStateDeleting,
	"deleted":  ReplicationTargetLifecycleStateDeleted,
	"failed":   ReplicationTargetLifecycleStateFailed,
}

// GetReplicationTargetLifecycleStateEnumValues Enumerates the set of values for ReplicationTargetLifecycleStateEnum
func GetReplicationTargetLifecycleStateEnumValues() []ReplicationTargetLifecycleStateEnum {
	values := make([]ReplicationTargetLifecycleStateEnum, 0)
	for _, v := range mappingReplicationTargetLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationTargetLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicationTargetLifecycleStateEnum
func GetReplicationTargetLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingReplicationTargetLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationTargetLifecycleStateEnum(val string) (ReplicationTargetLifecycleStateEnum, bool) {
	enum, ok := mappingReplicationTargetLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReplicationTargetDeltaStatusEnum Enum with underlying type: string
type ReplicationTargetDeltaStatusEnum string

// Set of constants representing the allowable values for ReplicationTargetDeltaStatusEnum
const (
	ReplicationTargetDeltaStatusIdle         ReplicationTargetDeltaStatusEnum = "IDLE"
	ReplicationTargetDeltaStatusCapturing    ReplicationTargetDeltaStatusEnum = "CAPTURING"
	ReplicationTargetDeltaStatusApplying     ReplicationTargetDeltaStatusEnum = "APPLYING"
	ReplicationTargetDeltaStatusServiceError ReplicationTargetDeltaStatusEnum = "SERVICE_ERROR"
	ReplicationTargetDeltaStatusUserError    ReplicationTargetDeltaStatusEnum = "USER_ERROR"
	ReplicationTargetDeltaStatusFailed       ReplicationTargetDeltaStatusEnum = "FAILED"
	ReplicationTargetDeltaStatusTransferring ReplicationTargetDeltaStatusEnum = "TRANSFERRING"
)

var mappingReplicationTargetDeltaStatusEnum = map[string]ReplicationTargetDeltaStatusEnum{
	"IDLE":          ReplicationTargetDeltaStatusIdle,
	"CAPTURING":     ReplicationTargetDeltaStatusCapturing,
	"APPLYING":      ReplicationTargetDeltaStatusApplying,
	"SERVICE_ERROR": ReplicationTargetDeltaStatusServiceError,
	"USER_ERROR":    ReplicationTargetDeltaStatusUserError,
	"FAILED":        ReplicationTargetDeltaStatusFailed,
	"TRANSFERRING":  ReplicationTargetDeltaStatusTransferring,
}

var mappingReplicationTargetDeltaStatusEnumLowerCase = map[string]ReplicationTargetDeltaStatusEnum{
	"idle":          ReplicationTargetDeltaStatusIdle,
	"capturing":     ReplicationTargetDeltaStatusCapturing,
	"applying":      ReplicationTargetDeltaStatusApplying,
	"service_error": ReplicationTargetDeltaStatusServiceError,
	"user_error":    ReplicationTargetDeltaStatusUserError,
	"failed":        ReplicationTargetDeltaStatusFailed,
	"transferring":  ReplicationTargetDeltaStatusTransferring,
}

// GetReplicationTargetDeltaStatusEnumValues Enumerates the set of values for ReplicationTargetDeltaStatusEnum
func GetReplicationTargetDeltaStatusEnumValues() []ReplicationTargetDeltaStatusEnum {
	values := make([]ReplicationTargetDeltaStatusEnum, 0)
	for _, v := range mappingReplicationTargetDeltaStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationTargetDeltaStatusEnumStringValues Enumerates the set of values in String for ReplicationTargetDeltaStatusEnum
func GetReplicationTargetDeltaStatusEnumStringValues() []string {
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

// GetMappingReplicationTargetDeltaStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationTargetDeltaStatusEnum(val string) (ReplicationTargetDeltaStatusEnum, bool) {
	enum, ok := mappingReplicationTargetDeltaStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
