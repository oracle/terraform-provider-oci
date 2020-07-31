// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WorkRequestSummary A description of workrequest status
type WorkRequestSummary struct {

	// type of the work request
	OperationType WorkRequestSummaryOperationTypeEnum `mandatory:"true" json:"operationType"`

	// The id of the work request.
	Id *string `mandatory:"true" json:"id"`

	// The ocid of the compartment that contains the work request. Work requests should be scoped to
	// the same compartment as the resource the work request affects. If the work request affects multiple resources,
	// and those resources are not in the same compartment, it is up to the service team to pick the primary
	// resource whose compartment should be used
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The resources affected by this work request.
	Resources []WorkRequestResource `mandatory:"true" json:"resources"`

	// Percentage of the request completed.
	PercentComplete *float32 `mandatory:"true" json:"percentComplete"`

	// The date and time the request was created, as described in
	// RFC 3339 (https://tools.ietf.org/rfc/rfc3339), section 14.29.
	TimeAccepted *common.SDKTime `mandatory:"true" json:"timeAccepted"`

	// status of current work request.
	Status WorkRequestStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The date and time the request was started, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339),
	// section 14.29.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the object was finished, as described in RFC 3339 (https://tools.ietf.org/rfc/rfc3339).
	TimeFinished *common.SDKTime `mandatory:"false" json:"timeFinished"`
}

func (m WorkRequestSummary) String() string {
	return common.PointerString(m)
}

// WorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type WorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestSummaryOperationTypeEnum
const (
	WorkRequestSummaryOperationTypeCreatePlatform    WorkRequestSummaryOperationTypeEnum = "CREATE_PLATFORM"
	WorkRequestSummaryOperationTypeUpdatePlatform    WorkRequestSummaryOperationTypeEnum = "UPDATE_PLATFORM"
	WorkRequestSummaryOperationTypeDeletePlatform    WorkRequestSummaryOperationTypeEnum = "DELETE_PLATFORM"
	WorkRequestSummaryOperationTypeScalePlatform     WorkRequestSummaryOperationTypeEnum = "SCALE_PLATFORM"
	WorkRequestSummaryOperationTypeStartPlatform     WorkRequestSummaryOperationTypeEnum = "START_PLATFORM"
	WorkRequestSummaryOperationTypeStopPlatform      WorkRequestSummaryOperationTypeEnum = "STOP_PLATFORM"
	WorkRequestSummaryOperationTypeCustomizePlatform WorkRequestSummaryOperationTypeEnum = "CUSTOMIZE_PLATFORM"
)

var mappingWorkRequestSummaryOperationType = map[string]WorkRequestSummaryOperationTypeEnum{
	"CREATE_PLATFORM":    WorkRequestSummaryOperationTypeCreatePlatform,
	"UPDATE_PLATFORM":    WorkRequestSummaryOperationTypeUpdatePlatform,
	"DELETE_PLATFORM":    WorkRequestSummaryOperationTypeDeletePlatform,
	"SCALE_PLATFORM":     WorkRequestSummaryOperationTypeScalePlatform,
	"START_PLATFORM":     WorkRequestSummaryOperationTypeStartPlatform,
	"STOP_PLATFORM":      WorkRequestSummaryOperationTypeStopPlatform,
	"CUSTOMIZE_PLATFORM": WorkRequestSummaryOperationTypeCustomizePlatform,
}

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumValues() []WorkRequestSummaryOperationTypeEnum {
	values := make([]WorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestSummaryOperationType {
		values = append(values, v)
	}
	return values
}
