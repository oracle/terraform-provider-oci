// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateVcnDetails The representation of CreateVcnDetails
type CreateVcnDetails struct {

	// The CIDR IP address block of the VCN.
	// Example: `10.0.0.0/16`
	CidrBlock *string `mandatory:"true" json:"cidrBlock"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the VCN.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// If you enable IPv6 for the VCN (see `isIpv6Enabled`), you may optionally provide an IPv6
	// /56 CIDR block from the supported ranges (see IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// The addresses in this block will be considered private and cannot be accessed
	// from the internet. The documentation refers to this as a *custom CIDR* for the VCN.
	// If you don't provide a custom CIDR for the VCN, Oracle assigns the VCN's IPv6 /56 CIDR block.
	// Regardless of whether you or Oracle assigns the `ipv6CidrBlock`,
	// Oracle *also* assigns the VCN an IPv6 CIDR block for the VCN's public IP address space
	// (see the `ipv6PublicCidrBlock` of the Vcn object). If you do
	// not assign a custom CIDR, Oracle uses the *same* Oracle-assigned CIDR for both the private
	// IP address space (`ipv6CidrBlock` in the `Vcn` object) and the public IP addreses space
	// (`ipv6PublicCidrBlock` in the `Vcn` object). This means that a given VNIC might use the same
	// IPv6 IP address for both private and public (internet) communication. You control whether
	// an IPv6 address can be used for internet communication by using the `isInternetAccessAllowed`
	// attribute in the Ipv6 object.
	// For important details about IPv6 addressing in a VCN, see IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// Example: `2001:0db8:0123::/48`
	Ipv6CidrBlock *string `mandatory:"false" json:"ipv6CidrBlock"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// A DNS label for the VCN, used in conjunction with the VNIC's hostname and
	// subnet's DNS label to form a fully qualified domain name (FQDN) for each VNIC
	// within this subnet (for example, `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Not required to be unique, but it's a best practice to set unique DNS labels
	// for VCNs in your tenancy. Must be an alphanumeric string that begins with a letter.
	// The value cannot be changed.
	// You must set this value if you want instances to be able to use hostnames to
	// resolve other instances in the VCN. Otherwise the Internet and VCN Resolver
	// will not work.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// Example: `vcn1`
	DnsLabel *string `mandatory:"false" json:"dnsLabel"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether IPv6 is enabled for the VCN. Default is `false`.
	// If enabled, Oracle will assign the VCN a IPv6 /56 CIDR block.
	// For important details about IPv6 addressing in a VCN, see IPv6 Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/ipv6.htm).
	// Example: `true`
	IsIpv6Enabled *bool `mandatory:"false" json:"isIpv6Enabled"`
}

func (m CreateVcnDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateVcnDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
