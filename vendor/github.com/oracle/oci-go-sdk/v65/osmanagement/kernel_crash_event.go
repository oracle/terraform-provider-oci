// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OS Management API
//
// API for the OS Management service. Use these API operations for working
// with Managed instances and Managed instance groups.
//

package osmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// KernelCrashEvent Information about a Kernel Crash.
type KernelCrashEvent struct {

	// OCID identifier of the event
	Id *string `mandatory:"true" json:"id"`

	// OCI identifier of the instance where the event occurred
	InstanceId *string `mandatory:"false" json:"instanceId"`

	// OCI identifier of the compartement where the instance is
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// OCID identifier of the instance tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// human readable description of the event
	Summary *string `mandatory:"false" json:"summary"`

	// Time of the occurrence of the event
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// Unique ID used to group event with the same characteristics together.
	// The list of such groups of event can be retrieved via /recurringEvents/{EventFingerprint}
	EventFingerprint *string `mandatory:"false" json:"eventFingerprint"`

	// Event occurrence count. Number of time the event has happen on the system.
	Count *int `mandatory:"false" json:"count"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// reason of the crash
	Reason *string `mandatory:"false" json:"reason"`

	// First occurrence time of the event
	TimeFirstOccurred *common.SDKTime `mandatory:"false" json:"timeFirstOccurred"`

	Vmcore *KernelVmCoreInformation `mandatory:"false" json:"vmcore"`

	Content *EventContent `mandatory:"false" json:"content"`

	System *CrashEventSystemInformation `mandatory:"false" json:"system"`
}

// GetId returns Id
func (m KernelCrashEvent) GetId() *string {
	return m.Id
}

// GetInstanceId returns InstanceId
func (m KernelCrashEvent) GetInstanceId() *string {
	return m.InstanceId
}

// GetCompartmentId returns CompartmentId
func (m KernelCrashEvent) GetCompartmentId() *string {
	return m.CompartmentId
}

// GetTenancyId returns TenancyId
func (m KernelCrashEvent) GetTenancyId() *string {
	return m.TenancyId
}

// GetSummary returns Summary
func (m KernelCrashEvent) GetSummary() *string {
	return m.Summary
}

// GetTimestamp returns Timestamp
func (m KernelCrashEvent) GetTimestamp() *common.SDKTime {
	return m.Timestamp
}

// GetEventFingerprint returns EventFingerprint
func (m KernelCrashEvent) GetEventFingerprint() *string {
	return m.EventFingerprint
}

// GetCount returns Count
func (m KernelCrashEvent) GetCount() *int {
	return m.Count
}

// GetFreeformTags returns FreeformTags
func (m KernelCrashEvent) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

// GetDefinedTags returns DefinedTags
func (m KernelCrashEvent) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

// GetSystemTags returns SystemTags
func (m KernelCrashEvent) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m KernelCrashEvent) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m KernelCrashEvent) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m KernelCrashEvent) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeKernelCrashEvent KernelCrashEvent
	s := struct {
		DiscriminatorParam string `json:"eventType"`
		MarshalTypeKernelCrashEvent
	}{
		"KERNEL_CRASH",
		(MarshalTypeKernelCrashEvent)(m),
	}

	return json.Marshal(&s)
}
