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

// CreateInternalPrivateIpDetails Details to create internal private ip
type CreateInternalPrivateIpDetails struct {

	// The internal system using this IP, if any
	InternalUseByName *string `mandatory:"true" json:"internalUseByName"`

	// ID of the subnet
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// User friendly name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Auto-delete floating private IP when VNIC is deleted (will auto-detach regardless of this setting)
	DeleteOnVnicDelete *bool `mandatory:"false" json:"deleteOnVnicDelete"`

	Mapping *MapInternalPrivateIpDetails `mandatory:"false" json:"mapping"`

	// The OCID of the VLAN that the FloatingPrivateIP belongs to
	VlanId *string `mandatory:"false" json:"vlanId"`

	// The private IP address. Next available IP is selected from available IPs in subnet if none is specified.
	PrivateIp *string `mandatory:"false" json:"privateIp"`

	// HostName for the Floating Private IP. Only the hostname label, not the FQDN.
	HostNameLabel *string `mandatory:"false" json:"hostNameLabel"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime CreateInternalPrivateIpDetailsLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`
}

func (m CreateInternalPrivateIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalPrivateIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalPrivateIpDetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetCreateInternalPrivateIpDetailsLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalPrivateIpDetailsLifetimeEnum Enum with underlying type: string
type CreateInternalPrivateIpDetailsLifetimeEnum string

// Set of constants representing the allowable values for CreateInternalPrivateIpDetailsLifetimeEnum
const (
	CreateInternalPrivateIpDetailsLifetimeEphemeral CreateInternalPrivateIpDetailsLifetimeEnum = "EPHEMERAL"
	CreateInternalPrivateIpDetailsLifetimeReserved  CreateInternalPrivateIpDetailsLifetimeEnum = "RESERVED"
)

var mappingCreateInternalPrivateIpDetailsLifetimeEnum = map[string]CreateInternalPrivateIpDetailsLifetimeEnum{
	"EPHEMERAL": CreateInternalPrivateIpDetailsLifetimeEphemeral,
	"RESERVED":  CreateInternalPrivateIpDetailsLifetimeReserved,
}

var mappingCreateInternalPrivateIpDetailsLifetimeEnumLowerCase = map[string]CreateInternalPrivateIpDetailsLifetimeEnum{
	"ephemeral": CreateInternalPrivateIpDetailsLifetimeEphemeral,
	"reserved":  CreateInternalPrivateIpDetailsLifetimeReserved,
}

// GetCreateInternalPrivateIpDetailsLifetimeEnumValues Enumerates the set of values for CreateInternalPrivateIpDetailsLifetimeEnum
func GetCreateInternalPrivateIpDetailsLifetimeEnumValues() []CreateInternalPrivateIpDetailsLifetimeEnum {
	values := make([]CreateInternalPrivateIpDetailsLifetimeEnum, 0)
	for _, v := range mappingCreateInternalPrivateIpDetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalPrivateIpDetailsLifetimeEnumStringValues Enumerates the set of values in String for CreateInternalPrivateIpDetailsLifetimeEnum
func GetCreateInternalPrivateIpDetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingCreateInternalPrivateIpDetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalPrivateIpDetailsLifetimeEnum(val string) (CreateInternalPrivateIpDetailsLifetimeEnum, bool) {
	enum, ok := mappingCreateInternalPrivateIpDetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
