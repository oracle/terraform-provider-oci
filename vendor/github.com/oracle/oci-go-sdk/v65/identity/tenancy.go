// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Tenancy The root compartment that contains all of your organization's compartments and other
// Oracle Cloud Infrastructure cloud resources. When you sign up for Oracle Cloud Infrastructure,
// Oracle creates a tenancy for your company, which is a secure and isolated partition
// where you can create, organize, and administer your cloud resources.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.cloud.oracle.com/Content/Identity/policiesgs/get-started-with-policies.htm).
type Tenancy struct {

	// The OCID of the tenancy.
	Id *string `mandatory:"false" json:"id"`

	// The name of the tenancy.
	Name *string `mandatory:"false" json:"name"`

	// The description of the tenancy.
	Description *string `mandatory:"false" json:"description"`

	// The region key for the tenancy's home region. For the full list of supported regions, see
	// Regions and Availability Domains (https://docs.cloud.oracle.com/Content/General/Concepts/regions.htm).
	// Example: `PHX`
	HomeRegionKey *string `mandatory:"false" json:"homeRegionKey"`

	// Url which refers to the UPI IDCS compatibility layer endpoint configured for this Tenant's home region.
	UpiIdcsCompatibilityLayerEndpoint *string `mandatory:"false" json:"upiIdcsCompatibilityLayerEndpoint"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Tenancy) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Tenancy) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
