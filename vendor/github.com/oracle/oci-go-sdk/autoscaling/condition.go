// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Condition A container for metric and action details
type Condition struct {
	Action *Action `mandatory:"true" json:"action"`

	Metric *Metric `mandatory:"true" json:"metric"`

	// A user-friendly name for the AutoScalingConfiguration condition details. Does not have to be unique, and
	// it's changeable. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Id of the condition that is assigned after creation
	Id *string `mandatory:"false" json:"id"`
}

func (m Condition) String() string {
	return common.PointerString(m)
}
