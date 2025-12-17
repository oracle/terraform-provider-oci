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

// CreateBatchContextDetails The data to create a batch context.
// If the value for a collection is absent or is explicitly provided as null, it will be converted to an empty value, i.e. "[]" or "{}" in json notation. This applies to nested collections as well.
type CreateBatchContextDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Network *CreateNetworkDetails `mandatory:"true" json:"network"`

	// List of fleet configurations related to the batch context.
	Fleets []CreateFleetDetails `mandatory:"true" json:"fleets"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// If not specified or provided as null or empty string, it will be generated as "<resourceType><timeCreated>", where timeCreated corresponds with the resource creation time in ISO 8601 basic format, i.e. omitting separating punctuation, at second-level precision and no UTC offset. Example: batchcontext20250914115623.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Summarized information about the batch context.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// List of job priority configurations related to the batch context.
	JobPriorityConfigurations []JobPriorityConfiguration `mandatory:"false" json:"jobPriorityConfigurations"`

	// Mapping of concurrent/shared resources used in job tasks to their limits.
	Entitlements map[string]int `mandatory:"false" json:"entitlements"`

	LoggingConfiguration LoggingConfiguration `mandatory:"false" json:"loggingConfiguration"`
}

func (m CreateBatchContextDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateBatchContextDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateBatchContextDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName               *string                           `json:"displayName"`
		Description               *string                           `json:"description"`
		FreeformTags              map[string]string                 `json:"freeformTags"`
		DefinedTags               map[string]map[string]interface{} `json:"definedTags"`
		JobPriorityConfigurations []JobPriorityConfiguration        `json:"jobPriorityConfigurations"`
		Entitlements              map[string]int                    `json:"entitlements"`
		LoggingConfiguration      loggingconfiguration              `json:"loggingConfiguration"`
		CompartmentId             *string                           `json:"compartmentId"`
		Network                   *CreateNetworkDetails             `json:"network"`
		Fleets                    []createfleetdetails              `json:"fleets"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Description = model.Description

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.JobPriorityConfigurations = make([]JobPriorityConfiguration, len(model.JobPriorityConfigurations))
	copy(m.JobPriorityConfigurations, model.JobPriorityConfigurations)
	m.Entitlements = model.Entitlements

	nn, e = model.LoggingConfiguration.UnmarshalPolymorphicJSON(model.LoggingConfiguration.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.LoggingConfiguration = nn.(LoggingConfiguration)
	} else {
		m.LoggingConfiguration = nil
	}

	m.CompartmentId = model.CompartmentId

	m.Network = model.Network

	m.Fleets = make([]CreateFleetDetails, len(model.Fleets))
	for i, n := range model.Fleets {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Fleets[i] = nn.(CreateFleetDetails)
		} else {
			m.Fleets[i] = nil
		}
	}
	return
}
