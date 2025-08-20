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

// UpdateOpensearchClusterPipelineDetails The configuration to update on an existing OpenSearch cluster pipeline. You can only edit capacity limits and pipeline configurations. You can't edit its network settings.
type UpdateOpensearchClusterPipelineDetails struct {

	// The name of the pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The number of OCPUs configured for each pipeline node.
	OcpuCount *int `mandatory:"false" json:"ocpuCount"`

	// The amount of memory in GB, for each pipeline node.
	MemoryGB *int `mandatory:"false" json:"memoryGB"`

	// The number of nodes configured for the pipeline.
	NodeCount *int `mandatory:"false" json:"nodeCount"`

	// The pipeline node shape.
	NodeShape *string `mandatory:"false" json:"nodeShape"`

	// The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	PipelineConfigurationBody *string `mandatory:"false" json:"pipelineConfigurationBody"`

	// The data prepper config in YAML format. The command accepts the data prepper config as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	DataPrepperConfigurationBody *string `mandatory:"false" json:"dataPrepperConfigurationBody"`

	// The OCID of the pipeline's VCN.
	VcnId *string `mandatory:"false" json:"vcnId"`

	// The OCID of the pipeline's subnet.
	SubnetId *string `mandatory:"false" json:"subnetId"`

	// The OCID for the compartment where the pipeline's VCN is located.
	VcnCompartmentId *string `mandatory:"false" json:"vcnCompartmentId"`

	// The OCID for the compartment where the pipeline's subnet is located.
	SubnetCompartmentId *string `mandatory:"false" json:"subnetCompartmentId"`

	// The OCID of the NSG where the pipeline private endpoint vnic will be attached.
	NsgId *string `mandatory:"false" json:"nsgId"`

	// The customer IP and the corresponding fully qualified domain name that the pipeline will connect to.
	ReverseConnectionEndpoints []OpensearchPipelineReverseConnectionEndpoint `mandatory:"false" json:"reverseConnectionEndpoints"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateOpensearchClusterPipelineDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOpensearchClusterPipelineDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
