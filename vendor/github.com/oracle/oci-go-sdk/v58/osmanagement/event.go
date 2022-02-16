// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Event Description of Event.
type Event interface {

	// OCID identifier of the event
	GetId() *string

	// OCI identifier of the instance where the event occurred
	GetInstanceId() *string

	// OCI identifier of the compartement where the instance is
	GetCompartmentId() *string

	// OCID identifier of the instance tenancy.
	GetTenancyId() *string

	// human readable description of the event
	GetSummary() *string

	// Time of the occurrence of the event
	GetTimestamp() *common.SDKTime

	// Unique ID used to group event with the same characteristics together.
	// The list of such groups of event can be retrieved via /recurringEvents/{EventFingerprint}
	GetEventFingerprint() *string

	// Event occurrence count. Number of time the event has happen on the system.
	GetCount() *int

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	GetSystemTags() map[string]map[string]interface{}
}

type event struct {
	JsonData         []byte
	Id               *string                           `mandatory:"true" json:"id"`
	InstanceId       *string                           `mandatory:"false" json:"instanceId"`
	CompartmentId    *string                           `mandatory:"false" json:"compartmentId"`
	TenancyId        *string                           `mandatory:"false" json:"tenancyId"`
	Summary          *string                           `mandatory:"false" json:"summary"`
	Timestamp        *common.SDKTime                   `mandatory:"false" json:"timestamp"`
	EventFingerprint *string                           `mandatory:"false" json:"eventFingerprint"`
	Count            *int                              `mandatory:"false" json:"count"`
	FreeformTags     map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags      map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	SystemTags       map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
	EventType        string                            `json:"eventType"`
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
	m.InstanceId = s.Model.InstanceId
	m.CompartmentId = s.Model.CompartmentId
	m.TenancyId = s.Model.TenancyId
	m.Summary = s.Model.Summary
	m.Timestamp = s.Model.Timestamp
	m.EventFingerprint = s.Model.EventFingerprint
	m.Count = s.Model.Count
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.SystemTags = s.Model.SystemTags
	m.EventType = s.Model.EventType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *event) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.EventType {
	case "KERNEL_OOPS":
		mm := KernelOopsEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "KERNEL_CRASH":
		mm := KernelCrashEvent{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetId returns Id
func (m event) GetId() *string {
	return m.Id
}

//GetInstanceId returns InstanceId
func (m event) GetInstanceId() *string {
	return m.InstanceId
}

//GetCompartmentId returns CompartmentId
func (m event) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTenancyId returns TenancyId
func (m event) GetTenancyId() *string {
	return m.TenancyId
}

//GetSummary returns Summary
func (m event) GetSummary() *string {
	return m.Summary
}

//GetTimestamp returns Timestamp
func (m event) GetTimestamp() *common.SDKTime {
	return m.Timestamp
}

//GetEventFingerprint returns EventFingerprint
func (m event) GetEventFingerprint() *string {
	return m.EventFingerprint
}

//GetCount returns Count
func (m event) GetCount() *int {
	return m.Count
}

//GetFreeformTags returns FreeformTags
func (m event) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m event) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m event) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m event) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m event) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
