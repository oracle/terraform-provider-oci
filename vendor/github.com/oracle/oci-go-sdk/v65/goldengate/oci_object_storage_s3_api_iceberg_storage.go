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

// OciObjectStorageS3ApiIcebergStorage Connection details of the OCI Object Storage (S3 Compatibility API) configuration used by the
// Iceberg connection.
type OciObjectStorageS3ApiIcebergStorage struct {

	// OCI Object Storage S3 Compatibility API endpoint URL.
	// Format: "https://<namespace>.compat.objectstorage.<region>.<domain>"
	// Example: "https://mynamespace.compat.objectstorage.us-ashburn-1.oraclecloud.com"
	Endpoint *string `mandatory:"true" json:"endpoint"`

	// Access Key ID from the OCI IAM user's Customer Secret Key pair used to authenticate to
	// OCI Object Storage via the S3 Compatibility API.
	// Note: Despite the "Id" suffix, this value is not an OCI OCID.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// Target OCI Object Storage bucket name where Iceberg stores
	// table metadata and data files.
	Bucket *string `mandatory:"true" json:"bucket"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key used for OCI Object Storage S3 Compatibility authentication.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`
}

func (m OciObjectStorageS3ApiIcebergStorage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciObjectStorageS3ApiIcebergStorage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciObjectStorageS3ApiIcebergStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciObjectStorageS3ApiIcebergStorage OciObjectStorageS3ApiIcebergStorage
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeOciObjectStorageS3ApiIcebergStorage
	}{
		"OCI_OBJECT_STORAGE_S3_API",
		(MarshalTypeOciObjectStorageS3ApiIcebergStorage)(m),
	}

	return json.Marshal(&s)
}
