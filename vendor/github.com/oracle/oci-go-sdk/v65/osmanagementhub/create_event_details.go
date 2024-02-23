// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for the operating system environments in your private data centers through a single management console. For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateEventDetails The data to create a Event.
type CreateEventDetails interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Event in.
	GetCompartmentId() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance or resource where the event occurred.
	GetResourceId() *string

	// True, if the event is managed by Autonomous Linux Service.
	GetIsManagedByAutonomousLinux() *bool

	// Summary of the event.
	GetEventSummary() *string

	// Details of an event.
	GetEventDetails() *string

	// Event occurred at
	GetTimeOccurred() *common.SDKTime

	GetSystemDetails() *SystemDetails

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type createeventdetails struct {
	JsonData                   []byte
	EventSummary               *string                           `mandatory:"false" json:"eventSummary"`
	EventDetails               *string                           `mandatory:"false" json:"eventDetails"`
	TimeOccurred               *common.SDKTime                   `mandatory:"false" json:"timeOccurred"`
	SystemDetails              *SystemDetails                    `mandatory:"false" json:"systemDetails"`
	FreeformTags               map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	ResourceId                 *string                           `mandatory:"true" json:"resourceId"`
	IsManagedByAutonomousLinux *bool                             `mandatory:"true" json:"isManagedByAutonomousLinux"`
	Type                       string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createeventdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateeventdetails createeventdetails
	s := struct {
		Model Unmarshalercreateeventdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.CompartmentId = s.Model.CompartmentId
	m.ResourceId = s.Model.ResourceId
	m.IsManagedByAutonomousLinux = s.Model.IsManagedByAutonomousLinux
	m.EventSummary = s.Model.EventSummary
	m.EventDetails = s.Model.EventDetails
	m.TimeOccurred = s.Model.TimeOccurred
	m.SystemDetails = s.Model.SystemDetails
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createeventdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "KSPLICE_UPDATE":
		mm := CreateKspliceUpdateEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANAGEMENT_STATION":
		mm := CreateManagementStationEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KERNEL_OOPS":
		mm := CreateKernelOopsEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOFTWARE_UPDATE":
		mm := CreateSoftwareUpdateEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXPLOIT_ATTEMPT":
		mm := CreateExploitAttemptEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AGENT":
		mm := CreateAgentEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOFTWARE_SOURCE":
		mm := CreateSoftwareSourceEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KERNEL_CRASH":
		mm := CreateKernelCrashEventDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateEventDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetEventSummary returns EventSummary
func (m createeventdetails) GetEventSummary() *string {
	return m.EventSummary
}

// GetEventDetails returns EventDetails
func (m createeventdetails) GetEventDetails() *string {
	return m.EventDetails
}

// GetTimeOccurred returns TimeOccurred
func (m createeventdetails) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

// GetSystemDetails returns SystemDetails
func (m createeventdetails) GetSystemDetails() *SystemDetails {
	return m.SystemDetails
}

// GetFreeformTags returns FreeformTags
func (m createeventdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m createeventdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetCompartmentId returns CompartmentId
func (m createeventdetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetResourceId returns ResourceId
func (m createeventdetails) GetResourceId() *string {
	return m.ResourceId
}

// GetIsManagedByAutonomousLinux returns IsManagedByAutonomousLinux
func (m createeventdetails) GetIsManagedByAutonomousLinux() *bool {
	return m.IsManagedByAutonomousLinux
}

func (m createeventdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createeventdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
