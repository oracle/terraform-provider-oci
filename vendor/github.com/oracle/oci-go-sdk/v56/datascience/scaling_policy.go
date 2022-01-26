// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ScalingPolicy The scaling policy to apply to each model of the deployment.
type ScalingPolicy interface {
}

type scalingpolicy struct {
	JsonData   []byte
	PolicyType string `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *scalingpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerscalingpolicy scalingpolicy
	s := struct {
		Model Unmarshalerscalingpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *scalingpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "FIXED_SIZE":
		mm := FixedSizeScalingPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m scalingpolicy) String() string {
	return common.PointerString(m)
}

// ScalingPolicyPolicyTypeEnum Enum with underlying type: string
type ScalingPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ScalingPolicyPolicyTypeEnum
const (
	ScalingPolicyPolicyTypeFixedSize ScalingPolicyPolicyTypeEnum = "FIXED_SIZE"
)

var mappingScalingPolicyPolicyType = map[string]ScalingPolicyPolicyTypeEnum{
	"FIXED_SIZE": ScalingPolicyPolicyTypeFixedSize,
}

// GetScalingPolicyPolicyTypeEnumValues Enumerates the set of values for ScalingPolicyPolicyTypeEnum
func GetScalingPolicyPolicyTypeEnumValues() []ScalingPolicyPolicyTypeEnum {
	values := make([]ScalingPolicyPolicyTypeEnum, 0)
	for _, v := range mappingScalingPolicyPolicyType {
		values = append(values, v)
	}
	return values
}
