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

// ManagedComputeClusterModelDeployInfrastructureConfigDetails Infrastructure configuration details for model deploy on managed compute cluster type compute target.
type ManagedComputeClusterModelDeployInfrastructureConfigDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Compute Target.
	ComputeTargetId *string `mandatory:"true" json:"computeTargetId"`

	ModelDeploymentResourceConfiguration *ManagedComputeClusterModelDeploymentResourceConfiguration `mandatory:"false" json:"modelDeploymentResourceConfiguration"`

	ScalingPolicy ManagedComputeClusterWorkloadScalingPolicy `mandatory:"false" json:"scalingPolicy"`
}

func (m ManagedComputeClusterModelDeployInfrastructureConfigDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedComputeClusterModelDeployInfrastructureConfigDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedComputeClusterModelDeployInfrastructureConfigDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedComputeClusterModelDeployInfrastructureConfigDetails ManagedComputeClusterModelDeployInfrastructureConfigDetails
	s := struct {
		DiscriminatorParam string `json:"infrastructureType"`
		MarshalTypeManagedComputeClusterModelDeployInfrastructureConfigDetails
	}{
		"MANAGED_COMPUTE_CLUSTER",
		(MarshalTypeManagedComputeClusterModelDeployInfrastructureConfigDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ManagedComputeClusterModelDeployInfrastructureConfigDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		ModelDeploymentResourceConfiguration *ManagedComputeClusterModelDeploymentResourceConfiguration `json:"modelDeploymentResourceConfiguration"`
		ScalingPolicy                        managedcomputeclusterworkloadscalingpolicy                 `json:"scalingPolicy"`
		ComputeTargetId                      *string                                                    `json:"computeTargetId"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.ModelDeploymentResourceConfiguration = model.ModelDeploymentResourceConfiguration

	nn, e = model.ScalingPolicy.UnmarshalPolymorphicJSON(model.ScalingPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ScalingPolicy = nn.(ManagedComputeClusterWorkloadScalingPolicy)
	} else {
		m.ScalingPolicy = nil
	}

	m.ComputeTargetId = model.ComputeTargetId

	return
}
