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

// ReplicationSummary Summary information for a replication.
type ReplicationSummary struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the replication.
	Id *string `mandatory:"true" json:"id"`

	// The current state of this replication.
	// This resource can be in a `FAILED` state if replication target is deleted instead of the replication resource.
	LifecycleState ReplicationSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `My replication`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the replication was created
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2020-02-04T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The availability domain the replication is in. The replication must be in the same availability domain as the source file system.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Locks associated with this resource.
	Locks []ResourceLock `mandatory:"false" json:"locks"`

	// Duration in minutes between replication snapshots.
	ReplicationInterval *int64 `mandatory:"false" json:"replicationInterval"`

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

	// The `snapshotTime` of the most recent recoverable replication snapshot
	// in RFC 3339 (https://tools.ietf.org/rfc/rfc3339) timestamp format.
	// Example: `2021-04-04T20:01:29.100Z`
	RecoveryPointTime *common.SDKTime `mandatory:"false" json:"recoveryPointTime"`
}

func (m ReplicationSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReplicationSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReplicationSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationSummaryLifecycleStateEnum Enum with underlying type: string
type ReplicationSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for ReplicationSummaryLifecycleStateEnum
const (
	ReplicationSummaryLifecycleStateCreating ReplicationSummaryLifecycleStateEnum = "CREATING"
	ReplicationSummaryLifecycleStateActive   ReplicationSummaryLifecycleStateEnum = "ACTIVE"
	ReplicationSummaryLifecycleStateDeleting ReplicationSummaryLifecycleStateEnum = "DELETING"
	ReplicationSummaryLifecycleStateDeleted  ReplicationSummaryLifecycleStateEnum = "DELETED"
	ReplicationSummaryLifecycleStateFailed   ReplicationSummaryLifecycleStateEnum = "FAILED"
)

var mappingReplicationSummaryLifecycleStateEnum = map[string]ReplicationSummaryLifecycleStateEnum{
	"CREATING": ReplicationSummaryLifecycleStateCreating,
	"ACTIVE":   ReplicationSummaryLifecycleStateActive,
	"DELETING": ReplicationSummaryLifecycleStateDeleting,
	"DELETED":  ReplicationSummaryLifecycleStateDeleted,
	"FAILED":   ReplicationSummaryLifecycleStateFailed,
}

var mappingReplicationSummaryLifecycleStateEnumLowerCase = map[string]ReplicationSummaryLifecycleStateEnum{
	"creating": ReplicationSummaryLifecycleStateCreating,
	"active":   ReplicationSummaryLifecycleStateActive,
	"deleting": ReplicationSummaryLifecycleStateDeleting,
	"deleted":  ReplicationSummaryLifecycleStateDeleted,
	"failed":   ReplicationSummaryLifecycleStateFailed,
}

// GetReplicationSummaryLifecycleStateEnumValues Enumerates the set of values for ReplicationSummaryLifecycleStateEnum
func GetReplicationSummaryLifecycleStateEnumValues() []ReplicationSummaryLifecycleStateEnum {
	values := make([]ReplicationSummaryLifecycleStateEnum, 0)
	for _, v := range mappingReplicationSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for ReplicationSummaryLifecycleStateEnum
func GetReplicationSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingReplicationSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationSummaryLifecycleStateEnum(val string) (ReplicationSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingReplicationSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
