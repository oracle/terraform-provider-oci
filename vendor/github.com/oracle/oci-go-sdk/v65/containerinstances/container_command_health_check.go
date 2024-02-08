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

// ContainerCommandHealthCheck Container Health Check with command type.
type ContainerCommandHealthCheck struct {

	// The list of strings that will be simplified to a single command for checking the status of the container.
	Command []string `mandatory:"true" json:"command"`

	// Health check name.
	Name *string `mandatory:"false" json:"name"`

	// The initial delay in seconds before start checking container health status.
	InitialDelayInSeconds *int `mandatory:"false" json:"initialDelayInSeconds"`

	// Number of seconds between two consecutive runs for checking container health.
	IntervalInSeconds *int `mandatory:"false" json:"intervalInSeconds"`

	// Number of consecutive failures at which we consider the check failed.
	FailureThreshold *int `mandatory:"false" json:"failureThreshold"`

	// Number of consecutive successes at which we consider the check succeeded again after it was in failure state.
	SuccessThreshold *int `mandatory:"false" json:"successThreshold"`

	// Length of waiting time in seconds before marking health check failed.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`

	// A message describing the current status in more details.
	StatusDetails *string `mandatory:"false" json:"statusDetails"`

	// Status of container
	Status ContainerHealthCheckStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default
	// action is KILL. If failure action is KILL, the container will be subject to the container restart policy.
	FailureAction ContainerHealthCheckFailureActionEnum `mandatory:"false" json:"failureAction,omitempty"`
}

// GetName returns Name
func (m ContainerCommandHealthCheck) GetName() *string {
	return m.Name
}

// GetInitialDelayInSeconds returns InitialDelayInSeconds
func (m ContainerCommandHealthCheck) GetInitialDelayInSeconds() *int {
	return m.InitialDelayInSeconds
}

// GetIntervalInSeconds returns IntervalInSeconds
func (m ContainerCommandHealthCheck) GetIntervalInSeconds() *int {
	return m.IntervalInSeconds
}

// GetFailureThreshold returns FailureThreshold
func (m ContainerCommandHealthCheck) GetFailureThreshold() *int {
	return m.FailureThreshold
}

// GetSuccessThreshold returns SuccessThreshold
func (m ContainerCommandHealthCheck) GetSuccessThreshold() *int {
	return m.SuccessThreshold
}

// GetTimeoutInSeconds returns TimeoutInSeconds
func (m ContainerCommandHealthCheck) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

// GetStatus returns Status
func (m ContainerCommandHealthCheck) GetStatus() ContainerHealthCheckStatusEnum {
	return m.Status
}

// GetStatusDetails returns StatusDetails
func (m ContainerCommandHealthCheck) GetStatusDetails() *string {
	return m.StatusDetails
}

// GetFailureAction returns FailureAction
func (m ContainerCommandHealthCheck) GetFailureAction() ContainerHealthCheckFailureActionEnum {
	return m.FailureAction
}

func (m ContainerCommandHealthCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerCommandHealthCheck) ValidateEnumValue() (bool, error) {
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

// MarshalJSON marshals to json representation
func (m ContainerCommandHealthCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerCommandHealthCheck ContainerCommandHealthCheck
	s := struct {
		DiscriminatorParam string `json:"healthCheckType"`
		MarshalTypeContainerCommandHealthCheck
	}{
		"COMMAND",
		(MarshalTypeContainerCommandHealthCheck)(m),
	}

	return json.Marshal(&s)
}
