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

// UpdateOciObjectStorageAccessKeyIcebergStorageDetails The information to update the OCI Object Storage configuration for an Oracle AI Data Catalog Iceberg catalog
// using access-key/secret-key credentials.
type UpdateOciObjectStorageAccessKeyIcebergStorageDetails struct {

	// The access key ID used to access the OCI Object Storage bucket.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the OCI Object Storage secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// The secret access key used to access the warehouse storage.
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content
	// of the configuration file containing additional properties for the REST catalog.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PropertiesSecretId *string `mandatory:"false" json:"propertiesSecretId"`

	// The base64 encoded content of the configuration file containing additional properties for the REST catalog.
	Properties *string `mandatory:"false" json:"properties"`
}

func (m UpdateOciObjectStorageAccessKeyIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOciObjectStorageAccessKeyIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOciObjectStorageAccessKeyIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOciObjectStorageAccessKeyIcebergStorageDetails UpdateOciObjectStorageAccessKeyIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateOciObjectStorageAccessKeyIcebergStorageDetails
	}{
		"OCI_OBJECT_STORAGE_ACCESS_KEY",
		(MarshalTypeUpdateOciObjectStorageAccessKeyIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
