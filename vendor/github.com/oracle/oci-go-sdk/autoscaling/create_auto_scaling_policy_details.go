// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/common"
)

// CreateAutoScalingPolicyDetails An AutoScalingConfiguration Policy creation details
type CreateAutoScalingPolicyDetails interface {

	// The capacity requirements of the Policy
	GetCapacity() *Capacity

	// A user-friendly name for the Policy. Does not have to be unique, and it's changeable. Avoid entering
	// confidential information.
	GetDisplayName() *string
}

type createautoscalingpolicydetails struct {
	JsonData    []byte
	Capacity    *Capacity `mandatory:"true" json:"capacity"`
	DisplayName *string   `mandatory:"false" json:"displayName"`
	PolicyType  string    `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *createautoscalingpolicydetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreateautoscalingpolicydetails createautoscalingpolicydetails
	s := struct {
		Model Unmarshalercreateautoscalingpolicydetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Capacity = s.Model.Capacity
	m.DisplayName = s.Model.DisplayName
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createautoscalingpolicydetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "threshold":
		mm := CreateThresholdPolicyDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCapacity returns Capacity
func (m createautoscalingpolicydetails) GetCapacity() *Capacity {
	return m.Capacity
}

//GetDisplayName returns DisplayName
func (m createautoscalingpolicydetails) GetDisplayName() *string {
	return m.DisplayName
}

func (m createautoscalingpolicydetails) String() string {
	return common.PointerString(m)
}
