// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ReplicationProgress Progress of a migration asset's replication process.
type ReplicationProgress struct {

	// Percentage of the current replication progress from 0 to 100.
	Percentage *int `mandatory:"true" json:"percentage"`

	// Status of the current replication progress. It can be None or InProgress.
	Status ReplicationProgressStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Start time of the current replication process
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// Start time of the last replication process. It can be Completed or Failed.
	TimeOflastReplicationStart *common.SDKTime `mandatory:"false" json:"timeOflastReplicationStart"`

	// End time of the last replication process. It can be Completed or Failed.
	TimeOfLastReplicationEnd *common.SDKTime `mandatory:"false" json:"timeOfLastReplicationEnd"`

	// End time of the last successful replication process, which has been completed.
	TimeOfLastReplicationSuccess *common.SDKTime `mandatory:"false" json:"timeOfLastReplicationSuccess"`

	// Status of the last replication task. It can be Completed or Failed.
	LastReplicationStatus ReplicationProgressLastReplicationStatusEnum `mandatory:"false" json:"lastReplicationStatus,omitempty"`

	// Error message if the last finished replication failed.
	LastReplicationError *string `mandatory:"false" json:"lastReplicationError"`
}

func (m ReplicationProgress) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReplicationProgress) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingReplicationProgressStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetReplicationProgressStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReplicationProgressLastReplicationStatusEnum(string(m.LastReplicationStatus)); !ok && m.LastReplicationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LastReplicationStatus: %s. Supported values are: %s.", m.LastReplicationStatus, strings.Join(GetReplicationProgressLastReplicationStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReplicationProgressStatusEnum Enum with underlying type: string
type ReplicationProgressStatusEnum string

// Set of constants representing the allowable values for ReplicationProgressStatusEnum
const (
	ReplicationProgressStatusNone       ReplicationProgressStatusEnum = "NONE"
	ReplicationProgressStatusInProgress ReplicationProgressStatusEnum = "IN_PROGRESS"
)

var mappingReplicationProgressStatusEnum = map[string]ReplicationProgressStatusEnum{
	"NONE":        ReplicationProgressStatusNone,
	"IN_PROGRESS": ReplicationProgressStatusInProgress,
}

var mappingReplicationProgressStatusEnumLowerCase = map[string]ReplicationProgressStatusEnum{
	"none":        ReplicationProgressStatusNone,
	"in_progress": ReplicationProgressStatusInProgress,
}

// GetReplicationProgressStatusEnumValues Enumerates the set of values for ReplicationProgressStatusEnum
func GetReplicationProgressStatusEnumValues() []ReplicationProgressStatusEnum {
	values := make([]ReplicationProgressStatusEnum, 0)
	for _, v := range mappingReplicationProgressStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationProgressStatusEnumStringValues Enumerates the set of values in String for ReplicationProgressStatusEnum
func GetReplicationProgressStatusEnumStringValues() []string {
	return []string{
		"NONE",
		"IN_PROGRESS",
	}
}

// GetMappingReplicationProgressStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationProgressStatusEnum(val string) (ReplicationProgressStatusEnum, bool) {
	enum, ok := mappingReplicationProgressStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReplicationProgressLastReplicationStatusEnum Enum with underlying type: string
type ReplicationProgressLastReplicationStatusEnum string

// Set of constants representing the allowable values for ReplicationProgressLastReplicationStatusEnum
const (
	ReplicationProgressLastReplicationStatusNone      ReplicationProgressLastReplicationStatusEnum = "NONE"
	ReplicationProgressLastReplicationStatusCompleted ReplicationProgressLastReplicationStatusEnum = "COMPLETED"
	ReplicationProgressLastReplicationStatusFailed    ReplicationProgressLastReplicationStatusEnum = "FAILED"
)

var mappingReplicationProgressLastReplicationStatusEnum = map[string]ReplicationProgressLastReplicationStatusEnum{
	"NONE":      ReplicationProgressLastReplicationStatusNone,
	"COMPLETED": ReplicationProgressLastReplicationStatusCompleted,
	"FAILED":    ReplicationProgressLastReplicationStatusFailed,
}

var mappingReplicationProgressLastReplicationStatusEnumLowerCase = map[string]ReplicationProgressLastReplicationStatusEnum{
	"none":      ReplicationProgressLastReplicationStatusNone,
	"completed": ReplicationProgressLastReplicationStatusCompleted,
	"failed":    ReplicationProgressLastReplicationStatusFailed,
}

// GetReplicationProgressLastReplicationStatusEnumValues Enumerates the set of values for ReplicationProgressLastReplicationStatusEnum
func GetReplicationProgressLastReplicationStatusEnumValues() []ReplicationProgressLastReplicationStatusEnum {
	values := make([]ReplicationProgressLastReplicationStatusEnum, 0)
	for _, v := range mappingReplicationProgressLastReplicationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetReplicationProgressLastReplicationStatusEnumStringValues Enumerates the set of values in String for ReplicationProgressLastReplicationStatusEnum
func GetReplicationProgressLastReplicationStatusEnumStringValues() []string {
	return []string{
		"NONE",
		"COMPLETED",
		"FAILED",
	}
}

// GetMappingReplicationProgressLastReplicationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReplicationProgressLastReplicationStatusEnum(val string) (ReplicationProgressLastReplicationStatusEnum, bool) {
	enum, ok := mappingReplicationProgressLastReplicationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
