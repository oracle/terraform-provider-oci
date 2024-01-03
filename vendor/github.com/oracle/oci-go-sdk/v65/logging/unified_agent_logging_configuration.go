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

// UnifiedAgentLoggingConfiguration Unified Agent logging service configuration object.
type UnifiedAgentLoggingConfiguration struct {

	// Logging source object.
	Sources []UnifiedAgentLoggingSource `mandatory:"true" json:"sources"`

	Destination *UnifiedAgentLoggingDestination `mandatory:"true" json:"destination"`

	// Logging filter object.
	Filter []UnifiedAgentLoggingFilter `mandatory:"false" json:"filter"`
}

func (m UnifiedAgentLoggingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnifiedAgentLoggingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
		Filter      []unifiedagentloggingfilter     `json:"filter"`
		Sources     []unifiedagentloggingsource     `json:"sources"`
		Destination *UnifiedAgentLoggingDestination `json:"destination"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Filter = make([]UnifiedAgentLoggingFilter, len(model.Filter))
	for i, n := range model.Filter {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Filter[i] = nn.(UnifiedAgentLoggingFilter)
		} else {
			m.Filter[i] = nil
		}
	}
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
