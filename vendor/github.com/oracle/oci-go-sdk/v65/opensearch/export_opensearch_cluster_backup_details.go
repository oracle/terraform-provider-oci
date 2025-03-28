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

// ExportOpensearchClusterBackupDetails Information about the cluster backup to export.
type ExportOpensearchClusterBackupDetails struct {

	// The Object Storage namespace for the cluster backup export operation.
	ObjectStorageNamespace *string `mandatory:"true" json:"objectStorageNamespace"`

	// The name of the Object Storage bucket for the cluster backup export operation.
	ObjectStorageBucketName *string `mandatory:"true" json:"objectStorageBucketName"`

	// The name of the snapshot for the cluster backup export operation.
	SnapshotName *string `mandatory:"true" json:"snapshotName"`

	// The name of the repository containing the snapshots for the cluster backup export operation.
	RepositoryName *string `mandatory:"true" json:"repositoryName"`

	// The prefix within object storage bucket for the cluster backup export operation.
	Prefix *string `mandatory:"true" json:"prefix"`

	// The OCID of the compartment where the Object Storage resources for the cluster backup are located.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The prefix within the Object Storage bucket for the cluster backup export operation.
	ObjectStoragePrefix *string `mandatory:"false" json:"objectStoragePrefix"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m ExportOpensearchClusterBackupDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExportOpensearchClusterBackupDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
