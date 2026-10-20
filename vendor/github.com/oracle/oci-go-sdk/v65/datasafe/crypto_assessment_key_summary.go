// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CryptoAssessmentKeySummary Summary of one cryptographic key observed for a crypto assessment.
type CryptoAssessmentKeySummary struct {

	// OCID of the crypto assessment that discovered the key.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// OCID of the target database associated with the key.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Identifier of the cryptographic key.
	KeyId *string `mandatory:"true" json:"keyId"`

	// Crypto feature for which the key is observed.
	Feature CryptoFeatureEnum `mandatory:"true" json:"feature"`

	// Cryptographic algorithm used by the key.
	Algorithm *string `mandatory:"true" json:"algorithm"`

	// Key creation time in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Type of cryptographic key observed for the assessment.
	KeyType CryptoAssessmentKeySummaryKeyTypeEnum `mandatory:"false" json:"keyType,omitempty"`

	// Current status of the cryptographic key as observed on the target.
	Status *string `mandatory:"false" json:"status"`

	// Primary keystore type observed for the key.
	KeystoreType CryptoKeystoreTypeEnum `mandatory:"false" json:"keystoreType,omitempty"`

	// Secondary keystore type observed for the key, if configured.
	SecondaryKeystoreType CryptoKeystoreTypeEnum `mandatory:"false" json:"secondaryKeystoreType,omitempty"`

	// Key cache setting observed for the key.
	KeyCache CryptoKeyCacheStatusEnum `mandatory:"false" json:"keyCache,omitempty"`

	// Wallet location observed for the key.
	WalletLocation *string `mandatory:"false" json:"walletLocation"`

	// Most recent key rotation time in RFC3339 format.
	TimeLastRotation *common.SDKTime `mandatory:"false" json:"timeLastRotation"`

	// Age of the key in whole days, calculated from timeCreated using the current UTC date.
	Age *int64 `mandatory:"false" json:"age"`
}

func (m CryptoAssessmentKeySummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentKeySummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoFeatureEnum(string(m.Feature)); !ok && m.Feature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Feature: %s. Supported values are: %s.", m.Feature, strings.Join(GetCryptoFeatureEnumStringValues(), ",")))
	}

	if _, ok := GetMappingCryptoAssessmentKeySummaryKeyTypeEnum(string(m.KeyType)); !ok && m.KeyType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyType: %s. Supported values are: %s.", m.KeyType, strings.Join(GetCryptoAssessmentKeySummaryKeyTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoKeystoreTypeEnum(string(m.KeystoreType)); !ok && m.KeystoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeystoreType: %s. Supported values are: %s.", m.KeystoreType, strings.Join(GetCryptoKeystoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoKeystoreTypeEnum(string(m.SecondaryKeystoreType)); !ok && m.SecondaryKeystoreType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SecondaryKeystoreType: %s. Supported values are: %s.", m.SecondaryKeystoreType, strings.Join(GetCryptoKeystoreTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoKeyCacheStatusEnum(string(m.KeyCache)); !ok && m.KeyCache != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for KeyCache: %s. Supported values are: %s.", m.KeyCache, strings.Join(GetCryptoKeyCacheStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentKeySummaryKeyTypeEnum Enum with underlying type: string
type CryptoAssessmentKeySummaryKeyTypeEnum string

// Set of constants representing the allowable values for CryptoAssessmentKeySummaryKeyTypeEnum
const (
	CryptoAssessmentKeySummaryKeyTypeMasterKey     CryptoAssessmentKeySummaryKeyTypeEnum = "MASTER_KEY"
	CryptoAssessmentKeySummaryKeyTypeEncryptionKey CryptoAssessmentKeySummaryKeyTypeEnum = "ENCRYPTION_KEY"
)

var mappingCryptoAssessmentKeySummaryKeyTypeEnum = map[string]CryptoAssessmentKeySummaryKeyTypeEnum{
	"MASTER_KEY":     CryptoAssessmentKeySummaryKeyTypeMasterKey,
	"ENCRYPTION_KEY": CryptoAssessmentKeySummaryKeyTypeEncryptionKey,
}

var mappingCryptoAssessmentKeySummaryKeyTypeEnumLowerCase = map[string]CryptoAssessmentKeySummaryKeyTypeEnum{
	"master_key":     CryptoAssessmentKeySummaryKeyTypeMasterKey,
	"encryption_key": CryptoAssessmentKeySummaryKeyTypeEncryptionKey,
}

// GetCryptoAssessmentKeySummaryKeyTypeEnumValues Enumerates the set of values for CryptoAssessmentKeySummaryKeyTypeEnum
func GetCryptoAssessmentKeySummaryKeyTypeEnumValues() []CryptoAssessmentKeySummaryKeyTypeEnum {
	values := make([]CryptoAssessmentKeySummaryKeyTypeEnum, 0)
	for _, v := range mappingCryptoAssessmentKeySummaryKeyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentKeySummaryKeyTypeEnumStringValues Enumerates the set of values in String for CryptoAssessmentKeySummaryKeyTypeEnum
func GetCryptoAssessmentKeySummaryKeyTypeEnumStringValues() []string {
	return []string{
		"MASTER_KEY",
		"ENCRYPTION_KEY",
	}
}

// GetMappingCryptoAssessmentKeySummaryKeyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentKeySummaryKeyTypeEnum(val string) (CryptoAssessmentKeySummaryKeyTypeEnum, bool) {
	enum, ok := mappingCryptoAssessmentKeySummaryKeyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
