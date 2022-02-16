// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// Capacity Capacity limits for the instance pool.
type Capacity struct {

	// For a threshold-based autoscaling policy, this value is the maximum number of instances the instance pool is allowed
	// to increase to (scale out).
	// For a schedule-based autoscaling policy, this value is not used.
	Max *int `mandatory:"false" json:"max"`

	// For a threshold-based autoscaling policy, this value is the minimum number of instances the instance pool is allowed
	// to decrease to (scale in).
	// For a schedule-based autoscaling policy, this value is not used.
	Min *int `mandatory:"false" json:"min"`

	// For a threshold-based autoscaling policy, this value is the initial number of instances to launch in the instance pool
	// immediately after autoscaling is enabled. After autoscaling retrieves performance metrics, the number of
	// instances is automatically adjusted from this initial number to a number that is based on the limits that
	// you set.
	// For a schedule-based autoscaling policy, this value is the target pool size to scale to when executing the schedule
	// that's defined in the autoscaling policy.
	Initial *int `mandatory:"false" json:"initial"`
}

func (m Capacity) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Capacity) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
