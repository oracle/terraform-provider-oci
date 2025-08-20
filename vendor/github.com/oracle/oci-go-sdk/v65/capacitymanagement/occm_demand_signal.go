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

// OccmDemandSignal An occm demand signal is a resource that communicates the forecasting needs of a customer to OCI in advance.
type OccmDemandSignal struct {

	// The OCID of the demand signal.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the tenancy from which the request to create the demand signal was made.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The different states associated with a demand signal.
	// CREATED -> A demand signal is by default created in this state.
	// SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that OCI can start working on it.
	// DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state.
	// IN_PROGRESS -> Once OCI starts working on a given demand signal. They transition it to IN_PROGRESS.
	// CANCELLED -> OCI can transition the demand signal to this state.
	// COMPLETED -> OCI will transition the demand signal to COMPLETED state once the quantities which OCI committed to deliver to you has been delivered.
	LifecycleDetails OccmDemandSignalLifecycleDetailsEnum `mandatory:"true" json:"lifecycleDetails"`

	// The display name of the demand signal.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The current lifecycle state of the resource.
	LifecycleState OccmDemandSignalLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The time when the demand signal was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time when the demand signal was last updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Meaningful text about the demand signal.
	Description *string `mandatory:"false" json:"description"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OccmDemandSignal) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OccmDemandSignal) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOccmDemandSignalLifecycleDetailsEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetOccmDemandSignalLifecycleDetailsEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOccmDemandSignalLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOccmDemandSignalLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OccmDemandSignalLifecycleDetailsEnum Enum with underlying type: string
type OccmDemandSignalLifecycleDetailsEnum string

// Set of constants representing the allowable values for OccmDemandSignalLifecycleDetailsEnum
const (
	OccmDemandSignalLifecycleDetailsCreated    OccmDemandSignalLifecycleDetailsEnum = "CREATED"
	OccmDemandSignalLifecycleDetailsSubmitted  OccmDemandSignalLifecycleDetailsEnum = "SUBMITTED"
	OccmDemandSignalLifecycleDetailsDeleted    OccmDemandSignalLifecycleDetailsEnum = "DELETED"
	OccmDemandSignalLifecycleDetailsInProgress OccmDemandSignalLifecycleDetailsEnum = "IN_PROGRESS"
	OccmDemandSignalLifecycleDetailsRejected   OccmDemandSignalLifecycleDetailsEnum = "REJECTED"
	OccmDemandSignalLifecycleDetailsCompleted  OccmDemandSignalLifecycleDetailsEnum = "COMPLETED"
)

var mappingOccmDemandSignalLifecycleDetailsEnum = map[string]OccmDemandSignalLifecycleDetailsEnum{
	"CREATED":     OccmDemandSignalLifecycleDetailsCreated,
	"SUBMITTED":   OccmDemandSignalLifecycleDetailsSubmitted,
	"DELETED":     OccmDemandSignalLifecycleDetailsDeleted,
	"IN_PROGRESS": OccmDemandSignalLifecycleDetailsInProgress,
	"REJECTED":    OccmDemandSignalLifecycleDetailsRejected,
	"COMPLETED":   OccmDemandSignalLifecycleDetailsCompleted,
}

var mappingOccmDemandSignalLifecycleDetailsEnumLowerCase = map[string]OccmDemandSignalLifecycleDetailsEnum{
	"created":     OccmDemandSignalLifecycleDetailsCreated,
	"submitted":   OccmDemandSignalLifecycleDetailsSubmitted,
	"deleted":     OccmDemandSignalLifecycleDetailsDeleted,
	"in_progress": OccmDemandSignalLifecycleDetailsInProgress,
	"rejected":    OccmDemandSignalLifecycleDetailsRejected,
	"completed":   OccmDemandSignalLifecycleDetailsCompleted,
}

// GetOccmDemandSignalLifecycleDetailsEnumValues Enumerates the set of values for OccmDemandSignalLifecycleDetailsEnum
func GetOccmDemandSignalLifecycleDetailsEnumValues() []OccmDemandSignalLifecycleDetailsEnum {
	values := make([]OccmDemandSignalLifecycleDetailsEnum, 0)
	for _, v := range mappingOccmDemandSignalLifecycleDetailsEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalLifecycleDetailsEnumStringValues Enumerates the set of values in String for OccmDemandSignalLifecycleDetailsEnum
func GetOccmDemandSignalLifecycleDetailsEnumStringValues() []string {
	return []string{
		"CREATED",
		"SUBMITTED",
		"DELETED",
		"IN_PROGRESS",
		"REJECTED",
		"COMPLETED",
	}
}

// GetMappingOccmDemandSignalLifecycleDetailsEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalLifecycleDetailsEnum(val string) (OccmDemandSignalLifecycleDetailsEnum, bool) {
	enum, ok := mappingOccmDemandSignalLifecycleDetailsEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OccmDemandSignalLifecycleStateEnum Enum with underlying type: string
type OccmDemandSignalLifecycleStateEnum string

// Set of constants representing the allowable values for OccmDemandSignalLifecycleStateEnum
const (
	OccmDemandSignalLifecycleStateCreating OccmDemandSignalLifecycleStateEnum = "CREATING"
	OccmDemandSignalLifecycleStateActive   OccmDemandSignalLifecycleStateEnum = "ACTIVE"
	OccmDemandSignalLifecycleStateUpdating OccmDemandSignalLifecycleStateEnum = "UPDATING"
	OccmDemandSignalLifecycleStateDeleted  OccmDemandSignalLifecycleStateEnum = "DELETED"
	OccmDemandSignalLifecycleStateDeleting OccmDemandSignalLifecycleStateEnum = "DELETING"
	OccmDemandSignalLifecycleStateFailed   OccmDemandSignalLifecycleStateEnum = "FAILED"
)

var mappingOccmDemandSignalLifecycleStateEnum = map[string]OccmDemandSignalLifecycleStateEnum{
	"CREATING": OccmDemandSignalLifecycleStateCreating,
	"ACTIVE":   OccmDemandSignalLifecycleStateActive,
	"UPDATING": OccmDemandSignalLifecycleStateUpdating,
	"DELETED":  OccmDemandSignalLifecycleStateDeleted,
	"DELETING": OccmDemandSignalLifecycleStateDeleting,
	"FAILED":   OccmDemandSignalLifecycleStateFailed,
}

var mappingOccmDemandSignalLifecycleStateEnumLowerCase = map[string]OccmDemandSignalLifecycleStateEnum{
	"creating": OccmDemandSignalLifecycleStateCreating,
	"active":   OccmDemandSignalLifecycleStateActive,
	"updating": OccmDemandSignalLifecycleStateUpdating,
	"deleted":  OccmDemandSignalLifecycleStateDeleted,
	"deleting": OccmDemandSignalLifecycleStateDeleting,
	"failed":   OccmDemandSignalLifecycleStateFailed,
}

// GetOccmDemandSignalLifecycleStateEnumValues Enumerates the set of values for OccmDemandSignalLifecycleStateEnum
func GetOccmDemandSignalLifecycleStateEnumValues() []OccmDemandSignalLifecycleStateEnum {
	values := make([]OccmDemandSignalLifecycleStateEnum, 0)
	for _, v := range mappingOccmDemandSignalLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOccmDemandSignalLifecycleStateEnumStringValues Enumerates the set of values in String for OccmDemandSignalLifecycleStateEnum
func GetOccmDemandSignalLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"UPDATING",
		"DELETED",
		"DELETING",
		"FAILED",
	}
}

// GetMappingOccmDemandSignalLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOccmDemandSignalLifecycleStateEnum(val string) (OccmDemandSignalLifecycleStateEnum, bool) {
	enum, ok := mappingOccmDemandSignalLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
