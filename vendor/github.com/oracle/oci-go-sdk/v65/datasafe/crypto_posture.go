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

// CryptoPosture Cryptographic posture details captured by the assessment.
type CryptoPosture struct {
	Tls *CryptoTlsPosture `mandatory:"false" json:"tls"`

	Nne *CryptoNnePosture `mandatory:"false" json:"nne"`

	Tde *CryptoTdePosture `mandatory:"false" json:"tde"`

	// Network encryption details.
	NetworkEncryption []string `mandatory:"false" json:"networkEncryption"`

	// Overall FIPS status for the assessment when the target uses common FIPS configuration.
	FipsStatus CryptoFeatureStatusEnum `mandatory:"false" json:"fipsStatus,omitempty"`

	// Common FIPS mode configured for the assessment when the target uses common FIPS configuration.
	FipsModeConfigured CryptoPostureFipsModeConfiguredEnum `mandatory:"false" json:"fipsModeConfigured,omitempty"`

	// Backup encryption status observed for the assessment.
	BackupStatus *string `mandatory:"false" json:"backupStatus"`

	// Number of encrypted backup pieces.
	EncryptedBackupPiecesCount *int `mandatory:"false" json:"encryptedBackupPiecesCount"`

	// Number of unencrypted backup pieces.
	UnencryptedBackupPiecesCount *int `mandatory:"false" json:"unencryptedBackupPiecesCount"`
}

func (m CryptoPosture) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoPosture) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCryptoFeatureStatusEnum(string(m.FipsStatus)); !ok && m.FipsStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FipsStatus: %s. Supported values are: %s.", m.FipsStatus, strings.Join(GetCryptoFeatureStatusEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoPostureFipsModeConfiguredEnum(string(m.FipsModeConfigured)); !ok && m.FipsModeConfigured != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for FipsModeConfigured: %s. Supported values are: %s.", m.FipsModeConfigured, strings.Join(GetCryptoPostureFipsModeConfiguredEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoPostureFipsModeConfiguredEnum Enum with underlying type: string
type CryptoPostureFipsModeConfiguredEnum string

// Set of constants representing the allowable values for CryptoPostureFipsModeConfiguredEnum
const (
	CryptoPostureFipsModeConfiguredFips1402      CryptoPostureFipsModeConfiguredEnum = "FIPS_140_2"
	CryptoPostureFipsModeConfiguredFips1403      CryptoPostureFipsModeConfiguredEnum = "FIPS_140_3"
	CryptoPostureFipsModeConfiguredNotConfigured CryptoPostureFipsModeConfiguredEnum = "NOT_CONFIGURED"
	CryptoPostureFipsModeConfiguredNotApplicable CryptoPostureFipsModeConfiguredEnum = "NOT_APPLICABLE"
	CryptoPostureFipsModeConfiguredNotSupported  CryptoPostureFipsModeConfiguredEnum = "NOT_SUPPORTED"
)

var mappingCryptoPostureFipsModeConfiguredEnum = map[string]CryptoPostureFipsModeConfiguredEnum{
	"FIPS_140_2":     CryptoPostureFipsModeConfiguredFips1402,
	"FIPS_140_3":     CryptoPostureFipsModeConfiguredFips1403,
	"NOT_CONFIGURED": CryptoPostureFipsModeConfiguredNotConfigured,
	"NOT_APPLICABLE": CryptoPostureFipsModeConfiguredNotApplicable,
	"NOT_SUPPORTED":  CryptoPostureFipsModeConfiguredNotSupported,
}

var mappingCryptoPostureFipsModeConfiguredEnumLowerCase = map[string]CryptoPostureFipsModeConfiguredEnum{
	"fips_140_2":     CryptoPostureFipsModeConfiguredFips1402,
	"fips_140_3":     CryptoPostureFipsModeConfiguredFips1403,
	"not_configured": CryptoPostureFipsModeConfiguredNotConfigured,
	"not_applicable": CryptoPostureFipsModeConfiguredNotApplicable,
	"not_supported":  CryptoPostureFipsModeConfiguredNotSupported,
}

// GetCryptoPostureFipsModeConfiguredEnumValues Enumerates the set of values for CryptoPostureFipsModeConfiguredEnum
func GetCryptoPostureFipsModeConfiguredEnumValues() []CryptoPostureFipsModeConfiguredEnum {
	values := make([]CryptoPostureFipsModeConfiguredEnum, 0)
	for _, v := range mappingCryptoPostureFipsModeConfiguredEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoPostureFipsModeConfiguredEnumStringValues Enumerates the set of values in String for CryptoPostureFipsModeConfiguredEnum
func GetCryptoPostureFipsModeConfiguredEnumStringValues() []string {
	return []string{
		"FIPS_140_2",
		"FIPS_140_3",
		"NOT_CONFIGURED",
		"NOT_APPLICABLE",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoPostureFipsModeConfiguredEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoPostureFipsModeConfiguredEnum(val string) (CryptoPostureFipsModeConfiguredEnum, bool) {
	enum, ok := mappingCryptoPostureFipsModeConfiguredEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
