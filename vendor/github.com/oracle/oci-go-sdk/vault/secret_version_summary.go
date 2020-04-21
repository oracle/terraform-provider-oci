// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Secrets Management API
//
// API for managing secrets.
//

package vault

import (
	"github.com/oracle/oci-go-sdk/common"
)

// SecretVersionSummary The secret version summary object, which doesn't include the contents of the secret.
type SecretVersionSummary struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" json:"secretId"`

	// A optional property indicating when the secret version was created, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// The content type of the secret version's secret contents.
	ContentType SecretVersionSummaryContentTypeEnum `mandatory:"false" json:"contentType,omitempty"`

	// The name of the secret version. A name is unique across versions of a secret.
	Name *string `mandatory:"false" json:"name"`

	// A list of possible rotation states for the secret version. A secret version marked `CURRENT` is currently in use. A secret version
	// marked `PENDING` is staged and available for use, but has not been applied on the target system and, therefore, has not been rotated
	// into current, active use. The secret most recently uploaded to a vault is always marked `LATEST`. (The first version of a secret is
	// always marked as both `CURRENT` and `LATEST`.) A secret version marked `PREVIOUS` is the secret version that was most recently marked
	// `CURRENT`, before the last secret version rotation. A secret version marked `DEPRECATED` is neither current, pending, nor the previous
	// one in use. Only secret versions marked `DEPRECATED` can be scheduled for deletion.
	Stages []SecretVersionSummaryStagesEnum `mandatory:"false" json:"stages,omitempty"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`
}

func (m SecretVersionSummary) String() string {
	return common.PointerString(m)
}

// SecretVersionSummaryContentTypeEnum Enum with underlying type: string
type SecretVersionSummaryContentTypeEnum string

// Set of constants representing the allowable values for SecretVersionSummaryContentTypeEnum
const (
	SecretVersionSummaryContentTypeBase64 SecretVersionSummaryContentTypeEnum = "BASE64"
)

var mappingSecretVersionSummaryContentType = map[string]SecretVersionSummaryContentTypeEnum{
	"BASE64": SecretVersionSummaryContentTypeBase64,
}

// GetSecretVersionSummaryContentTypeEnumValues Enumerates the set of values for SecretVersionSummaryContentTypeEnum
func GetSecretVersionSummaryContentTypeEnumValues() []SecretVersionSummaryContentTypeEnum {
	values := make([]SecretVersionSummaryContentTypeEnum, 0)
	for _, v := range mappingSecretVersionSummaryContentType {
		values = append(values, v)
	}
	return values
}

// SecretVersionSummaryStagesEnum Enum with underlying type: string
type SecretVersionSummaryStagesEnum string

// Set of constants representing the allowable values for SecretVersionSummaryStagesEnum
const (
	SecretVersionSummaryStagesCurrent    SecretVersionSummaryStagesEnum = "CURRENT"
	SecretVersionSummaryStagesPending    SecretVersionSummaryStagesEnum = "PENDING"
	SecretVersionSummaryStagesLatest     SecretVersionSummaryStagesEnum = "LATEST"
	SecretVersionSummaryStagesPrevious   SecretVersionSummaryStagesEnum = "PREVIOUS"
	SecretVersionSummaryStagesDeprecated SecretVersionSummaryStagesEnum = "DEPRECATED"
)

var mappingSecretVersionSummaryStages = map[string]SecretVersionSummaryStagesEnum{
	"CURRENT":    SecretVersionSummaryStagesCurrent,
	"PENDING":    SecretVersionSummaryStagesPending,
	"LATEST":     SecretVersionSummaryStagesLatest,
	"PREVIOUS":   SecretVersionSummaryStagesPrevious,
	"DEPRECATED": SecretVersionSummaryStagesDeprecated,
}

// GetSecretVersionSummaryStagesEnumValues Enumerates the set of values for SecretVersionSummaryStagesEnum
func GetSecretVersionSummaryStagesEnumValues() []SecretVersionSummaryStagesEnum {
	values := make([]SecretVersionSummaryStagesEnum, 0)
	for _, v := range mappingSecretVersionSummaryStages {
		values = append(values, v)
	}
	return values
}
