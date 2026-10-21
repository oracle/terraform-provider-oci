// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// GoldenGate API
//
// Use the Oracle Cloud Infrastructure GoldenGate APIs to perform data replication operations.
//

package goldengate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AzureDataLakeStorageDatabricksEngineArtifactLocation Represents an Azure Data Lake Storage Gen2 location used to share Databricks application artifacts.
type AzureDataLakeStorageDatabricksEngineArtifactLocation struct {

	// Azure Data Lake Storage Gen2 storage-account name.
	StorageAccount *string `mandatory:"true" json:"storageAccount"`

	// Azure Data Lake Storage Gen2 filesystem containing the application artifacts.
	// Example: "ggs-artifacts"
	Filesystem *string `mandatory:"true" json:"filesystem"`

	// Artifact path.
	Path *string `mandatory:"true" json:"path"`

	// Controls whether a staged application artifact replaces an existing file or object with
	// the same name at the configured destination. The artifact is typically a GGFAB pipeline
	// package or one of its dependency files.
	// When set to `true`, the existing file or object is replaced with the newly staged version.
	// When set to `false`, staging fails or skips the upload rather than replacing it. This
	// setting does not replace the entire target directory or prefix. It allows an updated
	// pipeline or dependency to be deployed when artifact names remain the same across versions.
	ShouldOverwriteExisting *bool `mandatory:"false" json:"shouldOverwriteExisting"`

	// Endpoint suffix. Defaults to dfs.core.windows.net.
	EndpointSuffix *string `mandatory:"false" json:"endpointSuffix"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the storage account access key.
	// Note: When provided, 'accountKey' field must not be provided.
	AccountKeySecretId *string `mandatory:"false" json:"accountKeySecretId"`
}

// GetShouldOverwriteExisting returns ShouldOverwriteExisting
func (m AzureDataLakeStorageDatabricksEngineArtifactLocation) GetShouldOverwriteExisting() *bool {
	return m.ShouldOverwriteExisting
}

func (m AzureDataLakeStorageDatabricksEngineArtifactLocation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AzureDataLakeStorageDatabricksEngineArtifactLocation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AzureDataLakeStorageDatabricksEngineArtifactLocation) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAzureDataLakeStorageDatabricksEngineArtifactLocation AzureDataLakeStorageDatabricksEngineArtifactLocation
	s := struct {
		DiscriminatorParam string `json:"locationType"`
		MarshalTypeAzureDataLakeStorageDatabricksEngineArtifactLocation
	}{
		"AZURE_DATE_LAKE_STORAGE",
		(MarshalTypeAzureDataLakeStorageDatabricksEngineArtifactLocation)(m),
	}

	return json.Marshal(&s)
}
