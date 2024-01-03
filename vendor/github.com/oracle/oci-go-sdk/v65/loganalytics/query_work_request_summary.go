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

// QueryWorkRequestSummary High level summary of query job work request.
type QueryWorkRequestSummary struct {

	// Unique OCID identifier to reference this query job work Request with.
	Id *string `mandatory:"true" json:"id"`

	// When the work request started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// Current execution mode for the job.
	Mode JobModeEnum `mandatory:"true" json:"mode"`

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// When the work request was accepted. Should match timeStarted in all cases.
	TimeAccepted *common.SDKTime `mandatory:"false" json:"timeAccepted"`

	// When the work request finished execution.
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`

	// When the work request will expire.
	TimeExpires *common.SDKTime `mandatory:"false" json:"timeExpires"`

	// Percentage progress completion of the query.
	PercentComplete *int `mandatory:"false" json:"percentComplete"`

	// Work request status.
	Status WorkRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Asynchronous action name.
	OperationType QueryOperationTypeEnum `mandatory:"false" json:"operationType,omitempty"`
}

func (m QueryWorkRequestSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryWorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingJobModeEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetJobModeEnumStringValues(), ",")))
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
