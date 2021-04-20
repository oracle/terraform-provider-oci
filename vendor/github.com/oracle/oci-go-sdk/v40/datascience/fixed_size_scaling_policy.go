// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science APIs to organize your data science work, access data and computing resources, and build, train, deploy, and manage models on Oracle Cloud.
//

package datascience

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v40/common"
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
