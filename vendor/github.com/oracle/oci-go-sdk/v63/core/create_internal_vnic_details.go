// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v63/common"
	"strings"
)

// CreateInternalVnicDetails This structure is used when creating vnic for internal clients.
// For more information about VNICs, see
// Virtual Network Interface Cards (VNICs) (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm).
type CreateInternalVnicDetails struct {

	// Whether the VNIC should be assigned a public IP address. Defaults to whether
	// the subnet is public or private. If not set and the VNIC is being created
	// in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the
	// Subnet), then no public IP address is assigned.
	// If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then
	// a public IP address is assigned. If set to true and
	// `prohibitPublicIpOnVnic` = true, an error is returned.
	// **Note:** This public IP address is associated with the primary private IP
	// on the VNIC. For more information, see
	// IP Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPaddresses.htm).
	// **Note:** There's a limit to the number of PublicIp
	// a VNIC or instance can have. If you try to create a secondary VNIC
	// with an assigned public IP for an instance that has already
	// reached its public IP limit, an error is returned. For information
	// about the public IP limits, see
	// Public IP Addresses (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingpublicIPs.htm).
	// Example: `false`
	AssignPublicIp *bool `mandatory:"false" json:"assignPublicIp"`

	// The availability domain of the instance.
	// Availability domain can not be provided if isServiceVnic is true, it is required otherwise.
	//   Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"false" json:"availabilityDomain"`

	// Not for general use!
	// Contact sic_vcn_us_grp@oracle.com before setting this flag.
	// Indicates that the Cavium should not enforce Internet ingress/egress throttling.
	// Defaults to `false`, in which case we do enforce that throttling.
	// At least one of subnetId OR the vlanId are required
	BypassInternetThrottle *bool `mandatory:"false" json:"bypassInternetThrottle"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname for the VNIC's primary private IP. Used for DNS. The value is the hostname
	// portion of the primary private IP's fully qualified domain name (FQDN)
	// (for example, `bminstance-1` in FQDN `bminstance-1.subnet123.vcn1.oraclevcn.com`).
	// Must be unique across all VNICs in the subnet and comply with
	// RFC 952 (https://tools.ietf.org/html/rfc952) and
	// RFC 1123 (https://tools.ietf.org/html/rfc1123).
	// The value appears in the Vnic object and also the
	// PrivateIp object returned by
	// ListPrivateIps and
	// GetPrivateIp.
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/dns.htm).
	// When launching an instance, use this `hostnameLabel` instead
	// of the deprecated `hostnameLabel` in
	// LaunchInstanceDetails.
	// If you provide both, the values must match.
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Indicates if the VNIC is associated with (and will be attached to) a BM instance.
	IsBmVnic *bool `mandatory:"false" json:"isBmVnic"`

	// Indicates if the VNIC is bridge which means cavium will ARP for its MAC address.
	IsBridgeVnic *bool `mandatory:"false" json:"isBridgeVnic"`

	// Indicates if this VNIC can issue GARP requests. False by default.
	IsGarpEnabled *bool `mandatory:"false" json:"isGarpEnabled"`

	// Indicates if MAC learning is enabled for the VNIC. The default is `false`.
	// When this flag is enabled, then VCN CP does not allocate MAC address,
	// hence MAC address will be set as null as part of the VNIC that is returned.
	IsMacLearningEnabled *bool `mandatory:"false" json:"isMacLearningEnabled"`

	// Indicates if the VNIC is managed by a internal partner team. And customer is not allowed
	// the perform update/delete operations on it directly.
	// Defaults to `False`
	IsManaged *bool `mandatory:"false" json:"isManaged"`

	// Indicates if the VNIC is primary which means it cannot be detached.
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// Indicates if the VNIC is a service vnic.
	IsServiceVnic *bool `mandatory:"false" json:"isServiceVnic"`

	// Only provided when no publicIpPoolId is specified.
	InternalPoolName CreateInternalVnicDetailsInternalPoolNameEnum `mandatory:"false" json:"internalPoolName,omitempty"`

	// The overlay MAC address of the instance
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// A list of the OCIDs of the network security groups (NSGs) to add the VNIC to. For more
	// information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// ID of the entity owning the VNIC. This is passed in the create vnic call.
	// If none is passed and if there is an attachment then the attached instanceId is the ownerId.
	OwnerId *string `mandatory:"false" json:"ownerId"`

	// A private IP address of your choice to assign to the VNIC. Must be an
	// available IP address within the subnet's CIDR. If you don't specify a
	// value, Oracle automatically assigns a private IP address from the subnet.
	// This is the VNIC's *primary* private IP address. The value appears in
	// the Vnic object and also the
	// PrivateIp object returned by
	// ListPrivateIps and
	// GetPrivateIp.
	// Example: `10.0.3.3`
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the pool object created by the current tenancy
	PublicIpPoolId *string `mandatory:"false" json:"publicIpPoolId"`

	// ID of the customer visible upstream resource that the VNIC is associated with. This property is
	// exposed to customers as part of API to list members of a network security group.
	// For example, if the VNIC is associated with a loadbalancer or dbsystem instance, then it needs
	// to be set to corresponding customer visible loadbalancer or dbsystem instance OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	// Note that the partner team creating/managing the VNIC is owner of this metadata.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Type of the customer visible upstream resource that the VNIC is associated with. This property can be
	// exposed to customers as part of API to list members of a network security group.
	// For example, it can be set as,
	//  - `loadbalancer` if corresponding resourceId is a loadbalancer instance's OCID
	//  - `dbsystem` if corresponding resourceId is a dbsystem instance's OCID
	// Note that the partner team creating/managing the VNIC is owner of this metadata.
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// Whether the source/destination check is disabled on the VNIC.
	// Defaults to `false`, which means the check is performed. For information
	// about why you would skip the source/destination check, see
	// Using a Private IP as a Route Target (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingroutetables.htm#privateip).
	// Example: `true`
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet to create the VNIC in. When launching an instance,
	// use this `subnetId` instead of the deprecated `subnetId` in
	// LaunchInstanceDetails.
	// At least one of them is required; if you provide both, the values must match.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VLAN that the VNIC belongs to
	VlanId *string `mandatory:"false" json:"vlanId"`

	// ID of the compartment
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// Indicates if generation of NAT IP should be skipped if the associated VNIC has this flag set to true.
	IsNatIpAllocationDisabled *bool `mandatory:"false" json:"isNatIpAllocationDisabled"`

	// Indicates if private IP creation should be blocked for external customers. Default to false.
	// For example, Exadata team can create private IP through internal api. External customers who call public api
	// are prohibited to add private IP to Exadata node.
	IsPrivateIpCreationBlocked *bool `mandatory:"false" json:"isPrivateIpCreationBlocked"`

	// Indicates if the VNIC should get a public IP.
	HasPublicIp *bool `mandatory:"false" json:"hasPublicIp"`

	// Indicates if the VNIC is connected to VNIC as a Service. Defaults to `False`.
	IsVnicServiceVnic *bool `mandatory:"false" json:"isVnicServiceVnic"`

	// MPLS label to be used with a VNIC connected to VNIC as a Service. Required if isVnicServiceVnic is `True`.
	ServiceMplsLabel *int `mandatory:"false" json:"serviceMplsLabel"`

	// Type of service VNIC. Feature or use case that is creating this service VNIC. Used for forecasting, resource limits enforcement, and capacity management.
	ServiceVnicType CreateInternalVnicDetailsServiceVnicTypeEnum `mandatory:"false" json:"serviceVnicType,omitempty"`

	// Shape of VNIC that will be used to allocate resource in the data plane once the VNIC is attached
	VnicShape CreateInternalVnicDetailsVnicShapeEnum `mandatory:"false" json:"vnicShape,omitempty"`
}

func (m CreateInternalVnicDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalVnicDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalVnicDetailsInternalPoolNameEnum(string(m.InternalPoolName)); !ok && m.InternalPoolName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalPoolName: %s. Supported values are: %s.", m.InternalPoolName, strings.Join(GetCreateInternalVnicDetailsInternalPoolNameEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateInternalVnicDetailsServiceVnicTypeEnum(string(m.ServiceVnicType)); !ok && m.ServiceVnicType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ServiceVnicType: %s. Supported values are: %s.", m.ServiceVnicType, strings.Join(GetCreateInternalVnicDetailsServiceVnicTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreateInternalVnicDetailsVnicShapeEnum(string(m.VnicShape)); !ok && m.VnicShape != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for VnicShape: %s. Supported values are: %s.", m.VnicShape, strings.Join(GetCreateInternalVnicDetailsVnicShapeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalVnicDetailsInternalPoolNameEnum Enum with underlying type: string
type CreateInternalVnicDetailsInternalPoolNameEnum string

// Set of constants representing the allowable values for CreateInternalVnicDetailsInternalPoolNameEnum
const (
	CreateInternalVnicDetailsInternalPoolNameExternal   CreateInternalVnicDetailsInternalPoolNameEnum = "EXTERNAL"
	CreateInternalVnicDetailsInternalPoolNameSociEgress CreateInternalVnicDetailsInternalPoolNameEnum = "SOCI_EGRESS"
)

var mappingCreateInternalVnicDetailsInternalPoolNameEnum = map[string]CreateInternalVnicDetailsInternalPoolNameEnum{
	"EXTERNAL":    CreateInternalVnicDetailsInternalPoolNameExternal,
	"SOCI_EGRESS": CreateInternalVnicDetailsInternalPoolNameSociEgress,
}

var mappingCreateInternalVnicDetailsInternalPoolNameEnumLowerCase = map[string]CreateInternalVnicDetailsInternalPoolNameEnum{
	"external":    CreateInternalVnicDetailsInternalPoolNameExternal,
	"soci_egress": CreateInternalVnicDetailsInternalPoolNameSociEgress,
}

// GetCreateInternalVnicDetailsInternalPoolNameEnumValues Enumerates the set of values for CreateInternalVnicDetailsInternalPoolNameEnum
func GetCreateInternalVnicDetailsInternalPoolNameEnumValues() []CreateInternalVnicDetailsInternalPoolNameEnum {
	values := make([]CreateInternalVnicDetailsInternalPoolNameEnum, 0)
	for _, v := range mappingCreateInternalVnicDetailsInternalPoolNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalVnicDetailsInternalPoolNameEnumStringValues Enumerates the set of values in String for CreateInternalVnicDetailsInternalPoolNameEnum
func GetCreateInternalVnicDetailsInternalPoolNameEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"SOCI_EGRESS",
	}
}

// GetMappingCreateInternalVnicDetailsInternalPoolNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalVnicDetailsInternalPoolNameEnum(val string) (CreateInternalVnicDetailsInternalPoolNameEnum, bool) {
	enum, ok := mappingCreateInternalVnicDetailsInternalPoolNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateInternalVnicDetailsServiceVnicTypeEnum Enum with underlying type: string
type CreateInternalVnicDetailsServiceVnicTypeEnum string

// Set of constants representing the allowable values for CreateInternalVnicDetailsServiceVnicTypeEnum
const (
	CreateInternalVnicDetailsServiceVnicTypePrivateEndpoint           CreateInternalVnicDetailsServiceVnicTypeEnum = "PRIVATE_ENDPOINT"
	CreateInternalVnicDetailsServiceVnicTypeReverseConnectionEndpoint CreateInternalVnicDetailsServiceVnicTypeEnum = "REVERSE_CONNECTION_ENDPOINT"
	CreateInternalVnicDetailsServiceVnicTypeRealVirtualRouter         CreateInternalVnicDetailsServiceVnicTypeEnum = "REAL_VIRTUAL_ROUTER"
	CreateInternalVnicDetailsServiceVnicTypePrivateDnsEndpoint        CreateInternalVnicDetailsServiceVnicTypeEnum = "PRIVATE_DNS_ENDPOINT"
)

var mappingCreateInternalVnicDetailsServiceVnicTypeEnum = map[string]CreateInternalVnicDetailsServiceVnicTypeEnum{
	"PRIVATE_ENDPOINT":            CreateInternalVnicDetailsServiceVnicTypePrivateEndpoint,
	"REVERSE_CONNECTION_ENDPOINT": CreateInternalVnicDetailsServiceVnicTypeReverseConnectionEndpoint,
	"REAL_VIRTUAL_ROUTER":         CreateInternalVnicDetailsServiceVnicTypeRealVirtualRouter,
	"PRIVATE_DNS_ENDPOINT":        CreateInternalVnicDetailsServiceVnicTypePrivateDnsEndpoint,
}

var mappingCreateInternalVnicDetailsServiceVnicTypeEnumLowerCase = map[string]CreateInternalVnicDetailsServiceVnicTypeEnum{
	"private_endpoint":            CreateInternalVnicDetailsServiceVnicTypePrivateEndpoint,
	"reverse_connection_endpoint": CreateInternalVnicDetailsServiceVnicTypeReverseConnectionEndpoint,
	"real_virtual_router":         CreateInternalVnicDetailsServiceVnicTypeRealVirtualRouter,
	"private_dns_endpoint":        CreateInternalVnicDetailsServiceVnicTypePrivateDnsEndpoint,
}

// GetCreateInternalVnicDetailsServiceVnicTypeEnumValues Enumerates the set of values for CreateInternalVnicDetailsServiceVnicTypeEnum
func GetCreateInternalVnicDetailsServiceVnicTypeEnumValues() []CreateInternalVnicDetailsServiceVnicTypeEnum {
	values := make([]CreateInternalVnicDetailsServiceVnicTypeEnum, 0)
	for _, v := range mappingCreateInternalVnicDetailsServiceVnicTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalVnicDetailsServiceVnicTypeEnumStringValues Enumerates the set of values in String for CreateInternalVnicDetailsServiceVnicTypeEnum
func GetCreateInternalVnicDetailsServiceVnicTypeEnumStringValues() []string {
	return []string{
		"PRIVATE_ENDPOINT",
		"REVERSE_CONNECTION_ENDPOINT",
		"REAL_VIRTUAL_ROUTER",
		"PRIVATE_DNS_ENDPOINT",
	}
}

// GetMappingCreateInternalVnicDetailsServiceVnicTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalVnicDetailsServiceVnicTypeEnum(val string) (CreateInternalVnicDetailsServiceVnicTypeEnum, bool) {
	enum, ok := mappingCreateInternalVnicDetailsServiceVnicTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreateInternalVnicDetailsVnicShapeEnum Enum with underlying type: string
type CreateInternalVnicDetailsVnicShapeEnum string

// Set of constants representing the allowable values for CreateInternalVnicDetailsVnicShapeEnum
const (
	CreateInternalVnicDetailsVnicShapeDynamic                         CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC"
	CreateInternalVnicDetailsVnicShapeFixed0040                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040"
	CreateInternalVnicDetailsVnicShapeFixed0060                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060"
	CreateInternalVnicDetailsVnicShapeFixed0060Psm                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060_PSM"
	CreateInternalVnicDetailsVnicShapeFixed0100                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100"
	CreateInternalVnicDetailsVnicShapeFixed0120                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120"
	CreateInternalVnicDetailsVnicShapeFixed01202x                     CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120_2X"
	CreateInternalVnicDetailsVnicShapeFixed0200                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200"
	CreateInternalVnicDetailsVnicShapeFixed0240                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0240"
	CreateInternalVnicDetailsVnicShapeFixed0480                       CreateInternalVnicDetailsVnicShapeEnum = "FIXED0480"
	CreateInternalVnicDetailsVnicShapeEntirehost                      CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST"
	CreateInternalVnicDetailsVnicShapeDynamic25g                      CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_25G"
	CreateInternalVnicDetailsVnicShapeFixed004025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_25G"
	CreateInternalVnicDetailsVnicShapeFixed010025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_25G"
	CreateInternalVnicDetailsVnicShapeFixed020025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_25G"
	CreateInternalVnicDetailsVnicShapeFixed040025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_25G"
	CreateInternalVnicDetailsVnicShapeFixed080025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_25G"
	CreateInternalVnicDetailsVnicShapeFixed160025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_25G"
	CreateInternalVnicDetailsVnicShapeFixed240025g                    CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_25G"
	CreateInternalVnicDetailsVnicShapeEntirehost25g                   CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_25G"
	CreateInternalVnicDetailsVnicShapeDynamicE125g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0040E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0070E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0070_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0140E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0140_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0280E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0280_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0560E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0560_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed1120E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1120_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed1680E125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1680_E1_25G"
	CreateInternalVnicDetailsVnicShapeEntirehostE125g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_E1_25G"
	CreateInternalVnicDetailsVnicShapeDynamicB125g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0040B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0060B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0120B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0240B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0240_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0480B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0480_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0960B125g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0960_B1_25G"
	CreateInternalVnicDetailsVnicShapeEntirehostB125g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_B1_25G"
	CreateInternalVnicDetailsVnicShapeMicroVmFixed0048E125g           CreateInternalVnicDetailsVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	CreateInternalVnicDetailsVnicShapeMicroLbFixed0001E125g           CreateInternalVnicDetailsVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	CreateInternalVnicDetailsVnicShapeVnicaasFixed0200                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FIXED0200"
	CreateInternalVnicDetailsVnicShapeVnicaasFixed0400                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FIXED0400"
	CreateInternalVnicDetailsVnicShapeVnicaasFixed0700                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FIXED0700"
	CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved10g           CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_10G"
	CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved25g           CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_NLB_APPROVED_25G"
	CreateInternalVnicDetailsVnicShapeVnicaasTelesis25g               CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_TELESIS_25G"
	CreateInternalVnicDetailsVnicShapeVnicaasTelesis10g               CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_TELESIS_10G"
	CreateInternalVnicDetailsVnicShapeVnicaasAmbassadorFixed0100      CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_AMBASSADOR_FIXED0100"
	CreateInternalVnicDetailsVnicShapeVnicaasTelesisGamma             CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_TELESIS_GAMMA"
	CreateInternalVnicDetailsVnicShapeVnicaasPrivatedns               CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_PRIVATEDNS"
	CreateInternalVnicDetailsVnicShapeVnicaasFwaas                    CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FWAAS"
	CreateInternalVnicDetailsVnicShapeVnicaasLbaasFree                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_LBAAS_FREE"
	CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g512k              CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_512K"
	CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g1m                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_1M"
	CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g2m                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_2M"
	CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g3m                CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_LBAAS_8G_3M"
	CreateInternalVnicDetailsVnicShapeDynamicE350g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0040E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0100E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0200E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0300E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0400E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0500E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0600E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0700E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0800E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0900E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1000E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1100E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1200E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1300E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1400E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1500E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1600E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1700E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1800E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1900E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2000E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2100E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2200E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2300E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2400E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2500E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2600E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2700E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2800E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2900E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3000E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3100E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3200E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3300E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3400E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3500E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3600E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3700E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3800E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3900E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed4000E350g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED4000_E3_50G"
	CreateInternalVnicDetailsVnicShapeEntirehostE350g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_E3_50G"
	CreateInternalVnicDetailsVnicShapeDynamicE450g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0040E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0100E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0200E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0300E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0300_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0400E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0500E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0500_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0600E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0600_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0700E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0700_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0800E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed0900E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0900_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1000E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1000_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1100E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1100_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1200E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1200_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1300E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1300_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1400E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1400_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1500E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1500_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1600E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1700E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1700_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1800E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1800_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed1900E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1900_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2000E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2000_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2100E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2100_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2200E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2200_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2300E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2300_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2400E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2500E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2500_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2600E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2600_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2700E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2700_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2800E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2800_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed2900E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2900_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3000E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3000_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3100E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3100_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3200E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3200_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3300E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3300_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3400E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3400_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3500E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3500_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3600E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3600_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3700E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3700_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3800E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3800_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed3900E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3900_E4_50G"
	CreateInternalVnicDetailsVnicShapeFixed4000E450g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED4000_E4_50G"
	CreateInternalVnicDetailsVnicShapeEntirehostE450g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_E4_50G"
	CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E350g           CreateInternalVnicDetailsVnicShapeEnum = "MICRO_VM_FIXED0050_E3_50G"
	CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E450g           CreateInternalVnicDetailsVnicShapeEnum = "MICRO_VM_FIXED0050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E350g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E3_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0025_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0075_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0125_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0150_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0175_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0225_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0250_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0275_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0325_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0350_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0375_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0425_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0475_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0525_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0550_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0575_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0625_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0650_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0675_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0725_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0750_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0775_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0825_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0850_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0875_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0925_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0950_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0975_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1025_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1075_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1125_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1150_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1175_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1225_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1250_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1275_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1325_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1375_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1425_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1450_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1475_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1525_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1550_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1575_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1625_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1650_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1725_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1750_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1850_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1875_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1925_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1950_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2025_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2125_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2150_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2175_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2275_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2325_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2350_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2375_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2450_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2475_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2550_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2625_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2650_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2750_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2775_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2850_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2875_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2925_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2950_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2975_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3025_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3075_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3125_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3225_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3250_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3325_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3375_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3450_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3525_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3575_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3625_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3675_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3750_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3825_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3850_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3875_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3975_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4025_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4125_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4225_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4250_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4275_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4350_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4375_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4425_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4550_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4575_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4625_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4650_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4675_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4725_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4750_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4875_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E450g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_E4_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0020A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0020_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0040A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0040_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0060A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0060_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0080A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0080_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0100_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0120A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0120_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0140A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0140_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0160A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0160_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0200_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0220A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0220_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0240A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0240_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0260A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0260_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0280A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0280_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0300_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0320A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0320_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0340A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0340_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0380A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0380_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0400_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0420A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0420_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0440A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0440_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0460A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0460_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0480A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0480_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0500_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0520A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0520_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0560A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0560_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0580A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0580_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0600_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0620A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0620_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0640A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0640_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0660A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0660_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0680A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0680_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0700_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0740A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0740_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0760A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0760_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0780A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0780_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0800_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0820A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0820_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0840A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0840_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0860A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0860_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0880A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0880_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0920A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0920_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0940A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0940_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0960A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0960_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0980A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0980_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1000_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1020A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1020_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1040A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1040_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1060A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1060_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1100_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1120A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1120_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1140A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1140_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1160A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1160_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1180A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1180_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1200_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1220A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1220_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1240A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1240_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1280A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1280_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1300_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1320A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1320_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1340A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1340_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1360A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1360_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1380A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1380_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1400_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1420A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1420_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1460A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1460_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1480A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1480_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1500_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1520A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1520_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1540A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1540_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1560A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1560_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1580A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1580_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1600_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1640A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1640_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1660A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1660_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1680A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1680_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1700_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1720A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1720_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1740A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1740_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1760A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1760_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1780A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1780_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1820A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1820_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1840A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1840_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1860A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1860_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1880A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1880_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1900_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1920A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1920_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1940A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1940_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1960A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1960_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2000_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2020A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2020_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2040A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2040_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2060A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2060_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2080A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2080_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2100_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2120A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2120_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2140A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2140_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2180A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2180_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2200_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2220A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2220_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2240A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2240_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2260A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2260_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2280A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2280_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2300_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2320A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2320_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2360A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2360_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2380A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2380_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2400_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2420A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2420_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2440A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2440_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2460A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2460_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2480A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2480_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2500_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2540A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2540_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2560A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2560_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2580A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2580_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2600_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2620A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2620_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2640A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2640_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2660A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2660_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2680A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2680_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2720A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2720_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2740A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2740_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2760A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2760_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2780A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2780_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2800_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2820A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2820_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2840A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2840_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2860A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2860_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2900_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2920A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2920_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2940A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2940_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2960A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2960_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2980A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2980_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3000_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3020A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3020_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3040A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3040_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3080A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3080_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3100_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3120A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3120_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3140A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3140_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3160A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3160_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3180A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3180_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3200_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3220A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3220_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3260A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3260_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3280A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3280_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3300_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3320A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3320_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3340A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3340_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3360A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3360_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3380A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3380_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3400_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3440A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3440_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3460A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3460_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3480A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3480_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3500_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3520A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3520_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3540A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3540_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3560A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3560_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3580A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3580_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3620A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3620_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3640A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3640_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3660A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3660_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3680A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3680_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3700_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3720A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3720_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3740A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3740_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3760A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3760_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3800_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3820A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3820_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3840A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3840_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3860A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3860_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3880A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3880_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3900_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3920A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3920_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3940A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3940_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3980A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3980_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4000_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4020A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4020_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4040A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4040_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4060A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4060_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4080A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4080_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4100_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4120A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4120_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4160A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4160_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4180A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4180_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4200_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4220A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4220_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4240A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4240_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4260A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4260_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4280A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4280_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4300_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4340A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4340_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4360A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4360_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4380A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4380_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4400_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4420A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4420_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4440A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4440_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4460A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4460_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4480A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4480_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4520A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4520_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4540A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4540_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4560A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4560_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4580A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4580_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4600_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4620A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4620_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4640A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4640_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4660A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4660_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4700_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4720A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4720_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4740A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4740_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4760A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4760_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4780A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4780_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4800_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4820A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4820_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4840A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4840_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4880A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4880_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4900_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4920A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4920_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4940A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4940_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4960A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4960_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4980A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4980_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000A150g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED5000_A1_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0090X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0090_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0180_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0270X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0270_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0360_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0450_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0540_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0630X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0630_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0720_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0810X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0810_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0990X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED0990_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1080_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1170X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1170_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1260_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1350_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1440_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1530X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1530_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1620_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1710X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1710_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1890X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1890_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED1980_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2070X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2070_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2160_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2340_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2430X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2430_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2520_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2610X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2610_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2790X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2790_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2880_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2970X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED2970_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3060_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3150_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3240_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3330X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3330_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3420_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3510X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3510_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3690X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3690_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3780_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3870X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3870_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED3960_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4140_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4230X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4230_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4320_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4410X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4410_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4590X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4590_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4680_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4770X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4770_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4860_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950X950g         CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_VM_FIXED4950_X9_50G"
	CreateInternalVnicDetailsVnicShapeDynamicA150g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0040A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0100A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0200A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0300A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0300_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0400A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0500A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0500_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0600A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0600_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0700A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0700_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0800A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed0900A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0900_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1000A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1000_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1100A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1100_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1200A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1200_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1300A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1300_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1400A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1400_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1500A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1500_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1600A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1700A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1700_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1800A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1800_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed1900A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1900_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2000A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2000_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2100A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2100_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2200A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2200_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2300A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2300_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2400A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2500A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2500_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2600A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2600_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2700A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2700_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2800A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2800_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed2900A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2900_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3000A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3000_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3100A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3100_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3200A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3200_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3300A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3300_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3400A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3400_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3500A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3500_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3600A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3600_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3700A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3700_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3800A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3800_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed3900A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3900_A1_50G"
	CreateInternalVnicDetailsVnicShapeFixed4000A150g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED4000_A1_50G"
	CreateInternalVnicDetailsVnicShapeEntirehostA150g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_A1_50G"
	CreateInternalVnicDetailsVnicShapeDynamicX950g                    CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed0040X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed0400X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed0800X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed1200X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1200_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed1600X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed2000X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2000_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed2400X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed2800X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED2800_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed3200X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3200_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed3600X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED3600_X9_50G"
	CreateInternalVnicDetailsVnicShapeFixed4000X950g                  CreateInternalVnicDetailsVnicShapeEnum = "FIXED4000_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0100X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0100_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0200X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0200_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0300X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0300_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0400X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0400_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0500X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0500_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0600X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0600_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0700X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0700_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0800X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0800_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed0900X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED0900_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1000X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1000_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1100X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1100_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1200X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1200_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1300X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1300_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1400X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1400_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1500X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1500_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1600X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1600_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1700X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1700_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1800X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1800_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed1900X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED1900_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2000X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2000_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2100X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2100_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2200X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2200_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2300X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2300_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2400X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2400_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2500X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2500_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2600X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2600_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2700X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2700_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2800X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2800_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed2900X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED2900_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3000X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3000_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3100X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3100_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3200X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3200_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3300X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3300_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3400X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3400_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3500X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3500_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3600X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3600_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3700X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3700_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3800X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3800_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed3900X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED3900_X9_50G"
	CreateInternalVnicDetailsVnicShapeStandardVmFixed4000X950g        CreateInternalVnicDetailsVnicShapeEnum = "STANDARD_VM_FIXED4000_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0025X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0025_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0050X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0075X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0075_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0100X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0100_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0125X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0125_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0150X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0150_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0175X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0175_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0200X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0200_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0225X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0225_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0250X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0275X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0275_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0300X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0300_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0325X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0325_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0350X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0350_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0375X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0375_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0400X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0400_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0425X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0425_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0450X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0450_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0475X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0475_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0500X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0525X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0525_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0550X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0550_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0575X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0575_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0600X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0625X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0625_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0650X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0650_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0675X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0675_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0700X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0725X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0725_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0750X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0750_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0775X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0775_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0800X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0825X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0825_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0850X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0850_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0875X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0875_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0900X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0925X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0925_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0950X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0950_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0975X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED0975_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1000X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1000_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1025X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1025_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1050X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1075X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1075_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1100X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1100_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1125X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1125_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1150X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1150_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1175X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1175_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1200X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1200_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1225X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1225_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1250X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1275X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1275_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1300X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1300_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1325X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1325_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1350X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1350_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1375X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1375_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1400X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1400_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1425X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1425_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1450X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1450_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1475X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1475_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1500X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1525X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1525_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1550X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1550_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1575X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1575_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1600X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1625X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1625_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1650X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1650_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1700X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1725X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1725_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1750X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1750_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1800X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1850X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1850_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1875X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1875_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1900X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1925X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1925_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1950X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED1950_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2000X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2000_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2025X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2025_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2050X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2100X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2100_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2125X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2125_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2150X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2150_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2175X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2175_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2200X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2200_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2250X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2275X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2275_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2300X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2300_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2325X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2325_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2350X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2350_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2375X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2375_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2400X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2400_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2450X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2450_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2475X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2475_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2500X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2550X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2550_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2600X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2625X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2625_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2650X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2650_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2700X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2750X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2750_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2775X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2775_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2800X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2850X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2850_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2875X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2875_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2900X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2925X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2925_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2950X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2950_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2975X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED2975_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3000X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3000_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3025X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3025_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3050X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3075X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3075_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3100X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3100_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3125X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3125_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3150X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3150_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3200X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3200_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3225X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3225_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3250X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3300X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3300_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3325X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3325_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3375X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3375_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3400X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3400_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3450X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3450_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3500X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3525X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3525_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3575X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3575_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3600X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3625X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3625_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3675X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3675_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3700X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3750X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3750_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3800X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3825X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3825_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3850X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3850_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3875X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3875_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3900X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3975X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED3975_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4000X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4000_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4025X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4025_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4050X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4050_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4100X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4100_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4125X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4125_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4200X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4200_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4225X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4225_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4250X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4250_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4275X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4275_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4300X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4300_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4350X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4350_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4375X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4375_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4400X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4400_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4425X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4425_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4500X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4500_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4550X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4550_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4575X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4575_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4600X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4600_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4625X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4625_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4650X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4650_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4675X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4675_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4700X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4700_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4725X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4725_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4750X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4750_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4800X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4800_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4875X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4875_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4900X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4900_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4950X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED4950_X9_50G"
	CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed5000X950g CreateInternalVnicDetailsVnicShapeEnum = "SUBCORE_STANDARD_VM_FIXED5000_X9_50G"
	CreateInternalVnicDetailsVnicShapeEntirehostX950g                 CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_X9_50G"
)

var mappingCreateInternalVnicDetailsVnicShapeEnum = map[string]CreateInternalVnicDetailsVnicShapeEnum{
	"DYNAMIC":                              CreateInternalVnicDetailsVnicShapeDynamic,
	"FIXED0040":                            CreateInternalVnicDetailsVnicShapeFixed0040,
	"FIXED0060":                            CreateInternalVnicDetailsVnicShapeFixed0060,
	"FIXED0060_PSM":                        CreateInternalVnicDetailsVnicShapeFixed0060Psm,
	"FIXED0100":                            CreateInternalVnicDetailsVnicShapeFixed0100,
	"FIXED0120":                            CreateInternalVnicDetailsVnicShapeFixed0120,
	"FIXED0120_2X":                         CreateInternalVnicDetailsVnicShapeFixed01202x,
	"FIXED0200":                            CreateInternalVnicDetailsVnicShapeFixed0200,
	"FIXED0240":                            CreateInternalVnicDetailsVnicShapeFixed0240,
	"FIXED0480":                            CreateInternalVnicDetailsVnicShapeFixed0480,
	"ENTIREHOST":                           CreateInternalVnicDetailsVnicShapeEntirehost,
	"DYNAMIC_25G":                          CreateInternalVnicDetailsVnicShapeDynamic25g,
	"FIXED0040_25G":                        CreateInternalVnicDetailsVnicShapeFixed004025g,
	"FIXED0100_25G":                        CreateInternalVnicDetailsVnicShapeFixed010025g,
	"FIXED0200_25G":                        CreateInternalVnicDetailsVnicShapeFixed020025g,
	"FIXED0400_25G":                        CreateInternalVnicDetailsVnicShapeFixed040025g,
	"FIXED0800_25G":                        CreateInternalVnicDetailsVnicShapeFixed080025g,
	"FIXED1600_25G":                        CreateInternalVnicDetailsVnicShapeFixed160025g,
	"FIXED2400_25G":                        CreateInternalVnicDetailsVnicShapeFixed240025g,
	"ENTIREHOST_25G":                       CreateInternalVnicDetailsVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":                       CreateInternalVnicDetailsVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":                     CreateInternalVnicDetailsVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":                    CreateInternalVnicDetailsVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":                       CreateInternalVnicDetailsVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":                     CreateInternalVnicDetailsVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":                    CreateInternalVnicDetailsVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G":            CreateInternalVnicDetailsVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0200":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0400,
	"VNICAAS_FIXED0700":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0700,
	"VNICAAS_NLB_APPROVED_10G":             CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved10g,
	"VNICAAS_NLB_APPROVED_25G":             CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved25g,
	"VNICAAS_TELESIS_25G":                  CreateInternalVnicDetailsVnicShapeVnicaasTelesis25g,
	"VNICAAS_TELESIS_10G":                  CreateInternalVnicDetailsVnicShapeVnicaasTelesis10g,
	"VNICAAS_AMBASSADOR_FIXED0100":         CreateInternalVnicDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"VNICAAS_TELESIS_GAMMA":                CreateInternalVnicDetailsVnicShapeVnicaasTelesisGamma,
	"VNICAAS_PRIVATEDNS":                   CreateInternalVnicDetailsVnicShapeVnicaasPrivatedns,
	"VNICAAS_FWAAS":                        CreateInternalVnicDetailsVnicShapeVnicaasFwaas,
	"VNICAAS_LBAAS_FREE":                   CreateInternalVnicDetailsVnicShapeVnicaasLbaasFree,
	"VNICAAS_LBAAS_8G_512K":                CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g512k,
	"VNICAAS_LBAAS_8G_1M":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g1m,
	"VNICAAS_LBAAS_8G_2M":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g2m,
	"VNICAAS_LBAAS_8G_3M":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g3m,
	"DYNAMIC_E3_50G":                       CreateInternalVnicDetailsVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":                     CreateInternalVnicDetailsVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":                    CreateInternalVnicDetailsVnicShapeEntirehostE350g,
	"DYNAMIC_E4_50G":                       CreateInternalVnicDetailsVnicShapeDynamicE450g,
	"FIXED0040_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0040E450g,
	"FIXED0100_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0100E450g,
	"FIXED0200_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0200E450g,
	"FIXED0300_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0300E450g,
	"FIXED0400_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0400E450g,
	"FIXED0500_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0500E450g,
	"FIXED0600_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0600E450g,
	"FIXED0700_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0700E450g,
	"FIXED0800_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0800E450g,
	"FIXED0900_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed0900E450g,
	"FIXED1000_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1000E450g,
	"FIXED1100_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1100E450g,
	"FIXED1200_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1200E450g,
	"FIXED1300_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1300E450g,
	"FIXED1400_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1400E450g,
	"FIXED1500_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1500E450g,
	"FIXED1600_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1600E450g,
	"FIXED1700_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1700E450g,
	"FIXED1800_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1800E450g,
	"FIXED1900_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed1900E450g,
	"FIXED2000_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2000E450g,
	"FIXED2100_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2100E450g,
	"FIXED2200_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2200E450g,
	"FIXED2300_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2300E450g,
	"FIXED2400_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2400E450g,
	"FIXED2500_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2500E450g,
	"FIXED2600_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2600E450g,
	"FIXED2700_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2700E450g,
	"FIXED2800_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2800E450g,
	"FIXED2900_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed2900E450g,
	"FIXED3000_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3000E450g,
	"FIXED3100_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3100E450g,
	"FIXED3200_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3200E450g,
	"FIXED3300_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3300E450g,
	"FIXED3400_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3400E450g,
	"FIXED3500_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3500E450g,
	"FIXED3600_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3600E450g,
	"FIXED3700_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3700E450g,
	"FIXED3800_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3800E450g,
	"FIXED3900_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed3900E450g,
	"FIXED4000_E4_50G":                     CreateInternalVnicDetailsVnicShapeFixed4000E450g,
	"ENTIREHOST_E4_50G":                    CreateInternalVnicDetailsVnicShapeEntirehostE450g,
	"MICRO_VM_FIXED0050_E3_50G":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E350g,
	"MICRO_VM_FIXED0050_E4_50G":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E450g,
	"SUBCORE_VM_FIXED0025_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E350g,
	"SUBCORE_VM_FIXED0050_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E350g,
	"SUBCORE_VM_FIXED0075_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E350g,
	"SUBCORE_VM_FIXED0100_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E350g,
	"SUBCORE_VM_FIXED0125_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E350g,
	"SUBCORE_VM_FIXED0150_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E350g,
	"SUBCORE_VM_FIXED0175_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E350g,
	"SUBCORE_VM_FIXED0200_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E350g,
	"SUBCORE_VM_FIXED0225_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E350g,
	"SUBCORE_VM_FIXED0250_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E350g,
	"SUBCORE_VM_FIXED0275_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E350g,
	"SUBCORE_VM_FIXED0300_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E350g,
	"SUBCORE_VM_FIXED0325_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E350g,
	"SUBCORE_VM_FIXED0350_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E350g,
	"SUBCORE_VM_FIXED0375_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E350g,
	"SUBCORE_VM_FIXED0400_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E350g,
	"SUBCORE_VM_FIXED0425_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E350g,
	"SUBCORE_VM_FIXED0450_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E350g,
	"SUBCORE_VM_FIXED0475_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E350g,
	"SUBCORE_VM_FIXED0500_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E350g,
	"SUBCORE_VM_FIXED0525_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E350g,
	"SUBCORE_VM_FIXED0550_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E350g,
	"SUBCORE_VM_FIXED0575_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E350g,
	"SUBCORE_VM_FIXED0600_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E350g,
	"SUBCORE_VM_FIXED0625_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E350g,
	"SUBCORE_VM_FIXED0650_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E350g,
	"SUBCORE_VM_FIXED0675_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E350g,
	"SUBCORE_VM_FIXED0700_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E350g,
	"SUBCORE_VM_FIXED0725_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E350g,
	"SUBCORE_VM_FIXED0750_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E350g,
	"SUBCORE_VM_FIXED0775_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E350g,
	"SUBCORE_VM_FIXED0800_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E350g,
	"SUBCORE_VM_FIXED0825_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E350g,
	"SUBCORE_VM_FIXED0850_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E350g,
	"SUBCORE_VM_FIXED0875_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E350g,
	"SUBCORE_VM_FIXED0900_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E350g,
	"SUBCORE_VM_FIXED0925_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E350g,
	"SUBCORE_VM_FIXED0950_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E350g,
	"SUBCORE_VM_FIXED0975_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E350g,
	"SUBCORE_VM_FIXED1000_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E350g,
	"SUBCORE_VM_FIXED1025_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E350g,
	"SUBCORE_VM_FIXED1050_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E350g,
	"SUBCORE_VM_FIXED1075_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E350g,
	"SUBCORE_VM_FIXED1100_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E350g,
	"SUBCORE_VM_FIXED1125_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E350g,
	"SUBCORE_VM_FIXED1150_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E350g,
	"SUBCORE_VM_FIXED1175_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E350g,
	"SUBCORE_VM_FIXED1200_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E350g,
	"SUBCORE_VM_FIXED1225_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E350g,
	"SUBCORE_VM_FIXED1250_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E350g,
	"SUBCORE_VM_FIXED1275_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E350g,
	"SUBCORE_VM_FIXED1300_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E350g,
	"SUBCORE_VM_FIXED1325_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E350g,
	"SUBCORE_VM_FIXED1350_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E350g,
	"SUBCORE_VM_FIXED1375_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E350g,
	"SUBCORE_VM_FIXED1400_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E350g,
	"SUBCORE_VM_FIXED1425_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E350g,
	"SUBCORE_VM_FIXED1450_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E350g,
	"SUBCORE_VM_FIXED1475_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E350g,
	"SUBCORE_VM_FIXED1500_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E350g,
	"SUBCORE_VM_FIXED1525_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E350g,
	"SUBCORE_VM_FIXED1550_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E350g,
	"SUBCORE_VM_FIXED1575_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E350g,
	"SUBCORE_VM_FIXED1600_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E350g,
	"SUBCORE_VM_FIXED1625_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E350g,
	"SUBCORE_VM_FIXED1650_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E350g,
	"SUBCORE_VM_FIXED1700_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E350g,
	"SUBCORE_VM_FIXED1725_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E350g,
	"SUBCORE_VM_FIXED1750_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E350g,
	"SUBCORE_VM_FIXED1800_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E350g,
	"SUBCORE_VM_FIXED1850_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E350g,
	"SUBCORE_VM_FIXED1875_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E350g,
	"SUBCORE_VM_FIXED1900_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E350g,
	"SUBCORE_VM_FIXED1925_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E350g,
	"SUBCORE_VM_FIXED1950_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E350g,
	"SUBCORE_VM_FIXED2000_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E350g,
	"SUBCORE_VM_FIXED2025_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E350g,
	"SUBCORE_VM_FIXED2050_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E350g,
	"SUBCORE_VM_FIXED2100_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E350g,
	"SUBCORE_VM_FIXED2125_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E350g,
	"SUBCORE_VM_FIXED2150_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E350g,
	"SUBCORE_VM_FIXED2175_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E350g,
	"SUBCORE_VM_FIXED2200_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E350g,
	"SUBCORE_VM_FIXED2250_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E350g,
	"SUBCORE_VM_FIXED2275_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E350g,
	"SUBCORE_VM_FIXED2300_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E350g,
	"SUBCORE_VM_FIXED2325_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E350g,
	"SUBCORE_VM_FIXED2350_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E350g,
	"SUBCORE_VM_FIXED2375_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E350g,
	"SUBCORE_VM_FIXED2400_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E350g,
	"SUBCORE_VM_FIXED2450_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E350g,
	"SUBCORE_VM_FIXED2475_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E350g,
	"SUBCORE_VM_FIXED2500_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E350g,
	"SUBCORE_VM_FIXED2550_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E350g,
	"SUBCORE_VM_FIXED2600_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E350g,
	"SUBCORE_VM_FIXED2625_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E350g,
	"SUBCORE_VM_FIXED2650_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E350g,
	"SUBCORE_VM_FIXED2700_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E350g,
	"SUBCORE_VM_FIXED2750_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E350g,
	"SUBCORE_VM_FIXED2775_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E350g,
	"SUBCORE_VM_FIXED2800_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E350g,
	"SUBCORE_VM_FIXED2850_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E350g,
	"SUBCORE_VM_FIXED2875_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E350g,
	"SUBCORE_VM_FIXED2900_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E350g,
	"SUBCORE_VM_FIXED2925_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E350g,
	"SUBCORE_VM_FIXED2950_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E350g,
	"SUBCORE_VM_FIXED2975_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E350g,
	"SUBCORE_VM_FIXED3000_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E350g,
	"SUBCORE_VM_FIXED3025_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E350g,
	"SUBCORE_VM_FIXED3050_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E350g,
	"SUBCORE_VM_FIXED3075_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E350g,
	"SUBCORE_VM_FIXED3100_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E350g,
	"SUBCORE_VM_FIXED3125_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E350g,
	"SUBCORE_VM_FIXED3150_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E350g,
	"SUBCORE_VM_FIXED3200_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E350g,
	"SUBCORE_VM_FIXED3225_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E350g,
	"SUBCORE_VM_FIXED3250_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E350g,
	"SUBCORE_VM_FIXED3300_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E350g,
	"SUBCORE_VM_FIXED3325_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E350g,
	"SUBCORE_VM_FIXED3375_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E350g,
	"SUBCORE_VM_FIXED3400_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E350g,
	"SUBCORE_VM_FIXED3450_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E350g,
	"SUBCORE_VM_FIXED3500_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E350g,
	"SUBCORE_VM_FIXED3525_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E350g,
	"SUBCORE_VM_FIXED3575_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E350g,
	"SUBCORE_VM_FIXED3600_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E350g,
	"SUBCORE_VM_FIXED3625_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E350g,
	"SUBCORE_VM_FIXED3675_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E350g,
	"SUBCORE_VM_FIXED3700_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E350g,
	"SUBCORE_VM_FIXED3750_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E350g,
	"SUBCORE_VM_FIXED3800_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E350g,
	"SUBCORE_VM_FIXED3825_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E350g,
	"SUBCORE_VM_FIXED3850_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E350g,
	"SUBCORE_VM_FIXED3875_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E350g,
	"SUBCORE_VM_FIXED3900_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E350g,
	"SUBCORE_VM_FIXED3975_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E350g,
	"SUBCORE_VM_FIXED4000_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E350g,
	"SUBCORE_VM_FIXED4025_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E350g,
	"SUBCORE_VM_FIXED4050_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E350g,
	"SUBCORE_VM_FIXED4100_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E350g,
	"SUBCORE_VM_FIXED4125_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E350g,
	"SUBCORE_VM_FIXED4200_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E350g,
	"SUBCORE_VM_FIXED4225_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E350g,
	"SUBCORE_VM_FIXED4250_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E350g,
	"SUBCORE_VM_FIXED4275_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E350g,
	"SUBCORE_VM_FIXED4300_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E350g,
	"SUBCORE_VM_FIXED4350_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E350g,
	"SUBCORE_VM_FIXED4375_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E350g,
	"SUBCORE_VM_FIXED4400_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E350g,
	"SUBCORE_VM_FIXED4425_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E350g,
	"SUBCORE_VM_FIXED4500_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E350g,
	"SUBCORE_VM_FIXED4550_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E350g,
	"SUBCORE_VM_FIXED4575_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E350g,
	"SUBCORE_VM_FIXED4600_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E350g,
	"SUBCORE_VM_FIXED4625_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E350g,
	"SUBCORE_VM_FIXED4650_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E350g,
	"SUBCORE_VM_FIXED4675_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E350g,
	"SUBCORE_VM_FIXED4700_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E350g,
	"SUBCORE_VM_FIXED4725_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E350g,
	"SUBCORE_VM_FIXED4750_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E350g,
	"SUBCORE_VM_FIXED4800_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E350g,
	"SUBCORE_VM_FIXED4875_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E350g,
	"SUBCORE_VM_FIXED4900_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E350g,
	"SUBCORE_VM_FIXED4950_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E350g,
	"SUBCORE_VM_FIXED5000_E3_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E350g,
	"SUBCORE_VM_FIXED0025_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E450g,
	"SUBCORE_VM_FIXED0050_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E450g,
	"SUBCORE_VM_FIXED0075_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E450g,
	"SUBCORE_VM_FIXED0100_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E450g,
	"SUBCORE_VM_FIXED0125_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E450g,
	"SUBCORE_VM_FIXED0150_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E450g,
	"SUBCORE_VM_FIXED0175_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E450g,
	"SUBCORE_VM_FIXED0200_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E450g,
	"SUBCORE_VM_FIXED0225_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E450g,
	"SUBCORE_VM_FIXED0250_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E450g,
	"SUBCORE_VM_FIXED0275_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E450g,
	"SUBCORE_VM_FIXED0300_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E450g,
	"SUBCORE_VM_FIXED0325_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E450g,
	"SUBCORE_VM_FIXED0350_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E450g,
	"SUBCORE_VM_FIXED0375_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E450g,
	"SUBCORE_VM_FIXED0400_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E450g,
	"SUBCORE_VM_FIXED0425_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E450g,
	"SUBCORE_VM_FIXED0450_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E450g,
	"SUBCORE_VM_FIXED0475_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E450g,
	"SUBCORE_VM_FIXED0500_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E450g,
	"SUBCORE_VM_FIXED0525_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E450g,
	"SUBCORE_VM_FIXED0550_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E450g,
	"SUBCORE_VM_FIXED0575_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E450g,
	"SUBCORE_VM_FIXED0600_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E450g,
	"SUBCORE_VM_FIXED0625_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E450g,
	"SUBCORE_VM_FIXED0650_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E450g,
	"SUBCORE_VM_FIXED0675_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E450g,
	"SUBCORE_VM_FIXED0700_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E450g,
	"SUBCORE_VM_FIXED0725_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E450g,
	"SUBCORE_VM_FIXED0750_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E450g,
	"SUBCORE_VM_FIXED0775_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E450g,
	"SUBCORE_VM_FIXED0800_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E450g,
	"SUBCORE_VM_FIXED0825_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E450g,
	"SUBCORE_VM_FIXED0850_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E450g,
	"SUBCORE_VM_FIXED0875_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E450g,
	"SUBCORE_VM_FIXED0900_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E450g,
	"SUBCORE_VM_FIXED0925_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E450g,
	"SUBCORE_VM_FIXED0950_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E450g,
	"SUBCORE_VM_FIXED0975_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E450g,
	"SUBCORE_VM_FIXED1000_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E450g,
	"SUBCORE_VM_FIXED1025_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E450g,
	"SUBCORE_VM_FIXED1050_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E450g,
	"SUBCORE_VM_FIXED1075_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E450g,
	"SUBCORE_VM_FIXED1100_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E450g,
	"SUBCORE_VM_FIXED1125_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E450g,
	"SUBCORE_VM_FIXED1150_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E450g,
	"SUBCORE_VM_FIXED1175_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E450g,
	"SUBCORE_VM_FIXED1200_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E450g,
	"SUBCORE_VM_FIXED1225_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E450g,
	"SUBCORE_VM_FIXED1250_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E450g,
	"SUBCORE_VM_FIXED1275_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E450g,
	"SUBCORE_VM_FIXED1300_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E450g,
	"SUBCORE_VM_FIXED1325_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E450g,
	"SUBCORE_VM_FIXED1350_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E450g,
	"SUBCORE_VM_FIXED1375_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E450g,
	"SUBCORE_VM_FIXED1400_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E450g,
	"SUBCORE_VM_FIXED1425_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E450g,
	"SUBCORE_VM_FIXED1450_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E450g,
	"SUBCORE_VM_FIXED1475_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E450g,
	"SUBCORE_VM_FIXED1500_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E450g,
	"SUBCORE_VM_FIXED1525_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E450g,
	"SUBCORE_VM_FIXED1550_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E450g,
	"SUBCORE_VM_FIXED1575_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E450g,
	"SUBCORE_VM_FIXED1600_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E450g,
	"SUBCORE_VM_FIXED1625_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E450g,
	"SUBCORE_VM_FIXED1650_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E450g,
	"SUBCORE_VM_FIXED1700_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E450g,
	"SUBCORE_VM_FIXED1725_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E450g,
	"SUBCORE_VM_FIXED1750_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E450g,
	"SUBCORE_VM_FIXED1800_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E450g,
	"SUBCORE_VM_FIXED1850_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E450g,
	"SUBCORE_VM_FIXED1875_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E450g,
	"SUBCORE_VM_FIXED1900_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E450g,
	"SUBCORE_VM_FIXED1925_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E450g,
	"SUBCORE_VM_FIXED1950_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E450g,
	"SUBCORE_VM_FIXED2000_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E450g,
	"SUBCORE_VM_FIXED2025_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E450g,
	"SUBCORE_VM_FIXED2050_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E450g,
	"SUBCORE_VM_FIXED2100_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E450g,
	"SUBCORE_VM_FIXED2125_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E450g,
	"SUBCORE_VM_FIXED2150_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E450g,
	"SUBCORE_VM_FIXED2175_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E450g,
	"SUBCORE_VM_FIXED2200_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E450g,
	"SUBCORE_VM_FIXED2250_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E450g,
	"SUBCORE_VM_FIXED2275_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E450g,
	"SUBCORE_VM_FIXED2300_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E450g,
	"SUBCORE_VM_FIXED2325_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E450g,
	"SUBCORE_VM_FIXED2350_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E450g,
	"SUBCORE_VM_FIXED2375_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E450g,
	"SUBCORE_VM_FIXED2400_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E450g,
	"SUBCORE_VM_FIXED2450_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E450g,
	"SUBCORE_VM_FIXED2475_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E450g,
	"SUBCORE_VM_FIXED2500_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E450g,
	"SUBCORE_VM_FIXED2550_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E450g,
	"SUBCORE_VM_FIXED2600_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E450g,
	"SUBCORE_VM_FIXED2625_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E450g,
	"SUBCORE_VM_FIXED2650_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E450g,
	"SUBCORE_VM_FIXED2700_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E450g,
	"SUBCORE_VM_FIXED2750_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E450g,
	"SUBCORE_VM_FIXED2775_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E450g,
	"SUBCORE_VM_FIXED2800_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E450g,
	"SUBCORE_VM_FIXED2850_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E450g,
	"SUBCORE_VM_FIXED2875_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E450g,
	"SUBCORE_VM_FIXED2900_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E450g,
	"SUBCORE_VM_FIXED2925_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E450g,
	"SUBCORE_VM_FIXED2950_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E450g,
	"SUBCORE_VM_FIXED2975_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E450g,
	"SUBCORE_VM_FIXED3000_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E450g,
	"SUBCORE_VM_FIXED3025_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E450g,
	"SUBCORE_VM_FIXED3050_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E450g,
	"SUBCORE_VM_FIXED3075_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E450g,
	"SUBCORE_VM_FIXED3100_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E450g,
	"SUBCORE_VM_FIXED3125_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E450g,
	"SUBCORE_VM_FIXED3150_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E450g,
	"SUBCORE_VM_FIXED3200_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E450g,
	"SUBCORE_VM_FIXED3225_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E450g,
	"SUBCORE_VM_FIXED3250_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E450g,
	"SUBCORE_VM_FIXED3300_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E450g,
	"SUBCORE_VM_FIXED3325_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E450g,
	"SUBCORE_VM_FIXED3375_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E450g,
	"SUBCORE_VM_FIXED3400_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E450g,
	"SUBCORE_VM_FIXED3450_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E450g,
	"SUBCORE_VM_FIXED3500_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E450g,
	"SUBCORE_VM_FIXED3525_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E450g,
	"SUBCORE_VM_FIXED3575_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E450g,
	"SUBCORE_VM_FIXED3600_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E450g,
	"SUBCORE_VM_FIXED3625_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E450g,
	"SUBCORE_VM_FIXED3675_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E450g,
	"SUBCORE_VM_FIXED3700_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E450g,
	"SUBCORE_VM_FIXED3750_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E450g,
	"SUBCORE_VM_FIXED3800_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E450g,
	"SUBCORE_VM_FIXED3825_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E450g,
	"SUBCORE_VM_FIXED3850_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E450g,
	"SUBCORE_VM_FIXED3875_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E450g,
	"SUBCORE_VM_FIXED3900_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E450g,
	"SUBCORE_VM_FIXED3975_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E450g,
	"SUBCORE_VM_FIXED4000_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E450g,
	"SUBCORE_VM_FIXED4025_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E450g,
	"SUBCORE_VM_FIXED4050_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E450g,
	"SUBCORE_VM_FIXED4100_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E450g,
	"SUBCORE_VM_FIXED4125_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E450g,
	"SUBCORE_VM_FIXED4200_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E450g,
	"SUBCORE_VM_FIXED4225_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E450g,
	"SUBCORE_VM_FIXED4250_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E450g,
	"SUBCORE_VM_FIXED4275_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E450g,
	"SUBCORE_VM_FIXED4300_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E450g,
	"SUBCORE_VM_FIXED4350_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E450g,
	"SUBCORE_VM_FIXED4375_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E450g,
	"SUBCORE_VM_FIXED4400_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E450g,
	"SUBCORE_VM_FIXED4425_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E450g,
	"SUBCORE_VM_FIXED4500_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E450g,
	"SUBCORE_VM_FIXED4550_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E450g,
	"SUBCORE_VM_FIXED4575_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E450g,
	"SUBCORE_VM_FIXED4600_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E450g,
	"SUBCORE_VM_FIXED4625_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E450g,
	"SUBCORE_VM_FIXED4650_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E450g,
	"SUBCORE_VM_FIXED4675_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E450g,
	"SUBCORE_VM_FIXED4700_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E450g,
	"SUBCORE_VM_FIXED4725_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E450g,
	"SUBCORE_VM_FIXED4750_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E450g,
	"SUBCORE_VM_FIXED4800_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E450g,
	"SUBCORE_VM_FIXED4875_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E450g,
	"SUBCORE_VM_FIXED4900_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E450g,
	"SUBCORE_VM_FIXED4950_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E450g,
	"SUBCORE_VM_FIXED5000_E4_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E450g,
	"SUBCORE_VM_FIXED0020_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0020A150g,
	"SUBCORE_VM_FIXED0040_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0040A150g,
	"SUBCORE_VM_FIXED0060_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0060A150g,
	"SUBCORE_VM_FIXED0080_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0080A150g,
	"SUBCORE_VM_FIXED0100_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100A150g,
	"SUBCORE_VM_FIXED0120_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0120A150g,
	"SUBCORE_VM_FIXED0140_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0140A150g,
	"SUBCORE_VM_FIXED0160_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0160A150g,
	"SUBCORE_VM_FIXED0180_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180A150g,
	"SUBCORE_VM_FIXED0200_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200A150g,
	"SUBCORE_VM_FIXED0220_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0220A150g,
	"SUBCORE_VM_FIXED0240_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0240A150g,
	"SUBCORE_VM_FIXED0260_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0260A150g,
	"SUBCORE_VM_FIXED0280_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0280A150g,
	"SUBCORE_VM_FIXED0300_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300A150g,
	"SUBCORE_VM_FIXED0320_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0320A150g,
	"SUBCORE_VM_FIXED0340_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0340A150g,
	"SUBCORE_VM_FIXED0360_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360A150g,
	"SUBCORE_VM_FIXED0380_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0380A150g,
	"SUBCORE_VM_FIXED0400_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400A150g,
	"SUBCORE_VM_FIXED0420_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0420A150g,
	"SUBCORE_VM_FIXED0440_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0440A150g,
	"SUBCORE_VM_FIXED0460_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0460A150g,
	"SUBCORE_VM_FIXED0480_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0480A150g,
	"SUBCORE_VM_FIXED0500_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500A150g,
	"SUBCORE_VM_FIXED0520_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0520A150g,
	"SUBCORE_VM_FIXED0540_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540A150g,
	"SUBCORE_VM_FIXED0560_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0560A150g,
	"SUBCORE_VM_FIXED0580_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0580A150g,
	"SUBCORE_VM_FIXED0600_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600A150g,
	"SUBCORE_VM_FIXED0620_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0620A150g,
	"SUBCORE_VM_FIXED0640_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0640A150g,
	"SUBCORE_VM_FIXED0660_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0660A150g,
	"SUBCORE_VM_FIXED0680_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0680A150g,
	"SUBCORE_VM_FIXED0700_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700A150g,
	"SUBCORE_VM_FIXED0720_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720A150g,
	"SUBCORE_VM_FIXED0740_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0740A150g,
	"SUBCORE_VM_FIXED0760_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0760A150g,
	"SUBCORE_VM_FIXED0780_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0780A150g,
	"SUBCORE_VM_FIXED0800_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800A150g,
	"SUBCORE_VM_FIXED0820_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0820A150g,
	"SUBCORE_VM_FIXED0840_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0840A150g,
	"SUBCORE_VM_FIXED0860_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0860A150g,
	"SUBCORE_VM_FIXED0880_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0880A150g,
	"SUBCORE_VM_FIXED0900_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900A150g,
	"SUBCORE_VM_FIXED0920_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0920A150g,
	"SUBCORE_VM_FIXED0940_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0940A150g,
	"SUBCORE_VM_FIXED0960_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0960A150g,
	"SUBCORE_VM_FIXED0980_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0980A150g,
	"SUBCORE_VM_FIXED1000_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000A150g,
	"SUBCORE_VM_FIXED1020_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1020A150g,
	"SUBCORE_VM_FIXED1040_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1040A150g,
	"SUBCORE_VM_FIXED1060_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1060A150g,
	"SUBCORE_VM_FIXED1080_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080A150g,
	"SUBCORE_VM_FIXED1100_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100A150g,
	"SUBCORE_VM_FIXED1120_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1120A150g,
	"SUBCORE_VM_FIXED1140_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1140A150g,
	"SUBCORE_VM_FIXED1160_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1160A150g,
	"SUBCORE_VM_FIXED1180_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1180A150g,
	"SUBCORE_VM_FIXED1200_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200A150g,
	"SUBCORE_VM_FIXED1220_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1220A150g,
	"SUBCORE_VM_FIXED1240_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1240A150g,
	"SUBCORE_VM_FIXED1260_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260A150g,
	"SUBCORE_VM_FIXED1280_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1280A150g,
	"SUBCORE_VM_FIXED1300_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300A150g,
	"SUBCORE_VM_FIXED1320_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1320A150g,
	"SUBCORE_VM_FIXED1340_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1340A150g,
	"SUBCORE_VM_FIXED1360_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1360A150g,
	"SUBCORE_VM_FIXED1380_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1380A150g,
	"SUBCORE_VM_FIXED1400_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400A150g,
	"SUBCORE_VM_FIXED1420_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1420A150g,
	"SUBCORE_VM_FIXED1440_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440A150g,
	"SUBCORE_VM_FIXED1460_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1460A150g,
	"SUBCORE_VM_FIXED1480_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1480A150g,
	"SUBCORE_VM_FIXED1500_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500A150g,
	"SUBCORE_VM_FIXED1520_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1520A150g,
	"SUBCORE_VM_FIXED1540_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1540A150g,
	"SUBCORE_VM_FIXED1560_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1560A150g,
	"SUBCORE_VM_FIXED1580_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1580A150g,
	"SUBCORE_VM_FIXED1600_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600A150g,
	"SUBCORE_VM_FIXED1620_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620A150g,
	"SUBCORE_VM_FIXED1640_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1640A150g,
	"SUBCORE_VM_FIXED1660_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1660A150g,
	"SUBCORE_VM_FIXED1680_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1680A150g,
	"SUBCORE_VM_FIXED1700_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700A150g,
	"SUBCORE_VM_FIXED1720_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1720A150g,
	"SUBCORE_VM_FIXED1740_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1740A150g,
	"SUBCORE_VM_FIXED1760_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1760A150g,
	"SUBCORE_VM_FIXED1780_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1780A150g,
	"SUBCORE_VM_FIXED1800_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800A150g,
	"SUBCORE_VM_FIXED1820_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1820A150g,
	"SUBCORE_VM_FIXED1840_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1840A150g,
	"SUBCORE_VM_FIXED1860_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1860A150g,
	"SUBCORE_VM_FIXED1880_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1880A150g,
	"SUBCORE_VM_FIXED1900_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900A150g,
	"SUBCORE_VM_FIXED1920_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1920A150g,
	"SUBCORE_VM_FIXED1940_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1940A150g,
	"SUBCORE_VM_FIXED1960_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1960A150g,
	"SUBCORE_VM_FIXED1980_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980A150g,
	"SUBCORE_VM_FIXED2000_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000A150g,
	"SUBCORE_VM_FIXED2020_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2020A150g,
	"SUBCORE_VM_FIXED2040_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2040A150g,
	"SUBCORE_VM_FIXED2060_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2060A150g,
	"SUBCORE_VM_FIXED2080_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2080A150g,
	"SUBCORE_VM_FIXED2100_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100A150g,
	"SUBCORE_VM_FIXED2120_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2120A150g,
	"SUBCORE_VM_FIXED2140_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2140A150g,
	"SUBCORE_VM_FIXED2160_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160A150g,
	"SUBCORE_VM_FIXED2180_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2180A150g,
	"SUBCORE_VM_FIXED2200_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200A150g,
	"SUBCORE_VM_FIXED2220_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2220A150g,
	"SUBCORE_VM_FIXED2240_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2240A150g,
	"SUBCORE_VM_FIXED2260_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2260A150g,
	"SUBCORE_VM_FIXED2280_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2280A150g,
	"SUBCORE_VM_FIXED2300_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300A150g,
	"SUBCORE_VM_FIXED2320_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2320A150g,
	"SUBCORE_VM_FIXED2340_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340A150g,
	"SUBCORE_VM_FIXED2360_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2360A150g,
	"SUBCORE_VM_FIXED2380_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2380A150g,
	"SUBCORE_VM_FIXED2400_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400A150g,
	"SUBCORE_VM_FIXED2420_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2420A150g,
	"SUBCORE_VM_FIXED2440_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2440A150g,
	"SUBCORE_VM_FIXED2460_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2460A150g,
	"SUBCORE_VM_FIXED2480_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2480A150g,
	"SUBCORE_VM_FIXED2500_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500A150g,
	"SUBCORE_VM_FIXED2520_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520A150g,
	"SUBCORE_VM_FIXED2540_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2540A150g,
	"SUBCORE_VM_FIXED2560_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2560A150g,
	"SUBCORE_VM_FIXED2580_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2580A150g,
	"SUBCORE_VM_FIXED2600_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600A150g,
	"SUBCORE_VM_FIXED2620_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2620A150g,
	"SUBCORE_VM_FIXED2640_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2640A150g,
	"SUBCORE_VM_FIXED2660_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2660A150g,
	"SUBCORE_VM_FIXED2680_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2680A150g,
	"SUBCORE_VM_FIXED2700_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700A150g,
	"SUBCORE_VM_FIXED2720_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2720A150g,
	"SUBCORE_VM_FIXED2740_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2740A150g,
	"SUBCORE_VM_FIXED2760_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2760A150g,
	"SUBCORE_VM_FIXED2780_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2780A150g,
	"SUBCORE_VM_FIXED2800_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800A150g,
	"SUBCORE_VM_FIXED2820_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2820A150g,
	"SUBCORE_VM_FIXED2840_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2840A150g,
	"SUBCORE_VM_FIXED2860_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2860A150g,
	"SUBCORE_VM_FIXED2880_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880A150g,
	"SUBCORE_VM_FIXED2900_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900A150g,
	"SUBCORE_VM_FIXED2920_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2920A150g,
	"SUBCORE_VM_FIXED2940_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2940A150g,
	"SUBCORE_VM_FIXED2960_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2960A150g,
	"SUBCORE_VM_FIXED2980_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2980A150g,
	"SUBCORE_VM_FIXED3000_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000A150g,
	"SUBCORE_VM_FIXED3020_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3020A150g,
	"SUBCORE_VM_FIXED3040_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3040A150g,
	"SUBCORE_VM_FIXED3060_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060A150g,
	"SUBCORE_VM_FIXED3080_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3080A150g,
	"SUBCORE_VM_FIXED3100_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100A150g,
	"SUBCORE_VM_FIXED3120_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3120A150g,
	"SUBCORE_VM_FIXED3140_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3140A150g,
	"SUBCORE_VM_FIXED3160_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3160A150g,
	"SUBCORE_VM_FIXED3180_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3180A150g,
	"SUBCORE_VM_FIXED3200_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200A150g,
	"SUBCORE_VM_FIXED3220_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3220A150g,
	"SUBCORE_VM_FIXED3240_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240A150g,
	"SUBCORE_VM_FIXED3260_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3260A150g,
	"SUBCORE_VM_FIXED3280_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3280A150g,
	"SUBCORE_VM_FIXED3300_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300A150g,
	"SUBCORE_VM_FIXED3320_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3320A150g,
	"SUBCORE_VM_FIXED3340_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3340A150g,
	"SUBCORE_VM_FIXED3360_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3360A150g,
	"SUBCORE_VM_FIXED3380_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3380A150g,
	"SUBCORE_VM_FIXED3400_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400A150g,
	"SUBCORE_VM_FIXED3420_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420A150g,
	"SUBCORE_VM_FIXED3440_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3440A150g,
	"SUBCORE_VM_FIXED3460_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3460A150g,
	"SUBCORE_VM_FIXED3480_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3480A150g,
	"SUBCORE_VM_FIXED3500_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500A150g,
	"SUBCORE_VM_FIXED3520_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3520A150g,
	"SUBCORE_VM_FIXED3540_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3540A150g,
	"SUBCORE_VM_FIXED3560_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3560A150g,
	"SUBCORE_VM_FIXED3580_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3580A150g,
	"SUBCORE_VM_FIXED3600_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600A150g,
	"SUBCORE_VM_FIXED3620_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3620A150g,
	"SUBCORE_VM_FIXED3640_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3640A150g,
	"SUBCORE_VM_FIXED3660_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3660A150g,
	"SUBCORE_VM_FIXED3680_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3680A150g,
	"SUBCORE_VM_FIXED3700_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700A150g,
	"SUBCORE_VM_FIXED3720_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3720A150g,
	"SUBCORE_VM_FIXED3740_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3740A150g,
	"SUBCORE_VM_FIXED3760_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3760A150g,
	"SUBCORE_VM_FIXED3780_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780A150g,
	"SUBCORE_VM_FIXED3800_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800A150g,
	"SUBCORE_VM_FIXED3820_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3820A150g,
	"SUBCORE_VM_FIXED3840_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3840A150g,
	"SUBCORE_VM_FIXED3860_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3860A150g,
	"SUBCORE_VM_FIXED3880_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3880A150g,
	"SUBCORE_VM_FIXED3900_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900A150g,
	"SUBCORE_VM_FIXED3920_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3920A150g,
	"SUBCORE_VM_FIXED3940_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3940A150g,
	"SUBCORE_VM_FIXED3960_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960A150g,
	"SUBCORE_VM_FIXED3980_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3980A150g,
	"SUBCORE_VM_FIXED4000_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000A150g,
	"SUBCORE_VM_FIXED4020_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4020A150g,
	"SUBCORE_VM_FIXED4040_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4040A150g,
	"SUBCORE_VM_FIXED4060_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4060A150g,
	"SUBCORE_VM_FIXED4080_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4080A150g,
	"SUBCORE_VM_FIXED4100_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100A150g,
	"SUBCORE_VM_FIXED4120_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4120A150g,
	"SUBCORE_VM_FIXED4140_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140A150g,
	"SUBCORE_VM_FIXED4160_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4160A150g,
	"SUBCORE_VM_FIXED4180_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4180A150g,
	"SUBCORE_VM_FIXED4200_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200A150g,
	"SUBCORE_VM_FIXED4220_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4220A150g,
	"SUBCORE_VM_FIXED4240_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4240A150g,
	"SUBCORE_VM_FIXED4260_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4260A150g,
	"SUBCORE_VM_FIXED4280_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4280A150g,
	"SUBCORE_VM_FIXED4300_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300A150g,
	"SUBCORE_VM_FIXED4320_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320A150g,
	"SUBCORE_VM_FIXED4340_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4340A150g,
	"SUBCORE_VM_FIXED4360_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4360A150g,
	"SUBCORE_VM_FIXED4380_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4380A150g,
	"SUBCORE_VM_FIXED4400_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400A150g,
	"SUBCORE_VM_FIXED4420_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4420A150g,
	"SUBCORE_VM_FIXED4440_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4440A150g,
	"SUBCORE_VM_FIXED4460_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4460A150g,
	"SUBCORE_VM_FIXED4480_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4480A150g,
	"SUBCORE_VM_FIXED4500_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500A150g,
	"SUBCORE_VM_FIXED4520_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4520A150g,
	"SUBCORE_VM_FIXED4540_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4540A150g,
	"SUBCORE_VM_FIXED4560_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4560A150g,
	"SUBCORE_VM_FIXED4580_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4580A150g,
	"SUBCORE_VM_FIXED4600_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600A150g,
	"SUBCORE_VM_FIXED4620_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4620A150g,
	"SUBCORE_VM_FIXED4640_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4640A150g,
	"SUBCORE_VM_FIXED4660_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4660A150g,
	"SUBCORE_VM_FIXED4680_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680A150g,
	"SUBCORE_VM_FIXED4700_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700A150g,
	"SUBCORE_VM_FIXED4720_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4720A150g,
	"SUBCORE_VM_FIXED4740_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4740A150g,
	"SUBCORE_VM_FIXED4760_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4760A150g,
	"SUBCORE_VM_FIXED4780_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4780A150g,
	"SUBCORE_VM_FIXED4800_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800A150g,
	"SUBCORE_VM_FIXED4820_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4820A150g,
	"SUBCORE_VM_FIXED4840_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4840A150g,
	"SUBCORE_VM_FIXED4860_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860A150g,
	"SUBCORE_VM_FIXED4880_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4880A150g,
	"SUBCORE_VM_FIXED4900_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900A150g,
	"SUBCORE_VM_FIXED4920_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4920A150g,
	"SUBCORE_VM_FIXED4940_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4940A150g,
	"SUBCORE_VM_FIXED4960_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4960A150g,
	"SUBCORE_VM_FIXED4980_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4980A150g,
	"SUBCORE_VM_FIXED5000_A1_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000A150g,
	"SUBCORE_VM_FIXED0090_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0090X950g,
	"SUBCORE_VM_FIXED0180_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180X950g,
	"SUBCORE_VM_FIXED0270_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0270X950g,
	"SUBCORE_VM_FIXED0360_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360X950g,
	"SUBCORE_VM_FIXED0450_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450X950g,
	"SUBCORE_VM_FIXED0540_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540X950g,
	"SUBCORE_VM_FIXED0630_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0630X950g,
	"SUBCORE_VM_FIXED0720_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720X950g,
	"SUBCORE_VM_FIXED0810_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0810X950g,
	"SUBCORE_VM_FIXED0900_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900X950g,
	"SUBCORE_VM_FIXED0990_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0990X950g,
	"SUBCORE_VM_FIXED1080_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080X950g,
	"SUBCORE_VM_FIXED1170_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1170X950g,
	"SUBCORE_VM_FIXED1260_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260X950g,
	"SUBCORE_VM_FIXED1350_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350X950g,
	"SUBCORE_VM_FIXED1440_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440X950g,
	"SUBCORE_VM_FIXED1530_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1530X950g,
	"SUBCORE_VM_FIXED1620_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620X950g,
	"SUBCORE_VM_FIXED1710_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1710X950g,
	"SUBCORE_VM_FIXED1800_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800X950g,
	"SUBCORE_VM_FIXED1890_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1890X950g,
	"SUBCORE_VM_FIXED1980_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980X950g,
	"SUBCORE_VM_FIXED2070_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2070X950g,
	"SUBCORE_VM_FIXED2160_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160X950g,
	"SUBCORE_VM_FIXED2250_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250X950g,
	"SUBCORE_VM_FIXED2340_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340X950g,
	"SUBCORE_VM_FIXED2430_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2430X950g,
	"SUBCORE_VM_FIXED2520_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520X950g,
	"SUBCORE_VM_FIXED2610_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2610X950g,
	"SUBCORE_VM_FIXED2700_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700X950g,
	"SUBCORE_VM_FIXED2790_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2790X950g,
	"SUBCORE_VM_FIXED2880_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880X950g,
	"SUBCORE_VM_FIXED2970_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2970X950g,
	"SUBCORE_VM_FIXED3060_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060X950g,
	"SUBCORE_VM_FIXED3150_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150X950g,
	"SUBCORE_VM_FIXED3240_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240X950g,
	"SUBCORE_VM_FIXED3330_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3330X950g,
	"SUBCORE_VM_FIXED3420_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420X950g,
	"SUBCORE_VM_FIXED3510_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3510X950g,
	"SUBCORE_VM_FIXED3600_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600X950g,
	"SUBCORE_VM_FIXED3690_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3690X950g,
	"SUBCORE_VM_FIXED3780_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780X950g,
	"SUBCORE_VM_FIXED3870_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3870X950g,
	"SUBCORE_VM_FIXED3960_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960X950g,
	"SUBCORE_VM_FIXED4050_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050X950g,
	"SUBCORE_VM_FIXED4140_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140X950g,
	"SUBCORE_VM_FIXED4230_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4230X950g,
	"SUBCORE_VM_FIXED4320_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320X950g,
	"SUBCORE_VM_FIXED4410_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4410X950g,
	"SUBCORE_VM_FIXED4500_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500X950g,
	"SUBCORE_VM_FIXED4590_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4590X950g,
	"SUBCORE_VM_FIXED4680_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680X950g,
	"SUBCORE_VM_FIXED4770_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4770X950g,
	"SUBCORE_VM_FIXED4860_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860X950g,
	"SUBCORE_VM_FIXED4950_X9_50G":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950X950g,
	"DYNAMIC_A1_50G":                       CreateInternalVnicDetailsVnicShapeDynamicA150g,
	"FIXED0040_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0040A150g,
	"FIXED0100_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0100A150g,
	"FIXED0200_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0200A150g,
	"FIXED0300_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0300A150g,
	"FIXED0400_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0400A150g,
	"FIXED0500_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0500A150g,
	"FIXED0600_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0600A150g,
	"FIXED0700_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0700A150g,
	"FIXED0800_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0800A150g,
	"FIXED0900_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed0900A150g,
	"FIXED1000_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1000A150g,
	"FIXED1100_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1100A150g,
	"FIXED1200_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1200A150g,
	"FIXED1300_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1300A150g,
	"FIXED1400_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1400A150g,
	"FIXED1500_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1500A150g,
	"FIXED1600_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1600A150g,
	"FIXED1700_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1700A150g,
	"FIXED1800_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1800A150g,
	"FIXED1900_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed1900A150g,
	"FIXED2000_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2000A150g,
	"FIXED2100_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2100A150g,
	"FIXED2200_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2200A150g,
	"FIXED2300_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2300A150g,
	"FIXED2400_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2400A150g,
	"FIXED2500_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2500A150g,
	"FIXED2600_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2600A150g,
	"FIXED2700_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2700A150g,
	"FIXED2800_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2800A150g,
	"FIXED2900_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed2900A150g,
	"FIXED3000_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3000A150g,
	"FIXED3100_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3100A150g,
	"FIXED3200_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3200A150g,
	"FIXED3300_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3300A150g,
	"FIXED3400_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3400A150g,
	"FIXED3500_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3500A150g,
	"FIXED3600_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3600A150g,
	"FIXED3700_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3700A150g,
	"FIXED3800_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3800A150g,
	"FIXED3900_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed3900A150g,
	"FIXED4000_A1_50G":                     CreateInternalVnicDetailsVnicShapeFixed4000A150g,
	"ENTIREHOST_A1_50G":                    CreateInternalVnicDetailsVnicShapeEntirehostA150g,
	"DYNAMIC_X9_50G":                       CreateInternalVnicDetailsVnicShapeDynamicX950g,
	"FIXED0040_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed0040X950g,
	"FIXED0400_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed0400X950g,
	"FIXED0800_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed0800X950g,
	"FIXED1200_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed1200X950g,
	"FIXED1600_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed1600X950g,
	"FIXED2000_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed2000X950g,
	"FIXED2400_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed2400X950g,
	"FIXED2800_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed2800X950g,
	"FIXED3200_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed3200X950g,
	"FIXED3600_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed3600X950g,
	"FIXED4000_X9_50G":                     CreateInternalVnicDetailsVnicShapeFixed4000X950g,
	"STANDARD_VM_FIXED0100_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0100X950g,
	"STANDARD_VM_FIXED0200_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0200X950g,
	"STANDARD_VM_FIXED0300_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0300X950g,
	"STANDARD_VM_FIXED0400_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0400X950g,
	"STANDARD_VM_FIXED0500_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0500X950g,
	"STANDARD_VM_FIXED0600_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0600X950g,
	"STANDARD_VM_FIXED0700_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0700X950g,
	"STANDARD_VM_FIXED0800_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0800X950g,
	"STANDARD_VM_FIXED0900_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0900X950g,
	"STANDARD_VM_FIXED1000_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1000X950g,
	"STANDARD_VM_FIXED1100_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1100X950g,
	"STANDARD_VM_FIXED1200_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1200X950g,
	"STANDARD_VM_FIXED1300_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1300X950g,
	"STANDARD_VM_FIXED1400_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1400X950g,
	"STANDARD_VM_FIXED1500_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1500X950g,
	"STANDARD_VM_FIXED1600_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1600X950g,
	"STANDARD_VM_FIXED1700_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1700X950g,
	"STANDARD_VM_FIXED1800_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1800X950g,
	"STANDARD_VM_FIXED1900_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1900X950g,
	"STANDARD_VM_FIXED2000_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2000X950g,
	"STANDARD_VM_FIXED2100_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2100X950g,
	"STANDARD_VM_FIXED2200_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2200X950g,
	"STANDARD_VM_FIXED2300_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2300X950g,
	"STANDARD_VM_FIXED2400_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2400X950g,
	"STANDARD_VM_FIXED2500_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2500X950g,
	"STANDARD_VM_FIXED2600_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2600X950g,
	"STANDARD_VM_FIXED2700_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2700X950g,
	"STANDARD_VM_FIXED2800_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2800X950g,
	"STANDARD_VM_FIXED2900_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2900X950g,
	"STANDARD_VM_FIXED3000_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3000X950g,
	"STANDARD_VM_FIXED3100_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3100X950g,
	"STANDARD_VM_FIXED3200_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3200X950g,
	"STANDARD_VM_FIXED3300_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3300X950g,
	"STANDARD_VM_FIXED3400_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3400X950g,
	"STANDARD_VM_FIXED3500_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3500X950g,
	"STANDARD_VM_FIXED3600_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3600X950g,
	"STANDARD_VM_FIXED3700_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3700X950g,
	"STANDARD_VM_FIXED3800_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3800X950g,
	"STANDARD_VM_FIXED3900_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3900X950g,
	"STANDARD_VM_FIXED4000_X9_50G":         CreateInternalVnicDetailsVnicShapeStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED0025_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"SUBCORE_STANDARD_VM_FIXED0050_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"SUBCORE_STANDARD_VM_FIXED0075_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"SUBCORE_STANDARD_VM_FIXED0100_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"SUBCORE_STANDARD_VM_FIXED0125_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"SUBCORE_STANDARD_VM_FIXED0150_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"SUBCORE_STANDARD_VM_FIXED0175_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"SUBCORE_STANDARD_VM_FIXED0200_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"SUBCORE_STANDARD_VM_FIXED0225_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"SUBCORE_STANDARD_VM_FIXED0250_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"SUBCORE_STANDARD_VM_FIXED0275_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"SUBCORE_STANDARD_VM_FIXED0300_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"SUBCORE_STANDARD_VM_FIXED0325_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"SUBCORE_STANDARD_VM_FIXED0350_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"SUBCORE_STANDARD_VM_FIXED0375_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"SUBCORE_STANDARD_VM_FIXED0400_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"SUBCORE_STANDARD_VM_FIXED0425_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"SUBCORE_STANDARD_VM_FIXED0450_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"SUBCORE_STANDARD_VM_FIXED0475_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"SUBCORE_STANDARD_VM_FIXED0500_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"SUBCORE_STANDARD_VM_FIXED0525_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"SUBCORE_STANDARD_VM_FIXED0550_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"SUBCORE_STANDARD_VM_FIXED0575_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"SUBCORE_STANDARD_VM_FIXED0600_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"SUBCORE_STANDARD_VM_FIXED0625_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"SUBCORE_STANDARD_VM_FIXED0650_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"SUBCORE_STANDARD_VM_FIXED0675_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"SUBCORE_STANDARD_VM_FIXED0700_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"SUBCORE_STANDARD_VM_FIXED0725_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"SUBCORE_STANDARD_VM_FIXED0750_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"SUBCORE_STANDARD_VM_FIXED0775_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"SUBCORE_STANDARD_VM_FIXED0800_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"SUBCORE_STANDARD_VM_FIXED0825_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"SUBCORE_STANDARD_VM_FIXED0850_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"SUBCORE_STANDARD_VM_FIXED0875_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"SUBCORE_STANDARD_VM_FIXED0900_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"SUBCORE_STANDARD_VM_FIXED0925_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"SUBCORE_STANDARD_VM_FIXED0950_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"SUBCORE_STANDARD_VM_FIXED0975_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"SUBCORE_STANDARD_VM_FIXED1000_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"SUBCORE_STANDARD_VM_FIXED1025_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"SUBCORE_STANDARD_VM_FIXED1050_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"SUBCORE_STANDARD_VM_FIXED1075_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"SUBCORE_STANDARD_VM_FIXED1100_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"SUBCORE_STANDARD_VM_FIXED1125_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"SUBCORE_STANDARD_VM_FIXED1150_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"SUBCORE_STANDARD_VM_FIXED1175_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"SUBCORE_STANDARD_VM_FIXED1200_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"SUBCORE_STANDARD_VM_FIXED1225_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"SUBCORE_STANDARD_VM_FIXED1250_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"SUBCORE_STANDARD_VM_FIXED1275_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"SUBCORE_STANDARD_VM_FIXED1300_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"SUBCORE_STANDARD_VM_FIXED1325_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"SUBCORE_STANDARD_VM_FIXED1350_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"SUBCORE_STANDARD_VM_FIXED1375_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"SUBCORE_STANDARD_VM_FIXED1400_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"SUBCORE_STANDARD_VM_FIXED1425_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"SUBCORE_STANDARD_VM_FIXED1450_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"SUBCORE_STANDARD_VM_FIXED1475_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"SUBCORE_STANDARD_VM_FIXED1500_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"SUBCORE_STANDARD_VM_FIXED1525_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"SUBCORE_STANDARD_VM_FIXED1550_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"SUBCORE_STANDARD_VM_FIXED1575_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"SUBCORE_STANDARD_VM_FIXED1600_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"SUBCORE_STANDARD_VM_FIXED1625_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"SUBCORE_STANDARD_VM_FIXED1650_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"SUBCORE_STANDARD_VM_FIXED1700_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"SUBCORE_STANDARD_VM_FIXED1725_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"SUBCORE_STANDARD_VM_FIXED1750_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"SUBCORE_STANDARD_VM_FIXED1800_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"SUBCORE_STANDARD_VM_FIXED1850_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"SUBCORE_STANDARD_VM_FIXED1875_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"SUBCORE_STANDARD_VM_FIXED1900_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"SUBCORE_STANDARD_VM_FIXED1925_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"SUBCORE_STANDARD_VM_FIXED1950_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"SUBCORE_STANDARD_VM_FIXED2000_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"SUBCORE_STANDARD_VM_FIXED2025_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"SUBCORE_STANDARD_VM_FIXED2050_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"SUBCORE_STANDARD_VM_FIXED2100_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"SUBCORE_STANDARD_VM_FIXED2125_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"SUBCORE_STANDARD_VM_FIXED2150_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"SUBCORE_STANDARD_VM_FIXED2175_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"SUBCORE_STANDARD_VM_FIXED2200_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"SUBCORE_STANDARD_VM_FIXED2250_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"SUBCORE_STANDARD_VM_FIXED2275_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"SUBCORE_STANDARD_VM_FIXED2300_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"SUBCORE_STANDARD_VM_FIXED2325_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"SUBCORE_STANDARD_VM_FIXED2350_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"SUBCORE_STANDARD_VM_FIXED2375_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"SUBCORE_STANDARD_VM_FIXED2400_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"SUBCORE_STANDARD_VM_FIXED2450_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"SUBCORE_STANDARD_VM_FIXED2475_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"SUBCORE_STANDARD_VM_FIXED2500_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"SUBCORE_STANDARD_VM_FIXED2550_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"SUBCORE_STANDARD_VM_FIXED2600_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"SUBCORE_STANDARD_VM_FIXED2625_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"SUBCORE_STANDARD_VM_FIXED2650_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"SUBCORE_STANDARD_VM_FIXED2700_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"SUBCORE_STANDARD_VM_FIXED2750_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"SUBCORE_STANDARD_VM_FIXED2775_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"SUBCORE_STANDARD_VM_FIXED2800_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"SUBCORE_STANDARD_VM_FIXED2850_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"SUBCORE_STANDARD_VM_FIXED2875_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"SUBCORE_STANDARD_VM_FIXED2900_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"SUBCORE_STANDARD_VM_FIXED2925_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"SUBCORE_STANDARD_VM_FIXED2950_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"SUBCORE_STANDARD_VM_FIXED2975_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"SUBCORE_STANDARD_VM_FIXED3000_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"SUBCORE_STANDARD_VM_FIXED3025_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"SUBCORE_STANDARD_VM_FIXED3050_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"SUBCORE_STANDARD_VM_FIXED3075_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"SUBCORE_STANDARD_VM_FIXED3100_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"SUBCORE_STANDARD_VM_FIXED3125_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"SUBCORE_STANDARD_VM_FIXED3150_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"SUBCORE_STANDARD_VM_FIXED3200_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"SUBCORE_STANDARD_VM_FIXED3225_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"SUBCORE_STANDARD_VM_FIXED3250_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"SUBCORE_STANDARD_VM_FIXED3300_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"SUBCORE_STANDARD_VM_FIXED3325_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"SUBCORE_STANDARD_VM_FIXED3375_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"SUBCORE_STANDARD_VM_FIXED3400_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"SUBCORE_STANDARD_VM_FIXED3450_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"SUBCORE_STANDARD_VM_FIXED3500_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"SUBCORE_STANDARD_VM_FIXED3525_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"SUBCORE_STANDARD_VM_FIXED3575_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"SUBCORE_STANDARD_VM_FIXED3600_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"SUBCORE_STANDARD_VM_FIXED3625_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"SUBCORE_STANDARD_VM_FIXED3675_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"SUBCORE_STANDARD_VM_FIXED3700_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"SUBCORE_STANDARD_VM_FIXED3750_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"SUBCORE_STANDARD_VM_FIXED3800_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"SUBCORE_STANDARD_VM_FIXED3825_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"SUBCORE_STANDARD_VM_FIXED3850_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"SUBCORE_STANDARD_VM_FIXED3875_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"SUBCORE_STANDARD_VM_FIXED3900_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"SUBCORE_STANDARD_VM_FIXED3975_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"SUBCORE_STANDARD_VM_FIXED4000_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"SUBCORE_STANDARD_VM_FIXED4025_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"SUBCORE_STANDARD_VM_FIXED4050_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"SUBCORE_STANDARD_VM_FIXED4100_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"SUBCORE_STANDARD_VM_FIXED4125_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"SUBCORE_STANDARD_VM_FIXED4200_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"SUBCORE_STANDARD_VM_FIXED4225_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"SUBCORE_STANDARD_VM_FIXED4250_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"SUBCORE_STANDARD_VM_FIXED4275_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"SUBCORE_STANDARD_VM_FIXED4300_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"SUBCORE_STANDARD_VM_FIXED4350_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"SUBCORE_STANDARD_VM_FIXED4375_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"SUBCORE_STANDARD_VM_FIXED4400_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"SUBCORE_STANDARD_VM_FIXED4425_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"SUBCORE_STANDARD_VM_FIXED4500_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"SUBCORE_STANDARD_VM_FIXED4550_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"SUBCORE_STANDARD_VM_FIXED4575_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"SUBCORE_STANDARD_VM_FIXED4600_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"SUBCORE_STANDARD_VM_FIXED4625_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"SUBCORE_STANDARD_VM_FIXED4650_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"SUBCORE_STANDARD_VM_FIXED4675_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"SUBCORE_STANDARD_VM_FIXED4700_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"SUBCORE_STANDARD_VM_FIXED4725_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"SUBCORE_STANDARD_VM_FIXED4750_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"SUBCORE_STANDARD_VM_FIXED4800_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"SUBCORE_STANDARD_VM_FIXED4875_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"SUBCORE_STANDARD_VM_FIXED4900_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"SUBCORE_STANDARD_VM_FIXED4950_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"SUBCORE_STANDARD_VM_FIXED5000_X9_50G": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"ENTIREHOST_X9_50G":                    CreateInternalVnicDetailsVnicShapeEntirehostX950g,
}

var mappingCreateInternalVnicDetailsVnicShapeEnumLowerCase = map[string]CreateInternalVnicDetailsVnicShapeEnum{
	"dynamic":                              CreateInternalVnicDetailsVnicShapeDynamic,
	"fixed0040":                            CreateInternalVnicDetailsVnicShapeFixed0040,
	"fixed0060":                            CreateInternalVnicDetailsVnicShapeFixed0060,
	"fixed0060_psm":                        CreateInternalVnicDetailsVnicShapeFixed0060Psm,
	"fixed0100":                            CreateInternalVnicDetailsVnicShapeFixed0100,
	"fixed0120":                            CreateInternalVnicDetailsVnicShapeFixed0120,
	"fixed0120_2x":                         CreateInternalVnicDetailsVnicShapeFixed01202x,
	"fixed0200":                            CreateInternalVnicDetailsVnicShapeFixed0200,
	"fixed0240":                            CreateInternalVnicDetailsVnicShapeFixed0240,
	"fixed0480":                            CreateInternalVnicDetailsVnicShapeFixed0480,
	"entirehost":                           CreateInternalVnicDetailsVnicShapeEntirehost,
	"dynamic_25g":                          CreateInternalVnicDetailsVnicShapeDynamic25g,
	"fixed0040_25g":                        CreateInternalVnicDetailsVnicShapeFixed004025g,
	"fixed0100_25g":                        CreateInternalVnicDetailsVnicShapeFixed010025g,
	"fixed0200_25g":                        CreateInternalVnicDetailsVnicShapeFixed020025g,
	"fixed0400_25g":                        CreateInternalVnicDetailsVnicShapeFixed040025g,
	"fixed0800_25g":                        CreateInternalVnicDetailsVnicShapeFixed080025g,
	"fixed1600_25g":                        CreateInternalVnicDetailsVnicShapeFixed160025g,
	"fixed2400_25g":                        CreateInternalVnicDetailsVnicShapeFixed240025g,
	"entirehost_25g":                       CreateInternalVnicDetailsVnicShapeEntirehost25g,
	"dynamic_e1_25g":                       CreateInternalVnicDetailsVnicShapeDynamicE125g,
	"fixed0040_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0040E125g,
	"fixed0070_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0070E125g,
	"fixed0140_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0140E125g,
	"fixed0280_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0280E125g,
	"fixed0560_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0560E125g,
	"fixed1120_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed1120E125g,
	"fixed1680_e1_25g":                     CreateInternalVnicDetailsVnicShapeFixed1680E125g,
	"entirehost_e1_25g":                    CreateInternalVnicDetailsVnicShapeEntirehostE125g,
	"dynamic_b1_25g":                       CreateInternalVnicDetailsVnicShapeDynamicB125g,
	"fixed0040_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0040B125g,
	"fixed0060_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0060B125g,
	"fixed0120_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0120B125g,
	"fixed0240_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0240B125g,
	"fixed0480_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0480B125g,
	"fixed0960_b1_25g":                     CreateInternalVnicDetailsVnicShapeFixed0960B125g,
	"entirehost_b1_25g":                    CreateInternalVnicDetailsVnicShapeEntirehostB125g,
	"micro_vm_fixed0048_e1_25g":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0048E125g,
	"micro_lb_fixed0001_e1_25g":            CreateInternalVnicDetailsVnicShapeMicroLbFixed0001E125g,
	"vnicaas_fixed0200":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0200,
	"vnicaas_fixed0400":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0400,
	"vnicaas_fixed0700":                    CreateInternalVnicDetailsVnicShapeVnicaasFixed0700,
	"vnicaas_nlb_approved_10g":             CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved10g,
	"vnicaas_nlb_approved_25g":             CreateInternalVnicDetailsVnicShapeVnicaasNlbApproved25g,
	"vnicaas_telesis_25g":                  CreateInternalVnicDetailsVnicShapeVnicaasTelesis25g,
	"vnicaas_telesis_10g":                  CreateInternalVnicDetailsVnicShapeVnicaasTelesis10g,
	"vnicaas_ambassador_fixed0100":         CreateInternalVnicDetailsVnicShapeVnicaasAmbassadorFixed0100,
	"vnicaas_telesis_gamma":                CreateInternalVnicDetailsVnicShapeVnicaasTelesisGamma,
	"vnicaas_privatedns":                   CreateInternalVnicDetailsVnicShapeVnicaasPrivatedns,
	"vnicaas_fwaas":                        CreateInternalVnicDetailsVnicShapeVnicaasFwaas,
	"vnicaas_lbaas_free":                   CreateInternalVnicDetailsVnicShapeVnicaasLbaasFree,
	"vnicaas_lbaas_8g_512k":                CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g512k,
	"vnicaas_lbaas_8g_1m":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g1m,
	"vnicaas_lbaas_8g_2m":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g2m,
	"vnicaas_lbaas_8g_3m":                  CreateInternalVnicDetailsVnicShapeVnicaasLbaas8g3m,
	"dynamic_e3_50g":                       CreateInternalVnicDetailsVnicShapeDynamicE350g,
	"fixed0040_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0040E350g,
	"fixed0100_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0100E350g,
	"fixed0200_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0200E350g,
	"fixed0300_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0300E350g,
	"fixed0400_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0400E350g,
	"fixed0500_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0500E350g,
	"fixed0600_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0600E350g,
	"fixed0700_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0700E350g,
	"fixed0800_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0800E350g,
	"fixed0900_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed0900E350g,
	"fixed1000_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1000E350g,
	"fixed1100_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1100E350g,
	"fixed1200_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1200E350g,
	"fixed1300_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1300E350g,
	"fixed1400_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1400E350g,
	"fixed1500_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1500E350g,
	"fixed1600_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1600E350g,
	"fixed1700_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1700E350g,
	"fixed1800_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1800E350g,
	"fixed1900_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed1900E350g,
	"fixed2000_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2000E350g,
	"fixed2100_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2100E350g,
	"fixed2200_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2200E350g,
	"fixed2300_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2300E350g,
	"fixed2400_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2400E350g,
	"fixed2500_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2500E350g,
	"fixed2600_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2600E350g,
	"fixed2700_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2700E350g,
	"fixed2800_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2800E350g,
	"fixed2900_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed2900E350g,
	"fixed3000_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3000E350g,
	"fixed3100_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3100E350g,
	"fixed3200_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3200E350g,
	"fixed3300_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3300E350g,
	"fixed3400_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3400E350g,
	"fixed3500_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3500E350g,
	"fixed3600_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3600E350g,
	"fixed3700_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3700E350g,
	"fixed3800_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3800E350g,
	"fixed3900_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed3900E350g,
	"fixed4000_e3_50g":                     CreateInternalVnicDetailsVnicShapeFixed4000E350g,
	"entirehost_e3_50g":                    CreateInternalVnicDetailsVnicShapeEntirehostE350g,
	"dynamic_e4_50g":                       CreateInternalVnicDetailsVnicShapeDynamicE450g,
	"fixed0040_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0040E450g,
	"fixed0100_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0100E450g,
	"fixed0200_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0200E450g,
	"fixed0300_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0300E450g,
	"fixed0400_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0400E450g,
	"fixed0500_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0500E450g,
	"fixed0600_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0600E450g,
	"fixed0700_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0700E450g,
	"fixed0800_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0800E450g,
	"fixed0900_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed0900E450g,
	"fixed1000_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1000E450g,
	"fixed1100_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1100E450g,
	"fixed1200_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1200E450g,
	"fixed1300_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1300E450g,
	"fixed1400_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1400E450g,
	"fixed1500_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1500E450g,
	"fixed1600_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1600E450g,
	"fixed1700_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1700E450g,
	"fixed1800_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1800E450g,
	"fixed1900_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed1900E450g,
	"fixed2000_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2000E450g,
	"fixed2100_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2100E450g,
	"fixed2200_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2200E450g,
	"fixed2300_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2300E450g,
	"fixed2400_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2400E450g,
	"fixed2500_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2500E450g,
	"fixed2600_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2600E450g,
	"fixed2700_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2700E450g,
	"fixed2800_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2800E450g,
	"fixed2900_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed2900E450g,
	"fixed3000_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3000E450g,
	"fixed3100_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3100E450g,
	"fixed3200_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3200E450g,
	"fixed3300_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3300E450g,
	"fixed3400_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3400E450g,
	"fixed3500_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3500E450g,
	"fixed3600_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3600E450g,
	"fixed3700_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3700E450g,
	"fixed3800_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3800E450g,
	"fixed3900_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed3900E450g,
	"fixed4000_e4_50g":                     CreateInternalVnicDetailsVnicShapeFixed4000E450g,
	"entirehost_e4_50g":                    CreateInternalVnicDetailsVnicShapeEntirehostE450g,
	"micro_vm_fixed0050_e3_50g":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E350g,
	"micro_vm_fixed0050_e4_50g":            CreateInternalVnicDetailsVnicShapeMicroVmFixed0050E450g,
	"subcore_vm_fixed0025_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E350g,
	"subcore_vm_fixed0050_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E350g,
	"subcore_vm_fixed0075_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E350g,
	"subcore_vm_fixed0100_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E350g,
	"subcore_vm_fixed0125_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E350g,
	"subcore_vm_fixed0150_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E350g,
	"subcore_vm_fixed0175_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E350g,
	"subcore_vm_fixed0200_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E350g,
	"subcore_vm_fixed0225_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E350g,
	"subcore_vm_fixed0250_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E350g,
	"subcore_vm_fixed0275_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E350g,
	"subcore_vm_fixed0300_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E350g,
	"subcore_vm_fixed0325_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E350g,
	"subcore_vm_fixed0350_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E350g,
	"subcore_vm_fixed0375_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E350g,
	"subcore_vm_fixed0400_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E350g,
	"subcore_vm_fixed0425_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E350g,
	"subcore_vm_fixed0450_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E350g,
	"subcore_vm_fixed0475_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E350g,
	"subcore_vm_fixed0500_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E350g,
	"subcore_vm_fixed0525_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E350g,
	"subcore_vm_fixed0550_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E350g,
	"subcore_vm_fixed0575_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E350g,
	"subcore_vm_fixed0600_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E350g,
	"subcore_vm_fixed0625_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E350g,
	"subcore_vm_fixed0650_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E350g,
	"subcore_vm_fixed0675_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E350g,
	"subcore_vm_fixed0700_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E350g,
	"subcore_vm_fixed0725_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E350g,
	"subcore_vm_fixed0750_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E350g,
	"subcore_vm_fixed0775_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E350g,
	"subcore_vm_fixed0800_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E350g,
	"subcore_vm_fixed0825_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E350g,
	"subcore_vm_fixed0850_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E350g,
	"subcore_vm_fixed0875_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E350g,
	"subcore_vm_fixed0900_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E350g,
	"subcore_vm_fixed0925_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E350g,
	"subcore_vm_fixed0950_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E350g,
	"subcore_vm_fixed0975_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E350g,
	"subcore_vm_fixed1000_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E350g,
	"subcore_vm_fixed1025_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E350g,
	"subcore_vm_fixed1050_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E350g,
	"subcore_vm_fixed1075_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E350g,
	"subcore_vm_fixed1100_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E350g,
	"subcore_vm_fixed1125_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E350g,
	"subcore_vm_fixed1150_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E350g,
	"subcore_vm_fixed1175_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E350g,
	"subcore_vm_fixed1200_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E350g,
	"subcore_vm_fixed1225_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E350g,
	"subcore_vm_fixed1250_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E350g,
	"subcore_vm_fixed1275_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E350g,
	"subcore_vm_fixed1300_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E350g,
	"subcore_vm_fixed1325_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E350g,
	"subcore_vm_fixed1350_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E350g,
	"subcore_vm_fixed1375_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E350g,
	"subcore_vm_fixed1400_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E350g,
	"subcore_vm_fixed1425_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E350g,
	"subcore_vm_fixed1450_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E350g,
	"subcore_vm_fixed1475_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E350g,
	"subcore_vm_fixed1500_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E350g,
	"subcore_vm_fixed1525_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E350g,
	"subcore_vm_fixed1550_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E350g,
	"subcore_vm_fixed1575_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E350g,
	"subcore_vm_fixed1600_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E350g,
	"subcore_vm_fixed1625_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E350g,
	"subcore_vm_fixed1650_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E350g,
	"subcore_vm_fixed1700_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E350g,
	"subcore_vm_fixed1725_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E350g,
	"subcore_vm_fixed1750_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E350g,
	"subcore_vm_fixed1800_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E350g,
	"subcore_vm_fixed1850_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E350g,
	"subcore_vm_fixed1875_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E350g,
	"subcore_vm_fixed1900_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E350g,
	"subcore_vm_fixed1925_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E350g,
	"subcore_vm_fixed1950_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E350g,
	"subcore_vm_fixed2000_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E350g,
	"subcore_vm_fixed2025_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E350g,
	"subcore_vm_fixed2050_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E350g,
	"subcore_vm_fixed2100_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E350g,
	"subcore_vm_fixed2125_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E350g,
	"subcore_vm_fixed2150_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E350g,
	"subcore_vm_fixed2175_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E350g,
	"subcore_vm_fixed2200_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E350g,
	"subcore_vm_fixed2250_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E350g,
	"subcore_vm_fixed2275_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E350g,
	"subcore_vm_fixed2300_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E350g,
	"subcore_vm_fixed2325_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E350g,
	"subcore_vm_fixed2350_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E350g,
	"subcore_vm_fixed2375_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E350g,
	"subcore_vm_fixed2400_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E350g,
	"subcore_vm_fixed2450_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E350g,
	"subcore_vm_fixed2475_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E350g,
	"subcore_vm_fixed2500_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E350g,
	"subcore_vm_fixed2550_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E350g,
	"subcore_vm_fixed2600_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E350g,
	"subcore_vm_fixed2625_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E350g,
	"subcore_vm_fixed2650_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E350g,
	"subcore_vm_fixed2700_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E350g,
	"subcore_vm_fixed2750_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E350g,
	"subcore_vm_fixed2775_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E350g,
	"subcore_vm_fixed2800_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E350g,
	"subcore_vm_fixed2850_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E350g,
	"subcore_vm_fixed2875_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E350g,
	"subcore_vm_fixed2900_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E350g,
	"subcore_vm_fixed2925_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E350g,
	"subcore_vm_fixed2950_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E350g,
	"subcore_vm_fixed2975_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E350g,
	"subcore_vm_fixed3000_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E350g,
	"subcore_vm_fixed3025_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E350g,
	"subcore_vm_fixed3050_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E350g,
	"subcore_vm_fixed3075_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E350g,
	"subcore_vm_fixed3100_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E350g,
	"subcore_vm_fixed3125_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E350g,
	"subcore_vm_fixed3150_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E350g,
	"subcore_vm_fixed3200_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E350g,
	"subcore_vm_fixed3225_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E350g,
	"subcore_vm_fixed3250_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E350g,
	"subcore_vm_fixed3300_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E350g,
	"subcore_vm_fixed3325_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E350g,
	"subcore_vm_fixed3375_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E350g,
	"subcore_vm_fixed3400_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E350g,
	"subcore_vm_fixed3450_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E350g,
	"subcore_vm_fixed3500_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E350g,
	"subcore_vm_fixed3525_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E350g,
	"subcore_vm_fixed3575_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E350g,
	"subcore_vm_fixed3600_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E350g,
	"subcore_vm_fixed3625_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E350g,
	"subcore_vm_fixed3675_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E350g,
	"subcore_vm_fixed3700_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E350g,
	"subcore_vm_fixed3750_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E350g,
	"subcore_vm_fixed3800_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E350g,
	"subcore_vm_fixed3825_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E350g,
	"subcore_vm_fixed3850_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E350g,
	"subcore_vm_fixed3875_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E350g,
	"subcore_vm_fixed3900_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E350g,
	"subcore_vm_fixed3975_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E350g,
	"subcore_vm_fixed4000_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E350g,
	"subcore_vm_fixed4025_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E350g,
	"subcore_vm_fixed4050_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E350g,
	"subcore_vm_fixed4100_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E350g,
	"subcore_vm_fixed4125_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E350g,
	"subcore_vm_fixed4200_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E350g,
	"subcore_vm_fixed4225_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E350g,
	"subcore_vm_fixed4250_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E350g,
	"subcore_vm_fixed4275_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E350g,
	"subcore_vm_fixed4300_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E350g,
	"subcore_vm_fixed4350_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E350g,
	"subcore_vm_fixed4375_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E350g,
	"subcore_vm_fixed4400_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E350g,
	"subcore_vm_fixed4425_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E350g,
	"subcore_vm_fixed4500_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E350g,
	"subcore_vm_fixed4550_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E350g,
	"subcore_vm_fixed4575_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E350g,
	"subcore_vm_fixed4600_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E350g,
	"subcore_vm_fixed4625_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E350g,
	"subcore_vm_fixed4650_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E350g,
	"subcore_vm_fixed4675_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E350g,
	"subcore_vm_fixed4700_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E350g,
	"subcore_vm_fixed4725_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E350g,
	"subcore_vm_fixed4750_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E350g,
	"subcore_vm_fixed4800_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E350g,
	"subcore_vm_fixed4875_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E350g,
	"subcore_vm_fixed4900_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E350g,
	"subcore_vm_fixed4950_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E350g,
	"subcore_vm_fixed5000_e3_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E350g,
	"subcore_vm_fixed0025_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0025E450g,
	"subcore_vm_fixed0050_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0050E450g,
	"subcore_vm_fixed0075_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0075E450g,
	"subcore_vm_fixed0100_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100E450g,
	"subcore_vm_fixed0125_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0125E450g,
	"subcore_vm_fixed0150_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0150E450g,
	"subcore_vm_fixed0175_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0175E450g,
	"subcore_vm_fixed0200_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200E450g,
	"subcore_vm_fixed0225_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0225E450g,
	"subcore_vm_fixed0250_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0250E450g,
	"subcore_vm_fixed0275_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0275E450g,
	"subcore_vm_fixed0300_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300E450g,
	"subcore_vm_fixed0325_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0325E450g,
	"subcore_vm_fixed0350_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0350E450g,
	"subcore_vm_fixed0375_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0375E450g,
	"subcore_vm_fixed0400_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400E450g,
	"subcore_vm_fixed0425_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0425E450g,
	"subcore_vm_fixed0450_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450E450g,
	"subcore_vm_fixed0475_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0475E450g,
	"subcore_vm_fixed0500_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500E450g,
	"subcore_vm_fixed0525_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0525E450g,
	"subcore_vm_fixed0550_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0550E450g,
	"subcore_vm_fixed0575_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0575E450g,
	"subcore_vm_fixed0600_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600E450g,
	"subcore_vm_fixed0625_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0625E450g,
	"subcore_vm_fixed0650_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0650E450g,
	"subcore_vm_fixed0675_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0675E450g,
	"subcore_vm_fixed0700_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700E450g,
	"subcore_vm_fixed0725_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0725E450g,
	"subcore_vm_fixed0750_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0750E450g,
	"subcore_vm_fixed0775_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0775E450g,
	"subcore_vm_fixed0800_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800E450g,
	"subcore_vm_fixed0825_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0825E450g,
	"subcore_vm_fixed0850_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0850E450g,
	"subcore_vm_fixed0875_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0875E450g,
	"subcore_vm_fixed0900_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900E450g,
	"subcore_vm_fixed0925_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0925E450g,
	"subcore_vm_fixed0950_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0950E450g,
	"subcore_vm_fixed0975_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0975E450g,
	"subcore_vm_fixed1000_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000E450g,
	"subcore_vm_fixed1025_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1025E450g,
	"subcore_vm_fixed1050_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1050E450g,
	"subcore_vm_fixed1075_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1075E450g,
	"subcore_vm_fixed1100_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100E450g,
	"subcore_vm_fixed1125_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1125E450g,
	"subcore_vm_fixed1150_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1150E450g,
	"subcore_vm_fixed1175_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1175E450g,
	"subcore_vm_fixed1200_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200E450g,
	"subcore_vm_fixed1225_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1225E450g,
	"subcore_vm_fixed1250_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1250E450g,
	"subcore_vm_fixed1275_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1275E450g,
	"subcore_vm_fixed1300_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300E450g,
	"subcore_vm_fixed1325_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1325E450g,
	"subcore_vm_fixed1350_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350E450g,
	"subcore_vm_fixed1375_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1375E450g,
	"subcore_vm_fixed1400_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400E450g,
	"subcore_vm_fixed1425_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1425E450g,
	"subcore_vm_fixed1450_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1450E450g,
	"subcore_vm_fixed1475_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1475E450g,
	"subcore_vm_fixed1500_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500E450g,
	"subcore_vm_fixed1525_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1525E450g,
	"subcore_vm_fixed1550_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1550E450g,
	"subcore_vm_fixed1575_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1575E450g,
	"subcore_vm_fixed1600_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600E450g,
	"subcore_vm_fixed1625_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1625E450g,
	"subcore_vm_fixed1650_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1650E450g,
	"subcore_vm_fixed1700_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700E450g,
	"subcore_vm_fixed1725_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1725E450g,
	"subcore_vm_fixed1750_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1750E450g,
	"subcore_vm_fixed1800_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800E450g,
	"subcore_vm_fixed1850_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1850E450g,
	"subcore_vm_fixed1875_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1875E450g,
	"subcore_vm_fixed1900_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900E450g,
	"subcore_vm_fixed1925_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1925E450g,
	"subcore_vm_fixed1950_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1950E450g,
	"subcore_vm_fixed2000_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000E450g,
	"subcore_vm_fixed2025_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2025E450g,
	"subcore_vm_fixed2050_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2050E450g,
	"subcore_vm_fixed2100_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100E450g,
	"subcore_vm_fixed2125_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2125E450g,
	"subcore_vm_fixed2150_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2150E450g,
	"subcore_vm_fixed2175_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2175E450g,
	"subcore_vm_fixed2200_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200E450g,
	"subcore_vm_fixed2250_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250E450g,
	"subcore_vm_fixed2275_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2275E450g,
	"subcore_vm_fixed2300_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300E450g,
	"subcore_vm_fixed2325_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2325E450g,
	"subcore_vm_fixed2350_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2350E450g,
	"subcore_vm_fixed2375_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2375E450g,
	"subcore_vm_fixed2400_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400E450g,
	"subcore_vm_fixed2450_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2450E450g,
	"subcore_vm_fixed2475_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2475E450g,
	"subcore_vm_fixed2500_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500E450g,
	"subcore_vm_fixed2550_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2550E450g,
	"subcore_vm_fixed2600_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600E450g,
	"subcore_vm_fixed2625_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2625E450g,
	"subcore_vm_fixed2650_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2650E450g,
	"subcore_vm_fixed2700_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700E450g,
	"subcore_vm_fixed2750_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2750E450g,
	"subcore_vm_fixed2775_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2775E450g,
	"subcore_vm_fixed2800_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800E450g,
	"subcore_vm_fixed2850_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2850E450g,
	"subcore_vm_fixed2875_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2875E450g,
	"subcore_vm_fixed2900_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900E450g,
	"subcore_vm_fixed2925_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2925E450g,
	"subcore_vm_fixed2950_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2950E450g,
	"subcore_vm_fixed2975_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2975E450g,
	"subcore_vm_fixed3000_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000E450g,
	"subcore_vm_fixed3025_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3025E450g,
	"subcore_vm_fixed3050_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3050E450g,
	"subcore_vm_fixed3075_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3075E450g,
	"subcore_vm_fixed3100_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100E450g,
	"subcore_vm_fixed3125_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3125E450g,
	"subcore_vm_fixed3150_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150E450g,
	"subcore_vm_fixed3200_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200E450g,
	"subcore_vm_fixed3225_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3225E450g,
	"subcore_vm_fixed3250_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3250E450g,
	"subcore_vm_fixed3300_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300E450g,
	"subcore_vm_fixed3325_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3325E450g,
	"subcore_vm_fixed3375_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3375E450g,
	"subcore_vm_fixed3400_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400E450g,
	"subcore_vm_fixed3450_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3450E450g,
	"subcore_vm_fixed3500_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500E450g,
	"subcore_vm_fixed3525_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3525E450g,
	"subcore_vm_fixed3575_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3575E450g,
	"subcore_vm_fixed3600_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600E450g,
	"subcore_vm_fixed3625_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3625E450g,
	"subcore_vm_fixed3675_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3675E450g,
	"subcore_vm_fixed3700_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700E450g,
	"subcore_vm_fixed3750_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3750E450g,
	"subcore_vm_fixed3800_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800E450g,
	"subcore_vm_fixed3825_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3825E450g,
	"subcore_vm_fixed3850_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3850E450g,
	"subcore_vm_fixed3875_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3875E450g,
	"subcore_vm_fixed3900_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900E450g,
	"subcore_vm_fixed3975_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3975E450g,
	"subcore_vm_fixed4000_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000E450g,
	"subcore_vm_fixed4025_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4025E450g,
	"subcore_vm_fixed4050_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050E450g,
	"subcore_vm_fixed4100_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100E450g,
	"subcore_vm_fixed4125_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4125E450g,
	"subcore_vm_fixed4200_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200E450g,
	"subcore_vm_fixed4225_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4225E450g,
	"subcore_vm_fixed4250_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4250E450g,
	"subcore_vm_fixed4275_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4275E450g,
	"subcore_vm_fixed4300_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300E450g,
	"subcore_vm_fixed4350_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4350E450g,
	"subcore_vm_fixed4375_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4375E450g,
	"subcore_vm_fixed4400_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400E450g,
	"subcore_vm_fixed4425_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4425E450g,
	"subcore_vm_fixed4500_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500E450g,
	"subcore_vm_fixed4550_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4550E450g,
	"subcore_vm_fixed4575_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4575E450g,
	"subcore_vm_fixed4600_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600E450g,
	"subcore_vm_fixed4625_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4625E450g,
	"subcore_vm_fixed4650_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4650E450g,
	"subcore_vm_fixed4675_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4675E450g,
	"subcore_vm_fixed4700_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700E450g,
	"subcore_vm_fixed4725_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4725E450g,
	"subcore_vm_fixed4750_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4750E450g,
	"subcore_vm_fixed4800_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800E450g,
	"subcore_vm_fixed4875_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4875E450g,
	"subcore_vm_fixed4900_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900E450g,
	"subcore_vm_fixed4950_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950E450g,
	"subcore_vm_fixed5000_e4_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000E450g,
	"subcore_vm_fixed0020_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0020A150g,
	"subcore_vm_fixed0040_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0040A150g,
	"subcore_vm_fixed0060_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0060A150g,
	"subcore_vm_fixed0080_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0080A150g,
	"subcore_vm_fixed0100_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0100A150g,
	"subcore_vm_fixed0120_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0120A150g,
	"subcore_vm_fixed0140_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0140A150g,
	"subcore_vm_fixed0160_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0160A150g,
	"subcore_vm_fixed0180_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180A150g,
	"subcore_vm_fixed0200_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0200A150g,
	"subcore_vm_fixed0220_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0220A150g,
	"subcore_vm_fixed0240_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0240A150g,
	"subcore_vm_fixed0260_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0260A150g,
	"subcore_vm_fixed0280_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0280A150g,
	"subcore_vm_fixed0300_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0300A150g,
	"subcore_vm_fixed0320_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0320A150g,
	"subcore_vm_fixed0340_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0340A150g,
	"subcore_vm_fixed0360_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360A150g,
	"subcore_vm_fixed0380_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0380A150g,
	"subcore_vm_fixed0400_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0400A150g,
	"subcore_vm_fixed0420_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0420A150g,
	"subcore_vm_fixed0440_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0440A150g,
	"subcore_vm_fixed0460_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0460A150g,
	"subcore_vm_fixed0480_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0480A150g,
	"subcore_vm_fixed0500_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0500A150g,
	"subcore_vm_fixed0520_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0520A150g,
	"subcore_vm_fixed0540_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540A150g,
	"subcore_vm_fixed0560_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0560A150g,
	"subcore_vm_fixed0580_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0580A150g,
	"subcore_vm_fixed0600_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0600A150g,
	"subcore_vm_fixed0620_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0620A150g,
	"subcore_vm_fixed0640_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0640A150g,
	"subcore_vm_fixed0660_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0660A150g,
	"subcore_vm_fixed0680_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0680A150g,
	"subcore_vm_fixed0700_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0700A150g,
	"subcore_vm_fixed0720_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720A150g,
	"subcore_vm_fixed0740_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0740A150g,
	"subcore_vm_fixed0760_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0760A150g,
	"subcore_vm_fixed0780_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0780A150g,
	"subcore_vm_fixed0800_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0800A150g,
	"subcore_vm_fixed0820_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0820A150g,
	"subcore_vm_fixed0840_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0840A150g,
	"subcore_vm_fixed0860_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0860A150g,
	"subcore_vm_fixed0880_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0880A150g,
	"subcore_vm_fixed0900_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900A150g,
	"subcore_vm_fixed0920_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0920A150g,
	"subcore_vm_fixed0940_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0940A150g,
	"subcore_vm_fixed0960_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0960A150g,
	"subcore_vm_fixed0980_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0980A150g,
	"subcore_vm_fixed1000_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1000A150g,
	"subcore_vm_fixed1020_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1020A150g,
	"subcore_vm_fixed1040_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1040A150g,
	"subcore_vm_fixed1060_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1060A150g,
	"subcore_vm_fixed1080_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080A150g,
	"subcore_vm_fixed1100_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1100A150g,
	"subcore_vm_fixed1120_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1120A150g,
	"subcore_vm_fixed1140_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1140A150g,
	"subcore_vm_fixed1160_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1160A150g,
	"subcore_vm_fixed1180_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1180A150g,
	"subcore_vm_fixed1200_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1200A150g,
	"subcore_vm_fixed1220_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1220A150g,
	"subcore_vm_fixed1240_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1240A150g,
	"subcore_vm_fixed1260_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260A150g,
	"subcore_vm_fixed1280_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1280A150g,
	"subcore_vm_fixed1300_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1300A150g,
	"subcore_vm_fixed1320_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1320A150g,
	"subcore_vm_fixed1340_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1340A150g,
	"subcore_vm_fixed1360_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1360A150g,
	"subcore_vm_fixed1380_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1380A150g,
	"subcore_vm_fixed1400_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1400A150g,
	"subcore_vm_fixed1420_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1420A150g,
	"subcore_vm_fixed1440_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440A150g,
	"subcore_vm_fixed1460_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1460A150g,
	"subcore_vm_fixed1480_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1480A150g,
	"subcore_vm_fixed1500_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1500A150g,
	"subcore_vm_fixed1520_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1520A150g,
	"subcore_vm_fixed1540_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1540A150g,
	"subcore_vm_fixed1560_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1560A150g,
	"subcore_vm_fixed1580_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1580A150g,
	"subcore_vm_fixed1600_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1600A150g,
	"subcore_vm_fixed1620_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620A150g,
	"subcore_vm_fixed1640_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1640A150g,
	"subcore_vm_fixed1660_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1660A150g,
	"subcore_vm_fixed1680_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1680A150g,
	"subcore_vm_fixed1700_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1700A150g,
	"subcore_vm_fixed1720_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1720A150g,
	"subcore_vm_fixed1740_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1740A150g,
	"subcore_vm_fixed1760_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1760A150g,
	"subcore_vm_fixed1780_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1780A150g,
	"subcore_vm_fixed1800_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800A150g,
	"subcore_vm_fixed1820_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1820A150g,
	"subcore_vm_fixed1840_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1840A150g,
	"subcore_vm_fixed1860_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1860A150g,
	"subcore_vm_fixed1880_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1880A150g,
	"subcore_vm_fixed1900_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1900A150g,
	"subcore_vm_fixed1920_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1920A150g,
	"subcore_vm_fixed1940_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1940A150g,
	"subcore_vm_fixed1960_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1960A150g,
	"subcore_vm_fixed1980_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980A150g,
	"subcore_vm_fixed2000_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2000A150g,
	"subcore_vm_fixed2020_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2020A150g,
	"subcore_vm_fixed2040_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2040A150g,
	"subcore_vm_fixed2060_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2060A150g,
	"subcore_vm_fixed2080_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2080A150g,
	"subcore_vm_fixed2100_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2100A150g,
	"subcore_vm_fixed2120_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2120A150g,
	"subcore_vm_fixed2140_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2140A150g,
	"subcore_vm_fixed2160_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160A150g,
	"subcore_vm_fixed2180_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2180A150g,
	"subcore_vm_fixed2200_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2200A150g,
	"subcore_vm_fixed2220_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2220A150g,
	"subcore_vm_fixed2240_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2240A150g,
	"subcore_vm_fixed2260_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2260A150g,
	"subcore_vm_fixed2280_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2280A150g,
	"subcore_vm_fixed2300_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2300A150g,
	"subcore_vm_fixed2320_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2320A150g,
	"subcore_vm_fixed2340_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340A150g,
	"subcore_vm_fixed2360_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2360A150g,
	"subcore_vm_fixed2380_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2380A150g,
	"subcore_vm_fixed2400_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2400A150g,
	"subcore_vm_fixed2420_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2420A150g,
	"subcore_vm_fixed2440_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2440A150g,
	"subcore_vm_fixed2460_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2460A150g,
	"subcore_vm_fixed2480_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2480A150g,
	"subcore_vm_fixed2500_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2500A150g,
	"subcore_vm_fixed2520_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520A150g,
	"subcore_vm_fixed2540_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2540A150g,
	"subcore_vm_fixed2560_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2560A150g,
	"subcore_vm_fixed2580_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2580A150g,
	"subcore_vm_fixed2600_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2600A150g,
	"subcore_vm_fixed2620_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2620A150g,
	"subcore_vm_fixed2640_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2640A150g,
	"subcore_vm_fixed2660_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2660A150g,
	"subcore_vm_fixed2680_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2680A150g,
	"subcore_vm_fixed2700_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700A150g,
	"subcore_vm_fixed2720_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2720A150g,
	"subcore_vm_fixed2740_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2740A150g,
	"subcore_vm_fixed2760_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2760A150g,
	"subcore_vm_fixed2780_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2780A150g,
	"subcore_vm_fixed2800_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2800A150g,
	"subcore_vm_fixed2820_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2820A150g,
	"subcore_vm_fixed2840_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2840A150g,
	"subcore_vm_fixed2860_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2860A150g,
	"subcore_vm_fixed2880_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880A150g,
	"subcore_vm_fixed2900_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2900A150g,
	"subcore_vm_fixed2920_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2920A150g,
	"subcore_vm_fixed2940_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2940A150g,
	"subcore_vm_fixed2960_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2960A150g,
	"subcore_vm_fixed2980_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2980A150g,
	"subcore_vm_fixed3000_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3000A150g,
	"subcore_vm_fixed3020_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3020A150g,
	"subcore_vm_fixed3040_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3040A150g,
	"subcore_vm_fixed3060_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060A150g,
	"subcore_vm_fixed3080_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3080A150g,
	"subcore_vm_fixed3100_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3100A150g,
	"subcore_vm_fixed3120_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3120A150g,
	"subcore_vm_fixed3140_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3140A150g,
	"subcore_vm_fixed3160_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3160A150g,
	"subcore_vm_fixed3180_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3180A150g,
	"subcore_vm_fixed3200_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3200A150g,
	"subcore_vm_fixed3220_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3220A150g,
	"subcore_vm_fixed3240_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240A150g,
	"subcore_vm_fixed3260_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3260A150g,
	"subcore_vm_fixed3280_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3280A150g,
	"subcore_vm_fixed3300_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3300A150g,
	"subcore_vm_fixed3320_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3320A150g,
	"subcore_vm_fixed3340_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3340A150g,
	"subcore_vm_fixed3360_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3360A150g,
	"subcore_vm_fixed3380_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3380A150g,
	"subcore_vm_fixed3400_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3400A150g,
	"subcore_vm_fixed3420_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420A150g,
	"subcore_vm_fixed3440_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3440A150g,
	"subcore_vm_fixed3460_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3460A150g,
	"subcore_vm_fixed3480_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3480A150g,
	"subcore_vm_fixed3500_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3500A150g,
	"subcore_vm_fixed3520_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3520A150g,
	"subcore_vm_fixed3540_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3540A150g,
	"subcore_vm_fixed3560_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3560A150g,
	"subcore_vm_fixed3580_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3580A150g,
	"subcore_vm_fixed3600_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600A150g,
	"subcore_vm_fixed3620_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3620A150g,
	"subcore_vm_fixed3640_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3640A150g,
	"subcore_vm_fixed3660_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3660A150g,
	"subcore_vm_fixed3680_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3680A150g,
	"subcore_vm_fixed3700_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3700A150g,
	"subcore_vm_fixed3720_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3720A150g,
	"subcore_vm_fixed3740_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3740A150g,
	"subcore_vm_fixed3760_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3760A150g,
	"subcore_vm_fixed3780_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780A150g,
	"subcore_vm_fixed3800_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3800A150g,
	"subcore_vm_fixed3820_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3820A150g,
	"subcore_vm_fixed3840_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3840A150g,
	"subcore_vm_fixed3860_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3860A150g,
	"subcore_vm_fixed3880_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3880A150g,
	"subcore_vm_fixed3900_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3900A150g,
	"subcore_vm_fixed3920_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3920A150g,
	"subcore_vm_fixed3940_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3940A150g,
	"subcore_vm_fixed3960_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960A150g,
	"subcore_vm_fixed3980_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3980A150g,
	"subcore_vm_fixed4000_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4000A150g,
	"subcore_vm_fixed4020_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4020A150g,
	"subcore_vm_fixed4040_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4040A150g,
	"subcore_vm_fixed4060_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4060A150g,
	"subcore_vm_fixed4080_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4080A150g,
	"subcore_vm_fixed4100_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4100A150g,
	"subcore_vm_fixed4120_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4120A150g,
	"subcore_vm_fixed4140_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140A150g,
	"subcore_vm_fixed4160_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4160A150g,
	"subcore_vm_fixed4180_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4180A150g,
	"subcore_vm_fixed4200_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4200A150g,
	"subcore_vm_fixed4220_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4220A150g,
	"subcore_vm_fixed4240_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4240A150g,
	"subcore_vm_fixed4260_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4260A150g,
	"subcore_vm_fixed4280_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4280A150g,
	"subcore_vm_fixed4300_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4300A150g,
	"subcore_vm_fixed4320_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320A150g,
	"subcore_vm_fixed4340_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4340A150g,
	"subcore_vm_fixed4360_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4360A150g,
	"subcore_vm_fixed4380_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4380A150g,
	"subcore_vm_fixed4400_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4400A150g,
	"subcore_vm_fixed4420_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4420A150g,
	"subcore_vm_fixed4440_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4440A150g,
	"subcore_vm_fixed4460_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4460A150g,
	"subcore_vm_fixed4480_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4480A150g,
	"subcore_vm_fixed4500_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500A150g,
	"subcore_vm_fixed4520_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4520A150g,
	"subcore_vm_fixed4540_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4540A150g,
	"subcore_vm_fixed4560_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4560A150g,
	"subcore_vm_fixed4580_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4580A150g,
	"subcore_vm_fixed4600_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4600A150g,
	"subcore_vm_fixed4620_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4620A150g,
	"subcore_vm_fixed4640_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4640A150g,
	"subcore_vm_fixed4660_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4660A150g,
	"subcore_vm_fixed4680_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680A150g,
	"subcore_vm_fixed4700_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4700A150g,
	"subcore_vm_fixed4720_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4720A150g,
	"subcore_vm_fixed4740_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4740A150g,
	"subcore_vm_fixed4760_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4760A150g,
	"subcore_vm_fixed4780_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4780A150g,
	"subcore_vm_fixed4800_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4800A150g,
	"subcore_vm_fixed4820_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4820A150g,
	"subcore_vm_fixed4840_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4840A150g,
	"subcore_vm_fixed4860_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860A150g,
	"subcore_vm_fixed4880_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4880A150g,
	"subcore_vm_fixed4900_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4900A150g,
	"subcore_vm_fixed4920_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4920A150g,
	"subcore_vm_fixed4940_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4940A150g,
	"subcore_vm_fixed4960_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4960A150g,
	"subcore_vm_fixed4980_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4980A150g,
	"subcore_vm_fixed5000_a1_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed5000A150g,
	"subcore_vm_fixed0090_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0090X950g,
	"subcore_vm_fixed0180_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0180X950g,
	"subcore_vm_fixed0270_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0270X950g,
	"subcore_vm_fixed0360_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0360X950g,
	"subcore_vm_fixed0450_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0450X950g,
	"subcore_vm_fixed0540_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0540X950g,
	"subcore_vm_fixed0630_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0630X950g,
	"subcore_vm_fixed0720_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0720X950g,
	"subcore_vm_fixed0810_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0810X950g,
	"subcore_vm_fixed0900_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0900X950g,
	"subcore_vm_fixed0990_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed0990X950g,
	"subcore_vm_fixed1080_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1080X950g,
	"subcore_vm_fixed1170_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1170X950g,
	"subcore_vm_fixed1260_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1260X950g,
	"subcore_vm_fixed1350_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1350X950g,
	"subcore_vm_fixed1440_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1440X950g,
	"subcore_vm_fixed1530_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1530X950g,
	"subcore_vm_fixed1620_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1620X950g,
	"subcore_vm_fixed1710_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1710X950g,
	"subcore_vm_fixed1800_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1800X950g,
	"subcore_vm_fixed1890_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1890X950g,
	"subcore_vm_fixed1980_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed1980X950g,
	"subcore_vm_fixed2070_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2070X950g,
	"subcore_vm_fixed2160_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2160X950g,
	"subcore_vm_fixed2250_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2250X950g,
	"subcore_vm_fixed2340_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2340X950g,
	"subcore_vm_fixed2430_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2430X950g,
	"subcore_vm_fixed2520_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2520X950g,
	"subcore_vm_fixed2610_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2610X950g,
	"subcore_vm_fixed2700_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2700X950g,
	"subcore_vm_fixed2790_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2790X950g,
	"subcore_vm_fixed2880_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2880X950g,
	"subcore_vm_fixed2970_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed2970X950g,
	"subcore_vm_fixed3060_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3060X950g,
	"subcore_vm_fixed3150_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3150X950g,
	"subcore_vm_fixed3240_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3240X950g,
	"subcore_vm_fixed3330_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3330X950g,
	"subcore_vm_fixed3420_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3420X950g,
	"subcore_vm_fixed3510_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3510X950g,
	"subcore_vm_fixed3600_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3600X950g,
	"subcore_vm_fixed3690_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3690X950g,
	"subcore_vm_fixed3780_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3780X950g,
	"subcore_vm_fixed3870_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3870X950g,
	"subcore_vm_fixed3960_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed3960X950g,
	"subcore_vm_fixed4050_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4050X950g,
	"subcore_vm_fixed4140_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4140X950g,
	"subcore_vm_fixed4230_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4230X950g,
	"subcore_vm_fixed4320_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4320X950g,
	"subcore_vm_fixed4410_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4410X950g,
	"subcore_vm_fixed4500_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4500X950g,
	"subcore_vm_fixed4590_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4590X950g,
	"subcore_vm_fixed4680_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4680X950g,
	"subcore_vm_fixed4770_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4770X950g,
	"subcore_vm_fixed4860_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4860X950g,
	"subcore_vm_fixed4950_x9_50g":          CreateInternalVnicDetailsVnicShapeSubcoreVmFixed4950X950g,
	"dynamic_a1_50g":                       CreateInternalVnicDetailsVnicShapeDynamicA150g,
	"fixed0040_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0040A150g,
	"fixed0100_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0100A150g,
	"fixed0200_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0200A150g,
	"fixed0300_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0300A150g,
	"fixed0400_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0400A150g,
	"fixed0500_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0500A150g,
	"fixed0600_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0600A150g,
	"fixed0700_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0700A150g,
	"fixed0800_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0800A150g,
	"fixed0900_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed0900A150g,
	"fixed1000_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1000A150g,
	"fixed1100_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1100A150g,
	"fixed1200_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1200A150g,
	"fixed1300_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1300A150g,
	"fixed1400_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1400A150g,
	"fixed1500_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1500A150g,
	"fixed1600_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1600A150g,
	"fixed1700_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1700A150g,
	"fixed1800_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1800A150g,
	"fixed1900_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed1900A150g,
	"fixed2000_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2000A150g,
	"fixed2100_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2100A150g,
	"fixed2200_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2200A150g,
	"fixed2300_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2300A150g,
	"fixed2400_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2400A150g,
	"fixed2500_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2500A150g,
	"fixed2600_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2600A150g,
	"fixed2700_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2700A150g,
	"fixed2800_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2800A150g,
	"fixed2900_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed2900A150g,
	"fixed3000_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3000A150g,
	"fixed3100_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3100A150g,
	"fixed3200_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3200A150g,
	"fixed3300_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3300A150g,
	"fixed3400_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3400A150g,
	"fixed3500_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3500A150g,
	"fixed3600_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3600A150g,
	"fixed3700_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3700A150g,
	"fixed3800_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3800A150g,
	"fixed3900_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed3900A150g,
	"fixed4000_a1_50g":                     CreateInternalVnicDetailsVnicShapeFixed4000A150g,
	"entirehost_a1_50g":                    CreateInternalVnicDetailsVnicShapeEntirehostA150g,
	"dynamic_x9_50g":                       CreateInternalVnicDetailsVnicShapeDynamicX950g,
	"fixed0040_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed0040X950g,
	"fixed0400_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed0400X950g,
	"fixed0800_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed0800X950g,
	"fixed1200_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed1200X950g,
	"fixed1600_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed1600X950g,
	"fixed2000_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed2000X950g,
	"fixed2400_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed2400X950g,
	"fixed2800_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed2800X950g,
	"fixed3200_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed3200X950g,
	"fixed3600_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed3600X950g,
	"fixed4000_x9_50g":                     CreateInternalVnicDetailsVnicShapeFixed4000X950g,
	"standard_vm_fixed0100_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0100X950g,
	"standard_vm_fixed0200_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0200X950g,
	"standard_vm_fixed0300_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0300X950g,
	"standard_vm_fixed0400_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0400X950g,
	"standard_vm_fixed0500_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0500X950g,
	"standard_vm_fixed0600_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0600X950g,
	"standard_vm_fixed0700_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0700X950g,
	"standard_vm_fixed0800_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0800X950g,
	"standard_vm_fixed0900_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed0900X950g,
	"standard_vm_fixed1000_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1000X950g,
	"standard_vm_fixed1100_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1100X950g,
	"standard_vm_fixed1200_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1200X950g,
	"standard_vm_fixed1300_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1300X950g,
	"standard_vm_fixed1400_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1400X950g,
	"standard_vm_fixed1500_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1500X950g,
	"standard_vm_fixed1600_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1600X950g,
	"standard_vm_fixed1700_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1700X950g,
	"standard_vm_fixed1800_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1800X950g,
	"standard_vm_fixed1900_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed1900X950g,
	"standard_vm_fixed2000_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2000X950g,
	"standard_vm_fixed2100_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2100X950g,
	"standard_vm_fixed2200_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2200X950g,
	"standard_vm_fixed2300_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2300X950g,
	"standard_vm_fixed2400_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2400X950g,
	"standard_vm_fixed2500_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2500X950g,
	"standard_vm_fixed2600_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2600X950g,
	"standard_vm_fixed2700_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2700X950g,
	"standard_vm_fixed2800_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2800X950g,
	"standard_vm_fixed2900_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed2900X950g,
	"standard_vm_fixed3000_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3000X950g,
	"standard_vm_fixed3100_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3100X950g,
	"standard_vm_fixed3200_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3200X950g,
	"standard_vm_fixed3300_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3300X950g,
	"standard_vm_fixed3400_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3400X950g,
	"standard_vm_fixed3500_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3500X950g,
	"standard_vm_fixed3600_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3600X950g,
	"standard_vm_fixed3700_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3700X950g,
	"standard_vm_fixed3800_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3800X950g,
	"standard_vm_fixed3900_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed3900X950g,
	"standard_vm_fixed4000_x9_50g":         CreateInternalVnicDetailsVnicShapeStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed0025_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0025X950g,
	"subcore_standard_vm_fixed0050_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0050X950g,
	"subcore_standard_vm_fixed0075_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0075X950g,
	"subcore_standard_vm_fixed0100_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0100X950g,
	"subcore_standard_vm_fixed0125_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0125X950g,
	"subcore_standard_vm_fixed0150_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0150X950g,
	"subcore_standard_vm_fixed0175_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0175X950g,
	"subcore_standard_vm_fixed0200_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0200X950g,
	"subcore_standard_vm_fixed0225_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0225X950g,
	"subcore_standard_vm_fixed0250_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0250X950g,
	"subcore_standard_vm_fixed0275_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0275X950g,
	"subcore_standard_vm_fixed0300_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0300X950g,
	"subcore_standard_vm_fixed0325_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0325X950g,
	"subcore_standard_vm_fixed0350_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0350X950g,
	"subcore_standard_vm_fixed0375_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0375X950g,
	"subcore_standard_vm_fixed0400_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0400X950g,
	"subcore_standard_vm_fixed0425_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0425X950g,
	"subcore_standard_vm_fixed0450_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0450X950g,
	"subcore_standard_vm_fixed0475_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0475X950g,
	"subcore_standard_vm_fixed0500_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0500X950g,
	"subcore_standard_vm_fixed0525_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0525X950g,
	"subcore_standard_vm_fixed0550_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0550X950g,
	"subcore_standard_vm_fixed0575_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0575X950g,
	"subcore_standard_vm_fixed0600_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0600X950g,
	"subcore_standard_vm_fixed0625_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0625X950g,
	"subcore_standard_vm_fixed0650_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0650X950g,
	"subcore_standard_vm_fixed0675_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0675X950g,
	"subcore_standard_vm_fixed0700_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0700X950g,
	"subcore_standard_vm_fixed0725_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0725X950g,
	"subcore_standard_vm_fixed0750_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0750X950g,
	"subcore_standard_vm_fixed0775_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0775X950g,
	"subcore_standard_vm_fixed0800_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0800X950g,
	"subcore_standard_vm_fixed0825_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0825X950g,
	"subcore_standard_vm_fixed0850_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0850X950g,
	"subcore_standard_vm_fixed0875_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0875X950g,
	"subcore_standard_vm_fixed0900_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0900X950g,
	"subcore_standard_vm_fixed0925_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0925X950g,
	"subcore_standard_vm_fixed0950_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0950X950g,
	"subcore_standard_vm_fixed0975_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed0975X950g,
	"subcore_standard_vm_fixed1000_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1000X950g,
	"subcore_standard_vm_fixed1025_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1025X950g,
	"subcore_standard_vm_fixed1050_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1050X950g,
	"subcore_standard_vm_fixed1075_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1075X950g,
	"subcore_standard_vm_fixed1100_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1100X950g,
	"subcore_standard_vm_fixed1125_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1125X950g,
	"subcore_standard_vm_fixed1150_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1150X950g,
	"subcore_standard_vm_fixed1175_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1175X950g,
	"subcore_standard_vm_fixed1200_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1200X950g,
	"subcore_standard_vm_fixed1225_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1225X950g,
	"subcore_standard_vm_fixed1250_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1250X950g,
	"subcore_standard_vm_fixed1275_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1275X950g,
	"subcore_standard_vm_fixed1300_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1300X950g,
	"subcore_standard_vm_fixed1325_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1325X950g,
	"subcore_standard_vm_fixed1350_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1350X950g,
	"subcore_standard_vm_fixed1375_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1375X950g,
	"subcore_standard_vm_fixed1400_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1400X950g,
	"subcore_standard_vm_fixed1425_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1425X950g,
	"subcore_standard_vm_fixed1450_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1450X950g,
	"subcore_standard_vm_fixed1475_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1475X950g,
	"subcore_standard_vm_fixed1500_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1500X950g,
	"subcore_standard_vm_fixed1525_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1525X950g,
	"subcore_standard_vm_fixed1550_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1550X950g,
	"subcore_standard_vm_fixed1575_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1575X950g,
	"subcore_standard_vm_fixed1600_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1600X950g,
	"subcore_standard_vm_fixed1625_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1625X950g,
	"subcore_standard_vm_fixed1650_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1650X950g,
	"subcore_standard_vm_fixed1700_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1700X950g,
	"subcore_standard_vm_fixed1725_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1725X950g,
	"subcore_standard_vm_fixed1750_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1750X950g,
	"subcore_standard_vm_fixed1800_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1800X950g,
	"subcore_standard_vm_fixed1850_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1850X950g,
	"subcore_standard_vm_fixed1875_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1875X950g,
	"subcore_standard_vm_fixed1900_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1900X950g,
	"subcore_standard_vm_fixed1925_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1925X950g,
	"subcore_standard_vm_fixed1950_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed1950X950g,
	"subcore_standard_vm_fixed2000_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2000X950g,
	"subcore_standard_vm_fixed2025_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2025X950g,
	"subcore_standard_vm_fixed2050_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2050X950g,
	"subcore_standard_vm_fixed2100_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2100X950g,
	"subcore_standard_vm_fixed2125_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2125X950g,
	"subcore_standard_vm_fixed2150_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2150X950g,
	"subcore_standard_vm_fixed2175_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2175X950g,
	"subcore_standard_vm_fixed2200_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2200X950g,
	"subcore_standard_vm_fixed2250_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2250X950g,
	"subcore_standard_vm_fixed2275_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2275X950g,
	"subcore_standard_vm_fixed2300_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2300X950g,
	"subcore_standard_vm_fixed2325_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2325X950g,
	"subcore_standard_vm_fixed2350_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2350X950g,
	"subcore_standard_vm_fixed2375_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2375X950g,
	"subcore_standard_vm_fixed2400_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2400X950g,
	"subcore_standard_vm_fixed2450_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2450X950g,
	"subcore_standard_vm_fixed2475_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2475X950g,
	"subcore_standard_vm_fixed2500_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2500X950g,
	"subcore_standard_vm_fixed2550_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2550X950g,
	"subcore_standard_vm_fixed2600_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2600X950g,
	"subcore_standard_vm_fixed2625_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2625X950g,
	"subcore_standard_vm_fixed2650_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2650X950g,
	"subcore_standard_vm_fixed2700_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2700X950g,
	"subcore_standard_vm_fixed2750_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2750X950g,
	"subcore_standard_vm_fixed2775_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2775X950g,
	"subcore_standard_vm_fixed2800_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2800X950g,
	"subcore_standard_vm_fixed2850_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2850X950g,
	"subcore_standard_vm_fixed2875_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2875X950g,
	"subcore_standard_vm_fixed2900_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2900X950g,
	"subcore_standard_vm_fixed2925_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2925X950g,
	"subcore_standard_vm_fixed2950_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2950X950g,
	"subcore_standard_vm_fixed2975_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed2975X950g,
	"subcore_standard_vm_fixed3000_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3000X950g,
	"subcore_standard_vm_fixed3025_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3025X950g,
	"subcore_standard_vm_fixed3050_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3050X950g,
	"subcore_standard_vm_fixed3075_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3075X950g,
	"subcore_standard_vm_fixed3100_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3100X950g,
	"subcore_standard_vm_fixed3125_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3125X950g,
	"subcore_standard_vm_fixed3150_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3150X950g,
	"subcore_standard_vm_fixed3200_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3200X950g,
	"subcore_standard_vm_fixed3225_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3225X950g,
	"subcore_standard_vm_fixed3250_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3250X950g,
	"subcore_standard_vm_fixed3300_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3300X950g,
	"subcore_standard_vm_fixed3325_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3325X950g,
	"subcore_standard_vm_fixed3375_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3375X950g,
	"subcore_standard_vm_fixed3400_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3400X950g,
	"subcore_standard_vm_fixed3450_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3450X950g,
	"subcore_standard_vm_fixed3500_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3500X950g,
	"subcore_standard_vm_fixed3525_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3525X950g,
	"subcore_standard_vm_fixed3575_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3575X950g,
	"subcore_standard_vm_fixed3600_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3600X950g,
	"subcore_standard_vm_fixed3625_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3625X950g,
	"subcore_standard_vm_fixed3675_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3675X950g,
	"subcore_standard_vm_fixed3700_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3700X950g,
	"subcore_standard_vm_fixed3750_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3750X950g,
	"subcore_standard_vm_fixed3800_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3800X950g,
	"subcore_standard_vm_fixed3825_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3825X950g,
	"subcore_standard_vm_fixed3850_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3850X950g,
	"subcore_standard_vm_fixed3875_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3875X950g,
	"subcore_standard_vm_fixed3900_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3900X950g,
	"subcore_standard_vm_fixed3975_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed3975X950g,
	"subcore_standard_vm_fixed4000_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4000X950g,
	"subcore_standard_vm_fixed4025_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4025X950g,
	"subcore_standard_vm_fixed4050_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4050X950g,
	"subcore_standard_vm_fixed4100_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4100X950g,
	"subcore_standard_vm_fixed4125_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4125X950g,
	"subcore_standard_vm_fixed4200_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4200X950g,
	"subcore_standard_vm_fixed4225_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4225X950g,
	"subcore_standard_vm_fixed4250_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4250X950g,
	"subcore_standard_vm_fixed4275_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4275X950g,
	"subcore_standard_vm_fixed4300_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4300X950g,
	"subcore_standard_vm_fixed4350_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4350X950g,
	"subcore_standard_vm_fixed4375_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4375X950g,
	"subcore_standard_vm_fixed4400_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4400X950g,
	"subcore_standard_vm_fixed4425_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4425X950g,
	"subcore_standard_vm_fixed4500_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4500X950g,
	"subcore_standard_vm_fixed4550_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4550X950g,
	"subcore_standard_vm_fixed4575_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4575X950g,
	"subcore_standard_vm_fixed4600_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4600X950g,
	"subcore_standard_vm_fixed4625_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4625X950g,
	"subcore_standard_vm_fixed4650_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4650X950g,
	"subcore_standard_vm_fixed4675_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4675X950g,
	"subcore_standard_vm_fixed4700_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4700X950g,
	"subcore_standard_vm_fixed4725_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4725X950g,
	"subcore_standard_vm_fixed4750_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4750X950g,
	"subcore_standard_vm_fixed4800_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4800X950g,
	"subcore_standard_vm_fixed4875_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4875X950g,
	"subcore_standard_vm_fixed4900_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4900X950g,
	"subcore_standard_vm_fixed4950_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed4950X950g,
	"subcore_standard_vm_fixed5000_x9_50g": CreateInternalVnicDetailsVnicShapeSubcoreStandardVmFixed5000X950g,
	"entirehost_x9_50g":                    CreateInternalVnicDetailsVnicShapeEntirehostX950g,
}

// GetCreateInternalVnicDetailsVnicShapeEnumValues Enumerates the set of values for CreateInternalVnicDetailsVnicShapeEnum
func GetCreateInternalVnicDetailsVnicShapeEnumValues() []CreateInternalVnicDetailsVnicShapeEnum {
	values := make([]CreateInternalVnicDetailsVnicShapeEnum, 0)
	for _, v := range mappingCreateInternalVnicDetailsVnicShapeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalVnicDetailsVnicShapeEnumStringValues Enumerates the set of values in String for CreateInternalVnicDetailsVnicShapeEnum
func GetCreateInternalVnicDetailsVnicShapeEnumStringValues() []string {
	return []string{
		"DYNAMIC",
		"FIXED0040",
		"FIXED0060",
		"FIXED0060_PSM",
		"FIXED0100",
		"FIXED0120",
		"FIXED0120_2X",
		"FIXED0200",
		"FIXED0240",
		"FIXED0480",
		"ENTIREHOST",
		"DYNAMIC_25G",
		"FIXED0040_25G",
		"FIXED0100_25G",
		"FIXED0200_25G",
		"FIXED0400_25G",
		"FIXED0800_25G",
		"FIXED1600_25G",
		"FIXED2400_25G",
		"ENTIREHOST_25G",
		"DYNAMIC_E1_25G",
		"FIXED0040_E1_25G",
		"FIXED0070_E1_25G",
		"FIXED0140_E1_25G",
		"FIXED0280_E1_25G",
		"FIXED0560_E1_25G",
		"FIXED1120_E1_25G",
		"FIXED1680_E1_25G",
		"ENTIREHOST_E1_25G",
		"DYNAMIC_B1_25G",
		"FIXED0040_B1_25G",
		"FIXED0060_B1_25G",
		"FIXED0120_B1_25G",
		"FIXED0240_B1_25G",
		"FIXED0480_B1_25G",
		"FIXED0960_B1_25G",
		"ENTIREHOST_B1_25G",
		"MICRO_VM_FIXED0048_E1_25G",
		"MICRO_LB_FIXED0001_E1_25G",
		"VNICAAS_FIXED0200",
		"VNICAAS_FIXED0400",
		"VNICAAS_FIXED0700",
		"VNICAAS_NLB_APPROVED_10G",
		"VNICAAS_NLB_APPROVED_25G",
		"VNICAAS_TELESIS_25G",
		"VNICAAS_TELESIS_10G",
		"VNICAAS_AMBASSADOR_FIXED0100",
		"VNICAAS_TELESIS_GAMMA",
		"VNICAAS_PRIVATEDNS",
		"VNICAAS_FWAAS",
		"VNICAAS_LBAAS_FREE",
		"VNICAAS_LBAAS_8G_512K",
		"VNICAAS_LBAAS_8G_1M",
		"VNICAAS_LBAAS_8G_2M",
		"VNICAAS_LBAAS_8G_3M",
		"DYNAMIC_E3_50G",
		"FIXED0040_E3_50G",
		"FIXED0100_E3_50G",
		"FIXED0200_E3_50G",
		"FIXED0300_E3_50G",
		"FIXED0400_E3_50G",
		"FIXED0500_E3_50G",
		"FIXED0600_E3_50G",
		"FIXED0700_E3_50G",
		"FIXED0800_E3_50G",
		"FIXED0900_E3_50G",
		"FIXED1000_E3_50G",
		"FIXED1100_E3_50G",
		"FIXED1200_E3_50G",
		"FIXED1300_E3_50G",
		"FIXED1400_E3_50G",
		"FIXED1500_E3_50G",
		"FIXED1600_E3_50G",
		"FIXED1700_E3_50G",
		"FIXED1800_E3_50G",
		"FIXED1900_E3_50G",
		"FIXED2000_E3_50G",
		"FIXED2100_E3_50G",
		"FIXED2200_E3_50G",
		"FIXED2300_E3_50G",
		"FIXED2400_E3_50G",
		"FIXED2500_E3_50G",
		"FIXED2600_E3_50G",
		"FIXED2700_E3_50G",
		"FIXED2800_E3_50G",
		"FIXED2900_E3_50G",
		"FIXED3000_E3_50G",
		"FIXED3100_E3_50G",
		"FIXED3200_E3_50G",
		"FIXED3300_E3_50G",
		"FIXED3400_E3_50G",
		"FIXED3500_E3_50G",
		"FIXED3600_E3_50G",
		"FIXED3700_E3_50G",
		"FIXED3800_E3_50G",
		"FIXED3900_E3_50G",
		"FIXED4000_E3_50G",
		"ENTIREHOST_E3_50G",
		"DYNAMIC_E4_50G",
		"FIXED0040_E4_50G",
		"FIXED0100_E4_50G",
		"FIXED0200_E4_50G",
		"FIXED0300_E4_50G",
		"FIXED0400_E4_50G",
		"FIXED0500_E4_50G",
		"FIXED0600_E4_50G",
		"FIXED0700_E4_50G",
		"FIXED0800_E4_50G",
		"FIXED0900_E4_50G",
		"FIXED1000_E4_50G",
		"FIXED1100_E4_50G",
		"FIXED1200_E4_50G",
		"FIXED1300_E4_50G",
		"FIXED1400_E4_50G",
		"FIXED1500_E4_50G",
		"FIXED1600_E4_50G",
		"FIXED1700_E4_50G",
		"FIXED1800_E4_50G",
		"FIXED1900_E4_50G",
		"FIXED2000_E4_50G",
		"FIXED2100_E4_50G",
		"FIXED2200_E4_50G",
		"FIXED2300_E4_50G",
		"FIXED2400_E4_50G",
		"FIXED2500_E4_50G",
		"FIXED2600_E4_50G",
		"FIXED2700_E4_50G",
		"FIXED2800_E4_50G",
		"FIXED2900_E4_50G",
		"FIXED3000_E4_50G",
		"FIXED3100_E4_50G",
		"FIXED3200_E4_50G",
		"FIXED3300_E4_50G",
		"FIXED3400_E4_50G",
		"FIXED3500_E4_50G",
		"FIXED3600_E4_50G",
		"FIXED3700_E4_50G",
		"FIXED3800_E4_50G",
		"FIXED3900_E4_50G",
		"FIXED4000_E4_50G",
		"ENTIREHOST_E4_50G",
		"MICRO_VM_FIXED0050_E3_50G",
		"MICRO_VM_FIXED0050_E4_50G",
		"SUBCORE_VM_FIXED0025_E3_50G",
		"SUBCORE_VM_FIXED0050_E3_50G",
		"SUBCORE_VM_FIXED0075_E3_50G",
		"SUBCORE_VM_FIXED0100_E3_50G",
		"SUBCORE_VM_FIXED0125_E3_50G",
		"SUBCORE_VM_FIXED0150_E3_50G",
		"SUBCORE_VM_FIXED0175_E3_50G",
		"SUBCORE_VM_FIXED0200_E3_50G",
		"SUBCORE_VM_FIXED0225_E3_50G",
		"SUBCORE_VM_FIXED0250_E3_50G",
		"SUBCORE_VM_FIXED0275_E3_50G",
		"SUBCORE_VM_FIXED0300_E3_50G",
		"SUBCORE_VM_FIXED0325_E3_50G",
		"SUBCORE_VM_FIXED0350_E3_50G",
		"SUBCORE_VM_FIXED0375_E3_50G",
		"SUBCORE_VM_FIXED0400_E3_50G",
		"SUBCORE_VM_FIXED0425_E3_50G",
		"SUBCORE_VM_FIXED0450_E3_50G",
		"SUBCORE_VM_FIXED0475_E3_50G",
		"SUBCORE_VM_FIXED0500_E3_50G",
		"SUBCORE_VM_FIXED0525_E3_50G",
		"SUBCORE_VM_FIXED0550_E3_50G",
		"SUBCORE_VM_FIXED0575_E3_50G",
		"SUBCORE_VM_FIXED0600_E3_50G",
		"SUBCORE_VM_FIXED0625_E3_50G",
		"SUBCORE_VM_FIXED0650_E3_50G",
		"SUBCORE_VM_FIXED0675_E3_50G",
		"SUBCORE_VM_FIXED0700_E3_50G",
		"SUBCORE_VM_FIXED0725_E3_50G",
		"SUBCORE_VM_FIXED0750_E3_50G",
		"SUBCORE_VM_FIXED0775_E3_50G",
		"SUBCORE_VM_FIXED0800_E3_50G",
		"SUBCORE_VM_FIXED0825_E3_50G",
		"SUBCORE_VM_FIXED0850_E3_50G",
		"SUBCORE_VM_FIXED0875_E3_50G",
		"SUBCORE_VM_FIXED0900_E3_50G",
		"SUBCORE_VM_FIXED0925_E3_50G",
		"SUBCORE_VM_FIXED0950_E3_50G",
		"SUBCORE_VM_FIXED0975_E3_50G",
		"SUBCORE_VM_FIXED1000_E3_50G",
		"SUBCORE_VM_FIXED1025_E3_50G",
		"SUBCORE_VM_FIXED1050_E3_50G",
		"SUBCORE_VM_FIXED1075_E3_50G",
		"SUBCORE_VM_FIXED1100_E3_50G",
		"SUBCORE_VM_FIXED1125_E3_50G",
		"SUBCORE_VM_FIXED1150_E3_50G",
		"SUBCORE_VM_FIXED1175_E3_50G",
		"SUBCORE_VM_FIXED1200_E3_50G",
		"SUBCORE_VM_FIXED1225_E3_50G",
		"SUBCORE_VM_FIXED1250_E3_50G",
		"SUBCORE_VM_FIXED1275_E3_50G",
		"SUBCORE_VM_FIXED1300_E3_50G",
		"SUBCORE_VM_FIXED1325_E3_50G",
		"SUBCORE_VM_FIXED1350_E3_50G",
		"SUBCORE_VM_FIXED1375_E3_50G",
		"SUBCORE_VM_FIXED1400_E3_50G",
		"SUBCORE_VM_FIXED1425_E3_50G",
		"SUBCORE_VM_FIXED1450_E3_50G",
		"SUBCORE_VM_FIXED1475_E3_50G",
		"SUBCORE_VM_FIXED1500_E3_50G",
		"SUBCORE_VM_FIXED1525_E3_50G",
		"SUBCORE_VM_FIXED1550_E3_50G",
		"SUBCORE_VM_FIXED1575_E3_50G",
		"SUBCORE_VM_FIXED1600_E3_50G",
		"SUBCORE_VM_FIXED1625_E3_50G",
		"SUBCORE_VM_FIXED1650_E3_50G",
		"SUBCORE_VM_FIXED1700_E3_50G",
		"SUBCORE_VM_FIXED1725_E3_50G",
		"SUBCORE_VM_FIXED1750_E3_50G",
		"SUBCORE_VM_FIXED1800_E3_50G",
		"SUBCORE_VM_FIXED1850_E3_50G",
		"SUBCORE_VM_FIXED1875_E3_50G",
		"SUBCORE_VM_FIXED1900_E3_50G",
		"SUBCORE_VM_FIXED1925_E3_50G",
		"SUBCORE_VM_FIXED1950_E3_50G",
		"SUBCORE_VM_FIXED2000_E3_50G",
		"SUBCORE_VM_FIXED2025_E3_50G",
		"SUBCORE_VM_FIXED2050_E3_50G",
		"SUBCORE_VM_FIXED2100_E3_50G",
		"SUBCORE_VM_FIXED2125_E3_50G",
		"SUBCORE_VM_FIXED2150_E3_50G",
		"SUBCORE_VM_FIXED2175_E3_50G",
		"SUBCORE_VM_FIXED2200_E3_50G",
		"SUBCORE_VM_FIXED2250_E3_50G",
		"SUBCORE_VM_FIXED2275_E3_50G",
		"SUBCORE_VM_FIXED2300_E3_50G",
		"SUBCORE_VM_FIXED2325_E3_50G",
		"SUBCORE_VM_FIXED2350_E3_50G",
		"SUBCORE_VM_FIXED2375_E3_50G",
		"SUBCORE_VM_FIXED2400_E3_50G",
		"SUBCORE_VM_FIXED2450_E3_50G",
		"SUBCORE_VM_FIXED2475_E3_50G",
		"SUBCORE_VM_FIXED2500_E3_50G",
		"SUBCORE_VM_FIXED2550_E3_50G",
		"SUBCORE_VM_FIXED2600_E3_50G",
		"SUBCORE_VM_FIXED2625_E3_50G",
		"SUBCORE_VM_FIXED2650_E3_50G",
		"SUBCORE_VM_FIXED2700_E3_50G",
		"SUBCORE_VM_FIXED2750_E3_50G",
		"SUBCORE_VM_FIXED2775_E3_50G",
		"SUBCORE_VM_FIXED2800_E3_50G",
		"SUBCORE_VM_FIXED2850_E3_50G",
		"SUBCORE_VM_FIXED2875_E3_50G",
		"SUBCORE_VM_FIXED2900_E3_50G",
		"SUBCORE_VM_FIXED2925_E3_50G",
		"SUBCORE_VM_FIXED2950_E3_50G",
		"SUBCORE_VM_FIXED2975_E3_50G",
		"SUBCORE_VM_FIXED3000_E3_50G",
		"SUBCORE_VM_FIXED3025_E3_50G",
		"SUBCORE_VM_FIXED3050_E3_50G",
		"SUBCORE_VM_FIXED3075_E3_50G",
		"SUBCORE_VM_FIXED3100_E3_50G",
		"SUBCORE_VM_FIXED3125_E3_50G",
		"SUBCORE_VM_FIXED3150_E3_50G",
		"SUBCORE_VM_FIXED3200_E3_50G",
		"SUBCORE_VM_FIXED3225_E3_50G",
		"SUBCORE_VM_FIXED3250_E3_50G",
		"SUBCORE_VM_FIXED3300_E3_50G",
		"SUBCORE_VM_FIXED3325_E3_50G",
		"SUBCORE_VM_FIXED3375_E3_50G",
		"SUBCORE_VM_FIXED3400_E3_50G",
		"SUBCORE_VM_FIXED3450_E3_50G",
		"SUBCORE_VM_FIXED3500_E3_50G",
		"SUBCORE_VM_FIXED3525_E3_50G",
		"SUBCORE_VM_FIXED3575_E3_50G",
		"SUBCORE_VM_FIXED3600_E3_50G",
		"SUBCORE_VM_FIXED3625_E3_50G",
		"SUBCORE_VM_FIXED3675_E3_50G",
		"SUBCORE_VM_FIXED3700_E3_50G",
		"SUBCORE_VM_FIXED3750_E3_50G",
		"SUBCORE_VM_FIXED3800_E3_50G",
		"SUBCORE_VM_FIXED3825_E3_50G",
		"SUBCORE_VM_FIXED3850_E3_50G",
		"SUBCORE_VM_FIXED3875_E3_50G",
		"SUBCORE_VM_FIXED3900_E3_50G",
		"SUBCORE_VM_FIXED3975_E3_50G",
		"SUBCORE_VM_FIXED4000_E3_50G",
		"SUBCORE_VM_FIXED4025_E3_50G",
		"SUBCORE_VM_FIXED4050_E3_50G",
		"SUBCORE_VM_FIXED4100_E3_50G",
		"SUBCORE_VM_FIXED4125_E3_50G",
		"SUBCORE_VM_FIXED4200_E3_50G",
		"SUBCORE_VM_FIXED4225_E3_50G",
		"SUBCORE_VM_FIXED4250_E3_50G",
		"SUBCORE_VM_FIXED4275_E3_50G",
		"SUBCORE_VM_FIXED4300_E3_50G",
		"SUBCORE_VM_FIXED4350_E3_50G",
		"SUBCORE_VM_FIXED4375_E3_50G",
		"SUBCORE_VM_FIXED4400_E3_50G",
		"SUBCORE_VM_FIXED4425_E3_50G",
		"SUBCORE_VM_FIXED4500_E3_50G",
		"SUBCORE_VM_FIXED4550_E3_50G",
		"SUBCORE_VM_FIXED4575_E3_50G",
		"SUBCORE_VM_FIXED4600_E3_50G",
		"SUBCORE_VM_FIXED4625_E3_50G",
		"SUBCORE_VM_FIXED4650_E3_50G",
		"SUBCORE_VM_FIXED4675_E3_50G",
		"SUBCORE_VM_FIXED4700_E3_50G",
		"SUBCORE_VM_FIXED4725_E3_50G",
		"SUBCORE_VM_FIXED4750_E3_50G",
		"SUBCORE_VM_FIXED4800_E3_50G",
		"SUBCORE_VM_FIXED4875_E3_50G",
		"SUBCORE_VM_FIXED4900_E3_50G",
		"SUBCORE_VM_FIXED4950_E3_50G",
		"SUBCORE_VM_FIXED5000_E3_50G",
		"SUBCORE_VM_FIXED0025_E4_50G",
		"SUBCORE_VM_FIXED0050_E4_50G",
		"SUBCORE_VM_FIXED0075_E4_50G",
		"SUBCORE_VM_FIXED0100_E4_50G",
		"SUBCORE_VM_FIXED0125_E4_50G",
		"SUBCORE_VM_FIXED0150_E4_50G",
		"SUBCORE_VM_FIXED0175_E4_50G",
		"SUBCORE_VM_FIXED0200_E4_50G",
		"SUBCORE_VM_FIXED0225_E4_50G",
		"SUBCORE_VM_FIXED0250_E4_50G",
		"SUBCORE_VM_FIXED0275_E4_50G",
		"SUBCORE_VM_FIXED0300_E4_50G",
		"SUBCORE_VM_FIXED0325_E4_50G",
		"SUBCORE_VM_FIXED0350_E4_50G",
		"SUBCORE_VM_FIXED0375_E4_50G",
		"SUBCORE_VM_FIXED0400_E4_50G",
		"SUBCORE_VM_FIXED0425_E4_50G",
		"SUBCORE_VM_FIXED0450_E4_50G",
		"SUBCORE_VM_FIXED0475_E4_50G",
		"SUBCORE_VM_FIXED0500_E4_50G",
		"SUBCORE_VM_FIXED0525_E4_50G",
		"SUBCORE_VM_FIXED0550_E4_50G",
		"SUBCORE_VM_FIXED0575_E4_50G",
		"SUBCORE_VM_FIXED0600_E4_50G",
		"SUBCORE_VM_FIXED0625_E4_50G",
		"SUBCORE_VM_FIXED0650_E4_50G",
		"SUBCORE_VM_FIXED0675_E4_50G",
		"SUBCORE_VM_FIXED0700_E4_50G",
		"SUBCORE_VM_FIXED0725_E4_50G",
		"SUBCORE_VM_FIXED0750_E4_50G",
		"SUBCORE_VM_FIXED0775_E4_50G",
		"SUBCORE_VM_FIXED0800_E4_50G",
		"SUBCORE_VM_FIXED0825_E4_50G",
		"SUBCORE_VM_FIXED0850_E4_50G",
		"SUBCORE_VM_FIXED0875_E4_50G",
		"SUBCORE_VM_FIXED0900_E4_50G",
		"SUBCORE_VM_FIXED0925_E4_50G",
		"SUBCORE_VM_FIXED0950_E4_50G",
		"SUBCORE_VM_FIXED0975_E4_50G",
		"SUBCORE_VM_FIXED1000_E4_50G",
		"SUBCORE_VM_FIXED1025_E4_50G",
		"SUBCORE_VM_FIXED1050_E4_50G",
		"SUBCORE_VM_FIXED1075_E4_50G",
		"SUBCORE_VM_FIXED1100_E4_50G",
		"SUBCORE_VM_FIXED1125_E4_50G",
		"SUBCORE_VM_FIXED1150_E4_50G",
		"SUBCORE_VM_FIXED1175_E4_50G",
		"SUBCORE_VM_FIXED1200_E4_50G",
		"SUBCORE_VM_FIXED1225_E4_50G",
		"SUBCORE_VM_FIXED1250_E4_50G",
		"SUBCORE_VM_FIXED1275_E4_50G",
		"SUBCORE_VM_FIXED1300_E4_50G",
		"SUBCORE_VM_FIXED1325_E4_50G",
		"SUBCORE_VM_FIXED1350_E4_50G",
		"SUBCORE_VM_FIXED1375_E4_50G",
		"SUBCORE_VM_FIXED1400_E4_50G",
		"SUBCORE_VM_FIXED1425_E4_50G",
		"SUBCORE_VM_FIXED1450_E4_50G",
		"SUBCORE_VM_FIXED1475_E4_50G",
		"SUBCORE_VM_FIXED1500_E4_50G",
		"SUBCORE_VM_FIXED1525_E4_50G",
		"SUBCORE_VM_FIXED1550_E4_50G",
		"SUBCORE_VM_FIXED1575_E4_50G",
		"SUBCORE_VM_FIXED1600_E4_50G",
		"SUBCORE_VM_FIXED1625_E4_50G",
		"SUBCORE_VM_FIXED1650_E4_50G",
		"SUBCORE_VM_FIXED1700_E4_50G",
		"SUBCORE_VM_FIXED1725_E4_50G",
		"SUBCORE_VM_FIXED1750_E4_50G",
		"SUBCORE_VM_FIXED1800_E4_50G",
		"SUBCORE_VM_FIXED1850_E4_50G",
		"SUBCORE_VM_FIXED1875_E4_50G",
		"SUBCORE_VM_FIXED1900_E4_50G",
		"SUBCORE_VM_FIXED1925_E4_50G",
		"SUBCORE_VM_FIXED1950_E4_50G",
		"SUBCORE_VM_FIXED2000_E4_50G",
		"SUBCORE_VM_FIXED2025_E4_50G",
		"SUBCORE_VM_FIXED2050_E4_50G",
		"SUBCORE_VM_FIXED2100_E4_50G",
		"SUBCORE_VM_FIXED2125_E4_50G",
		"SUBCORE_VM_FIXED2150_E4_50G",
		"SUBCORE_VM_FIXED2175_E4_50G",
		"SUBCORE_VM_FIXED2200_E4_50G",
		"SUBCORE_VM_FIXED2250_E4_50G",
		"SUBCORE_VM_FIXED2275_E4_50G",
		"SUBCORE_VM_FIXED2300_E4_50G",
		"SUBCORE_VM_FIXED2325_E4_50G",
		"SUBCORE_VM_FIXED2350_E4_50G",
		"SUBCORE_VM_FIXED2375_E4_50G",
		"SUBCORE_VM_FIXED2400_E4_50G",
		"SUBCORE_VM_FIXED2450_E4_50G",
		"SUBCORE_VM_FIXED2475_E4_50G",
		"SUBCORE_VM_FIXED2500_E4_50G",
		"SUBCORE_VM_FIXED2550_E4_50G",
		"SUBCORE_VM_FIXED2600_E4_50G",
		"SUBCORE_VM_FIXED2625_E4_50G",
		"SUBCORE_VM_FIXED2650_E4_50G",
		"SUBCORE_VM_FIXED2700_E4_50G",
		"SUBCORE_VM_FIXED2750_E4_50G",
		"SUBCORE_VM_FIXED2775_E4_50G",
		"SUBCORE_VM_FIXED2800_E4_50G",
		"SUBCORE_VM_FIXED2850_E4_50G",
		"SUBCORE_VM_FIXED2875_E4_50G",
		"SUBCORE_VM_FIXED2900_E4_50G",
		"SUBCORE_VM_FIXED2925_E4_50G",
		"SUBCORE_VM_FIXED2950_E4_50G",
		"SUBCORE_VM_FIXED2975_E4_50G",
		"SUBCORE_VM_FIXED3000_E4_50G",
		"SUBCORE_VM_FIXED3025_E4_50G",
		"SUBCORE_VM_FIXED3050_E4_50G",
		"SUBCORE_VM_FIXED3075_E4_50G",
		"SUBCORE_VM_FIXED3100_E4_50G",
		"SUBCORE_VM_FIXED3125_E4_50G",
		"SUBCORE_VM_FIXED3150_E4_50G",
		"SUBCORE_VM_FIXED3200_E4_50G",
		"SUBCORE_VM_FIXED3225_E4_50G",
		"SUBCORE_VM_FIXED3250_E4_50G",
		"SUBCORE_VM_FIXED3300_E4_50G",
		"SUBCORE_VM_FIXED3325_E4_50G",
		"SUBCORE_VM_FIXED3375_E4_50G",
		"SUBCORE_VM_FIXED3400_E4_50G",
		"SUBCORE_VM_FIXED3450_E4_50G",
		"SUBCORE_VM_FIXED3500_E4_50G",
		"SUBCORE_VM_FIXED3525_E4_50G",
		"SUBCORE_VM_FIXED3575_E4_50G",
		"SUBCORE_VM_FIXED3600_E4_50G",
		"SUBCORE_VM_FIXED3625_E4_50G",
		"SUBCORE_VM_FIXED3675_E4_50G",
		"SUBCORE_VM_FIXED3700_E4_50G",
		"SUBCORE_VM_FIXED3750_E4_50G",
		"SUBCORE_VM_FIXED3800_E4_50G",
		"SUBCORE_VM_FIXED3825_E4_50G",
		"SUBCORE_VM_FIXED3850_E4_50G",
		"SUBCORE_VM_FIXED3875_E4_50G",
		"SUBCORE_VM_FIXED3900_E4_50G",
		"SUBCORE_VM_FIXED3975_E4_50G",
		"SUBCORE_VM_FIXED4000_E4_50G",
		"SUBCORE_VM_FIXED4025_E4_50G",
		"SUBCORE_VM_FIXED4050_E4_50G",
		"SUBCORE_VM_FIXED4100_E4_50G",
		"SUBCORE_VM_FIXED4125_E4_50G",
		"SUBCORE_VM_FIXED4200_E4_50G",
		"SUBCORE_VM_FIXED4225_E4_50G",
		"SUBCORE_VM_FIXED4250_E4_50G",
		"SUBCORE_VM_FIXED4275_E4_50G",
		"SUBCORE_VM_FIXED4300_E4_50G",
		"SUBCORE_VM_FIXED4350_E4_50G",
		"SUBCORE_VM_FIXED4375_E4_50G",
		"SUBCORE_VM_FIXED4400_E4_50G",
		"SUBCORE_VM_FIXED4425_E4_50G",
		"SUBCORE_VM_FIXED4500_E4_50G",
		"SUBCORE_VM_FIXED4550_E4_50G",
		"SUBCORE_VM_FIXED4575_E4_50G",
		"SUBCORE_VM_FIXED4600_E4_50G",
		"SUBCORE_VM_FIXED4625_E4_50G",
		"SUBCORE_VM_FIXED4650_E4_50G",
		"SUBCORE_VM_FIXED4675_E4_50G",
		"SUBCORE_VM_FIXED4700_E4_50G",
		"SUBCORE_VM_FIXED4725_E4_50G",
		"SUBCORE_VM_FIXED4750_E4_50G",
		"SUBCORE_VM_FIXED4800_E4_50G",
		"SUBCORE_VM_FIXED4875_E4_50G",
		"SUBCORE_VM_FIXED4900_E4_50G",
		"SUBCORE_VM_FIXED4950_E4_50G",
		"SUBCORE_VM_FIXED5000_E4_50G",
		"SUBCORE_VM_FIXED0020_A1_50G",
		"SUBCORE_VM_FIXED0040_A1_50G",
		"SUBCORE_VM_FIXED0060_A1_50G",
		"SUBCORE_VM_FIXED0080_A1_50G",
		"SUBCORE_VM_FIXED0100_A1_50G",
		"SUBCORE_VM_FIXED0120_A1_50G",
		"SUBCORE_VM_FIXED0140_A1_50G",
		"SUBCORE_VM_FIXED0160_A1_50G",
		"SUBCORE_VM_FIXED0180_A1_50G",
		"SUBCORE_VM_FIXED0200_A1_50G",
		"SUBCORE_VM_FIXED0220_A1_50G",
		"SUBCORE_VM_FIXED0240_A1_50G",
		"SUBCORE_VM_FIXED0260_A1_50G",
		"SUBCORE_VM_FIXED0280_A1_50G",
		"SUBCORE_VM_FIXED0300_A1_50G",
		"SUBCORE_VM_FIXED0320_A1_50G",
		"SUBCORE_VM_FIXED0340_A1_50G",
		"SUBCORE_VM_FIXED0360_A1_50G",
		"SUBCORE_VM_FIXED0380_A1_50G",
		"SUBCORE_VM_FIXED0400_A1_50G",
		"SUBCORE_VM_FIXED0420_A1_50G",
		"SUBCORE_VM_FIXED0440_A1_50G",
		"SUBCORE_VM_FIXED0460_A1_50G",
		"SUBCORE_VM_FIXED0480_A1_50G",
		"SUBCORE_VM_FIXED0500_A1_50G",
		"SUBCORE_VM_FIXED0520_A1_50G",
		"SUBCORE_VM_FIXED0540_A1_50G",
		"SUBCORE_VM_FIXED0560_A1_50G",
		"SUBCORE_VM_FIXED0580_A1_50G",
		"SUBCORE_VM_FIXED0600_A1_50G",
		"SUBCORE_VM_FIXED0620_A1_50G",
		"SUBCORE_VM_FIXED0640_A1_50G",
		"SUBCORE_VM_FIXED0660_A1_50G",
		"SUBCORE_VM_FIXED0680_A1_50G",
		"SUBCORE_VM_FIXED0700_A1_50G",
		"SUBCORE_VM_FIXED0720_A1_50G",
		"SUBCORE_VM_FIXED0740_A1_50G",
		"SUBCORE_VM_FIXED0760_A1_50G",
		"SUBCORE_VM_FIXED0780_A1_50G",
		"SUBCORE_VM_FIXED0800_A1_50G",
		"SUBCORE_VM_FIXED0820_A1_50G",
		"SUBCORE_VM_FIXED0840_A1_50G",
		"SUBCORE_VM_FIXED0860_A1_50G",
		"SUBCORE_VM_FIXED0880_A1_50G",
		"SUBCORE_VM_FIXED0900_A1_50G",
		"SUBCORE_VM_FIXED0920_A1_50G",
		"SUBCORE_VM_FIXED0940_A1_50G",
		"SUBCORE_VM_FIXED0960_A1_50G",
		"SUBCORE_VM_FIXED0980_A1_50G",
		"SUBCORE_VM_FIXED1000_A1_50G",
		"SUBCORE_VM_FIXED1020_A1_50G",
		"SUBCORE_VM_FIXED1040_A1_50G",
		"SUBCORE_VM_FIXED1060_A1_50G",
		"SUBCORE_VM_FIXED1080_A1_50G",
		"SUBCORE_VM_FIXED1100_A1_50G",
		"SUBCORE_VM_FIXED1120_A1_50G",
		"SUBCORE_VM_FIXED1140_A1_50G",
		"SUBCORE_VM_FIXED1160_A1_50G",
		"SUBCORE_VM_FIXED1180_A1_50G",
		"SUBCORE_VM_FIXED1200_A1_50G",
		"SUBCORE_VM_FIXED1220_A1_50G",
		"SUBCORE_VM_FIXED1240_A1_50G",
		"SUBCORE_VM_FIXED1260_A1_50G",
		"SUBCORE_VM_FIXED1280_A1_50G",
		"SUBCORE_VM_FIXED1300_A1_50G",
		"SUBCORE_VM_FIXED1320_A1_50G",
		"SUBCORE_VM_FIXED1340_A1_50G",
		"SUBCORE_VM_FIXED1360_A1_50G",
		"SUBCORE_VM_FIXED1380_A1_50G",
		"SUBCORE_VM_FIXED1400_A1_50G",
		"SUBCORE_VM_FIXED1420_A1_50G",
		"SUBCORE_VM_FIXED1440_A1_50G",
		"SUBCORE_VM_FIXED1460_A1_50G",
		"SUBCORE_VM_FIXED1480_A1_50G",
		"SUBCORE_VM_FIXED1500_A1_50G",
		"SUBCORE_VM_FIXED1520_A1_50G",
		"SUBCORE_VM_FIXED1540_A1_50G",
		"SUBCORE_VM_FIXED1560_A1_50G",
		"SUBCORE_VM_FIXED1580_A1_50G",
		"SUBCORE_VM_FIXED1600_A1_50G",
		"SUBCORE_VM_FIXED1620_A1_50G",
		"SUBCORE_VM_FIXED1640_A1_50G",
		"SUBCORE_VM_FIXED1660_A1_50G",
		"SUBCORE_VM_FIXED1680_A1_50G",
		"SUBCORE_VM_FIXED1700_A1_50G",
		"SUBCORE_VM_FIXED1720_A1_50G",
		"SUBCORE_VM_FIXED1740_A1_50G",
		"SUBCORE_VM_FIXED1760_A1_50G",
		"SUBCORE_VM_FIXED1780_A1_50G",
		"SUBCORE_VM_FIXED1800_A1_50G",
		"SUBCORE_VM_FIXED1820_A1_50G",
		"SUBCORE_VM_FIXED1840_A1_50G",
		"SUBCORE_VM_FIXED1860_A1_50G",
		"SUBCORE_VM_FIXED1880_A1_50G",
		"SUBCORE_VM_FIXED1900_A1_50G",
		"SUBCORE_VM_FIXED1920_A1_50G",
		"SUBCORE_VM_FIXED1940_A1_50G",
		"SUBCORE_VM_FIXED1960_A1_50G",
		"SUBCORE_VM_FIXED1980_A1_50G",
		"SUBCORE_VM_FIXED2000_A1_50G",
		"SUBCORE_VM_FIXED2020_A1_50G",
		"SUBCORE_VM_FIXED2040_A1_50G",
		"SUBCORE_VM_FIXED2060_A1_50G",
		"SUBCORE_VM_FIXED2080_A1_50G",
		"SUBCORE_VM_FIXED2100_A1_50G",
		"SUBCORE_VM_FIXED2120_A1_50G",
		"SUBCORE_VM_FIXED2140_A1_50G",
		"SUBCORE_VM_FIXED2160_A1_50G",
		"SUBCORE_VM_FIXED2180_A1_50G",
		"SUBCORE_VM_FIXED2200_A1_50G",
		"SUBCORE_VM_FIXED2220_A1_50G",
		"SUBCORE_VM_FIXED2240_A1_50G",
		"SUBCORE_VM_FIXED2260_A1_50G",
		"SUBCORE_VM_FIXED2280_A1_50G",
		"SUBCORE_VM_FIXED2300_A1_50G",
		"SUBCORE_VM_FIXED2320_A1_50G",
		"SUBCORE_VM_FIXED2340_A1_50G",
		"SUBCORE_VM_FIXED2360_A1_50G",
		"SUBCORE_VM_FIXED2380_A1_50G",
		"SUBCORE_VM_FIXED2400_A1_50G",
		"SUBCORE_VM_FIXED2420_A1_50G",
		"SUBCORE_VM_FIXED2440_A1_50G",
		"SUBCORE_VM_FIXED2460_A1_50G",
		"SUBCORE_VM_FIXED2480_A1_50G",
		"SUBCORE_VM_FIXED2500_A1_50G",
		"SUBCORE_VM_FIXED2520_A1_50G",
		"SUBCORE_VM_FIXED2540_A1_50G",
		"SUBCORE_VM_FIXED2560_A1_50G",
		"SUBCORE_VM_FIXED2580_A1_50G",
		"SUBCORE_VM_FIXED2600_A1_50G",
		"SUBCORE_VM_FIXED2620_A1_50G",
		"SUBCORE_VM_FIXED2640_A1_50G",
		"SUBCORE_VM_FIXED2660_A1_50G",
		"SUBCORE_VM_FIXED2680_A1_50G",
		"SUBCORE_VM_FIXED2700_A1_50G",
		"SUBCORE_VM_FIXED2720_A1_50G",
		"SUBCORE_VM_FIXED2740_A1_50G",
		"SUBCORE_VM_FIXED2760_A1_50G",
		"SUBCORE_VM_FIXED2780_A1_50G",
		"SUBCORE_VM_FIXED2800_A1_50G",
		"SUBCORE_VM_FIXED2820_A1_50G",
		"SUBCORE_VM_FIXED2840_A1_50G",
		"SUBCORE_VM_FIXED2860_A1_50G",
		"SUBCORE_VM_FIXED2880_A1_50G",
		"SUBCORE_VM_FIXED2900_A1_50G",
		"SUBCORE_VM_FIXED2920_A1_50G",
		"SUBCORE_VM_FIXED2940_A1_50G",
		"SUBCORE_VM_FIXED2960_A1_50G",
		"SUBCORE_VM_FIXED2980_A1_50G",
		"SUBCORE_VM_FIXED3000_A1_50G",
		"SUBCORE_VM_FIXED3020_A1_50G",
		"SUBCORE_VM_FIXED3040_A1_50G",
		"SUBCORE_VM_FIXED3060_A1_50G",
		"SUBCORE_VM_FIXED3080_A1_50G",
		"SUBCORE_VM_FIXED3100_A1_50G",
		"SUBCORE_VM_FIXED3120_A1_50G",
		"SUBCORE_VM_FIXED3140_A1_50G",
		"SUBCORE_VM_FIXED3160_A1_50G",
		"SUBCORE_VM_FIXED3180_A1_50G",
		"SUBCORE_VM_FIXED3200_A1_50G",
		"SUBCORE_VM_FIXED3220_A1_50G",
		"SUBCORE_VM_FIXED3240_A1_50G",
		"SUBCORE_VM_FIXED3260_A1_50G",
		"SUBCORE_VM_FIXED3280_A1_50G",
		"SUBCORE_VM_FIXED3300_A1_50G",
		"SUBCORE_VM_FIXED3320_A1_50G",
		"SUBCORE_VM_FIXED3340_A1_50G",
		"SUBCORE_VM_FIXED3360_A1_50G",
		"SUBCORE_VM_FIXED3380_A1_50G",
		"SUBCORE_VM_FIXED3400_A1_50G",
		"SUBCORE_VM_FIXED3420_A1_50G",
		"SUBCORE_VM_FIXED3440_A1_50G",
		"SUBCORE_VM_FIXED3460_A1_50G",
		"SUBCORE_VM_FIXED3480_A1_50G",
		"SUBCORE_VM_FIXED3500_A1_50G",
		"SUBCORE_VM_FIXED3520_A1_50G",
		"SUBCORE_VM_FIXED3540_A1_50G",
		"SUBCORE_VM_FIXED3560_A1_50G",
		"SUBCORE_VM_FIXED3580_A1_50G",
		"SUBCORE_VM_FIXED3600_A1_50G",
		"SUBCORE_VM_FIXED3620_A1_50G",
		"SUBCORE_VM_FIXED3640_A1_50G",
		"SUBCORE_VM_FIXED3660_A1_50G",
		"SUBCORE_VM_FIXED3680_A1_50G",
		"SUBCORE_VM_FIXED3700_A1_50G",
		"SUBCORE_VM_FIXED3720_A1_50G",
		"SUBCORE_VM_FIXED3740_A1_50G",
		"SUBCORE_VM_FIXED3760_A1_50G",
		"SUBCORE_VM_FIXED3780_A1_50G",
		"SUBCORE_VM_FIXED3800_A1_50G",
		"SUBCORE_VM_FIXED3820_A1_50G",
		"SUBCORE_VM_FIXED3840_A1_50G",
		"SUBCORE_VM_FIXED3860_A1_50G",
		"SUBCORE_VM_FIXED3880_A1_50G",
		"SUBCORE_VM_FIXED3900_A1_50G",
		"SUBCORE_VM_FIXED3920_A1_50G",
		"SUBCORE_VM_FIXED3940_A1_50G",
		"SUBCORE_VM_FIXED3960_A1_50G",
		"SUBCORE_VM_FIXED3980_A1_50G",
		"SUBCORE_VM_FIXED4000_A1_50G",
		"SUBCORE_VM_FIXED4020_A1_50G",
		"SUBCORE_VM_FIXED4040_A1_50G",
		"SUBCORE_VM_FIXED4060_A1_50G",
		"SUBCORE_VM_FIXED4080_A1_50G",
		"SUBCORE_VM_FIXED4100_A1_50G",
		"SUBCORE_VM_FIXED4120_A1_50G",
		"SUBCORE_VM_FIXED4140_A1_50G",
		"SUBCORE_VM_FIXED4160_A1_50G",
		"SUBCORE_VM_FIXED4180_A1_50G",
		"SUBCORE_VM_FIXED4200_A1_50G",
		"SUBCORE_VM_FIXED4220_A1_50G",
		"SUBCORE_VM_FIXED4240_A1_50G",
		"SUBCORE_VM_FIXED4260_A1_50G",
		"SUBCORE_VM_FIXED4280_A1_50G",
		"SUBCORE_VM_FIXED4300_A1_50G",
		"SUBCORE_VM_FIXED4320_A1_50G",
		"SUBCORE_VM_FIXED4340_A1_50G",
		"SUBCORE_VM_FIXED4360_A1_50G",
		"SUBCORE_VM_FIXED4380_A1_50G",
		"SUBCORE_VM_FIXED4400_A1_50G",
		"SUBCORE_VM_FIXED4420_A1_50G",
		"SUBCORE_VM_FIXED4440_A1_50G",
		"SUBCORE_VM_FIXED4460_A1_50G",
		"SUBCORE_VM_FIXED4480_A1_50G",
		"SUBCORE_VM_FIXED4500_A1_50G",
		"SUBCORE_VM_FIXED4520_A1_50G",
		"SUBCORE_VM_FIXED4540_A1_50G",
		"SUBCORE_VM_FIXED4560_A1_50G",
		"SUBCORE_VM_FIXED4580_A1_50G",
		"SUBCORE_VM_FIXED4600_A1_50G",
		"SUBCORE_VM_FIXED4620_A1_50G",
		"SUBCORE_VM_FIXED4640_A1_50G",
		"SUBCORE_VM_FIXED4660_A1_50G",
		"SUBCORE_VM_FIXED4680_A1_50G",
		"SUBCORE_VM_FIXED4700_A1_50G",
		"SUBCORE_VM_FIXED4720_A1_50G",
		"SUBCORE_VM_FIXED4740_A1_50G",
		"SUBCORE_VM_FIXED4760_A1_50G",
		"SUBCORE_VM_FIXED4780_A1_50G",
		"SUBCORE_VM_FIXED4800_A1_50G",
		"SUBCORE_VM_FIXED4820_A1_50G",
		"SUBCORE_VM_FIXED4840_A1_50G",
		"SUBCORE_VM_FIXED4860_A1_50G",
		"SUBCORE_VM_FIXED4880_A1_50G",
		"SUBCORE_VM_FIXED4900_A1_50G",
		"SUBCORE_VM_FIXED4920_A1_50G",
		"SUBCORE_VM_FIXED4940_A1_50G",
		"SUBCORE_VM_FIXED4960_A1_50G",
		"SUBCORE_VM_FIXED4980_A1_50G",
		"SUBCORE_VM_FIXED5000_A1_50G",
		"SUBCORE_VM_FIXED0090_X9_50G",
		"SUBCORE_VM_FIXED0180_X9_50G",
		"SUBCORE_VM_FIXED0270_X9_50G",
		"SUBCORE_VM_FIXED0360_X9_50G",
		"SUBCORE_VM_FIXED0450_X9_50G",
		"SUBCORE_VM_FIXED0540_X9_50G",
		"SUBCORE_VM_FIXED0630_X9_50G",
		"SUBCORE_VM_FIXED0720_X9_50G",
		"SUBCORE_VM_FIXED0810_X9_50G",
		"SUBCORE_VM_FIXED0900_X9_50G",
		"SUBCORE_VM_FIXED0990_X9_50G",
		"SUBCORE_VM_FIXED1080_X9_50G",
		"SUBCORE_VM_FIXED1170_X9_50G",
		"SUBCORE_VM_FIXED1260_X9_50G",
		"SUBCORE_VM_FIXED1350_X9_50G",
		"SUBCORE_VM_FIXED1440_X9_50G",
		"SUBCORE_VM_FIXED1530_X9_50G",
		"SUBCORE_VM_FIXED1620_X9_50G",
		"SUBCORE_VM_FIXED1710_X9_50G",
		"SUBCORE_VM_FIXED1800_X9_50G",
		"SUBCORE_VM_FIXED1890_X9_50G",
		"SUBCORE_VM_FIXED1980_X9_50G",
		"SUBCORE_VM_FIXED2070_X9_50G",
		"SUBCORE_VM_FIXED2160_X9_50G",
		"SUBCORE_VM_FIXED2250_X9_50G",
		"SUBCORE_VM_FIXED2340_X9_50G",
		"SUBCORE_VM_FIXED2430_X9_50G",
		"SUBCORE_VM_FIXED2520_X9_50G",
		"SUBCORE_VM_FIXED2610_X9_50G",
		"SUBCORE_VM_FIXED2700_X9_50G",
		"SUBCORE_VM_FIXED2790_X9_50G",
		"SUBCORE_VM_FIXED2880_X9_50G",
		"SUBCORE_VM_FIXED2970_X9_50G",
		"SUBCORE_VM_FIXED3060_X9_50G",
		"SUBCORE_VM_FIXED3150_X9_50G",
		"SUBCORE_VM_FIXED3240_X9_50G",
		"SUBCORE_VM_FIXED3330_X9_50G",
		"SUBCORE_VM_FIXED3420_X9_50G",
		"SUBCORE_VM_FIXED3510_X9_50G",
		"SUBCORE_VM_FIXED3600_X9_50G",
		"SUBCORE_VM_FIXED3690_X9_50G",
		"SUBCORE_VM_FIXED3780_X9_50G",
		"SUBCORE_VM_FIXED3870_X9_50G",
		"SUBCORE_VM_FIXED3960_X9_50G",
		"SUBCORE_VM_FIXED4050_X9_50G",
		"SUBCORE_VM_FIXED4140_X9_50G",
		"SUBCORE_VM_FIXED4230_X9_50G",
		"SUBCORE_VM_FIXED4320_X9_50G",
		"SUBCORE_VM_FIXED4410_X9_50G",
		"SUBCORE_VM_FIXED4500_X9_50G",
		"SUBCORE_VM_FIXED4590_X9_50G",
		"SUBCORE_VM_FIXED4680_X9_50G",
		"SUBCORE_VM_FIXED4770_X9_50G",
		"SUBCORE_VM_FIXED4860_X9_50G",
		"SUBCORE_VM_FIXED4950_X9_50G",
		"DYNAMIC_A1_50G",
		"FIXED0040_A1_50G",
		"FIXED0100_A1_50G",
		"FIXED0200_A1_50G",
		"FIXED0300_A1_50G",
		"FIXED0400_A1_50G",
		"FIXED0500_A1_50G",
		"FIXED0600_A1_50G",
		"FIXED0700_A1_50G",
		"FIXED0800_A1_50G",
		"FIXED0900_A1_50G",
		"FIXED1000_A1_50G",
		"FIXED1100_A1_50G",
		"FIXED1200_A1_50G",
		"FIXED1300_A1_50G",
		"FIXED1400_A1_50G",
		"FIXED1500_A1_50G",
		"FIXED1600_A1_50G",
		"FIXED1700_A1_50G",
		"FIXED1800_A1_50G",
		"FIXED1900_A1_50G",
		"FIXED2000_A1_50G",
		"FIXED2100_A1_50G",
		"FIXED2200_A1_50G",
		"FIXED2300_A1_50G",
		"FIXED2400_A1_50G",
		"FIXED2500_A1_50G",
		"FIXED2600_A1_50G",
		"FIXED2700_A1_50G",
		"FIXED2800_A1_50G",
		"FIXED2900_A1_50G",
		"FIXED3000_A1_50G",
		"FIXED3100_A1_50G",
		"FIXED3200_A1_50G",
		"FIXED3300_A1_50G",
		"FIXED3400_A1_50G",
		"FIXED3500_A1_50G",
		"FIXED3600_A1_50G",
		"FIXED3700_A1_50G",
		"FIXED3800_A1_50G",
		"FIXED3900_A1_50G",
		"FIXED4000_A1_50G",
		"ENTIREHOST_A1_50G",
		"DYNAMIC_X9_50G",
		"FIXED0040_X9_50G",
		"FIXED0400_X9_50G",
		"FIXED0800_X9_50G",
		"FIXED1200_X9_50G",
		"FIXED1600_X9_50G",
		"FIXED2000_X9_50G",
		"FIXED2400_X9_50G",
		"FIXED2800_X9_50G",
		"FIXED3200_X9_50G",
		"FIXED3600_X9_50G",
		"FIXED4000_X9_50G",
		"STANDARD_VM_FIXED0100_X9_50G",
		"STANDARD_VM_FIXED0200_X9_50G",
		"STANDARD_VM_FIXED0300_X9_50G",
		"STANDARD_VM_FIXED0400_X9_50G",
		"STANDARD_VM_FIXED0500_X9_50G",
		"STANDARD_VM_FIXED0600_X9_50G",
		"STANDARD_VM_FIXED0700_X9_50G",
		"STANDARD_VM_FIXED0800_X9_50G",
		"STANDARD_VM_FIXED0900_X9_50G",
		"STANDARD_VM_FIXED1000_X9_50G",
		"STANDARD_VM_FIXED1100_X9_50G",
		"STANDARD_VM_FIXED1200_X9_50G",
		"STANDARD_VM_FIXED1300_X9_50G",
		"STANDARD_VM_FIXED1400_X9_50G",
		"STANDARD_VM_FIXED1500_X9_50G",
		"STANDARD_VM_FIXED1600_X9_50G",
		"STANDARD_VM_FIXED1700_X9_50G",
		"STANDARD_VM_FIXED1800_X9_50G",
		"STANDARD_VM_FIXED1900_X9_50G",
		"STANDARD_VM_FIXED2000_X9_50G",
		"STANDARD_VM_FIXED2100_X9_50G",
		"STANDARD_VM_FIXED2200_X9_50G",
		"STANDARD_VM_FIXED2300_X9_50G",
		"STANDARD_VM_FIXED2400_X9_50G",
		"STANDARD_VM_FIXED2500_X9_50G",
		"STANDARD_VM_FIXED2600_X9_50G",
		"STANDARD_VM_FIXED2700_X9_50G",
		"STANDARD_VM_FIXED2800_X9_50G",
		"STANDARD_VM_FIXED2900_X9_50G",
		"STANDARD_VM_FIXED3000_X9_50G",
		"STANDARD_VM_FIXED3100_X9_50G",
		"STANDARD_VM_FIXED3200_X9_50G",
		"STANDARD_VM_FIXED3300_X9_50G",
		"STANDARD_VM_FIXED3400_X9_50G",
		"STANDARD_VM_FIXED3500_X9_50G",
		"STANDARD_VM_FIXED3600_X9_50G",
		"STANDARD_VM_FIXED3700_X9_50G",
		"STANDARD_VM_FIXED3800_X9_50G",
		"STANDARD_VM_FIXED3900_X9_50G",
		"STANDARD_VM_FIXED4000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0775_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0825_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED0975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED1950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2175_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2475_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2775_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2925_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED2975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3075_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3150_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3325_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3450_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3525_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3825_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3850_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED3975_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4000_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4025_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4050_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4100_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4125_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4200_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4225_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4250_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4275_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4300_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4350_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4375_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4400_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4425_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4500_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4550_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4575_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4600_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4625_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4650_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4675_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4700_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4725_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4750_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4800_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4875_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4900_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED4950_X9_50G",
		"SUBCORE_STANDARD_VM_FIXED5000_X9_50G",
		"ENTIREHOST_X9_50G",
	}
}

// GetMappingCreateInternalVnicDetailsVnicShapeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalVnicDetailsVnicShapeEnum(val string) (CreateInternalVnicDetailsVnicShapeEnum, bool) {
	enum, ok := mappingCreateInternalVnicDetailsVnicShapeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
