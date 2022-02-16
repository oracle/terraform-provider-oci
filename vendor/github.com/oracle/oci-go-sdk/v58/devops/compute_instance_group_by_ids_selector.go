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

// ComputeInstanceGroupByIdsSelector Specifies the Compute instance group environment by listing the OCIDs of the compute instances.
type ComputeInstanceGroupByIdsSelector struct {

	// Compute instance OCID identifiers that are members of this group.
	ComputeInstanceIds []string `mandatory:"true" json:"computeInstanceIds"`
}

func (m ComputeInstanceGroupByIdsSelector) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeInstanceGroupByIdsSelector) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ComputeInstanceGroupByIdsSelector) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeComputeInstanceGroupByIdsSelector ComputeInstanceGroupByIdsSelector
	s := struct {
		DiscriminatorParam string `json:"selectorType"`
		MarshalTypeComputeInstanceGroupByIdsSelector
	}{
		"INSTANCE_IDS",
		(MarshalTypeComputeInstanceGroupByIdsSelector)(m),
	}

	return json.Marshal(&s)
}
