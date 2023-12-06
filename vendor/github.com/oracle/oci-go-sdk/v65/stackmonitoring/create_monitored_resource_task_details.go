// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateMonitoredResourceTaskDetails The request details for the stack monitoring resource task.
type CreateMonitoredResourceTaskDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment identifier.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	TaskDetails MonitoredResourceTaskDetails `mandatory:"true" json:"taskDetails"`

	// Name of the task. If not provided by default the following names will be taken
	// OCI tasks - namespace plus timestamp.
	Name *string `mandatory:"false" json:"name"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateMonitoredResourceTaskDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateMonitoredResourceTaskDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateMonitoredResourceTaskDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name          *string                           `json:"name"`
		FreeformTags  map[string]string                 `json:"freeformTags"`
		DefinedTags   map[string]map[string]interface{} `json:"definedTags"`
		CompartmentId *string                           `json:"compartmentId"`
		TaskDetails   monitoredresourcetaskdetails      `json:"taskDetails"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.CompartmentId = model.CompartmentId

	nn, e = model.TaskDetails.UnmarshalPolymorphicJSON(model.TaskDetails.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.TaskDetails = nn.(MonitoredResourceTaskDetails)
	} else {
		m.TaskDetails = nil
	}

	return
}
