// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Core Services API
//
// APIs for Networking Service, Compute Service, and Block Volume Service.
//

package core

import (
	"github.com/oracle/oci-go-sdk/common"
)

// ServiceGateway A Service Gateway.
// Service Gateway is a Gateway that interfaces a Virtual Cloud Network (VCN) and multiple Service networks. With
// the use of Service Gateway, public Services can be reached without the need to traverse through the Internet.
// The list of OCI Services that are supported by Service Gateway is the
// following: Object Storage
// For more information, see Service Gateway Overview (https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Concepts/servicegateway.htm).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/policygetstarted.htm).
type ServiceGateway struct {

	// Boolean to allow/disallow traffic through Service Gateway. This will be False by default
	BlockTraffic *bool `mandatory:"true" json:"blockTraffic"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment containing the Service Gateway.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Service Gateway's Oracle ID ([OCID])(/Content/General/Concepts/identifiers.htm).
	Id *string `mandatory:"true" json:"id"`

	// The Service Gateway's current state.
	LifecycleState ServiceGatewayLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// List of objects of Service OCID and name. These are the Services which have been enabled on the Service Gateway.
	Services []ServiceIdResponseDetails `mandatory:"true" json:"services"`

	// The OCID (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no
	// predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// The date and time the Service Gateway was created, in the format defined by RFC3339.
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m ServiceGateway) String() string {
	return common.PointerString(m)
}

// ServiceGatewayLifecycleStateEnum Enum with underlying type: string
type ServiceGatewayLifecycleStateEnum string

// Set of constants representing the allowable values for ServiceGatewayLifecycleState
const (
	ServiceGatewayLifecycleStateCreating ServiceGatewayLifecycleStateEnum = "CREATING"
	ServiceGatewayLifecycleStateActive   ServiceGatewayLifecycleStateEnum = "ACTIVE"
	ServiceGatewayLifecycleStateFailed   ServiceGatewayLifecycleStateEnum = "FAILED"
	ServiceGatewayLifecycleStateDeleting ServiceGatewayLifecycleStateEnum = "DELETING"
	ServiceGatewayLifecycleStateDeleted  ServiceGatewayLifecycleStateEnum = "DELETED"
)

var mappingServiceGatewayLifecycleState = map[string]ServiceGatewayLifecycleStateEnum{
	"CREATING": ServiceGatewayLifecycleStateCreating,
	"ACTIVE":   ServiceGatewayLifecycleStateActive,
	"FAILED":   ServiceGatewayLifecycleStateFailed,
	"DELETING": ServiceGatewayLifecycleStateDeleting,
	"DELETED":  ServiceGatewayLifecycleStateDeleted,
}

// GetServiceGatewayLifecycleStateEnumValues Enumerates the set of values for ServiceGatewayLifecycleState
func GetServiceGatewayLifecycleStateEnumValues() []ServiceGatewayLifecycleStateEnum {
	values := make([]ServiceGatewayLifecycleStateEnum, 0)
	for _, v := range mappingServiceGatewayLifecycleState {
		values = append(values, v)
	}
	return values
}
