// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateLoadBalancerDetails Configuration details to update a load balancer.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateLoadBalancerDetails struct {

	// The user-friendly display name for the load balancer. It does not have to be unique, and it is changeable.
	// Avoid entering confidential information.
	// Example: `example_load_balancer`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Whether the load balancer has an IPv4 or IPv6 IP address.
	//   If "IPV4", the service assigns an IPv4 address and the load balancer supports IPv4 traffic.
	//   If "IPV6", the service assigns an IPv6 address and the load balancer supports IPv6 traffic.
	//   Example: "ipMode":"IPV6"
	IpMode UpdateLoadBalancerDetailsIpModeEnum `mandatory:"false" json:"ipMode,omitempty"`

	// Used to disambiguate which subnet prefix should be used to create an IPv6 LB.
	// Example: "2002::1234:abcd:ffff:c0a8:101/64"
	Ipv6SubnetCidr *string `mandatory:"false" json:"ipv6SubnetCidr"`

	// An array of reserved Ips.
	ReservedIps []ReservedIp `mandatory:"false" json:"reservedIps"`

	// Whether or not the load balancer has delete protection enabled.
	// If "true", the loadbalancer will be protected against deletion if configured to accept traffic.
	// If "false", the loadbalancer will not be protected against deletion.
	// If null or unset, the value for delete protection will not be changed.
	// Example: `true`
	IsDeleteProtectionEnabled *bool `mandatory:"false" json:"isDeleteProtectionEnabled"`

	// Whether or not the load balancer has the Request Id feature enabled for HTTP listeners.
	// If "true", the load balancer will attach a unique request id header to every request
	// passed through from the load balancer to load balancer backends. This same request id
	// header also will be added to the response the lb received from the backend handling
	// the request before the load balancer returns the response to the requestor. The name
	// of the unique request id header is set the by value of requestIdHeader.
	// If "false", the loadbalancer not add this unique request id header to either the request
	// passed through to the load balancer backends nor to the reponse returned to the user.
	// New load balancers have the Request Id feature enabled unless isRequestIdEnabled is set to False.
	// Example: `true`
	IsRequestIdEnabled *bool `mandatory:"false" json:"isRequestIdEnabled"`

	// If isRequestIdEnabled is true then this field contains the name of the header field
	// that contains the unique request id that is attached to every request from
	// the load balancer to the load balancer backends and to every response from the load
	// balancer.
	// If a request to the load balancer already contains a header with same name as specified
	// in requestIdHeader then the load balancer will not change the value of that field.
	// If isRequestIdEnabled is false then this field is ignored.
	// **Notes:**
	// * Unless the header name is "" it must start with "X-" prefix.
	// * Setting the header name to "" will set it to the default: X-Request-Id.
	RequestIdHeader *string `mandatory:"false" json:"requestIdHeader"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Extended Defined tags for ZPR for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit", "usagetype" : "zpr"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`
}

func (m UpdateLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateLoadBalancerDetailsIpModeEnum(string(m.IpMode)); !ok && m.IpMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpMode: %s. Supported values are: %s.", m.IpMode, strings.Join(GetUpdateLoadBalancerDetailsIpModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateLoadBalancerDetailsIpModeEnum Enum with underlying type: string
type UpdateLoadBalancerDetailsIpModeEnum string

// Set of constants representing the allowable values for UpdateLoadBalancerDetailsIpModeEnum
const (
	UpdateLoadBalancerDetailsIpModeIpv4 UpdateLoadBalancerDetailsIpModeEnum = "IPV4"
	UpdateLoadBalancerDetailsIpModeIpv6 UpdateLoadBalancerDetailsIpModeEnum = "IPV6"
)

var mappingUpdateLoadBalancerDetailsIpModeEnum = map[string]UpdateLoadBalancerDetailsIpModeEnum{
	"IPV4": UpdateLoadBalancerDetailsIpModeIpv4,
	"IPV6": UpdateLoadBalancerDetailsIpModeIpv6,
}

var mappingUpdateLoadBalancerDetailsIpModeEnumLowerCase = map[string]UpdateLoadBalancerDetailsIpModeEnum{
	"ipv4": UpdateLoadBalancerDetailsIpModeIpv4,
	"ipv6": UpdateLoadBalancerDetailsIpModeIpv6,
}

// GetUpdateLoadBalancerDetailsIpModeEnumValues Enumerates the set of values for UpdateLoadBalancerDetailsIpModeEnum
func GetUpdateLoadBalancerDetailsIpModeEnumValues() []UpdateLoadBalancerDetailsIpModeEnum {
	values := make([]UpdateLoadBalancerDetailsIpModeEnum, 0)
	for _, v := range mappingUpdateLoadBalancerDetailsIpModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateLoadBalancerDetailsIpModeEnumStringValues Enumerates the set of values in String for UpdateLoadBalancerDetailsIpModeEnum
func GetUpdateLoadBalancerDetailsIpModeEnumStringValues() []string {
	return []string{
		"IPV4",
		"IPV6",
	}
}

// GetMappingUpdateLoadBalancerDetailsIpModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateLoadBalancerDetailsIpModeEnum(val string) (UpdateLoadBalancerDetailsIpModeEnum, bool) {
	enum, ok := mappingUpdateLoadBalancerDetailsIpModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
