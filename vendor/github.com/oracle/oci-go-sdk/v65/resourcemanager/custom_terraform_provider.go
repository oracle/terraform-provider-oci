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

// CustomTerraformProvider Location information about custom Terraform providers for a stack.
// For more information, see Custom Providers (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Concepts/resourcemanager.htm#features__custom-providers).
// Note: Older stacks must be explicitly updated to use Terraform Registry (`isThirdPartyProviderExperienceEnabled=true`).
// See UpdateStack. For more information, see
// Using Terraform Registry with Older Stacks (https://docs.cloud.oracle.com/iaas/Content/ResourceManager/Tasks/update-stack-tf-reg.htm).
type CustomTerraformProvider struct {

	// The name of the region that contains the bucket you want.
	// For information about regions, see Regions and Availability Domains (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
	// Example: `us-phoenix-1`
	Region *string `mandatory:"true" json:"region"`

	// The Object Storage namespace that contains the bucket you want.
	// For information about Object Storage namespaces, see Understanding Object Storage Namespaces (https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/understandingnamespaces.htm).
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the bucket that contains the binary files for the custom Terraform providers.
	// For information about buckets, see Managing Buckets (https://docs.cloud.oracle.com/iaas/Content/Object/Tasks/managingbuckets.htm).
	BucketName *string `mandatory:"true" json:"bucketName"`
}

func (m CustomTerraformProvider) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomTerraformProvider) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
