// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// NetworkLoadBalancer API
//
// This describes the network load balancer API.
//

package networkloadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateNetworkLoadBalancerDetails Configuration details to update a network load balancer.
// **Caution:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type UpdateNetworkLoadBalancerDetails struct {

	// The user-friendly display name for the network load balancer, which does not have to be unique and can be changed.
	// Avoid entering confidential information.
	// Example: `example_network_load_balancer`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// This parameter can be enabled only if backends are compute OCIDs. When enabled, the skipSourceDestinationCheck parameter is automatically
	// enabled on the load balancer VNIC, and packets are sent to the backend with the entire IP header intact.
	IsPreserveSourceDestination *bool `mandatory:"false" json:"isPreserveSourceDestination"`

	// This can only be enabled when NLB is working in transparent mode with source destination header preservation enabled.
	// This removes the additional dependency from NLB backends(like Firewalls) to perform SNAT.
	IsSymmetricHashEnabled *bool `mandatory:"false" json:"isSymmetricHashEnabled"`

	// IP version associated with the NLB.
	NlbIpVersion NlbIpVersionEnum `mandatory:"false" json:"nlbIpVersion,omitempty"`

	// IPv6 subnet prefix selection. If Ipv6 subnet prefix is passed, Nlb Ipv6 Address would be assign within the cidr block. NLB has to be dual or single stack ipv6 to support this.
	SubnetIpv6Cidr *string `mandatory:"false" json:"subnetIpv6Cidr"`

	// IPv6 address to be assigned to the network load balancer being created.
	// This IP address has to be part of one of the prefixes supported by the subnet.
	// Example: "2607:9b80:9a0a:9a7e:abcd:ef01:2345:6789"
	AssignedIpv6 *string `mandatory:"false" json:"assignedIpv6"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// ZPR tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"oracle-zpr": {"td": {"value": "42", "mode": "audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`
}

func (m UpdateNetworkLoadBalancerDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateNetworkLoadBalancerDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNlbIpVersionEnum(string(m.NlbIpVersion)); !ok && m.NlbIpVersion != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NlbIpVersion: %s. Supported values are: %s.", m.NlbIpVersion, strings.Join(GetNlbIpVersionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
