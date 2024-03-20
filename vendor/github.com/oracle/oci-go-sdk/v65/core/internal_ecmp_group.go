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

// InternalEcmpGroup An ECMP Group is a list of next hops that are referenced by multiple ECMP routes. Each next hop has a targetIp and associated weight.
type InternalEcmpGroup struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the InternalEcmpGroup.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The InternalEcmpGroup's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The InternalEcmpGroup's current state.
	LifecycleState InternalEcmpGroupLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The type of the gateway owning InternalEcmpGroup.
	GatewayType InternalEcmpGroupGatewayTypeEnum `mandatory:"true" json:"gatewayType"`

	// List of nextHopEntries consisting TargetIps with the associated weight.
	NextHopEntries []NextHopEntryDetails `mandatory:"true" json:"nextHopEntries"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The date and time the InternalEcmpGroup was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m InternalEcmpGroup) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalEcmpGroup) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalEcmpGroupLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalEcmpGroupLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingInternalEcmpGroupGatewayTypeEnum(string(m.GatewayType)); !ok && m.GatewayType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GatewayType: %s. Supported values are: %s.", m.GatewayType, strings.Join(GetInternalEcmpGroupGatewayTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalEcmpGroupLifecycleStateEnum Enum with underlying type: string
type InternalEcmpGroupLifecycleStateEnum string

// Set of constants representing the allowable values for InternalEcmpGroupLifecycleStateEnum
const (
	InternalEcmpGroupLifecycleStateProvisioning InternalEcmpGroupLifecycleStateEnum = "PROVISIONING"
	InternalEcmpGroupLifecycleStateAvailable    InternalEcmpGroupLifecycleStateEnum = "AVAILABLE"
	InternalEcmpGroupLifecycleStateTerminating  InternalEcmpGroupLifecycleStateEnum = "TERMINATING"
	InternalEcmpGroupLifecycleStateTerminated   InternalEcmpGroupLifecycleStateEnum = "TERMINATED"
	InternalEcmpGroupLifecycleStateUpdating     InternalEcmpGroupLifecycleStateEnum = "UPDATING"
)

var mappingInternalEcmpGroupLifecycleStateEnum = map[string]InternalEcmpGroupLifecycleStateEnum{
	"PROVISIONING": InternalEcmpGroupLifecycleStateProvisioning,
	"AVAILABLE":    InternalEcmpGroupLifecycleStateAvailable,
	"TERMINATING":  InternalEcmpGroupLifecycleStateTerminating,
	"TERMINATED":   InternalEcmpGroupLifecycleStateTerminated,
	"UPDATING":     InternalEcmpGroupLifecycleStateUpdating,
}

var mappingInternalEcmpGroupLifecycleStateEnumLowerCase = map[string]InternalEcmpGroupLifecycleStateEnum{
	"provisioning": InternalEcmpGroupLifecycleStateProvisioning,
	"available":    InternalEcmpGroupLifecycleStateAvailable,
	"terminating":  InternalEcmpGroupLifecycleStateTerminating,
	"terminated":   InternalEcmpGroupLifecycleStateTerminated,
	"updating":     InternalEcmpGroupLifecycleStateUpdating,
}

// GetInternalEcmpGroupLifecycleStateEnumValues Enumerates the set of values for InternalEcmpGroupLifecycleStateEnum
func GetInternalEcmpGroupLifecycleStateEnumValues() []InternalEcmpGroupLifecycleStateEnum {
	values := make([]InternalEcmpGroupLifecycleStateEnum, 0)
	for _, v := range mappingInternalEcmpGroupLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalEcmpGroupLifecycleStateEnumStringValues Enumerates the set of values in String for InternalEcmpGroupLifecycleStateEnum
func GetInternalEcmpGroupLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
	}
}

// GetMappingInternalEcmpGroupLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalEcmpGroupLifecycleStateEnum(val string) (InternalEcmpGroupLifecycleStateEnum, bool) {
	enum, ok := mappingInternalEcmpGroupLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalEcmpGroupGatewayTypeEnum Enum with underlying type: string
type InternalEcmpGroupGatewayTypeEnum string

// Set of constants representing the allowable values for InternalEcmpGroupGatewayTypeEnum
const (
	InternalEcmpGroupGatewayTypeServicegateway       InternalEcmpGroupGatewayTypeEnum = "SERVICEGATEWAY"
	InternalEcmpGroupGatewayTypeNatgateway           InternalEcmpGroupGatewayTypeEnum = "NATGATEWAY"
	InternalEcmpGroupGatewayTypePrivateaccessgateway InternalEcmpGroupGatewayTypeEnum = "PRIVATEACCESSGATEWAY"
)

var mappingInternalEcmpGroupGatewayTypeEnum = map[string]InternalEcmpGroupGatewayTypeEnum{
	"SERVICEGATEWAY":       InternalEcmpGroupGatewayTypeServicegateway,
	"NATGATEWAY":           InternalEcmpGroupGatewayTypeNatgateway,
	"PRIVATEACCESSGATEWAY": InternalEcmpGroupGatewayTypePrivateaccessgateway,
}

var mappingInternalEcmpGroupGatewayTypeEnumLowerCase = map[string]InternalEcmpGroupGatewayTypeEnum{
	"servicegateway":       InternalEcmpGroupGatewayTypeServicegateway,
	"natgateway":           InternalEcmpGroupGatewayTypeNatgateway,
	"privateaccessgateway": InternalEcmpGroupGatewayTypePrivateaccessgateway,
}

// GetInternalEcmpGroupGatewayTypeEnumValues Enumerates the set of values for InternalEcmpGroupGatewayTypeEnum
func GetInternalEcmpGroupGatewayTypeEnumValues() []InternalEcmpGroupGatewayTypeEnum {
	values := make([]InternalEcmpGroupGatewayTypeEnum, 0)
	for _, v := range mappingInternalEcmpGroupGatewayTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalEcmpGroupGatewayTypeEnumStringValues Enumerates the set of values in String for InternalEcmpGroupGatewayTypeEnum
func GetInternalEcmpGroupGatewayTypeEnumStringValues() []string {
	return []string{
		"SERVICEGATEWAY",
		"NATGATEWAY",
		"PRIVATEACCESSGATEWAY",
	}
}

// GetMappingInternalEcmpGroupGatewayTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalEcmpGroupGatewayTypeEnum(val string) (InternalEcmpGroupGatewayTypeEnum, bool) {
	enum, ok := mappingInternalEcmpGroupGatewayTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
