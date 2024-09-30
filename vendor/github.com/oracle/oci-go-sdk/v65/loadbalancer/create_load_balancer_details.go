// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateLoadBalancerDetails The configuration details for creating a load balancer.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type CreateLoadBalancerDetails struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which to create the load balancer.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// A template that determines the total pre-provisioned bandwidth (ingress plus egress).
	// To get a list of available shapes, use the ListShapes
	// operation.
	// Example: `flexible`
	// NOTE: After May 2023, Fixed shapes - 10Mbps, 100Mbps, 400Mbps, 8000Mbps would be deprecated and only shape
	//       allowed would be `Flexible`
	ShapeName *string `mandatory:"true" json:"shapeName"`

	// An array of subnet OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	SubnetIds []string `mandatory:"true" json:"subnetIds"`

	// The configuration details to create load balancer using Flexible shape. This is required only if shapeName is `Flexible`.
	ShapeDetails *ShapeDetails `mandatory:"false" json:"shapeDetails"`

	// Whether the load balancer has a VCN-local (private) IP address.
	// If "true", the service assigns a private IP address to the load balancer.
	// If "false", the service assigns a public IP address to the load balancer.
	// A public load balancer is accessible from the internet, depending on your VCN's
	// security list rules (https://docs.cloud.oracle.com/Content/Network/Concepts/securitylists.htm). For more information about public and
	// private load balancers, see How Load Balancing Works (https://docs.cloud.oracle.com/Content/Balance/Concepts/balanceoverview.htm#how-load-balancing-works).
	// Example: `true`
	IsPrivate *bool `mandatory:"false" json:"isPrivate"`

	// Whether or not the load balancer has delete protection enabled.
	// If "true", the loadbalancer will be protected against deletion if configured to accept traffic.
	// If "false", the loadbalancer will not be protected against deletion.
	// Delete protection will not be enabled unless a value of "true" is provided.
	// Example: `true`
	IsDeleteProtectionEnabled *bool `mandatory:"false" json:"isDeleteProtectionEnabled"`

	// Whether the load balancer has an IPv4 or IPv6 IP address.
	// If "IPV4", the service assigns an IPv4 address and the load balancer supports IPv4 traffic.
	// If "IPV6", the service assigns an IPv6 address and the load balancer supports IPv6 traffic.
	// Example: "ipMode":"IPV6"
	IpMode CreateLoadBalancerDetailsIpModeEnum `mandatory:"false" json:"ipMode,omitempty"`

	// Whether or not the load balancer has the Request Id feature enabled for HTTP listeners.
	// If "true", the load balancer will attach a unique request id header to every request
	// passed through from the load balancer to load balancer backends. This same request id
	// header also will be added to the response the lb received from the backend handling
	// the request before the load balancer returns the response to the requestor. The name
	// of the unique request id header is set the by value of requestIdHeader.
	// If "false", the loadbalancer not add this unique request id header to either the request
	// passed through to the load balancer backends nor to the reponse returned to the user.
	// New load balancers have the Request Id feature disabled unless isRequestIdEnabled is set to true.
	// Example: `true`
	IsRequestIdEnabled *bool `mandatory:"false" json:"isRequestIdEnabled"`

	// If isRequestIdEnabled is true then this field contains the name of the header field
	// that contains the unique request id that is attached to every request from
	// the load balancer to the load balancer backends and to every response from the load
	// balancer.
	// If a request to the load balancer already contains a header with same name as specified
	// in requestIdHeader then the load balancer will not change the value of that field.
	// If isRequestIdEnabled is false then this field is ignored.
	// If this field is not set or is set to "" then this field defaults to X-Request-Id
	// **Notes:**
	// * Unless the header name is "" it must start with "X-" prefix.
	// * Setting the header name to "" will set it to the default: X-Request-Id.
	RequestIdHeader *string `mandatory:"false" json:"requestIdHeader"`

	// An array of reserved Ips.
	ReservedIps []ReservedIp `mandatory:"false" json:"reservedIps"`

	Listeners map[string]ListenerDetails `mandatory:"false" json:"listeners"`

	Hostnames map[string]HostnameDetails `mandatory:"false" json:"hostnames"`

	BackendSets map[string]BackendSetDetails `mandatory:"false" json:"backendSets"`

	// An array of NSG OCIDs (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) associated with this load balancer.
	// During the load balancer's creation, the service adds the new load balancer to the specified NSGs.
	// The benefits of using NSGs with the load balancer include:
	// *  NSGs define network security rules to govern ingress and egress traffic for the load balancer.
	// *  The network security rules of other resources can reference the NSGs associated with the load balancer
	//    to ensure access.
	// Example: `["ocid1.nsg.oc1.phx.unique_ID"]`
	NetworkSecurityGroupIds []string `mandatory:"false" json:"networkSecurityGroupIds"`

	Certificates map[string]CertificateDetails `mandatory:"false" json:"certificates"`

	SslCipherSuites map[string]SslCipherSuiteDetails `mandatory:"false" json:"sslCipherSuites"`

	PathRouteSets map[string]PathRouteSetDetails `mandatory:"false" json:"pathRouteSets"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Extended Defined tags for ZPR for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit", "usagetype" : "zpr"}}}`
	ZprTags map[string]map[string]interface{} `mandatory:"false" json:"zprTags"`

	RuleSets map[string]RuleSetDetails `mandatory:"false" json:"ruleSets"`
}

