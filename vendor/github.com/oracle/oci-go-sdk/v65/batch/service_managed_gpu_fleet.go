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

// ServiceManagedGpuFleet Service managed GPU Fleet configuration of the batch context.
type ServiceManagedGpuFleet struct {

	// Name of the service managed GPU fleet.
	Name *string `mandatory:"true" json:"name"`

	Shape GpuFleetShape `mandatory:"true" json:"shape"`

	// Maximum number of concurrent tasks for the service managed GPU fleet.
	MaxConcurrentTasks *int `mandatory:"true" json:"maxConcurrentTasks"`

	// Current state of the service managed GPU fleet configuration.
	State *string `mandatory:"false" json:"state"`

	// A message that describes the current state of the service managed GPU fleet configuration in more detail.
	Details *string `mandatory:"false" json:"details"`
}

func (m ServiceManagedGpuFleet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceManagedGpuFleet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ServiceManagedGpuFleet) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeServiceManagedGpuFleet ServiceManagedGpuFleet
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeServiceManagedGpuFleet
	}{
		"SERVICE_MANAGED_GPU_FLEET",
		(MarshalTypeServiceManagedGpuFleet)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ServiceManagedGpuFleet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		State              *string       `json:"state"`
		Details            *string       `json:"details"`
		Name               *string       `json:"name"`
		Shape              gpufleetshape `json:"shape"`
		MaxConcurrentTasks *int          `json:"maxConcurrentTasks"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.State = model.State

	m.Details = model.Details

	m.Name = model.Name

	nn, e = model.Shape.UnmarshalPolymorphicJSON(model.Shape.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Shape = nn.(GpuFleetShape)
	} else {
		m.Shape = nil
	}

	m.MaxConcurrentTasks = model.MaxConcurrentTasks

	return
}
