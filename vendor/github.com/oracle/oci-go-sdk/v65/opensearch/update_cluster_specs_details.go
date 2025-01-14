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

// UpdateClusterSpecsDetails update cluster specs in Kiev.
type UpdateClusterSpecsDetails struct {

	// OCID of the Opensearch Cluster.
	ClusterId *string `mandatory:"true" json:"clusterId"`

	// Updated value of OCPU's configured for master nodes of the cluster.
	MasterNodeHostOcpuCount *int `mandatory:"false" json:"masterNodeHostOcpuCount"`

	// Updated value of memory for master nodes in the cluster (in GB).
	MasterNodeHostMemoryGB *int `mandatory:"false" json:"masterNodeHostMemoryGB"`

	// Updated value of OCPU's configured for data nodes of the cluster.
	DataNodeHostOcpuCount *int `mandatory:"false" json:"dataNodeHostOcpuCount"`

	// Updated value of memory for data nodes in the cluster (in GB).
	DataNodeHostMemoryGB *int `mandatory:"false" json:"dataNodeHostMemoryGB"`

	// Updated version of the software the cluster is currently running.
	SoftwareVersion *string `mandatory:"false" json:"softwareVersion"`

	// Updated version of the dashboard software the cluster is currently running.
	DashboardSoftwareVersion *string `mandatory:"false" json:"dashboardSoftwareVersion"`

	// Updated Private endpoint of cluster.
	ClusterPrivateEndpoint *string `mandatory:"false" json:"clusterPrivateEndpoint"`
}

func (m UpdateClusterSpecsDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateClusterSpecsDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
