// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Core Services API
//
// API covering the Networking (https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/overview.htm),
// Compute (https://docs.cloud.oracle.com/iaas/Content/Compute/Concepts/computeoverview.htm), and
// Block Volume (https://docs.cloud.oracle.com/iaas/Content/Block/Concepts/overview.htm) services. Use this API
// to manage resources such as virtual cloud networks (VCNs), compute instances, and
// block storage volumes.
//

package core

import (
	"github.com/oracle/oci-go-sdk/v25/common"
)

// PrivateAccessGateway Required for Oracle services that offer customers private endpoints for private access to the
// service.
// The service VCN requires a private access gateway (PAG) to handle the traffic to and from
// private endpoints in customer VCNs (see PrivateEndpoint).
// After creating the gateway, update the route tables in your service VCN to send all traffic
// destined for private endpoints to this gateway.
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/Content/Identity/Concepts/policygetstarted.htm).
// **Warning:** Oracle recommends that you avoid using any confidential information when you
// supply string values using the API.
type PrivateAccessGateway struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the PAG.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment that contains the PAG.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the service VCN that the PAG belongs to.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The date and time the PAG was created, in the format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The PAG's current lifecycle state.
	LifecycleState PrivateAccessGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// A user-friendly name. Does not have to be unique.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`
}

func (m PrivateAccessGateway) String() string {
	return common.PointerString(m)
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

var mappingPrivateAccessGatewayLifecycleState = map[string]PrivateAccessGatewayLifecycleStateEnum{
	"PROVISIONING": PrivateAccessGatewayLifecycleStateProvisioning,
	"AVAILABLE":    PrivateAccessGatewayLifecycleStateAvailable,
	"TERMINATING":  PrivateAccessGatewayLifecycleStateTerminating,
	"TERMINATED":   PrivateAccessGatewayLifecycleStateTerminated,
	"UPDATING":     PrivateAccessGatewayLifecycleStateUpdating,
	"FAILED":       PrivateAccessGatewayLifecycleStateFailed,
}

// GetPrivateAccessGatewayLifecycleStateEnumValues Enumerates the set of values for PrivateAccessGatewayLifecycleStateEnum
func GetPrivateAccessGatewayLifecycleStateEnumValues() []PrivateAccessGatewayLifecycleStateEnum {
	values := make([]PrivateAccessGatewayLifecycleStateEnum, 0)
	for _, v := range mappingPrivateAccessGatewayLifecycleState {
		values = append(values, v)
	}
	return values
}
