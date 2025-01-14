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

	// The OCID of the compartment where the cluster pipeline is located.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

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

	// The current state of the cluster backup.
	LifecycleState OpensearchClusterPipelineLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The name of the cluster pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

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

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
