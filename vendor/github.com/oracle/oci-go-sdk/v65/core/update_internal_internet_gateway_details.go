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

// UpdateInternalInternetGatewayDetails The representation of UpdateInternalInternetGatewayDetails
type UpdateInternalInternetGatewayDetails struct {

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

	// Whether the gateway is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the Internet Gateway is using.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// * **FailOverToInternet:** This is the default fail over policy which allows traffic
	// to flow over internet when the back bone fails.
	// * **NoFailOver:** This disables any fail over policy when the back bone fails.
	BackboneFailOverPolicy UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum `mandatory:"false" json:"backboneFailOverPolicy,omitempty"`
}

func (m UpdateInternalInternetGatewayDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInternalInternetGatewayDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum(string(m.BackboneFailOverPolicy)); !ok && m.BackboneFailOverPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackboneFailOverPolicy: %s. Supported values are: %s.", m.BackboneFailOverPolicy, strings.Join(GetUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum Enum with underlying type: string
type UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum string

// Set of constants representing the allowable values for UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
const (
	UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = "FAILOVERTOINTERNET"
	UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover         UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = "NOFAILOVER"
)

var mappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum = map[string]UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum{
	"FAILOVERTOINTERNET": UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet,
	"NOFAILOVER":         UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover,
}

var mappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumLowerCase = map[string]UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum{
	"failovertointernet": UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyFailovertointernet,
	"nofailover":         UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyNofailover,
}

// GetUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumValues Enumerates the set of values for UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
func GetUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumValues() []UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum {
	values := make([]UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum, 0)
	for _, v := range mappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues Enumerates the set of values in String for UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum
func GetUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumStringValues() []string {
	return []string{
		"FAILOVERTOINTERNET",
		"NOFAILOVER",
	}
}

// GetMappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum(val string) (UpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnum, bool) {
	enum, ok := mappingUpdateInternalInternetGatewayDetailsBackboneFailOverPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
