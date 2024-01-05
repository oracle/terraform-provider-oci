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

// ServiceGateway Represents a router that lets your VCN privately access specific Oracle services such as Object
// Storage without exposing the VCN to the public internet. Traffic leaving the VCN and destined
// for a supported Oracle service (see ListServices) is
// routed through the service gateway and does not traverse the internet. The instances in the VCN
// do not need to have public IP addresses nor be in a public subnet. The VCN does not need an internet gateway
// for this traffic. For more information, see
// Access to Oracle Services: Service Gateway (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/servicegateway.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
type ServiceGateway struct {

	// Whether the service gateway blocks all traffic through it. The default is `false`. When
	// this is `true`, traffic is not routed to any services, regardless of route rules.
	// Example: `true`
	BlockTraffic *bool `mandatory:"true" json:"blockTraffic"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the
	// service gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the service gateway.
	Id *string `mandatory:"true" json:"id"`

	// The service gateway's current state.
	LifecycleState ServiceGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of the Service objects enabled for this service gateway.
	// The list can be empty. You can enable a particular `Service` by using
	// AttachServiceId or
	// UpdateServiceGateway.
	Services []ServiceIdResponseDetails `mandatory:"true" json:"services"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN the service gateway
	// belongs to.
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

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the route table the service gateway is using.
	// For information about why you would associate a route table with a service gateway, see
	// Transit Routing: Private Access to Oracle Services (https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/transitroutingoracleservices.htm).
	RouteTableId *string `mandatory:"false" json:"routeTableId"`

	// The date and time the service gateway was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ServiceGateway) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ServiceGateway) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingServiceGatewayLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetServiceGatewayLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ServiceGatewayLifecycleStateEnum Enum with underlying type: string
type ServiceGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for ServiceGatewayLifecycleStateEnum
const (
	ServiceGatewayLifecycleStateProvisioning ServiceGatewayLifecycleStateEnum = "PROVISIONING"
	ServiceGatewayLifecycleStateAvailable    ServiceGatewayLifecycleStateEnum = "AVAILABLE"
	ServiceGatewayLifecycleStateTerminating  ServiceGatewayLifecycleStateEnum = "TERMINATING"
	ServiceGatewayLifecycleStateTerminated   ServiceGatewayLifecycleStateEnum = "TERMINATED"
)

var mappingServiceGatewayLifecycleStateEnum = map[string]ServiceGatewayLifecycleStateEnum{
	"PROVISIONING": ServiceGatewayLifecycleStateProvisioning,
	"AVAILABLE":    ServiceGatewayLifecycleStateAvailable,
	"TERMINATING":  ServiceGatewayLifecycleStateTerminating,
	"TERMINATED":   ServiceGatewayLifecycleStateTerminated,
}

var mappingServiceGatewayLifecycleStateEnumLowerCase = map[string]ServiceGatewayLifecycleStateEnum{
	"provisioning": ServiceGatewayLifecycleStateProvisioning,
	"available":    ServiceGatewayLifecycleStateAvailable,
	"terminating":  ServiceGatewayLifecycleStateTerminating,
	"terminated":   ServiceGatewayLifecycleStateTerminated,
}

// GetServiceGatewayLifecycleStateEnumValues Enumerates the set of values for ServiceGatewayLifecycleStateEnum
func GetServiceGatewayLifecycleStateEnumValues() []ServiceGatewayLifecycleStateEnum {
	values := make([]ServiceGatewayLifecycleStateEnum, 0)
	for _, v := range mappingServiceGatewayLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetServiceGatewayLifecycleStateEnumStringValues Enumerates the set of values in String for ServiceGatewayLifecycleStateEnum
func GetServiceGatewayLifecycleStateEnumStringValues() []string {
	return []string{
		"PROVISIONING",
		"AVAILABLE",
		"TERMINATING",
		"TERMINATED",
	}
}

// GetMappingServiceGatewayLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingServiceGatewayLifecycleStateEnum(val string) (ServiceGatewayLifecycleStateEnum, bool) {
	enum, ok := mappingServiceGatewayLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
