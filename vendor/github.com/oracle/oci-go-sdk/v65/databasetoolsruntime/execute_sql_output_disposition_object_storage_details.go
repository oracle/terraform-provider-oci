// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlOutputDispositionObjectStorageDetails Describes how the result of a statement is stored in Object Storage
type ExecuteSqlOutputDispositionObjectStorageDetails struct {

	// The name of the object storage namespace
	Namespace *string `mandatory:"true" json:"namespace"`

	// The name of the object storage bucket
	BucketName *string `mandatory:"true" json:"bucketName"`

	// The name of the object template (can contain statementId placeholder, for example; query_{statementId}.csv )
	ObjectName *string `mandatory:"true" json:"objectName"`

	// The content type
	ContentType *string `mandatory:"false" json:"contentType"`

	// The content disposition
	ContentDisposition *string `mandatory:"false" json:"contentDisposition"`

	// The content encoding
	ContentEncoding *string `mandatory:"false" json:"contentEncoding"`

	// The Cache-Control header
	CacheControl *string `mandatory:"false" json:"cacheControl"`

	// The content language
	ContentLanguage *string `mandatory:"false" json:"contentLanguage"`

	// The object metadata
	OpcMeta map[string]string `mandatory:"false" json:"opcMeta"`

	// The Encryption Algorithm
	OpcSseCustomerAlgorithm *string `mandatory:"false" json:"opcSseCustomerAlgorithm"`

	// The base64-encoded 256-bit encryption key to use
	OpcSseCustomerKey *string `mandatory:"false" json:"opcSseCustomerKey"`

	// The base64-encoded SHA256 hash of the encryption key
	OpcSseCustomerKeySha256 *string `mandatory:"false" json:"opcSseCustomerKeySha256"`

	// The OCID of a master encryption key
	OpcSseKmsKeyId *string `mandatory:"false" json:"opcSseKmsKeyId"`

	// The storage tier that the object should be stored in
	StorageTier ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum `mandatory:"false" json:"storageTier,omitempty"`
}

func (m ExecuteSqlOutputDispositionObjectStorageDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlOutputDispositionObjectStorageDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum(string(m.StorageTier)); !ok && m.StorageTier != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StorageTier: %s. Supported values are: %s.", m.StorageTier, strings.Join(GetExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlOutputDispositionObjectStorageDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlOutputDispositionObjectStorageDetails ExecuteSqlOutputDispositionObjectStorageDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlOutputDispositionObjectStorageDetails
	}{
		"OBJECT_STORAGE",
		(MarshalTypeExecuteSqlOutputDispositionObjectStorageDetails)(m),
	}

	return json.Marshal(&s)
}

// ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum Enum with underlying type: string
type ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum string

// Set of constants representing the allowable values for ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum
const (
	ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierStandard         ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum = "STANDARD"
	ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierInfrequentAccess ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum = "INFREQUENT_ACCESS"
	ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierArchive          ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum = "ARCHIVE"
)

var mappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum = map[string]ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum{
	"STANDARD":          ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierStandard,
	"INFREQUENT_ACCESS": ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierInfrequentAccess,
	"ARCHIVE":           ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierArchive,
}

var mappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumLowerCase = map[string]ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum{
	"standard":          ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierStandard,
	"infrequent_access": ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierInfrequentAccess,
	"archive":           ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierArchive,
}

// GetExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumValues Enumerates the set of values for ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum
func GetExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumValues() []ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum {
	values := make([]ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum, 0)
	for _, v := range mappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum {
		values = append(values, v)
	}
	return values
}

// GetExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumStringValues Enumerates the set of values in String for ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum
func GetExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumStringValues() []string {
	return []string{
		"STANDARD",
		"INFREQUENT_ACCESS",
		"ARCHIVE",
	}
}

// GetMappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum(val string) (ExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnum, bool) {
	enum, ok := mappingExecuteSqlOutputDispositionObjectStorageDetailsStorageTierEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
