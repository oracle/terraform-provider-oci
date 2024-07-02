// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.cloud.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOccCapacityRequestDetails The details required for making an update call for capacity requests.
type UpdateOccCapacityRequestDetails struct {

	// The display name of the capacity request.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The subset of request states available for updating the capacity request.
	RequestState UpdateOccCapacityRequestDetailsRequestStateEnum `mandatory:"false" json:"requestState,omitempty"`
}

func (m UpdateOccCapacityRequestDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOccCapacityRequestDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateOccCapacityRequestDetailsRequestStateEnum(string(m.RequestState)); !ok && m.RequestState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RequestState: %s. Supported values are: %s.", m.RequestState, strings.Join(GetUpdateOccCapacityRequestDetailsRequestStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateOccCapacityRequestDetailsRequestStateEnum Enum with underlying type: string
type UpdateOccCapacityRequestDetailsRequestStateEnum string

// Set of constants representing the allowable values for UpdateOccCapacityRequestDetailsRequestStateEnum
const (
	UpdateOccCapacityRequestDetailsRequestStateSubmitted UpdateOccCapacityRequestDetailsRequestStateEnum = "SUBMITTED"
	UpdateOccCapacityRequestDetailsRequestStateCancelled UpdateOccCapacityRequestDetailsRequestStateEnum = "CANCELLED"
)

var mappingUpdateOccCapacityRequestDetailsRequestStateEnum = map[string]UpdateOccCapacityRequestDetailsRequestStateEnum{
	"SUBMITTED": UpdateOccCapacityRequestDetailsRequestStateSubmitted,
	"CANCELLED": UpdateOccCapacityRequestDetailsRequestStateCancelled,
}

var mappingUpdateOccCapacityRequestDetailsRequestStateEnumLowerCase = map[string]UpdateOccCapacityRequestDetailsRequestStateEnum{
	"submitted": UpdateOccCapacityRequestDetailsRequestStateSubmitted,
	"cancelled": UpdateOccCapacityRequestDetailsRequestStateCancelled,
}

// GetUpdateOccCapacityRequestDetailsRequestStateEnumValues Enumerates the set of values for UpdateOccCapacityRequestDetailsRequestStateEnum
func GetUpdateOccCapacityRequestDetailsRequestStateEnumValues() []UpdateOccCapacityRequestDetailsRequestStateEnum {
	values := make([]UpdateOccCapacityRequestDetailsRequestStateEnum, 0)
	for _, v := range mappingUpdateOccCapacityRequestDetailsRequestStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateOccCapacityRequestDetailsRequestStateEnumStringValues Enumerates the set of values in String for UpdateOccCapacityRequestDetailsRequestStateEnum
func GetUpdateOccCapacityRequestDetailsRequestStateEnumStringValues() []string {
	return []string{
		"SUBMITTED",
		"CANCELLED",
	}
}

// GetMappingUpdateOccCapacityRequestDetailsRequestStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateOccCapacityRequestDetailsRequestStateEnum(val string) (UpdateOccCapacityRequestDetailsRequestStateEnum, bool) {
	enum, ok := mappingUpdateOccCapacityRequestDetailsRequestStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
