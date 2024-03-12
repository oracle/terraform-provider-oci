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

// ThresholdBasedAutoScalingPolicyDetails Details for a threshold-based autoscaling policy to enable on the model deployment.
// In a threshold-based autoscaling policy, an autoscaling action is triggered when a performance metric meets
// or exceeds a threshold.
type ThresholdBasedAutoScalingPolicyDetails struct {

	// The list of autoscaling policy rules.
	Rules []MetricExpressionRule `mandatory:"true" json:"rules"`

	// For a threshold-based autoscaling policy, this value is the maximum number of instances the model deployment is allowed
	// to increase to (scale out).
	MaximumInstanceCount *int `mandatory:"true" json:"maximumInstanceCount"`

	// For a threshold-based autoscaling policy, this value is the minimum number of instances the model deployment is allowed
	// to decrease to (scale in).
	MinimumInstanceCount *int `mandatory:"true" json:"minimumInstanceCount"`

	// For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the model deployment
	// immediately after autoscaling is enabled. Note that anytime this value is updated, the number of instances will be reset
	// to this value. After autoscaling retrieves performance metrics, the number of instances is automatically adjusted from
	// this initial number to a number that is based on the limits that you set.
	InitialInstanceCount *int `mandatory:"true" json:"initialInstanceCount"`
}

func (m ThresholdBasedAutoScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ThresholdBasedAutoScalingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ThresholdBasedAutoScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeThresholdBasedAutoScalingPolicyDetails ThresholdBasedAutoScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"autoScalingPolicyType"`
		MarshalTypeThresholdBasedAutoScalingPolicyDetails
	}{
		"THRESHOLD",
		(MarshalTypeThresholdBasedAutoScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ThresholdBasedAutoScalingPolicyDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Rules                []metricexpressionrule `json:"rules"`
		MaximumInstanceCount *int                   `json:"maximumInstanceCount"`
		MinimumInstanceCount *int                   `json:"minimumInstanceCount"`
		InitialInstanceCount *int                   `json:"initialInstanceCount"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Rules = make([]MetricExpressionRule, len(model.Rules))
	for i, n := range model.Rules {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.Rules[i] = nn.(MetricExpressionRule)
		} else {
			m.Rules[i] = nil
		}
	}
	m.MaximumInstanceCount = model.MaximumInstanceCount

	m.MinimumInstanceCount = model.MinimumInstanceCount

	m.InitialInstanceCount = model.InitialInstanceCount

	return
}
