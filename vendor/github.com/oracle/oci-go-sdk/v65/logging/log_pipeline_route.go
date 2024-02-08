// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.cloud.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineRoute Describes one Route within the Pipeline resource. All Routes process the same log events from the Pipeline inputs.
// Each Route processes events using the list of Functions and forwards the resulting data to configured Destinations.
type LogPipelineRoute struct {

	// List of destinations for the pipeline.
	Destinations []LogPipelineDestination `mandatory:"true" json:"destinations"`

	// Name of Log Pipeline Route.
	Name *string `mandatory:"false" json:"name"`

	// A list of Log Pipeline functions.
	Functions []LogPipelineFunction `mandatory:"false" json:"functions"`
}

func (m LogPipelineRoute) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineRoute) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LogPipelineRoute) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name         *string                  `json:"name"`
		Functions    []logpipelinefunction    `json:"functions"`
		Destinations []logpipelinedestination `json:"destinations"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Functions = make([]LogPipelineFunction, len(model.Functions))
	for i, n := range model.Functions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Functions[i] = nn.(LogPipelineFunction)
		} else {
			m.Functions[i] = nil
		}
	}
	m.Destinations = make([]LogPipelineDestination, len(model.Destinations))
	for i, n := range model.Destinations {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Destinations[i] = nn.(LogPipelineDestination)
		} else {
			m.Destinations[i] = nil
		}
	}
	return
}
