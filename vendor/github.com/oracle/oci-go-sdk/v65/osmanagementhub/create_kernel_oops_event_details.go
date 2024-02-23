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

// CreateKernelOopsEventDetails The data to create kernel oops event
type CreateKernelOopsEventDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the Event in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the managed instance or resource where the event occurred.
	ResourceId *string `mandatory:"true" json:"resourceId"`

	// True, if the event is managed by Autonomous Linux Service.
	IsManagedByAutonomousLinux *bool `mandatory:"true" json:"isManagedByAutonomousLinux"`

	Data *KernelEventData `mandatory:"true" json:"data"`

	// Summary of the event.
	EventSummary *string `mandatory:"false" json:"eventSummary"`

	// Details of an event.
	EventDetails *string `mandatory:"false" json:"eventDetails"`

	// Event occurred at
	TimeOccurred *common.SDKTime `mandatory:"false" json:"timeOccurred"`

	SystemDetails *SystemDetails `mandatory:"false" json:"systemDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

// GetEventSummary returns EventSummary
func (m CreateKernelOopsEventDetails) GetEventSummary() *string {
	return m.EventSummary
}

// GetEventDetails returns EventDetails
func (m CreateKernelOopsEventDetails) GetEventDetails() *string {
	return m.EventDetails
}

// GetTimeOccurred returns TimeOccurred
func (m CreateKernelOopsEventDetails) GetTimeOccurred() *common.SDKTime {
	return m.TimeOccurred
}

// GetCompartmentId returns CompartmentId
func (m CreateKernelOopsEventDetails) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetResourceId returns ResourceId
func (m CreateKernelOopsEventDetails) GetResourceId() *string {
	return m.ResourceId
}

// GetSystemDetails returns SystemDetails
func (m CreateKernelOopsEventDetails) GetSystemDetails() *SystemDetails {
	return m.SystemDetails
}

// GetIsManagedByAutonomousLinux returns IsManagedByAutonomousLinux
func (m CreateKernelOopsEventDetails) GetIsManagedByAutonomousLinux() *bool {
	return m.IsManagedByAutonomousLinux
}

// GetFreeformTags returns FreeformTags
func (m CreateKernelOopsEventDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m CreateKernelOopsEventDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateKernelOopsEventDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateKernelOopsEventDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateKernelOopsEventDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateKernelOopsEventDetails CreateKernelOopsEventDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateKernelOopsEventDetails
	}{
		"KERNEL_OOPS",
		(MarshalTypeCreateKernelOopsEventDetails)(m),
	}

	return json.Marshal(&s)
}