func (m CreateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateLoadBalancerDetailsIpModeEnum(string(m.IpMode)); !ok && m.IpMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpMode: %s. Supported values are: %s.", m.IpMode, strings.Join(GetCreateLoadBalancerDetailsIpModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateLoadBalancerDetailsIpModeEnum Enum with underlying type: string
type CreateLoadBalancerDetailsIpModeEnum string

// Set of constants representing the allowable values for CreateLoadBalancerDetailsIpModeEnum
const (
	CreateLoadBalancerDetailsIpModeIpv4 CreateLoadBalancerDetailsIpModeEnum = "IPV4"
	CreateLoadBalancerDetailsIpModeIpv6 CreateLoadBalancerDetailsIpModeEnum = "IPV6"
)

var mappingCreateLoadBalancerDetailsIpModeEnum = map[string]CreateLoadBalancerDetailsIpModeEnum{
	"IPV4": CreateLoadBalancerDetailsIpModeIpv4,
	"IPV6": CreateLoadBalancerDetailsIpModeIpv6,
}

var mappingCreateLoadBalancerDetailsIpModeEnumLowerCase = map[string]CreateLoadBalancerDetailsIpModeEnum{
	"ipv4": CreateLoadBalancerDetailsIpModeIpv4,
	"ipv6": CreateLoadBalancerDetailsIpModeIpv6,
}

// GetCreateLoadBalancerDetailsIpModeEnumValues Enumerates the set of values for CreateLoadBalancerDetailsIpModeEnum
func GetCreateLoadBalancerDetailsIpModeEnumValues() []CreateLoadBalancerDetailsIpModeEnum {
	values := make([]CreateLoadBalancerDetailsIpModeEnum, 0)
	for _, v := range mappingCreateLoadBalancerDetailsIpModeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateLoadBalancerDetailsIpModeEnumStringValues Enumerates the set of values in String for CreateLoadBalancerDetailsIpModeEnum
func GetCreateLoadBalancerDetailsIpModeEnumStringValues() []string {
	return []string{
		"IPV4",
		"IPV6",
	}
}

// GetMappingCreateLoadBalancerDetailsIpModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateLoadBalancerDetailsIpModeEnum(val string) (CreateLoadBalancerDetailsIpModeEnum, bool) {
	enum, ok := mappingCreateLoadBalancerDetailsIpModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
