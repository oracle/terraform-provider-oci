// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Monitoring API
//
// Use the Network Monitoring API to troubleshoot routing and security issues for resources such as virtual cloud networks (VCNs) and compute instances. For more information, see the console
// documentation for the Network Path Analyzer (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/path_analyzer.htm) tool.
//

package vnmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PublicIpPool A public IP pool is a set of public IP addresses represented as one or more IPv4 CIDR blocks.      Resources like load balancers and compute instances can be allocated public IP addresses from a public IP pool.
type PublicIpPool struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing this pool.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP pool.
	Id *string `mandatory:"true" json:"id"`

	// The date and time the public IP pool was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The CIDR blocks added to this pool. This could be all or a portion of a BYOIP CIDR block.
	CidrBlocks []string `mandatory:"false" json:"cidrBlocks"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The public IP pool's current state.
	LifecycleState PublicIpPoolLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

func (m PublicIpPool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PublicIpPool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPublicIpPoolLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPublicIpPoolLifecycleStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PublicIpPoolLifecycleStateEnum Enum with underlying type: string
type PublicIpPoolLifecycleStateEnum string

// Set of constants representing the allowable values for PublicIpPoolLifecycleStateEnum
const (
	PublicIpPoolLifecycleStateInactive PublicIpPoolLifecycleStateEnum = "INACTIVE"
	PublicIpPoolLifecycleStateUpdating PublicIpPoolLifecycleStateEnum = "UPDATING"
	PublicIpPoolLifecycleStateActive   PublicIpPoolLifecycleStateEnum = "ACTIVE"
	PublicIpPoolLifecycleStateDeleting PublicIpPoolLifecycleStateEnum = "DELETING"
	PublicIpPoolLifecycleStateDeleted  PublicIpPoolLifecycleStateEnum = "DELETED"
)

var mappingPublicIpPoolLifecycleStateEnum = map[string]PublicIpPoolLifecycleStateEnum{
	"INACTIVE": PublicIpPoolLifecycleStateInactive,
	"UPDATING": PublicIpPoolLifecycleStateUpdating,
	"ACTIVE":   PublicIpPoolLifecycleStateActive,
	"DELETING": PublicIpPoolLifecycleStateDeleting,
	"DELETED":  PublicIpPoolLifecycleStateDeleted,
}

var mappingPublicIpPoolLifecycleStateEnumLowerCase = map[string]PublicIpPoolLifecycleStateEnum{
	"inactive": PublicIpPoolLifecycleStateInactive,
	"updating": PublicIpPoolLifecycleStateUpdating,
	"active":   PublicIpPoolLifecycleStateActive,
	"deleting": PublicIpPoolLifecycleStateDeleting,
	"deleted":  PublicIpPoolLifecycleStateDeleted,
}

// GetPublicIpPoolLifecycleStateEnumValues Enumerates the set of values for PublicIpPoolLifecycleStateEnum
func GetPublicIpPoolLifecycleStateEnumValues() []PublicIpPoolLifecycleStateEnum {
	values := make([]PublicIpPoolLifecycleStateEnum, 0)
	for _, v := range mappingPublicIpPoolLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPublicIpPoolLifecycleStateEnumStringValues Enumerates the set of values in String for PublicIpPoolLifecycleStateEnum
func GetPublicIpPoolLifecycleStateEnumStringValues() []string {
	return []string{
		"INACTIVE",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingPublicIpPoolLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPublicIpPoolLifecycleStateEnum(val string) (PublicIpPoolLifecycleStateEnum, bool) {
	enum, ok := mappingPublicIpPoolLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
