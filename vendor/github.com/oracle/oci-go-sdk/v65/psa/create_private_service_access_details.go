// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// PrivateServiceAccess Control Plane API
//
// Use the PrivateServiceAccess Control Plane API to manage privateServiceAccess.
//

package psa

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePrivateServiceAccessDetails Details to create a private service access.
type CreatePrivateServiceAccessDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// private service access.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN's
	// subnet where the private service access's VNIC will reside.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// A unique service identifier for which the private service access was created.
	ServiceId *string `mandatory:"true" json:"serviceId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Security attributes (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/zpr-artifacts.htm#security-attributes) are labels
	// for a resource that can be referenced in a Zero Trust Packet Routing (https://docs.oracle.com/iaas/Content/zero-trust-packet-routing/overview.htm)
	// (ZPR) policy to control access to ZPR-supported resources.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A description of this private service accesss.
	Description *string `mandatory:"false" json:"description"`

	// A list of the OCIDs of the network security groups (NSGs) to add the private
	// service access's VNIC to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private IPv4 address (in the consumer's VCN) that represents the access point for the
	// associated service.
	Ipv4Ip *string `mandatory:"false" json:"ipv4Ip"`
}

func (m CreatePrivateServiceAccessDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePrivateServiceAccessDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
