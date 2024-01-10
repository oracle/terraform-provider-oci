// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryWorkRequest Job details outlining parameters specified when job was submitted.
type QueryWorkRequest struct {

	// Unique OCID identifier to reference this query job work Request with.
	Id *string `mandatory:"true" json:"id"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// When the job was started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// Current execution mode for the job.
	Mode JobModeEnum `mandatory:"true" json:"mode"`

	// Default subsystem to qualify fields with in the queryString if not specified.
	SubSystem SubSystemNameEnum `mandatory:"true" json:"subSystem"`

	// Display version of the user speciified queryString.
	DisplayQueryString *string `mandatory:"true" json:"displayQueryString"`

	// Internal version of the user specified queryString.
	InternalQueryString *string `mandatory:"true" json:"internalQueryString"`

	// When the work request was accepted. Should match timeStarted in all cases.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// When the job finished execution.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// When the job will expire.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Percentage progress completion of the query.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Work request status.
	Status WorkRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Asynchronous action name.
	OperationType QueryOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`

	// When the job was put in to the background.
	TimeBackgroundAt *common.SDKTime `mandatory:"false" json:"timeBackgroundAt"`

	TimeFilter *TimeRange `mandatory:"false" json:"timeFilter"`

	// List of filters applied when the query executed.
	ScopeFilters []ScopeFilter `mandatory:"false" json:"scopeFilters"`
}

func (m QueryWorkRequest) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryWorkRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetJobModeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubSystemNameEnum(string(m.SubSystem)); !ok && m.SubSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubSystem: %s. Supported values are: %s.", m.SubSystem, strings.Join(GetSubSystemNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingQueryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetQueryOperationTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
