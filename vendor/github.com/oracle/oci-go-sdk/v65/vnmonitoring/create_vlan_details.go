// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVlanDetails The representation of CreateVlanDetails
type CreateVlanDetails struct {

	// The range of IPv4 addresses that will be used for layer 3 communication with
	// hosts outside the VLAN. The CIDR must maintain the following rules -
	// 1. The CIDR block is valid and correctly formatted.
	// 2. The new range is within one of the parent VCN ranges.
	// Example: `192.0.2.0/24`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the VLAN.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN to contain the VLAN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Controls whether the VLAN is regional or specific to an availability domain.
	// A regional VLAN has the flexibility to implement failover across availability domains.
	// Previously, all VLANs were AD-specific.
	// To create a regional VLAN, omit this attribute. Resources created subsequently in this
	// VLAN (such as a Compute instance) can be created in any availability domain in the region.
	// To create an AD-specific VLAN, use this attribute to specify the availability domain.
	// Resources created in this VLAN must be in that availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A list of the OCIDs of the network security groups (NSGs) to add all VNICs in the VLAN to. For more
	// information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the VLAN will use. If you don't provide a value,
	// the VLAN uses the VCN's default route table.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The IEEE 802.1Q VLAN tag for this VLAN. The value must be unique across all
	// VLANs in the VCN. If you don't provide a value, Oracle assigns one.
	// You cannot change the value later. VLAN tag 0 is reserved for use by Oracle.
	VlanTag *int `mandatory:"false" json:"vlanTag"`
}

func (m CreateVlanDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVlanDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
