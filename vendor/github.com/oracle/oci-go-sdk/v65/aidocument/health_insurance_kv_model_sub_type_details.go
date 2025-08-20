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

// HealthInsuranceKvModelSubTypeDetails Health insurance card KV model sub type details
type HealthInsuranceKvModelSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType HealthInsuranceKvModelSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m HealthInsuranceKvModelSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m HealthInsuranceKvModelSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetHealthInsuranceKvModelSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// HealthInsuranceKvModelSubTypeDetailsModelTypeEnum Enum with underlying type: string
type HealthInsuranceKvModelSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for HealthInsuranceKvModelSubTypeDetailsModelTypeEnum
const (
	HealthInsuranceKvModelSubTypeDetailsModelTypeKeyValueExtraction         HealthInsuranceKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	HealthInsuranceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction HealthInsuranceKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnum = map[string]HealthInsuranceKvModelSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         HealthInsuranceKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": HealthInsuranceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnumLowerCase = map[string]HealthInsuranceKvModelSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         HealthInsuranceKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": HealthInsuranceKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetHealthInsuranceKvModelSubTypeDetailsModelTypeEnumValues Enumerates the set of values for HealthInsuranceKvModelSubTypeDetailsModelTypeEnum
func GetHealthInsuranceKvModelSubTypeDetailsModelTypeEnumValues() []HealthInsuranceKvModelSubTypeDetailsModelTypeEnum {
	values := make([]HealthInsuranceKvModelSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthInsuranceKvModelSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for HealthInsuranceKvModelSubTypeDetailsModelTypeEnum
func GetHealthInsuranceKvModelSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnum(val string) (HealthInsuranceKvModelSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingHealthInsuranceKvModelSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum
const (
	HealthInsuranceKvModelSubTypeDetailsModelSubTypeReceipt           HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = "RECEIPT"
	HealthInsuranceKvModelSubTypeDetailsModelSubTypeInvoice           HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = "INVOICE"
	HealthInsuranceKvModelSubTypeDetailsModelSubTypePassport          HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = "PASSPORT"
	HealthInsuranceKvModelSubTypeDetailsModelSubTypeDriverLicense     HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = "DRIVER_LICENSE"
	HealthInsuranceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum = map[string]HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum{
	"RECEIPT":             HealthInsuranceKvModelSubTypeDetailsModelSubTypeReceipt,
	"INVOICE":             HealthInsuranceKvModelSubTypeDetailsModelSubTypeInvoice,
	"PASSPORT":            HealthInsuranceKvModelSubTypeDetailsModelSubTypePassport,
	"DRIVER_LICENSE":      HealthInsuranceKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": HealthInsuranceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

var mappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumLowerCase = map[string]HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum{
	"receipt":             HealthInsuranceKvModelSubTypeDetailsModelSubTypeReceipt,
	"invoice":             HealthInsuranceKvModelSubTypeDetailsModelSubTypeInvoice,
	"passport":            HealthInsuranceKvModelSubTypeDetailsModelSubTypePassport,
	"driver_license":      HealthInsuranceKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"health_insurance_id": HealthInsuranceKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

// GetHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum
func GetHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumValues() []HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum {
	values := make([]HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum
func GetHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum(val string) (HealthInsuranceKvModelSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingHealthInsuranceKvModelSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
