// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management Hub API
//
// Use the OS Management Hub API to manage and monitor updates and patches for instances in OCI, your private data center, or 3rd-party clouds.
// For more information, see Overview of OS Management Hub (https://docs.oracle.com/iaas/osmh/doc/overview.htm).
//

package osmanagementhub

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KernelOopsEvent Information about the kernel OOPS event. A kernel OOPS event occurs when the kernel detects an erroneous state inside itself.
type KernelOopsEvent struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the event.
	Id *string `mandatory:"true" json:"id"`

	// Summary of the event.
	EventSummary *string `mandatory:"true" json:"eventSummary"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The date and time the Event was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	Data *KernelEventData `mandatory:"true" json:"data"`

	// Details of an event.
	EventDetails *string `mandatory:"false" json:"eventDetails"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance or resource where the event occurred.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	SystemDetails *SystemDetails `mandatory:"false" json:"systemDetails"`

	// The date and time that the event occurred.
	TimeOccurred *common.SDKTime `mandatory:"false" json:"timeOccurred"`

	// The date and time that the event was updated (in RFC 3339 (https://tools.ietf.org/html/rfc3339) format).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Describes the current state of the event in more detail. For example, the
	// message can provide actionable information for a resource in the 'FAILED' state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Indicates whether the event occurred on a resource that is managed by the Autonomous Linux service.
	IsManagedByAutonomousLinux *bool `mandatory:"false" json:"isManagedByAutonomousLinux"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the event.
	LifecycleState EventLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`
}

// GetId returns Id
func (m KernelOopsEvent) GetId() *string {
	return m.Id
}

// GetEventSummary returns EventSummary
func (m KernelOopsEvent) GetEventSummary() *string {
	return m.EventSummary
}

// GetCompartmentId returns CompartmentId
func (m KernelOopsEvent) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetEventDetails returns EventDetails
func (m KernelOopsEvent) GetEventDetails() *string {
	return m.EventDetails
}

// GetResourceId returns ResourceId
func (m KernelOopsEvent) GetResourceId() *string {
	return m.ResourceId
}

// GetSystemDetails returns SystemDetails
func (m KernelOopsEvent) GetSystemDetails() *SystemDetails {
	return m.SystemDetails
}

// GetTimeOccurred returns TimeOccurred
func (m KernelOopsEvent) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

// GetTimeCreated returns TimeCreated
func (m KernelOopsEvent) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

// GetTimeUpdated returns TimeUpdated
func (m KernelOopsEvent) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

// GetLifecycleState returns LifecycleState
func (m KernelOopsEvent) GetLifecycleState() EventLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m KernelOopsEvent) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetIsManagedByAutonomousLinux returns IsManagedByAutonomousLinux
func (m KernelOopsEvent) GetIsManagedByAutonomousLinux() *bool {
	return m.IsManagedByAutonomousLinux
}

// GetFreeformTags returns FreeformTags
func (m KernelOopsEvent) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m KernelOopsEvent) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m KernelOopsEvent) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m KernelOopsEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KernelOopsEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingEventLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetEventLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KernelOopsEvent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKernelOopsEvent KernelOopsEvent
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeKernelOopsEvent
	}{
		"KERNEL_OOPS",
		(MarshalTypeKernelOopsEvent)(m),
	}

	return json.Marshal(&s)
}
