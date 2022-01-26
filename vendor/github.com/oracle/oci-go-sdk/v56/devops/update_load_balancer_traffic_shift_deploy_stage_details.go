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

// UpdateLoadBalancerTrafficShiftDeployStageDetails Specifies load balancer traffic shift stage.
type UpdateLoadBalancerTrafficShiftDeployStageDetails struct {

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	BlueBackendIps *BackendSetIpCollection `mandatory:"false" json:"blueBackendIps"`

	GreenBackendIps *BackendSetIpCollection `mandatory:"false" json:"greenBackendIps"`

	RolloutPolicy *LoadBalancerTrafficShiftRolloutPolicy `mandatory:"false" json:"rolloutPolicy"`

	LoadBalancerConfig *LoadBalancerConfig `mandatory:"false" json:"loadBalancerConfig"`

	RollbackPolicy DeployStageRollbackPolicy `mandatory:"false" json:"rollbackPolicy"`

	// Specifies the target or destination backend set.
	TrafficShiftTarget LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum `mandatory:"false" json:"trafficShiftTarget,omitempty"`
}

//GetDescription returns Description
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m UpdateLoadBalancerTrafficShiftDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateLoadBalancerTrafficShiftDeployStageDetails UpdateLoadBalancerTrafficShiftDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeUpdateLoadBalancerTrafficShiftDeployStageDetails
	}{
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		(MarshalTypeUpdateLoadBalancerTrafficShiftDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *UpdateLoadBalancerTrafficShiftDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                                                   `json:"description"`
		DisplayName                      *string                                                   `json:"displayName"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection                         `json:"deployStagePredecessorCollection"`
		FreeformTags                     map[string]string                                         `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}                         `json:"definedTags"`
		BlueBackendIps                   *BackendSetIpCollection                                   `json:"blueBackendIps"`
		GreenBackendIps                  *BackendSetIpCollection                                   `json:"greenBackendIps"`
		TrafficShiftTarget               LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum `json:"trafficShiftTarget"`
		RolloutPolicy                    *LoadBalancerTrafficShiftRolloutPolicy                    `json:"rolloutPolicy"`
		LoadBalancerConfig               *LoadBalancerConfig                                       `json:"loadBalancerConfig"`
		RollbackPolicy                   deploystagerollbackpolicy                                 `json:"rollbackPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.BlueBackendIps = model.BlueBackendIps

	m.GreenBackendIps = model.GreenBackendIps

	m.TrafficShiftTarget = model.TrafficShiftTarget

	m.RolloutPolicy = model.RolloutPolicy

	m.LoadBalancerConfig = model.LoadBalancerConfig

	nn, e = model.RollbackPolicy.UnmarshalPolymorphicJSON(model.RollbackPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RollbackPolicy = nn.(DeployStageRollbackPolicy)
	} else {
		m.RollbackPolicy = nil
	}

	return
}
