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

// CreateComputeClusterDetails The information about a new Compute Cluster.
type CreateComputeClusterDetails struct {

	// The identifier of the compartment used with the Compute Cluster.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The Compute Cluster name, which can be changed.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The minimum number of executors.
	MinExecutorCount *int `mandatory:"true" json:"minExecutorCount"`

	// The maximum number of executors.
	MaxExecutorCount *int `mandatory:"true" json:"maxExecutorCount"`

	// The description of the Compute Cluster.
	Description *string `mandatory:"false" json:"description"`

	// Cluster node type encodes the node shape and associated resources.
	NodeType *string `mandatory:"false" json:"nodeType"`

	// Optional Driver's node type which encodes the Driver node shape and associated resources.
	DriverNodeType *string `mandatory:"false" json:"driverNodeType"`

	// The shape of the Compute Cluster driver instance.
	DriverShape *string `mandatory:"false" json:"driverShape"`

	DriverShapeConfig *ShapeConfig `mandatory:"false" json:"driverShapeConfig"`

	// The shape of the Compute Cluster executor instance.
	ExecutorShape *string `mandatory:"false" json:"executorShape"`

	ExecutorShapeConfig *ShapeConfig `mandatory:"false" json:"executorShapeConfig"`

	// The OCID of OCI Lake.
	LakeId *string `mandatory:"false" json:"lakeId"`

	// The OCID of OCI Lake.
	UserTenantId *string `mandatory:"false" json:"userTenantId"`

	// The OCID of OCI Lake.
	UserCompartmentId *string `mandatory:"false" json:"userCompartmentId"`

	// The OCID of OCI Lake.
	IdlSvcTenantId *string `mandatory:"false" json:"idlSvcTenantId"`

	// The OCID of Hub.
	HubId *string `mandatory:"false" json:"hubId"`

	// The paths to init scripts that will be executed in the order of definition
	InitScripts []string `mandatory:"false" json:"initScripts"`

	// The YAML file paths to import into the Compute Cluster.
	YamlPaths []string `mandatory:"false" json:"yamlPaths"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// The Spark version utilized to run the application.
	SparkVersion *string `mandatory:"false" json:"sparkVersion"`

	// The Spark configuration passed to the running process.
	// See https://spark.apache.org/docs/latest/configuration.html#available-properties.
	// Example: { "spark.app.name" : "My App Name", "spark.shuffle.io.maxRetries" : "4" }
	// Note: Not all Spark properties are permitted to be set.  Attempting to set a property that is
	// not allowed to be overwritten will cause a 400 status to be returned.
	SparkAdvancedConfigurations map[string]string `mandatory:"false" json:"sparkAdvancedConfigurations"`

	// The environment variables passed to the running process.
	// These variables are set inside driver and executor pods for the user application to consume.
	// Example: { "key1" : "value1", "key2" : "value2" }
	EnvironmentVariables map[string]string `mandatory:"false" json:"environmentVariables"`

	// Optional timeout value in minutes used to auto stop Compute Clusters. A cluster will be auto stopped after inactivity for this amount of time period.
	// If this value is not set, the cluster will not be auto stopped.
	IdleComputeClusterTimeoutInMinutes *int `mandatory:"false" json:"idleComputeClusterTimeoutInMinutes"`

	// WorkspaceKey of the cluster.
	WorkspaceKey *string `mandatory:"false" json:"workspaceKey"`

	// HubProxyEndpoint of the cluster.
	HubProxyEndpoint *string `mandatory:"false" json:"hubProxyEndpoint"`

	// List of networkConfigurationIds for the cluster
	NetworkConfigurationIds []string `mandatory:"false" json:"networkConfigurationIds"`

	// Cluster Key of the cluster.
	ClusterKey *string `mandatory:"false" json:"clusterKey"`

	// Specify the logId to publish spark diagnostic logs.
	LogId *string `mandatory:"false" json:"logId"`

	// Async Operation Key for the operation on the cluster.
	AsyncOperationKey *string `mandatory:"false" json:"asyncOperationKey"`

	// Indicates whether this compute cluster should be created as default cluster
	IsDefaultComputeCluster *bool `mandatory:"false" json:"isDefaultComputeCluster"`

	// Compute Cluster Type.
	ClusterType *string `mandatory:"false" json:"clusterType"`

	// Tag slug from user tenancy passed from Datalake as Base64.getEncoder().encodeToString(data)
	UserTenancyTagSlug *string `mandatory:"false" json:"userTenancyTagSlug"`

	// The delegation token type.
	DelegationTokenType DelegationTokenTypeEnum `mandatory:"false" json:"delegationTokenType,omitempty"`

	// The tenant ID of owner service principal.
	OwnerServicePrincipalTenantId *string `mandatory:"false" json:"ownerServicePrincipalTenantId"`

	Subscription []SubscriptionDetails `mandatory:"false" json:"subscription"`

	// URL to give callback.
	CallbackUrl *string `mandatory:"false" json:"callbackUrl"`
}

func (m CreateComputeClusterDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateComputeClusterDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingDelegationTokenTypeEnum(string(m.DelegationTokenType)); !ok && m.DelegationTokenType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DelegationTokenType: %s. Supported values are: %s.", m.DelegationTokenType, strings.Join(GetDelegationTokenTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
