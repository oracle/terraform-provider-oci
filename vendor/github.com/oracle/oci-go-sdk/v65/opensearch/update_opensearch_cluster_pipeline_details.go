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

// UpdateOpensearchClusterPipelineDetails The configuration to update on an existing OpenSearch cluster pipeline. You can only edit capaccity limits and pipeline configuration. You can't edit its name or network settings.
type UpdateOpensearchClusterPipelineDetails struct {

	// The name of the pipeline. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The maximum pipeline capacity, in OCPUs.
	MaxOcpuCount *int `mandatory:"false" json:"maxOcpuCount"`

	// The minimum pipeline capacity, in OCPUs.
	MinOcpuCount *int `mandatory:"false" json:"minOcpuCount"`

	// The maximum amount of memory in GB, for the pipeline.
	MaxMemoryGB *int `mandatory:"false" json:"maxMemoryGB"`

	// The minimum amount of memory in GB, for the pipeline.
	MinMemoryGB *int `mandatory:"false" json:"minMemoryGB"`

	// The pipeline configuration in YAML format. The command accepts the pipeline configuration as a string or within a .yaml file. If you provide the configuration as a string, each new line must be escaped with \.
	PipelineConfigurationBody *string `mandatory:"false" json:"pipelineConfigurationBody"`
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
