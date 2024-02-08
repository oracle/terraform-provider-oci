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
		common.Logf("Recieved unsupported enum value for ScalingPolicy: %s.", m.PolicyType)
		return *m, nil
	}
}

func (m scalingpolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m scalingpolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScalingPolicyPolicyTypeEnum Enum with underlying type: string
type ScalingPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for ScalingPolicyPolicyTypeEnum
const (
	ScalingPolicyPolicyTypeFixedSize ScalingPolicyPolicyTypeEnum = "FIXED_SIZE"
)

var mappingScalingPolicyPolicyTypeEnum = map[string]ScalingPolicyPolicyTypeEnum{
	"FIXED_SIZE": ScalingPolicyPolicyTypeFixedSize,
}

var mappingScalingPolicyPolicyTypeEnumLowerCase = map[string]ScalingPolicyPolicyTypeEnum{
	"fixed_size": ScalingPolicyPolicyTypeFixedSize,
}

// GetScalingPolicyPolicyTypeEnumValues Enumerates the set of values for ScalingPolicyPolicyTypeEnum
func GetScalingPolicyPolicyTypeEnumValues() []ScalingPolicyPolicyTypeEnum {
	values := make([]ScalingPolicyPolicyTypeEnum, 0)
	for _, v := range mappingScalingPolicyPolicyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScalingPolicyPolicyTypeEnumStringValues Enumerates the set of values in String for ScalingPolicyPolicyTypeEnum
func GetScalingPolicyPolicyTypeEnumStringValues() []string {
	return []string{
		"FIXED_SIZE",
	}
}

// GetMappingScalingPolicyPolicyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScalingPolicyPolicyTypeEnum(val string) (ScalingPolicyPolicyTypeEnum, bool) {
	enum, ok := mappingScalingPolicyPolicyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
