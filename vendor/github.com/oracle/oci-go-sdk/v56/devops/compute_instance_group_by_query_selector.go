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

// ComputeInstanceGroupByQuerySelector Specifies the Compute instance group environment filtered by DSL expression of the compute instances.
type ComputeInstanceGroupByQuerySelector struct {

	// Region identifier referred by the deployment environment. Region identifiers are listed at https://docs.oracle.com/en-us/iaas/Content/General/Concepts/regions.htm
	Region *string `mandatory:"true" json:"region"`

	// Query expression confirming to the OCI Search Language syntax to select compute instances for the group. The language is documented at https://docs.oracle.com/en-us/iaas/Content/Search/Concepts/querysyntax.htm
	Query *string `mandatory:"true" json:"query"`
}

func (m ComputeInstanceGroupByQuerySelector) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupByQuerySelector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupByQuerySelector ComputeInstanceGroupByQuerySelector
	s := struct {
		DiscriminatorParam string `json:"selectorType"`
		MarshalTypeComputeInstanceGroupByQuerySelector
	}{
		"INSTANCE_QUERY",
		(MarshalTypeComputeInstanceGroupByQuerySelector)(m),
	}

	return json.Marshal(&s)
}
