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

// CreateInternalEcmpGroupDetails Details to create an internal ecmp group.
type CreateInternalEcmpGroupDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the ecmp group
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// They type of the gateway owning InternalEcmpGroup.
	GatewayType CreateInternalEcmpGroupDetailsGatewayTypeEnum `mandatory:"true" json:"gatewayType"`

	// List of nextHopEntries consisting TargetIps with the associated weight.
	NextHopEntries []NextHopEntryDetails `mandatory:"true" json:"nextHopEntries"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`
}

func (m CreateInternalEcmpGroupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInternalEcmpGroupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCreateInternalEcmpGroupDetailsGatewayTypeEnum(string(m.GatewayType)); !ok && m.GatewayType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for GatewayType: %s. Supported values are: %s.", m.GatewayType, strings.Join(GetCreateInternalEcmpGroupDetailsGatewayTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInternalEcmpGroupDetailsGatewayTypeEnum Enum with underlying type: string
type CreateInternalEcmpGroupDetailsGatewayTypeEnum string

// Set of constants representing the allowable values for CreateInternalEcmpGroupDetailsGatewayTypeEnum
const (
	CreateInternalEcmpGroupDetailsGatewayTypeServicegateway       CreateInternalEcmpGroupDetailsGatewayTypeEnum = "SERVICEGATEWAY"
	CreateInternalEcmpGroupDetailsGatewayTypeNatgateway           CreateInternalEcmpGroupDetailsGatewayTypeEnum = "NATGATEWAY"
	CreateInternalEcmpGroupDetailsGatewayTypePrivateaccessgateway CreateInternalEcmpGroupDetailsGatewayTypeEnum = "PRIVATEACCESSGATEWAY"
)

var mappingCreateInternalEcmpGroupDetailsGatewayTypeEnum = map[string]CreateInternalEcmpGroupDetailsGatewayTypeEnum{
	"SERVICEGATEWAY":       CreateInternalEcmpGroupDetailsGatewayTypeServicegateway,
	"NATGATEWAY":           CreateInternalEcmpGroupDetailsGatewayTypeNatgateway,
	"PRIVATEACCESSGATEWAY": CreateInternalEcmpGroupDetailsGatewayTypePrivateaccessgateway,
}

var mappingCreateInternalEcmpGroupDetailsGatewayTypeEnumLowerCase = map[string]CreateInternalEcmpGroupDetailsGatewayTypeEnum{
	"servicegateway":       CreateInternalEcmpGroupDetailsGatewayTypeServicegateway,
	"natgateway":           CreateInternalEcmpGroupDetailsGatewayTypeNatgateway,
	"privateaccessgateway": CreateInternalEcmpGroupDetailsGatewayTypePrivateaccessgateway,
}

// GetCreateInternalEcmpGroupDetailsGatewayTypeEnumValues Enumerates the set of values for CreateInternalEcmpGroupDetailsGatewayTypeEnum
func GetCreateInternalEcmpGroupDetailsGatewayTypeEnumValues() []CreateInternalEcmpGroupDetailsGatewayTypeEnum {
	values := make([]CreateInternalEcmpGroupDetailsGatewayTypeEnum, 0)
	for _, v := range mappingCreateInternalEcmpGroupDetailsGatewayTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInternalEcmpGroupDetailsGatewayTypeEnumStringValues Enumerates the set of values in String for CreateInternalEcmpGroupDetailsGatewayTypeEnum
func GetCreateInternalEcmpGroupDetailsGatewayTypeEnumStringValues() []string {
	return []string{
		"SERVICEGATEWAY",
		"NATGATEWAY",
		"PRIVATEACCESSGATEWAY",
	}
}

// GetMappingCreateInternalEcmpGroupDetailsGatewayTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInternalEcmpGroupDetailsGatewayTypeEnum(val string) (CreateInternalEcmpGroupDetailsGatewayTypeEnum, bool) {
	enum, ok := mappingCreateInternalEcmpGroupDetailsGatewayTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
