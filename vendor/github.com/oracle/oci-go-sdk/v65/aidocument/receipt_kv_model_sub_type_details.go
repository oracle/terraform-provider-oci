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

// ReceiptKvModelSubTypeDetails Receipt KV model sub type details
type ReceiptKvModelSubTypeDetails struct {

	// Sub type model based on the model type.
	// The allowed values are:
	// - `PRE_TRAINED_KEY_VALUE_EXTRACTION`
	// - `PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION`
	ModelType ReceiptKvModelSubTypeDetailsModelTypeEnum `mandatory:"true" json:"modelType"`

	// The model sub type for KEY_VALUE_EXTRACTION
	// The allowed values are:
	// - `RECEIPT`
	// - `INVOICE`
	// - `PASSPORT`
	// - `DRIVER_LICENSE`
	// - `HEALTH_INSURANCE_ID`
	ModelSubType ReceiptKvModelSubTypeDetailsModelSubTypeEnum `mandatory:"true" json:"modelSubType"`
}

func (m ReceiptKvModelSubTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ReceiptKvModelSubTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingReceiptKvModelSubTypeDetailsModelTypeEnum(string(m.ModelType)); !ok && m.ModelType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelType: %s. Supported values are: %s.", m.ModelType, strings.Join(GetReceiptKvModelSubTypeDetailsModelTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingReceiptKvModelSubTypeDetailsModelSubTypeEnum(string(m.ModelSubType)); !ok && m.ModelSubType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ModelSubType: %s. Supported values are: %s.", m.ModelSubType, strings.Join(GetReceiptKvModelSubTypeDetailsModelSubTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ReceiptKvModelSubTypeDetailsModelTypeEnum Enum with underlying type: string
type ReceiptKvModelSubTypeDetailsModelTypeEnum string

// Set of constants representing the allowable values for ReceiptKvModelSubTypeDetailsModelTypeEnum
const (
	ReceiptKvModelSubTypeDetailsModelTypeKeyValueExtraction         ReceiptKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_KEY_VALUE_EXTRACTION"
	ReceiptKvModelSubTypeDetailsModelTypeDocumentElementsExtraction ReceiptKvModelSubTypeDetailsModelTypeEnum = "PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION"
)

var mappingReceiptKvModelSubTypeDetailsModelTypeEnum = map[string]ReceiptKvModelSubTypeDetailsModelTypeEnum{
	"PRE_TRAINED_KEY_VALUE_EXTRACTION":         ReceiptKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION": ReceiptKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

var mappingReceiptKvModelSubTypeDetailsModelTypeEnumLowerCase = map[string]ReceiptKvModelSubTypeDetailsModelTypeEnum{
	"pre_trained_key_value_extraction":         ReceiptKvModelSubTypeDetailsModelTypeKeyValueExtraction,
	"pre_trained_document_elements_extraction": ReceiptKvModelSubTypeDetailsModelTypeDocumentElementsExtraction,
}

// GetReceiptKvModelSubTypeDetailsModelTypeEnumValues Enumerates the set of values for ReceiptKvModelSubTypeDetailsModelTypeEnum
func GetReceiptKvModelSubTypeDetailsModelTypeEnumValues() []ReceiptKvModelSubTypeDetailsModelTypeEnum {
	values := make([]ReceiptKvModelSubTypeDetailsModelTypeEnum, 0)
	for _, v := range mappingReceiptKvModelSubTypeDetailsModelTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReceiptKvModelSubTypeDetailsModelTypeEnumStringValues Enumerates the set of values in String for ReceiptKvModelSubTypeDetailsModelTypeEnum
func GetReceiptKvModelSubTypeDetailsModelTypeEnumStringValues() []string {
	return []string{
		"PRE_TRAINED_KEY_VALUE_EXTRACTION",
		"PRE_TRAINED_DOCUMENT_ELEMENTS_EXTRACTION",
	}
}

// GetMappingReceiptKvModelSubTypeDetailsModelTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReceiptKvModelSubTypeDetailsModelTypeEnum(val string) (ReceiptKvModelSubTypeDetailsModelTypeEnum, bool) {
	enum, ok := mappingReceiptKvModelSubTypeDetailsModelTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// ReceiptKvModelSubTypeDetailsModelSubTypeEnum Enum with underlying type: string
type ReceiptKvModelSubTypeDetailsModelSubTypeEnum string

// Set of constants representing the allowable values for ReceiptKvModelSubTypeDetailsModelSubTypeEnum
const (
	ReceiptKvModelSubTypeDetailsModelSubTypeReceipt           ReceiptKvModelSubTypeDetailsModelSubTypeEnum = "RECEIPT"
	ReceiptKvModelSubTypeDetailsModelSubTypeInvoice           ReceiptKvModelSubTypeDetailsModelSubTypeEnum = "INVOICE"
	ReceiptKvModelSubTypeDetailsModelSubTypePassport          ReceiptKvModelSubTypeDetailsModelSubTypeEnum = "PASSPORT"
	ReceiptKvModelSubTypeDetailsModelSubTypeDriverLicense     ReceiptKvModelSubTypeDetailsModelSubTypeEnum = "DRIVER_LICENSE"
	ReceiptKvModelSubTypeDetailsModelSubTypeHealthInsuranceId ReceiptKvModelSubTypeDetailsModelSubTypeEnum = "HEALTH_INSURANCE_ID"
)

var mappingReceiptKvModelSubTypeDetailsModelSubTypeEnum = map[string]ReceiptKvModelSubTypeDetailsModelSubTypeEnum{
	"RECEIPT":             ReceiptKvModelSubTypeDetailsModelSubTypeReceipt,
	"INVOICE":             ReceiptKvModelSubTypeDetailsModelSubTypeInvoice,
	"PASSPORT":            ReceiptKvModelSubTypeDetailsModelSubTypePassport,
	"DRIVER_LICENSE":      ReceiptKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"HEALTH_INSURANCE_ID": ReceiptKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

var mappingReceiptKvModelSubTypeDetailsModelSubTypeEnumLowerCase = map[string]ReceiptKvModelSubTypeDetailsModelSubTypeEnum{
	"receipt":             ReceiptKvModelSubTypeDetailsModelSubTypeReceipt,
	"invoice":             ReceiptKvModelSubTypeDetailsModelSubTypeInvoice,
	"passport":            ReceiptKvModelSubTypeDetailsModelSubTypePassport,
	"driver_license":      ReceiptKvModelSubTypeDetailsModelSubTypeDriverLicense,
	"health_insurance_id": ReceiptKvModelSubTypeDetailsModelSubTypeHealthInsuranceId,
}

// GetReceiptKvModelSubTypeDetailsModelSubTypeEnumValues Enumerates the set of values for ReceiptKvModelSubTypeDetailsModelSubTypeEnum
func GetReceiptKvModelSubTypeDetailsModelSubTypeEnumValues() []ReceiptKvModelSubTypeDetailsModelSubTypeEnum {
	values := make([]ReceiptKvModelSubTypeDetailsModelSubTypeEnum, 0)
	for _, v := range mappingReceiptKvModelSubTypeDetailsModelSubTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetReceiptKvModelSubTypeDetailsModelSubTypeEnumStringValues Enumerates the set of values in String for ReceiptKvModelSubTypeDetailsModelSubTypeEnum
func GetReceiptKvModelSubTypeDetailsModelSubTypeEnumStringValues() []string {
	return []string{
		"RECEIPT",
		"INVOICE",
		"PASSPORT",
		"DRIVER_LICENSE",
		"HEALTH_INSURANCE_ID",
	}
}

// GetMappingReceiptKvModelSubTypeDetailsModelSubTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingReceiptKvModelSubTypeDetailsModelSubTypeEnum(val string) (ReceiptKvModelSubTypeDetailsModelSubTypeEnum, bool) {
	enum, ok := mappingReceiptKvModelSubTypeDetailsModelSubTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
