// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ActionList A list of Action objects associated with a rule.
type ActionList struct {

	// A list of one or more Action objects.
	Actions []Action `mandatory:"true" json:"actions"`
}

func (m ActionList) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActionList) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *ActionList) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions []action `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Actions = make([]Action, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(Action)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
