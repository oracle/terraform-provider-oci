// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Logging Management API
//
// Use the Logging Management API to create, read, list, update, move and delete
// log groups, log objects, log saved searches, and agent configurations.
// For more information, see Logging Overview (https://docs.oracle.com/iaas/Content/Logging/Concepts/loggingoverview.htm).
//

package logging

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogPipelineDestinationResource Represents a log pipeline destination resource.
type LogPipelineDestinationResource struct {

	// The OCID of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment that the resource belongs to.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	Destination LogPipelineDestinationResponse `mandatory:"true" json:"destination"`

	// The pipeline state.
	LifecycleState LogPipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Time the resource was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time the resource was last modified.
	TimeLastModified *common.SDKTime `mandatory:"true" json:"timeLastModified"`
}

func (m LogPipelineDestinationResource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogPipelineDestinationResource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLogPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLogPipelineLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *LogPipelineDestinationResource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Id               *string                        `json:"id"`
		CompartmentId    *string                        `json:"compartmentId"`
		Destination      logpipelinedestinationresponse `json:"destination"`
		LifecycleState   LogPipelineLifecycleStateEnum  `json:"lifecycleState"`
		TimeCreated      *common.SDKTime                `json:"timeCreated"`
		TimeLastModified *common.SDKTime                `json:"timeLastModified"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	nn, e = model.Destination.UnmarshalPolymorphicJSON(model.Destination.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Destination = nn.(LogPipelineDestinationResponse)
	} else {
		m.Destination = nil
	}

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	m.TimeLastModified = model.TimeLastModified

	return
}
