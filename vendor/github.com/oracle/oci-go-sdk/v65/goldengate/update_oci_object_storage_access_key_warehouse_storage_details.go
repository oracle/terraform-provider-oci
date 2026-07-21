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

// UpdateOciObjectStorageAccessKeyWarehouseStorageDetails The information to update the OCI Object Storage warehouse storage configuration for Oracle AI Data Catalog Storage
// Connection using access-key/secret-key credentials.
type UpdateOciObjectStorageAccessKeyWarehouseStorageDetails struct {

	// The name of the region. e.g.: us-ashburn-1
	// If the region is not provided, backend will default to the default region.
	Region *string `mandatory:"false" json:"region"`

	// The access key ID used to access the OCI Object Storage bucket.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the OCI Object Storage secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// The secret access key used to access the warehouse storage.
	SecretAccessKey *string `mandatory:"false" json:"secretAccessKey"`
}

func (m UpdateOciObjectStorageAccessKeyWarehouseStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOciObjectStorageAccessKeyWarehouseStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOciObjectStorageAccessKeyWarehouseStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOciObjectStorageAccessKeyWarehouseStorageDetails UpdateOciObjectStorageAccessKeyWarehouseStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateOciObjectStorageAccessKeyWarehouseStorageDetails
	}{
		"OCI_OBJECT_STORAGE_ACCESS_KEY",
		(MarshalTypeUpdateOciObjectStorageAccessKeyWarehouseStorageDetails)(m),
	}

	return json.Marshal(&s)
}
