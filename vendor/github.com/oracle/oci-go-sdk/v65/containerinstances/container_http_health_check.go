// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// ContainerHttpHealthCheck Container Health Check HTTP type.
type ContainerHttpHealthCheck struct {

	// Container health check HTTP path.
	Path *string `mandatory:"true" json:"path"`

	// Container health check HTTP port.
	Port *int `mandatory:"true" json:"port"`

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

	// If set to true, this health check runs first while other HealthChecks wait for this one to complete.
	// If this becomes Healthy then other health checks are started.
	// If it becomes Unhealthy the container is killed.
	// At max only 1 healthCheck can have this field set to true.
	IsStartupCheck *bool `mandatory:"false" json:"isStartupCheck"`

	// Container health check HTTP headers.
	Headers []HealthCheckHttpHeader `mandatory:"false" json:"headers"`

	// Status of container
	Status ContainerHealthCheckStatusEnum `mandatory:"false" json:"status,omitempty"`

	// The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default
	// action is KILL. If failure action is KILL, the container will be subject to the container restart policy.
	FailureAction ContainerHealthCheckFailureActionEnum `mandatory:"false" json:"failureAction,omitempty"`

	// Container health check HTTP port.
	Scheme ContainerHttpHealthCheckSchemeTypeEnum `mandatory:"false" json:"scheme,omitempty"`
}

// GetName returns Name
func (m ContainerHttpHealthCheck) GetName() *string {
	return m.Name
}

// GetInitialDelayInSeconds returns InitialDelayInSeconds
func (m ContainerHttpHealthCheck) GetInitialDelayInSeconds() *int {
	return m.InitialDelayInSeconds
}

// GetIntervalInSeconds returns IntervalInSeconds
func (m ContainerHttpHealthCheck) GetIntervalInSeconds() *int {
	return m.IntervalInSeconds
}

// GetFailureThreshold returns FailureThreshold
func (m ContainerHttpHealthCheck) GetFailureThreshold() *int {
	return m.FailureThreshold
}

// GetSuccessThreshold returns SuccessThreshold
func (m ContainerHttpHealthCheck) GetSuccessThreshold() *int {
	return m.SuccessThreshold
}

// GetTimeoutInSeconds returns TimeoutInSeconds
func (m ContainerHttpHealthCheck) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

// GetStatus returns Status
func (m ContainerHttpHealthCheck) GetStatus() ContainerHealthCheckStatusEnum {
	return m.Status
}

// GetStatusDetails returns StatusDetails
func (m ContainerHttpHealthCheck) GetStatusDetails() *string {
	return m.StatusDetails
}

// GetFailureAction returns FailureAction
func (m ContainerHttpHealthCheck) GetFailureAction() ContainerHealthCheckFailureActionEnum {
	return m.FailureAction
}

// GetIsStartupCheck returns IsStartupCheck
func (m ContainerHttpHealthCheck) GetIsStartupCheck() *bool {
	return m.IsStartupCheck
}

func (m ContainerHttpHealthCheck) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContainerHttpHealthCheck) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerHealthCheckStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetContainerHealthCheckStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContainerHealthCheckFailureActionEnum(string(m.FailureAction)); !ok && m.FailureAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FailureAction: %s. Supported values are: %s.", m.FailureAction, strings.Join(GetContainerHealthCheckFailureActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingContainerHttpHealthCheckSchemeTypeEnum(string(m.Scheme)); !ok && m.Scheme != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scheme: %s. Supported values are: %s.", m.Scheme, strings.Join(GetContainerHttpHealthCheckSchemeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ContainerHttpHealthCheck) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeContainerHttpHealthCheck ContainerHttpHealthCheck
	s := struct {
		DiscriminatorParam string `json:"healthCheckType"`
		MarshalTypeContainerHttpHealthCheck
	}{
		"HTTP",
		(MarshalTypeContainerHttpHealthCheck)(m),
	}

	return json.Marshal(&s)
}
