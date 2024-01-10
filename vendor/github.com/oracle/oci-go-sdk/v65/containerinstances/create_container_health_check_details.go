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

// CreateContainerHealthCheckDetails Container Health Check is used to check and report the status of a container.
type CreateContainerHealthCheckDetails interface {

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

	// The action will be triggered when the container health check fails. There are two types of action: KILL or NONE. The default
	// action is KILL. If failure action is KILL, the container will be subject to the container restart policy.
	GetFailureAction() ContainerHealthCheckFailureActionEnum
}

type createcontainerhealthcheckdetails struct {
	JsonData              []byte
	Name                  *string                               `mandatory:"false" json:"name"`
	InitialDelayInSeconds *int                                  `mandatory:"false" json:"initialDelayInSeconds"`
	IntervalInSeconds     *int                                  `mandatory:"false" json:"intervalInSeconds"`
	FailureThreshold      *int                                  `mandatory:"false" json:"failureThreshold"`
	SuccessThreshold      *int                                  `mandatory:"false" json:"successThreshold"`
	TimeoutInSeconds      *int                                  `mandatory:"false" json:"timeoutInSeconds"`
	FailureAction         ContainerHealthCheckFailureActionEnum `mandatory:"false" json:"failureAction,omitempty"`
	HealthCheckType       string                                `json:"healthCheckType"`
}

// UnmarshalJSON unmarshals json
func (m *createcontainerhealthcheckdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatecontainerhealthcheckdetails createcontainerhealthcheckdetails
	s := struct {
		Model Unmarshalercreatecontainerhealthcheckdetails
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
	m.FailureAction = s.Model.FailureAction
	m.HealthCheckType = s.Model.HealthCheckType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createcontainerhealthcheckdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.HealthCheckType {
	case "TCP":
		mm := CreateContainerTcpHealthCheckDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "HTTP":
		mm := CreateContainerHttpHealthCheckDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMMAND":
		mm := CreateContainerCommandHealthCheckDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateContainerHealthCheckDetails: %s.", m.HealthCheckType)
		return *m, nil
	}
}

// GetName returns Name
func (m createcontainerhealthcheckdetails) GetName() *string {
	return m.Name
}

// GetInitialDelayInSeconds returns InitialDelayInSeconds
func (m createcontainerhealthcheckdetails) GetInitialDelayInSeconds() *int {
	return m.InitialDelayInSeconds
}

// GetIntervalInSeconds returns IntervalInSeconds
func (m createcontainerhealthcheckdetails) GetIntervalInSeconds() *int {
	return m.IntervalInSeconds
}

// GetFailureThreshold returns FailureThreshold
func (m createcontainerhealthcheckdetails) GetFailureThreshold() *int {
	return m.FailureThreshold
}

// GetSuccessThreshold returns SuccessThreshold
func (m createcontainerhealthcheckdetails) GetSuccessThreshold() *int {
	return m.SuccessThreshold
}

// GetTimeoutInSeconds returns TimeoutInSeconds
func (m createcontainerhealthcheckdetails) GetTimeoutInSeconds() *int {
	return m.TimeoutInSeconds
}

// GetFailureAction returns FailureAction
func (m createcontainerhealthcheckdetails) GetFailureAction() ContainerHealthCheckFailureActionEnum {
	return m.FailureAction
}

func (m createcontainerhealthcheckdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createcontainerhealthcheckdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingContainerHealthCheckFailureActionEnum(string(m.FailureAction)); !ok && m.FailureAction != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FailureAction: %s. Supported values are: %s.", m.FailureAction, strings.Join(GetContainerHealthCheckFailureActionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
