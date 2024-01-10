// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ComputeCostEstimation Cost estimation for compute
type ComputeCostEstimation struct {

	// OCPU per hour
	OcpuPerHour *float32 `mandatory:"true" json:"ocpuPerHour"`

	// Gigabyte per hour
	MemoryGbPerHour *float32 `mandatory:"true" json:"memoryGbPerHour"`

	// GPU per hour
	GpuPerHour *float32 `mandatory:"true" json:"gpuPerHour"`

	// Total per hour
	TotalPerHour *float32 `mandatory:"true" json:"totalPerHour"`

	// OCPU per hour by subscription
	OcpuPerHourBySubscription *float32 `mandatory:"false" json:"ocpuPerHourBySubscription"`

	// Gigabyte per hour by subscription
	MemoryGbPerHourBySubscription *float32 `mandatory:"false" json:"memoryGbPerHourBySubscription"`

	// GPU per hour by subscription
	GpuPerHourBySubscription *float32 `mandatory:"false" json:"gpuPerHourBySubscription"`

	// Total usage per hour by subscription
	TotalPerHourBySubscription *float32 `mandatory:"false" json:"totalPerHourBySubscription"`

	// Total number of OCPUs
	OcpuCount *float32 `mandatory:"false" json:"ocpuCount"`

	// Total usage of memory
	MemoryAmountGb *float32 `mandatory:"false" json:"memoryAmountGb"`

	// Total number of GPU
	GpuCount *float32 `mandatory:"false" json:"gpuCount"`
}

func (m ComputeCostEstimation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ComputeCostEstimation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
