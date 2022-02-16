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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// CreateLoadBalancerTrafficShiftDeployStageDetails Specifies load balancer traffic shift stage.
type CreateLoadBalancerTrafficShiftDeployStageDetails struct {

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"true" json:"deployStagePredecessorCollection"`

	BlueBackendIps *BackendSetIpCollection `mandatory:"true" json:"blueBackendIps"`

	GreenBackendIps *BackendSetIpCollection `mandatory:"true" json:"greenBackendIps"`

	RolloutPolicy *LoadBalancerTrafficShiftRolloutPolicy `mandatory:"true" json:"rolloutPolicy"`

	LoadBalancerConfig *LoadBalancerConfig `mandatory:"true" json:"loadBalancerConfig"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	RollbackPolicy DeployStageRollbackPolicy `mandatory:"false" json:"rollbackPolicy"`

	// Specifies the target or destination backend set.
	TrafficShiftTarget LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum `mandatory:"true" json:"trafficShiftTarget"`
}

//GetDescription returns Description
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetDisplayName() *string {
	return m.DisplayName
}

//GetDeployPipelineId returns DeployPipelineId
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

func (m CreateLoadBalancerTrafficShiftDeployStageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum(string(m.TrafficShiftTarget)); !ok && m.TrafficShiftTarget != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TrafficShiftTarget: %s. Supported values are: %s.", m.TrafficShiftTarget, strings.Join(GetLoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateLoadBalancerTrafficShiftDeployStageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateLoadBalancerTrafficShiftDeployStageDetails CreateLoadBalancerTrafficShiftDeployStageDetails
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeCreateLoadBalancerTrafficShiftDeployStageDetails
	}{
		"LOAD_BALANCER_TRAFFIC_SHIFT",
		(MarshalTypeCreateLoadBalancerTrafficShiftDeployStageDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *CreateLoadBalancerTrafficShiftDeployStageDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                                                   `json:"description"`
		DisplayName                      *string                                                   `json:"displayName"`
		FreeformTags                     map[string]string                                         `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{}                         `json:"definedTags"`
		RollbackPolicy                   deploystagerollbackpolicy                                 `json:"rollbackPolicy"`
		DeployPipelineId                 *string                                                   `json:"deployPipelineId"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection                         `json:"deployStagePredecessorCollection"`
		BlueBackendIps                   *BackendSetIpCollection                                   `json:"blueBackendIps"`
		GreenBackendIps                  *BackendSetIpCollection                                   `json:"greenBackendIps"`
		TrafficShiftTarget               LoadBalancerTrafficShiftDeployStageTrafficShiftTargetEnum `json:"trafficShiftTarget"`
		RolloutPolicy                    *LoadBalancerTrafficShiftRolloutPolicy                    `json:"rolloutPolicy"`
		LoadBalancerConfig               *LoadBalancerConfig                                       `json:"loadBalancerConfig"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	nn, e = model.RollbackPolicy.UnmarshalPolymorphicJSON(model.RollbackPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.RollbackPolicy = nn.(DeployStageRollbackPolicy)
	} else {
		m.RollbackPolicy = nil
	}

	m.DeployPipelineId = model.DeployPipelineId

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.BlueBackendIps = model.BlueBackendIps

	m.GreenBackendIps = model.GreenBackendIps

	m.TrafficShiftTarget = model.TrafficShiftTarget

	m.RolloutPolicy = model.RolloutPolicy

	m.LoadBalancerConfig = model.LoadBalancerConfig

	return
}
