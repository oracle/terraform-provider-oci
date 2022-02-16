// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// TriggerBuildPipelineAction The action to trigger a build pipeline.
type TriggerBuildPipelineAction struct {

	// The OCID of the build pipeline to be triggered.
	BuildPipelineId *string `mandatory:"true" json:"buildPipelineId"`

	Filter Filter `mandatory:"false" json:"filter"`
}

//GetFilter returns Filter
func (m TriggerBuildPipelineAction) GetFilter() Filter {
	return m.Filter
}

func (m TriggerBuildPipelineAction) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TriggerBuildPipelineAction) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TriggerBuildPipelineAction) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTriggerBuildPipelineAction TriggerBuildPipelineAction
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeTriggerBuildPipelineAction
	}{
		"TRIGGER_BUILD_PIPELINE",
		(MarshalTypeTriggerBuildPipelineAction)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *TriggerBuildPipelineAction) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Filter          filter  `json:"filter"`
		BuildPipelineId *string `json:"buildPipelineId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Filter.UnmarshalPolymorphicJSON(model.Filter.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Filter = nn.(Filter)
	} else {
		m.Filter = nil
	}

	m.BuildPipelineId = model.BuildPipelineId

	return
}
