// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// UnifiedAgentTailLogSource Tail log source object.
type UnifiedAgentTailLogSource struct {

	// Unique name for the source.
	Name *string `mandatory:"true" json:"name"`

	// Absolute paths for log source files. Wildcards can be used.
	Paths []string `mandatory:"true" json:"paths"`

	Parser UnifiedAgentParser `mandatory:"false" json:"parser"`

	AdvancedOptions *UnifiedAgentTailSourceAdvancedOptions `mandatory:"false" json:"advancedOptions"`
}

// GetName returns Name
func (m UnifiedAgentTailLogSource) GetName() *string {
	return m.Name
}

func (m UnifiedAgentTailLogSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentTailLogSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentTailLogSource) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentTailLogSource UnifiedAgentTailLogSource
	s := struct {
		DiscriminatorParam string `json:"sourceType"`
		MarshalTypeUnifiedAgentTailLogSource
	}{
		"LOG_TAIL",
		(MarshalTypeUnifiedAgentTailLogSource)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAgentTailLogSource) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Parser          unifiedagentparser                     `json:"parser"`
		AdvancedOptions *UnifiedAgentTailSourceAdvancedOptions `json:"advancedOptions"`
		Name            *string                                `json:"name"`
		Paths           []string                               `json:"paths"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Parser.UnmarshalPolymorphicJSON(model.Parser.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Parser = nn.(UnifiedAgentParser)
	} else {
		m.Parser = nil
	}

	m.AdvancedOptions = model.AdvancedOptions

	m.Name = model.Name

	m.Paths = make([]string, len(model.Paths))
	copy(m.Paths, model.Paths)
	return
}
