// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RefreshActivity An environment refresh copies data from a source environment to a target environment, making a copy of the source environment onto the target environment. For more information, see Refreshing an Environment (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/refresh-environment.htm).
type RefreshActivity struct {

	// The unique identifier (OCID) of the refresh activity. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name for the refresh activity. Can be changed later.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Fusion environment that is the source environment for the refresh.
	SourceFusionEnvironmentId *string `mandatory:"true" json:"sourceFusionEnvironmentId"`

	// The current state of the refreshActivity.
	LifecycleState RefreshActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Service availability / impact during refresh activity execution up down
	ServiceAvailability RefreshActivityServiceAvailabilityEnum `mandatory:"true" json:"serviceAvailability"`

	// The time the refresh activity is scheduled to start. An RFC3339 formatted datetime string.
	TimeScheduledStart *common.SDKTime `mandatory:"true" json:"timeScheduledStart"`

	// The time the refresh activity is scheduled to end. An RFC3339 formatted datetime string.
	TimeExpectedFinish *common.SDKTime `mandatory:"true" json:"timeExpectedFinish"`

	// The date and time of the most recent source environment backup used for the environment refresh.
	TimeOfRestorationPoint *common.SDKTime `mandatory:"false" json:"timeOfRestorationPoint"`

	// The time the refresh activity actually completed / cancelled / failed. An RFC3339 formatted datetime string.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// The time the refresh activity record was created. An RFC3339 formatted datetime string.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// The time the refresh activity record was updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Represents if the customer opted for Data Masking or not during refreshActivity.
	IsDataMaskingOpted *bool `mandatory:"false" json:"isDataMaskingOpted"`

	// Details of refresh investigation information, each item represents a different issue.
	RefreshIssueDetailsList []RefreshIssueDetails `mandatory:"false" json:"refreshIssueDetailsList"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails RefreshActivityLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m RefreshActivity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RefreshActivity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRefreshActivityLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetRefreshActivityLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingRefreshActivityServiceAvailabilityEnum(string(m.ServiceAvailability)); !ok && m.ServiceAvailability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceAvailability: %s. Supported values are: %s.", m.ServiceAvailability, strings.Join(GetRefreshActivityServiceAvailabilityEnumStringValues(), ",")))
	}

	if _, ok := GetMappingRefreshActivityLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetRefreshActivityLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RefreshActivityLifecycleStateEnum Enum with underlying type: string
type RefreshActivityLifecycleStateEnum string

// Set of constants representing the allowable values for RefreshActivityLifecycleStateEnum
const (
	RefreshActivityLifecycleStateAccepted       RefreshActivityLifecycleStateEnum = "ACCEPTED"
	RefreshActivityLifecycleStateInProgress     RefreshActivityLifecycleStateEnum = "IN_PROGRESS"
	RefreshActivityLifecycleStateNeedsAttention RefreshActivityLifecycleStateEnum = "NEEDS_ATTENTION"
	RefreshActivityLifecycleStateFailed         RefreshActivityLifecycleStateEnum = "FAILED"
	RefreshActivityLifecycleStateSucceeded      RefreshActivityLifecycleStateEnum = "SUCCEEDED"
	RefreshActivityLifecycleStateCanceled       RefreshActivityLifecycleStateEnum = "CANCELED"
)

var mappingRefreshActivityLifecycleStateEnum = map[string]RefreshActivityLifecycleStateEnum{
	"ACCEPTED":        RefreshActivityLifecycleStateAccepted,
	"IN_PROGRESS":     RefreshActivityLifecycleStateInProgress,
	"NEEDS_ATTENTION": RefreshActivityLifecycleStateNeedsAttention,
	"FAILED":          RefreshActivityLifecycleStateFailed,
	"SUCCEEDED":       RefreshActivityLifecycleStateSucceeded,
	"CANCELED":        RefreshActivityLifecycleStateCanceled,
}

var mappingRefreshActivityLifecycleStateEnumLowerCase = map[string]RefreshActivityLifecycleStateEnum{
	"accepted":        RefreshActivityLifecycleStateAccepted,
	"in_progress":     RefreshActivityLifecycleStateInProgress,
	"needs_attention": RefreshActivityLifecycleStateNeedsAttention,
	"failed":          RefreshActivityLifecycleStateFailed,
	"succeeded":       RefreshActivityLifecycleStateSucceeded,
	"canceled":        RefreshActivityLifecycleStateCanceled,
}

// GetRefreshActivityLifecycleStateEnumValues Enumerates the set of values for RefreshActivityLifecycleStateEnum
func GetRefreshActivityLifecycleStateEnumValues() []RefreshActivityLifecycleStateEnum {
	values := make([]RefreshActivityLifecycleStateEnum, 0)
	for _, v := range mappingRefreshActivityLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetRefreshActivityLifecycleStateEnumStringValues Enumerates the set of values in String for RefreshActivityLifecycleStateEnum
func GetRefreshActivityLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"NEEDS_ATTENTION",
		"FAILED",
		"SUCCEEDED",
		"CANCELED",
	}
}

// GetMappingRefreshActivityLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRefreshActivityLifecycleStateEnum(val string) (RefreshActivityLifecycleStateEnum, bool) {
	enum, ok := mappingRefreshActivityLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RefreshActivityServiceAvailabilityEnum Enum with underlying type: string
type RefreshActivityServiceAvailabilityEnum string

// Set of constants representing the allowable values for RefreshActivityServiceAvailabilityEnum
const (
	RefreshActivityServiceAvailabilityAvailable   RefreshActivityServiceAvailabilityEnum = "AVAILABLE"
	RefreshActivityServiceAvailabilityUnavailable RefreshActivityServiceAvailabilityEnum = "UNAVAILABLE"
)

var mappingRefreshActivityServiceAvailabilityEnum = map[string]RefreshActivityServiceAvailabilityEnum{
	"AVAILABLE":   RefreshActivityServiceAvailabilityAvailable,
	"UNAVAILABLE": RefreshActivityServiceAvailabilityUnavailable,
}

var mappingRefreshActivityServiceAvailabilityEnumLowerCase = map[string]RefreshActivityServiceAvailabilityEnum{
	"available":   RefreshActivityServiceAvailabilityAvailable,
	"unavailable": RefreshActivityServiceAvailabilityUnavailable,
}

// GetRefreshActivityServiceAvailabilityEnumValues Enumerates the set of values for RefreshActivityServiceAvailabilityEnum
func GetRefreshActivityServiceAvailabilityEnumValues() []RefreshActivityServiceAvailabilityEnum {
	values := make([]RefreshActivityServiceAvailabilityEnum, 0)
	for _, v := range mappingRefreshActivityServiceAvailabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetRefreshActivityServiceAvailabilityEnumStringValues Enumerates the set of values in String for RefreshActivityServiceAvailabilityEnum
func GetRefreshActivityServiceAvailabilityEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"UNAVAILABLE",
	}
}

// GetMappingRefreshActivityServiceAvailabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRefreshActivityServiceAvailabilityEnum(val string) (RefreshActivityServiceAvailabilityEnum, bool) {
	enum, ok := mappingRefreshActivityServiceAvailabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// RefreshActivityLifecycleDetailsEnum Enum with underlying type: string
type RefreshActivityLifecycleDetailsEnum string

// Set of constants representing the allowable values for RefreshActivityLifecycleDetailsEnum
const (
	RefreshActivityLifecycleDetailsNone               RefreshActivityLifecycleDetailsEnum = "NONE"
	RefreshActivityLifecycleDetailsRollbackaccepted   RefreshActivityLifecycleDetailsEnum = "ROLLBACKACCEPTED"
	RefreshActivityLifecycleDetailsRollbackinprogress RefreshActivityLifecycleDetailsEnum = "ROLLBACKINPROGRESS"
	RefreshActivityLifecycleDetailsRollbacksucceeded  RefreshActivityLifecycleDetailsEnum = "ROLLBACKSUCCEEDED"
	RefreshActivityLifecycleDetailsRollbackfailed     RefreshActivityLifecycleDetailsEnum = "ROLLBACKFAILED"
)

var mappingRefreshActivityLifecycleDetailsEnum = map[string]RefreshActivityLifecycleDetailsEnum{
	"NONE":               RefreshActivityLifecycleDetailsNone,
	"ROLLBACKACCEPTED":   RefreshActivityLifecycleDetailsRollbackaccepted,
	"ROLLBACKINPROGRESS": RefreshActivityLifecycleDetailsRollbackinprogress,
	"ROLLBACKSUCCEEDED":  RefreshActivityLifecycleDetailsRollbacksucceeded,
	"ROLLBACKFAILED":     RefreshActivityLifecycleDetailsRollbackfailed,
}

var mappingRefreshActivityLifecycleDetailsEnumLowerCase = map[string]RefreshActivityLifecycleDetailsEnum{
	"none":               RefreshActivityLifecycleDetailsNone,
	"rollbackaccepted":   RefreshActivityLifecycleDetailsRollbackaccepted,
	"rollbackinprogress": RefreshActivityLifecycleDetailsRollbackinprogress,
	"rollbacksucceeded":  RefreshActivityLifecycleDetailsRollbacksucceeded,
	"rollbackfailed":     RefreshActivityLifecycleDetailsRollbackfailed,
}

// GetRefreshActivityLifecycleDetailsEnumValues Enumerates the set of values for RefreshActivityLifecycleDetailsEnum
func GetRefreshActivityLifecycleDetailsEnumValues() []RefreshActivityLifecycleDetailsEnum {
	values := make([]RefreshActivityLifecycleDetailsEnum, 0)
	for _, v := range mappingRefreshActivityLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetRefreshActivityLifecycleDetailsEnumStringValues Enumerates the set of values in String for RefreshActivityLifecycleDetailsEnum
func GetRefreshActivityLifecycleDetailsEnumStringValues() []string {
	return []string{
		"NONE",
		"ROLLBACKACCEPTED",
		"ROLLBACKINPROGRESS",
		"ROLLBACKSUCCEEDED",
		"ROLLBACKFAILED",
	}
}

// GetMappingRefreshActivityLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRefreshActivityLifecycleDetailsEnum(val string) (RefreshActivityLifecycleDetailsEnum, bool) {
	enum, ok := mappingRefreshActivityLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
