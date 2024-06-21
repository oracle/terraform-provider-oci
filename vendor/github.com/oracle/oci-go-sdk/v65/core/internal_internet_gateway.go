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

// InternalInternetGateway Represents a router that connects the edge of a VCN with the Internet. For an example scenario
// that uses an internet gateway, see
// Typical Networking Service Scenarios (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm#scenarios).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type InternalInternetGateway struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the internet gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The internet gateway's Oracle ID (OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
	Id *string `mandatory:"true" json:"id"`

	// The internet gateway's current state.
	LifecycleState InternalInternetGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the Internet Gateway belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Security Attributes for this resource. This is unique to ZPR, and helps identify which resources are allowed to be accessed by what permission controls.
	// Example: `{"Oracle-DataSecurity-ZPR": {"MaxEgressCount": {"value":"42","mode":"audit"}}}`
	SecurityAttributes map[string]map[string]interface{} `mandatory:"false" json:"securityAttributes"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether the gateway is enabled. When the gateway is disabled, traffic is not
	// routed to/from the Internet, regardless of route rules.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The date and time the internet gateway was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the Internet Gateway is using.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// * **FailOverToInternet:** This is the default fail over policy which allows traffic
	// to flow over internet when the back bone fails.
	// * **NoFailOver:** This disables any fail over policy when the the back bone fails.
	BackboneFailOverPolicy InternalInternetGatewayBackboneFailOverPolicyEnum `mandatory:"false" json:"backboneFailOverPolicy,omitempty"`
}

func (m InternalInternetGateway) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternalInternetGateway) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInternalInternetGatewayLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetInternalInternetGatewayLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingInternalInternetGatewayBackboneFailOverPolicyEnum(string(m.BackboneFailOverPolicy)); !ok && m.BackboneFailOverPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackboneFailOverPolicy: %s. Supported values are: %s.", m.BackboneFailOverPolicy, strings.Join(GetInternalInternetGatewayBackboneFailOverPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InternalInternetGatewayLifecycleStateEnum Enum with underlying type: string
type InternalInternetGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for InternalInternetGatewayLifecycleStateEnum
const (
	InternalInternetGatewayLifecycleStateProvisioning InternalInternetGatewayLifecycleStateEnum = "PROVISIONING"
	InternalInternetGatewayLifecycleStateAvailable    InternalInternetGatewayLifecycleStateEnum = "AVAILABLE"
	InternalInternetGatewayLifecycleStateTerminating  InternalInternetGatewayLifecycleStateEnum = "TERMINATING"
	InternalInternetGatewayLifecycleStateTerminated   InternalInternetGatewayLifecycleStateEnum = "TERMINATED"
)

var mappingInternalInternetGatewayLifecycleStateEnum = map[string]InternalInternetGatewayLifecycleStateEnum{
	"PROVISIONING": InternalInternetGatewayLifecycleStateProvisioning,
	"AVAILABLE":    InternalInternetGatewayLifecycleStateAvailable,
	"TERMINATING":  InternalInternetGatewayLifecycleStateTerminating,
	"TERMINATED":   InternalInternetGatewayLifecycleStateTerminated,
}

var mappingInternalInternetGatewayLifecycleStateEnumLowerCase = map[string]InternalInternetGatewayLifecycleStateEnum{
	"provisioning": InternalInternetGatewayLifecycleStateProvisioning,
	"available":    InternalInternetGatewayLifecycleStateAvailable,
	"terminating":  InternalInternetGatewayLifecycleStateTerminating,
	"terminated":   InternalInternetGatewayLifecycleStateTerminated,
}

// GetInternalInternetGatewayLifecycleStateEnumValues Enumerates the set of values for InternalInternetGatewayLifecycleStateEnum
func GetInternalInternetGatewayLifecycleStateEnumValues() []InternalInternetGatewayLifecycleStateEnum {
	values := make([]InternalInternetGatewayLifecycleStateEnum, 0)
	for _, v := range mappingInternalInternetGatewayLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalInternetGatewayLifecycleStateEnumStringValues Enumerates the set of values in String for InternalInternetGatewayLifecycleStateEnum
func GetInternalInternetGatewayLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingInternalInternetGatewayLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalInternetGatewayLifecycleStateEnum(val string) (InternalInternetGatewayLifecycleStateEnum, bool) {
	enum, ok := mappingInternalInternetGatewayLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// InternalInternetGatewayBackboneFailOverPolicyEnum Enum with underlying type: string
type InternalInternetGatewayBackboneFailOverPolicyEnum string

// Set of constants representing the allowable values for InternalInternetGatewayBackboneFailOverPolicyEnum
const (
	InternalInternetGatewayBackboneFailOverPolicyFailovertointernet InternalInternetGatewayBackboneFailOverPolicyEnum = "FAILOVERTOINTERNET"
	InternalInternetGatewayBackboneFailOverPolicyNofailover         InternalInternetGatewayBackboneFailOverPolicyEnum = "NOFAILOVER"
)

var mappingInternalInternetGatewayBackboneFailOverPolicyEnum = map[string]InternalInternetGatewayBackboneFailOverPolicyEnum{
	"FAILOVERTOINTERNET": InternalInternetGatewayBackboneFailOverPolicyFailovertointernet,
	"NOFAILOVER":         InternalInternetGatewayBackboneFailOverPolicyNofailover,
}

var mappingInternalInternetGatewayBackboneFailOverPolicyEnumLowerCase = map[string]InternalInternetGatewayBackboneFailOverPolicyEnum{
	"failovertointernet": InternalInternetGatewayBackboneFailOverPolicyFailovertointernet,
	"nofailover":         InternalInternetGatewayBackboneFailOverPolicyNofailover,
}

// GetInternalInternetGatewayBackboneFailOverPolicyEnumValues Enumerates the set of values for InternalInternetGatewayBackboneFailOverPolicyEnum
func GetInternalInternetGatewayBackboneFailOverPolicyEnumValues() []InternalInternetGatewayBackboneFailOverPolicyEnum {
	values := make([]InternalInternetGatewayBackboneFailOverPolicyEnum, 0)
	for _, v := range mappingInternalInternetGatewayBackboneFailOverPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetInternalInternetGatewayBackboneFailOverPolicyEnumStringValues Enumerates the set of values in String for InternalInternetGatewayBackboneFailOverPolicyEnum
func GetInternalInternetGatewayBackboneFailOverPolicyEnumStringValues() []string {
	return []string{
		"FAILOVERTOINTERNET",
		"NOFAILOVER",
	}
}

// GetMappingInternalInternetGatewayBackboneFailOverPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInternalInternetGatewayBackboneFailOverPolicyEnum(val string) (InternalInternetGatewayBackboneFailOverPolicyEnum, bool) {
	enum, ok := mappingInternalInternetGatewayBackboneFailOverPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
