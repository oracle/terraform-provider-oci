// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FirmwareReport A FirmwareReport represents an asynchronous request to generate
// a firmware report for a tenancy or compartment and tracks the lifecycle
// of that report.
type FirmwareReport struct {

	// The availability domain the firmware report was requested for.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment to run the firmware report against.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the firmware report.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the firmware report.
	LifecycleState FirmwareReportLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the firmware report was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	ResultLocationDetails ResultLocationDetails `mandatory:"true" json:"resultLocationDetails"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The date and time the firmware report reached the FAILED or COMPLETE lifecycle states, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCompleted *common.SDKTime `mandatory:"false" json:"timeCompleted"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m FirmwareReport) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FirmwareReport) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFirmwareReportLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetFirmwareReportLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *FirmwareReport) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCompleted         *common.SDKTime                   `json:"timeCompleted"`
		FreeformTags          map[string]string                 `json:"freeformTags"`
		DefinedTags           map[string]map[string]interface{} `json:"definedTags"`
		AvailabilityDomain    *string                           `json:"availabilityDomain"`
		CompartmentId         *string                           `json:"compartmentId"`
		Id                    *string                           `json:"id"`
		LifecycleState        FirmwareReportLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated           *common.SDKTime                   `json:"timeCreated"`
		ResultLocationDetails resultlocationdetails             `json:"resultLocationDetails"`
		DisplayName           *string                           `json:"displayName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCompleted = model.TimeCompleted

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.AvailabilityDomain = model.AvailabilityDomain

	m.CompartmentId = model.CompartmentId

	m.Id = model.Id

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	nn, e = model.ResultLocationDetails.UnmarshalPolymorphicJSON(model.ResultLocationDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ResultLocationDetails = nn.(ResultLocationDetails)
	} else {
		m.ResultLocationDetails = nil
	}

	m.DisplayName = model.DisplayName

	return
}

// FirmwareReportLifecycleStateEnum Enum with underlying type: string
type FirmwareReportLifecycleStateEnum string

// Set of constants representing the allowable values for FirmwareReportLifecycleStateEnum
const (
	FirmwareReportLifecycleStateAccepted   FirmwareReportLifecycleStateEnum = "ACCEPTED"
	FirmwareReportLifecycleStateInProgress FirmwareReportLifecycleStateEnum = "IN_PROGRESS"
	FirmwareReportLifecycleStateWaiting    FirmwareReportLifecycleStateEnum = "WAITING"
	FirmwareReportLifecycleStateFailed     FirmwareReportLifecycleStateEnum = "FAILED"
	FirmwareReportLifecycleStateSucceeded  FirmwareReportLifecycleStateEnum = "SUCCEEDED"
	FirmwareReportLifecycleStateCancelling FirmwareReportLifecycleStateEnum = "CANCELLING"
	FirmwareReportLifecycleStateCancelled  FirmwareReportLifecycleStateEnum = "CANCELLED"
)

var mappingFirmwareReportLifecycleStateEnum = map[string]FirmwareReportLifecycleStateEnum{
	"ACCEPTED":    FirmwareReportLifecycleStateAccepted,
	"IN_PROGRESS": FirmwareReportLifecycleStateInProgress,
	"WAITING":     FirmwareReportLifecycleStateWaiting,
	"FAILED":      FirmwareReportLifecycleStateFailed,
	"SUCCEEDED":   FirmwareReportLifecycleStateSucceeded,
	"CANCELLING":  FirmwareReportLifecycleStateCancelling,
	"CANCELLED":   FirmwareReportLifecycleStateCancelled,
}

var mappingFirmwareReportLifecycleStateEnumLowerCase = map[string]FirmwareReportLifecycleStateEnum{
	"accepted":    FirmwareReportLifecycleStateAccepted,
	"in_progress": FirmwareReportLifecycleStateInProgress,
	"waiting":     FirmwareReportLifecycleStateWaiting,
	"failed":      FirmwareReportLifecycleStateFailed,
	"succeeded":   FirmwareReportLifecycleStateSucceeded,
	"cancelling":  FirmwareReportLifecycleStateCancelling,
	"cancelled":   FirmwareReportLifecycleStateCancelled,
}

// GetFirmwareReportLifecycleStateEnumValues Enumerates the set of values for FirmwareReportLifecycleStateEnum
func GetFirmwareReportLifecycleStateEnumValues() []FirmwareReportLifecycleStateEnum {
	values := make([]FirmwareReportLifecycleStateEnum, 0)
	for _, v := range mappingFirmwareReportLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetFirmwareReportLifecycleStateEnumStringValues Enumerates the set of values in String for FirmwareReportLifecycleStateEnum
func GetFirmwareReportLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"IN_PROGRESS",
		"WAITING",
		"FAILED",
		"SUCCEEDED",
		"CANCELLING",
		"CANCELLED",
	}
}

// GetMappingFirmwareReportLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFirmwareReportLifecycleStateEnum(val string) (FirmwareReportLifecycleStateEnum, bool) {
	enum, ok := mappingFirmwareReportLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
