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

// CreateNatGatewayDetails The representation of CreateNatGatewayDetails
type CreateNatGatewayDetails struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the
	// NAT gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the gateway belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Whether the NAT gateway blocks traffic through it. The default is `false`.
	// Example: `true`
	BlockTraffic *bool `mandatory:"false" json:"blockTraffic"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the public IP address associated with the NAT gateway.
	PublicIpId *string `mandatory:"false" json:"publicIpId"`

	// The name of the Oracle managed public IP Pool from which the IP address associated with the NAT gateway is allocated.
	InternalPublicIpPoolName CreateNatGatewayDetailsInternalPublicIpPoolNameEnum `mandatory:"false" json:"internalPublicIpPoolName,omitempty"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table used by the NAT gateway.
	// If you don't specify a route table here, the NAT gateway is created without an associated route
	// table. The Networking service does NOT automatically associate the attached VCN's default route table
	// with the NAT gateway.
	RouteTableId *string `mandatory:"false" json:"routeTableId"`
}

func (m CreateNatGatewayDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateNatGatewayDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnum(string(m.InternalPublicIpPoolName)); !ok && m.InternalPublicIpPoolName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for InternalPublicIpPoolName: %s. Supported values are: %s.", m.InternalPublicIpPoolName, strings.Join(GetCreateNatGatewayDetailsInternalPublicIpPoolNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateNatGatewayDetailsInternalPublicIpPoolNameEnum Enum with underlying type: string
type CreateNatGatewayDetailsInternalPublicIpPoolNameEnum string

// Set of constants representing the allowable values for CreateNatGatewayDetailsInternalPublicIpPoolNameEnum
const (
	CreateNatGatewayDetailsInternalPublicIpPoolNameExternal   CreateNatGatewayDetailsInternalPublicIpPoolNameEnum = "EXTERNAL"
	CreateNatGatewayDetailsInternalPublicIpPoolNameSociEgress CreateNatGatewayDetailsInternalPublicIpPoolNameEnum = "SOCI_EGRESS"
)

var mappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnum = map[string]CreateNatGatewayDetailsInternalPublicIpPoolNameEnum{
	"EXTERNAL":    CreateNatGatewayDetailsInternalPublicIpPoolNameExternal,
	"SOCI_EGRESS": CreateNatGatewayDetailsInternalPublicIpPoolNameSociEgress,
}

var mappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnumLowerCase = map[string]CreateNatGatewayDetailsInternalPublicIpPoolNameEnum{
	"external":    CreateNatGatewayDetailsInternalPublicIpPoolNameExternal,
	"soci_egress": CreateNatGatewayDetailsInternalPublicIpPoolNameSociEgress,
}

// GetCreateNatGatewayDetailsInternalPublicIpPoolNameEnumValues Enumerates the set of values for CreateNatGatewayDetailsInternalPublicIpPoolNameEnum
func GetCreateNatGatewayDetailsInternalPublicIpPoolNameEnumValues() []CreateNatGatewayDetailsInternalPublicIpPoolNameEnum {
	values := make([]CreateNatGatewayDetailsInternalPublicIpPoolNameEnum, 0)
	for _, v := range mappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateNatGatewayDetailsInternalPublicIpPoolNameEnumStringValues Enumerates the set of values in String for CreateNatGatewayDetailsInternalPublicIpPoolNameEnum
func GetCreateNatGatewayDetailsInternalPublicIpPoolNameEnumStringValues() []string {
	return []string{
		"EXTERNAL",
		"SOCI_EGRESS",
	}
}

// GetMappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnum(val string) (CreateNatGatewayDetailsInternalPublicIpPoolNameEnum, bool) {
	enum, ok := mappingCreateNatGatewayDetailsInternalPublicIpPoolNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
