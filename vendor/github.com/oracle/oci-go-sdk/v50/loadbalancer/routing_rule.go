// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v50/common"
)

// RoutingRule A routing rule examines an incoming request, routing matching requests to the specified backend set.
// Routing rules apply only to HTTP and HTTPS requests. They have no effect on TCP requests.
type RoutingRule struct {

	// A unique name for the routing policy rule. Avoid entering confidential information.
	Name *string `mandatory:"true" json:"name"`

	// A routing rule to evaluate defined conditions against the incoming HTTP request and perform an action.
	Condition *string `mandatory:"true" json:"condition"`

	// A list of actions to be applied when conditions of the routing rule are met.
	Actions []Action `mandatory:"true" json:"actions"`
}

func (m RoutingRule) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *RoutingRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name      *string  `json:"name"`
		Condition *string  `json:"condition"`
		Actions   []action `json:"actions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Condition = model.Condition

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
