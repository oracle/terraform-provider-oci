// Copyright (c) 2016, 2018, 2021, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// API for the Resource Manager service.
// Use this API to install, configure, and manage resources via the "infrastructure-as-code" model.
// For more information, see
// Overview of Resource Manager (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm).
//

package resourcemanager

import (
	"github.com/oracle/oci-go-sdk/v36/common"
)

// DetectStackDriftDetails The details for detecting drift in a stack
type DetectStackDriftDetails struct {

	// The list of resources in the specified stack to detect drift for. Each resource is identified by a resource address,
	// which is a case-insensitive string derived from the resource type and name specified in the stack's Terraform configuration plus an optional index.
	// For example, the resource address for the fourth Compute instance with the name "test_instance" is oci_core_instance.test_instance3.
	// For more details and examples of resource addresses, see the Terraform documentation at Resource spec (https://www.terraform.io/docs/internals/resource-addressing.html#examples).
	ResourceAddresses []string `mandatory:"false" json:"resourceAddresses"`
}

func (m DetectStackDriftDetails) String() string {
	return common.PointerString(m)
}
