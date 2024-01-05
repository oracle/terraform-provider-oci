// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Container Instance API
//
// A description of the Container Instance API
//

package containerinstances

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ContainerHealthCheck Type of container health check which could be either HTTP, TCP, or Command.
type ContainerHealthCheck interface {

	// Health check name.
	GetName() *string

	// The initial delay in seconds before start checking container health status.
	GetInitialDelayInSeconds() *int

	// Number of seconds between two consecutive runs for checking container health.
	GetIntervalInSeconds() *int

	// Number of consecutive failures at which we consider the check failed.
	GetFailureThreshold() *int

	// Number of consecutive successes at which we consider the check succeeded again after it was in failure state.
	GetSuccessThreshold() *int

	// Length of waiting time in seconds before marking health check failed.
	GetTimeoutInSeconds() *int

	// Status of container
	GetStatus() ContainerHealthCheckStatusEnum

	// A message describing the current status in more details.
	GetStatusDetails() *string

	// The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default
	// action is KILL. If failure action is KILL, the container will be subject to the container restart policy.
	GetFailureAction() ContainerHealthCheckFailureActionEnum
}

type containerhealthcheck struct {
	JsonData              []byte
	Name                  *string                               `mandatory:"false" json:"name"`
	InitialDelayInSeconds *int                                  `mandatory:"false" json:"initialDelayInSeconds"`
	IntervalInSeconds     *int                                  `mandatory:"false" json:"intervalInSeconds"`
	FailureThreshold      *int                                  `mandatory:"false" json:"failureThreshold"`
	SuccessThreshold      *int                                  `mandatory:"false" json:"successThreshold"`
	TimeoutInSeconds      *int                                  `mandatory:"false" json:"timeoutInSeconds"`
	Status                ContainerHealthCheckStatusEnum        `mandatory:"false" json:"status,omitempty"`
	StatusDetails         *string                               `mandatory:"false" json:"statusDetails"`
	FailureAction         ContainerHealthCheckFailureActionEnum `mandatory:"false" json:"failureAction,omitempty"`
	HealthCheckType       string                                `json:"healthCheckType"`
}

// UnmarshalJSON unmarshals json
func (m *containerhealthcheck) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercontainerhealthcheck containerhealthcheck
	s := struct {
		Model Unmarshalercontainerhealthcheck
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.InitialDelayInSeconds = s.Model.InitialDelayInSeconds
	m.IntervalInSeconds = s.Model.IntervalInSeconds
	m.FailureThreshold = s.Model.FailureThreshold
	m.SuccessThreshold = s.Model.SuccessThreshold
	m.TimeoutInSeconds = s.Model.TimeoutInSeconds
	m.Status = s.Model.Status
	m.StatusDetails = s.Model.StatusDetails
	m.FailureAction = s.Model.FailureAction
	m.HealthCheckType = s.Model.HealthCheckType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *containerhealthcheck) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.HealthCheckType {
	case "TCP":
		mm := ContainerTcpHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := ContainerHttpHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMMAND":
		mm := ContainerCommandHealthCheck{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ContainerHealthCheck: %s.", m.HealthCheckType)
		return *m, nil
	}
}

// GetName returns Name
func (m containerhealthcheck) GetName() *string {
	return m.Name
}

// GetInitialDelayInSeconds returns InitialDelayInSeconds
func (m containerhealthcheck) GetInitialDelayInSeconds() *int {
	return m.InitialDelayInSeconds
}

// GetIntervalInSeconds returns IntervalInSeconds
func (m containerhealthcheck) GetIntervalInSeconds() *int {
	return m.IntervalInSeconds
}

// GetFailureThreshold returns FailureThreshold
func (m containerhealthcheck) GetFailureThreshold() *int {
	return m.FailureThreshold
}

// GetSuccessThreshold returns SuccessThreshold
func (m containerhealthcheck) GetSuccessThreshold() *int {
	return m.SuccessThreshold
}

// GetTimeoutInSeconds returns TimeoutInSeconds
func (m containerhealthcheck) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

// GetStatus returns Status
func (m containerhealthcheck) GetStatus() ContainerHealthCheckStatusEnum {
	return m.Status
}

// GetStatusDetails returns StatusDetails
func (m containerhealthcheck) GetStatusDetails() *string {
	return m.StatusDetails
}

// GetFailureAction returns FailureAction
func (m containerhealthcheck) GetFailureAction() ContainerHealthCheckFailureActionEnum {
	return m.FailureAction
}

func (m containerhealthcheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m containerhealthcheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerHealthCheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetContainerHealthCheckStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContainerHealthCheckFailureActionEnum(string(m.FailureAction)); !ok && m.FailureAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FailureAction: %s. Supported values are: %s.", m.FailureAction, strings.Join(GetContainerHealthCheckFailureActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ContainerHealthCheckStatusEnum Enum with underlying type: string
type ContainerHealthCheckStatusEnum string

// Set of constants representing the allowable values for ContainerHealthCheckStatusEnum
const (
	ContainerHealthCheckStatusHealthy   ContainerHealthCheckStatusEnum = "HEALTHY"
	ContainerHealthCheckStatusUnhealthy ContainerHealthCheckStatusEnum = "UNHEALTHY"
	ContainerHealthCheckStatusUnknown   ContainerHealthCheckStatusEnum = "UNKNOWN"
)

var mappingContainerHealthCheckStatusEnum = map[string]ContainerHealthCheckStatusEnum{
	"HEALTHY":   ContainerHealthCheckStatusHealthy,
	"UNHEALTHY": ContainerHealthCheckStatusUnhealthy,
	"UNKNOWN":   ContainerHealthCheckStatusUnknown,
}

var mappingContainerHealthCheckStatusEnumLowerCase = map[string]ContainerHealthCheckStatusEnum{
	"healthy":   ContainerHealthCheckStatusHealthy,
	"unhealthy": ContainerHealthCheckStatusUnhealthy,
	"unknown":   ContainerHealthCheckStatusUnknown,
}

// GetContainerHealthCheckStatusEnumValues Enumerates the set of values for ContainerHealthCheckStatusEnum
func GetContainerHealthCheckStatusEnumValues() []ContainerHealthCheckStatusEnum {
	values := make([]ContainerHealthCheckStatusEnum, 0)
	for _, v := range mappingContainerHealthCheckStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetContainerHealthCheckStatusEnumStringValues Enumerates the set of values in String for ContainerHealthCheckStatusEnum
func GetContainerHealthCheckStatusEnumStringValues() []string {
	return []string{
		"HEALTHY",
		"UNHEALTHY",
		"UNKNOWN",
	}
}

// GetMappingContainerHealthCheckStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContainerHealthCheckStatusEnum(val string) (ContainerHealthCheckStatusEnum, bool) {
	enum, ok := mappingContainerHealthCheckStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
