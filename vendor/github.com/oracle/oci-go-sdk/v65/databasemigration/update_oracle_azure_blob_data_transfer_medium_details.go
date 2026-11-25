// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateOracleAzureBlobDataTransferMediumDetails Azure Blob Storage details used to transfer Data Pump dump files for Oracle migrations.
type UpdateOracleAzureBlobDataTransferMediumDetails struct {

	// Azure Storage account name.
	StorageAccountName *string `mandatory:"false" json:"storageAccountName"`

	// Azure Blob container name.
	ContainerName *string `mandatory:"false" json:"containerName"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OCI Vault secret that contains the Azure Storage Account Key.
	StorageAccountKeySecretId *string `mandatory:"false" json:"storageAccountKeySecretId"`
}

func (m UpdateOracleAzureBlobDataTransferMediumDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleAzureBlobDataTransferMediumDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleAzureBlobDataTransferMediumDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleAzureBlobDataTransferMediumDetails UpdateOracleAzureBlobDataTransferMediumDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOracleAzureBlobDataTransferMediumDetails
	}{
		"AZURE_BLOB",
		(MarshalTypeUpdateOracleAzureBlobDataTransferMediumDetails)(m),
	}

	return json.Marshal(&s)
}
