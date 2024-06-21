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

// CreateInternalInternetGatewayDetails The representation of CreateInternalInternetGatewayDetails
type CreateInternalInternetGatewayDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the internet gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Whether the gateway is enabled upon creation.
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the Internet Gateway is attached to.
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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the Internet Gateway is using.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// * **FailOverToInternet:** This is the default fail over policy which allows traffic
	// to flow over internet when the back bone fails.
	// * **NoFailOver:** This policy disables any fail over when the back bone fails.
	BackboneFailOverPolicy CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum `mandatory:"false" json:"backboneFailOverPolicy,omitempty"`
}

func (m CreateInternalInternetGatewayDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalInternetGatewayDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum(string(m.BackboneFailOverPolicy)); !ok && m.BackboneFailOverPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackboneFailOverPolicy: %s. Supported values are: %s.", m.BackboneFailOverPolicy, strings.Join(GetCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum Enum with underlying type: string
type CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum string

// Set of constants representing the allowable values for CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
const (
	CreateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = "FAILOVERTOINTERNET"
	CreateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover         CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = "NOFAILOVER"
)

var mappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = map[string]CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum{
	"FAILOVERTOINTERNET": CreateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet,
	"NOFAILOVER":         CreateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover,
}

var mappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumLowerCase = map[string]CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum{
	"failovertointernet": CreateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet,
	"nofailover":         CreateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover,
}

// GetCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumValues Enumerates the set of values for CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
func GetCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumValues() []CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum {
	values := make([]CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum, 0)
	for _, v := range mappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues Enumerates the set of values in String for CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
func GetCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues() []string {
	return []string{
		"FAILOVERTOINTERNET",
		"NOFAILOVER",
	}
}

// GetMappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum(val string) (CreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum, bool) {
	enum, ok := mappingCreateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
