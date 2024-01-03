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

// ResizeOpensearchClusterHorizontalDetails The node count configuration to update on an existing OpenSearch cluster for horizontal resizing (https://docs.cloud.oracle.com/iaas/Content/search-opensearch/Tasks/resizingacluster.htm#horizontalresize).
type ResizeOpensearchClusterHorizontalDetails struct {

	// The number of master nodes to configure for the cluster.
	MasterNodeCount *int `mandatory:"false" json:"masterNodeCount"`

	// The number of data nodes to configure for the cluster.
	DataNodeCount *int `mandatory:"false" json:"dataNodeCount"`

	// The number of OpenSearch Dashboard nodes to configure for the cluster.
	OpendashboardNodeCount *int `mandatory:"false" json:"opendashboardNodeCount"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ResizeOpensearchClusterHorizontalDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResizeOpensearchClusterHorizontalDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
