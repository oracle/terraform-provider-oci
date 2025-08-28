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

// AmazonS3IcebergStorage Represents an Amazon S3 storage used in the Iceberg connection.
type AmazonS3IcebergStorage struct {

	// Access key ID to access the Amazon S3 bucket.
	AccessKeyId *string `mandatory:"true" json:"accessKeyId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the Secret Access Key is stored.
	SecretAccessKeySecretId *string `mandatory:"true" json:"secretAccessKeySecretId"`

	// The AMAZON region where the S3 bucket is hosted.
	// e.g.: 'us-east-2'
	Region *string `mandatory:"true" json:"region"`

	// S3 bucket where Iceberg stores metadata and data files.
	Bucket *string `mandatory:"true" json:"bucket"`

	// The endpoint URL of the Amazon S3 storage service.
	// e.g.: 'https://s3.amazonaws.com'
	Endpoint *string `mandatory:"false" json:"endpoint"`

	// The scheme of the storage.
	SchemeType AmazonS3IcebergStorageSchemeTypeEnum `mandatory:"true" json:"schemeType"`
}

func (m AmazonS3IcebergStorage) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AmazonS3IcebergStorage) ValidateEnumValue() (bool, error) {
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
func (m AmazonS3IcebergStorage) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAmazonS3IcebergStorage AmazonS3IcebergStorage
	s := struct {
		DiscriminatorParam string `json:"storageType"`
		MarshalTypeAmazonS3IcebergStorage
	}{
		"AMAZON_S3",
		(MarshalTypeAmazonS3IcebergStorage)(m),
	}

	return json.Marshal(&s)
}

// AmazonS3IcebergStorageSchemeTypeEnum Enum with underlying type: string
type AmazonS3IcebergStorageSchemeTypeEnum string

// Set of constants representing the allowable values for AmazonS3IcebergStorageSchemeTypeEnum
const (
	AmazonS3IcebergStorageSchemeTypeS3  AmazonS3IcebergStorageSchemeTypeEnum = "S3"
	AmazonS3IcebergStorageSchemeTypeS3a AmazonS3IcebergStorageSchemeTypeEnum = "S3A"
)

var mappingAmazonS3IcebergStorageSchemeTypeEnum = map[string]AmazonS3IcebergStorageSchemeTypeEnum{
	"S3":  AmazonS3IcebergStorageSchemeTypeS3,
	"S3A": AmazonS3IcebergStorageSchemeTypeS3a,
}

var mappingAmazonS3IcebergStorageSchemeTypeEnumLowerCase = map[string]AmazonS3IcebergStorageSchemeTypeEnum{
	"s3":  AmazonS3IcebergStorageSchemeTypeS3,
	"s3a": AmazonS3IcebergStorageSchemeTypeS3a,
}

// GetAmazonS3IcebergStorageSchemeTypeEnumValues Enumerates the set of values for AmazonS3IcebergStorageSchemeTypeEnum
func GetAmazonS3IcebergStorageSchemeTypeEnumValues() []AmazonS3IcebergStorageSchemeTypeEnum {
	values := make([]AmazonS3IcebergStorageSchemeTypeEnum, 0)
	for _, v := range mappingAmazonS3IcebergStorageSchemeTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetAmazonS3IcebergStorageSchemeTypeEnumStringValues Enumerates the set of values in String for AmazonS3IcebergStorageSchemeTypeEnum
func GetAmazonS3IcebergStorageSchemeTypeEnumStringValues() []string {
	return []string{
		"S3",
		"S3A",
	}
}

// GetMappingAmazonS3IcebergStorageSchemeTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAmazonS3IcebergStorageSchemeTypeEnum(val string) (AmazonS3IcebergStorageSchemeTypeEnum, bool) {
	enum, ok := mappingAmazonS3IcebergStorageSchemeTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
