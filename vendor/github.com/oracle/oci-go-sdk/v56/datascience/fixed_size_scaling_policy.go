// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// FixedSizeScalingPolicy The fixed size scaling policy.
type FixedSizeScalingPolicy struct {

	// The number of instances for the model deployment.
	InstanceCount *int `mandatory:"true" json:"instanceCount"`
}

func (m FixedSizeScalingPolicy) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m FixedSizeScalingPolicy) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeFixedSizeScalingPolicy FixedSizeScalingPolicy
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeFixedSizeScalingPolicy
	}{
		"FIXED_SIZE",
		(MarshalTypeFixedSizeScalingPolicy)(m),
	}

	return json.Marshal(&s)
}
