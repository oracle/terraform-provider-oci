// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// EnableReverseConnectionsDetails Details for enabling reverse connections for a private endpoint.
type EnableReverseConnectionsDetails struct {

	// A list of IP addresses in the customer VCN to be used as the source IPs for reverse connection packets
	// traveling from the service's VCN to the customer's VCN. If no list is specified or
	// an empty list is provided, an IP address will be chosen from the customer subnet's CIDR.
	ReverseConnectionsSourceIps []ReverseConnectionsSourceIpDetails `mandatory:"false" json:"reverseConnectionsSourceIps"`

	// Whether a DNS proxy should be configured for the reverse connection. If the service
	// does not intend to use DNS FQDN to communicate to customer endpoints, set this to `false`.
	// Example: `false`
	IsProxyEnabled *bool `mandatory:"false" json:"isProxyEnabled"`

	// The IP address in the service VCN to be used to reach the DNS proxy that resolves the
	// customer FQDN for reverse connections. If no value is provided, an available IP address will
	// be chosen from the service subnet's CIDR.
	DnsProxyIp *string `mandatory:"false" json:"dnsProxyIp"`

	// The context in which the DNS proxy will resolve the DNS queries. The default is `SERVICE`.
	// For example: if the service does not know the specific DNS zones for the customer VCNs, set
	// this to `CUSTOMER`, and set `excludedDnsZones` to the list of DNS zones in your service
	// provider VCN.
	// Allowed values:
	//  * `SERVICE`: All DNS queries will be resolved within the service VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	//  * `CUSTOMER`: All DNS queries will be resolved within the customer VCN's DNS context,
	//    unless the FQDN belongs to one of zones in the `excludedDnsZones` list.
	DefaultDnsResolutionContext EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum `mandatory:"false" json:"defaultDnsResolutionContext,omitempty"`

	// List of DNS zones to exclude from the default DNS resolution context.
	ExcludedDnsZones []string `mandatory:"false" json:"excludedDnsZones"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the service's subnet where
	// the DNS proxy endpoint will be created.
	ServiceSubnetId *string `mandatory:"false" json:"serviceSubnetId"`

	// A list of the OCIDs of the network security groups that the reverse connection's VNIC belongs to.
	// For more information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// Number of customer endpoints that the service provider expects to establish connections to using this RCE. The default is 0.
	// When non-zero value is specified, reverse connection configuration will be allocated with a list of CIDRs, from
	// which NAT IP addresses will be allocated. These list of CIDRs will not be shared by other reverse
	// connection enabled private endpoints.
	// When zero is specified, reverse connection configuration will get NAT IP addresses from common pool of CIDRs,
	// which will be shared with other reverse connection enabled private endpoints.
	// If the private endpoint was enabled with reverse connection with 0 already, the field is not updatable.
	// The size may not be updated with smaller number than previously specified value, but may be increased.
	CustomerEndpointsSize *int `mandatory:"false" json:"customerEndpointsSize"`

	// Layer 4 transport protocol to be used when resolving DNS queries within the default DNS resolution context.
	DefaultDnsContextTransport EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum `mandatory:"false" json:"defaultDnsContextTransport,omitempty"`

	// List of CIDRs that this reverse connection configuration will allocate the NAT IP addresses from.
	// CIDRs on this list should not be shared by other reverse connection enabled private endpoints.
	// When not specified, if the customerEndpointSize is non null, reverse connection configuration will get
	// NAT IP addresses from the dedicated pool of CIDRs, else will get specified from the common pool of CIDRs.
	// This field cannot be specified if the customerEndpointsSize field is non null and vice versa.
	ReverseConnectionNatIpCidrs []string `mandatory:"false" json:"reverseConnectionNatIpCidrs"`
}

func (m EnableReverseConnectionsDetails) String() string {
	return common.PointerString(m)
}

// EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum Enum with underlying type: string
type EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum string

// Set of constants representing the allowable values for EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum
const (
	EnableReverseConnectionsDetailsDefaultDnsResolutionContextService  EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum = "SERVICE"
	EnableReverseConnectionsDetailsDefaultDnsResolutionContextCustomer EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum = "CUSTOMER"
)

var mappingEnableReverseConnectionsDetailsDefaultDnsResolutionContext = map[string]EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum{
	"SERVICE":  EnableReverseConnectionsDetailsDefaultDnsResolutionContextService,
	"CUSTOMER": EnableReverseConnectionsDetailsDefaultDnsResolutionContextCustomer,
}

// GetEnableReverseConnectionsDetailsDefaultDnsResolutionContextEnumValues Enumerates the set of values for EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum
func GetEnableReverseConnectionsDetailsDefaultDnsResolutionContextEnumValues() []EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum {
	values := make([]EnableReverseConnectionsDetailsDefaultDnsResolutionContextEnum, 0)
	for _, v := range mappingEnableReverseConnectionsDetailsDefaultDnsResolutionContext {
		values = append(values, v)
	}
	return values
}

// EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum Enum with underlying type: string
type EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum string

// Set of constants representing the allowable values for EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum
const (
	EnableReverseConnectionsDetailsDefaultDnsContextTransportTcp EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum = "TCP"
	EnableReverseConnectionsDetailsDefaultDnsContextTransportUdp EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum = "UDP"
)

var mappingEnableReverseConnectionsDetailsDefaultDnsContextTransport = map[string]EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum{
	"TCP": EnableReverseConnectionsDetailsDefaultDnsContextTransportTcp,
	"UDP": EnableReverseConnectionsDetailsDefaultDnsContextTransportUdp,
}

// GetEnableReverseConnectionsDetailsDefaultDnsContextTransportEnumValues Enumerates the set of values for EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum
func GetEnableReverseConnectionsDetailsDefaultDnsContextTransportEnumValues() []EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum {
	values := make([]EnableReverseConnectionsDetailsDefaultDnsContextTransportEnum, 0)
	for _, v := range mappingEnableReverseConnectionsDetailsDefaultDnsContextTransport {
		values = append(values, v)
	}
	return values
}
