// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// Use the Core Services API to manage resources such as virtual cloud networks (VCNs),
// compute instances, and block storage volumes. For more information, see the console
// documentation for the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services.
// The required permissions are documented in the
// Details for the Core Services (https://docs.cloud.oracle.com/iaas/Content/Identity/Reference/corepolicyreference.htm) article.
//

package core

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternalPrivateIp An internal private IP
type InternalPrivateIp struct {

	// Unique identifier of a floating private IP
	Id *string `mandatory:"true" json:"id"`

	// ID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The IP address of the floating private IP
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The current state of the floating private IP
	State InternalPrivateIpStateEnum `mandatory:"true" json:"state"`

	// ID of the subnet
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// User friendly name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// HostName for the Floating Private IP. Only the hostname label, not the FQDN.
	HostnameLabel *string `mandatory:"false" json:"hostnameLabel"`

	// Auto-delete floating private IP when VNIC is deleted, will auto-unmap when VNIC is deleted regardless of this setting
	DeleteOnVnicDelete *bool `mandatory:"false" json:"deleteOnVnicDelete"`

	// The internal system using this IP, if any
	InternalUseByName *string `mandatory:"false" json:"internalUseByName"`

	Mapping *InternalPrivateIpMapping `mandatory:"false" json:"mapping"`

	// The OCID of the VLAN that the FloatingPrivateIP belongs to
	VlanId *string `mandatory:"false" json:"vlanId"`

	// Creation time of the entity
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Whether this private IP is the primary one on the VNIC.
	// Primary private IPs are unassigned and deleted automatically when the VNIC is terminated.
	// Example: `true`
	IsPrimary *bool `mandatory:"false" json:"isPrimary"`

	// The internal name of the private IP's Availability Domain.
	InternalAvailabilityDomain *string `mandatory:"false" json:"internalAvailabilityDomain"`

	// The VNIC's OCID.
	VnicId *string `mandatory:"false" json:"vnicId"`

	// State of the IP address. If an IP address is assigned to a VNIC it is ASSIGNED otherwise AVAILABLE
	IpState InternalPrivateIpIpStateEnum `mandatory:"false" json:"ipState,omitempty"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime InternalPrivateIpLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`
}

