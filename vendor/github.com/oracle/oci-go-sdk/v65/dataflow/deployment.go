// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Flow API
//
// Use the Data Flow APIs to run any Apache Spark application at any scale without deploying or managing any infrastructure.
//

package dataflow

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// Deployment Deployment response.
type Deployment struct {

	// The provision identifier that is immutable on creation.
	ComputeClusterId *string `mandatory:"true" json:"computeClusterId"`

	// The unique identifier of a compute cluster context.
	Id *string `mandatory:"true" json:"id"`

	// Time of object creation
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Endpoint url.
	EndpointUrl *string `mandatory:"true" json:"endpointUrl"`

	// Status of deployment.
	LifecycleState DeploymentLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Deployment type
	Type DeploymentTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Details about the deployment
	LifecycleStateDetails *string `mandatory:"false" json:"lifecycleStateDetails"`

	// The unique identifier of a async execution on this deployment
	ExecutionId *string `mandatory:"false" json:"executionId"`

	// Additional configuration passed to the running process.
	// Example: { "agentFlowKey" : "AF1, "endpointKey" : "EK1" }
	Configuration *string `mandatory:"false" json:"configuration"`

	OauthConfiguration *OAuthConfiguration `mandatory:"false" json:"oauthConfiguration"`
}

func (m Deployment) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Deployment) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDeploymentLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetDeploymentLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingDeploymentTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetDeploymentTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DeploymentLifecycleStateEnum Enum with underlying type: string
type DeploymentLifecycleStateEnum string

// Set of constants representing the allowable values for DeploymentLifecycleStateEnum
const (
	DeploymentLifecycleStateCreating DeploymentLifecycleStateEnum = "CREATING"
	DeploymentLifecycleStateActive   DeploymentLifecycleStateEnum = "ACTIVE"
	DeploymentLifecycleStateDeleting DeploymentLifecycleStateEnum = "DELETING"
	DeploymentLifecycleStateFailed   DeploymentLifecycleStateEnum = "FAILED"
)

var mappingDeploymentLifecycleStateEnum = map[string]DeploymentLifecycleStateEnum{
	"CREATING": DeploymentLifecycleStateCreating,
	"ACTIVE":   DeploymentLifecycleStateActive,
	"DELETING": DeploymentLifecycleStateDeleting,
	"FAILED":   DeploymentLifecycleStateFailed,
}

var mappingDeploymentLifecycleStateEnumLowerCase = map[string]DeploymentLifecycleStateEnum{
	"creating": DeploymentLifecycleStateCreating,
	"active":   DeploymentLifecycleStateActive,
	"deleting": DeploymentLifecycleStateDeleting,
	"failed":   DeploymentLifecycleStateFailed,
}

// GetDeploymentLifecycleStateEnumValues Enumerates the set of values for DeploymentLifecycleStateEnum
func GetDeploymentLifecycleStateEnumValues() []DeploymentLifecycleStateEnum {
	values := make([]DeploymentLifecycleStateEnum, 0)
	for _, v := range mappingDeploymentLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetDeploymentLifecycleStateEnumStringValues Enumerates the set of values in String for DeploymentLifecycleStateEnum
func GetDeploymentLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"ACTIVE",
		"DELETING",
		"FAILED",
	}
}

// GetMappingDeploymentLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDeploymentLifecycleStateEnum(val string) (DeploymentLifecycleStateEnum, bool) {
	enum, ok := mappingDeploymentLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
