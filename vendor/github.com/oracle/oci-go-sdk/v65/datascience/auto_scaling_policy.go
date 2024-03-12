// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AutoScalingPolicy The scaling policy to enable autoscaling on the model deployment.
type AutoScalingPolicy struct {

	// The list of autoscaling policy details.
	AutoScalingPolicies []AutoScalingPolicyDetails `mandatory:"true" json:"autoScalingPolicies"`

	// For threshold-based autoscaling policies, this value is the minimum period of time to wait between scaling actions.
	// The cooldown period gives the system time to stabilize before rescaling. The minimum value is 600 seconds, which
	// is also the default. The cooldown period starts when the model deployment becomes ACTIVE after the scaling operation.
	CoolDownInSeconds *int `mandatory:"false" json:"coolDownInSeconds"`

	// Whether the autoscaling policy is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m AutoScalingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AutoScalingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AutoScalingPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAutoScalingPolicy AutoScalingPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeAutoScalingPolicy
	}{
		"AUTOSCALING",
		(MarshalTypeAutoScalingPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *AutoScalingPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		CoolDownInSeconds   *int                       `json:"coolDownInSeconds"`
		IsEnabled           *bool                      `json:"isEnabled"`
		AutoScalingPolicies []autoscalingpolicydetails `json:"autoScalingPolicies"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.CoolDownInSeconds = model.CoolDownInSeconds

	m.IsEnabled = model.IsEnabled

	m.AutoScalingPolicies = make([]AutoScalingPolicyDetails, len(model.AutoScalingPolicies))
	for i, n := range model.AutoScalingPolicies {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.AutoScalingPolicies[i] = nn.(AutoScalingPolicyDetails)
		} else {
			m.AutoScalingPolicies[i] = nil
		}
	}
	return
}
