// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DeployStageRollbackPolicy Specifies the rollback policy. This is initiated on the failure of certain stage types.
type DeployStageRollbackPolicy interface {
}

type deploystagerollbackpolicy struct {
	JsonData   []byte
	PolicyType string `json:"policyType"`
}

// UnmarshalJSON unmarshals json
func (m *deploystagerollbackpolicy) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdeploystagerollbackpolicy deploystagerollbackpolicy
	s := struct {
		Model Unmarshalerdeploystagerollbackpolicy
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.PolicyType = s.Model.PolicyType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *deploystagerollbackpolicy) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.PolicyType {
	case "NO_STAGE_ROLLBACK_POLICY":
		mm := NoDeployStageRollbackPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "AUTOMATED_STAGE_ROLLBACK_POLICY":
		mm := AutomatedDeployStageRollbackPolicy{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m deploystagerollbackpolicy) String() string {
	return common.PointerString(m)
}

// DeployStageRollbackPolicyPolicyTypeEnum Enum with underlying type: string
type DeployStageRollbackPolicyPolicyTypeEnum string

// Set of constants representing the allowable values for DeployStageRollbackPolicyPolicyTypeEnum
const (
	DeployStageRollbackPolicyPolicyTypeAutomatedStageRollbackPolicy DeployStageRollbackPolicyPolicyTypeEnum = "AUTOMATED_STAGE_ROLLBACK_POLICY"
	DeployStageRollbackPolicyPolicyTypeNoStageRollbackPolicy        DeployStageRollbackPolicyPolicyTypeEnum = "NO_STAGE_ROLLBACK_POLICY"
)

var mappingDeployStageRollbackPolicyPolicyType = map[string]DeployStageRollbackPolicyPolicyTypeEnum{
	"AUTOMATED_STAGE_ROLLBACK_POLICY": DeployStageRollbackPolicyPolicyTypeAutomatedStageRollbackPolicy,
	"NO_STAGE_ROLLBACK_POLICY":        DeployStageRollbackPolicyPolicyTypeNoStageRollbackPolicy,
}

// GetDeployStageRollbackPolicyPolicyTypeEnumValues Enumerates the set of values for DeployStageRollbackPolicyPolicyTypeEnum
func GetDeployStageRollbackPolicyPolicyTypeEnumValues() []DeployStageRollbackPolicyPolicyTypeEnum {
	values := make([]DeployStageRollbackPolicyPolicyTypeEnum, 0)
	for _, v := range mappingDeployStageRollbackPolicyPolicyType {
		values = append(values, v)
	}
	return values
}
