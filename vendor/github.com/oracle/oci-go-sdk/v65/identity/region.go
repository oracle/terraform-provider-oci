// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Region A localized geographic area, such as Phoenix, AZ. Oracle Cloud Infrastructure is hosted in regions and Availability
// Domains. A region is composed of several Availability Domains. An Availability Domain is one or more data centers
// located within a region. For more information, see Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access,
// see Get Started with Policies (https://docs.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm).
type Region struct {

	// The key of the region. See Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm) for
	// the full list of supported 3-letter region codes.
	// Example: `PHX`
	Key *string `mandatory:"false" json:"key"`

	// The name of the region. See Regions and Availability Domains (https://docs.oracle.com/iaas/Content/General/Concepts/regions.htm)
	// for the full list of supported region names.
	// Example: `us-phoenix-1`
	Name *string `mandatory:"false" json:"name"`
}

func (m Region) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Region) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
