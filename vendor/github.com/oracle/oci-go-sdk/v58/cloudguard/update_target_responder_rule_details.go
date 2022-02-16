// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Guard APIs
//
// A description of the Cloud Guard APIs
//

package cloudguard

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateTargetResponderRuleDetails Details of ResponderRule.
type UpdateTargetResponderRuleDetails struct {
	Condition Condition `mandatory:"false" json:"condition"`

	// Configurations associated with the ResponderRule
	Configurations []ResponderConfiguration `mandatory:"false" json:"configurations"`

	// Execution Mode for ResponderRule
	Mode ResponderModeTypesEnum `mandatory:"false" json:"mode,omitempty"`
}

func (m UpdateTargetResponderRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateTargetResponderRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingResponderModeTypesEnum(string(m.Mode)); !ok && m.Mode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Mode: %s. Supported values are: %s.", m.Mode, strings.Join(GetResponderModeTypesEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *UpdateTargetResponderRuleDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Condition      condition                `json:"condition"`
		Configurations []ResponderConfiguration `json:"configurations"`
		Mode           ResponderModeTypesEnum   `json:"mode"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Condition.UnmarshalPolymorphicJSON(model.Condition.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Condition = nn.(Condition)
	} else {
		m.Condition = nil
	}

	m.Configurations = make([]ResponderConfiguration, len(model.Configurations))
	for i, n := range model.Configurations {
		m.Configurations[i] = n
	}

	m.Mode = model.Mode

	return
}
