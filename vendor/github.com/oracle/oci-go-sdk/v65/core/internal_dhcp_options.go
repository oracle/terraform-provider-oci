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

// InternalDhcpOptions A set of DHCP options.
type InternalDhcpOptions struct {

	// The OCID of the compartment containing the set of DHCP options.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Oracle ID (OCID) for the set of DHCP options.
	Id *string `mandatory:"true" json:"id"`

	// The current state of the set of DHCP options.
	State InternalDhcpOptionsStateEnum `mandatory:"true" json:"state"`

	// Date and time the set of DHCP options was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The collection of individual DHCP options.
	Options []InternalDhcpOption `mandatory:"true" json:"options"`

	// A user-friendly name. Does not have to be unique, and it's unchangeable.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Date and time the set of DHCP options was last modified (same as `TimeCreated`
	// if not yet modified).
	ModifiedTime *common.SDKTime `mandatory:"false" json:"modifiedTime"`
}

func (m InternalDhcpOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalDhcpOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalDhcpOptionsStateEnum(string(m.State)); !ok && m.State != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for State: %s. Supported values are: %s.", m.State, strings.Join(GetInternalDhcpOptionsStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalDhcpOptionsStateEnum Enum with underlying type: string
type InternalDhcpOptionsStateEnum string

// Set of constants representing the allowable values for InternalDhcpOptionsStateEnum
const (
	InternalDhcpOptionsStateProvisioning InternalDhcpOptionsStateEnum = "PROVISIONING"
	InternalDhcpOptionsStateAvailable    InternalDhcpOptionsStateEnum = "AVAILABLE"
	InternalDhcpOptionsStateTerminating  InternalDhcpOptionsStateEnum = "TERMINATING"
	InternalDhcpOptionsStateTerminated   InternalDhcpOptionsStateEnum = "TERMINATED"
)

var mappingInternalDhcpOptionsStateEnum = map[string]InternalDhcpOptionsStateEnum{
	"PROVISIONING": InternalDhcpOptionsStateProvisioning,
	"AVAILABLE":    InternalDhcpOptionsStateAvailable,
	"TERMINATING":  InternalDhcpOptionsStateTerminating,
	"TERMINATED":   InternalDhcpOptionsStateTerminated,
}

var mappingInternalDhcpOptionsStateEnumLowerCase = map[string]InternalDhcpOptionsStateEnum{
	"provisioning": InternalDhcpOptionsStateProvisioning,
	"available":    InternalDhcpOptionsStateAvailable,
	"terminating":  InternalDhcpOptionsStateTerminating,
	"terminated":   InternalDhcpOptionsStateTerminated,
}

// GetInternalDhcpOptionsStateEnumValues Enumerates the set of values for InternalDhcpOptionsStateEnum
func GetInternalDhcpOptionsStateEnumValues() []InternalDhcpOptionsStateEnum {
	values := make([]InternalDhcpOptionsStateEnum, 0)
	for _, v := range mappingInternalDhcpOptionsStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalDhcpOptionsStateEnumStringValues Enumerates the set of values in String for InternalDhcpOptionsStateEnum
func GetInternalDhcpOptionsStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalDhcpOptionsStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalDhcpOptionsStateEnum(val string) (InternalDhcpOptionsStateEnum, bool) {
	enum, ok := mappingInternalDhcpOptionsStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
