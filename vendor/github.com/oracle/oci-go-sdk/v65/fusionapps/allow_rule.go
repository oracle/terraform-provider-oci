// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AllowRule An object that represents the action of configuring an access control rule. Access control rules permit access
// to application resources based on user-specified match conditions. This rule applies only to HTTP listeners.
// **NOTES:**
// *  If you do not specify any access control rules, the default rule is to allow all traffic.
// *  If you add access control rules, the load balancer denies any traffic that does not match the rules.
// *  Maximum of two match conditions can be specified in a rule.
// *  You can specify this rule only with the following `RuleCondition` combinations:
// *  `SOURCE_IP_ADDRESS`
// *  `SOURCE_VCN_ID`
// *  `SOURCE_VCN_ID", "SOURCE_VCN_IP_ADDRESS`
type AllowRule struct {
	Conditions []RuleCondition `mandatory:"true" json:"conditions"`

	// A brief description of the access control rule. Avoid entering confidential information.
	// example: `192.168.0.0/16 and 2001:db8::/32 are trusted clients. Whitelist them.`
	Description *string `mandatory:"false" json:"description"`
}

func (m AllowRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AllowRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AllowRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAllowRule AllowRule
	s := struct {
		DiscriminatorParam string `json:"action"`
		MarshalTypeAllowRule
	}{
		"ALLOW",
		(MarshalTypeAllowRule)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AllowRule) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description *string         `json:"description"`
		Conditions  []rulecondition `json:"conditions"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.Conditions = make([]RuleCondition, len(model.Conditions))
	for i, n := range model.Conditions {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Conditions[i] = nn.(RuleCondition)
		} else {
			m.Conditions[i] = nil
		}
	}
	return
}
