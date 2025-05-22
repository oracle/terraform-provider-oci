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

// UpdatePrivateIpNextHopDetails The data to update private IP's nextHop configuration.
type UpdatePrivateIpNextHopDetails struct {

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// VNICaaS will flow-hash traffic that matches the service protocol and port. Sending an
	// empty list mean you want to remove all service protocol ports.
	ServiceProtocolPorts []PrivateIpNextHopProtocolPort `mandatory:"false" json:"serviceProtocolPorts"`

	// Details of nextHop targets.
	Targets []PrivateIpNextHopTarget `mandatory:"false" json:"targets"`

	// Turns on/off flow stickiness for the private IP's nextHop. The default is 'false'.
	IsFlowStickinessEnabled *bool `mandatory:"false" json:"isFlowStickinessEnabled"`

	// Hashing config to use while selecting target
	// DEFAULT_HASHING: The choice of hashing method is determined by the dataplane and may evolve over time based on implementation decisions
	// MODULAR_HASHING: Uses modular hashing
	// CONSISTENT_HASHING: Uses consistent hashing
	TargetSelectionHashConfig UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum `mandatory:"false" json:"targetSelectionHashConfig,omitempty"`

	// Forwarding configuration for a private IP's nextHop. The default is 'DEFAULT'.
	// DEFAULT: Default behavior where packets are flow hashed to a range of ports.
	// SKIP_PORT_SHARDING: Packets will skip port sharding.
	// SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER: Packets will skip port sharding and a wildcard listener will be used.
	// RCE_PROXY: Packets will be forwarded to RCE
	NextHopForwardingConfig UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum `mandatory:"false" json:"nextHopForwardingConfig,omitempty"`
}

func (m UpdatePrivateIpNextHopDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdatePrivateIpNextHopDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum(string(m.TargetSelectionHashConfig)); !ok && m.TargetSelectionHashConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetSelectionHashConfig: %s. Supported values are: %s.", m.TargetSelectionHashConfig, strings.Join(GetUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum(string(m.NextHopForwardingConfig)); !ok && m.NextHopForwardingConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextHopForwardingConfig: %s. Supported values are: %s.", m.NextHopForwardingConfig, strings.Join(GetUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum Enum with underlying type: string
type UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum string

// Set of constants representing the allowable values for UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
const (
	UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "DEFAULT_HASHING"
	UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "MODULAR_HASHING"
	UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "CONSISTENT_HASHING"
)

var mappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = map[string]UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum{
	"DEFAULT_HASHING":    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing,
	"MODULAR_HASHING":    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing,
	"CONSISTENT_HASHING": UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing,
}

var mappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumLowerCase = map[string]UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum{
	"default_hashing":    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing,
	"modular_hashing":    UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing,
	"consistent_hashing": UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing,
}

// GetUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumValues Enumerates the set of values for UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
func GetUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumValues() []UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum {
	values := make([]UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum, 0)
	for _, v := range mappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues Enumerates the set of values in String for UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
func GetUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues() []string {
	return []string{
		"DEFAULT_HASHING",
		"MODULAR_HASHING",
		"CONSISTENT_HASHING",
	}
}

// GetMappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum(val string) (UpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum, bool) {
	enum, ok := mappingUpdatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum Enum with underlying type: string
type UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum string

// Set of constants representing the allowable values for UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
const (
	UpdatePrivateIpNextHopDetailsNextHopForwardingConfigDefault                              UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "DEFAULT"
	UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding                     UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "SKIP_PORT_SHARDING"
	UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER"
	UpdatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy                             UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "RCE_PROXY"
)

var mappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = map[string]UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum{
	"DEFAULT":            UpdatePrivateIpNextHopDetailsNextHopForwardingConfigDefault,
	"SKIP_PORT_SHARDING": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding,
	"SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener,
	"RCE_PROXY": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy,
}

var mappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumLowerCase = map[string]UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum{
	"default":            UpdatePrivateIpNextHopDetailsNextHopForwardingConfigDefault,
	"skip_port_sharding": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding,
	"skip_port_sharding_with_wildcard_listener": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener,
	"rce_proxy": UpdatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy,
}

// GetUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumValues Enumerates the set of values for UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
func GetUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumValues() []UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum {
	values := make([]UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum, 0)
	for _, v := range mappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues Enumerates the set of values in String for UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
func GetUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SKIP_PORT_SHARDING",
		"SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER",
		"RCE_PROXY",
	}
}

// GetMappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum(val string) (UpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnum, bool) {
	enum, ok := mappingUpdatePrivateIpNextHopDetailsNextHopForwardingConfigEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
