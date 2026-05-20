// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Analytics API
//
// Use the Resource Analytics API to manage Resource Analytics Instances.
//

package resourceanalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TenancyAttachment A TenancyAttachment is a customers' tenancy attached to a ResourceAnalyticsInstance. Attached tenancies will be included in analytics collection.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to
// an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type TenancyAttachment struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the TenancyAttachment.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ResourceAnalyticsInstance associated with this TenancyAttachment.
	ResourceAnalyticsInstanceId *string `mandatory:"true" json:"resourceAnalyticsInstanceId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy associated with this TenancyAttachment.
	TenancyId *string `mandatory:"true" json:"tenancyId"`

	// Whether the tenancy is the tenancy used when creating Resource Analytics Instance.
	IsReportingTenancy *bool `mandatory:"true" json:"isReportingTenancy"`

	// List of monitored regions with their data population status.
	MonitoredRegions []TenancyAttachmentMonitoredRegionSummary `mandatory:"true" json:"monitoredRegions"`

	// The overall status of the data population from the tenancy.
	DataPopulationStatus TenancyAttachmentDataPopulationStatusEnum `mandatory:"true" json:"dataPopulationStatus"`

	// The date and time the TenancyAttachment was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current state of the TenancyAttachment.
	LifecycleState TenancyAttachmentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A description of the tenancy.
	Description *string `mandatory:"false" json:"description"`

	// The date and time the data population tasks started, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeDataPopulationStarted *common.SDKTime `mandatory:"false" json:"timeDataPopulationStarted"`

	// The date and time the data population tasks completed, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeDataPopulationEnded *common.SDKTime `mandatory:"false" json:"timeDataPopulationEnded"`

	// The date and time the TenancyAttachment was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message that describes the current state of the TenancyAttachment in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m TenancyAttachment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TenancyAttachment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTenancyAttachmentDataPopulationStatusEnum(string(m.DataPopulationStatus)); !ok && m.DataPopulationStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataPopulationStatus: %s. Supported values are: %s.", m.DataPopulationStatus, strings.Join(GetTenancyAttachmentDataPopulationStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTenancyAttachmentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetTenancyAttachmentLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TenancyAttachmentDataPopulationStatusEnum Enum with underlying type: string
type TenancyAttachmentDataPopulationStatusEnum string

// Set of constants representing the allowable values for TenancyAttachmentDataPopulationStatusEnum
const (
	TenancyAttachmentDataPopulationStatusAccepted           TenancyAttachmentDataPopulationStatusEnum = "ACCEPTED"
	TenancyAttachmentDataPopulationStatusInProgress         TenancyAttachmentDataPopulationStatusEnum = "IN_PROGRESS"
	TenancyAttachmentDataPopulationStatusWaiting            TenancyAttachmentDataPopulationStatusEnum = "WAITING"
	TenancyAttachmentDataPopulationStatusFailed             TenancyAttachmentDataPopulationStatusEnum = "FAILED"
	TenancyAttachmentDataPopulationStatusSucceeded          TenancyAttachmentDataPopulationStatusEnum = "SUCCEEDED"
	TenancyAttachmentDataPopulationStatusCanceling          TenancyAttachmentDataPopulationStatusEnum = "CANCELING"
	TenancyAttachmentDataPopulationStatusCanceled           TenancyAttachmentDataPopulationStatusEnum = "CANCELED"
	TenancyAttachmentDataPopulationStatusNeedsAttention     TenancyAttachmentDataPopulationStatusEnum = "NEEDS_ATTENTION"
	TenancyAttachmentDataPopulationStatusPartiallySucceeded TenancyAttachmentDataPopulationStatusEnum = "PARTIALLY_SUCCEEDED"
)

var mappingTenancyAttachmentDataPopulationStatusEnum = map[string]TenancyAttachmentDataPopulationStatusEnum{
	"ACCEPTED":            TenancyAttachmentDataPopulationStatusAccepted,
	"IN_PROGRESS":         TenancyAttachmentDataPopulationStatusInProgress,
	"WAITING":             TenancyAttachmentDataPopulationStatusWaiting,
	"FAILED":              TenancyAttachmentDataPopulationStatusFailed,
	"SUCCEEDED":           TenancyAttachmentDataPopulationStatusSucceeded,
	"CANCELING":           TenancyAttachmentDataPopulationStatusCanceling,
	"CANCELED":            TenancyAttachmentDataPopulationStatusCanceled,
	"NEEDS_ATTENTION":     TenancyAttachmentDataPopulationStatusNeedsAttention,
	"PARTIALLY_SUCCEEDED": TenancyAttachmentDataPopulationStatusPartiallySucceeded,
}

var mappingTenancyAttachmentDataPopulationStatusEnumLowerCase = map[string]TenancyAttachmentDataPopulationStatusEnum{
	"accepted":            TenancyAttachmentDataPopulationStatusAccepted,
	"in_progress":         TenancyAttachmentDataPopulationStatusInProgress,
	"waiting":             TenancyAttachmentDataPopulationStatusWaiting,
	"failed":              TenancyAttachmentDataPopulationStatusFailed,
	"succeeded":           TenancyAttachmentDataPopulationStatusSucceeded,
	"canceling":           TenancyAttachmentDataPopulationStatusCanceling,
	"canceled":            TenancyAttachmentDataPopulationStatusCanceled,
	"needs_attention":     TenancyAttachmentDataPopulationStatusNeedsAttention,
	"partially_succeeded": TenancyAttachmentDataPopulationStatusPartiallySucceeded,
}

// GetTenancyAttachmentDataPopulationStatusEnumValues Enumerates the set of values for TenancyAttachmentDataPopulationStatusEnum
func GetTenancyAttachmentDataPopulationStatusEnumValues() []TenancyAttachmentDataPopulationStatusEnum {
	values := make([]TenancyAttachmentDataPopulationStatusEnum, 0)
	for _, v := range mappingTenancyAttachmentDataPopulationStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetTenancyAttachmentDataPopulationStatusEnumStringValues Enumerates the set of values in String for TenancyAttachmentDataPopulationStatusEnum
func GetTenancyAttachmentDataPopulationStatusEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELING",
		"CANCELED",
		"NEEDS_ATTENTION",
		"PARTIALLY_SUCCEEDED",
	}
}

// GetMappingTenancyAttachmentDataPopulationStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTenancyAttachmentDataPopulationStatusEnum(val string) (TenancyAttachmentDataPopulationStatusEnum, bool) {
	enum, ok := mappingTenancyAttachmentDataPopulationStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// TenancyAttachmentLifecycleStateEnum Enum with underlying type: string
type TenancyAttachmentLifecycleStateEnum string

// Set of constants representing the allowable values for TenancyAttachmentLifecycleStateEnum
const (
	TenancyAttachmentLifecycleStateCreating       TenancyAttachmentLifecycleStateEnum = "CREATING"
	TenancyAttachmentLifecycleStateUpdating       TenancyAttachmentLifecycleStateEnum = "UPDATING"
	TenancyAttachmentLifecycleStateActive         TenancyAttachmentLifecycleStateEnum = "ACTIVE"
	TenancyAttachmentLifecycleStateNeedsAttention TenancyAttachmentLifecycleStateEnum = "NEEDS_ATTENTION"
	TenancyAttachmentLifecycleStateDeleting       TenancyAttachmentLifecycleStateEnum = "DELETING"
	TenancyAttachmentLifecycleStateDeleted        TenancyAttachmentLifecycleStateEnum = "DELETED"
	TenancyAttachmentLifecycleStateFailed         TenancyAttachmentLifecycleStateEnum = "FAILED"
)

var mappingTenancyAttachmentLifecycleStateEnum = map[string]TenancyAttachmentLifecycleStateEnum{
	"CREATING":        TenancyAttachmentLifecycleStateCreating,
	"UPDATING":        TenancyAttachmentLifecycleStateUpdating,
	"ACTIVE":          TenancyAttachmentLifecycleStateActive,
	"NEEDS_ATTENTION": TenancyAttachmentLifecycleStateNeedsAttention,
	"DELETING":        TenancyAttachmentLifecycleStateDeleting,
	"DELETED":         TenancyAttachmentLifecycleStateDeleted,
	"FAILED":          TenancyAttachmentLifecycleStateFailed,
}

var mappingTenancyAttachmentLifecycleStateEnumLowerCase = map[string]TenancyAttachmentLifecycleStateEnum{
	"creating":        TenancyAttachmentLifecycleStateCreating,
	"updating":        TenancyAttachmentLifecycleStateUpdating,
	"active":          TenancyAttachmentLifecycleStateActive,
	"needs_attention": TenancyAttachmentLifecycleStateNeedsAttention,
	"deleting":        TenancyAttachmentLifecycleStateDeleting,
	"deleted":         TenancyAttachmentLifecycleStateDeleted,
	"failed":          TenancyAttachmentLifecycleStateFailed,
}

// GetTenancyAttachmentLifecycleStateEnumValues Enumerates the set of values for TenancyAttachmentLifecycleStateEnum
func GetTenancyAttachmentLifecycleStateEnumValues() []TenancyAttachmentLifecycleStateEnum {
	values := make([]TenancyAttachmentLifecycleStateEnum, 0)
	for _, v := range mappingTenancyAttachmentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetTenancyAttachmentLifecycleStateEnumStringValues Enumerates the set of values in String for TenancyAttachmentLifecycleStateEnum
func GetTenancyAttachmentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"NEEDS_ATTENTION",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingTenancyAttachmentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTenancyAttachmentLifecycleStateEnum(val string) (TenancyAttachmentLifecycleStateEnum, bool) {
	enum, ok := mappingTenancyAttachmentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
