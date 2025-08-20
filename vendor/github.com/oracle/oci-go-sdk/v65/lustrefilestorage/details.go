// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// File Storage with Lustre API
//
// Use the File Storage with Lustre API to manage Lustre file systems and related resources. For more information, see File Storage with Lustre (https://docs.oracle.com/iaas/Content/lustre/home.htm).
//

package lustrefilestorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Details ManagementCell details are captured in this JSON object.
type Details struct {

	// JSON object version. It helps in deploying read/write phase when JSON object is updated.
	Version *int64 `mandatory:"false" json:"version"`

	// Management plane load balancer endpoint for ManagementCell
	MpLoadBalancerEndpoint *string `mandatory:"false" json:"mpLoadBalancerEndpoint"`

	// Total Cell capacity available for creating new filesystems on the cell. Measured in GB. For create request, this will mapped to availableCapacityInGBs and need not to be part of request params.
	TotalCapacityInGBs *int64 `mandatory:"false" json:"totalCapacityInGBs"`
}

func (m Details) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Details) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
