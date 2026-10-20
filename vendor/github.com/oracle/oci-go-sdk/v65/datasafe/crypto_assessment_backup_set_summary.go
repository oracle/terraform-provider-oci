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

// CryptoAssessmentBackupSetSummary Summary of one backup set observed for a crypto assessment.
type CryptoAssessmentBackupSetSummary struct {

	// OCID of the crypto assessment that discovered the backup set.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// OCID of the target database associated with the backup set.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Backup set key.
	BackupSetKey *string `mandatory:"true" json:"backupSetKey"`

	// Backup set stamp.
	SetStamp *string `mandatory:"true" json:"setStamp"`

	// Backup type observed for the set.
	BackupType CryptoAssessmentBackupSetSummaryBackupTypeEnum `mandatory:"true" json:"backupType"`

	// Current status of the backup set.
	Status CryptoAssessmentBackupSetSummaryStatusEnum `mandatory:"true" json:"status"`

	// Number of backup pieces in the set.
	BackupPieces *int `mandatory:"true" json:"backupPieces"`

	// Indicates whether the backup set is encrypted.
	IsEncrypted *bool `mandatory:"true" json:"isEncrypted"`

	// Backup set size in gigabytes.
	SizeInGBs *float64 `mandatory:"true" json:"sizeInGBs"`

	// The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Encryption algorithm observed for the backup set when encryption is enabled.
	AlgorithmObserved *string `mandatory:"false" json:"algorithmObserved"`

	// Cipher mode observed for the backup set when encryption is enabled.
	CipherModeObserved *string `mandatory:"false" json:"cipherModeObserved"`

	// Indicates whether the backup set is compressed.
	IsCompressed *bool `mandatory:"false" json:"isCompressed"`

	// Backup set creation time in RFC3339 format.
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`
}

func (m CryptoAssessmentBackupSetSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentBackupSetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentBackupSetSummaryBackupTypeEnum(string(m.BackupType)); !ok && m.BackupType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BackupType: %s. Supported values are: %s.", m.BackupType, strings.Join(GetCryptoAssessmentBackupSetSummaryBackupTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentBackupSetSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoAssessmentBackupSetSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentBackupSetSummaryBackupTypeEnum Enum with underlying type: string
type CryptoAssessmentBackupSetSummaryBackupTypeEnum string

// Set of constants representing the allowable values for CryptoAssessmentBackupSetSummaryBackupTypeEnum
const (
	CryptoAssessmentBackupSetSummaryBackupTypeDatafile     CryptoAssessmentBackupSetSummaryBackupTypeEnum = "DATAFILE"
	CryptoAssessmentBackupSetSummaryBackupTypeArchived     CryptoAssessmentBackupSetSummaryBackupTypeEnum = "ARCHIVED"
	CryptoAssessmentBackupSetSummaryBackupTypeIncremental  CryptoAssessmentBackupSetSummaryBackupTypeEnum = "INCREMENTAL"
	CryptoAssessmentBackupSetSummaryBackupTypeNotSupported CryptoAssessmentBackupSetSummaryBackupTypeEnum = "NOT_SUPPORTED"
)

var mappingCryptoAssessmentBackupSetSummaryBackupTypeEnum = map[string]CryptoAssessmentBackupSetSummaryBackupTypeEnum{
	"DATAFILE":      CryptoAssessmentBackupSetSummaryBackupTypeDatafile,
	"ARCHIVED":      CryptoAssessmentBackupSetSummaryBackupTypeArchived,
	"INCREMENTAL":   CryptoAssessmentBackupSetSummaryBackupTypeIncremental,
	"NOT_SUPPORTED": CryptoAssessmentBackupSetSummaryBackupTypeNotSupported,
}

var mappingCryptoAssessmentBackupSetSummaryBackupTypeEnumLowerCase = map[string]CryptoAssessmentBackupSetSummaryBackupTypeEnum{
	"datafile":      CryptoAssessmentBackupSetSummaryBackupTypeDatafile,
	"archived":      CryptoAssessmentBackupSetSummaryBackupTypeArchived,
	"incremental":   CryptoAssessmentBackupSetSummaryBackupTypeIncremental,
	"not_supported": CryptoAssessmentBackupSetSummaryBackupTypeNotSupported,
}

// GetCryptoAssessmentBackupSetSummaryBackupTypeEnumValues Enumerates the set of values for CryptoAssessmentBackupSetSummaryBackupTypeEnum
func GetCryptoAssessmentBackupSetSummaryBackupTypeEnumValues() []CryptoAssessmentBackupSetSummaryBackupTypeEnum {
	values := make([]CryptoAssessmentBackupSetSummaryBackupTypeEnum, 0)
	for _, v := range mappingCryptoAssessmentBackupSetSummaryBackupTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentBackupSetSummaryBackupTypeEnumStringValues Enumerates the set of values in String for CryptoAssessmentBackupSetSummaryBackupTypeEnum
func GetCryptoAssessmentBackupSetSummaryBackupTypeEnumStringValues() []string {
	return []string{
		"DATAFILE",
		"ARCHIVED",
		"INCREMENTAL",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoAssessmentBackupSetSummaryBackupTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentBackupSetSummaryBackupTypeEnum(val string) (CryptoAssessmentBackupSetSummaryBackupTypeEnum, bool) {
	enum, ok := mappingCryptoAssessmentBackupSetSummaryBackupTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentBackupSetSummaryStatusEnum Enum with underlying type: string
type CryptoAssessmentBackupSetSummaryStatusEnum string

// Set of constants representing the allowable values for CryptoAssessmentBackupSetSummaryStatusEnum
const (
	CryptoAssessmentBackupSetSummaryStatusAvailable    CryptoAssessmentBackupSetSummaryStatusEnum = "AVAILABLE"
	CryptoAssessmentBackupSetSummaryStatusExpired      CryptoAssessmentBackupSetSummaryStatusEnum = "EXPIRED"
	CryptoAssessmentBackupSetSummaryStatusDeleted      CryptoAssessmentBackupSetSummaryStatusEnum = "DELETED"
	CryptoAssessmentBackupSetSummaryStatusNotSupported CryptoAssessmentBackupSetSummaryStatusEnum = "NOT_SUPPORTED"
)

var mappingCryptoAssessmentBackupSetSummaryStatusEnum = map[string]CryptoAssessmentBackupSetSummaryStatusEnum{
	"AVAILABLE":     CryptoAssessmentBackupSetSummaryStatusAvailable,
	"EXPIRED":       CryptoAssessmentBackupSetSummaryStatusExpired,
	"DELETED":       CryptoAssessmentBackupSetSummaryStatusDeleted,
	"NOT_SUPPORTED": CryptoAssessmentBackupSetSummaryStatusNotSupported,
}

var mappingCryptoAssessmentBackupSetSummaryStatusEnumLowerCase = map[string]CryptoAssessmentBackupSetSummaryStatusEnum{
	"available":     CryptoAssessmentBackupSetSummaryStatusAvailable,
	"expired":       CryptoAssessmentBackupSetSummaryStatusExpired,
	"deleted":       CryptoAssessmentBackupSetSummaryStatusDeleted,
	"not_supported": CryptoAssessmentBackupSetSummaryStatusNotSupported,
}

// GetCryptoAssessmentBackupSetSummaryStatusEnumValues Enumerates the set of values for CryptoAssessmentBackupSetSummaryStatusEnum
func GetCryptoAssessmentBackupSetSummaryStatusEnumValues() []CryptoAssessmentBackupSetSummaryStatusEnum {
	values := make([]CryptoAssessmentBackupSetSummaryStatusEnum, 0)
	for _, v := range mappingCryptoAssessmentBackupSetSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentBackupSetSummaryStatusEnumStringValues Enumerates the set of values in String for CryptoAssessmentBackupSetSummaryStatusEnum
func GetCryptoAssessmentBackupSetSummaryStatusEnumStringValues() []string {
	return []string{
		"AVAILABLE",
		"EXPIRED",
		"DELETED",
		"NOT_SUPPORTED",
	}
}

// GetMappingCryptoAssessmentBackupSetSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentBackupSetSummaryStatusEnum(val string) (CryptoAssessmentBackupSetSummaryStatusEnum, bool) {
	enum, ok := mappingCryptoAssessmentBackupSetSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
