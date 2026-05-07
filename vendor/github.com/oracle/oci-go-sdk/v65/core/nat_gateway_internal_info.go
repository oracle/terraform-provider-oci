// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NatGatewayInternalInfo Internal operational information about a NAT gateway.
type NatGatewayInternalInfo struct {

	// The OCID for the NAT gateway's compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID for the NAT gateway.
	Id *string `mandatory:"true" json:"id"`

	// Whether the NAT gateway blocks traffic through it. The default is `false`.
	// Example: `false`
	BlockTraffic *bool `mandatory:"true" json:"blockTraffic"`

	// The NAT gateway's current state.
	LifecycleState NatGatewayInternalInfoLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the NAT gateway was created, in the format defined by RFC3339.
	// Example: '2016-08-25T21:10:29.600Z'
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the VCN the NAT gateway
	// belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The public IP pool
	InternalPublicIpPoolName NatGatewayInternalInfoInternalPublicIpPoolNameEnum `mandatory:"false" json:"internalPublicIpPoolName,omitempty"`

	// Set to `FAILOVER_TO_INTERNET` (the default) to allow the replication of
	// customer data from flowing over public Internet,
	// set to `DO_NOT_FAILOVER_TO_INTERNET` to prevent replication of customer
	// data from flowing over public Internet.
	BackboneFailoverPolicy NatGatewayInternalInfoBackboneFailoverPolicyEnum `mandatory:"false" json:"backboneFailoverPolicy,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The IP address associated with the NAT gateway.
	NatIp *string `mandatory:"false" json:"natIp"`

	// The MPLS label which identifies this NAT gateway in encapsulated
	// traffic sent to either the NAT egress or ingress redirectors.
	// This label is scoped by the egress and ingress redirector IPs.
	Label *int `mandatory:"false" json:"label"`

	// The destination substrate IPv4 address, in dot-decimal notation,
	// used to encapsulate egress traffic routed to this NAT gateway
	// (e.g. from Caviums).
	EgressVip *string `mandatory:"false" json:"egressVip"`

	// The destination substrate IPv4 address, in dot-decimal notation,
	// used to encapsulate ingress traffic routed to this NAT gateway
	// (e.g. from the Internet ingress redirector).
	IngressVip *string `mandatory:"false" json:"ingressVip"`

	// indicates the VIP type (Fleet Name) of the associated NatGW
	NatVipType *string `mandatory:"false" json:"natVipType"`

	NextHop *SubstrateRoute `mandatory:"false" json:"nextHop"`

	// The OCID of the Generic Gateway associated with the NAT gateway.
	GgwId *string `mandatory:"false" json:"ggwId"`

	// The OCID of Public IP associated with the NAT gateway.
	NatIpOcid *string `mandatory:"false" json:"natIpOcid"`

	// Whether the NAT gateway was intended to be deleted. The default is `false`.
	// Example: `false`
	PendingDelete *bool `mandatory:"false" json:"pendingDelete"`

	// The Entity ID of the Natgateway
	EntityId *int64 `mandatory:"false" json:"entityId"`

	// The pod identifier associated with the NAT gateway.
	PodId *string `mandatory:"false" json:"podId"`

	// The OCID of the VN entity attachment associated with the NAT gateway.
	VnEntityAttachmentId *string `mandatory:"false" json:"vnEntityAttachmentId"`

	// Per AD vip details of the natgateway pod.
	PerAdNgwVipDetails []PerAdNgwVipDetails `mandatory:"false" json:"perAdNgwVipDetails"`
}

func (m NatGatewayInternalInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatGatewayInternalInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingNatGatewayInternalInfoLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetNatGatewayInternalInfoLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNatGatewayInternalInfoInternalPublicIpPoolNameEnum(string(m.InternalPublicIpPoolName)); !ok && m.InternalPublicIpPoolName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalPublicIpPoolName: %s. Supported values are: %s.", m.InternalPublicIpPoolName, strings.Join(GetNatGatewayInternalInfoInternalPublicIpPoolNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingNatGatewayInternalInfoBackboneFailoverPolicyEnum(string(m.BackboneFailoverPolicy)); !ok && m.BackboneFailoverPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackboneFailoverPolicy: %s. Supported values are: %s.", m.BackboneFailoverPolicy, strings.Join(GetNatGatewayInternalInfoBackboneFailoverPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// NatGatewayInternalInfoInternalPublicIpPoolNameEnum Enum with underlying type: string
type NatGatewayInternalInfoInternalPublicIpPoolNameEnum string

// Set of constants representing the allowable values for NatGatewayInternalInfoInternalPublicIpPoolNameEnum
const (
	NatGatewayInternalInfoInternalPublicIpPoolNameExternal   NatGatewayInternalInfoInternalPublicIpPoolNameEnum = "EXTERNAL"
	NatGatewayInternalInfoInternalPublicIpPoolNameSociEgress NatGatewayInternalInfoInternalPublicIpPoolNameEnum = "SOCI_EGRESS"
)

var mappingNatGatewayInternalInfoInternalPublicIpPoolNameEnum = map[string]NatGatewayInternalInfoInternalPublicIpPoolNameEnum{
	"EXTERNAL":    NatGatewayInternalInfoInternalPublicIpPoolNameExternal,
	"SOCI_EGRESS": NatGatewayInternalInfoInternalPublicIpPoolNameSociEgress,
}

var mappingNatGatewayInternalInfoInternalPublicIpPoolNameEnumLowerCase = map[string]NatGatewayInternalInfoInternalPublicIpPoolNameEnum{
	"external":    NatGatewayInternalInfoInternalPublicIpPoolNameExternal,
	"soci_egress": NatGatewayInternalInfoInternalPublicIpPoolNameSociEgress,
}

// GetNatGatewayInternalInfoInternalPublicIpPoolNameEnumValues Enumerates the set of values for NatGatewayInternalInfoInternalPublicIpPoolNameEnum
func GetNatGatewayInternalInfoInternalPublicIpPoolNameEnumValues() []NatGatewayInternalInfoInternalPublicIpPoolNameEnum {
	values := make([]NatGatewayInternalInfoInternalPublicIpPoolNameEnum, 0)
	for _, v := range mappingNatGatewayInternalInfoInternalPublicIpPoolNameEnum {
		values = append(values, v)
	}
	return values
}

// GetNatGatewayInternalInfoInternalPublicIpPoolNameEnumStringValues Enumerates the set of values in String for NatGatewayInternalInfoInternalPublicIpPoolNameEnum
func GetNatGatewayInternalInfoInternalPublicIpPoolNameEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"SOCI_EGRESS",
	}
}

// GetMappingNatGatewayInternalInfoInternalPublicIpPoolNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNatGatewayInternalInfoInternalPublicIpPoolNameEnum(val string) (NatGatewayInternalInfoInternalPublicIpPoolNameEnum, bool) {
	enum, ok := mappingNatGatewayInternalInfoInternalPublicIpPoolNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NatGatewayInternalInfoLifecycleStateEnum Enum with underlying type: string
type NatGatewayInternalInfoLifecycleStateEnum string

// Set of constants representing the allowable values for NatGatewayInternalInfoLifecycleStateEnum
const (
	NatGatewayInternalInfoLifecycleStateProvisioning NatGatewayInternalInfoLifecycleStateEnum = "PROVISIONING"
	NatGatewayInternalInfoLifecycleStateAvailable    NatGatewayInternalInfoLifecycleStateEnum = "AVAILABLE"
	NatGatewayInternalInfoLifecycleStateTerminating  NatGatewayInternalInfoLifecycleStateEnum = "TERMINATING"
	NatGatewayInternalInfoLifecycleStateTerminated   NatGatewayInternalInfoLifecycleStateEnum = "TERMINATED"
)

var mappingNatGatewayInternalInfoLifecycleStateEnum = map[string]NatGatewayInternalInfoLifecycleStateEnum{
	"PROVISIONING": NatGatewayInternalInfoLifecycleStateProvisioning,
	"AVAILABLE":    NatGatewayInternalInfoLifecycleStateAvailable,
	"TERMINATING":  NatGatewayInternalInfoLifecycleStateTerminating,
	"TERMINATED":   NatGatewayInternalInfoLifecycleStateTerminated,
}

var mappingNatGatewayInternalInfoLifecycleStateEnumLowerCase = map[string]NatGatewayInternalInfoLifecycleStateEnum{
	"provisioning": NatGatewayInternalInfoLifecycleStateProvisioning,
	"available":    NatGatewayInternalInfoLifecycleStateAvailable,
	"terminating":  NatGatewayInternalInfoLifecycleStateTerminating,
	"terminated":   NatGatewayInternalInfoLifecycleStateTerminated,
}

// GetNatGatewayInternalInfoLifecycleStateEnumValues Enumerates the set of values for NatGatewayInternalInfoLifecycleStateEnum
func GetNatGatewayInternalInfoLifecycleStateEnumValues() []NatGatewayInternalInfoLifecycleStateEnum {
	values := make([]NatGatewayInternalInfoLifecycleStateEnum, 0)
	for _, v := range mappingNatGatewayInternalInfoLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetNatGatewayInternalInfoLifecycleStateEnumStringValues Enumerates the set of values in String for NatGatewayInternalInfoLifecycleStateEnum
func GetNatGatewayInternalInfoLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingNatGatewayInternalInfoLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNatGatewayInternalInfoLifecycleStateEnum(val string) (NatGatewayInternalInfoLifecycleStateEnum, bool) {
	enum, ok := mappingNatGatewayInternalInfoLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// NatGatewayInternalInfoBackboneFailoverPolicyEnum Enum with underlying type: string
type NatGatewayInternalInfoBackboneFailoverPolicyEnum string

// Set of constants representing the allowable values for NatGatewayInternalInfoBackboneFailoverPolicyEnum
const (
	NatGatewayInternalInfoBackboneFailoverPolicyFailoverToInternet      NatGatewayInternalInfoBackboneFailoverPolicyEnum = "FAILOVER_TO_INTERNET"
	NatGatewayInternalInfoBackboneFailoverPolicyDoNotFailoverToInternet NatGatewayInternalInfoBackboneFailoverPolicyEnum = "DO_NOT_FAILOVER_TO_INTERNET"
)

var mappingNatGatewayInternalInfoBackboneFailoverPolicyEnum = map[string]NatGatewayInternalInfoBackboneFailoverPolicyEnum{
	"FAILOVER_TO_INTERNET":        NatGatewayInternalInfoBackboneFailoverPolicyFailoverToInternet,
	"DO_NOT_FAILOVER_TO_INTERNET": NatGatewayInternalInfoBackboneFailoverPolicyDoNotFailoverToInternet,
}

var mappingNatGatewayInternalInfoBackboneFailoverPolicyEnumLowerCase = map[string]NatGatewayInternalInfoBackboneFailoverPolicyEnum{
	"failover_to_internet":        NatGatewayInternalInfoBackboneFailoverPolicyFailoverToInternet,
	"do_not_failover_to_internet": NatGatewayInternalInfoBackboneFailoverPolicyDoNotFailoverToInternet,
}

// GetNatGatewayInternalInfoBackboneFailoverPolicyEnumValues Enumerates the set of values for NatGatewayInternalInfoBackboneFailoverPolicyEnum
func GetNatGatewayInternalInfoBackboneFailoverPolicyEnumValues() []NatGatewayInternalInfoBackboneFailoverPolicyEnum {
	values := make([]NatGatewayInternalInfoBackboneFailoverPolicyEnum, 0)
	for _, v := range mappingNatGatewayInternalInfoBackboneFailoverPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetNatGatewayInternalInfoBackboneFailoverPolicyEnumStringValues Enumerates the set of values in String for NatGatewayInternalInfoBackboneFailoverPolicyEnum
func GetNatGatewayInternalInfoBackboneFailoverPolicyEnumStringValues() []string {
	return []string{
		"FAILOVER_TO_INTERNET",
		"DO_NOT_FAILOVER_TO_INTERNET",
	}
}

// GetMappingNatGatewayInternalInfoBackboneFailoverPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingNatGatewayInternalInfoBackboneFailoverPolicyEnum(val string) (NatGatewayInternalInfoBackboneFailoverPolicyEnum, bool) {
	enum, ok := mappingNatGatewayInternalInfoBackboneFailoverPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
