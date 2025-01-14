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

// CreateOpensearchClusterPipelineDetails The configuration details for a new OpenSearch cluster pipeline.
type CreateOpensearchClusterPipelineDetails struct {

	// The name of the cluster pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The maximum pipeline capacity, in OCPUs.
	MaxOcpuCount *int `mandatory:"true" json:"maxOcpuCount"`

	// The minimum pipeline capacity, in OCPUs.
	MinOcpuCount *int `mandatory:"true" json:"minOcpuCount"`

	// The maximum amount of memory in GB, for the pipeline.
	MaxMemoryGB *int `mandatory:"true" json:"maxMemoryGB"`

	// The minimum amount of memory in GB, for the pipeline.
	MinMemoryGB *int `mandatory:"true" json:"minMemoryGB"`

	// The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	PipelineConfigurationBody *string `mandatory:"true" json:"pipelineConfigurationBody"`

	// The OCID of the compartment to create the pipeline in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the pipeline's VCN.
	VcnId *string `mandatory:"true" json:"vcnId"`

	// The OCID of the pipeline's subnet.
	SubnetId *string `mandatory:"true" json:"subnetId"`

	// The OCID for the compartment where the pipeline's VCN is located.
	VcnCompartmentId *string `mandatory:"true" json:"vcnCompartmentId"`

	// The OCID for the compartment where the pipwline's subnet is located.
	SubnetCompartmentId *string `mandatory:"true" json:"subnetCompartmentId"`

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

func (m CreateOpensearchClusterPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateOpensearchClusterPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
