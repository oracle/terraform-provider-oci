// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Vault Service Secret Retrieval API
//
// API for retrieving secrets from vaults.
//

package secrets

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecretBundle The contents of the secret, properties of the secret (and secret version), and user-provided contextual metadata for the secret.
type SecretBundle struct {

	// The OCID of the secret.
	SecretId *string `mandatory:"true" json:"secretId"`

	// The version number of the secret.
	VersionNumber *int64 `mandatory:"true" json:"versionNumber"`

	// The time when the secret bundle was created.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The name of the secret version. Labels are unique across the different versions of a particular secret.
	VersionName *string `mandatory:"false" json:"versionName"`

	SecretBundleContent SecretBundleContentDetails `mandatory:"false" json:"secretBundleContent"`

	// An optional property indicating when to delete the secret version, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfDeletion *common.SDKTime `mandatory:"false" json:"timeOfDeletion"`

	// An optional property indicating when the secret version will expire, expressed in RFC 3339 (https://tools.ietf.org/html/rfc3339) timestamp format.
	// Example: `2019-04-03T21:10:29.600Z`
	TimeOfExpiry *common.SDKTime `mandatory:"false" json:"timeOfExpiry"`

	// A list of possible rotation states for the secret version.
	Stages []SecretBundleStagesEnum `mandatory:"false" json:"stages,omitempty"`

	// Customer-provided contextual metadata for the secret.
	Metadata map[string]interface{} `mandatory:"false" json:"metadata"`
}

func (m SecretBundle) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecretBundle) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.Stages {
		if _, ok := GetMappingSecretBundleStagesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Stages: %s. Supported values are: %s.", val, strings.Join(GetSecretBundleStagesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *SecretBundle) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		TimeCreated         *common.SDKTime            `json:"timeCreated"`
		VersionName         *string                    `json:"versionName"`
		SecretBundleContent secretbundlecontentdetails `json:"secretBundleContent"`
		TimeOfDeletion      *common.SDKTime            `json:"timeOfDeletion"`
		TimeOfExpiry        *common.SDKTime            `json:"timeOfExpiry"`
		Stages              []SecretBundleStagesEnum   `json:"stages"`
		Metadata            map[string]interface{}     `json:"metadata"`
		SecretId            *string                    `json:"secretId"`
		VersionNumber       *int64                     `json:"versionNumber"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.TimeCreated = model.TimeCreated

	m.VersionName = model.VersionName

	nn, e = model.SecretBundleContent.UnmarshalPolymorphicJSON(model.SecretBundleContent.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.SecretBundleContent = nn.(SecretBundleContentDetails)
	} else {
		m.SecretBundleContent = nil
	}

	m.TimeOfDeletion = model.TimeOfDeletion

	m.TimeOfExpiry = model.TimeOfExpiry

	m.Stages = make([]SecretBundleStagesEnum, len(model.Stages))
	for i, n := range model.Stages {
		m.Stages[i] = n
	}

	m.Metadata = model.Metadata

	m.SecretId = model.SecretId

	m.VersionNumber = model.VersionNumber

	return
}

// SecretBundleStagesEnum Enum with underlying type: string
type SecretBundleStagesEnum string

// Set of constants representing the allowable values for SecretBundleStagesEnum
const (
	SecretBundleStagesCurrent    SecretBundleStagesEnum = "CURRENT"
	SecretBundleStagesPending    SecretBundleStagesEnum = "PENDING"
	SecretBundleStagesLatest     SecretBundleStagesEnum = "LATEST"
	SecretBundleStagesPrevious   SecretBundleStagesEnum = "PREVIOUS"
	SecretBundleStagesDeprecated SecretBundleStagesEnum = "DEPRECATED"
)

var mappingSecretBundleStagesEnum = map[string]SecretBundleStagesEnum{
	"CURRENT":    SecretBundleStagesCurrent,
	"PENDING":    SecretBundleStagesPending,
	"LATEST":     SecretBundleStagesLatest,
	"PREVIOUS":   SecretBundleStagesPrevious,
	"DEPRECATED": SecretBundleStagesDeprecated,
}

var mappingSecretBundleStagesEnumLowerCase = map[string]SecretBundleStagesEnum{
	"current":    SecretBundleStagesCurrent,
	"pending":    SecretBundleStagesPending,
	"latest":     SecretBundleStagesLatest,
	"previous":   SecretBundleStagesPrevious,
	"deprecated": SecretBundleStagesDeprecated,
}

// GetSecretBundleStagesEnumValues Enumerates the set of values for SecretBundleStagesEnum
func GetSecretBundleStagesEnumValues() []SecretBundleStagesEnum {
	values := make([]SecretBundleStagesEnum, 0)
	for _, v := range mappingSecretBundleStagesEnum {
		values = append(values, v)
	}
	return values
}

// GetSecretBundleStagesEnumStringValues Enumerates the set of values in String for SecretBundleStagesEnum
func GetSecretBundleStagesEnumStringValues() []string {
	return []string{
		"CURRENT",
		"PENDING",
		"LATEST",
		"PREVIOUS",
		"DEPRECATED",
	}
}

// GetMappingSecretBundleStagesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecretBundleStagesEnum(val string) (SecretBundleStagesEnum, bool) {
	enum, ok := mappingSecretBundleStagesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
