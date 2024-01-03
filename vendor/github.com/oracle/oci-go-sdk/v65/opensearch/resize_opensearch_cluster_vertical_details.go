// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// ResizeOpensearchClusterVerticalDetails The OCPU and memory configuration to update on an existing OpenSearch cluster for vertical resizing (https://docs.cloud.oracle.com/iaas/Content/search-opensearch/Tasks/resizingacluster.htm#vertical).
type ResizeOpensearchClusterVerticalDetails struct {

	// The number of OCPUs to configure for the cluster's master nodes.
	MasterNodeHostOcpuCount *int `mandatory:"false" json:"masterNodeHostOcpuCount"`

	// The amount of memory in GB, to configure for the cluster's master nodes.
	MasterNodeHostMemoryGB *int `mandatory:"false" json:"masterNodeHostMemoryGB"`

	// The number of OCPUs to configure for the cluster's data nodes.
	DataNodeHostOcpuCount *int `mandatory:"false" json:"dataNodeHostOcpuCount"`

	// The amount of memory in GB, to configure for the cluster's data nodes.
	DataNodeHostMemoryGB *int `mandatory:"false" json:"dataNodeHostMemoryGB"`

	// The amount of storage in GB, to configure per node for the cluster's data nodes.
	DataNodeStorageGB *int `mandatory:"false" json:"dataNodeStorageGB"`

	// The number of OCPUs to configure for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostOcpuCount *int `mandatory:"false" json:"opendashboardNodeHostOcpuCount"`

	// The amount of memory in GB, to configure for the cluster's OpenSearch Dashboard nodes.
	OpendashboardNodeHostMemoryGB *int `mandatory:"false" json:"opendashboardNodeHostMemoryGB"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ResizeOpensearchClusterVerticalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResizeOpensearchClusterVerticalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
