// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OpenSearch Service API
//
// The OpenSearch service API provides access to OCI Search Service with OpenSearch.
//

package opensearch

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OpensearchClusterPipeline An OpenSearch cluster Pipeline resource. An cluster is set of instances that provide OpenSearch functionality in OCI Search Service with OpenSearch.
// For more information, see Cluster Pipelines (https://docs.cloud.oracle.com/iaas/Content/search-opensearch/Concepts/ociopensearchpipeline.htm).
type OpensearchClusterPipeline struct {

	// The OCID of the cluster pipeline.
	Id *string `mandatory:"true" json:"id"`

	// The name of the pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment where the pipeline is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the pipeline's VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the pipeline's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID for the compartment where the pipeline's VCN is located.
	VcnCompartmentId *string `mandatory:"true" json:"vcnCompartmentId"`

	// The OCID for the compartment where the pipwline's subnet is located.
	SubnetCompartmentId *string `mandatory:"true" json:"subnetCompartmentId"`

	// The maximum pipeline capacity, in OCPUs.
	MaxOcpuCount *int `mandatory:"true" json:"maxOcpuCount"`

	// The maximum pipeline capacity, in OCPUs.
	MinOcpuCount *int `mandatory:"true" json:"minOcpuCount"`

	// The maximum amount of memory in GB, for the pipeline.
	MaxMemoryGB *int `mandatory:"true" json:"maxMemoryGB"`

	// The minimum amount of memory in GB, for the pipeline.
	MinMemoryGB *int `mandatory:"true" json:"minMemoryGB"`

	// The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	PipelineConfigurationBody *string `mandatory:"true" json:"pipelineConfigurationBody"`

	// The fully qualified domain name (FQDN) for the cluster's API endpoint.
	OpensearchPipelineFqdn *string `mandatory:"true" json:"opensearchPipelineFqdn"`

	// The pipeline's private IP address.
	OpensearchPipelinePrivateIp *string `mandatory:"true" json:"opensearchPipelinePrivateIp"`

	// The current state of the cluster backup.
	LifecycleState OpensearchClusterPipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the cluster pipeline was created. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The amount of time in milliseconds since the pipeline was updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m OpensearchClusterPipeline) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchClusterPipeline) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterPipelineLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OpensearchClusterPipelineLifecycleStateEnum Enum with underlying type: string
type OpensearchClusterPipelineLifecycleStateEnum string

// Set of constants representing the allowable values for OpensearchClusterPipelineLifecycleStateEnum
const (
	OpensearchClusterPipelineLifecycleStateCreating OpensearchClusterPipelineLifecycleStateEnum = "CREATING"
	OpensearchClusterPipelineLifecycleStateUpdating OpensearchClusterPipelineLifecycleStateEnum = "UPDATING"
	OpensearchClusterPipelineLifecycleStateActive   OpensearchClusterPipelineLifecycleStateEnum = "ACTIVE"
	OpensearchClusterPipelineLifecycleStateDeleting OpensearchClusterPipelineLifecycleStateEnum = "DELETING"
	OpensearchClusterPipelineLifecycleStateDeleted  OpensearchClusterPipelineLifecycleStateEnum = "DELETED"
	OpensearchClusterPipelineLifecycleStateFailed   OpensearchClusterPipelineLifecycleStateEnum = "FAILED"
)

var mappingOpensearchClusterPipelineLifecycleStateEnum = map[string]OpensearchClusterPipelineLifecycleStateEnum{
	"CREATING": OpensearchClusterPipelineLifecycleStateCreating,
	"UPDATING": OpensearchClusterPipelineLifecycleStateUpdating,
	"ACTIVE":   OpensearchClusterPipelineLifecycleStateActive,
	"DELETING": OpensearchClusterPipelineLifecycleStateDeleting,
	"DELETED":  OpensearchClusterPipelineLifecycleStateDeleted,
	"FAILED":   OpensearchClusterPipelineLifecycleStateFailed,
}

var mappingOpensearchClusterPipelineLifecycleStateEnumLowerCase = map[string]OpensearchClusterPipelineLifecycleStateEnum{
	"creating": OpensearchClusterPipelineLifecycleStateCreating,
	"updating": OpensearchClusterPipelineLifecycleStateUpdating,
	"active":   OpensearchClusterPipelineLifecycleStateActive,
	"deleting": OpensearchClusterPipelineLifecycleStateDeleting,
	"deleted":  OpensearchClusterPipelineLifecycleStateDeleted,
	"failed":   OpensearchClusterPipelineLifecycleStateFailed,
}

// GetOpensearchClusterPipelineLifecycleStateEnumValues Enumerates the set of values for OpensearchClusterPipelineLifecycleStateEnum
func GetOpensearchClusterPipelineLifecycleStateEnumValues() []OpensearchClusterPipelineLifecycleStateEnum {
	values := make([]OpensearchClusterPipelineLifecycleStateEnum, 0)
	for _, v := range mappingOpensearchClusterPipelineLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetOpensearchClusterPipelineLifecycleStateEnumStringValues Enumerates the set of values in String for OpensearchClusterPipelineLifecycleStateEnum
func GetOpensearchClusterPipelineLifecycleStateEnumStringValues() []string {
	return []string{
		"CREATING",
		"UPDATING",
		"ACTIVE",
		"DELETING",
		"DELETED",
		"FAILED",
	}
}

// GetMappingOpensearchClusterPipelineLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOpensearchClusterPipelineLifecycleStateEnum(val string) (OpensearchClusterPipelineLifecycleStateEnum, bool) {
	enum, ok := mappingOpensearchClusterPipelineLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
