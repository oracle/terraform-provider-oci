// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps APIs to create a DevOps project to group the pipelines,  add reference to target deployment environments, add artifacts to deploy,  and create deployment pipelines needed to deploy your software.
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v45/common"
)

// ApprovalPolicy Specifies the approval policy.
type ApprovalPolicy interface {
}

type approvalpolicy struct {
	JsonData           []byte
	ApprovalPolicyType string `json:"approvalPolicyType"`
}

// UnmarshalJSON unmarshals json
func (m *approvalpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapprovalpolicy approvalpolicy
	s := struct {
		Model Unmarshalerapprovalpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ApprovalPolicyType = s.Model.ApprovalPolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *approvalpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ApprovalPolicyType {
	case "COUNT_BASED_APPROVAL":
		mm := CountBasedApprovalPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m approvalpolicy) String() string {
	return common.PointerString(m)
}

// ApprovalPolicyApprovalPolicyTypeEnum Enum with underlying type: string
type ApprovalPolicyApprovalPolicyTypeEnum string

// Set of constants representing the allowable values for ApprovalPolicyApprovalPolicyTypeEnum
const (
	ApprovalPolicyApprovalPolicyTypeCountBasedApproval ApprovalPolicyApprovalPolicyTypeEnum = "COUNT_BASED_APPROVAL"
)

var mappingApprovalPolicyApprovalPolicyType = map[string]ApprovalPolicyApprovalPolicyTypeEnum{
	"COUNT_BASED_APPROVAL": ApprovalPolicyApprovalPolicyTypeCountBasedApproval,
}

// GetApprovalPolicyApprovalPolicyTypeEnumValues Enumerates the set of values for ApprovalPolicyApprovalPolicyTypeEnum
func GetApprovalPolicyApprovalPolicyTypeEnumValues() []ApprovalPolicyApprovalPolicyTypeEnum {
	values := make([]ApprovalPolicyApprovalPolicyTypeEnum, 0)
	for _, v := range mappingApprovalPolicyApprovalPolicyType {
		values = append(values, v)
	}
	return values
}
