// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OciControlCenterCp API
//
// A description of the OciControlCenterCp API
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInternalOccCapacityRequestDetails The details required for making an internal API update call for the capacity requests.
type UpdateInternalOccCapacityRequestDetails struct {

	// The subset of request states available internally for updating the capacity request.
	RequestState UpdateInternalOccCapacityRequestDetailsRequestStateEnum `mandatory:"false" json:"requestState,omitempty"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m UpdateInternalOccCapacityRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalOccCapacityRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalOccCapacityRequestDetailsRequestStateEnum(string(m.RequestState)); !ok && m.RequestState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestState: %s. Supported values are: %s.", m.RequestState, strings.Join(GetUpdateInternalOccCapacityRequestDetailsRequestStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalOccCapacityRequestDetailsRequestStateEnum Enum with underlying type: string
type UpdateInternalOccCapacityRequestDetailsRequestStateEnum string

// Set of constants representing the allowable values for UpdateInternalOccCapacityRequestDetailsRequestStateEnum
const (
	UpdateInternalOccCapacityRequestDetailsRequestStateResolved   UpdateInternalOccCapacityRequestDetailsRequestStateEnum = "RESOLVED"
	UpdateInternalOccCapacityRequestDetailsRequestStateRejected   UpdateInternalOccCapacityRequestDetailsRequestStateEnum = "REJECTED"
	UpdateInternalOccCapacityRequestDetailsRequestStateInProgress UpdateInternalOccCapacityRequestDetailsRequestStateEnum = "IN_PROGRESS"
)

var mappingUpdateInternalOccCapacityRequestDetailsRequestStateEnum = map[string]UpdateInternalOccCapacityRequestDetailsRequestStateEnum{
	"RESOLVED":    UpdateInternalOccCapacityRequestDetailsRequestStateResolved,
	"REJECTED":    UpdateInternalOccCapacityRequestDetailsRequestStateRejected,
	"IN_PROGRESS": UpdateInternalOccCapacityRequestDetailsRequestStateInProgress,
}

var mappingUpdateInternalOccCapacityRequestDetailsRequestStateEnumLowerCase = map[string]UpdateInternalOccCapacityRequestDetailsRequestStateEnum{
	"resolved":    UpdateInternalOccCapacityRequestDetailsRequestStateResolved,
	"rejected":    UpdateInternalOccCapacityRequestDetailsRequestStateRejected,
	"in_progress": UpdateInternalOccCapacityRequestDetailsRequestStateInProgress,
}

// GetUpdateInternalOccCapacityRequestDetailsRequestStateEnumValues Enumerates the set of values for UpdateInternalOccCapacityRequestDetailsRequestStateEnum
func GetUpdateInternalOccCapacityRequestDetailsRequestStateEnumValues() []UpdateInternalOccCapacityRequestDetailsRequestStateEnum {
	values := make([]UpdateInternalOccCapacityRequestDetailsRequestStateEnum, 0)
	for _, v := range mappingUpdateInternalOccCapacityRequestDetailsRequestStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalOccCapacityRequestDetailsRequestStateEnumStringValues Enumerates the set of values in String for UpdateInternalOccCapacityRequestDetailsRequestStateEnum
func GetUpdateInternalOccCapacityRequestDetailsRequestStateEnumStringValues() []string {
	return []string{
		"RESOLVED",
		"REJECTED",
		"IN_PROGRESS",
	}
}

// GetMappingUpdateInternalOccCapacityRequestDetailsRequestStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalOccCapacityRequestDetailsRequestStateEnum(val string) (UpdateInternalOccCapacityRequestDetailsRequestStateEnum, bool) {
	enum, ok := mappingUpdateInternalOccCapacityRequestDetailsRequestStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
