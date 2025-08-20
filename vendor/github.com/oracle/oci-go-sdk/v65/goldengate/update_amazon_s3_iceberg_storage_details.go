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

// UpdateAmazonS3IcebergStorageDetails The information to update the Amazon S3 storage used in the Iceberg connection.
type UpdateAmazonS3IcebergStorageDetails struct {

	// The endpoint URL of the Amazon S3 storage service.
	// e.g.: 'https://s3.amazonaws.com'
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// Access key ID to access the Amazon S3 bucket.
	AccessKeyId *string `mandatory:"false" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
	SecretAccessKeySecretId *string `mandatory:"false" json:"secretAccessKeySecretId"`

	// The AMAZON region where the S3 bucket is hosted.
	// e.g.: 'us-east-2'
	Region *string `mandatory:"false" json:"region"`

	// S3 bucket where Iceberg stores metadata and data files.
	Bucket *string `mandatory:"false" json:"bucket"`

	// The scheme of the storage.
	SchemeType AmazonS3IcebergStorageSchemeTypeEnum `mandatory:"false" json:"schemeType,omitempty"`
}

func (m UpdateAmazonS3IcebergStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAmazonS3IcebergStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAmazonS3IcebergStorageSchemeTypeEnum(string(m.SchemeType)); !ok && m.SchemeType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SchemeType: %s. Supported values are: %s.", m.SchemeType, strings.Join(GetAmazonS3IcebergStorageSchemeTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateAmazonS3IcebergStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateAmazonS3IcebergStorageDetails UpdateAmazonS3IcebergStorageDetails
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeUpdateAmazonS3IcebergStorageDetails
	}{
		"AMAZON_S3",
		(MarshalTypeUpdateAmazonS3IcebergStorageDetails)(m),
	}

	return json.Marshal(&s)
}
