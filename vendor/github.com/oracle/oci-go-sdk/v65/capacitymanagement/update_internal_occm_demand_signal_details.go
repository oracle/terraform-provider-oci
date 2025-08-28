// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OCI Control Center Capacity Management API
//
// OCI Control Center (OCC) Capacity Management enables you to manage capacity requests in realms where OCI Control Center Capacity Management is available. For more information, see OCI Control Center (https://docs.oracle.com/iaas/Content/control-center/home.htm).
//

package capacitymanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInternalOccmDemandSignalDetails An internal model to update the demand signal state.
type UpdateInternalOccmDemandSignalDetails struct {

	// The subset of demand signal states available for operators for updating the demand signal.
	// IN_PROGRESS -> Transitions the demand signal to IN_PROGRESS state.
	// REJECTED -> Transitions the demand signal to REJECTED state.
	// COMPLETED -> This will transition the demand signal to COMPLETED state.
	LifecycleDetails UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m UpdateInternalOccmDemandSignalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalOccmDemandSignalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum Enum with underlying type: string
type UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum string

// Set of constants representing the allowable values for UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum
const (
	UpdateInternalOccmDemandSignalDetailsLifecycleDetailsInProgress UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum = "IN_PROGRESS"
	UpdateInternalOccmDemandSignalDetailsLifecycleDetailsCompleted  UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum = "COMPLETED"
	UpdateInternalOccmDemandSignalDetailsLifecycleDetailsRejected   UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum = "REJECTED"
)

var mappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum = map[string]UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum{
	"IN_PROGRESS": UpdateInternalOccmDemandSignalDetailsLifecycleDetailsInProgress,
	"COMPLETED":   UpdateInternalOccmDemandSignalDetailsLifecycleDetailsCompleted,
	"REJECTED":    UpdateInternalOccmDemandSignalDetailsLifecycleDetailsRejected,
}

var mappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumLowerCase = map[string]UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum{
	"in_progress": UpdateInternalOccmDemandSignalDetailsLifecycleDetailsInProgress,
	"completed":   UpdateInternalOccmDemandSignalDetailsLifecycleDetailsCompleted,
	"rejected":    UpdateInternalOccmDemandSignalDetailsLifecycleDetailsRejected,
}

// GetUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumValues Enumerates the set of values for UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum
func GetUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumValues() []UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum {
	values := make([]UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum, 0)
	for _, v := range mappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumStringValues Enumerates the set of values in String for UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum
func GetUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumStringValues() []string {
	return []string{
		"IN_PROGRESS",
		"COMPLETED",
		"REJECTED",
	}
}

// GetMappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum(val string) (UpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnum, bool) {
	enum, ok := mappingUpdateInternalOccmDemandSignalDetailsLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
