// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Manager API
//
// Use the Resource Manager API to automate deployment and operations for all Oracle Cloud Infrastructure resources.
// Using the infrastructure-as-code (IaC) model, the service is based on Terraform, an open source industry standard that lets DevOps engineers develop and deploy their infrastructure anywhere.
// For more information, see
// the Resource Manager documentation (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/home.htm).
//

package resourcemanager

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DetectStackDriftDetails Details for detecting drift in a stack.
type DetectStackDriftDetails struct {

	// The list of resources in the specified stack to detect drift for. Each resource is identified by a resource address,
	// which is a string derived from the resource type and name specified in the stack's Terraform configuration plus an optional index.
	// For example, the resource address for the fourth Compute instance with the name "test_instance" is oci_core_instance.test_instance3.
	// For more details and examples of resource addresses, see the Terraform documentation at Resource spec (https://www.terraform.io/docs/internals/resource-addressing.html#examples).
	ResourceAddresses []string `mandatory:"false" json:"resourceAddresses"`

	// Specifies whether or not to upgrade provider versions.
	// Within the version constraints of your Terraform configuration, use the latest versions available from the source of Terraform providers.
	// For more information about this option, see Dependency Lock File (terraform.io) (https://www.terraform.io/language/files/dependency-lock).
	IsProviderUpgradeRequired *bool `mandatory:"false" json:"isProviderUpgradeRequired"`
}

func (m DetectStackDriftDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DetectStackDriftDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
