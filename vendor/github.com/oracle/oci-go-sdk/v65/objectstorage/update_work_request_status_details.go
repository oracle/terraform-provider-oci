// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.cloud.oracle.com/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.cloud.oracle.com/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateWorkRequestStatusDetails The status of the specified work request.
type UpdateWorkRequestStatusDetails struct {

	// The status of the specified work request.
	State UpdateWorkRequestStatusDetailsStateEnum `mandatory:"false" json:"state,omitempty"`

	// Percentage of the work request completed.
	PercentComplete *float32 `mandatory:"false" json:"percentComplete"`

	// list of work request logs.
	WorkRequestLogs []WorkRequestLogEntry `mandatory:"false" json:"workRequestLogs"`

	// list of work request errors.
	WorkRequestErrors []WorkRequestError `mandatory:"false" json:"workRequestErrors"`
}

func (m UpdateWorkRequestStatusDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateWorkRequestStatusDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateWorkRequestStatusDetailsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetUpdateWorkRequestStatusDetailsStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateWorkRequestStatusDetailsStateEnum Enum with underlying type: string
type UpdateWorkRequestStatusDetailsStateEnum string

// Set of constants representing the allowable values for UpdateWorkRequestStatusDetailsStateEnum
const (
	UpdateWorkRequestStatusDetailsStateAccepted   UpdateWorkRequestStatusDetailsStateEnum = "ACCEPTED"
	UpdateWorkRequestStatusDetailsStateInProgress UpdateWorkRequestStatusDetailsStateEnum = "IN_PROGRESS"
	UpdateWorkRequestStatusDetailsStateFailed     UpdateWorkRequestStatusDetailsStateEnum = "FAILED"
	UpdateWorkRequestStatusDetailsStateCompleted  UpdateWorkRequestStatusDetailsStateEnum = "COMPLETED"
	UpdateWorkRequestStatusDetailsStateCanceling  UpdateWorkRequestStatusDetailsStateEnum = "CANCELING"
	UpdateWorkRequestStatusDetailsStateCanceled   UpdateWorkRequestStatusDetailsStateEnum = "CANCELED"
)

var mappingUpdateWorkRequestStatusDetailsStateEnum = map[string]UpdateWorkRequestStatusDetailsStateEnum{
	"ACCEPTED":    UpdateWorkRequestStatusDetailsStateAccepted,
	"IN_PROGRESS": UpdateWorkRequestStatusDetailsStateInProgress,
	"FAILED":      UpdateWorkRequestStatusDetailsStateFailed,
	"COMPLETED":   UpdateWorkRequestStatusDetailsStateCompleted,
	"CANCELING":   UpdateWorkRequestStatusDetailsStateCanceling,
	"CANCELED":    UpdateWorkRequestStatusDetailsStateCanceled,
}

var mappingUpdateWorkRequestStatusDetailsStateEnumLowerCase = map[string]UpdateWorkRequestStatusDetailsStateEnum{
	"accepted":    UpdateWorkRequestStatusDetailsStateAccepted,
	"in_progress": UpdateWorkRequestStatusDetailsStateInProgress,
	"failed":      UpdateWorkRequestStatusDetailsStateFailed,
	"completed":   UpdateWorkRequestStatusDetailsStateCompleted,
	"canceling":   UpdateWorkRequestStatusDetailsStateCanceling,
	"canceled":    UpdateWorkRequestStatusDetailsStateCanceled,
}

// GetUpdateWorkRequestStatusDetailsStateEnumValues Enumerates the set of values for UpdateWorkRequestStatusDetailsStateEnum
func GetUpdateWorkRequestStatusDetailsStateEnumValues() []UpdateWorkRequestStatusDetailsStateEnum {
	values := make([]UpdateWorkRequestStatusDetailsStateEnum, 0)
	for _, v := range mappingUpdateWorkRequestStatusDetailsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateWorkRequestStatusDetailsStateEnumStringValues Enumerates the set of values in String for UpdateWorkRequestStatusDetailsStateEnum
func GetUpdateWorkRequestStatusDetailsStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"FAILED",
		"COMPLETED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingUpdateWorkRequestStatusDetailsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateWorkRequestStatusDetailsStateEnum(val string) (UpdateWorkRequestStatusDetailsStateEnum, bool) {
	enum, ok := mappingUpdateWorkRequestStatusDetailsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
