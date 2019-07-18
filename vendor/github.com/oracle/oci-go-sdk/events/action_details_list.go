// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Events API
//
// API for the Events Service. Use this API to manage rules and actions that create automation
// in your tenancy. For more information, see Overview of Events (https://docs.cloud.oracle.com/iaas/Content/Events/Concepts/eventsoverview.htm).
//

package events

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// ActionDetailsList A list of ActionDetails objects to create for a rule.
type ActionDetailsList struct {

	// A list of one or more ActionDetails objects.
	Actions []ActionDetails `mandatory:"true" json:"actions"`
}

func (m ActionDetailsList) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ActionDetailsList) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Actions []actiondetails `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	m.Actions = make([]ActionDetails, len(model.Actions))
	for i, n := range model.Actions {
		nn, err := n.UnmarshalPolymorphicJSON(n.JsonData)
		if err != nil {
			return err
		}
		if nn != nil {
			m.Actions[i] = nn.(ActionDetails)
		} else {
			m.Actions[i] = nil
		}
	}
	return
}
