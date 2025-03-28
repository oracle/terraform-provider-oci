// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Blockchain Platform Control Plane API
//
// Blockchain Platform Control Plane API
//

package blockchain

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OcpuUtilizationInfo Ocpu utilization for a VM host
type OcpuUtilizationInfo struct {

	// Host name of VM
	Host *string `mandatory:"false" json:"host"`

	// Number of OCPU utilized
	OcpuUtilizationNumber *float32 `mandatory:"false" json:"ocpuUtilizationNumber"`

	// Number of total OCPU capacity on the host
	OcpuCapacityNumber *float32 `mandatory:"false" json:"ocpuCapacityNumber"`
}

func (m OcpuUtilizationInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OcpuUtilizationInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
