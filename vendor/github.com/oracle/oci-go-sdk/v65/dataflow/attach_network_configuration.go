// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AttachNetworkConfiguration A Data Flow AttachNetworkConfiguration object.
type AttachNetworkConfiguration struct {

	// WorkspaceKey of the cluster.
	WorkspaceKey *string `mandatory:"true" json:"workspaceKey"`

	// The OCID of the Compute Cluster.
	ComputeClusterId *string `mandatory:"true" json:"computeClusterId"`

	// Lifecycle State of Network Configuration.
	LifecycleState *string `mandatory:"true" json:"lifecycleState"`

	// Detailed State of Network Configuration.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`
}

func (m AttachNetworkConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AttachNetworkConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
