// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// IpAnycast Global Anycast also known as IP Anycast or Anycast routing, is an IP network addressing scheme that allows multiple servers to share the same IP address, allowing for multiple physical destination servers to be logically identified by a single IP address.
type IpAnycast struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the `IpAnycast` resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the `IpAnycast` resource.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The `IpAnycast` resource's current state.
	LifecycleState IpAnycastLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

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

	// Brief summary about the details of Anycast resource being created.
	Description *string `mandatory:"false" json:"description"`

	// The list of BYOIP Range OCIDs resources created for BYOIP Prefix.
	ByoipRangeIds []string `mandatory:"false" json:"byoipRangeIds"`

	// The IP CIDR block being imported from on-premises to the Oracle cloud.
	Prefix *string `mandatory:"false" json:"prefix"`

	// The IP address of the CIDR for Prefix Monitoring.
	MonitorIp *string `mandatory:"false" json:"monitorIp"`

	// The date and time the `IpAnycast` resource was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the `IpAnycast` resource was last updated, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m IpAnycast) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m IpAnycast) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingIpAnycastLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetIpAnycastLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// IpAnycastLifecycleStateEnum Enum with underlying type: string
type IpAnycastLifecycleStateEnum string

// Set of constants representing the allowable values for IpAnycastLifecycleStateEnum
const (
	IpAnycastLifecycleStateActive  IpAnycastLifecycleStateEnum = "ACTIVE"
	IpAnycastLifecycleStateDeleted IpAnycastLifecycleStateEnum = "DELETED"
)

var mappingIpAnycastLifecycleStateEnum = map[string]IpAnycastLifecycleStateEnum{
	"ACTIVE":  IpAnycastLifecycleStateActive,
	"DELETED": IpAnycastLifecycleStateDeleted,
}

var mappingIpAnycastLifecycleStateEnumLowerCase = map[string]IpAnycastLifecycleStateEnum{
	"active":  IpAnycastLifecycleStateActive,
	"deleted": IpAnycastLifecycleStateDeleted,
}

// GetIpAnycastLifecycleStateEnumValues Enumerates the set of values for IpAnycastLifecycleStateEnum
func GetIpAnycastLifecycleStateEnumValues() []IpAnycastLifecycleStateEnum {
	values := make([]IpAnycastLifecycleStateEnum, 0)
	for _, v := range mappingIpAnycastLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetIpAnycastLifecycleStateEnumStringValues Enumerates the set of values in String for IpAnycastLifecycleStateEnum
func GetIpAnycastLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETED",
	}
}

// GetMappingIpAnycastLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingIpAnycastLifecycleStateEnum(val string) (IpAnycastLifecycleStateEnum, bool) {
	enum, ok := mappingIpAnycastLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
