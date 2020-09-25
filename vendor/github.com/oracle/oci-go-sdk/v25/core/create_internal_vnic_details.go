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

// CreateInternalVnicDetails This structure is used when creating vnic for internal clients.
// For more information about VNICs, see
// Virtual Network Interface Cards (VNICs) (https://docs.cloud.oracle.com/Content/Network/Tasks/managingVNICs.htm).
type CreateInternalVnicDetails struct {

	// Indicates if the VNIC is primary which means it cannot be detached.
	IsPrimary *bool `mandatory:"true" json:"isPrimary"`

	// Whether the VNIC should be assigned a public IP address. Defaults to whether
	// the subnet is public or private. If not set and the VNIC is being created
	// in a private subnet (that is, where `prohibitPublicIpOnVnic` = true in the
	// Subnet), then no public IP address is assigned.
	// If not set and the subnet is public (`prohibitPublicIpOnVnic` = false), then
	// a public IP address is assigned. If set to true and
	// `prohibitPublicIpOnVnic` = true, an error is returned.
	// **Note:** This public IP address is associated with the primary private IP
	// on the VNIC. For more information, see
	// IP Addresses (https://docs.cloud.oracle.com/Content/Network/Tasks/managingIPaddresses.htm).
	// **Note:** There's a limit to the number of PublicIp
	// a VNIC or instance can have. If you try to create a secondary VNIC
	// with an assigned public IP for an instance that has already
	// reached its public IP limit, an error is returned. For information
	// about the public IP limits, see
	// Public IP Addresses (https://docs.cloud.oracle.com/Content/Network/Tasks/managingpublicIPs.htm).
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
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name for the VNIC. Does not have to be unique.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
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
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/Content/Network/Concepts/dns.htm).
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

	// OCID of the pool object created by the current tenancy
	PublicIpPoolId *string `mandatory:"false" json:"publicIpPoolId"`

	// ID of the customer visible upstream resource that the VNIC is associated with. This property is
	// exposed to customers as part of API to list members of a network security group.
	// For example, if the VNIC is associated with a loadbalancer or dbsystem instance, then it needs
	// to be set to corresponding customer visible loadbalancer or dbsystem instance OCID.
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
	// Using a Private IP as a Route Target (https://docs.cloud.oracle.com/Content/Network/Tasks/managingroutetables.htm#privateip).
	// Example: `true`
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`

	// The OCID of the subnet to create the VNIC in. When launching an instance,
	// use this `subnetId` instead of the deprecated `subnetId` in
	// LaunchInstanceDetails.
	// At least one of them is required; if you provide both, the values must match.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The OCID of the VLAN that the VNIC belongs to
	VlanId *string `mandatory:"false" json:"vlanId"`

	// Shape of VNIC that will be used to allocate resource in the data plane once the VNIC is attached
	VnicShape CreateInternalVnicDetailsVnicShapeEnum `mandatory:"false" json:"vnicShape,omitempty"`
}

func (m CreateInternalVnicDetails) String() string {
	return common.PointerString(m)
}

// CreateInternalVnicDetailsInternalPoolNameEnum Enum with underlying type: string
type CreateInternalVnicDetailsInternalPoolNameEnum string

// Set of constants representing the allowable values for CreateInternalVnicDetailsInternalPoolNameEnum
const (
	CreateInternalVnicDetailsInternalPoolNameExternal   CreateInternalVnicDetailsInternalPoolNameEnum = "EXTERNAL"
	CreateInternalVnicDetailsInternalPoolNameSociEgress CreateInternalVnicDetailsInternalPoolNameEnum = "SOCI_EGRESS"
)

var mappingCreateInternalVnicDetailsInternalPoolName = map[string]CreateInternalVnicDetailsInternalPoolNameEnum{
	"EXTERNAL":    CreateInternalVnicDetailsInternalPoolNameExternal,
	"SOCI_EGRESS": CreateInternalVnicDetailsInternalPoolNameSociEgress,
}

// GetCreateInternalVnicDetailsInternalPoolNameEnumValues Enumerates the set of values for CreateInternalVnicDetailsInternalPoolNameEnum
func GetCreateInternalVnicDetailsInternalPoolNameEnumValues() []CreateInternalVnicDetailsInternalPoolNameEnum {
	values := make([]CreateInternalVnicDetailsInternalPoolNameEnum, 0)
	for _, v := range mappingCreateInternalVnicDetailsInternalPoolName {
		values = append(values, v)
	}
	return values
}

// CreateInternalVnicDetailsVnicShapeEnum Enum with underlying type: string
type CreateInternalVnicDetailsVnicShapeEnum string

// Set of constants representing the allowable values for CreateInternalVnicDetailsVnicShapeEnum
const (
	CreateInternalVnicDetailsVnicShapeDynamic               CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC"
	CreateInternalVnicDetailsVnicShapeFixed0040             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040"
	CreateInternalVnicDetailsVnicShapeFixed0060             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060"
	CreateInternalVnicDetailsVnicShapeFixed0060Psm          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060_PSM"
	CreateInternalVnicDetailsVnicShapeFixed0100             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100"
	CreateInternalVnicDetailsVnicShapeFixed0120             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120"
	CreateInternalVnicDetailsVnicShapeFixed01202x           CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120_2X"
	CreateInternalVnicDetailsVnicShapeFixed0200             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200"
	CreateInternalVnicDetailsVnicShapeFixed0240             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0240"
	CreateInternalVnicDetailsVnicShapeFixed0480             CreateInternalVnicDetailsVnicShapeEnum = "FIXED0480"
	CreateInternalVnicDetailsVnicShapeEntirehost            CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST"
	CreateInternalVnicDetailsVnicShapeDynamic25g            CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_25G"
	CreateInternalVnicDetailsVnicShapeFixed004025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_25G"
	CreateInternalVnicDetailsVnicShapeFixed010025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_25G"
	CreateInternalVnicDetailsVnicShapeFixed020025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_25G"
	CreateInternalVnicDetailsVnicShapeFixed040025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_25G"
	CreateInternalVnicDetailsVnicShapeFixed080025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_25G"
	CreateInternalVnicDetailsVnicShapeFixed160025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_25G"
	CreateInternalVnicDetailsVnicShapeFixed240025g          CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_25G"
	CreateInternalVnicDetailsVnicShapeEntirehost25g         CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_25G"
	CreateInternalVnicDetailsVnicShapeDynamicE125g          CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0040E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0070E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0070_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0140E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0140_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0280E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0280_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0560E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0560_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed1120E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1120_E1_25G"
	CreateInternalVnicDetailsVnicShapeFixed1680E125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1680_E1_25G"
	CreateInternalVnicDetailsVnicShapeEntirehostE125g       CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_E1_25G"
	CreateInternalVnicDetailsVnicShapeDynamicB125g          CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0040B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0060B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0060_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0120B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0120_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0240B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0240_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0480B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0480_B1_25G"
	CreateInternalVnicDetailsVnicShapeFixed0960B125g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0960_B1_25G"
	CreateInternalVnicDetailsVnicShapeEntirehostB125g       CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_B1_25G"
	CreateInternalVnicDetailsVnicShapeMicroVmFixed0048E125g CreateInternalVnicDetailsVnicShapeEnum = "MICRO_VM_FIXED0048_E1_25G"
	CreateInternalVnicDetailsVnicShapeMicroLbFixed0001E125g CreateInternalVnicDetailsVnicShapeEnum = "MICRO_LB_FIXED0001_E1_25G"
	CreateInternalVnicDetailsVnicShapeVnicaasFixed0200      CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FIXED0200"
	CreateInternalVnicDetailsVnicShapeVnicaasFixed0400      CreateInternalVnicDetailsVnicShapeEnum = "VNICAAS_FIXED0400"
	CreateInternalVnicDetailsVnicShapeDynamicE350g          CreateInternalVnicDetailsVnicShapeEnum = "DYNAMIC_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0040E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0040_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0100E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0200E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0300E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0400E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0500E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0600E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0700E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0800E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed0900E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED0900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1000E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1100E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1200E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1300E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1400E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1500E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1600E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1700E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1800E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed1900E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED1900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2000E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2100E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2200E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2300E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2400E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2500E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2600E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2700E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2800E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed2900E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED2900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3000E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3000_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3100E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3100_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3200E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3200_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3300E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3300_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3400E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3400_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3500E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3500_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3600E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3600_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3700E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3700_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3800E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3800_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed3900E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED3900_E3_50G"
	CreateInternalVnicDetailsVnicShapeFixed4000E350g        CreateInternalVnicDetailsVnicShapeEnum = "FIXED4000_E3_50G"
	CreateInternalVnicDetailsVnicShapeEntirehostE350g       CreateInternalVnicDetailsVnicShapeEnum = "ENTIREHOST_E3_50G"
)

var mappingCreateInternalVnicDetailsVnicShape = map[string]CreateInternalVnicDetailsVnicShapeEnum{
	"DYNAMIC":                   CreateInternalVnicDetailsVnicShapeDynamic,
	"FIXED0040":                 CreateInternalVnicDetailsVnicShapeFixed0040,
	"FIXED0060":                 CreateInternalVnicDetailsVnicShapeFixed0060,
	"FIXED0060_PSM":             CreateInternalVnicDetailsVnicShapeFixed0060Psm,
	"FIXED0100":                 CreateInternalVnicDetailsVnicShapeFixed0100,
	"FIXED0120":                 CreateInternalVnicDetailsVnicShapeFixed0120,
	"FIXED0120_2X":              CreateInternalVnicDetailsVnicShapeFixed01202x,
	"FIXED0200":                 CreateInternalVnicDetailsVnicShapeFixed0200,
	"FIXED0240":                 CreateInternalVnicDetailsVnicShapeFixed0240,
	"FIXED0480":                 CreateInternalVnicDetailsVnicShapeFixed0480,
	"ENTIREHOST":                CreateInternalVnicDetailsVnicShapeEntirehost,
	"DYNAMIC_25G":               CreateInternalVnicDetailsVnicShapeDynamic25g,
	"FIXED0040_25G":             CreateInternalVnicDetailsVnicShapeFixed004025g,
	"FIXED0100_25G":             CreateInternalVnicDetailsVnicShapeFixed010025g,
	"FIXED0200_25G":             CreateInternalVnicDetailsVnicShapeFixed020025g,
	"FIXED0400_25G":             CreateInternalVnicDetailsVnicShapeFixed040025g,
	"FIXED0800_25G":             CreateInternalVnicDetailsVnicShapeFixed080025g,
	"FIXED1600_25G":             CreateInternalVnicDetailsVnicShapeFixed160025g,
	"FIXED2400_25G":             CreateInternalVnicDetailsVnicShapeFixed240025g,
	"ENTIREHOST_25G":            CreateInternalVnicDetailsVnicShapeEntirehost25g,
	"DYNAMIC_E1_25G":            CreateInternalVnicDetailsVnicShapeDynamicE125g,
	"FIXED0040_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed0040E125g,
	"FIXED0070_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed0070E125g,
	"FIXED0140_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed0140E125g,
	"FIXED0280_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed0280E125g,
	"FIXED0560_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed0560E125g,
	"FIXED1120_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed1120E125g,
	"FIXED1680_E1_25G":          CreateInternalVnicDetailsVnicShapeFixed1680E125g,
	"ENTIREHOST_E1_25G":         CreateInternalVnicDetailsVnicShapeEntirehostE125g,
	"DYNAMIC_B1_25G":            CreateInternalVnicDetailsVnicShapeDynamicB125g,
	"FIXED0040_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0040B125g,
	"FIXED0060_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0060B125g,
	"FIXED0120_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0120B125g,
	"FIXED0240_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0240B125g,
	"FIXED0480_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0480B125g,
	"FIXED0960_B1_25G":          CreateInternalVnicDetailsVnicShapeFixed0960B125g,
	"ENTIREHOST_B1_25G":         CreateInternalVnicDetailsVnicShapeEntirehostB125g,
	"MICRO_VM_FIXED0048_E1_25G": CreateInternalVnicDetailsVnicShapeMicroVmFixed0048E125g,
	"MICRO_LB_FIXED0001_E1_25G": CreateInternalVnicDetailsVnicShapeMicroLbFixed0001E125g,
	"VNICAAS_FIXED0200":         CreateInternalVnicDetailsVnicShapeVnicaasFixed0200,
	"VNICAAS_FIXED0400":         CreateInternalVnicDetailsVnicShapeVnicaasFixed0400,
	"DYNAMIC_E3_50G":            CreateInternalVnicDetailsVnicShapeDynamicE350g,
	"FIXED0040_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0040E350g,
	"FIXED0100_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0100E350g,
	"FIXED0200_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0200E350g,
	"FIXED0300_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0300E350g,
	"FIXED0400_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0400E350g,
	"FIXED0500_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0500E350g,
	"FIXED0600_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0600E350g,
	"FIXED0700_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0700E350g,
	"FIXED0800_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0800E350g,
	"FIXED0900_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed0900E350g,
	"FIXED1000_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1000E350g,
	"FIXED1100_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1100E350g,
	"FIXED1200_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1200E350g,
	"FIXED1300_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1300E350g,
	"FIXED1400_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1400E350g,
	"FIXED1500_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1500E350g,
	"FIXED1600_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1600E350g,
	"FIXED1700_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1700E350g,
	"FIXED1800_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1800E350g,
	"FIXED1900_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed1900E350g,
	"FIXED2000_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2000E350g,
	"FIXED2100_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2100E350g,
	"FIXED2200_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2200E350g,
	"FIXED2300_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2300E350g,
	"FIXED2400_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2400E350g,
	"FIXED2500_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2500E350g,
	"FIXED2600_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2600E350g,
	"FIXED2700_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2700E350g,
	"FIXED2800_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2800E350g,
	"FIXED2900_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed2900E350g,
	"FIXED3000_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3000E350g,
	"FIXED3100_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3100E350g,
	"FIXED3200_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3200E350g,
	"FIXED3300_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3300E350g,
	"FIXED3400_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3400E350g,
	"FIXED3500_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3500E350g,
	"FIXED3600_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3600E350g,
	"FIXED3700_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3700E350g,
	"FIXED3800_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3800E350g,
	"FIXED3900_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed3900E350g,
	"FIXED4000_E3_50G":          CreateInternalVnicDetailsVnicShapeFixed4000E350g,
	"ENTIREHOST_E3_50G":         CreateInternalVnicDetailsVnicShapeEntirehostE350g,
}

// GetCreateInternalVnicDetailsVnicShapeEnumValues Enumerates the set of values for CreateInternalVnicDetailsVnicShapeEnum
func GetCreateInternalVnicDetailsVnicShapeEnumValues() []CreateInternalVnicDetailsVnicShapeEnum {
	values := make([]CreateInternalVnicDetailsVnicShapeEnum, 0)
	for _, v := range mappingCreateInternalVnicDetailsVnicShape {
		values = append(values, v)
	}
	return values
}
