// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// CreateServiceManagedGpuFleetDetails Service managed GPU Fleet configuration of the batch context.
type CreateServiceManagedGpuFleetDetails struct {

	// Name of the service managed GPU fleet.
	Name *string `mandatory:"true" json:"name"`

	Shape CreateGpuFleetShapeDetails `mandatory:"true" json:"shape"`

	// Maximum number of concurrent tasks for the service managed GPU fleet.
	MaxConcurrentTasks *int `mandatory:"true" json:"maxConcurrentTasks"`
}

func (m CreateServiceManagedGpuFleetDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateServiceManagedGpuFleetDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateServiceManagedGpuFleetDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateServiceManagedGpuFleetDetails CreateServiceManagedGpuFleetDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeCreateServiceManagedGpuFleetDetails
	}{
		"SERVICE_MANAGED_GPU_FLEET",
		(MarshalTypeCreateServiceManagedGpuFleetDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateServiceManagedGpuFleetDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name               *string                    `json:"name"`
		Shape              creategpufleetshapedetails `json:"shape"`
		MaxConcurrentTasks *int                       `json:"maxConcurrentTasks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	nn, e = model.Shape.UnmarshalPolymorphicJSON(model.Shape.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Shape = nn.(CreateGpuFleetShapeDetails)
	} else {
		m.Shape = nil
	}

	m.MaxConcurrentTasks = model.MaxConcurrentTasks

	return
}
