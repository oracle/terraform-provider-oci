// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ActionGroupDetails Action Group details.
type ActionGroupDetails interface {

	// Name of the ActionGroup.
	GetDisplayName() *string

	// Product associated.
	// Only applicable if actionGroup type is PRODUCT.
	GetProduct() *string

	// LifeCycle Operation.
	GetLifecycleOperation() *string

	// Unique producer Id at Action Group Level
	GetActivityId() *string

	// Status of the Job at Action Group Level.
	GetStatus() JobStatusEnum

	// The time the Scheduler Job started. An RFC3339 formatted datetime string.
	GetTimeStarted() *common.SDKTime

	// The time the Scheduler Job ended. An RFC3339 formatted datetime string.
	GetTimeEnded() *common.SDKTime
}

type actiongroupdetails struct {
	JsonData           []byte
	DisplayName        *string         `mandatory:"false" json:"displayName"`
	Product            *string         `mandatory:"false" json:"product"`
	LifecycleOperation *string         `mandatory:"false" json:"lifecycleOperation"`
	ActivityId         *string         `mandatory:"false" json:"activityId"`
	Status             JobStatusEnum   `mandatory:"false" json:"status,omitempty"`
	TimeStarted        *common.SDKTime `mandatory:"false" json:"timeStarted"`
	TimeEnded          *common.SDKTime `mandatory:"false" json:"timeEnded"`
	Kind               string          `json:"kind"`
}

// UnmarshalJSON unmarshals json
func (m *actiongroupdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshaleractiongroupdetails actiongroupdetails
	s := struct {
		Model Unmarshaleractiongroupdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.DisplayName = s.Model.DisplayName
	m.Product = s.Model.Product
	m.LifecycleOperation = s.Model.LifecycleOperation
	m.ActivityId = s.Model.ActivityId
	m.Status = s.Model.Status
	m.TimeStarted = s.Model.TimeStarted
	m.TimeEnded = s.Model.TimeEnded
	m.Kind = s.Model.Kind

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *actiongroupdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Kind {
	case "FLEET_USING_RUNBOOK":
		mm := FleetBasedActionGroupDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for ActionGroupDetails: %s.", m.Kind)
		return *m, nil
	}
}

// GetDisplayName returns DisplayName
func (m actiongroupdetails) GetDisplayName() *string {
	return m.DisplayName
}

// GetProduct returns Product
func (m actiongroupdetails) GetProduct() *string {
	return m.Product
}

// GetLifecycleOperation returns LifecycleOperation
func (m actiongroupdetails) GetLifecycleOperation() *string {
	return m.LifecycleOperation
}

// GetActivityId returns ActivityId
func (m actiongroupdetails) GetActivityId() *string {
	return m.ActivityId
}

// GetStatus returns Status
func (m actiongroupdetails) GetStatus() JobStatusEnum {
	return m.Status
}

// GetTimeStarted returns TimeStarted
func (m actiongroupdetails) GetTimeStarted() *common.SDKTime {
	return m.TimeStarted
}

// GetTimeEnded returns TimeEnded
func (m actiongroupdetails) GetTimeEnded() *common.SDKTime {
	return m.TimeEnded
}

func (m actiongroupdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m actiongroupdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingJobStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetJobStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ActionGroupDetailsKindEnum Enum with underlying type: string
type ActionGroupDetailsKindEnum string

// Set of constants representing the allowable values for ActionGroupDetailsKindEnum
const (
	ActionGroupDetailsKindFleetUsingRunbook ActionGroupDetailsKindEnum = "FLEET_USING_RUNBOOK"
)

var mappingActionGroupDetailsKindEnum = map[string]ActionGroupDetailsKindEnum{
	"FLEET_USING_RUNBOOK": ActionGroupDetailsKindFleetUsingRunbook,
}

var mappingActionGroupDetailsKindEnumLowerCase = map[string]ActionGroupDetailsKindEnum{
	"fleet_using_runbook": ActionGroupDetailsKindFleetUsingRunbook,
}

// GetActionGroupDetailsKindEnumValues Enumerates the set of values for ActionGroupDetailsKindEnum
func GetActionGroupDetailsKindEnumValues() []ActionGroupDetailsKindEnum {
	values := make([]ActionGroupDetailsKindEnum, 0)
	for _, v := range mappingActionGroupDetailsKindEnum {
		values = append(values, v)
	}
	return values
}

// GetActionGroupDetailsKindEnumStringValues Enumerates the set of values in String for ActionGroupDetailsKindEnum
func GetActionGroupDetailsKindEnumStringValues() []string {
	return []string{
		"FLEET_USING_RUNBOOK",
	}
}

// GetMappingActionGroupDetailsKindEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionGroupDetailsKindEnum(val string) (ActionGroupDetailsKindEnum, bool) {
	enum, ok := mappingActionGroupDetailsKindEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
