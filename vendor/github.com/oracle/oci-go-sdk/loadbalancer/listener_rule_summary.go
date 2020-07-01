// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/common"
)

// ListenerRuleSummary The attributes of a rule associated with the specified listener, and the name of the rule set that the rule
// belongs to.
type ListenerRuleSummary struct {

	// A rule object that applies to the listener.
	Rule Rule `mandatory:"false" json:"rule"`

	// The name of the rule set that the rule belongs to.
	RuleSetName *string `mandatory:"false" json:"ruleSetName"`
}

func (m ListenerRuleSummary) String() string {
	return common.PointerString(m)
}

// UnmarshalJSON unmarshals from json
func (m *ListenerRuleSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Rule        rule    `json:"rule"`
		RuleSetName *string `json:"ruleSetName"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Rule.UnmarshalPolymorphicJSON(model.Rule.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Rule = nn.(Rule)
	} else {
		m.Rule = nil
	}

	m.RuleSetName = model.RuleSetName

	return
}
