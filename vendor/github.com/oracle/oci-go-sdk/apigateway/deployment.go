// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// API Gateway API
//
// API for the API Gateway service. Use this API to manage gateways, deployments, and related items.
// For more information, see
// Overview of API Gateway (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayoverview.htm).
//

package apigateway

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Deployment A deployment deploys an API on a gateway. Avoid entering confidential information.
// For more information, see
// API Gateway Concepts (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Concepts/apigatewayconcepts.htm).
type Deployment struct {

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource.
	GatewayId *string `mandatory:"true" json:"gatewayId"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment in which the
	// resource is created.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// A path on which to deploy all routes contained in the API
	// deployment specification. For more information, see
	// Deploying an API on an API Gateway by Creating an API
	// Deployment (https://docs.cloud.oracle.com/iaas/Content/APIGateway/Tasks/apigatewaycreatingdeployment.htm).
	PathPrefix *string `mandatory:"true" json:"pathPrefix"`

	// The endpoint to access this deployment on the gateway.
	Endpoint *string `mandatory:"true" json:"endpoint"`

	Specification *ApiSpecification `mandatory:"true" json:"specification"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// Example: `My new resource`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The time this resource was created. An RFC3339 formatted datetime string.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The time this resource was last updated. An RFC3339 formatted datetime string.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The current state of the deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message describing the current state in more detail.
	// For example, can be used to provide actionable information for a
	// resource in a Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// Free-form tags for this resource. Each tag is a simple key-value pair
	// with no predefined name, type, or namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see
	// Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m Deployment) String() string {
	return common.PointerString(m)
}

// DeploymentLifecycleStateEnum Enum with underlying type: string
type DeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentLifecycleStateEnum
const (
	DeploymentLifecycleStateCreating DeploymentLifecycleStateEnum = "CREATING"
	DeploymentLifecycleStateActive   DeploymentLifecycleStateEnum = "ACTIVE"
	DeploymentLifecycleStateUpdating DeploymentLifecycleStateEnum = "UPDATING"
	DeploymentLifecycleStateDeleting DeploymentLifecycleStateEnum = "DELETING"
	DeploymentLifecycleStateDeleted  DeploymentLifecycleStateEnum = "DELETED"
	DeploymentLifecycleStateFailed   DeploymentLifecycleStateEnum = "FAILED"
)

var mappingDeploymentLifecycleState = map[string]DeploymentLifecycleStateEnum{
	"CREATING": DeploymentLifecycleStateCreating,
	"ACTIVE":   DeploymentLifecycleStateActive,
	"UPDATING": DeploymentLifecycleStateUpdating,
	"DELETING": DeploymentLifecycleStateDeleting,
	"DELETED":  DeploymentLifecycleStateDeleted,
	"FAILED":   DeploymentLifecycleStateFailed,
}

// GetDeploymentLifecycleStateEnumValues Enumerates the set of values for DeploymentLifecycleStateEnum
func GetDeploymentLifecycleStateEnumValues() []DeploymentLifecycleStateEnum {
	values := make([]DeploymentLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentLifecycleState {
		values = append(values, v)
	}
	return values
}
