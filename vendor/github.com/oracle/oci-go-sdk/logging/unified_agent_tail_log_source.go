// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// loggingManagementControlplane API
//
// loggingManagementControlplane API specification
//

package logging

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// UnifiedAgentTailLogSource tail log source object.
type UnifiedAgentTailLogSource struct {

	// unique name for the source
	Name *string `mandatory:"true" json:"name"`

	Paths []string `mandatory:"false" json:"paths"`

	Parser UnifiedAgentParser `mandatory:"false" json:"parser"`
}

//GetName returns Name
func (m UnifiedAgentTailLogSource) GetName() *string {
	return m.Name
}

func (m UnifiedAgentTailLogSource) String() string {
	return common.PointerString(m)
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
		Paths  []string           `json:"paths"`
		Parser unifiedagentparser `json:"parser"`
		Name   *string            `json:"name"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Paths = make([]string, len(model.Paths))
	for i, n := range model.Paths {
		m.Paths[i] = n
	}

	nn, e = model.Parser.UnmarshalPolymorphicJSON(model.Parser.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Parser = nn.(UnifiedAgentParser)
	} else {
		m.Parser = nil
	}

	m.Name = model.Name

	return
}
