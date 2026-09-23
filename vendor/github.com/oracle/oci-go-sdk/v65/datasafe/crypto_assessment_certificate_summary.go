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

// CryptoAssessmentCertificateSummary Summary of one certificate discovered across targets in a compartment.
type CryptoAssessmentCertificateSummary struct {

	// OCID of the crypto assessment that discovered the certificate.
	AssessmentId *string `mandatory:"true" json:"assessmentId"`

	// Type of the crypto assessment that discovered the certificate.
	AssessmentType CryptoAssessmentTypeEnum `mandatory:"true" json:"assessmentType"`

	// OCID of the compartment that contains the certificate row.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// OCID of the target database associated with the certificate.
	TargetId *string `mandatory:"true" json:"targetId"`

	// Type of certificate.
	CertificateType CryptoAssessmentCertificateSummaryCertificateTypeEnum `mandatory:"true" json:"certificateType"`

	// Subject of the certificate.
	Subject *string `mandatory:"true" json:"subject"`

	// Issuer of the certificate.
	Issuer *string `mandatory:"true" json:"issuer"`

	// Certificate validity status.
	Status CryptoAssessmentCertificateSummaryStatusEnum `mandatory:"true" json:"status"`

	// Certificate serial number.
	SerialNumber *string `mandatory:"true" json:"serialNumber"`

	// Certificate validity start time in RFC3339 format.
	TimeValidFrom *common.SDKTime `mandatory:"true" json:"timeValidFrom"`

	// Certificate validity end time in RFC3339 format.
	TimeValidUntil *common.SDKTime `mandatory:"true" json:"timeValidUntil"`

	// Public key type and size.
	PublicKeyType *string `mandatory:"true" json:"publicKeyType"`

	// Signature algorithm used by the certificate.
	SignatureAlgorithm *string `mandatory:"true" json:"signatureAlgorithm"`

	// The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	TimeLastAssessed *common.SDKTime `mandatory:"false" json:"timeLastAssessed"`

	// Wallet location where the certificate was discovered, if available.
	WalletLocation *string `mandatory:"false" json:"walletLocation"`

	// Age of the certificate in whole days, calculated from timeValidFrom using the current UTC date.
	Age *int64 `mandatory:"false" json:"age"`

	// Expiry bucket populated for the certificate when the assessment runs. Supported values are 0_15, 15_30, 30_60, 60_90, and 90_PLUS.
	ExpiryBucket *string `mandatory:"false" json:"expiryBucket"`

	// Number of whole days until certificate expiration, calculated from timeValidUntil. Negative values indicate already expired certificates.
	DaysToExpiry *int `mandatory:"false" json:"daysToExpiry"`
}

func (m CryptoAssessmentCertificateSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CryptoAssessmentCertificateSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingCryptoAssessmentTypeEnum(string(m.AssessmentType)); !ok && m.AssessmentType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AssessmentType: %s. Supported values are: %s.", m.AssessmentType, strings.Join(GetCryptoAssessmentTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentCertificateSummaryCertificateTypeEnum(string(m.CertificateType)); !ok && m.CertificateType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CertificateType: %s. Supported values are: %s.", m.CertificateType, strings.Join(GetCryptoAssessmentCertificateSummaryCertificateTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingCryptoAssessmentCertificateSummaryStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetCryptoAssessmentCertificateSummaryStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CryptoAssessmentCertificateSummaryCertificateTypeEnum Enum with underlying type: string
type CryptoAssessmentCertificateSummaryCertificateTypeEnum string

// Set of constants representing the allowable values for CryptoAssessmentCertificateSummaryCertificateTypeEnum
const (
	CryptoAssessmentCertificateSummaryCertificateTypeServer  CryptoAssessmentCertificateSummaryCertificateTypeEnum = "SERVER"
	CryptoAssessmentCertificateSummaryCertificateTypeUser    CryptoAssessmentCertificateSummaryCertificateTypeEnum = "USER"
	CryptoAssessmentCertificateSummaryCertificateTypeTrusted CryptoAssessmentCertificateSummaryCertificateTypeEnum = "TRUSTED"
)

var mappingCryptoAssessmentCertificateSummaryCertificateTypeEnum = map[string]CryptoAssessmentCertificateSummaryCertificateTypeEnum{
	"SERVER":  CryptoAssessmentCertificateSummaryCertificateTypeServer,
	"USER":    CryptoAssessmentCertificateSummaryCertificateTypeUser,
	"TRUSTED": CryptoAssessmentCertificateSummaryCertificateTypeTrusted,
}

var mappingCryptoAssessmentCertificateSummaryCertificateTypeEnumLowerCase = map[string]CryptoAssessmentCertificateSummaryCertificateTypeEnum{
	"server":  CryptoAssessmentCertificateSummaryCertificateTypeServer,
	"user":    CryptoAssessmentCertificateSummaryCertificateTypeUser,
	"trusted": CryptoAssessmentCertificateSummaryCertificateTypeTrusted,
}

// GetCryptoAssessmentCertificateSummaryCertificateTypeEnumValues Enumerates the set of values for CryptoAssessmentCertificateSummaryCertificateTypeEnum
func GetCryptoAssessmentCertificateSummaryCertificateTypeEnumValues() []CryptoAssessmentCertificateSummaryCertificateTypeEnum {
	values := make([]CryptoAssessmentCertificateSummaryCertificateTypeEnum, 0)
	for _, v := range mappingCryptoAssessmentCertificateSummaryCertificateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentCertificateSummaryCertificateTypeEnumStringValues Enumerates the set of values in String for CryptoAssessmentCertificateSummaryCertificateTypeEnum
func GetCryptoAssessmentCertificateSummaryCertificateTypeEnumStringValues() []string {
	return []string{
		"SERVER",
		"USER",
		"TRUSTED",
	}
}

// GetMappingCryptoAssessmentCertificateSummaryCertificateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentCertificateSummaryCertificateTypeEnum(val string) (CryptoAssessmentCertificateSummaryCertificateTypeEnum, bool) {
	enum, ok := mappingCryptoAssessmentCertificateSummaryCertificateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// CryptoAssessmentCertificateSummaryStatusEnum Enum with underlying type: string
type CryptoAssessmentCertificateSummaryStatusEnum string

// Set of constants representing the allowable values for CryptoAssessmentCertificateSummaryStatusEnum
const (
	CryptoAssessmentCertificateSummaryStatusValid        CryptoAssessmentCertificateSummaryStatusEnum = "VALID"
	CryptoAssessmentCertificateSummaryStatusExpiringSoon CryptoAssessmentCertificateSummaryStatusEnum = "EXPIRING_SOON"
	CryptoAssessmentCertificateSummaryStatusExpired      CryptoAssessmentCertificateSummaryStatusEnum = "EXPIRED"
	CryptoAssessmentCertificateSummaryStatusInvalid      CryptoAssessmentCertificateSummaryStatusEnum = "INVALID"
	CryptoAssessmentCertificateSummaryStatusInUse        CryptoAssessmentCertificateSummaryStatusEnum = "IN_USE"
)

var mappingCryptoAssessmentCertificateSummaryStatusEnum = map[string]CryptoAssessmentCertificateSummaryStatusEnum{
	"VALID":         CryptoAssessmentCertificateSummaryStatusValid,
	"EXPIRING_SOON": CryptoAssessmentCertificateSummaryStatusExpiringSoon,
	"EXPIRED":       CryptoAssessmentCertificateSummaryStatusExpired,
	"INVALID":       CryptoAssessmentCertificateSummaryStatusInvalid,
	"IN_USE":        CryptoAssessmentCertificateSummaryStatusInUse,
}

var mappingCryptoAssessmentCertificateSummaryStatusEnumLowerCase = map[string]CryptoAssessmentCertificateSummaryStatusEnum{
	"valid":         CryptoAssessmentCertificateSummaryStatusValid,
	"expiring_soon": CryptoAssessmentCertificateSummaryStatusExpiringSoon,
	"expired":       CryptoAssessmentCertificateSummaryStatusExpired,
	"invalid":       CryptoAssessmentCertificateSummaryStatusInvalid,
	"in_use":        CryptoAssessmentCertificateSummaryStatusInUse,
}

// GetCryptoAssessmentCertificateSummaryStatusEnumValues Enumerates the set of values for CryptoAssessmentCertificateSummaryStatusEnum
func GetCryptoAssessmentCertificateSummaryStatusEnumValues() []CryptoAssessmentCertificateSummaryStatusEnum {
	values := make([]CryptoAssessmentCertificateSummaryStatusEnum, 0)
	for _, v := range mappingCryptoAssessmentCertificateSummaryStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetCryptoAssessmentCertificateSummaryStatusEnumStringValues Enumerates the set of values in String for CryptoAssessmentCertificateSummaryStatusEnum
func GetCryptoAssessmentCertificateSummaryStatusEnumStringValues() []string {
	return []string{
		"VALID",
		"EXPIRING_SOON",
		"EXPIRED",
		"INVALID",
		"IN_USE",
	}
}

// GetMappingCryptoAssessmentCertificateSummaryStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCryptoAssessmentCertificateSummaryStatusEnum(val string) (CryptoAssessmentCertificateSummaryStatusEnum, bool) {
	enum, ok := mappingCryptoAssessmentCertificateSummaryStatusEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
