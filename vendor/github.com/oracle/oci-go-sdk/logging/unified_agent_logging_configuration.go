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

// UnifiedAgentLoggingConfiguration Unified Agent logging service configuration object.
type UnifiedAgentLoggingConfiguration struct {
	Sources []UnifiedAgentLoggingSource `mandatory:"false" json:"sources"`

	Destination *UnifiedAgentLoggingDestination `mandatory:"false" json:"destination"`
}

func (m UnifiedAgentLoggingConfiguration) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UnifiedAgentLoggingConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUnifiedAgentLoggingConfiguration UnifiedAgentLoggingConfiguration
	s := struct {
		DiscriminatorParam string `json:"configurationType"`
		MarshalTypeUnifiedAgentLoggingConfiguration
	}{
		"LOGGING",
		(MarshalTypeUnifiedAgentLoggingConfiguration)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UnifiedAgentLoggingConfiguration) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Sources     []unifiedagentloggingsource     `json:"sources"`
		Destination *UnifiedAgentLoggingDestination `json:"destination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Sources = make([]UnifiedAgentLoggingSource, len(model.Sources))
	for i, n := range model.Sources {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Sources[i] = nn.(UnifiedAgentLoggingSource)
		} else {
			m.Sources[i] = nil
		}
	}

	m.Destination = model.Destination

	return
}