func (m InternalPrivateIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalPrivateIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalPrivateIpStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetInternalPrivateIpStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingInternalPrivateIpIpStateEnum(string(m.IpState)); !ok && m.IpState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IpState: %s. Supported values are: %s.", m.IpState, strings.Join(GetInternalPrivateIpIpStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalPrivateIpLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetInternalPrivateIpLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalPrivateIpStateEnum Enum with underlying type: string
type InternalPrivateIpStateEnum string

// Set of constants representing the allowable values for InternalPrivateIpStateEnum
const (
	InternalPrivateIpStateProvisioning InternalPrivateIpStateEnum = "PROVISIONING"
	InternalPrivateIpStateAvailable    InternalPrivateIpStateEnum = "AVAILABLE"
	InternalPrivateIpStateTerminating  InternalPrivateIpStateEnum = "TERMINATING"
	InternalPrivateIpStateTerminated   InternalPrivateIpStateEnum = "TERMINATED"
)

var mappingInternalPrivateIpStateEnum = map[string]InternalPrivateIpStateEnum{
	"PROVISIONING": InternalPrivateIpStateProvisioning,
	"AVAILABLE":    InternalPrivateIpStateAvailable,
	"TERMINATING":  InternalPrivateIpStateTerminating,
	"TERMINATED":   InternalPrivateIpStateTerminated,
}

var mappingInternalPrivateIpStateEnumLowerCase = map[string]InternalPrivateIpStateEnum{
	"provisioning": InternalPrivateIpStateProvisioning,
	"available":    InternalPrivateIpStateAvailable,
	"terminating":  InternalPrivateIpStateTerminating,
	"terminated":   InternalPrivateIpStateTerminated,
}

// GetInternalPrivateIpStateEnumValues Enumerates the set of values for InternalPrivateIpStateEnum
func GetInternalPrivateIpStateEnumValues() []InternalPrivateIpStateEnum {
	values := make([]InternalPrivateIpStateEnum, 0)
	for _, v := range mappingInternalPrivateIpStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalPrivateIpStateEnumStringValues Enumerates the set of values in String for InternalPrivateIpStateEnum
func GetInternalPrivateIpStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalPrivateIpStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalPrivateIpStateEnum(val string) (InternalPrivateIpStateEnum, bool) {
	enum, ok := mappingInternalPrivateIpStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalPrivateIpIpStateEnum Enum with underlying type: string
type InternalPrivateIpIpStateEnum string

// Set of constants representing the allowable values for InternalPrivateIpIpStateEnum
const (
	InternalPrivateIpIpStateAssigned  InternalPrivateIpIpStateEnum = "ASSIGNED"
	InternalPrivateIpIpStateAvailable InternalPrivateIpIpStateEnum = "AVAILABLE"
)

var mappingInternalPrivateIpIpStateEnum = map[string]InternalPrivateIpIpStateEnum{
	"ASSIGNED":  InternalPrivateIpIpStateAssigned,
	"AVAILABLE": InternalPrivateIpIpStateAvailable,
}

var mappingInternalPrivateIpIpStateEnumLowerCase = map[string]InternalPrivateIpIpStateEnum{
	"assigned":  InternalPrivateIpIpStateAssigned,
	"available": InternalPrivateIpIpStateAvailable,
}

// GetInternalPrivateIpIpStateEnumValues Enumerates the set of values for InternalPrivateIpIpStateEnum
func GetInternalPrivateIpIpStateEnumValues() []InternalPrivateIpIpStateEnum {
	values := make([]InternalPrivateIpIpStateEnum, 0)
	for _, v := range mappingInternalPrivateIpIpStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalPrivateIpIpStateEnumStringValues Enumerates the set of values in String for InternalPrivateIpIpStateEnum
func GetInternalPrivateIpIpStateEnumStringValues() []string {
	return []string{
		"ASSIGNED",
		"AVAILABLE",
	}
}

// GetMappingInternalPrivateIpIpStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalPrivateIpIpStateEnum(val string) (InternalPrivateIpIpStateEnum, bool) {
	enum, ok := mappingInternalPrivateIpIpStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalPrivateIpLifetimeEnum Enum with underlying type: string
type InternalPrivateIpLifetimeEnum string

// Set of constants representing the allowable values for InternalPrivateIpLifetimeEnum
const (
	InternalPrivateIpLifetimeEphemeral InternalPrivateIpLifetimeEnum = "EPHEMERAL"
	InternalPrivateIpLifetimeReserved  InternalPrivateIpLifetimeEnum = "RESERVED"
)

var mappingInternalPrivateIpLifetimeEnum = map[string]InternalPrivateIpLifetimeEnum{
	"EPHEMERAL": InternalPrivateIpLifetimeEphemeral,
	"RESERVED":  InternalPrivateIpLifetimeReserved,
}

var mappingInternalPrivateIpLifetimeEnumLowerCase = map[string]InternalPrivateIpLifetimeEnum{
	"ephemeral": InternalPrivateIpLifetimeEphemeral,
	"reserved":  InternalPrivateIpLifetimeReserved,
}

// GetInternalPrivateIpLifetimeEnumValues Enumerates the set of values for InternalPrivateIpLifetimeEnum
func GetInternalPrivateIpLifetimeEnumValues() []InternalPrivateIpLifetimeEnum {
	values := make([]InternalPrivateIpLifetimeEnum, 0)
	for _, v := range mappingInternalPrivateIpLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalPrivateIpLifetimeEnumStringValues Enumerates the set of values in String for InternalPrivateIpLifetimeEnum
func GetInternalPrivateIpLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingInternalPrivateIpLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalPrivateIpLifetimeEnum(val string) (InternalPrivateIpLifetimeEnum, bool) {
	enum, ok := mappingInternalPrivateIpLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
