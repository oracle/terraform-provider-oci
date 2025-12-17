// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Batch API
//
// Use the Batch Control Plane API to encapsulate and manage all aspects of computationally intensive jobs.
//

package batch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityContext Security context for container runtime configuration.
// See also docs (https://docs.oracle.com/en-us/iaas/api/#/en/container-instances/20210415/datatypes/LinuxSecurityContext).
type SecurityContext struct {

	// User ID for running processes inside the container.
	RunAsUser *int `mandatory:"false" json:"runAsUser"`

	// Group ID for running processes inside the container.
	RunAsGroup *int `mandatory:"false" json:"runAsGroup"`

	// A special supplemental group ID that applies to all containers in a pod.
	FsGroup *int `mandatory:"false" json:"fsGroup"`
}

func (m SecurityContext) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityContext) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
