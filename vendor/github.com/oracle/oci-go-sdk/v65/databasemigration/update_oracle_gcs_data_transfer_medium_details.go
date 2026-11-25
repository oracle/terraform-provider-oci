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

// UpdateOracleGcsDataTransferMediumDetails Google Cloud Storage details used to transfer Data Pump dump files for Oracle migrations.
type UpdateOracleGcsDataTransferMediumDetails struct {

	// Google Cloud Storage bucket name.
	BucketName *string `mandatory:"false" json:"bucketName"`

	// GCS HMAC access-key ID.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of OCI Vault secret that contains the GCS HMAC secret access key.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`
}

func (m UpdateOracleGcsDataTransferMediumDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateOracleGcsDataTransferMediumDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateOracleGcsDataTransferMediumDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateOracleGcsDataTransferMediumDetails UpdateOracleGcsDataTransferMediumDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeUpdateOracleGcsDataTransferMediumDetails
	}{
		"GCS",
		(MarshalTypeUpdateOracleGcsDataTransferMediumDetails)(m),
	}

	return json.Marshal(&s)
}
