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

// CreateAzureDataLakeStorageIcebergStorageDetails The information about a new Azure Data Lake Storage storage used in the Iceberg connection.
type CreateAzureDataLakeStorageIcebergStorageDetails struct {

	// Sets the Azure storage account name.
	AccountName *string `mandatory:"true" json:"accountName"`

	// The Azure Blob Storage container where Iceberg tables are stored.
	Container *string `mandatory:"true" json:"container"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the account key is stored.
	AccountKeySecretId *string `mandatory:"true" json:"accountKeySecretId"`

	// The Azure Blob Storage endpoint where Iceberg data is stored.
	// e.g.: 'https://my-azure-storage-account.blob.core.windows.net'
	Endpoint *string `mandatory:"false" json:"endpoint"`
}

func (m CreateAzureDataLakeStorageIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAzureDataLakeStorageIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAzureDataLakeStorageIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAzureDataLakeStorageIcebergStorageDetails CreateAzureDataLakeStorageIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeCreateAzureDataLakeStorageIcebergStorageDetails
	}{
		"AZURE_DATA_LAKE_STORAGE",
		(MarshalTypeCreateAzureDataLakeStorageIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
