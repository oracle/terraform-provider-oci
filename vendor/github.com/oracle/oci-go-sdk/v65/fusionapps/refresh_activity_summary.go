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

// RefreshActivitySummary Summary of the refresh activity.
type RefreshActivitySummary struct {

	// The unique identifier (OCID) of the refresh activity. Can't be changed after creation.
	Id *string `mandatory:"true" json:"id"`

	// A friendly name for the refresh activity. Can be changed later.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the Fusion environment that is the source environment for the refresh.
	SourceFusionEnvironmentId *string `mandatory:"true" json:"sourceFusionEnvironmentId"`

	// The current state of the refresh activity. Valid values are Scheduled, In progress , Failed, Completed.
	LifecycleState RefreshActivityLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time the refresh activity is scheduled to start. An RFC3339 formatted datetime string.
	TimeScheduledStart *common.SDKTime `mandatory:"true" json:"timeScheduledStart"`

	// The time the refresh activity is scheduled to end. An RFC3339 formatted datetime string.
	TimeExpectedFinish *common.SDKTime `mandatory:"true" json:"timeExpectedFinish"`

	// Service availability / impact during refresh activity execution, up down
	ServiceAvailability RefreshActivityServiceAvailabilityEnum `mandatory:"true" json:"serviceAvailability"`

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

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails RefreshActivityLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// Details of refresh investigation information, each item represents a different issue.
	RefreshIssueDetailsList []RefreshIssueDetails `mandatory:"false" json:"refreshIssueDetailsList"`
}

func (m RefreshActivitySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RefreshActivitySummary) ValidateEnumValue() (bool, error) {
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
