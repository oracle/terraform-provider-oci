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

// UpdateDeployEnvironmentDetails The information to be updated.
type UpdateDeployEnvironmentDetails interface {

	// Optional description about the deployment environment.
	GetDescription() *string

	// Deployment environment display name. Avoid entering confidential information.
	GetDisplayName() *string

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	GetFreeformTags() map[string]string

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	GetDefinedTags() map[string]map[string]interface{}
}

type updatedeployenvironmentdetails struct {
	JsonData              []byte
	Description           *string                           `mandatory:"false" json:"description"`
	DisplayName           *string                           `mandatory:"false" json:"displayName"`
	FreeformTags          map[string]string                 `mandatory:"false" json:"freeformTags"`
	DefinedTags           map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
	DeployEnvironmentType string                            `json:"deployEnvironmentType"`
}

// UnmarshalJSON unmarshals json
func (m *updatedeployenvironmentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatedeployenvironmentdetails updatedeployenvironmentdetails
	s := struct {
		Model Unmarshalerupdatedeployenvironmentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.DisplayName = s.Model.DisplayName
	m.FreeformTags = s.Model.FreeformTags
	m.DefinedTags = s.Model.DefinedTags
	m.DeployEnvironmentType = s.Model.DeployEnvironmentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatedeployenvironmentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.DeployEnvironmentType {
	case "FUNCTION":
		mm := UpdateFunctionDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "COMPUTE_INSTANCE_GROUP":
		mm := UpdateComputeInstanceGroupDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "OKE_CLUSTER":
		mm := UpdateOkeClusterDeployEnvironmentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetDescription returns Description
func (m updatedeployenvironmentdetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m updatedeployenvironmentdetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetFreeformTags returns FreeformTags
func (m updatedeployenvironmentdetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m updatedeployenvironmentdetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m updatedeployenvironmentdetails) String() string {
	return common.PointerString(m)
}
