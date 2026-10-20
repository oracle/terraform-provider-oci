// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataVerificationDetail Summary/status information used to enable and drive the Data Verification UI.
type DataVerificationDetail struct {

	// The OCID of the resource being referenced.
	MigrationId *string `mandatory:"true" json:"migrationId"`

	// Overall lifecycle state of Data Verification generation for the migration.
	// This is a single, feature-level state (all Data Verification reports are generated as one batch).
	LifecycleState DataVerificationDetailLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// A human-readable string that provides more details about the current lifecycle state.
	// This field is intended to follow OCI patterns and may include an internal error code and message
	// when `lifecycleState` is `FAILED` (for example, `DMS-xxxx: <message>`).
	// When `lifecycleState` is `SUCCEEDED`, this field may contain a non-failing disclaimer, such as
	// unavailable table row-count comparisons caused by missing optimizer statistics.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m DataVerificationDetail) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataVerificationDetail) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDataVerificationDetailLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDataVerificationDetailLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DataVerificationDetailLifecycleStateEnum Enum with underlying type: string
type DataVerificationDetailLifecycleStateEnum string

// Set of constants representing the allowable values for DataVerificationDetailLifecycleStateEnum
const (
	DataVerificationDetailLifecycleStateAccepted       DataVerificationDetailLifecycleStateEnum = "ACCEPTED"
	DataVerificationDetailLifecycleStateInProgress     DataVerificationDetailLifecycleStateEnum = "IN_PROGRESS"
	DataVerificationDetailLifecycleStateWaiting        DataVerificationDetailLifecycleStateEnum = "WAITING"
	DataVerificationDetailLifecycleStateFailed         DataVerificationDetailLifecycleStateEnum = "FAILED"
	DataVerificationDetailLifecycleStateSucceeded      DataVerificationDetailLifecycleStateEnum = "SUCCEEDED"
	DataVerificationDetailLifecycleStateCanceling      DataVerificationDetailLifecycleStateEnum = "CANCELING"
	DataVerificationDetailLifecycleStateCanceled       DataVerificationDetailLifecycleStateEnum = "CANCELED"
	DataVerificationDetailLifecycleStateNeedsAttention DataVerificationDetailLifecycleStateEnum = "NEEDS_ATTENTION"
)

var mappingDataVerificationDetailLifecycleStateEnum = map[string]DataVerificationDetailLifecycleStateEnum{
	"ACCEPTED":        DataVerificationDetailLifecycleStateAccepted,
	"IN_PROGRESS":     DataVerificationDetailLifecycleStateInProgress,
	"WAITING":         DataVerificationDetailLifecycleStateWaiting,
	"FAILED":          DataVerificationDetailLifecycleStateFailed,
	"SUCCEEDED":       DataVerificationDetailLifecycleStateSucceeded,
	"CANCELING":       DataVerificationDetailLifecycleStateCanceling,
	"CANCELED":        DataVerificationDetailLifecycleStateCanceled,
	"NEEDS_ATTENTION": DataVerificationDetailLifecycleStateNeedsAttention,
}

var mappingDataVerificationDetailLifecycleStateEnumLowerCase = map[string]DataVerificationDetailLifecycleStateEnum{
	"accepted":        DataVerificationDetailLifecycleStateAccepted,
	"in_progress":     DataVerificationDetailLifecycleStateInProgress,
	"waiting":         DataVerificationDetailLifecycleStateWaiting,
	"failed":          DataVerificationDetailLifecycleStateFailed,
	"succeeded":       DataVerificationDetailLifecycleStateSucceeded,
	"canceling":       DataVerificationDetailLifecycleStateCanceling,
	"canceled":        DataVerificationDetailLifecycleStateCanceled,
	"needs_attention": DataVerificationDetailLifecycleStateNeedsAttention,
}

// GetDataVerificationDetailLifecycleStateEnumValues Enumerates the set of values for DataVerificationDetailLifecycleStateEnum
func GetDataVerificationDetailLifecycleStateEnumValues() []DataVerificationDetailLifecycleStateEnum {
	values := make([]DataVerificationDetailLifecycleStateEnum, 0)
	for _, v := range mappingDataVerificationDetailLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDataVerificationDetailLifecycleStateEnumStringValues Enumerates the set of values in String for DataVerificationDetailLifecycleStateEnum
func GetDataVerificationDetailLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
	}
}

// GetMappingDataVerificationDetailLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDataVerificationDetailLifecycleStateEnum(val string) (DataVerificationDetailLifecycleStateEnum, bool) {
	enum, ok := mappingDataVerificationDetailLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
