// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WorkRequestSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingWorkRequestSummaryOperationTypeEnum(string(m.OperationType)); !ok && m.OperationType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OperationType: %s. Supported values are: %s.", m.OperationType, strings.Join(GetWorkRequestSummaryOperationTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingWorkRequestStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetWorkRequestStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WorkRequestSummaryOperationTypeEnum Enum with underlying type: string
type WorkRequestSummaryOperationTypeEnum string

// Set of constants representing the allowable values for WorkRequestSummaryOperationTypeEnum
const (
	WorkRequestSummaryOperationTypeCreatePlatform    WorkRequestSummaryOperationTypeEnum = "CREATE_PLATFORM"
	WorkRequestSummaryOperationTypeUpdatePlatform    WorkRequestSummaryOperationTypeEnum = "UPDATE_PLATFORM"
	WorkRequestSummaryOperationTypeUpgradePlatform   WorkRequestSummaryOperationTypeEnum = "UPGRADE_PLATFORM"
	WorkRequestSummaryOperationTypeDeletePlatform    WorkRequestSummaryOperationTypeEnum = "DELETE_PLATFORM"
	WorkRequestSummaryOperationTypeScalePlatform     WorkRequestSummaryOperationTypeEnum = "SCALE_PLATFORM"
	WorkRequestSummaryOperationTypeStartPlatform     WorkRequestSummaryOperationTypeEnum = "START_PLATFORM"
	WorkRequestSummaryOperationTypeStopPlatform      WorkRequestSummaryOperationTypeEnum = "STOP_PLATFORM"
	WorkRequestSummaryOperationTypeCustomizePlatform WorkRequestSummaryOperationTypeEnum = "CUSTOMIZE_PLATFORM"
	WorkRequestSummaryOperationTypeScaleStorage      WorkRequestSummaryOperationTypeEnum = "SCALE_STORAGE"
)

var mappingWorkRequestSummaryOperationTypeEnum = map[string]WorkRequestSummaryOperationTypeEnum{
	"CREATE_PLATFORM":    WorkRequestSummaryOperationTypeCreatePlatform,
	"UPDATE_PLATFORM":    WorkRequestSummaryOperationTypeUpdatePlatform,
	"UPGRADE_PLATFORM":   WorkRequestSummaryOperationTypeUpgradePlatform,
	"DELETE_PLATFORM":    WorkRequestSummaryOperationTypeDeletePlatform,
	"SCALE_PLATFORM":     WorkRequestSummaryOperationTypeScalePlatform,
	"START_PLATFORM":     WorkRequestSummaryOperationTypeStartPlatform,
	"STOP_PLATFORM":      WorkRequestSummaryOperationTypeStopPlatform,
	"CUSTOMIZE_PLATFORM": WorkRequestSummaryOperationTypeCustomizePlatform,
	"SCALE_STORAGE":      WorkRequestSummaryOperationTypeScaleStorage,
}

var mappingWorkRequestSummaryOperationTypeEnumLowerCase = map[string]WorkRequestSummaryOperationTypeEnum{
	"create_platform":    WorkRequestSummaryOperationTypeCreatePlatform,
	"update_platform":    WorkRequestSummaryOperationTypeUpdatePlatform,
	"upgrade_platform":   WorkRequestSummaryOperationTypeUpgradePlatform,
	"delete_platform":    WorkRequestSummaryOperationTypeDeletePlatform,
	"scale_platform":     WorkRequestSummaryOperationTypeScalePlatform,
	"start_platform":     WorkRequestSummaryOperationTypeStartPlatform,
	"stop_platform":      WorkRequestSummaryOperationTypeStopPlatform,
	"customize_platform": WorkRequestSummaryOperationTypeCustomizePlatform,
	"scale_storage":      WorkRequestSummaryOperationTypeScaleStorage,
}

// GetWorkRequestSummaryOperationTypeEnumValues Enumerates the set of values for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumValues() []WorkRequestSummaryOperationTypeEnum {
	values := make([]WorkRequestSummaryOperationTypeEnum, 0)
	for _, v := range mappingWorkRequestSummaryOperationTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWorkRequestSummaryOperationTypeEnumStringValues Enumerates the set of values in String for WorkRequestSummaryOperationTypeEnum
func GetWorkRequestSummaryOperationTypeEnumStringValues() []string {
	return []string{
		"CREATE_PLATFORM",
		"UPDATE_PLATFORM",
		"UPGRADE_PLATFORM",
		"DELETE_PLATFORM",
		"SCALE_PLATFORM",
		"START_PLATFORM",
		"STOP_PLATFORM",
		"CUSTOMIZE_PLATFORM",
		"SCALE_STORAGE",
	}
}

// GetMappingWorkRequestSummaryOperationTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWorkRequestSummaryOperationTypeEnum(val string) (WorkRequestSummaryOperationTypeEnum, bool) {
	enum, ok := mappingWorkRequestSummaryOperationTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
