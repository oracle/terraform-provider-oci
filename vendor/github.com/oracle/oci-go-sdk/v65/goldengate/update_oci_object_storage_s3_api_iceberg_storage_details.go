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

// UpdateOciObjectStorageS3ApiIcebergStorageDetails The information to update the OCI Object Storage (S3 Compatibility API) configuration used by the
// Iceberg connection.
type UpdateOciObjectStorageS3ApiIcebergStorageDetails struct {

	// OCI Object Storage S3 Compatibility API endpoint URL.
	// Format: "https://<namespace>.compat.objectstorage.<region>.<domain>"
	// Example: "https://mynamespace.compat.objectstorage.us-ashburn-1.oraclecloud.com"
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// Access Key ID from the OCI IAM user's Customer Secret Key pair used to authenticate to
	// OCI Object Storage via the S3 Compatibility API.
	// Note: Despite the "Id" suffix, this value is not an OCI OCID.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key used for OCI Object Storage S3 Compatibility authentication.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// Secret Access Key from the OCI IAM user's Customer Secret Key pair used to authenticate to
	// OCI Object Storage via the S3 Compatibility API.
	// Deprecated: This field is deprecated and replaced by "secretAccessKeySecretId".
	// This change follows the GoldenGate "Plain Text Fields in Connections" deprecation:
	// https://docs.oracle.com/en-us/iaas/Content/servicechanges.htm#servicechanges_topic-GoldenGate
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`

	// Target OCI Object Storage bucket name where Iceberg stores
	// table metadata and data files.
	Bucket *string `mandatory:"false" json:"bucket"`
}

func (m UpdateOciObjectStorageS3ApiIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOciObjectStorageS3ApiIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOciObjectStorageS3ApiIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOciObjectStorageS3ApiIcebergStorageDetails UpdateOciObjectStorageS3ApiIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateOciObjectStorageS3ApiIcebergStorageDetails
	}{
		"OCI_OBJECT_STORAGE_S3_API",
		(MarshalTypeUpdateOciObjectStorageS3ApiIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
