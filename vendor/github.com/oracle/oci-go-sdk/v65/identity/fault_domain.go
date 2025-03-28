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

// FaultDomain A Fault Domain is a logical grouping of hardware and infrastructure within an Availability Domain that can become
// unavailable in its entirety either due to hardware failure such as Top-of-rack (TOR) switch failure or due to
// planned software maintenance such as security updates that reboot your instances.
type FaultDomain struct {

	// The name of the Fault Domain.
	Name *string `mandatory:"false" json:"name"`

	// The OCID of the Fault Domain.
	Id *string `mandatory:"false" json:"id"`

	// The OCID of the compartment. Currently only tenancy (root) compartment can be provided.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The name of the availabilityDomain where the Fault Domain belongs.
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`
}

func (m FaultDomain) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FaultDomain) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
