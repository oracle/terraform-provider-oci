// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// ManagedComputeClusterWorkloadAutoScalingPolicy The scaling policy to enable autoscaling of workload on the managed compute cluster type compute target .
type ManagedComputeClusterWorkloadAutoScalingPolicy struct {

	// The list of autoscaling policy details.
	AutoScalingPolicies []ManagedComputeClusterWorkloadAutoScalingPolicyDetails `mandatory:"true" json:"autoScalingPolicies"`

	// Whether the autoscaling policy is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`
}

func (m ManagedComputeClusterWorkloadAutoScalingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedComputeClusterWorkloadAutoScalingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedComputeClusterWorkloadAutoScalingPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedComputeClusterWorkloadAutoScalingPolicy ManagedComputeClusterWorkloadAutoScalingPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeManagedComputeClusterWorkloadAutoScalingPolicy
	}{
		"AUTOSCALING",
		(MarshalTypeManagedComputeClusterWorkloadAutoScalingPolicy)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ManagedComputeClusterWorkloadAutoScalingPolicy) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		IsEnabled           *bool                                                   `json:"isEnabled"`
		AutoScalingPolicies []managedcomputeclusterworkloadautoscalingpolicydetails `json:"autoScalingPolicies"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.IsEnabled = model.IsEnabled

	m.AutoScalingPolicies = make([]ManagedComputeClusterWorkloadAutoScalingPolicyDetails, len(model.AutoScalingPolicies))
	for i, n := range model.AutoScalingPolicies {
		nn, e = n.UnmarshalPolymorphicJSON(n.JsonData)
		if e != nil {
			return e
		}
		if nn != nil {
			m.AutoScalingPolicies[i] = nn.(ManagedComputeClusterWorkloadAutoScalingPolicyDetails)
		} else {
			m.AutoScalingPolicies[i] = nil
		}
	}
	return
}
