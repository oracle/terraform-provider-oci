// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubsettingReportSummary Summary information for a report generated from a data subsetting operation
type SubsettingReportSummary struct {

	// The OCID of the subsetting report
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that contains the subsetting report
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the subsetting work request that resulted in this subsetting report
	SubsettingWorkRequestId *string `mandatory:"true" json:"subsettingWorkRequestId"`

	// The OCID of the subsetting policy used
	SubsettingPolicyId *string `mandatory:"true" json:"subsettingPolicyId"`

	// The OCID of the target database subsetted
	TargetId *string `mandatory:"true" json:"targetId"`

	// The total number of subsetted objects
	TotalSubsettedObjects *int64 `mandatory:"true" json:"totalSubsettedObjects"`

	// The date and time data subsetting started, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeSubsettingStarted *common.SDKTime `mandatory:"true" json:"timeSubsettingStarted"`

	// The date and time data subsetting finished, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeSubsettingFinished *common.SDKTime `mandatory:"true" json:"timeSubsettingFinished"`

	// The current state of the subsetting report
	LifecycleState SubsettingReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The status of the subsetting job
	SubsettingStatus SubsettingReportSummarySubsettingStatusEnum `mandatory:"true" json:"subsettingStatus"`

	// The OCID of the masking report associated with this subsetting report
	MaskingReportId *string `mandatory:"false" json:"maskingReportId"`

	// The OCID of the masking policy associated with this subsetting policy
	MaskingPolicyId *string `mandatory:"false" json:"maskingPolicyId"`

	// The OCID of the masking work request triggered after this subsetting job
	MaskingWorkRequestId *string `mandatory:"false" json:"maskingWorkRequestId"`

	// The total number of subsetted schemas
	TotalSubsettedSchemas *int64 `mandatory:"false" json:"totalSubsettedSchemas"`

	// The count of rows reduced in the subsetting job
	TotalSubsettedRows *int64 `mandatory:"false" json:"totalSubsettedRows"`

	// The size of the target database before subsetting in KBs
	DatabaseSizeBeforeSubsettingInKBs *string `mandatory:"false" json:"databaseSizeBeforeSubsettingInKBs"`

	// The size of the target database after subsetting in KBs
	DatabaseSizeAfterSubsettingInKBs *string `mandatory:"false" json:"databaseSizeAfterSubsettingInKBs"`

	// The date and time the subsetting report was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339)
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Indicates if redo logging was enabled during the subsetting operation
	IsRedoLoggingEnabled *bool `mandatory:"false" json:"isRedoLoggingEnabled"`

	// Indicates if statistics gathering was enabled during the subsetting operation
	IsRefreshStatsEnabled *bool `mandatory:"false" json:"isRefreshStatsEnabled"`

	// Indicates if parallel execution was enabled during the subsetting operation
	ParallelDegree *string `mandatory:"false" json:"parallelDegree"`

	// Indicates how invalid objects were recompiled post the subsetting operation
	Recompile *string `mandatory:"false" json:"recompile"`

	// The total number of errors in pre-subsetting script
	TotalPreSubsettingScriptErrors *int64 `mandatory:"false" json:"totalPreSubsettingScriptErrors"`

	// The total number of errors in post-subsetting script
	TotalPostSubsettingScriptErrors *int64 `mandatory:"false" json:"totalPostSubsettingScriptErrors"`
}

func (m SubsettingReportSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubsettingReportSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubsettingReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetSubsettingReportLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubsettingReportSummarySubsettingStatusEnum(string(m.SubsettingStatus)); !ok && m.SubsettingStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubsettingStatus: %s. Supported values are: %s.", m.SubsettingStatus, strings.Join(GetSubsettingReportSummarySubsettingStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SubsettingReportSummarySubsettingStatusEnum Enum with underlying type: string
type SubsettingReportSummarySubsettingStatusEnum string

// Set of constants representing the allowable values for SubsettingReportSummarySubsettingStatusEnum
const (
	SubsettingReportSummarySubsettingStatusFailed  SubsettingReportSummarySubsettingStatusEnum = "FAILED"
	SubsettingReportSummarySubsettingStatusSuccess SubsettingReportSummarySubsettingStatusEnum = "SUCCESS"
)

var mappingSubsettingReportSummarySubsettingStatusEnum = map[string]SubsettingReportSummarySubsettingStatusEnum{
	"FAILED":  SubsettingReportSummarySubsettingStatusFailed,
	"SUCCESS": SubsettingReportSummarySubsettingStatusSuccess,
}

var mappingSubsettingReportSummarySubsettingStatusEnumLowerCase = map[string]SubsettingReportSummarySubsettingStatusEnum{
	"failed":  SubsettingReportSummarySubsettingStatusFailed,
	"success": SubsettingReportSummarySubsettingStatusSuccess,
}

// GetSubsettingReportSummarySubsettingStatusEnumValues Enumerates the set of values for SubsettingReportSummarySubsettingStatusEnum
func GetSubsettingReportSummarySubsettingStatusEnumValues() []SubsettingReportSummarySubsettingStatusEnum {
	values := make([]SubsettingReportSummarySubsettingStatusEnum, 0)
	for _, v := range mappingSubsettingReportSummarySubsettingStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetSubsettingReportSummarySubsettingStatusEnumStringValues Enumerates the set of values in String for SubsettingReportSummarySubsettingStatusEnum
func GetSubsettingReportSummarySubsettingStatusEnumStringValues() []string {
	return []string{
		"FAILED",
		"SUCCESS",
	}
}

// GetMappingSubsettingReportSummarySubsettingStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSubsettingReportSummarySubsettingStatusEnum(val string) (SubsettingReportSummarySubsettingStatusEnum, bool) {
	enum, ok := mappingSubsettingReportSummarySubsettingStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
