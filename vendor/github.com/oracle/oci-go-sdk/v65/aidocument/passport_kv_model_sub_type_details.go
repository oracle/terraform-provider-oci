// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Document Understanding API
//
// Document AI helps customers perform various analysis on their documents. If a customer has lots of documents, they can process them in batch using asynchronous API endpoints.
//

package aidocument

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PassportKvModelSubTypeDetails Passport KV model sub type details
type PassportKvModelSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType PassportKvModelSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType PassportKvModelSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m PassportKvModelSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PassportKvModelSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPassportKvModelSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetPassportKvModelSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPassportKvModelSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetPassportKvModelSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PassportKvModelSubTypeDetailsModelTypeEnum Enum with underlying type: string
type PassportKvModelSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for PassportKvModelSubTypeDetailsModelTypeEnum
const (
	PassportKvModelSubTypeDetailsModelTypeKeyValueExtraction         PassportKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	PassportKvModelSubTypeDetailsModelTypeDocumentElementsExtraction PassportKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingPassportKvModelSubTypeDetailsModelTypeEnum = map[string]PassportKvModelSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         PassportKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": PassportKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingPassportKvModelSubTypeDetailsModelTypeEnumLowerCase = map[string]PassportKvModelSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         PassportKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": PassportKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetPassportKvModelSubTypeDetailsModelTypeEnumValues Enumerates the set of values for PassportKvModelSubTypeDetailsModelTypeEnum
func GetPassportKvModelSubTypeDetailsModelTypeEnumValues() []PassportKvModelSubTypeDetailsModelTypeEnum {
	values := make([]PassportKvModelSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingPassportKvModelSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPassportKvModelSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for PassportKvModelSubTypeDetailsModelTypeEnum
func GetPassportKvModelSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingPassportKvModelSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPassportKvModelSubTypeDetailsModelTypeEnum(val string) (PassportKvModelSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingPassportKvModelSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PassportKvModelSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type PassportKvModelSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for PassportKvModelSubTypeDetailsModelSubTypeEnum
const (
	PassportKvModelSubTypeDetailsModelSubTypeReceipt           PassportKvModelSubTypeDetailsModelSubTypeEnum = "RECEIPT"
	PassportKvModelSubTypeDetailsModelSubTypeInvoice           PassportKvModelSubTypeDetailsModelSubTypeEnum = "INVOICE"
	PassportKvModelSubTypeDetailsModelSubTypePassport          PassportKvModelSubTypeDetailsModelSubTypeEnum = "PASSPORT"
	PassportKvModelSubTypeDetailsModelSubTypeDriverLicense     PassportKvModelSubTypeDetailsModelSubTypeEnum = "DRIVER_LICENSE"
	PassportKvModelSubTypeDetailsModelSubTypeHealthInsuranceId PassportKvModelSubTypeDetailsModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingPassportKvModelSubTypeDetailsModelSubTypeEnum = map[string]PassportKvModelSubTypeDetailsModelSubTypeEnum{
	"RECEIPT":             PassportKvModelSubTypeDetailsModelSubTypeReceipt,
	"INVOICE":             PassportKvModelSubTypeDetailsModelSubTypeInvoice,
	"PASSPORT":            PassportKvModelSubTypeDetailsModelSubTypePassport,
	"DRIVER_LICENSE":      PassportKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": PassportKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

var mappingPassportKvModelSubTypeDetailsModelSubTypeEnumLowerCase = map[string]PassportKvModelSubTypeDetailsModelSubTypeEnum{
	"receipt":             PassportKvModelSubTypeDetailsModelSubTypeReceipt,
	"invoice":             PassportKvModelSubTypeDetailsModelSubTypeInvoice,
	"passport":            PassportKvModelSubTypeDetailsModelSubTypePassport,
	"driver_license":      PassportKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"health_insurance_id": PassportKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

// GetPassportKvModelSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for PassportKvModelSubTypeDetailsModelSubTypeEnum
func GetPassportKvModelSubTypeDetailsModelSubTypeEnumValues() []PassportKvModelSubTypeDetailsModelSubTypeEnum {
	values := make([]PassportKvModelSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingPassportKvModelSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPassportKvModelSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for PassportKvModelSubTypeDetailsModelSubTypeEnum
func GetPassportKvModelSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingPassportKvModelSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPassportKvModelSubTypeDetailsModelSubTypeEnum(val string) (PassportKvModelSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingPassportKvModelSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
