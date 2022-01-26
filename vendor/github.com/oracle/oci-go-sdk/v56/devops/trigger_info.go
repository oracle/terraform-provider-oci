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
	"github.com/oracle/oci-go-sdk/v56/common"
)

// TriggerInfo Trigger details that need to be used for the BuildRun
type TriggerInfo struct {

	// The list of actions that are to be performed for this Trigger
	Actions []TriggerAction `mandatory:"true" json:"actions"`

	// Name for Trigger.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m TriggerInfo) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *TriggerInfo) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		DisplayName *string         `json:"displayName"`
		Actions     []triggeraction `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.DisplayName = model.DisplayName

	m.Actions = make([]TriggerAction, len(model.Actions))
	for i, n := range model.Actions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Actions[i] = nn.(TriggerAction)
		} else {
			m.Actions[i] = nil
		}
	}

	return
}
