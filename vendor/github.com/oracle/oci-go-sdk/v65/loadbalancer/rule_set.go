// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuleSet A named set of rules associated with a load balancer. Rules are objects that represent actions to apply to a listener,
// such as adding, altering, or removing HTTP headers. For more information, see
// Managing Rule Sets (https://docs.oracle.com/iaas/Content/Balance/Tasks/managingrulesets.htm).
type RuleSet struct {

	// The name for this set of rules. It must be unique and it cannot be changed. Avoid entering
	// confidential information.
	// Example: `example_rule_set`
	Name *string `mandatory:"true" json:"name"`

	// An array of rules that compose the rule set.
	Items []Rule `mandatory:"true" json:"items"`
}

func (m RuleSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *RuleSet) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Name  *string `json:"name"`
		Items []rule  `json:"items"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Name = model.Name

	m.Items = make([]Rule, len(model.Items))
	for i, n := range model.Items {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Items[i] = nn.(Rule)
		} else {
			m.Items[i] = nil
		}
	}
	return
}
