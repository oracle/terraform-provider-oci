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

// CreateBatchTaskDetails A batch task contains common properties for all types of job tasks.
type CreateBatchTaskDetails interface {

	// The name of the batch task. It must be unique within its parent batch job.
	GetName() *string

	// An optional description that provides additional context next to the displayName.
	GetDescription() *string

	// A list of resources (for example licences) this task needs for its execution.
	GetEntitlementClaims() []string

	// A list of tasks from the same job this task depends on referenced by name.
	GetDependencies() []string

	// Environment variables to use for the task execution.
	GetEnvironmentVariables() []EnvironmentVariable
}

type createbatchtaskdetails struct {
	JsonData             []byte
	Description          *string               `mandatory:"false" json:"description"`
	EntitlementClaims    []string              `mandatory:"false" json:"entitlementClaims"`
	Dependencies         []string              `mandatory:"false" json:"dependencies"`
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`
	Name                 *string               `mandatory:"true" json:"name"`
	Type                 string                `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createbatchtaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatebatchtaskdetails createbatchtaskdetails
	s := struct {
		Model Unmarshalercreatebatchtaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.EntitlementClaims = s.Model.EntitlementClaims
	m.Dependencies = s.Model.Dependencies
	m.EnvironmentVariables = s.Model.EnvironmentVariables
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createbatchtaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "COMPUTE":
		mm := CreateComputeTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateBatchTaskDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createbatchtaskdetails) GetDescription() *string {
	return m.Description
}

// GetEntitlementClaims returns EntitlementClaims
func (m createbatchtaskdetails) GetEntitlementClaims() []string {
	return m.EntitlementClaims
}

// GetDependencies returns Dependencies
func (m createbatchtaskdetails) GetDependencies() []string {
	return m.Dependencies
}

// GetEnvironmentVariables returns EnvironmentVariables
func (m createbatchtaskdetails) GetEnvironmentVariables() []EnvironmentVariable {
	return m.EnvironmentVariables
}

// GetName returns Name
func (m createbatchtaskdetails) GetName() *string {
	return m.Name
}

func (m createbatchtaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createbatchtaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateBatchTaskDetailsTypeEnum Enum with underlying type: string
type CreateBatchTaskDetailsTypeEnum string

// Set of constants representing the allowable values for CreateBatchTaskDetailsTypeEnum
const (
	CreateBatchTaskDetailsTypeCompute CreateBatchTaskDetailsTypeEnum = "COMPUTE"
)

var mappingCreateBatchTaskDetailsTypeEnum = map[string]CreateBatchTaskDetailsTypeEnum{
	"COMPUTE": CreateBatchTaskDetailsTypeCompute,
}

var mappingCreateBatchTaskDetailsTypeEnumLowerCase = map[string]CreateBatchTaskDetailsTypeEnum{
	"compute": CreateBatchTaskDetailsTypeCompute,
}

// GetCreateBatchTaskDetailsTypeEnumValues Enumerates the set of values for CreateBatchTaskDetailsTypeEnum
func GetCreateBatchTaskDetailsTypeEnumValues() []CreateBatchTaskDetailsTypeEnum {
	values := make([]CreateBatchTaskDetailsTypeEnum, 0)
	for _, v := range mappingCreateBatchTaskDetailsTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateBatchTaskDetailsTypeEnumStringValues Enumerates the set of values in String for CreateBatchTaskDetailsTypeEnum
func GetCreateBatchTaskDetailsTypeEnumStringValues() []string {
	return []string{
		"COMPUTE",
	}
}

// GetMappingCreateBatchTaskDetailsTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateBatchTaskDetailsTypeEnum(val string) (CreateBatchTaskDetailsTypeEnum, bool) {
	enum, ok := mappingCreateBatchTaskDetailsTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
