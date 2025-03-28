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

// OpensearchClusterPipelineSummary The summary of information about an OpenSearch cluster Pipeline.
type OpensearchClusterPipelineSummary struct {

	// The OCID of the cluster pipeline.
	Id *string `mandatory:"true" json:"id"`

	// The name of the cluster pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID of the compartment where the cluster pipeline is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The number of OCPUs configured for each pipeline node.
	OcpuCount *int `mandatory:"true" json:"ocpuCount"`

	// The amount of memory in GB, for each pipeline node.
	MemoryGB *int `mandatory:"true" json:"memoryGB"`

	// The number of nodes configured for the pipeline.
	NodeCount *int `mandatory:"true" json:"nodeCount"`

	// The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	PipelineConfigurationBody *string `mandatory:"true" json:"pipelineConfigurationBody"`

	// The data prepper config in YAML format. The command accepts the data prepper config as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	DataPrepperConfigurationBody *string `mandatory:"true" json:"dataPrepperConfigurationBody"`

	// The current state of the cluster pipeline.
	LifecycleState OpensearchClusterPipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The current state of the pipeline.
	PipelineMode OpensearchClusterPipelinePipelineModeEnum `mandatory:"true" json:"pipelineMode"`

	// The OCID of the pipeline's VCN.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID of the pipeline's subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The date and time the cluster pipeline was created. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the cluster pipeline was updated. Format defined by RFC3339 (https://tools.ietf.org/html/rfc3339).
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

func (m OpensearchClusterPipelineSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OpensearchClusterPipelineSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingOpensearchClusterPipelineLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetOpensearchClusterPipelineLifecycleStateEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOpensearchClusterPipelinePipelineModeEnum(string(m.PipelineMode)); !ok && m.PipelineMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PipelineMode: %s. Supported values are: %s.", m.PipelineMode, strings.Join(GetOpensearchClusterPipelinePipelineModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
