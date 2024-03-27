// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.cloud.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Event An event is an occurrence of something of interest on a managed instance,
// such as a kernel crash, software package update, or software source
// update.
type Event interface {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
	GetId() *string

	// Summary of the event.
	GetEventSummary() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	GetCompartmentId() *string

	// The date and time the Event was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeCreated() *common.SDKTime

	// The current state of the event.
	GetLifecycleState() EventLifecycleStateEnum

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Details of an event.
	GetEventDetails() *string

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance or resource where the event occurred.
	GetResourceId() *string

	GetSystemDetails() *SystemDetails

	// The date and time that the event occurred.
	GetTimeOccurred() *common.SDKTime

	// The date and time that the event was updated (in RFC 3339 (https://tools.ietf.org/html/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	GetTimeUpdated() *common.SDKTime

	// Describes the current state of the event in more detail. For example, the
	// message can provide actionable information for a resource in the 'FAILED' state.
	GetLifecycleDetails() *string

	// Indicates whether the event occurred on a resource that is managed by the Autonomous Linux service.
	GetIsManagedByAutonomousLinux() *bool

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type event struct {
	JsonData                   []byte
	EventDetails               *string                           `mandatory:"false" json:"eventDetails"`
	ResourceId                 *string                           `mandatory:"false" json:"resourceId"`
	SystemDetails              *SystemDetails                    `mandatory:"false" json:"systemDetails"`
	TimeOccurred               *common.SDKTime                   `mandatory:"false" json:"timeOccurred"`
	TimeUpdated                *common.SDKTime                   `mandatory:"false" json:"timeUpdated"`
	LifecycleDetails           *string                           `mandatory:"false" json:"lifecycleDetails"`
	IsManagedByAutonomousLinux *bool                             `mandatory:"false" json:"isManagedByAutonomousLinux"`
	SystemTags                 map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	Id                         *string                           `mandatory:"true" json:"id"`
	EventSummary               *string                           `mandatory:"true" json:"eventSummary"`
	CompartmentId              *string                           `mandatory:"true" json:"compartmentId"`
	TimeCreated                *common.SDKTime                   `mandatory:"true" json:"timeCreated"`
	LifecycleState             EventLifecycleStateEnum           `mandatory:"true" json:"lifecycleState"`
	FreeformTags               map[string]string                 `mandatory:"true" json:"freeformTags"`
	DefinedTags                map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`
	Type                       string                            `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *event) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerevent event
	s := struct {
		Model Unmarshalerevent
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.EventSummary = s.Model.EventSummary
	m.CompartmentId = s.Model.CompartmentId
	m.TimeCreated = s.Model.TimeCreated
	m.LifecycleState = s.Model.LifecycleState
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.EventDetails = s.Model.EventDetails
	m.ResourceId = s.Model.ResourceId
	m.SystemDetails = s.Model.SystemDetails
	m.TimeOccurred = s.Model.TimeOccurred
	m.TimeUpdated = s.Model.TimeUpdated
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.IsManagedByAutonomousLinux = s.Model.IsManagedByAutonomousLinux
	m.SystemTags = s.Model.SystemTags
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *event) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "SOFTWARE_UPDATE":
		mm := SoftwareUpdateEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KERNEL_OOPS":
		mm := KernelOopsEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "MANAGEMENT_STATION":
		mm := ManagementStationEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SOFTWARE_SOURCE":
		mm := SoftwareSourceEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KERNEL_CRASH":
		mm := KernelCrashEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "EXPLOIT_ATTEMPT":
		mm := ExploitAttemptEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AGENT":
		mm := AgentEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KSPLICE_UPDATE":
		mm := KspliceUpdateEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for Event: %s.", m.Type)
		return *m, nil
	}
}

// GetEventDetails returns EventDetails
func (m event) GetEventDetails() *string {
	return m.EventDetails
}

// GetResourceId returns ResourceId
func (m event) GetResourceId() *string {
	return m.ResourceId
}

// GetSystemDetails returns SystemDetails
func (m event) GetSystemDetails() *SystemDetails {
	return m.SystemDetails
}

// GetTimeOccurred returns TimeOccurred
func (m event) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

// GetTimeUpdated returns TimeUpdated
func (m event) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleDetails returns LifecycleDetails
func (m event) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetIsManagedByAutonomousLinux returns IsManagedByAutonomousLinux
func (m event) GetIsManagedByAutonomousLinux() *bool {
	return m.IsManagedByAutonomousLinux
}

// GetSystemTags returns SystemTags
func (m event) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

// GetId returns Id
func (m event) GetId() *string {
	return m.Id
}

// GetEventSummary returns EventSummary
func (m event) GetEventSummary() *string {
	return m.EventSummary
}

// GetCompartmentId returns CompartmentId
func (m event) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTimeCreated returns TimeCreated
func (m event) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetLifecycleState returns LifecycleState
func (m event) GetLifecycleState() EventLifecycleStateEnum {
	return m.LifecycleState
}

// GetFreeformTags returns FreeformTags
func (m event) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m event) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m event) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m event) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingEventLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEventLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// EventLifecycleStateEnum Enum with underlying type: string
type EventLifecycleStateEnum string

// Set of constants representing the allowable values for EventLifecycleStateEnum
const (
	EventLifecycleStateCreating EventLifecycleStateEnum = "CREATING"
	EventLifecycleStateUpdating EventLifecycleStateEnum = "UPDATING"
	EventLifecycleStateActive   EventLifecycleStateEnum = "ACTIVE"
	EventLifecycleStateDeleting EventLifecycleStateEnum = "DELETING"
	EventLifecycleStateDeleted  EventLifecycleStateEnum = "DELETED"
	EventLifecycleStateFailed   EventLifecycleStateEnum = "FAILED"
)

var mappingEventLifecycleStateEnum = map[string]EventLifecycleStateEnum{
	"CREATING": EventLifecycleStateCreating,
	"UPDATING": EventLifecycleStateUpdating,
	"ACTIVE":   EventLifecycleStateActive,
	"DELETING": EventLifecycleStateDeleting,
	"DELETED":  EventLifecycleStateDeleted,
	"FAILED":   EventLifecycleStateFailed,
}

var mappingEventLifecycleStateEnumLowerCase = map[string]EventLifecycleStateEnum{
	"creating": EventLifecycleStateCreating,
	"updating": EventLifecycleStateUpdating,
	"active":   EventLifecycleStateActive,
	"deleting": EventLifecycleStateDeleting,
	"deleted":  EventLifecycleStateDeleted,
	"failed":   EventLifecycleStateFailed,
}

// GetEventLifecycleStateEnumValues Enumerates the set of values for EventLifecycleStateEnum
func GetEventLifecycleStateEnumValues() []EventLifecycleStateEnum {
	values := make([]EventLifecycleStateEnum, 0)
	for _, v := range mappingEventLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetEventLifecycleStateEnumStringValues Enumerates the set of values in String for EventLifecycleStateEnum
func GetEventLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingEventLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingEventLifecycleStateEnum(val string) (EventLifecycleStateEnum, bool) {
	enum, ok := mappingEventLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
