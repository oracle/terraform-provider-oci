// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FleetResourceBasedActionGroupDetails A string variable that holds a value
type FleetResourceBasedActionGroupDetails struct {

	// Name of the ActionGroup.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Product associated.
	// Only applicable if actionGroup type is PRODUCT.
	Product *string `mandatory:"false" json:"product"`

	// LifeCycle Operation.
	LifecycleOperation *string `mandatory:"false" json:"lifecycleOperation"`

	// Unique producer Id at Action Group Level
	ActivityId *string `mandatory:"false" json:"activityId"`

	// The time the Scheduler Job started. An RFC3339 formatted datetime string.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The time the Scheduler Job ended. An RFC3339 formatted datetime string.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// ID of the fleet
	FleetId *string `mandatory:"false" json:"fleetId"`

	// Sequence of the Action Group.
	// Action groups will be executed in a seuential order.
	// All Action Groups having the same sequence will be executed parallely.
	// If no value is provided a default value of 1 will be given.
	Sequence *int `mandatory:"false" json:"sequence"`

	// ID of the runbook
	RunbookId *string `mandatory:"false" json:"runbookId"`

	// Name of the runbook version
	RunbookVersionName *string `mandatory:"false" json:"runbookVersionName"`

	// Emergency override flag.
	IsEmergencyOverride *bool `mandatory:"false" json:"isEmergencyOverride"`

	// List of selected fleet resources.
	SelectedFleetResources []string `mandatory:"false" json:"selectedFleetResources"`

	// Status of the Job at Action Group Level.
	Status JobStatusEnum `mandatory:"false" json:"status,omitempty"`

	// Subject Type
	SubjectType SubjectTypeEnum `mandatory:"false" json:"subjectType,omitempty"`
}

// GetDisplayName returns DisplayName
func (m FleetResourceBasedActionGroupDetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetProduct returns Product
func (m FleetResourceBasedActionGroupDetails) GetProduct() *string {
	return m.Product
}

// GetLifecycleOperation returns LifecycleOperation
func (m FleetResourceBasedActionGroupDetails) GetLifecycleOperation() *string {
	return m.LifecycleOperation
}

// GetActivityId returns ActivityId
func (m FleetResourceBasedActionGroupDetails) GetActivityId() *string {
	return m.ActivityId
}

// GetStatus returns Status
func (m FleetResourceBasedActionGroupDetails) GetStatus() JobStatusEnum {
	return m.Status
}

// GetTimeStarted returns TimeStarted
func (m FleetResourceBasedActionGroupDetails) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeEnded returns TimeEnded
func (m FleetResourceBasedActionGroupDetails) GetTimeEnded() *common.SDKTime {
	return m.TimeEnded
}

func (m FleetResourceBasedActionGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FleetResourceBasedActionGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingSubjectTypeEnum(string(m.SubjectType)); !ok && m.SubjectType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubjectType: %s. Supported values are: %s.", m.SubjectType, strings.Join(GetSubjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m FleetResourceBasedActionGroupDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFleetResourceBasedActionGroupDetails FleetResourceBasedActionGroupDetails
	s := struct {
		DiscriminatorParam string `json:"kind"`
		MarshalTypeFleetResourceBasedActionGroupDetails
	}{
		"FLEET_RESOURCES_USING_RUNBOOK",
		(MarshalTypeFleetResourceBasedActionGroupDetails)(m),
	}

	return json.Marshal(&s)
}
