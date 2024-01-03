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

// InternalIpv6 An *IPv6* is a conceptual term that refers to an IPv6 address and related properties.
// The `IPv6` object is the API representation of an IPv6.
// IPv6 can be created and assigned to any particular VNICs which are inside an IPv6 enabled subnet.
type InternalIpv6 struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment containing the IPv6.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A user-friendly name. Does not have to be unique, and it's changeable. Avoid
	// entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The IPv6's OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The IPv6 address of the `IPv6` object. The address is within the IPv6 prefix
	// of the VNIC's subnet.
	// Example: `2001:0db8:0123:4567:89ab:cdef:1234:5678`
	IpAddress *string `mandatory:"true" json:"ipAddress"`

	// The IPv6's current state.
	LifecycleState InternalIpv6LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID of the subnet the VNIC is in.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The date and time the IPv6 was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The OCID of the VNIC the IPv6 is assigned to. The VNIC and IPv6 must be in the same subnet.
	VnicId *string `mandatory:"true" json:"vnicId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The hostname of the IPv6 address. Only the hostname label, not the FQDN.
	Hostname *string `mandatory:"false" json:"hostname"`

	// Whether IPv6 is usable for intenet communication. Internet access via IPv6 will not be allowed for
	// private subnet the same way as IPv4. Internet access will be enabled by default for a public subnet.
	// If VCN has IPv6 enabled with a custom IPv6 prefix, a different public IPv6 address will be assigned
	// for a particular IPv6.
	IsInternetAccessAllowed *bool `mandatory:"false" json:"isInternetAccessAllowed"`

	// The IPv6 address to be used for an internet communication. The address is within the IPv6 public prefix
	// of the VNIC's subnet. The `publicIpAddress` is always the same `ipAddress` if the VCN does not have
	// a custom IPv6 prefix (i.e. use Oracle provided prefix). If VCN has a custom IPv6 prefix, the `publicIpAddress`
	// is drawn from the subnet public IPv6 prefix by translating the IPv6 address prefix to the public IPv6 prefix.
	// If `publicIpAddress` is not available, internet access is not permitted for this particular IPv6.
	PublicIpAddress *string `mandatory:"false" json:"publicIpAddress"`
}

func (m InternalIpv6) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalIpv6) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalIpv6LifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalIpv6LifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalIpv6LifecycleStateEnum Enum with underlying type: string
type InternalIpv6LifecycleStateEnum string

// Set of constants representing the allowable values for InternalIpv6LifecycleStateEnum
const (
	InternalIpv6LifecycleStateProvisioning InternalIpv6LifecycleStateEnum = "PROVISIONING"
	InternalIpv6LifecycleStateAvailable    InternalIpv6LifecycleStateEnum = "AVAILABLE"
	InternalIpv6LifecycleStateTerminating  InternalIpv6LifecycleStateEnum = "TERMINATING"
	InternalIpv6LifecycleStateTerminated   InternalIpv6LifecycleStateEnum = "TERMINATED"
)

var mappingInternalIpv6LifecycleStateEnum = map[string]InternalIpv6LifecycleStateEnum{
	"PROVISIONING": InternalIpv6LifecycleStateProvisioning,
	"AVAILABLE":    InternalIpv6LifecycleStateAvailable,
	"TERMINATING":  InternalIpv6LifecycleStateTerminating,
	"TERMINATED":   InternalIpv6LifecycleStateTerminated,
}

var mappingInternalIpv6LifecycleStateEnumLowerCase = map[string]InternalIpv6LifecycleStateEnum{
	"provisioning": InternalIpv6LifecycleStateProvisioning,
	"available":    InternalIpv6LifecycleStateAvailable,
	"terminating":  InternalIpv6LifecycleStateTerminating,
	"terminated":   InternalIpv6LifecycleStateTerminated,
}

// GetInternalIpv6LifecycleStateEnumValues Enumerates the set of values for InternalIpv6LifecycleStateEnum
func GetInternalIpv6LifecycleStateEnumValues() []InternalIpv6LifecycleStateEnum {
	values := make([]InternalIpv6LifecycleStateEnum, 0)
	for _, v := range mappingInternalIpv6LifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalIpv6LifecycleStateEnumStringValues Enumerates the set of values in String for InternalIpv6LifecycleStateEnum
func GetInternalIpv6LifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalIpv6LifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalIpv6LifecycleStateEnum(val string) (InternalIpv6LifecycleStateEnum, bool) {
	enum, ok := mappingInternalIpv6LifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
