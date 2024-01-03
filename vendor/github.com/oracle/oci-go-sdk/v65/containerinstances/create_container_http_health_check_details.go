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

// CreateContainerHttpHealthCheckDetails Container Health Check HTTP type.
type CreateContainerHttpHealthCheckDetails struct {

	// Container health check HTTP path.
	Path *string `mandatory:"true" json:"path"`

	// Container health check HTTP port.
	Port *int `mandatory:"true" json:"port"`

	// Health check name.
	Name *string `mandatory:"false" json:"name"`

	// If set to true, this health check runs first while other HealthChecks wait for this one to complete.
	// If this becomes Healthy then other health checks are started.
	// If it becomes Unhealthy the container is killed.
	// At max only 1 healthCheck can have this field set to true.
	IsStartupCheck *bool `mandatory:"false" json:"isStartupCheck"`

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

	// Container health check HTTP headers.
	Headers []HealthCheckHttpHeader `mandatory:"false" json:"headers"`

	// The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default
	// action is KILL. If failure action is KILL, the container will be subject to the container restart policy.
	FailureAction ContainerHealthCheckFailureActionEnum `mandatory:"false" json:"failureAction,omitempty"`

	// Container health check HTTP port.
	Scheme ContainerHttpHealthCheckSchemeTypeEnum `mandatory:"false" json:"scheme,omitempty"`
}

// GetName returns Name
func (m CreateContainerHttpHealthCheckDetails) GetName() *string {
	return m.Name
}

// GetIsStartupCheck returns IsStartupCheck
func (m CreateContainerHttpHealthCheckDetails) GetIsStartupCheck() *bool {
	return m.IsStartupCheck
}

// GetInitialDelayInSeconds returns InitialDelayInSeconds
func (m CreateContainerHttpHealthCheckDetails) GetInitialDelayInSeconds() *int {
	return m.InitialDelayInSeconds
}

// GetIntervalInSeconds returns IntervalInSeconds
func (m CreateContainerHttpHealthCheckDetails) GetIntervalInSeconds() *int {
	return m.IntervalInSeconds
}

// GetFailureThreshold returns FailureThreshold
func (m CreateContainerHttpHealthCheckDetails) GetFailureThreshold() *int {
	return m.FailureThreshold
}

// GetSuccessThreshold returns SuccessThreshold
func (m CreateContainerHttpHealthCheckDetails) GetSuccessThreshold() *int {
	return m.SuccessThreshold
}

// GetTimeoutInSeconds returns TimeoutInSeconds
func (m CreateContainerHttpHealthCheckDetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

// GetFailureAction returns FailureAction
func (m CreateContainerHttpHealthCheckDetails) GetFailureAction() ContainerHealthCheckFailureActionEnum {
	return m.FailureAction
}

func (m CreateContainerHttpHealthCheckDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateContainerHttpHealthCheckDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

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
func (m CreateContainerHttpHealthCheckDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateContainerHttpHealthCheckDetails CreateContainerHttpHealthCheckDetails
	s := struct {
		DiscriminatorParam string `json:"healthCheckType"`
		MarshalTypeCreateContainerHttpHealthCheckDetails
	}{
		"HTTP",
		(MarshalTypeCreateContainerHttpHealthCheckDetails)(m),
	}

	return json.Marshal(&s)
}
