// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.us-phoenix-1.oraclecloud.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateLoadBalancerDetails The configuration details for creating a load balancer.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
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

	// Whether the load balancer has a VCN-local (private) IP address.
	// If "true", the service assigns a private IP address to the load balancer.
	// If "false", the service assigns a public IP address to the load balancer.
	// A public load balancer is accessible from the internet, depending on your VCN's
	// security list rules (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/securitylists.htm). For more information about public and
	// private load balancers, see How Load Balancing Works (https://docs.us-phoenix-1.oraclecloud.com/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	Listeners map[string]ListenerDetails `mandatory:"false" json:"listeners"`

	Hostnames map[string]HostnameDetails `mandatory:"false" json:"hostnames"`

	BackendSets map[string]BackendSetDetails `mandatory:"false" json:"backendSets"`

	Certificates map[string]CertificateDetails `mandatory:"false" json:"certificates"`

	PathRouteSets map[string]PathRouteSetDetails `mandatory:"false" json:"pathRouteSets"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	RuleSets map[string]RuleSetDetails `mandatory:"false" json:"ruleSets"`
}

func (m CreateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}
