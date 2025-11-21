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

// CreateBatchJobDetails The data to create a batch job.
// If the value for a collection is absent or is explicitly provided as null, it will be converted to an empty value, i.e. "[]" or "{}" in json notation. This applies to nested collections as well.
type CreateBatchJobDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent batch job pool.
	BatchJobPoolId *string `mandatory:"true" json:"batchJobPoolId"`

	// A list of tasks to be executed within this job.
	// Maximum number of tasks a job can have in total, across all nested levels, is 1000.
	Tasks []CreateBatchTaskDetails `mandatory:"true" json:"tasks"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// If not specified or provided as null or empty string, it will be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchjob20250914115623.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// An optional description that provides additional context next to the displayName.
	Description *string `mandatory:"false" json:"description"`

	// An environment variables to use for the job's tasks (can be overridden by task's environment variables).
	EnvironmentVariables []EnvironmentVariable `mandatory:"false" json:"environmentVariables"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBatchJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBatchJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateBatchJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName          *string                           `json:"displayName"`
		Description          *string                           `json:"description"`
		EnvironmentVariables []EnvironmentVariable             `json:"environmentVariables"`
		FreeformTags         map[string]string                 `json:"freeformTags"`
		DefinedTags          map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId        *string                           `json:"compartmentId"`
		BatchJobPoolId       *string                           `json:"batchJobPoolId"`
		Tasks                []createbatchtaskdetails          `json:"tasks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.EnvironmentVariables = make([]EnvironmentVariable, len(model.EnvironmentVariables))
	copy(m.EnvironmentVariables, model.EnvironmentVariables)
	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	m.BatchJobPoolId = model.BatchJobPoolId

	m.Tasks = make([]CreateBatchTaskDetails, len(model.Tasks))
	for i, n := range model.Tasks {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Tasks[i] = nn.(CreateBatchTaskDetails)
		} else {
			m.Tasks[i] = nil
		}
	}
	return
}
