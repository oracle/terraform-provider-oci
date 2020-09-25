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

// InternalVnic This is a vnic type only used in operations with overlay customers and RCE.
// It defines additonal properties: isManaged, resourceType, resourceId, isBMVnic, isGarpEnabled and isServiceVnic
type InternalVnic struct {

	// The VNIC's availability domain.
	// Example: `Uocm:PHX-AD-1`
	AvailabilityDomain *string `mandatory:"true" json:"availabilityDomain"`

	// The OCID of the compartment containing the VNIC.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the VNIC.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the VNIC.
	LifecycleState InternalVnicLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the subnet the VNIC is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the VNIC was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Indicates if the VNIC is managed by a internal partner team. And customer is not allowed
	// the perform update/delete operations on it directly.
	// Defaults to `False`
	IsManaged *bool `mandatory:"false" json:"isManaged"`

	// Type of the customer visible upstream resource that the VNIC is associated with. This property can be
	// exposed to customers as part of API to list members of a network security group.
	// For example, it can be set as,
	//  - `loadbalancer` if corresponding resourceId is a loadbalancer instance's OCID
	//  - `dbsystem` if corresponding resourceId is a dbsystem instance's OCID
	// Note that the partner team creating/managing the VNIC is owner of this metadata.
	// type:
	ResourceType *string `mandatory:"false" json:"resourceType"`

	// ID of the customer visible upstream resource that the VNIC is associated with. This property is
	// exposed to customers as part of API to list members of a network security group.
	// For example, if the VNIC is associated with a loadbalancer or dbsystem instance, then it needs
	// to be set to corresponding customer visible loadbalancer or dbsystem instance OCID.
	// Note that the partner team creating/managing the VNIC is owner of this metadata.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// Indicates if the VNIC is associated with (and will be attached to) a BM instance.
	IsBmVnic *bool `mandatory:"false" json:"isBmVnic"`

	// Indicates if the VNIC is a service vnic.
	IsServiceVnic *bool `mandatory:"false" json:"isServiceVnic"`

	// Indicates if this VNIC can issue GARP requests. False by default.
	IsGarpEnabled *bool `mandatory:"false" json:"isGarpEnabled"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique.
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
	// For more information, see
	// DNS in Your Virtual Cloud Network (https://docs.cloud.oracle.com/Content/Network/Concepts/dns.htm).
	// Example: `bminstance-1`
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Whether the VNIC is the primary VNIC (the VNIC that is automatically created
	// and attached during instance launch).
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The MAC address of the VNIC.
	// Example: `00:00:00:00:00:01`
	MacAddress *string `mandatory:"false" json:"macAddress"`

	// A list of the OCIDs of the network security groups that the VNIC belongs to. For more
	// information about NSGs, see
	// NetworkSecurityGroup.
	NsgIds []string `mandatory:"false" json:"nsgIds"`

	// The private IP address of the primary `privateIp` object on the VNIC.
	// The address is within the CIDR of the VNIC's subnet.
	// **Note: ** This is null if the VNIC is in a subnet that has `isLearningEnabled` = `true`.
	// Example: `10.0.3.3`
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// The public IP address of the VNIC, if one is assigned.
	PublicIp *string `mandatory:"false" json:"publicIp"`

	// Whether the source/destination check is disabled on the VNIC.
	// Defaults to `false`, which means the check is performed. For information
	// about why you would skip the source/destination check, see
	// Using a Private IP as a Route Target (https://docs.cloud.oracle.com/Content/Network/Tasks/managingroutetables.htm#privateip).
	// Example: `true`
	SkipSourceDestCheck *bool `mandatory:"false" json:"skipSourceDestCheck"`
}

func (m InternalVnic) String() string {
	return common.PointerString(m)
}

// InternalVnicLifecycleStateEnum Enum with underlying type: string
type InternalVnicLifecycleStateEnum string

// Set of constants representing the allowable values for InternalVnicLifecycleStateEnum
const (
	InternalVnicLifecycleStateProvisioning InternalVnicLifecycleStateEnum = "PROVISIONING"
	InternalVnicLifecycleStateAvailable    InternalVnicLifecycleStateEnum = "AVAILABLE"
	InternalVnicLifecycleStateTerminating  InternalVnicLifecycleStateEnum = "TERMINATING"
	InternalVnicLifecycleStateTerminated   InternalVnicLifecycleStateEnum = "TERMINATED"
)

var mappingInternalVnicLifecycleState = map[string]InternalVnicLifecycleStateEnum{
	"PROVISIONING": InternalVnicLifecycleStateProvisioning,
	"AVAILABLE":    InternalVnicLifecycleStateAvailable,
	"TERMINATING":  InternalVnicLifecycleStateTerminating,
	"TERMINATED":   InternalVnicLifecycleStateTerminated,
}

// GetInternalVnicLifecycleStateEnumValues Enumerates the set of values for InternalVnicLifecycleStateEnum
func GetInternalVnicLifecycleStateEnumValues() []InternalVnicLifecycleStateEnum {
	values := make([]InternalVnicLifecycleStateEnum, 0)
	for _, v := range mappingInternalVnicLifecycleState {
		values = append(values, v)
	}
	return values
}
