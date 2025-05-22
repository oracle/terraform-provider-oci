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

// CreatePrivateIpNextHopDetails The data to create private IP's nextHop configuration.
type CreatePrivateIpNextHopDetails struct {

	// Details of nextHop targets.
	Targets []PrivateIpNextHopTarget `mandatory:"true" json:"targets"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// VNICaaS will flow-hash traffic that matches the service protocol and port.
	ServiceProtocolPorts []PrivateIpNextHopProtocolPort `mandatory:"false" json:"serviceProtocolPorts"`

	// Turns on/off flow stickiness for the private IP's nextHop. The default is 'false'.
	IsFlowStickinessEnabled *bool `mandatory:"false" json:"isFlowStickinessEnabled"`

	// Hashing config to use while selecting target
	// DEFAULT_HASHING: The choice of hashing method is determined by the dataplane and may evolve over time based on implementation decisions
	// MODULAR_HASHING: Uses modular hashing
	// CONSISTENT_HASHING: Uses consistent hashing
	TargetSelectionHashConfig CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum `mandatory:"false" json:"targetSelectionHashConfig,omitempty"`

	// Forwarding configuration for a private IP's nextHop. The default is 'DEFAULT'.
	// DEFAULT: Default behavior where packets are flow hashed to a range of ports.
	// SKIP_PORT_SHARDING: Packets will skip port sharding.
	// SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER: Packets will skip port sharding and a wildcard listener will be used.
	// RCE_PROXY: Packets will be forwarded to RCE
	NextHopForwardingConfig CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum `mandatory:"false" json:"nextHopForwardingConfig,omitempty"`
}

func (m CreatePrivateIpNextHopDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePrivateIpNextHopDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum(string(m.TargetSelectionHashConfig)); !ok && m.TargetSelectionHashConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetSelectionHashConfig: %s. Supported values are: %s.", m.TargetSelectionHashConfig, strings.Join(GetCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum(string(m.NextHopForwardingConfig)); !ok && m.NextHopForwardingConfig != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for NextHopForwardingConfig: %s. Supported values are: %s.", m.NextHopForwardingConfig, strings.Join(GetCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum Enum with underlying type: string
type CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum string

// Set of constants representing the allowable values for CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
const (
	CreatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "DEFAULT_HASHING"
	CreatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "MODULAR_HASHING"
	CreatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = "CONSISTENT_HASHING"
)

var mappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum = map[string]CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum{
	"DEFAULT_HASHING":    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing,
	"MODULAR_HASHING":    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing,
	"CONSISTENT_HASHING": CreatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing,
}

var mappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumLowerCase = map[string]CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum{
	"default_hashing":    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigDefaultHashing,
	"modular_hashing":    CreatePrivateIpNextHopDetailsTargetSelectionHashConfigModularHashing,
	"consistent_hashing": CreatePrivateIpNextHopDetailsTargetSelectionHashConfigConsistentHashing,
}

// GetCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumValues Enumerates the set of values for CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
func GetCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumValues() []CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum {
	values := make([]CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum, 0)
	for _, v := range mappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues Enumerates the set of values in String for CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum
func GetCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumStringValues() []string {
	return []string{
		"DEFAULT_HASHING",
		"MODULAR_HASHING",
		"CONSISTENT_HASHING",
	}
}

// GetMappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum(val string) (CreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnum, bool) {
	enum, ok := mappingCreatePrivateIpNextHopDetailsTargetSelectionHashConfigEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum Enum with underlying type: string
type CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum string

// Set of constants representing the allowable values for CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
const (
	CreatePrivateIpNextHopDetailsNextHopForwardingConfigDefault                              CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "DEFAULT"
	CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding                     CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "SKIP_PORT_SHARDING"
	CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER"
	CreatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy                             CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = "RCE_PROXY"
)

var mappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum = map[string]CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum{
	"DEFAULT":            CreatePrivateIpNextHopDetailsNextHopForwardingConfigDefault,
	"SKIP_PORT_SHARDING": CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding,
	"SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER": CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener,
	"RCE_PROXY": CreatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy,
}

var mappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumLowerCase = map[string]CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum{
	"default":            CreatePrivateIpNextHopDetailsNextHopForwardingConfigDefault,
	"skip_port_sharding": CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortSharding,
	"skip_port_sharding_with_wildcard_listener": CreatePrivateIpNextHopDetailsNextHopForwardingConfigSkipPortShardingWithWildcardListener,
	"rce_proxy": CreatePrivateIpNextHopDetailsNextHopForwardingConfigRceProxy,
}

// GetCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumValues Enumerates the set of values for CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
func GetCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumValues() []CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum {
	values := make([]CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum, 0)
	for _, v := range mappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum {
		values = append(values, v)
	}
	return values
}

// GetCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues Enumerates the set of values in String for CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum
func GetCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumStringValues() []string {
	return []string{
		"DEFAULT",
		"SKIP_PORT_SHARDING",
		"SKIP_PORT_SHARDING_WITH_WILDCARD_LISTENER",
		"RCE_PROXY",
	}
}

// GetMappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum(val string) (CreatePrivateIpNextHopDetailsNextHopForwardingConfigEnum, bool) {
	enum, ok := mappingCreatePrivateIpNextHopDetailsNextHopForwardingConfigEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
