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

// UpdateAzureDataLakeStorageIcebergStorageDetails The information to update the Azure Data Lake Storage storage used in the Iceberg connection.
type UpdateAzureDataLakeStorageIcebergStorageDetails struct {

	// Sets the Azure storage account name.
	AccountName *string `mandatory:"false" json:"accountName"`

	// The Azure Blob Storage container where Iceberg tables are stored.
	Container *string `mandatory:"false" json:"container"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored.
	AccountKeySecretId *string `mandatory:"false" json:"accountKeySecretId"`

	// The Azure Blob Storage endpoint where Iceberg data is stored.
	// e.g.: 'https://my-azure-storage-account.blob.core.windows.net'
	Endpoint *string `mandatory:"false" json:"endpoint"`
}

func (m UpdateAzureDataLakeStorageIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAzureDataLakeStorageIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAzureDataLakeStorageIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAzureDataLakeStorageIcebergStorageDetails UpdateAzureDataLakeStorageIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateAzureDataLakeStorageIcebergStorageDetails
	}{
		"AZURE_DATA_LAKE_STORAGE",
		(MarshalTypeUpdateAzureDataLakeStorageIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
