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

// UpdateInternalPrivateIpDetails Details to update internal private ip
type UpdateInternalPrivateIpDetails struct {

	// User friendly name
	DisplayName *string `mandatory:"false" json:"displayName"`

	// HostName for the Floating Private IP. Only the hostname label, not the FQDN.
	HostNameLabel *string `mandatory:"false" json:"hostNameLabel"`

	// Lifetime of the IP address.
	// There are two types of IPv6 IPs:
	//  - Ephemeral
	//  - Reserved
	Lifetime UpdateInternalPrivateIpDetailsLifetimeEnum `mandatory:"false" json:"lifetime,omitempty"`

	// Auto-delete floating private IP when VNIC is deleted (will auto-detach regardless of this setting)
	DeleteOnVnicDelete *bool `mandatory:"false" json:"deleteOnVnicDelete"`
}

func (m UpdateInternalPrivateIpDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalPrivateIpDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalPrivateIpDetailsLifetimeEnum(string(m.Lifetime)); !ok && m.Lifetime != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Lifetime: %s. Supported values are: %s.", m.Lifetime, strings.Join(GetUpdateInternalPrivateIpDetailsLifetimeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalPrivateIpDetailsLifetimeEnum Enum with underlying type: string
type UpdateInternalPrivateIpDetailsLifetimeEnum string

// Set of constants representing the allowable values for UpdateInternalPrivateIpDetailsLifetimeEnum
const (
	UpdateInternalPrivateIpDetailsLifetimeEphemeral UpdateInternalPrivateIpDetailsLifetimeEnum = "EPHEMERAL"
	UpdateInternalPrivateIpDetailsLifetimeReserved  UpdateInternalPrivateIpDetailsLifetimeEnum = "RESERVED"
)

var mappingUpdateInternalPrivateIpDetailsLifetimeEnum = map[string]UpdateInternalPrivateIpDetailsLifetimeEnum{
	"EPHEMERAL": UpdateInternalPrivateIpDetailsLifetimeEphemeral,
	"RESERVED":  UpdateInternalPrivateIpDetailsLifetimeReserved,
}

var mappingUpdateInternalPrivateIpDetailsLifetimeEnumLowerCase = map[string]UpdateInternalPrivateIpDetailsLifetimeEnum{
	"ephemeral": UpdateInternalPrivateIpDetailsLifetimeEphemeral,
	"reserved":  UpdateInternalPrivateIpDetailsLifetimeReserved,
}

// GetUpdateInternalPrivateIpDetailsLifetimeEnumValues Enumerates the set of values for UpdateInternalPrivateIpDetailsLifetimeEnum
func GetUpdateInternalPrivateIpDetailsLifetimeEnumValues() []UpdateInternalPrivateIpDetailsLifetimeEnum {
	values := make([]UpdateInternalPrivateIpDetailsLifetimeEnum, 0)
	for _, v := range mappingUpdateInternalPrivateIpDetailsLifetimeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalPrivateIpDetailsLifetimeEnumStringValues Enumerates the set of values in String for UpdateInternalPrivateIpDetailsLifetimeEnum
func GetUpdateInternalPrivateIpDetailsLifetimeEnumStringValues() []string {
	return []string{
		"EPHEMERAL",
		"RESERVED",
	}
}

// GetMappingUpdateInternalPrivateIpDetailsLifetimeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalPrivateIpDetailsLifetimeEnum(val string) (UpdateInternalPrivateIpDetailsLifetimeEnum, bool) {
	enum, ok := mappingUpdateInternalPrivateIpDetailsLifetimeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
