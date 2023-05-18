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

// PrivateAccessGateway Required for Oracle services that offer customers private endpoints for private access to the
// service.
// The service VCN requires a private access gateway (PAG) to handle the traffic to and from
// private endpoints in customer VCNs (see PrivateEndpoint).
// After creating the gateway, update the route tables in your service VCN to send all traffic
// destined for private endpoints to this gateway.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type PrivateAccessGateway struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the PAG.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the PAG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service VCN that the PAG belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The date and time the PAG was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The PAG's current lifecycle state.
	LifecycleState PrivateAccessGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m PrivateAccessGateway) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PrivateAccessGateway) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPrivateAccessGatewayLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetPrivateAccessGatewayLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PrivateAccessGatewayLifecycleStateEnum Enum with underlying type: string
type PrivateAccessGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for PrivateAccessGatewayLifecycleStateEnum
const (
	PrivateAccessGatewayLifecycleStateProvisioning PrivateAccessGatewayLifecycleStateEnum = "PROVISIONING"
	PrivateAccessGatewayLifecycleStateAvailable    PrivateAccessGatewayLifecycleStateEnum = "AVAILABLE"
	PrivateAccessGatewayLifecycleStateTerminating  PrivateAccessGatewayLifecycleStateEnum = "TERMINATING"
	PrivateAccessGatewayLifecycleStateTerminated   PrivateAccessGatewayLifecycleStateEnum = "TERMINATED"
	PrivateAccessGatewayLifecycleStateUpdating     PrivateAccessGatewayLifecycleStateEnum = "UPDATING"
	PrivateAccessGatewayLifecycleStateFailed       PrivateAccessGatewayLifecycleStateEnum = "FAILED"
)

var mappingPrivateAccessGatewayLifecycleStateEnum = map[string]PrivateAccessGatewayLifecycleStateEnum{
	"PROVISIONING": PrivateAccessGatewayLifecycleStateProvisioning,
	"AVAILABLE":    PrivateAccessGatewayLifecycleStateAvailable,
	"TERMINATING":  PrivateAccessGatewayLifecycleStateTerminating,
	"TERMINATED":   PrivateAccessGatewayLifecycleStateTerminated,
	"UPDATING":     PrivateAccessGatewayLifecycleStateUpdating,
	"FAILED":       PrivateAccessGatewayLifecycleStateFailed,
}

var mappingPrivateAccessGatewayLifecycleStateEnumLowerCase = map[string]PrivateAccessGatewayLifecycleStateEnum{
	"provisioning": PrivateAccessGatewayLifecycleStateProvisioning,
	"available":    PrivateAccessGatewayLifecycleStateAvailable,
	"terminating":  PrivateAccessGatewayLifecycleStateTerminating,
	"terminated":   PrivateAccessGatewayLifecycleStateTerminated,
	"updating":     PrivateAccessGatewayLifecycleStateUpdating,
	"failed":       PrivateAccessGatewayLifecycleStateFailed,
}

// GetPrivateAccessGatewayLifecycleStateEnumValues Enumerates the set of values for PrivateAccessGatewayLifecycleStateEnum
func GetPrivateAccessGatewayLifecycleStateEnumValues() []PrivateAccessGatewayLifecycleStateEnum {
	values := make([]PrivateAccessGatewayLifecycleStateEnum, 0)
	for _, v := range mappingPrivateAccessGatewayLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetPrivateAccessGatewayLifecycleStateEnumStringValues Enumerates the set of values in String for PrivateAccessGatewayLifecycleStateEnum
func GetPrivateAccessGatewayLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
		"UPDATING",
		"FAILED",
	}
}

// GetMappingPrivateAccessGatewayLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPrivateAccessGatewayLifecycleStateEnum(val string) (PrivateAccessGatewayLifecycleStateEnum, bool) {
	enum, ok := mappingPrivateAccessGatewayLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
