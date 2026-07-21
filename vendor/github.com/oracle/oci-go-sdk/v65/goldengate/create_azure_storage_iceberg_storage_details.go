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

// CreateAzureStorageIcebergStorageDetails The information about a new Azure Storage storage used in the Iceberg connection for an Oracle AI Data Catalog Iceberg catalogs.
type CreateAzureStorageIcebergStorageDetails struct {

	// The Azure Storage account name.
	AccountName *string `mandatory:"true" json:"accountName"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the Azure Storage account key.
	AccountKeySecretId *string `mandatory:"false" json:"accountKeySecretId"`

	// The Azure Storage account key.
	AccountKey *string `mandatory:"false" json:"accountKey"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret that stores the content
	// of the configuration file containing additional properties for the REST catalog.
	// See documentation: https://docs.oracle.com/en-us/iaas/Content/Identity/Tasks/managingcredentials.htm
	PropertiesSecretId *string `mandatory:"false" json:"propertiesSecretId"`

	// The base64 encoded content of the configuration file containing additional properties for the REST catalog.
	Properties *string `mandatory:"false" json:"properties"`
}

func (m CreateAzureStorageIcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateAzureStorageIcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CreateAzureStorageIcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCreateAzureStorageIcebergStorageDetails CreateAzureStorageIcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeCreateAzureStorageIcebergStorageDetails
	}{
		"AZURE_STORAGE",
		(MarshalTypeCreateAzureStorageIcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
