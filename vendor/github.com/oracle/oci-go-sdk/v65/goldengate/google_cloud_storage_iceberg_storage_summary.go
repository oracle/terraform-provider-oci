// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// GoogleCloudStorageIcebergStorageSummary Summary of the Google Cloud Storage storage used in the Iceberg connection.
type GoogleCloudStorageIcebergStorageSummary struct {

	// Google Cloud Storage bucket where Iceberg stores metadata and data files.
	Bucket *string `mandatory:"true" json:"bucket"`

	// The Google Cloud Project where the bucket exists.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the content of the service account key file is stored,
	// which contains the credentials required to use Google Cloud Storage.
	ServiceAccountKeyFileSecretId *string `mandatory:"true" json:"serviceAccountKeyFileSecretId"`
}

func (m GoogleCloudStorageIcebergStorageSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GoogleCloudStorageIcebergStorageSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m GoogleCloudStorageIcebergStorageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeGoogleCloudStorageIcebergStorageSummary GoogleCloudStorageIcebergStorageSummary
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeGoogleCloudStorageIcebergStorageSummary
	}{
		"GOOGLE_CLOUD_STORAGE",
		(MarshalTypeGoogleCloudStorageIcebergStorageSummary)(m),
	}

	return json.Marshal(&s)
}
