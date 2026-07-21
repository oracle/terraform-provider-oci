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

// OciObjectStorageAccessKeyWarehouseStorage Represents OCI Object Storage warehouse storage configuration for Oracle AI Data Catalog Storage Connection
// using access-key/secret-key credentials.
type OciObjectStorageAccessKeyWarehouseStorage struct {

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the default region.
	Region *string `mandatory:"true" json:"region"`

	// The access key ID used to access the OCI Object Storage bucket.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the OCI Object Storage secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`
}

func (m OciObjectStorageAccessKeyWarehouseStorage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OciObjectStorageAccessKeyWarehouseStorage) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m OciObjectStorageAccessKeyWarehouseStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeOciObjectStorageAccessKeyWarehouseStorage OciObjectStorageAccessKeyWarehouseStorage
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeOciObjectStorageAccessKeyWarehouseStorage
	}{
		"OCI_OBJECT_STORAGE_ACCESS_KEY",
		(MarshalTypeOciObjectStorageAccessKeyWarehouseStorage)(m),
	}

	return json.Marshal(&s)
}
