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

// ManagedComputeClusterWorkloadFixedSizeScalingPolicy The fixed size scaling policy for workload scaling on managed compute cluster type compute target.
type ManagedComputeClusterWorkloadFixedSizeScalingPolicy struct {

	// The number of instances of the workload.
	InstanceCount *int `mandatory:"true" json:"instanceCount"`
}

func (m ManagedComputeClusterWorkloadFixedSizeScalingPolicy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ManagedComputeClusterWorkloadFixedSizeScalingPolicy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ManagedComputeClusterWorkloadFixedSizeScalingPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManagedComputeClusterWorkloadFixedSizeScalingPolicy ManagedComputeClusterWorkloadFixedSizeScalingPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeManagedComputeClusterWorkloadFixedSizeScalingPolicy
	}{
		"FIXED_SIZE",
		(MarshalTypeManagedComputeClusterWorkloadFixedSizeScalingPolicy)(m),
	}

	return json.Marshal(&s)
}
