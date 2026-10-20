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

// CryptoAssessmentWalletSummary Wallet details for one crypto feature.
type CryptoAssessmentWalletSummary struct {

	// OCID of the crypto assessment that discovered the wallet.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// OCID of the target database associated with the wallet.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Crypto feature for which wallet details are reported.
	Feature CryptoAssessmentWalletSummaryFeatureEnum `mandatory:"true" json:"feature"`

	// Wallet path for the feature.
	WalletLocation *string `mandatory:"true" json:"walletLocation"`

	// Whether wallet auto-login is enabled.
	AutoLogin CryptoAssessmentWalletSummaryAutoLoginEnum `mandatory:"true" json:"autoLogin"`

	// Wallet encryption algorithm.
	WalletEncryptionAlgorithm *string `mandatory:"true" json:"walletEncryptionAlgorithm"`

	// The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Wallet creation time in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m CryptoAssessmentWalletSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentWalletSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentWalletSummaryFeatureEnum(string(m.Feature)); !ok && m.Feature != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Feature: %s. Supported values are: %s.", m.Feature, strings.Join(GetCryptoAssessmentWalletSummaryFeatureEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentWalletSummaryAutoLoginEnum(string(m.AutoLogin)); !ok && m.AutoLogin != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoLogin: %s. Supported values are: %s.", m.AutoLogin, strings.Join(GetCryptoAssessmentWalletSummaryAutoLoginEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentWalletSummaryFeatureEnum Enum with underlying type: string
type CryptoAssessmentWalletSummaryFeatureEnum string

// Set of constants representing the allowable values for CryptoAssessmentWalletSummaryFeatureEnum
const (
	CryptoAssessmentWalletSummaryFeatureTde   CryptoAssessmentWalletSummaryFeatureEnum = "TDE"
	CryptoAssessmentWalletSummaryFeatureTls   CryptoAssessmentWalletSummaryFeatureEnum = "TLS"
	CryptoAssessmentWalletSummaryFeatureNne   CryptoAssessmentWalletSummaryFeatureEnum = "NNE"
	CryptoAssessmentWalletSummaryFeatureZdlra CryptoAssessmentWalletSummaryFeatureEnum = "ZDLRA"
)

var mappingCryptoAssessmentWalletSummaryFeatureEnum = map[string]CryptoAssessmentWalletSummaryFeatureEnum{
	"TDE":   CryptoAssessmentWalletSummaryFeatureTde,
	"TLS":   CryptoAssessmentWalletSummaryFeatureTls,
	"NNE":   CryptoAssessmentWalletSummaryFeatureNne,
	"ZDLRA": CryptoAssessmentWalletSummaryFeatureZdlra,
}

var mappingCryptoAssessmentWalletSummaryFeatureEnumLowerCase = map[string]CryptoAssessmentWalletSummaryFeatureEnum{
	"tde":   CryptoAssessmentWalletSummaryFeatureTde,
	"tls":   CryptoAssessmentWalletSummaryFeatureTls,
	"nne":   CryptoAssessmentWalletSummaryFeatureNne,
	"zdlra": CryptoAssessmentWalletSummaryFeatureZdlra,
}

// GetCryptoAssessmentWalletSummaryFeatureEnumValues Enumerates the set of values for CryptoAssessmentWalletSummaryFeatureEnum
func GetCryptoAssessmentWalletSummaryFeatureEnumValues() []CryptoAssessmentWalletSummaryFeatureEnum {
	values := make([]CryptoAssessmentWalletSummaryFeatureEnum, 0)
	for _, v := range mappingCryptoAssessmentWalletSummaryFeatureEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentWalletSummaryFeatureEnumStringValues Enumerates the set of values in String for CryptoAssessmentWalletSummaryFeatureEnum
func GetCryptoAssessmentWalletSummaryFeatureEnumStringValues() []string {
	return []string{
		"TDE",
		"TLS",
		"NNE",
		"ZDLRA",
	}
}

// GetMappingCryptoAssessmentWalletSummaryFeatureEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentWalletSummaryFeatureEnum(val string) (CryptoAssessmentWalletSummaryFeatureEnum, bool) {
	enum, ok := mappingCryptoAssessmentWalletSummaryFeatureEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentWalletSummaryAutoLoginEnum Enum with underlying type: string
type CryptoAssessmentWalletSummaryAutoLoginEnum string

// Set of constants representing the allowable values for CryptoAssessmentWalletSummaryAutoLoginEnum
const (
	CryptoAssessmentWalletSummaryAutoLoginEnabled       CryptoAssessmentWalletSummaryAutoLoginEnum = "ENABLED"
	CryptoAssessmentWalletSummaryAutoLoginDisabled      CryptoAssessmentWalletSummaryAutoLoginEnum = "DISABLED"
	CryptoAssessmentWalletSummaryAutoLoginNotApplicable CryptoAssessmentWalletSummaryAutoLoginEnum = "NOT_APPLICABLE"
	CryptoAssessmentWalletSummaryAutoLoginNotSupported  CryptoAssessmentWalletSummaryAutoLoginEnum = "NOT_SUPPORTED"
)

var mappingCryptoAssessmentWalletSummaryAutoLoginEnum = map[string]CryptoAssessmentWalletSummaryAutoLoginEnum{
	"ENABLED":        CryptoAssessmentWalletSummaryAutoLoginEnabled,
	"DISABLED":       CryptoAssessmentWalletSummaryAutoLoginDisabled,
	"NOT_APPLICABLE": CryptoAssessmentWalletSummaryAutoLoginNotApplicable,
	"NOT_SUPPORTED":  CryptoAssessmentWalletSummaryAutoLoginNotSupported,
}

var mappingCryptoAssessmentWalletSummaryAutoLoginEnumLowerCase = map[string]CryptoAssessmentWalletSummaryAutoLoginEnum{
	"enabled":        CryptoAssessmentWalletSummaryAutoLoginEnabled,
	"disabled":       CryptoAssessmentWalletSummaryAutoLoginDisabled,
	"not_applicable": CryptoAssessmentWalletSummaryAutoLoginNotApplicable,
	"not_supported":  CryptoAssessmentWalletSummaryAutoLoginNotSupported,
}

// GetCryptoAssessmentWalletSummaryAutoLoginEnumValues Enumerates the set of values for CryptoAssessmentWalletSummaryAutoLoginEnum
func GetCryptoAssessmentWalletSummaryAutoLoginEnumValues() []CryptoAssessmentWalletSummaryAutoLoginEnum {
	values := make([]CryptoAssessmentWalletSummaryAutoLoginEnum, 0)
	for _, v := range mappingCryptoAssessmentWalletSummaryAutoLoginEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentWalletSummaryAutoLoginEnumStringValues Enumerates the set of values in String for CryptoAssessmentWalletSummaryAutoLoginEnum
func GetCryptoAssessmentWalletSummaryAutoLoginEnumStringValues() []string {
	return []string{
		"ENABLED",
		"DISABLED",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoAssessmentWalletSummaryAutoLoginEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentWalletSummaryAutoLoginEnum(val string) (CryptoAssessmentWalletSummaryAutoLoginEnum, bool) {
	enum, ok := mappingCryptoAssessmentWalletSummaryAutoLoginEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
