// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing Service API
//
// API for the Load Balancing Service
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateLoadBalancerDetails The configuration details for creating a load balancer.
type CreateLoadBalancerDetails struct {

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment in which to create the load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A template that determines the total pre-provisioned bandwidth (ingress plus egress).
	// To get a list of available shapes, use the ListShapes
	// operation.
	// Example: `100Mbps`
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// An array of subnet OCIDs (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
	SubnetIds []string `mandatory:"true" json:"subnetIds"`

	BackendSets map[string]BackendSetDetails `mandatory:"false" json:"backendSets"`

	Certificates map[string]CertificateDetails `mandatory:"false" json:"certificates"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	Hostnames map[string]HostnameDetails `mandatory:"false" json:"hostnames"`

	// Whether the load balancer has a VCN-local (private) IP address.
	// If "true", the service assigns a private IP address to the load balancer. The load balancer requires only one subnet
	// to host both the primary and secondary load balancers. The private IP address is local to the subnet. The load balancer
	// is accessible only from within the VCN that contains the associated subnet, or as further restricted by your security
	// list rules. The load balancer can route traffic to any backend server that is reachable from the VCN.
	// For a private load balancer, both the primary and secondary load balancer hosts are within the same Availability Domain.
	// If "false", the service assigns a public IP address to the load balancer. A load balancer with a public IP address
	// requires two subnets, each in a different Availability Domain. One subnet hosts the primary load balancer and the other
	// hosts the secondary (standby) load balancer. A public load balancer is accessible from the internet, depending on your
	// VCN's security list rules (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm).
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	Listeners map[string]ListenerDetails `mandatory:"false" json:"listeners"`

	PathRouteSets map[string]PathRouteSetDetails `mandatory:"false" json:"pathRouteSets"`
}

func (m CreateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}
