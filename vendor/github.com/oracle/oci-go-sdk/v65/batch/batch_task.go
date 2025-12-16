// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BatchTask A batch task contains common properties for all types of job tasks.
type BatchTask interface {

	// The UUID of batch task.
	GetId() *string

	// The name of the batch task. It must be unique within its parent batch job.
	GetName() *string

	// A list of resources (for example licences) this task needs for its execution.
	GetEntitlementClaims() []string

	// A list of tasks from the same job this task depends on referenced by name.
	GetDependencies() []string

	// Environment variables to use for the task execution.
	GetEnvironmentVariables() []EnvironmentVariable

	// An optional description that provides additional context next to the displayName.
	GetDescription() *string

	// The current state of the batch task.
	GetLifecycleState() BatchTaskLifecycleStateEnum

	// A message that describes the current state of the batch task in more detail. For example, can be used to provide actionable information for a resource in the Failed state.
	GetLifecycleDetails() *string
}

type batchtask struct {
	JsonData             []byte
	Description          *string                     `mandatory:"false" json:"description"`
	LifecycleState       BatchTaskLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
	LifecycleDetails     *string                     `mandatory:"false" json:"lifecycleDetails"`
	Id                   *string                     `mandatory:"true" json:"id"`
	Name                 *string                     `mandatory:"true" json:"name"`
	EntitlementClaims    []string                    `mandatory:"true" json:"entitlementClaims"`
	Dependencies         []string                    `mandatory:"true" json:"dependencies"`
	EnvironmentVariables []EnvironmentVariable       `mandatory:"true" json:"environmentVariables"`
	Type                 string                      `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *batchtask) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerbatchtask batchtask
	s := struct {
		Model Unmarshalerbatchtask
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Id = s.Model.Id
	m.Name = s.Model.Name
	m.EntitlementClaims = s.Model.EntitlementClaims
	m.Dependencies = s.Model.Dependencies
	m.EnvironmentVariables = s.Model.EnvironmentVariables
	m.Description = s.Model.Description
	m.LifecycleState = s.Model.LifecycleState
	m.LifecycleDetails = s.Model.LifecycleDetails
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *batchtask) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE":
		mm := ComputeTask{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for BatchTask: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m batchtask) GetDescription() *string {
	return m.Description
}

// GetLifecycleState returns LifecycleState
func (m batchtask) GetLifecycleState() BatchTaskLifecycleStateEnum {
	return m.LifecycleState
}

// GetLifecycleDetails returns LifecycleDetails
func (m batchtask) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

// GetId returns Id
func (m batchtask) GetId() *string {
	return m.Id
}

// GetName returns Name
func (m batchtask) GetName() *string {
	return m.Name
}

// GetEntitlementClaims returns EntitlementClaims
func (m batchtask) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m batchtask) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m batchtask) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

func (m batchtask) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m batchtask) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingBatchTaskLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetBatchTaskLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BatchTaskLifecycleStateEnum Enum with underlying type: string
type BatchTaskLifecycleStateEnum string

// Set of constants representing the allowable values for BatchTaskLifecycleStateEnum
const (
	BatchTaskLifecycleStateAccepted   BatchTaskLifecycleStateEnum = "ACCEPTED"
	BatchTaskLifecycleStateWaiting    BatchTaskLifecycleStateEnum = "WAITING"
	BatchTaskLifecycleStateInProgress BatchTaskLifecycleStateEnum = "IN_PROGRESS"
	BatchTaskLifecycleStateSucceeded  BatchTaskLifecycleStateEnum = "SUCCEEDED"
	BatchTaskLifecycleStateFailed     BatchTaskLifecycleStateEnum = "FAILED"
	BatchTaskLifecycleStateCanceling  BatchTaskLifecycleStateEnum = "CANCELING"
	BatchTaskLifecycleStateCanceled   BatchTaskLifecycleStateEnum = "CANCELED"
)

var mappingBatchTaskLifecycleStateEnum = map[string]BatchTaskLifecycleStateEnum{
	"ACCEPTED":    BatchTaskLifecycleStateAccepted,
	"WAITING":     BatchTaskLifecycleStateWaiting,
	"IN_PROGRESS": BatchTaskLifecycleStateInProgress,
	"SUCCEEDED":   BatchTaskLifecycleStateSucceeded,
	"FAILED":      BatchTaskLifecycleStateFailed,
	"CANCELING":   BatchTaskLifecycleStateCanceling,
	"CANCELED":    BatchTaskLifecycleStateCanceled,
}

var mappingBatchTaskLifecycleStateEnumLowerCase = map[string]BatchTaskLifecycleStateEnum{
	"accepted":    BatchTaskLifecycleStateAccepted,
	"waiting":     BatchTaskLifecycleStateWaiting,
	"in_progress": BatchTaskLifecycleStateInProgress,
	"succeeded":   BatchTaskLifecycleStateSucceeded,
	"failed":      BatchTaskLifecycleStateFailed,
	"canceling":   BatchTaskLifecycleStateCanceling,
	"canceled":    BatchTaskLifecycleStateCanceled,
}

// GetBatchTaskLifecycleStateEnumValues Enumerates the set of values for BatchTaskLifecycleStateEnum
func GetBatchTaskLifecycleStateEnumValues() []BatchTaskLifecycleStateEnum {
	values := make([]BatchTaskLifecycleStateEnum, 0)
	for _, v := range mappingBatchTaskLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskLifecycleStateEnumStringValues Enumerates the set of values in String for BatchTaskLifecycleStateEnum
func GetBatchTaskLifecycleStateEnumStringValues() []string {
	return []string{
		"ACCEPTED",
		"WAITING",
		"IN_PROGRESS",
		"SUCCEEDED",
		"FAILED",
		"CANCELING",
		"CANCELED",
	}
}

// GetMappingBatchTaskLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskLifecycleStateEnum(val string) (BatchTaskLifecycleStateEnum, bool) {
	enum, ok := mappingBatchTaskLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// BatchTaskTypeEnum Enum with underlying type: string
type BatchTaskTypeEnum string

// Set of constants representing the allowable values for BatchTaskTypeEnum
const (
	BatchTaskTypeCompute BatchTaskTypeEnum = "COMPUTE"
)

var mappingBatchTaskTypeEnum = map[string]BatchTaskTypeEnum{
	"COMPUTE": BatchTaskTypeCompute,
}

var mappingBatchTaskTypeEnumLowerCase = map[string]BatchTaskTypeEnum{
	"compute": BatchTaskTypeCompute,
}

// GetBatchTaskTypeEnumValues Enumerates the set of values for BatchTaskTypeEnum
func GetBatchTaskTypeEnumValues() []BatchTaskTypeEnum {
	values := make([]BatchTaskTypeEnum, 0)
	for _, v := range mappingBatchTaskTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetBatchTaskTypeEnumStringValues Enumerates the set of values in String for BatchTaskTypeEnum
func GetBatchTaskTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingBatchTaskTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBatchTaskTypeEnum(val string) (BatchTaskTypeEnum, bool) {
	enum, ok := mappingBatchTaskTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
