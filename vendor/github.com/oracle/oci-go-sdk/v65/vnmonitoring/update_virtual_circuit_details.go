// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateVirtualCircuitDetails The representation of UpdateVirtualCircuitDetails
type UpdateVirtualCircuitDetails struct {

	// The provisioned data rate of the connection. To get a list of the
	// available bandwidth levels (that is, shapes), see
	// ListFastConnectProviderVirtualCircuitBandwidthShapes.
	// To be updated only by the customer who owns the virtual circuit.
	BandwidthShapeName *string `mandatory:"false" json:"bandwidthShapeName"`

	// An array of mappings, each containing properties for a cross-connect or
	// cross-connect group associated with this virtual circuit.
	// The customer and provider can update different properties in the mapping
	// depending on the situation. See the description of the
	// CrossConnectMapping.
	CrossConnectMappings []CrossConnectMapping `mandatory:"false" json:"crossConnectMappings"`

	// The BGP ASN of the network at the other end of the BGP
	// session from Oracle.
	// If the BGP session is from the customer's edge router to Oracle, the
	// required value is the customer's ASN, and it can be updated only
	// by the customer.
	// If the BGP session is from the provider's edge router to Oracle, the
	// required value is the provider's ASN, and it can be updated only
	// by the provider.
	CustomerBgpAsn *int `mandatory:"false" json:"customerBgpAsn"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Drg
	// that this private virtual circuit uses.
	// To be updated only by the customer who owns the virtual circuit.
	GatewayId *string `mandatory:"false" json:"gatewayId"`

	// The provider's state in relation to this virtual circuit. Relevant only
	// if the customer is using FastConnect via a provider. ACTIVE
	// means the provider has provisioned the virtual circuit from their
	// end. INACTIVE means the provider has not yet provisioned the virtual
	// circuit, or has de-provisioned it.
	// To be updated only by the provider.
	ProviderState UpdateVirtualCircuitDetailsProviderStateEnum `mandatory:"false" json:"providerState,omitempty"`

	// Provider-supplied reference information about this virtual circuit.
	// Relevant only if the customer is using FastConnect via a provider.
	// To be updated only by the provider.
	ReferenceComment *string `mandatory:"false" json:"referenceComment"`
}

func (m UpdateVirtualCircuitDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVirtualCircuitDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateVirtualCircuitDetailsProviderStateEnum(string(m.ProviderState)); !ok && m.ProviderState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProviderState: %s. Supported values are: %s.", m.ProviderState, strings.Join(GetUpdateVirtualCircuitDetailsProviderStateEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateVirtualCircuitDetailsProviderStateEnum Enum with underlying type: string
type UpdateVirtualCircuitDetailsProviderStateEnum string

// Set of constants representing the allowable values for UpdateVirtualCircuitDetailsProviderStateEnum
const (
	UpdateVirtualCircuitDetailsProviderStateActive   UpdateVirtualCircuitDetailsProviderStateEnum = "ACTIVE"
	UpdateVirtualCircuitDetailsProviderStateInactive UpdateVirtualCircuitDetailsProviderStateEnum = "INACTIVE"
)

var mappingUpdateVirtualCircuitDetailsProviderStateEnum = map[string]UpdateVirtualCircuitDetailsProviderStateEnum{
	"ACTIVE":   UpdateVirtualCircuitDetailsProviderStateActive,
	"INACTIVE": UpdateVirtualCircuitDetailsProviderStateInactive,
}

var mappingUpdateVirtualCircuitDetailsProviderStateEnumLowerCase = map[string]UpdateVirtualCircuitDetailsProviderStateEnum{
	"active":   UpdateVirtualCircuitDetailsProviderStateActive,
	"inactive": UpdateVirtualCircuitDetailsProviderStateInactive,
}

// GetUpdateVirtualCircuitDetailsProviderStateEnumValues Enumerates the set of values for UpdateVirtualCircuitDetailsProviderStateEnum
func GetUpdateVirtualCircuitDetailsProviderStateEnumValues() []UpdateVirtualCircuitDetailsProviderStateEnum {
	values := make([]UpdateVirtualCircuitDetailsProviderStateEnum, 0)
	for _, v := range mappingUpdateVirtualCircuitDetailsProviderStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateVirtualCircuitDetailsProviderStateEnumStringValues Enumerates the set of values in String for UpdateVirtualCircuitDetailsProviderStateEnum
func GetUpdateVirtualCircuitDetailsProviderStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"INACTIVE",
	}
}

// GetMappingUpdateVirtualCircuitDetailsProviderStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateVirtualCircuitDetailsProviderStateEnum(val string) (UpdateVirtualCircuitDetailsProviderStateEnum, bool) {
	enum, ok := mappingUpdateVirtualCircuitDetailsProviderStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
