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

// ReverseConnectionNatIp The current NAT IP address that corresponds to a specific customer IP address.
// To establish a reverse connection to a customer IP address, use the NAT IP
// address as the destination.
type ReverseConnectionNatIp struct {

	// The reverse connection NAT IP's current state.
	LifecycleState ReverseConnectionNatIpLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the reverse connection NAT IP was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The customer's IP address that corresponds to the reverse connection NAT IP address.
	ReverseConnectionCustomerIp *string `mandatory:"true" json:"reverseConnectionCustomerIp"`

	// The reverse connection NAT IP address corresonding to the customer IP and private endpoint.
	// Use this address as the destination when establishing a reverse connection to a customer's
	// IP address.
	ReverseConnectionNatIp *string `mandatory:"true" json:"reverseConnectionNatIp"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint
	// associated with the reverse connection.
	PrivateEndpointId *string `mandatory:"true" json:"privateEndpointId"`
}

func (m ReverseConnectionNatIp) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReverseConnectionNatIp) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReverseConnectionNatIpLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetReverseConnectionNatIpLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReverseConnectionNatIpLifecycleStateEnum Enum with underlying type: string
type ReverseConnectionNatIpLifecycleStateEnum string

// Set of constants representing the allowable values for ReverseConnectionNatIpLifecycleStateEnum
const (
	ReverseConnectionNatIpLifecycleStateProvisioning ReverseConnectionNatIpLifecycleStateEnum = "PROVISIONING"
	ReverseConnectionNatIpLifecycleStateAvailable    ReverseConnectionNatIpLifecycleStateEnum = "AVAILABLE"
	ReverseConnectionNatIpLifecycleStateTerminating  ReverseConnectionNatIpLifecycleStateEnum = "TERMINATING"
	ReverseConnectionNatIpLifecycleStateTerminated   ReverseConnectionNatIpLifecycleStateEnum = "TERMINATED"
)

var mappingReverseConnectionNatIpLifecycleStateEnum = map[string]ReverseConnectionNatIpLifecycleStateEnum{
	"PROVISIONING": ReverseConnectionNatIpLifecycleStateProvisioning,
	"AVAILABLE":    ReverseConnectionNatIpLifecycleStateAvailable,
	"TERMINATING":  ReverseConnectionNatIpLifecycleStateTerminating,
	"TERMINATED":   ReverseConnectionNatIpLifecycleStateTerminated,
}

var mappingReverseConnectionNatIpLifecycleStateEnumLowerCase = map[string]ReverseConnectionNatIpLifecycleStateEnum{
	"provisioning": ReverseConnectionNatIpLifecycleStateProvisioning,
	"available":    ReverseConnectionNatIpLifecycleStateAvailable,
	"terminating":  ReverseConnectionNatIpLifecycleStateTerminating,
	"terminated":   ReverseConnectionNatIpLifecycleStateTerminated,
}

// GetReverseConnectionNatIpLifecycleStateEnumValues Enumerates the set of values for ReverseConnectionNatIpLifecycleStateEnum
func GetReverseConnectionNatIpLifecycleStateEnumValues() []ReverseConnectionNatIpLifecycleStateEnum {
	values := make([]ReverseConnectionNatIpLifecycleStateEnum, 0)
	for _, v := range mappingReverseConnectionNatIpLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetReverseConnectionNatIpLifecycleStateEnumStringValues Enumerates the set of values in String for ReverseConnectionNatIpLifecycleStateEnum
func GetReverseConnectionNatIpLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingReverseConnectionNatIpLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReverseConnectionNatIpLifecycleStateEnum(val string) (ReverseConnectionNatIpLifecycleStateEnum, bool) {
	enum, ok := mappingReverseConnectionNatIpLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
